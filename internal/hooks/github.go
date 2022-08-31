package hooks

type GithubOrganization struct {
	IssuesURL string `json:"issues_url"`
	ReposURL  string `json:"repos_url"`
}

type GithubSender struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

type GithubRelease struct {
	AssetsURL string `json:"assets_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
}

type GithubEvent struct {
	GithubOrganization `json:"organization,omitempty"`
	GithubSender       `json:"sender,omitempty"`
	GithubRelease      `json:"release,omitempty"`
	Action             string `json:"action,omitempty"`
}
