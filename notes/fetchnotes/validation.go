package fetchnotes

import (
	"errors"
	"strings"
	"time"
)

func validateInput(jsInput jsonInput) (fetchInput, map[string]error) {
	ftInput := fetchInput{}
	errs := make(map[string]error)
	start, startErr := validateTime(jsInput.Start)
	if startErr != nil {
		errs["start"] = startErr
	}
	end, endErr := validateTime(jsInput.End)
	if endErr != nil {
		errs["end"] = endErr
	}
	if len(errs) > 0 {
		return ftInput, errs
	}
	ftInput.Start = start
	ftInput.End = end
	if ftInput.End.Before(ftInput.Start) || ftInput.End.Equal(ftInput.Start) {
		errs["end"] = errors.New("End date must come after start.")
	}
	return ftInput, errs
}

func validateTime(str string) (time.Time, error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return time.Time{}, errors.New("Time string is required.")
	}
	vTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Time{}, errors.New("Invalid RFC3339 time string.")
	}
	return vTime, nil
}
