package mapper

func Match(diffs, logs []string) []string {
    // Very naive match logic for now
    return append(diffs, logs...)
}
