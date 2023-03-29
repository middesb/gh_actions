package common

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
	ID        string `json:"id"`
	TreeID    string `json:"tree_id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Author    GHUser `json:"author"`
	Committer GHUser `json:"committer"`
}

type GHWorkflowRun struct {
	ID         int          `json:"id"`
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
	CreatedAt  string       `json:"created_at"`
	UpdatedAt  string       `json:"updated_at"`
	RunAttempt int          `json:"run_attempt"`
	Actor      GHActor      `json:"actor"`
	Repository GHRepository `json:"repository"`
	HeadCommit GHHeadCommit `json:"head_commit"`
}

type GHWorkflowRunInfo struct {
	TotalCount int             `json:"total_count"`
	Runs       []GHWorkflowRun `json:"workflow_runs"`
}

type GHJobStep struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Conclusion  string `json:"conclusion"`
	Number      int    `json:"number"`
	StartedAt   string `json:"started_at"`
	CompletedAt string `json:"completed_at"`
}

type GHJob struct {
	ID          int         `json:"id"`
	RunID       int         `json:"run_id"`
	RunAttempt  int         `json:"run_attempt"`
	Status      string      `json:"status"`
	Conclusion  string      `json:"conclusion"`
	CreatedAt   string      `json:"created_at"`
	StartedAt   string      `json:"started_at"`
	CompletedAt string      `json:"completed_at"`
	Name        string      `json:"name"`
	Steps       []GHJobStep `json:"steps"`
	CheckRunURL string      `json:"check_run_url"`
	Labels      []string    `json:"labels"`
	RunnerID    int         `json:"runner_id"`
	RunnerName  string      `json:"runner_name"`
}

type GHJobInfo struct {
	ID         int     `json:"id"`
	Jobs       []GHJob `json:"jobs"`
	Status     string  `json:"status"`
	Conclusion string  `json:"conclusion"`
	RunID      int     `json:"run_id"`
	WorkflowID int     `json:"workflow_id"`
}

type GHJobLog struct {
	ID   string   `json:"id"`
	Logs []string `json:"logs"`
}
type GHRunLogs struct {
	ID      int      `json:"id"`
	RunLogs []string `json:"run_logs"`
}

type GHResult[T any] struct {
	Url        string    `json:"url"`
	Params     GHRequest `json:"params"`
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	Result     T         `json:"result"`
}

type GHRequest struct {
	Action   string `json:"action"`
	Host     string `json:"host"`
	Token    string `json:"token"`
	RepoPath string `json:"repo_path"`
	ID       string `json:"id"`
}
