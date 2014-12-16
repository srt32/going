package validators

import (
	"fmt"
	"time"

	"github.com/markbates/going/validate"
)

type TimeIsBeforeTime struct {
	FirstName  string
	FirstTime  time.Time
	SecondName string
	SecondTime time.Time
}

func (v *TimeIsBeforeTime) IsValid(errors *validate.Errors) {
	if v.FirstTime.UnixNano() > v.SecondTime.UnixNano() {
		errors.Add(generateKey(v.FirstName), fmt.Sprintf("%s must be before %s.", v.FirstName, v.SecondName))
	}
}
