package fetchnotes

import (
	"nstream/data/mock"
	"testing"
	"time"
)

func TestValidInput(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	jsInput := jsonInput{yesterday.Format(time.RFC3339), now.Format(time.RFC3339)}
	ftInput, errs := validateInput(jsInput)
	if len(errs) > 0 || mock.HasZeroValues(ftInput) {
		t.Fail()
	}
}

func TestStartIsRequired(t *testing.T) {
	input := jsonInput{"", time.Now().Format(time.RFC3339)}
	_, errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["start"].Error() != "Time string is required." {
		t.Fail()
	}
}

func TestEndIsRequired(t *testing.T) {
	input := jsonInput{time.Now().Format(time.RFC3339), ""}
	_, errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["end"].Error() != "Time string is required." {
		t.Fail()
	}
}

func TestEndMustComeAfterStart(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	input := jsonInput{now.Format(time.RFC3339), yesterday.Format(time.RFC3339)}
	_, errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["end"].Error() != "End date must come after start." {
		t.Fail()
	}
}

func TestDatesMustNotBeEqual(t *testing.T) {
	sNow := time.Now().Format(time.RFC3339)
	input := jsonInput{sNow, sNow}
	_, errs := validateInput(input)
	if len(errs) != 1 {
		t.FailNow()
	}
	if errs["end"].Error() != "End date must come after start." {
		t.Fail()
	}
}

func TestTimeParsingWithValidTime(t *testing.T) {
	now := time.Now()
	sNow := now.Format(time.RFC3339)
	vTime, err := validateTime(sNow)
	if err != nil || vTime.IsZero() {
		t.Fail()
	}
}

func TestTimeParsingWithSpaces(t *testing.T) {
	vTime, err := validateTime("    ")
	if !vTime.IsZero() || err.Error() != "Time string is required." {
		t.Fail()
	}
}
