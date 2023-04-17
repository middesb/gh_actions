package actions

import (
	"errors"
	"gh_actions/common"
	"gh_actions/rest"
	"gh_actions/util"
	"net/http"
	"testing"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestGetRunsByWorkflowId(t *testing.T) {

	rest.Client = &MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(
			"Error from web server",
		)
	}

	var ghRequest common.GHRequest
	wfid := "29051699"
	ghRequest.ID = wfid
	var isCompleted bool
	ghResult := GetRunsByWorkflowId(ghRequest, isCompleted)
	if ghResult.StatusCode > 399 {

		t.Logf(" SUCCEDED,  %v, got ->  ", ghResult.Params)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := GetRunsByWorkflowIdResponse()

		return &resp, nil

	}

	ghResult = GetRunsByWorkflowId(ghRequest, isCompleted)
	util.Log(ghResult)
	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", result.TotalCount)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
}

func TestGetWorkflowRunByRunId(t *testing.T) {
	rest.Client = &MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}
		return resp, errors.New(
			"Error from web server",
		)
	}

	var ghRequest common.GHRequest
	runId := "4604017430"
	ghRequest.ID = runId
	ghResult := GetWorkflowRunByRunId(ghRequest)
	if ghResult.StatusCode > 399 {

		t.Logf(" SUCCEDED, got ->  %v", ghResult)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := GetWorkflowRunByRunIdResponse("1")

		return &resp, nil

	}

	ghResult = GetWorkflowRunByRunId(ghRequest)

	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", result.ID)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
}

func TestGetWorkflowRunsSince(t *testing.T) {
	rest.Client = &MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}
		return resp, errors.New(
			"Error from web server",
		)
	}

	var ghRequest common.GHRequest
	sinceTimeStamp := "2023-04-04T04:35:12Z"
	tt, err := util.ParseGHADate(sinceTimeStamp)

	util.Log("tt:", tt, "err:", err)

	ghRequest.ID = "29051699"
	ghResult := GetWorkflowRunsSince(ghRequest, sinceTimeStamp)
	if ghResult.StatusCode > 399 {

		t.Logf(" SUCCEDED, got ->  %v", ghResult)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := GetWorkflowRunsSinceResponse(sinceTimeStamp)

		return &resp, nil

	}

	ghResult = GetWorkflowRunsSince(ghRequest, sinceTimeStamp)

	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", result.TotalCount)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	sinceTimeStamp = "2023-04-04T04:35:13Z"
	ghResult = GetWorkflowRunsSince(ghRequest, sinceTimeStamp)

	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", result.TotalCount)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
}

func TestGetJobsByRunId(t *testing.T) {
	rest.Client = &MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}
		return resp, errors.New(
			"Error from web server",
		)
	}

	var ghRequest common.GHRequest
	runId := "4604017430"
	ghRequest.ID = runId
	ghResult := GetJobsByRunId(ghRequest)
	if ghResult.StatusCode > 399 {

		t.Logf(" SUCCEDED, got ->  %v", ghResult)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := GetJobsByRunIdResponse()

		return &resp, nil

	}

	ghResult = GetJobsByRunId(ghRequest)

	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", len(result.Jobs))
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
}

func TestGetJobLogsByJobId(t *testing.T) {
	rest.Client = &MockClient{}

	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}

		return resp, errors.New(
			"Error from web server",
		)

	}

	var ghRequest common.GHRequest
	jobId := "4604017430"
	ghRequest.ID = jobId
	ghResult := GetJobLogsByJobId(ghRequest)
	if ghResult.StatusCode > 399 {

		t.Logf(" SUCCEDED, got ->  %v", ghResult)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := GetJobLogsByJobIdResponse()

		return &resp, nil

	}

	ghResult = GetJobLogsByJobId(ghRequest)

	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", len(result.Logs))
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
}

func TestRunWorkflowById(t *testing.T) {
	rest.Client = &MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}
		return resp, errors.New(
			"Error from web server",
		)

	}

	var ghRequest common.GHRequest
	wfId := "28229045"
	ghRequest.ID = wfId
	ghResult := RunWorkflowById(ghRequest)
	if ghResult.StatusCode > 399 {

		t.Logf(" SUCCEDED, got ->  %v", ghResult)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := RunWorkflowByIdResponse()

		return &resp, nil

	}

	ghResult = RunWorkflowById(ghRequest)

	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", result)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
}

func TestRerunWorkflowByRunId(t *testing.T) {
	rest.Client = &MockClient{}

	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}

		return resp, errors.New(
			"Error from web server",
		)

	}

	var ghRequest common.GHRequest
	runId := "4467742027"
	ghRequest.ID = runId
	ghResult := RerunWorkflowByRunId(ghRequest)
	if ghResult.StatusCode > 399 {

		t.Logf(" SUCCEDED, got ->  %v", ghResult)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := RerunWorkflowByRunIdResponse()

		return &resp, nil

	}

	ghResult = RerunWorkflowByRunId(ghRequest)

	if ghResult.StatusCode < 400 {
		result := ghResult.Result
		t.Logf(" SUCCEDED,  got ->  %v", result)
	} else {
		t.Errorf(" FAILED,   got -> %v", ghResult)
	}
}

func TestRerunWorkflow(t *testing.T) {
	rest.Client = &MockClient{}
	doFuncCount := 0

	GetDoFunc = func(*http.Request) (*http.Response, error) {
		doFuncCount++
		if doFuncCount == 1 {
			resp := GetWorkflowRunByRunIdResponse("2")
			util.Log("resp 1:", resp.StatusCode)
			return &resp, nil

		} else if doFuncCount == 2 {
			resp := RerunWorkflowByRunIdResponse()
			util.Log("resp 2:", resp.StatusCode)
			return &resp, nil

		} else if doFuncCount == 3 {
			resp := GetWorkflowRunByRunIdResponse("3")
			util.Log("resp 3:", resp.StatusCode)
			return &resp, nil

		} else {
			return nil, errors.New(
				"Error from web server",
			)
		}
	}
	var ghRequest common.GHRequest
	runId := "4467742027"
	ghRequest.ID = runId
	returnVal := RerunWorkflow(ghRequest)

	if returnVal == 0 {
		t.Logf(" SUCCEDED,  got ->  %v", returnVal)
	} else {
		t.Errorf(" FAILED,   got -> %v", returnVal)
	}
}

func TestRunCompleteWorkflow(t *testing.T) {
	rest.Client = &MockClient{}
	doFuncCount := 0

	GetDoFunc = func(*http.Request) (*http.Response, error) {
		doFuncCount++
		if doFuncCount == 1 {
			resp := GetWorkflowRunByRunIdResponse("1")
			util.Log("resp 1:", resp.StatusCode)
			return &resp, nil

		} else if doFuncCount == 2 {
			resp := RunWorkflowByIdResponse()
			util.Log("resp 2:", resp.StatusCode)
			return &resp, nil

		} else if doFuncCount == 3 {
			sinceTimeStamp := "2023-04-04T04:35:13Z"
			resp := GetWorkflowRunsSinceResponse(sinceTimeStamp)
			util.Log("resp 3:", resp.StatusCode, resp)
			return &resp, nil

		} else if doFuncCount == 4 {
			resp := GetJobsByRunIdResponse()
			util.Log("resp :", doFuncCount, resp.StatusCode)
			return &resp, nil

		} else if doFuncCount == 5 || doFuncCount == 6 || doFuncCount == 7 {
			resp := GetJobLogsByJobIdResponse()
			util.Log("resp :", doFuncCount, resp.StatusCode)
			return &resp, nil

		} else {
			return nil, errors.New(
				"Error from web server",
			)
		}
	}
	var ghRequest common.GHRequest
	wfid := "29051699"
	ghRequest.ID = wfid
	returnVal := RunCompleteWorkflow(ghRequest)

	if returnVal == 0 {
		t.Logf(" SUCCEDED,  got ->  %v", returnVal)
	} else {
		t.Errorf(" FAILED,   got -> %v", returnVal)
	}
}
