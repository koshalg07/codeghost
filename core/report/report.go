package report

import (
    "strings"
)

func Generate(lines []string) string {
    var b strings.Builder
    b.WriteString("# CodeGhost Report\n\n")
    b.WriteString("## Summary\n\n")
    b.WriteString("Potential root cause inferred from diff and logs.\n\n")
    b.WriteString("## Details\n")
    for _, line := range lines {
        b.WriteString("- " + line + "\n")
    }
    b.WriteString("\n## Suggested Fix\n\n> (Placeholder: Use heuristics or LLM in Phase 2)")
    return b.String()
}
