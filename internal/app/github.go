package collector

import (
	"fmt"
	"time"

	"context"

	"github.com/google/go-github/v28/github" // with go modules enabled
	"golang.org/x/oauth2"
)

var (
	client *github.Client
	ctx    = context.Background()
)

// GetClient Get github API client
func GetClient(token string) *github.Client {
	// get go-github client
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
	return client
}

// LastCommitDays Get how many days till today the last commit
// was pushed onto Github. If error caused, return -1
func LastCommitDays(token string, owner string, repo string) int64 {
	if client == nil {
		client = GetClient(token)
	}

	commitInfo, _, err := client.Repositories.ListCommits(ctx, owner, repo, nil)
	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		// os.Exit(1)
		return -1
	}

	// get the latest commit
	var lastCommit = commitInfo[0]
	var lastCommitDate = *lastCommit.Commit.Committer.Date

	// compare to commit time with now
	now := time.Now().UTC()
	fmt.Printf("Now: %v\n", now)
	since := time.Since(lastCommitDate)
	fmt.Printf("Since: %v\n", since)
	// convert since to days
	days := int64(since.Hours() / 24)
	fmt.Printf("last commmit was %d days ago.\n", days)
	return days
}
