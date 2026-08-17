package submissions

type InitResult struct {
	RepoUrl string
}

type PushEvent struct {
	Ref        string `json:"ref"` // "refs/heads/main" - check against default branch
	Before     string `json:"before"`
	After      string `json:"after"`   // new HEAD sha
	Deleted    bool   `json:"deleted"` // true = branch deletion, skip processing
	Forced     bool   `json:"forced"`
	Repository struct {
		ID            int64  `json:"id"`
		Name          string `json:"name"`      // "Submission-labib0x9-tcp-echo-server-template"
		FullName      string `json:"full_name"` // "ZERO9xz/Submission-labib0x9-tcp-echo-server-template"
		CloneURL      string `json:"clone_url"`
		SSHURL        string `json:"ssh_url"`
		DefaultBranch string `json:"default_branch"`
		Private       bool   `json:"private"`
	} `json:"repository"`
	HeadCommit struct {
		ID       string   `json:"id"` // exact commit sha to check out
		Distinct bool     `json:"distinct"`
		Modified []string `json:"modified"`
	} `json:"head_commit"`
	Sender struct {
		Login string `json:"login"` // who pushed — cross-check against submitter
	} `json:"sender"`
	Installation struct {
		ID int64 `json:"id"`
	} `json:"installation"`
}
