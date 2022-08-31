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

type GithubCommit struct {
	Author struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"author,omitempty"`
	Committer struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"committer,omitempty"`
	Message string `json:"message,omitempty"`
	ID      string `json:"id,omitempty"`
	URL     string `json:"url,omitempty"`
}

type GithubEvent struct {
	GithubOrganization `json:"organization,omitempty"`
	GithubSender       `json:"sender,omitempty"`
	GithubRelease      `json:"release,omitempty"`
	GithubCommits      []GithubCommit `json:"commits,omitempty"`
	Action             string         `json:"action,omitempty"`
}
