package client

type LoginCredential struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

type Repo struct {
	// User ...
	User string `json:"user"`
	// Name of the repo
	Name string `json:"name"`
	// Namespace of the repo
	Namespace string `json:"namespace"`
	// RepoType is type of the repo, e.g. 'image'
	RepoType string `json:"repository_type"`
	// Status ...
	Status int `json:"status"`
	// Description ...
	Description string `json:"description"`
	// IsPrivate indicates whether the repo is private
	IsPrivate bool `json:"is_private"`
	// IsAutomated ...
	IsAutomated bool `json:"is_automated"`
	// CanEdit ...
	CanEdit bool `json:"can_edit"`
	// StarCount ..
	StarCount int `json:"star_count"`
	// PullCount ...
	PullCount int `json:"pull_count"`
}

// ReposResp is response of repo list request
type ReposResp struct {
	// Count is total number of repos
	Count int `json:"count"`
	// Next is the URL of the next page
	Next string `json:"next"`
	// Previous is the URL of the previous page
	Previous string `json:"previous"`
	// Repos is repo list
	Repos []Repo `json:"results"`
}

// Tag describes a tag in DockerHub
type Tag struct {
	// Name of the tag
	Name string `json:"name"`
	// FullSize is size of the image
	FullSize int64 `json:"full_size"`
}

// TagsResp is response of tag list request
type TagsResp struct {
	// Count is total number of repos
	Count int `json:"count"`
	// Next is the URL of the next page
	Next string `json:"next"`
	// Previous is the URL of the previous page
	Previous string `json:"previous"`
	// Repos is tags list
	Tags []Tag `json:"results"`
}

type TokenResp struct {
	Token string `json:"token"`
}