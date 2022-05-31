package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type ProjectedDateValue struct {
	Date  *time.Time
	IsSet *bool
}

// NewProjectedDate creates a new instance of a ProjectedDateValue, allocating memory for both properties which can then be set using the Value pattern
func NewProjectedDate() *ProjectedDateValue {
	var t time.Time
	var b bool = false
	return &ProjectedDateValue{
		Date:  &t,
		IsSet: &b,
	}
}

// Type implements pflag.Value
func (v ProjectedDateValue) Type() string {
	return "ProjectedDateValue"
}

var (
	reExtendedDuration = regexp.MustCompile(`^(\d+)([my])$`)
)

// Set implements pflag.Value
func (v ProjectedDateValue) Set(s string) error {
	m := reExtendedDuration.FindStringSubmatch(s)
	if m == nil || len(m) == 0 {
		return fmt.Errorf("not a valid extended duration; must end in m (months), y (years)")
	}
	if n, err := strconv.Atoi(m[1]); err != nil {
		return err
	} else {
		now := time.Now()
		nowLocation := now.Location()
		var d time.Time
		switch m[2] {
		case "m":
			newMonth := int(now.Month()) + n
			q := newMonth / 12
			r := newMonth % 12
			d = time.Date(now.Year()+q, time.Month(r), now.Day(), 0, 0, 0, 0, nowLocation)
		case "y":
			d = time.Date(now.Year()+n, now.Month(), now.Day(), 0, 0, 0, 0, nowLocation)
		default:
			return fmt.Errorf("not a valid extended duration flag")
		}
		*v.Date = d
		*v.IsSet = true
	}
	return nil
	//}
}
