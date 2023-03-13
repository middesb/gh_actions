package rest

import (
	"time"
)

type GHActor struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}
type GHRepository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
}

type GHUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GHHeadCommit struct {
	ID        string    `json:"id"`
	TreeID    string    `json:"tree_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Author    GHUser    `json:"author"`
	Committer GHUser    `json:"committer"`
}

type GHWorkflowRun struct {
	ID         int64        `json:"id"`
	Name       string       `json:"name"`
	NodeID     string       `json:"node_id"`
	HeadBranch string       `json:"head_branch"`
	HeadSha    string       `json:"head_sha"`
	Path       string       `json:"path"`
	RunNumber  int          `json:"run_number"`
	Event      string       `json:"event"`
	Status     string       `json:"status"`
	Conclusion string       `json:"conclusion"`
	WorkflowID int          `json:"workflow_id"`
	HTMLURL    string       `json:"html_url"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	RunAttempt int          `json:"run_attempt"`
	Actor      GHActor      `json:"actor"`
	Repository GHRepository `json:"repository"`
	HeadCommit GHHeadCommit `json:"head_commit"`
}

type GHWorkflowRunInfo struct {
	Runs []GHWorkflowRun `json:"workflow_runs"`
}
