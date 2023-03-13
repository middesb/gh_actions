package main

import (
	"gh_actions/rest"
	"os"
	"time"
)

const (
	GITHUB_GET_RUNS = "GITHUB_GET_RUNS"
)

func main() {

	args := os.Args

	action := args[4]
	rest.Log("action: ", action)

	if action == GITHUB_GET_RUNS {
		rest.Log(time.Now().Format(time.RFC3339))
		rest.Log(time.Parse(time.RFC3339, "2023-03-12T07:53:23Z"))
		runs, err := GetGithubActionsWorkflowRuns(args)
		if err != "" {
			rest.Log("Error occurred  - GITHUB_GET_RUNS request  [ERROR] -", err)
		} else {
			rest.Log("run size:", len(runs.Runs))
			rest.Log("ID:", runs.Runs[0].ID, "WorkflowID:", runs.Runs[0].WorkflowID, "Actor:", runs.Runs[0].Actor.Login)
		}
	} else if action == rest.GITHUB_DISPATCH {
		status, err := DispatchGitHubActionsWorkflow(args)
		if err != "" {
			rest.Log("Error occurred  - GITHUB_DISPATCH request  [ERROR] -", err)
		}
		rest.Log(status)
	}

}
func GetGithubActionsWorkflowRuns(args []string) (*rest.GHWorkflowRunInfo, string) {
	host := args[1]
	rest.Log("host: ", host)
	token := args[2]
	repoPath := args[3]
	rest.Log("repoPath: ", repoPath)
	workflowId := args[5]
	tailUrl := "actions/workflows/" + workflowId + "/runs?status=success&per_page=1"

	url := rest.BuildUrl(host, repoPath, tailUrl)

	rest.Log("url:", url)

	resp, errStr := rest.Get(url, token)
	if errStr != "" {
		return nil, "Error in GET request for GITHUB_GET_RUNS [ERROR] -" + errStr
	}

	runs := new(rest.GHWorkflowRunInfo)

	errStr = rest.FetchObjectFromBody(resp, runs)
	if errStr == "" {
		return runs, ""
	}

	return nil, "Error in GET request for GITHUB_GET_RUNS [ERROR] Unmarshalling -" + errStr

}
func DispatchGitHubActionsWorkflow(args []string) (int, string) {
	host := args[1]
	rest.Log("host: ", host)
	token := args[2]
	repoPath := args[3]
	rest.Log("repoPath: ", repoPath)
	workflowId := args[5]
	tailUrl := "/actions/workflows/" + workflowId + "/dispatches"

	url := rest.BuildUrl(host, repoPath, tailUrl)
	rest.Log("url:", url)

	jsonBytes := []byte(`{"ref":"main"}`)

	status, errStr := rest.Post(url, token, jsonBytes)

	if errStr != "" {
		return status, "Error POST request for GITHUB_DISPATCH [ERROR] -" + errStr
	}
	return status, ""
}
