package logs

func Parse(path string) []string {
    // TODO: Parse JSON/plaintext logs, extract error messages/timestamps
    return []string{"[ERROR] nil pointer in foo.go line 42", "[WARN] retry failed in bar.go"}
}
