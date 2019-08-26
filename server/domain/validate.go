package domain

import (
	"github.com/pkg/errors"
	"regexp"
)

var re = regexp.MustCompile(`^[0-9a-z._-]{3,128}$`)

func GetProblemsWithName(name string) (error) {
	if re.MatchString(name){
		return nil
	}

	return errors.New("Name must be 3-128 characters long and contain only a-z, '-', '.','_' and digits")


}
