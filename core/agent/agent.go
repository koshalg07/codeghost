package agent

import (
	"codeghost/core/diff"
	"codeghost/core/logs"
	"codeghost/core/report"
	"codeghost/core/mapper"
)

func Analyze(diffPath, logPath string) string {
	d := diff.Parse(diffPath)
	l := logs.Parse(logPath)
	if d == nil || l == nil {
		return "Failed to parse diff or log files"
	}
	events := mapper.Match(d, l)
	if events == nil {
		return "No matching events found"
	}
	return report.Generate(events)
}