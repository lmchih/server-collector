package collector

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// If not providing environment variables, fallback to default
const (
	TARGET_SERVER = "127.0.0.1"
	SOURCE_OWNER  = "lmchih"
	SOURCE_REPO   = "server-collector"
	SOURCE_BRANCH = "master"
	UNUSED_DAYS   = 3
)

type EnvVars struct {
	targetServer string
	sourceOwner  string
	sourceRepo   string
	sourceBranch string
	unusedDays   int64
}

// GetEnvs configruation is supposed to be injected from k8s yaml
// or ducker run -e
func GetEnvs() (*EnvVars, error) {
	// default env from yaml
	log.Printf("TARGET_SERVER=%s\n", os.Getenv("TARGET_SERVER"))
	log.Printf("SOURCE_OWNER=%s\n", os.Getenv("SOURCE_OWNER"))
	log.Printf("SOURCE_REPO=%s\n", os.Getenv("SOURCE_REPO"))
	log.Printf("SOURCE_BRANCH=%s\n", os.Getenv("SOURCE_BRANCH"))
	log.Printf("UNUSED_DAYS=%s\n", os.Getenv("UNUSED_DAYS"))

	envVars := EnvVars{}
	envVars.targetServer = os.Getenv("TARGET_SERVER")
	if envVars.targetServer == "" {
		envVars.targetServer = TARGET_SERVER
	}
	envVars.sourceOwner = os.Getenv("SOURCE_OWNER")
	if envVars.sourceOwner == "" {
		envVars.sourceOwner = SOURCE_OWNER
	}
	envVars.sourceRepo = os.Getenv("SOURCE_REPO")
	if envVars.sourceRepo == "" {
		envVars.sourceRepo = SOURCE_REPO
	}
	envVars.sourceBranch = os.Getenv("SOURCE_BRANCH")
	if envVars.sourceBranch == "" {
		envVars.sourceBranch = SOURCE_BRANCH
	}

	if os.Getenv("UNUSED_DAYS") == "" {
		envVars.unusedDays = UNUSED_DAYS
	} else {
		i, err := strconv.ParseInt(os.Getenv("UNUSED_DAYS"), 10, 64)
		if err != nil {
			panic(err)
		}
		envVars.unusedDays = int64(i)
	}

	return &envVars, nil
}

func Handler() {
	fmt.Println("Start server-collector")

	envs, err := GetEnvs()
	log.Printf("envs: %v, err: %v", envs, err)
	if err != nil {
		log.Fatal(err)
	}

	// if isUnused(envs) {
	// 	terminate(envs.targetServer)
	// }

	terminate(envs.targetServer)
}
