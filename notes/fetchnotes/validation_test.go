package fetchnotes

import (
	"testing"
	"time"
)

func TestStartIsRequired(t *testing.T) {
	input := fetchInput{time.Time{}, time.Now()}
	errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["start"].Error() != "Start date is required." {
		t.Fail()
	}
}

func TestEndIsRequired(t *testing.T) {
	input := fetchInput{time.Now(), time.Time{}}
	errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["end"].Error() != "End date is required." {
		t.Fail()
	}
}

func TestEndMustComeAfterStart(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	input := fetchInput{now, yesterday}
	errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["end"].Error() != "End date must come after start." {
		t.Fail()
	}
}

func TestDatesMustNotBeEqual(t *testing.T) {
	now := time.Now()
	input := fetchInput{now, now}
	errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["end"].Error() != "End date must come after start." {
		t.Fail()
	}
}
