package collector

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/pborman/getopt"
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
	AccessToken = "495fe0558bb01d3635bfa9e93f5ebecc83f85387"
	// SourceOwner github repo owner
	SourceOwner = "lmchih"
	// SourceRepo github repo name
	SourceRepo = "server-collector"
	// SourceBranch repo target branch (actually not used)
	SourceBranch = "master"
	// CheckFrequency how often of running check (in seconds)
	CheckFrequency = 120
	// UnusedDays the expiration days before the server going to be turn off
	UnusedDays = 3
)

var (
	targetServer   = flag.String("target", "127.0.0.1", "The target server want to be monitored")
	accessToken    = flag.String("token", "", "Your(organization)'s Github Personal Access Token")
	sourceOwner    = flag.String("owner", "lmchih", "The Github repo owner")
	sourceRepo     = flag.String("repo", "server-collector", "The Github repo name")
	sourceBranch   = flag.String("branch", "master", "The Github repo branch")
	checkFrequency = flag.Int64("check-freq", 120, "How often of running check (in seconds)")
	unusedDays     = flag.Int64("unused-days", 3, "How long is considered unused")
)

// YamlConf for the yaml format configuration file
type YamlConf struct {
	Version        string `yaml:"version"`
	AccessToken    string `yaml:"accessToken"`
	ServerIP       string `yaml:"serverIP"`
	SourceOwner    string `yaml:"sourceOwner"`
	SourceRepo     string `yaml:"sourceRepo"`
	SourceBranch   string `yaml:"sourceBranch"`
	CheckFrequency int64  `yaml:"checkFrequency"`
	UnusedDays     int64  `yaml:"unusedDays"`
}

// Options the standard input option
type Options struct {
	IP         string
	Token      string
	Owner      string
	Repo       string
	Branch     string
	CheckFreq  int64
	UnusedDays int64
}

// Envars for the configuration passed from environment variables
type Envars struct {
	targetServer   string
	accessToken    string
	sourceOwner    string
	sourceRepo     string
	sourceBranch   string
	checkFrequency int64
	unusedDays     int64
}

func configToOption(yaml *YamlConf) *Options {
	v := reflect.ValueOf(*yaml)
	options := Options{
		IP:         v.FieldByName("ServerIP").String(),
		Token:      v.FieldByName("AccessToken").String(),
		Owner:      v.FieldByName("SourceOwner").String(),
		Repo:       v.FieldByName("SourceRepo").String(),
		Branch:     v.FieldByName("SourceBranch").String(),
		CheckFreq:  v.FieldByName("CheckFrequency").Int(),
		UnusedDays: v.FieldByName("UnusedDays").Int(),
	}

	return &options
}

// GetEnvs configuration is supposed to be injected from k8s ConfigMap or ducker run -e
func GetEnvs() (*Envars, error) {
	// default env from yaml
	log.Printf("TARGET_SERVER=%s\n", os.Getenv("TARGET_SERVER"))
	log.Printf("ACCESS_TOKEN=%s\n", os.Getenv("ACCESS_TOKEN"))
	log.Printf("SOURCE_OWNER=%s\n", os.Getenv("SOURCE_OWNER"))
	log.Printf("SOURCE_REPO=%s\n", os.Getenv("SOURCE_REPO"))
	log.Printf("SOURCE_BRANCH=%s\n", os.Getenv("SOURCE_BRANCH"))
	log.Printf("CHECK_FREQUENCY=%s\n", os.Getenv("CHECK_FREQUENCY"))
	log.Printf("UNUSED_DAYS=%s\n", os.Getenv("UNUSED_DAYS"))

	envars := Envars{}
	// if not getting any values from environment, fall back
	// to default constant values
	envars.targetServer = os.Getenv("TARGET_SERVER")
	if envars.targetServer == "" {
		envars.targetServer = TargetServer
	}
	envars.accessToken = os.Getenv("ACCESS_TOKEN")
	if envars.accessToken == "" {
		envars.accessToken = AccessToken
	}
	envars.sourceOwner = os.Getenv("SOURCE_OWNER")
	if envars.sourceOwner == "" {
		envars.sourceOwner = SourceOwner
	}
	envars.sourceRepo = os.Getenv("SOURCE_REPO")
	if envars.sourceRepo == "" {
		envars.sourceRepo = SourceRepo
	}
	envars.sourceBranch = os.Getenv("SOURCE_BRANCH")
	if envars.sourceBranch == "" {
		envars.sourceBranch = SourceBranch
	}
	if os.Getenv("CHECK_FREQUENCY") == "" {
		envars.checkFrequency = CheckFrequency
	} else {
		i, err := strconv.ParseInt(os.Getenv("CHECK_FREQUENCY"), 10, 64)
		if err != nil {
			panic(err)
		}
		envars.checkFrequency = int64(i)
	}
	if os.Getenv("UNUSED_DAYS") == "" {
		envars.unusedDays = UnusedDays
	} else {
		i, err := strconv.ParseInt(os.Getenv("UNUSED_DAYS"), 10, 64)
		if err != nil {
			panic(err)
		}
		envars.unusedDays = int64(i)
	}

	return &envars, nil
}

// Read configuration from the user editted yaml file
func (c *YamlConf) getConf(yamlPath string) *YamlConf {
	yamlFile, err := ioutil.ReadFile(yamlPath)
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

// BinaryEntry binary main handler
func BinaryEntry() {
	optYaml := getopt.StringLong("from-file", 'f', "", "The path of configuration file. Support yaml only")
	optHelp := getopt.BoolLong("help", 'h', "Help")
	optTarget := getopt.StringLong("ip", 'i', "127.0.0.1", "Support localhost only")
	optToken := getopt.StringLong("token", 't', "", "Your personal/organization Github Personal Access Token")
	optOwner := getopt.StringLong("owner", 'o', "lmchih", "Github repo owner: https://github.com/{owner}/{repo}")
	optRepo := getopt.StringLong("repo", 'r', "server-collector", "Github repo name: https://github.com/{owner}/{repo}")
	optBranch := getopt.StringLong("branch", 'b', "master", "Github repo branch (Support master only)")
	optCheck := getopt.Int64('c', 120, "Seconds between every check")
	optUnused := getopt.Int64Long("unused-days", 'u', 3, "Days considered unused")

	getopt.Parse()

	if *optYaml != "" {
		fmt.Printf("Yaml: %v\n", *optYaml)
		BinaryYamlEntry(*optYaml)
		return
	}

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}

	if *optToken == "" {
		fmt.Println("Github Access Token is required.")
		getopt.Usage()
		os.Exit(0)
	}

	fmt.Println("IP: " + *optTarget)
	fmt.Println("Token: " + *optToken)
	fmt.Println("Owner: " + *optOwner)
	fmt.Println("Repo: " + *optRepo)
	fmt.Println("Branch: " + *optBranch)
	fmt.Printf("Check Frequency: %d\n", *optCheck)
	fmt.Printf("Unused Days: %d\n", *optUnused)

	options := Options{
		IP:         *optTarget,
		Token:      *optToken,
		Owner:      *optOwner,
		Repo:       *optRepo,
		Branch:     *optBranch,
		CheckFreq:  *optCheck,
		UnusedDays: *optUnused,
	}

	BinaryOptEntry(&options)
}

// BinaryOptEntry run with options
func BinaryOptEntry(opts *Options) {
	log.Println("Check Github last commit date")
	BinaryRunCheck(opts)

	// Start to run the routine job
	for range time.Tick(time.Duration(opts.CheckFreq) * time.Second) {
		log.Println("Check Github last commit date")
		BinaryRunCheck(opts)
	}
}

// BinaryYamlEntry run with yaml file
func BinaryYamlEntry(yamlPath string) {
	fmt.Println("Start server-collector")

	var c YamlConf
	c.getConf(yamlPath)

	// fmt.Println(c.Version)

	// Run once first right before the routine
	log.Println("Check Github last commit date")
	options := configToOption(&c)
	BinaryRunCheck(options)

	// Start to run the routine job
	for range time.Tick(time.Duration(options.CheckFreq) * time.Second) {
		log.Println("Check Github last commit date")
		BinaryRunCheck(options)
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
	for range time.Tick(time.Duration(envs.checkFrequency) * time.Second) {
		log.Println("Check Github last commit date")
		ContainerRunCheck(envs)
	}
}

// RunCheck check github commit days
func RunCheck(a interface{}) {
	// TODO:
}

// BinaryRunCheck check github commit status
func BinaryRunCheck(o *Options) {
	// routinely check last commit date
	var days = LastCommitDays(o.Token, o.Owner, o.Repo)
	if days == -1 {
		log.Println("Cannot retrieve Github info.")
		return
	}
	if isUnused(days, o.UnusedDays) {
		shutdownCommand()
	}
}

// ContainerRunCheck check github commit status
func ContainerRunCheck(e *Envars) {
	days := LastCommitDays(e.accessToken, e.sourceOwner, e.sourceRepo)
	// if older than expiration, terminate the server.
	if days == -1 {
		log.Println("Cannot retrieve Github info.")
		return
	}
	if isUnused(days, e.unusedDays) {
		terminate(e.targetServer)
	}
}

func isUnused(days int64, expiration int64) bool {
	if days >= expiration {
		return true
	}
	return false
}
