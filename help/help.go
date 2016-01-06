// autoflag-golang/help project help.go
package help

/*
github.com/mpfos/autoflag-golang/help/help.go - provides a help flag to golang programs
*/

import (
	"errors"
	"flag"
	"fmt"
)

type Page struct {
	Name    string
	Usage   string
	Accepts string
	Outputs string
	Does    string
}

type Entry struct {
	Option      string
	Description string
}

var (
	ShowHelp    bool
	preamble    *Page
	helpEntries []Entry
)

func init() {
	e := Add("help", "list program options")

	flag.BoolVar(&ShowHelp, e.Option, false, e.Description)
}

func SetPreamble(page *Page) (err error) {
	if page != nil {
		preamble = page
		err = nil
	} else {
		err = errors.New("SetPreamble error: received a nil Page pointer")
	}
	return
}

func UnsetPreamble() (err error) {
	preamble = nil
	err = nil
	return
}

func PrintHelp() (err error) {
	if preamble != nil {
		fmt.Println(*(FormatPreamble()))
		fmt.Println(*(FormatEntries()))
		err = nil
	} else {
		err = errors.New("PrintHelp error: preamble not set before PrintHelp() called.")
	}
	return
}

func FormatPreamble() (pstr *string) {
	str := ""
	if len(preamble.Name) > 0 {
		str = str + preamble.Name + "\n\n"
	}
	if len(preamble.Usage) > 0 {
		str = str + "usage: " + preamble.Usage + "\n\n"
	}
	if len(preamble.Accepts) > 0 {
		str = str + "accepts: " + preamble.Accepts + "\n\n"
	}
	if len(preamble.Does) > 0 {
		str = str + "does: " + preamble.Does + "\n\n"
	}
	if len(preamble.Outputs) > 0 {
		str = str + "output: " + preamble.Outputs + "\n"
	}
	return &str
}

func FormatEntries() (pstr *string) {
	max := 0
	str := ""
	for i := range helpEntries {
		l := len(helpEntries[i].Option)
		if l > max {
			max = l
		}
	}
	if max > 0 {
		str = str + "OPTIONS\n\n"
		flagformat := fmt.Sprintf("-%%-%ds\t%%s\n", max)
		for i := range helpEntries {
			str = str + fmt.Sprintf(flagformat, helpEntries[i].Option, helpEntries[i].Description)
		}
	}
	return &str
}

func Add(opt string, desc string) (e Entry) {
	e = Entry{}
	e.Option = opt
	e.Description = desc
	AddEntry(e)
	return
}

func AddEntry(e Entry) {
	helpEntries = append(helpEntries, e)
}
