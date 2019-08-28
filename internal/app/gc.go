package collector

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

// NOTICE: For generic and decoupling purpose, you might always want to inject these from the
// environment variables or some kinds of configuration yaml.
// If any of them are not properly provided, fallback to default.
const (
	// TargetServer is assumed localhost
	TargetServer = "127.0.0.1"
	// AccessToken place your own(or organization) github Personal Access Token
	// NOTICE: Make sure your are not pushing your token onto github if your repository is public
	AccessToken = "8992518d8cda5290ba387739837588662d6806e4"
	// SourceOwner github repo owner
	SourceOwner = "lmchih"
	// SourceRepo github repo name
	SourceRepo = "server-collector"
	// SourceBranch repo target branch (actually not used)
	SourceBranch = "master"
	// UnusedDays the expiration days before the server going to be turn off
	UnusedDays = 3
)

// YamlConf for the yaml format configuration file
type YamlConf struct {
	Version      string `yaml:"version"`
	AccessToken  string `yaml:"accessToken"`
	ServerIP     string `yaml:"serverIP"`
	SourceOwner  string `yaml:"sourceOwner"`
	SourceRepo   string `yaml:"sourceRepo"`
	SourceBranch string `yaml:"sourceBranch"`
	UnusedDays   int64  `yaml:"unusedDays"`
}

// EnvVars for the configuration passed from environment variables
type EnvVars struct {
	targetServer string
	accessToken  string
	sourceOwner  string
	sourceRepo   string
	sourceBranch string
	unusedDays   int64
}

// GetEnvs configuration is supposed to be injected from k8s ConfigMap or ducker run -e
func GetEnvs() (*EnvVars, error) {
	// default env from yaml
	log.Printf("TARGET_SERVER=%s\n", os.Getenv("TARGET_SERVER"))
	log.Printf("ACCESS_TOKEN=%s\n", os.Getenv("ACCESS_TOKEN"))
	log.Printf("SOURCE_OWNER=%s\n", os.Getenv("SOURCE_OWNER"))
	log.Printf("SOURCE_REPO=%s\n", os.Getenv("SOURCE_REPO"))
	log.Printf("SOURCE_BRANCH=%s\n", os.Getenv("SOURCE_BRANCH"))
	log.Printf("UNUSED_DAYS=%s\n", os.Getenv("UNUSED_DAYS"))

	envVars := EnvVars{}
	// if not getting any values from environment, fall back
	// to default constant values
	envVars.targetServer = os.Getenv("TARGET_SERVER")
	if envVars.targetServer == "" {
		envVars.targetServer = TargetServer
	}
	envVars.accessToken = os.Getenv("ACCESS_TOKEN")
	if envVars.accessToken == "" {
		envVars.accessToken = AccessToken
	}
	envVars.sourceOwner = os.Getenv("SOURCE_OWNER")
	if envVars.sourceOwner == "" {
		envVars.sourceOwner = SourceOwner
	}
	envVars.sourceRepo = os.Getenv("SOURCE_REPO")
	if envVars.sourceRepo == "" {
		envVars.sourceRepo = SourceRepo
	}
	envVars.sourceBranch = os.Getenv("SOURCE_BRANCH")
	if envVars.sourceBranch == "" {
		envVars.sourceBranch = SourceBranch
	}

	if os.Getenv("UNUSED_DAYS") == "" {
		envVars.unusedDays = UnusedDays
	} else {
		i, err := strconv.ParseInt(os.Getenv("UNUSED_DAYS"), 10, 64)
		if err != nil {
			panic(err)
		}
		envVars.unusedDays = int64(i)
	}

	return &envVars, nil
}

// Read configuration from the user editted yaml file
func (c *YamlConf) getConf() *YamlConf {
	yamlFile, err := ioutil.ReadFile("deployments/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v\n", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Printf("Unmarshal err: %v\n", err)
	}
	fmt.Printf("%v\n", c)
	return c
}

// BinaryEntry main handle method
func BinaryEntry() {
	fmt.Println("Start server-collector")

	var c YamlConf
	c.getConf()

	fmt.Println(c.Version)

	// Run once first right before the routine
	log.Println("Check Github last commit date")
	BinaryRunCheck(&c)

	// Start to run the routine job
	for range time.Tick(time.Duration(3) * time.Second) {
		log.Println("Check Github last commit date")
		BinaryRunCheck(&c)
	}
}

// ContainerEntry main handle method
func ContainerEntry() {
	envs, err := GetEnvs()
	log.Printf("envs: %v, err: %v", envs, err)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Check Github last commit date")
	ContainerRunCheck(envs)

	// Start to run the routine job
	for range time.Tick(time.Duration(30) * time.Second) {
		log.Println("Check Github last commit date")
		ContainerRunCheck(envs)
	}
}

// RunCheck check github commit days
func RunCheck(a interface{}) {
	// TODO:
}

// BinaryRunCheck check github commit status
func BinaryRunCheck(c *YamlConf) {
	// routinely check last commit date
	var days = lastCommitDays(c.AccessToken, c.SourceOwner, c.SourceRepo)
	if isUnused(days, c.UnusedDays) {
		shutdownCommand()
	}
}

// ContainerRunCheck check github commit status
func ContainerRunCheck(e *EnvVars) {
	fmt.Println("ContainerRunCheck")
	// fmt.Printf((e.accessToken))
	// days := lastCommitDays(e.accessToken, e.sourceOwner, e.sourceRepo)

	// if older than expiration, terminate the server.
	// if isUnused(days, e.unusedDays) {
	// 	terminate(e.targetServer)
	// }
	terminate(e.targetServer)
}

func isUnused(days int64, expiration int64) bool {
	if days >= expiration {
		return true
	}
	return false
}
