package bitbucket

import "net/http"

//go:generate counterfeiter . Client

type Client interface {
	CurrentUser(*http.Client) (string, error)
	Teams(*http.Client, Role) ([]string, error)
	Projects(*http.Client) ([]string, error)
	Repository(client *http.Client, owner string, repository string) (bool, error)
}
