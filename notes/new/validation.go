package new

import "strings"

func validateInput(input NewNoteInput) (NewNoteInput, map[string]error) {
	errs := make(map[string]error)
	vContent, _ := validateContent(input.Content)
	input.Content = vContent
	return input, errs
}

func validateContent(content string) (string, error) {
	content = strings.TrimSpace(content)
	return content, nil
}
