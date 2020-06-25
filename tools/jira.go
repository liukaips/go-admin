package tools

import (
	"container/list"
	"github.com/andygrunwald/go-jira"
	"strings"
)

func versionlist() *list.List {
	jiraClient, err := jira.NewClient(nil, "https://jira.shopee.io/")
	if err != nil {
		panic(err)
	}
	version_list := list.New()
	res, err := jiraClient.Authentication.AcquireSessionCookie("kai.liu@shopee.com", "Liukaiok$1234")
	if err != nil || res == false {
		panic(err)
	}
	// issue, _, err := jiraClient.Issue.Get("AIRPAY-32164", nil)
	issue, _, err := jiraClient.Project.Get("AIRPAY")
	if err != nil {
		panic(err)
	}
	for i := range issue.Versions {
		version := issue.Versions[i].Name
		if strings.HasPrefix(version, "ps_") {
			version_list.PushBack(version)
		}

	}
	return version_list
}
