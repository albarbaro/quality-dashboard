package api

import (
	"log"
	"net/http"

	"github.com/andygrunwald/go-jira"
	client "github.com/flacatus/qe-dashboard-backend/pkg/api/apis/jira"
)

// Version godoc
// @Summary Version
// @Description returns quality backend version
// @Tags Version API
// @Produce json
// @Router /api/version [get]
// @Success 200 {object} api.MapResponse
func (s *Server) jiraIssueKnown(w http.ResponseWriter, r *http.Request) {
	factory := client.NewTotClientFactory()
	jiraClient, err := factory.NewJiraClient()
	if err != nil {
		log.Fatal(err)
	}
	var issues []jira.Issue

	// append the jira issues to []jira.Issue
	appendFunc := func(i jira.Issue) (err error) {
		issues = append(issues, i)
		return err
	}

	// In this example, we'll search for all the issues with the provided JQL filter and Print the Story Points
	err = jiraClient.Issue.SearchPages("labels in (appstudio-e2e-tests-known-issues) AND status not in (resolved, closed)", nil, appendFunc)
	if err != nil {
		log.Fatal(err)
	}
	s.JSONResponse(w, r, issues)
}
