package main

import (
	"gh_actions/rest"
	"time"
)

const (
	GITHUB_GET_RUNS = "GITHUB_GET_RUNS"
)

func main() {

	ghRequest := rest.BuildRequest()

	if ghRequest.Action == rest.TRIGGER_WORKFLOW {

		TriggerWorkflow(ghRequest)

	} else if ghRequest.Action == rest.RERUN_WORKFLOW {

		RerunWorkflow(ghRequest)

	} else if ghRequest.Action == rest.GET_COMPLETED_RUNS_BY_WFID {

		GetCompletedRunsByWorkflowId(ghRequest)

	} else if ghRequest.Action == rest.GET_WF_RUN_BY_RUNID {

		ghResult := rest.GetWorkflowRunByRunId(ghRequest)
		rest.LogJson(ghResult)

	} else if ghRequest.Action == rest.GET_JOBS_BY_RUNID {

		ghResult := rest.GetJobsByRunId(ghRequest)
		rest.LogJson(ghResult)

	} else if ghRequest.Action == rest.GET_JOB_LOG_BY_JOBID {

		ghResult := rest.GetJobLogsByJobId(ghRequest)
		rest.LogJson(ghResult)

	} else if ghRequest.Action == rest.GET_RUN_LOGS_BY_RUNID {

		ghResult := rest.GetRunLogsByRunId(ghRequest)
		rest.LogJson(ghResult)

	}
	rest.Log("End - " + ghRequest.Action + " ==============================")
}

func TriggerWorkflow(ghRequest rest.GHRequest) {

	// fetch last run for the workflow Id
	ghResultRuns := rest.GetRunsByWorkflowId(ghRequest, false)
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
		ghResultInt := rest.RunWorkflowById(ghRequest)

		if ghResultInt.StatusCode > 199 && ghResultInt.StatusCode < 400 {

			retries := 5
			for retries > 0 {
				time.Sleep(5 * time.Second)
				// fetch last run for the workflow Id
				ghResultRuns = rest.GetWorkflowRunsSince(ghRequest, lastRunTimeStamp)

				for ghResultRuns.StatusCode > 199 && ghResultRuns.StatusCode < 400 {
					runsInfo := ghResultRuns.Result

					// fetch triggered Run
					run := rest.FetchTriggeredRun(ghRequest.ID, lastRunTimeStamp, runsInfo)
					// if run exists
					if run != nil {
						var result rest.GHResult[rest.GHWorkflowRun]
						result.Result = *run
						rest.LogJson(result)
						return
					}
				}
				retries--
			}

			ghResultInt.Message = "Error  - " + ghRequest.Action + " request failed  fetching Last  run for WorkflowId " + ghRequest.ID + " [Retry Count  Remaining - 0] -"
			ghResultInt.StatusCode = 400
			return
		}
		rest.LogJson(ghResultInt)
		return
	}

	rest.LogJson(ghResultRuns)
}

func RerunWorkflow(ghRequest rest.GHRequest) {

	// fetch last run for the workflow Id
	ghResultRun := rest.GetWorkflowRunByRunId(ghRequest)
	if ghResultRun.StatusCode > 199 && ghResultRun.StatusCode < 400 {

		run := ghResultRun.Result

		lastRunAttempt := run.RunAttempt

		// trigger workflow
		ghResultInt := rest.RerunWorkflowByRunId(ghRequest)

		if ghResultInt.StatusCode > 199 && ghResultInt.StatusCode < 400 {

			retries := 5
			for retries > 0 {
				time.Sleep(5 * time.Second)
				ghResultRun := rest.GetWorkflowRunByRunId(ghRequest)
				run := ghResultRun.Result

				if ghResultRun.StatusCode > 199 && ghResultRun.StatusCode < 400 && run.RunAttempt <= lastRunAttempt {
					rest.LogJson(ghResultRun)
					return
				}
				retries--
			}

			ghResultInt.Message = "Error  - " + ghRequest.Action + " request failed  fetching Last  run for runId " + ghRequest.ID + " [Retry Count  Remaining - 0]"
			ghResultInt.StatusCode = 400

		}
		rest.LogJson(ghResultInt)
	}
	rest.LogJson(ghResultRun)

}

func GetCompletedRunsByWorkflowId(ghRequest rest.GHRequest) {

	// fetch completed runs for the workflow Id
	ghResult := rest.GetRunsByWorkflowId(ghRequest, true)

	rest.LogJson(ghResult)
}
