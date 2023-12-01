# Github Status Exporter

[![Docker](https://github.com/jakubenglicky/github-status-exporter/actions/workflows/build.yml/badge.svg?branch=main)](https://hub.docker.com/repository/docker/jakubenglicky/github-status-exporter/general)

Github Status Exporter for Prometheus.

## Usage:
```bash
docker run -p 8080:8080 jakubenglicky/github-status-exporter
```
Visit http://localhost:8080/metrics

## Alerts
```
- name: Github Status
  rules:
    - alert: GithubComponentIsDown
      expr: github_component_status == 0
      for: 2m
      labels:
        severity: warning
      annotations:
        title: Github Components Outages
        description: Github component '{{ $labels.component }}' has problem for more than 2 minutes. Visit www.githubstatus.com for more information.
        summary: Github component '{{ $labels.component }}' has problem.
```

## Metrics
```
# HELP github_component_status Status of Github component
# TYPE github_component_status gauge
github_component_status{component="API Requests"} 1
github_component_status{component="Actions"} 1
github_component_status{component="Codespaces"} 1
github_component_status{component="Copilot"} 1
github_component_status{component="Git Operations"} 1
github_component_status{component="Issues"} 1
github_component_status{component="Packages"} 1
github_component_status{component="Pages"} 1
github_component_status{component="Pull Requests"} 1
github_component_status{component="Visit www.githubstatus.com for more information"} 1
github_component_status{component="Webhooks"} 1
```