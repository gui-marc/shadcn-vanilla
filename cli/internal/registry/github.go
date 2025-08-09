package registry

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type GitHubRegistry struct {
	Owner  string
	Repo   string
	Branch string
}

func NewGitHubRegistry(owner, repo, branch string) *GitHubRegistry {
	return &GitHubRegistry{Owner: owner, Repo: repo, Branch: branch}
}

// FetchFile downloads the raw file content from the GitHub repo
func (r *GitHubRegistry) FetchFile(ctx context.Context, path string) ([]byte, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", r.Owner, r.Repo, r.Branch, path)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch file %s: %s", path, resp.Status)
	}
	return io.ReadAll(resp.Body)
}
