package main

import (
	"gh_actions/actions"
	"gh_actions/common"
	"gh_actions/util"
)

func main() {

	ghRequest := util.BuildRequest()

	util.Log("Begin - " + ghRequest.Action + " ==============================")

	if ghRequest.Action == common.RUN_COMPLETE_WORKFLOW {

		actions.RunCompleteWorkflow(ghRequest)

	} else if ghRequest.Action == common.TRIGGER_WORKFLOW_BY_WFID {

		actions.TriggerWorkflow(ghRequest)

	} else if ghRequest.Action == common.RERUN_WORKFLOW_BY_RUNID {

		actions.RerunWorkflow(ghRequest)

	} else if ghRequest.Action == common.GET_COMPLETED_RUNS_BY_WFID {

		actions.GetCompletedRunsByWorkflowId(ghRequest)

	} else if ghRequest.Action == common.GET_WF_RUN_BY_RUNID {

		ghResult := actions.GetWorkflowRunByRunId(ghRequest)
		run := ghResult.Result
		util.WriteResult(run)

	} else if ghRequest.Action == common.GET_JOBS_BY_RUNID {

		ghResult := actions.GetJobsByRunId(ghRequest)
		util.LogJson(ghResult)

	} else if ghRequest.Action == common.GET_JOB_LOG_BY_JOBID {

		ghResult := actions.GetJobLogsByJobId(ghRequest)
		util.LogJson(ghResult)

	} else if ghRequest.Action == common.GET_RUN_LOGS_BY_RUNID {

		ghResult := actions.GetRunLogsByRunId(ghRequest)
		util.LogJson(ghResult)

	}
	util.Log("End - " + ghRequest.Action + " ==============================")
}
