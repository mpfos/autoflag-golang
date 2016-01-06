# autoflag-golang
This is a set of packages to automate simple command line flags, like help, version, verbose, etc for golang programs.

USAGE

importing the packages into your golang program will enable command line flags, like this:

import (
	"github.com/mpfos/autoflag-golang/help"
	"github.com/mpfos/autoflag-golang/version"
	"github.com/mpfos/autoflag-golang/verbose"
)

which will enable the -help flag, the -version flag, and the -verbose flag.  It will also provide the facility to add more help entries to be displayed when the help flag is set on the command line,  and provide version number handling routines.