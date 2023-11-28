package main

import "github.com/prometheus/client_golang/prometheus"

type Monitor struct {
	Registry                *prometheus.Registry
	GithubComponentStatus *prometheus.GaugeVec
}

func NewMonitor() *Monitor {
	reg := prometheus.NewRegistry()
	monitor := &Monitor{
		Registry: reg,

		GithubComponentStatus: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "github_component_status",
			Help: "Status of Github component",
		}, []string{"component", "status"}),
	}

	reg.MustRegister(monitor.GithubComponentStatus)

	return monitor
}
