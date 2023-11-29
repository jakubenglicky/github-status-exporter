# Github Status Exporter

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
        title: Github component '{{ $labels.component }}' is in state '{{ $labels.status }}'
        description: Github component '{{ $labels.component }}' is in state '{{ $labels.status }}' for more than 2 minutes. Github has some issue.
        summary: Github component '{{ $labels.component }}' is in state '{{ $labels.status }}'
```

## Metrics
```
# HELP github_component_status Status of Github component
# TYPE github_component_status gauge
github_component_status{component="API Requests",status="operational"} 1
github_component_status{component="Actions",status="operational"} 1
github_component_status{component="Codespaces",status="operational"} 1
github_component_status{component="Copilot",status="operational"} 1
github_component_status{component="Git Operations",status="operational"} 1
github_component_status{component="Issues",status="operational"} 1
github_component_status{component="Packages",status="operational"} 1
github_component_status{component="Pages",status="operational"} 1
github_component_status{component="Pull Requests",status="operational"} 1
github_component_status{component="Visit www.githubstatus.com for more information",status="operational"} 1
github_component_status{component="Webhooks",status="operational"} 1
```