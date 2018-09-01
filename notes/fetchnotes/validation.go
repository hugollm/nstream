package fetchnotes

import (
	"errors"
)

func validateInput(input fetchInput) (errs map[string]error) {
	errs = make(map[string]error)
	if input.Start.IsZero() {
		errs["start"] = errors.New("Start date is required.")
	}
	if input.End.IsZero() {
		errs["end"] = errors.New("End date is required.")
	}
	if len(errs) > 0 {
		return errs
	}
	if input.End.Before(input.Start) || input.End.Equal(input.Start) {
		errs["end"] = errors.New("End date must come after start.")
	}
	return errs
}
