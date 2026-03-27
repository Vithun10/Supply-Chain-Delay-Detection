// Package dateparser provides flexible date parsing for the supply-chain CSV dataset.
// It tries multiple common date layouts until one succeeds.
package dateparser

import (
	"fmt"
	"time"
)

// supportedLayouts lists all date formats the dataset may use.
var supportedLayouts = []string{
	"2006-01-02",
	"01/02/2006",
	"02-01-2006",
	"2006/01/02",
	"Jan 2, 2006",
	"January 2, 2006",
	"2 Jan 2006",
}

// Parse attempts to parse s against each supported layout and returns the first
// successful result.  An error is returned only if every layout fails.
func Parse(s string) (time.Time, error) {
	for _, layout := range supportedLayouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("dateparser: cannot parse %q with any known layout", s)
}
