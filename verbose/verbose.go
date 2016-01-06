// autoflag-golang/verbose project verbose.go
package verbose

import (
	"flag"

	"github.com/mpfos/autoflag-golang/help"
)

var (
	ShowVerbose bool
)

func init() {
	e := help.Add("verbose", "verbose output - show more detail")

	flag.BoolVar(&ShowVerbose, e.Option, false, e.Description)
}
