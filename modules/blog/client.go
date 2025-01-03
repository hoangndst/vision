package blog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hoangndst/vision/server/util/logging"
)

type Client struct {
	githubToken string
}

func NewClient(githubToken string) *Client {
	return &Client{
		githubToken: githubToken,
	}
}

// GetBlogRaw gets the raw data of a blog from the GitHub repository.
func (c *Client) GetBlogRaw(ctx context.Context, path string) (string, error) {
	logger := logging.GetLogger(ctx)
	const url = "https://raw.githubusercontent.com/hoangndst/hoangndst-homepage/blog/%s"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(url, path), nil)
	if err != nil {
		logger.Error("error creating request", err)
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.githubToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("error response status code: ", resp.StatusCode)
		return "", err
	}
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("error reading response body: ", err)
		return "", err
	}
	return string(rawData), nil
}

// GetAllBlogs gets all blogs from the GitHub repository.
func (c *Client) GetAllBlogs(ctx context.Context) ([]Blog, error) {
	logger := logging.GetLogger(ctx)
	const url = "https://api.github.com/repos/hoangndst/hoangndst-homepage/git/trees/blog?recursive=1"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Error("error creating request", err)
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.githubToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("error response status code: ", resp.StatusCode)
		return nil, err
	}
	var fileTree BlogTree
	if err := json.NewDecoder(resp.Body).Decode(&fileTree); err != nil {
		logger.Error("error decoding response: ", err)
		return nil, err
	}

	var blogs []Blog
	for _, file := range fileTree.Tree {
		// if path not end with .mdx, skip
		if !strings.HasSuffix(file.Path, ".mdx") {
			continue
		}
		rawData, err := c.GetBlogRaw(ctx, file.Path)
		if err != nil {
			logger.Error("error getting raw data", err)
		}
		blogs = append(blogs, Blog{
			Path:    file.Path,
			RawData: rawData,
		})
	}
	return blogs, nil
}
