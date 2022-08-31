package hooks

type GithubOrganization struct {
	IssuesURL string `json:"issues_url"`
	ReposURL  string `json:"repos_url"`
}

type GithubFullPayload struct {
	GithubOrganization `json:"organization"`
	Action             string `json:"action"`
}

type GithubEvent struct {
	GithubFullPayload `json:"fullPayload"`
}
