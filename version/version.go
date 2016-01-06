// version project version.go
package version

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mpfos/autoflag-golang/help"
)

var (

	//flag to indicate if -version was seen by the flag package
	ShowVersion bool

	//current revision of package "version"
	//this shows the simple way to store a version number in your main program
	//after using this package
	Version = V{0, 0, 1}
)

const (
	Limit_Major    = 999
	Limit_Minor    = 999
	Limit_Subminor = 999
)

type V struct {
	Major    int
	Minor    int
	Subminor int
}

func init() {
	e := help.Add("version", "display version information")

	flag.BoolVar(&ShowVersion, e.Option, false, e.Description)
}

func New(major int, minor int, subminor int) *V {
	v := V{Major: major, Minor: minor, Subminor: subminor}
	return &v
}

func (v *V) Next() (out *V, err error) {
	err = nil
	if v != nil {
		subm := (*v).Subminor
		min := (*v).Minor
		maj := (*v).Major

		subm++
		if subm > Limit_Subminor {
			subm = 0
			min++
			if min > Limit_Minor {
				min = 0
				maj++
				if maj > Limit_Major {
					out = nil
					err = errors.New("Next error: version number overflowed limits")
					return
				}
			}
		}
		newv := V{Major: maj, Minor: min, Subminor: subm}
		out = &newv
	} else {
		out = nil
		err = errors.New("Next error: method receiver was nil")
	}
	return
}

func (v *V) Previous() (out *V, err error) {
	err = nil
	if v != nil {
		subm := (*v).Subminor
		min := (*v).Minor
		maj := (*v).Major

		subm--
		if subm < 0 {
			subm = Limit_Subminor
			min--
			if min < 0 {
				min = Limit_Minor
				maj--
				if maj < 0 {
					out = nil
					err = errors.New("Previous error: version number underflowed limits")
					return
				}
			}
		}
		newv := V{Major: maj, Minor: min, Subminor: subm}
		out = &newv
	} else {
		out = nil
		err = errors.New("Next error: method receiver was nil")
	}
	return
}
func (v *V) Print(title string) (err error) {
	err = nil
	s := ""

	if v != nil {
		s, err = v.String(title)
		if err != nil {
			return
		}
		fmt.Println(s)
	} else {
		err = errors.New("Print error: method requires a non-nil method receiver")
	}
	return
}

func (v *V) String(title string) (out string, err error) {
	err = nil
	out = ""

	if v != nil {
		out = fmt.Sprintf("%s Version %d.%d.%d",
			title, (*v).Major, (*v).Minor, (*v).Subminor)
	} else {
		err = errors.New("String error: method requires a non-nil method receiver")
	}
	return
}
