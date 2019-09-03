package collector

import (
	"fmt"
	"testing"

	. "github.com/lmchih/server-collector/internal/app"
)

func TestGetClient(t *testing.T) {
	var token = ""
	client := GetClient(token)

	if client == nil {
		t.Error("FAILED. Not able to get Github client\n")
	} else {
		t.Logf("PASSED. Github client is %v\n", client)
	}
}

func TestLastCommitDays(t *testing.T) {
	var token = ""
	var owner = "lmchih"
	var repo = "server-collector"

	days := LastCommitDays(token, owner, repo)
	fmt.Printf("days: %v\n", days)
	if days < 0 {
		t.Errorf("FAILED. Not able to get Github info %d\n", days)
	} else {
		t.Logf("PASSED. Your commit is older then %v days\n", days)
	}
}
