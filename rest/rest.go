package rest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetRunsByWorkflowId(ghRequest GHRequest, isCompleted bool) GHResult[GHWorkflowRunInfo] {

	parStr := ""
	if isCompleted {
		parStr = "?status=success"
	}
	tailUrl := "actions/workflows/" + ghRequest.ID + "/runs" + parStr

	url := BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result GHResult[GHWorkflowRunInfo]

	result.Params = ghRequest
	result.Url = url

	resp, errStr := Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	runs := new(GHWorkflowRunInfo)
	errStr = FetchObjectFromBody(resp, runs)

	if errStr == "" {
		result.Result = *runs
		result.StatusCode = resp.StatusCode
		return result
	}
	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}
func GetWorkflowRunByRunId(ghRequest GHRequest) GHResult[GHWorkflowRun] {

	tailUrl := "/actions/runs/" + ghRequest.ID

	url := BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result GHResult[GHWorkflowRun]

	result.Params = ghRequest
	result.Url = url

	resp, errStr := Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	run := new(GHWorkflowRun)

	errStr = FetchObjectFromBody(resp, run)

	if errStr == "" {
		result.Result = *run
		result.StatusCode = resp.StatusCode
		return result
	}

	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}
func GetWorkflowRunsSince(ghRequest GHRequest, sinceTimeStamp string) GHResult[GHWorkflowRunInfo] {
	// sinceTimeStamp format  2023-03-14T07:54:23Z

	tailUrl := "/actions/workflows/" + ghRequest.ID + "/runs?created=%3E" + sinceTimeStamp
	url := BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)
	var result GHResult[GHWorkflowRunInfo]

	result.Params = ghRequest
	result.Url = url

	resp, errStr := Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}
	runs := new(GHWorkflowRunInfo)

	errStr = FetchObjectFromBody(resp, runs)

	if errStr == "" {
		result.Result = *runs
		result.StatusCode = resp.StatusCode
		return result
	}

	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}

func FetchTriggeredRun(workflowId string, lastRunTimestamp string, runInfo GHWorkflowRunInfo) *GHWorkflowRun {
	for i := 0; i < len(runInfo.Runs); i++ {
		run := runInfo.Runs[i]
		workflowIdStr := strconv.Itoa(run.WorkflowID)
		//	Log("WorkflowID:", workflowIdStr, "inworkflowid:", workflowId, workflowIdStr == workflowId)
		if workflowIdStr == workflowId && run.CreatedAt != lastRunTimestamp {
			//		Log("Debug : ID:", run.ID, "WorkflowID:", run.WorkflowID, "inworkflowid:", workflowId, "lastrunTimestamp:", lastRunTimestamp, "Actor:", run.Actor.Login, "Created:", run.CreatedAt, "Status:", run.Status)

			return &run
		}
	}
	return nil
}
func GetJobsByRunId(ghRequest GHRequest) GHResult[GHJobInfo] {

	tailUrl := "actions/runs/" + ghRequest.ID + "/jobs"

	url := BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result GHResult[GHJobInfo]
	result.Params = ghRequest
	result.Url = url

	resp, errStr := Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	jobInfo := new(GHJobInfo)

	errStr = FetchObjectFromBody(resp, jobInfo)

	if errStr == "" {
		result.Result = *jobInfo
		result.StatusCode = resp.StatusCode
		return result
	}

	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}
func GetJobLogsByJobId(ghRequest GHRequest) GHResult[GHJobLog] {

	tailUrl := "actions/jobs/" + ghRequest.ID + "/logs"

	url := BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result GHResult[GHJobLog]
	result.Params = ghRequest
	result.Url = url

	resp, errStr := Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	jobLogs := new(GHJobLog)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}
	splitLines := strings.Split(strings.ReplaceAll(string(body), "\r\n", "\n"), "\n")

	jobLogs.Logs = splitLines

	jobLogs.ID = ghRequest.ID

	result.Result = *jobLogs
	result.StatusCode = resp.StatusCode
	return result

}
func GetRunLogsByRunId(ghRequest GHRequest) GHResult[[]string] {

	ghResultJobs := GetJobsByRunId(ghRequest)
	jobs := ghResultJobs.Result.Jobs

	var runLogs GHResult[[]string]

	runLogs.Params = ghRequest
	runLogs.Url = ghResultJobs.Url

	temp := ""
	for i := 0; i < len(jobs); i++ {
		jobIdStr := strconv.Itoa(jobs[i].ID)
		ghRequest.ID = jobIdStr
		ghJobLogs := GetJobLogsByJobId(ghRequest)
		temp = temp + strings.Join(ghJobLogs.Result.Logs, "\r\n")
		runLogs.Message = ghJobLogs.Message
		runLogs.StatusCode = ghJobLogs.StatusCode

	}
	splitLines := strings.Split(strings.ReplaceAll(temp, "\r\n", "\n"), "\n")
	runLogs.Result = splitLines

	return runLogs
}
func RunWorkflowById(ghRequest GHRequest) GHResult[int] {

	tailUrl := "/actions/workflows/" + ghRequest.ID + "/dispatches"

	url := BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result GHResult[int]
	result.Params = ghRequest
	result.Url = url

	jsonBytes := []byte(`{"ref":"main"}`)

	status, errStr := Post(url, ghRequest.Token, jsonBytes)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}
	result.Result = status
	result.StatusCode = status
	return result
}
func RerunWorkflowByRunId(ghRequest GHRequest) GHResult[int] {

	tailUrl := "/actions/runs/" + ghRequest.ID + "/rerun"

	url := BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result GHResult[int]
	result.Params = ghRequest
	result.Url = url

	jsonBytes := []byte(`{"ref":"main"}`)

	status, errStr := Post(url, ghRequest.Token, jsonBytes)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}
	result.Result = status
	result.StatusCode = status
	return result
}

func Get(url string, bearerToken string) (*http.Response, string) {
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + bearerToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, "Error creating request]n[ERROR] -" + err.Error()
	}
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, "Error on invoking GET call client.\n[ERROR] -" + err.Error()
	}

	return resp, ""
}

func Post(url string, bearerToken string, jsonBytes []byte) (int, string) {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + bearerToken

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return -1, "Error while creating the request for POST bytes:" + err.Error()
	}
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	Log(resp.StatusCode)
	if err != nil || resp.StatusCode > 299 {
		return resp.StatusCode, "Error on response.\n[ERROR] -" + resp.Status
	}

	//errStr := LogResponse(resp)

	return resp.StatusCode, ""
}
