package queue

type PushMessage struct {
	Id       string `json:"id"`
	RepoName string `json:"repo_name"`
	Username string `json:"username"`
}
