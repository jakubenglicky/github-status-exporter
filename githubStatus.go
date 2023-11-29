package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const GithubStatusSummaryUrl = "https://www.githubstatus.com/api/v2/summary.json"

type Response struct {
	Components []Component `json:"components"`
}

type Component struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (c Component) IsOperational() bool {
	return c.Status == "operational"
}

func GetGithubStatusComponents() ([]Component, error) {
	resp, err := http.Get(GithubStatusSummaryUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get github status: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read github status response: %w", err)
	}

	var result Response
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal github status response: %w", err)
	}

	return result.Components, nil
}
