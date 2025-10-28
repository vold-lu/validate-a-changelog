package validator

import "strings"

type ValidationError struct {
	Issues []ValidationIssue `json:"issues"`
}

func (v *ValidationError) Error() string {
	var sb strings.Builder

	for _, issue := range v.Issues {
		sb.WriteString(issue.String() + "\n")
	}

	return sb.String()
}

func (v *ValidationError) pushIssue(version, section, error string) {
	v.Issues = append(v.Issues, ValidationIssue{
		Version: version,
		Section: section,
		Error:   error,
	})
}

func (v *ValidationError) hasIssues() bool {
	return len(v.Issues) > 0
}

type ValidationIssue struct {
	// Version contains the version where the error happens (when possible)
	Version string `json:"version"`
	// Section contains the section where the error happens (when possible)
	Section string `json:"section"`
	// Error is the human formatted error message
	Error string `json:"error"`
}

func (vi *ValidationIssue) String() string {
	var sb strings.Builder

	sb.WriteString("[")

	if vi.Version != "" {
		sb.WriteString("version: " + vi.Version + ", ")
	} else {
		sb.WriteString("version: n/a, ")
	}
	if vi.Section != "" {
		sb.WriteString("section: " + vi.Section)
	} else {
		sb.WriteString("section: n/a")
	}

	sb.WriteString("]: " + vi.Error)

	return sb.String()
}
