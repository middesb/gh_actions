package actions

import (
	"gh_actions/common"
	"gh_actions/rest"
	"gh_actions/util"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

/*
Fetch All Runs Information By Workflow ID
*/
func GetRunsByWorkflowId(ghRequest common.GHRequest, isCompleted bool) common.GHResult[common.GHWorkflowRunInfo] {

	parStr := ""
	if isCompleted {
		parStr = "?status=success"
	}
	tailUrl := "actions/workflows/" + ghRequest.ID + "/runs" + parStr

	url := util.BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result common.GHResult[common.GHWorkflowRunInfo]

	result.Params = ghRequest
	result.Url = url

	resp, errStr := rest.Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	runs := new(common.GHWorkflowRunInfo)
	errStr = util.FetchObjectFromBody(resp, runs)

	if errStr == "" {
		result.Result = *runs
		result.StatusCode = resp.StatusCode
		return result
	}
	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}

/*
Fetch Workflow Run by Specific Run ID
*/
func GetWorkflowRunByRunId(ghRequest common.GHRequest) common.GHResult[common.GHWorkflowRun] {

	tailUrl := "/actions/runs/" + ghRequest.ID

	url := util.BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result common.GHResult[common.GHWorkflowRun]

	result.Params = ghRequest
	result.Url = url

	resp, errStr := rest.Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	run := new(common.GHWorkflowRun)

	errStr = util.FetchObjectFromBody(resp, run)

	if errStr == "" {
		result.Result = *run
		result.StatusCode = resp.StatusCode
		return result
	}

	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}

/*
Fetch Workflow Runs Info since a specific Time
Format for Specific Time  -  2023-03-14T07:54:23Z
*/
func GetWorkflowRunsSince(ghRequest common.GHRequest, sinceTimeStamp string) common.GHResult[common.GHWorkflowRunInfo] {
	// sinceTimeStamp format  2023-03-14T07:54:23Z

	tailUrl := "/actions/workflows/" + ghRequest.ID + "/runs?created=%3E" + sinceTimeStamp
	url := util.BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)
	var result common.GHResult[common.GHWorkflowRunInfo]

	result.Params = ghRequest
	result.Url = url

	resp, errStr := rest.Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}
	runs := new(common.GHWorkflowRunInfo)

	errStr = util.FetchObjectFromBody(resp, runs)

	if errStr == "" {
		result.Result = *runs
		result.StatusCode = resp.StatusCode
		return result
	}

	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}

/*
Helper func to filter  run
*/
func FetchTriggeredRun(workflowid string, lastRunTimestamp string, runInfo common.GHWorkflowRunInfo) *common.GHWorkflowRun {
	for i := 0; i < len(runInfo.Runs); i++ {
		run := runInfo.Runs[i]
		workflowIdStr := strconv.Itoa(run.WorkflowID)
		//	Log("WorkflowID:", workflowIdStr, "inworkflowid:", workflowId, workflowIdStr == workflowId)
		if workflowIdStr == workflowid && run.CreatedAt != lastRunTimestamp {
			//		Log("Debug : ID:", run.ID, "WorkflowID:", run.WorkflowID, "inworkflowid:", workflowId, "lastrunTimestamp:", lastRunTimestamp, "Actor:", run.Actor.Login, "Created:", run.CreatedAt, "Status:", run.Status)

			return &run
		}
	}
	return nil
}

/*
Fetch All Jobs Information By Run ID
*/
func GetJobsByRunId(ghRequest common.GHRequest) common.GHResult[common.GHJobInfo] {

	tailUrl := "actions/runs/" + ghRequest.ID + "/jobs"

	url := util.BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result common.GHResult[common.GHJobInfo]
	result.Params = ghRequest
	result.Url = url

	resp, errStr := rest.Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	jobInfo := new(common.GHJobInfo)

	errStr = util.FetchObjectFromBody(resp, jobInfo)

	if errStr == "" {
		result.Result = *jobInfo
		result.StatusCode = resp.StatusCode
		return result
	}

	result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] - Unmarshalling -" + errStr
	result.StatusCode = 400
	return result

}

/*
Fetch Logs for  By Job ID
*/
func GetJobLogsByJobId(ghRequest common.GHRequest) common.GHResult[common.GHJobLog] {

	tailUrl := "actions/jobs/" + ghRequest.ID + "/logs"

	url := util.BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result common.GHResult[common.GHJobLog]
	result.Params = ghRequest
	result.Url = url

	resp, errStr := rest.Get(url, ghRequest.Token)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}

	jobLogs := new(common.GHJobLog)

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

/*
Fetch All Logs for all Jobs in a Run
*/
func GetRunLogsByRunId(ghRequest common.GHRequest) common.GHResult[[]string] {

	ghResultJobs := GetJobsByRunId(ghRequest)
	jobs := ghResultJobs.Result.Jobs

	var runLogs common.GHResult[[]string]

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

/*
Helper to dispatch a Workflow called from Trigger Workflow
*/
func RunWorkflowById(ghRequest common.GHRequest) common.GHResult[int] {

	tailUrl := "/actions/workflows/" + ghRequest.ID + "/dispatches"

	url := util.BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result common.GHResult[int]
	result.Params = ghRequest
	result.Url = url

	jsonBytes := []byte(`{"ref":"main"}`)

	status, errStr := rest.Post(url, ghRequest.Token, jsonBytes)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}
	result.Result = status
	result.StatusCode = status
	return result
}

/*
Helper to ReRun a specific Run  By Run ID called from RerunWorkflow
*/
func RerunWorkflowByRunId(ghRequest common.GHRequest) common.GHResult[int] {

	tailUrl := "/actions/runs/" + ghRequest.ID + "/rerun"

	url := util.BuildUrl(ghRequest.Host, ghRequest.RepoPath, tailUrl)

	var result common.GHResult[int]
	result.Params = ghRequest
	result.Url = url

	jsonBytes := []byte(`{"ref":"main"}`)

	status, errStr := rest.Post(url, ghRequest.Token, jsonBytes)

	if errStr != "" {
		result.Message = "Error in GET request for " + ghRequest.Action + " [ERROR] -" + errStr
		result.StatusCode = 400
		return result
	}
	result.Result = status
	result.StatusCode = status
	return result
}

/*
Trigger A Workflow  by workflow Id and also Fetch and Log triggerred run
*/
func TriggerWorkflow(ghRequest common.GHRequest) {

	// fetch last run for the workflow Id
	ghResultRuns := GetRunsByWorkflowId(ghRequest, false)
	if ghResultRuns.StatusCode > 199 && ghResultRuns.StatusCode < 400 {
		// init the last run timestamp as current
		lastRunTimeStamp := time.Now().UTC().Format(time.RFC3339)
		runsInfo := ghResultRuns.Result

		runs := runsInfo.Runs

		if runsInfo.TotalCount > 0 {
			// if run exists then fetch the last timestamp from it
			lastRunTimeStamp = runs[0].CreatedAt
		}

		// trigger workflow
		ghResultInt := RunWorkflowById(ghRequest)

		if ghResultInt.StatusCode > 199 && ghResultInt.StatusCode < 400 {

			retries := 5
			for retries > 0 {
				time.Sleep(5 * time.Second)
				// fetch last run for the workflow Id
				ghResultRuns = GetWorkflowRunsSince(ghRequest, lastRunTimeStamp)

				for ghResultRuns.StatusCode > 199 && ghResultRuns.StatusCode < 400 {
					runsInfo := ghResultRuns.Result

					// fetch triggered Run
					run := FetchTriggeredRun(ghRequest.ID, lastRunTimeStamp, runsInfo)
					// if run exists
					if run != nil {
						var result common.GHResult[common.GHWorkflowRun]
						result.Result = *run
						result.Url = ghResultInt.Url
						result.StatusCode = ghResultRuns.StatusCode
						result.Params = ghRequest
						util.LogJson(result)
						return
					}
				}
				retries--
			}

			ghResultInt.Message = "Error  - " + ghRequest.Action + " request failed  fetching Last  run for WorkflowId " + ghRequest.ID + " [Retry Count  Remaining - 0] -"
			ghResultInt.StatusCode = 400

		}
		util.LogJson(ghResultInt)
		return
	}

	util.LogJson(ghResultRuns)
}

/*
Rerun workflow and fetch and Log the Run results
*/
func RerunWorkflow(ghRequest common.GHRequest) {

	// fetch last run for the workflow Id
	ghResultRun := GetWorkflowRunByRunId(ghRequest)
	if ghResultRun.StatusCode > 199 && ghResultRun.StatusCode < 400 {

		run := ghResultRun.Result

		lastRunAttempt := run.RunAttempt

		// trigger workflow
		ghResultInt := RerunWorkflowByRunId(ghRequest)

		if ghResultInt.StatusCode > 199 && ghResultInt.StatusCode < 400 {

			retries := 5
			for retries > 0 {
				time.Sleep(5 * time.Second)
				ghResultRun = GetWorkflowRunByRunId(ghRequest)
				run := ghResultRun.Result

				if ghResultRun.StatusCode > 199 && ghResultRun.StatusCode < 400 && run.RunAttempt > lastRunAttempt {
					ghResultRun.Url = ghResultInt.Url
					util.LogJson(ghResultRun)
					return
				}
				retries--
			}

			ghResultInt.Message = "Error  - " + ghRequest.Action + " request failed  fetching Last  run for runId " + ghRequest.ID + " [Retry Count  Remaining - 0]"
			ghResultInt.StatusCode = 400

		}
		util.LogJson(ghResultInt)
	}
	util.LogJson(ghResultRun)

}

/*
Fetches all completed runs for a given Workflow Id
*/
func GetCompletedRunsByWorkflowId(ghRequest common.GHRequest) {

	// fetch completed runs for the workflow Id
	ghResult := GetRunsByWorkflowId(ghRequest, true)

	util.LogJson(ghResult)
}

/*
Run the Workflow and wait for all the jobs/steps to complete log information periodically until workflow completes
*/
func RunCompleteWorkflow(ghRequest common.GHRequest) {

	// trigger workflow
	// fetch last run for the workflow Id
	ghResultRuns := GetRunsByWorkflowId(ghRequest, false)
	if ghResultRuns.StatusCode > 199 && ghResultRuns.StatusCode < 400 {
		// init the last run timestamp as current
		lastRunTimeStamp := time.Now().UTC().Format(time.RFC3339)
		runsInfo := ghResultRuns.Result

		runs := runsInfo.Runs

		if runsInfo.TotalCount > 0 {
			// if run exists then fetch the last timestamp from it
			lastRunTimeStamp = runs[0].CreatedAt
		}

		// trigger workflow
		ghResultInt := RunWorkflowById(ghRequest)

		if ghResultInt.StatusCode > 199 && ghResultInt.StatusCode < 400 {

			retries := 5
			for retries > 0 {
				time.Sleep(5 * time.Second)
				// fetch last run for the workflow Id
				ghResultRuns = GetWorkflowRunsSince(ghRequest, lastRunTimeStamp)

				for ghResultRuns.StatusCode > 199 && ghResultRuns.StatusCode < 400 {
					runsInfo := ghResultRuns.Result

					// fetch triggered Run
					run := FetchTriggeredRun(ghRequest.ID, lastRunTimeStamp, runsInfo)
					// if run exists
					if run != nil {
						var result common.GHResult[common.GHWorkflowRun]
						result.Result = *run

						ghStatusResult := GetRunJobStatus(ghRequest, *run)
						util.LogJson(ghStatusResult)
						return
					}
				}
				retries--
			}

			ghResultInt.Message = "Error  - " + ghRequest.Action + " request failed  fetching Last  run for WorkflowId " + ghRequest.ID + " [Retry Count  Remaining - 0] -"
			ghResultInt.StatusCode = 400

		}
		util.LogJson(ghResultInt)
		return
	}

	util.LogJson(ghResultRuns)

}

/*
Helper function to fetch Jobstatus for a Run called from func RunCompleteWorkflow
*/
func GetRunJobStatus(ghRequest common.GHRequest, run common.GHWorkflowRun) common.GHResult[common.GHJobInfo] {

	ghRequest.ID = strconv.Itoa(run.ID)

	for {
		ghResult := GetJobsByRunId(ghRequest)
		ghJobInfo := ghResult.Result
		if ghResult.StatusCode > 199 && ghResult.StatusCode < 400 {
			jobCount := 0
			for i := 0; i < len(ghJobInfo.Jobs); i++ {
				job := ghJobInfo.Jobs[i]
				if job.Conclusion == "success" {
					jobCount++
				}

			}
			util.Log("jobCount:", jobCount, " Jobs : ", len(ghJobInfo.Jobs))
			if jobCount == len(ghJobInfo.Jobs) {
				ghJobInfo.Conclusion = "success"
				ghJobInfo.Status = "completed"
				ghJobInfo.RunID = run.ID
				ghJobInfo.WorkflowID = run.WorkflowID
			} else {
				ghJobInfo.Conclusion = "in_progress"
				ghJobInfo.Status = "in_progress"
				ghJobInfo.RunID = run.ID
				ghJobInfo.WorkflowID = run.WorkflowID
			}
			ghResult.Result = ghJobInfo
		} else {
			ghJobInfo.Conclusion = "failed"
			ghJobInfo.Status = "failed"

		}
		ghResult.Result = ghJobInfo
		util.LogJson(ghResult)
		if ghJobInfo.Conclusion == "success" || ghJobInfo.Conclusion == "failed" {
			util.Log("Returning...")
			return ghResult
		}
		time.Sleep(5 * time.Second)
	}

}
