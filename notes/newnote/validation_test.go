package newnote

import (
	"testing"
)

func TestValidateInputTrimsContent(t *testing.T) {
	input := NewNoteInput{"  Lorem ipsum.  \n"}
	input, errs := validateInput(input)
	if len(errs) > 0 || input.Content != "Lorem ipsum." {
		t.Fail()
	}
}

func TestContentIsTrimmed(t *testing.T) {
	content := "  Lorem ipsum.  \n"
	vContent, err := validateContent(content)
	if err != nil || vContent != "Lorem ipsum." {
		t.Fail()
	}
}
