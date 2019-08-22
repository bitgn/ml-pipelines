package domain

import (
	"github.com/pkg/errors"
	"regexp"
)

var re = regexp.MustCompile(`^[0-9A-Za-z._-]{2,38}$`)

func GetProblemsWithID(id string) (error) {
	if re.MatchString(id){
		return nil
	}

	return errors.New("ID must be 3-38 characters long and contain only A-Z,a-z, '-', '.','_' and digits")


}
