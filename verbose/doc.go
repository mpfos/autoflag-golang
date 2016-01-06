// autoflag-golang/verbose project doc.go
// Copyright 2016 Michael Fabio

/*

autoflag-golang/verbose Provides automated verbose flag for golang programs

PURPOSE

Intended for command line utilities, and general Golang applications, this package
provides an easy way to add command line verbose flags to golang programs. When
imported the package adds to the list of program options stored by
autoflag-golang/help and sets up the flag and help handler.

INITIALIZATION

	On import and initialization, verbose.go installs a flag handler for the
	-verbose flag, installs a help entry in package help, and provides the
	public routines listed above.

TODO

	What's left todo? Shorter -v option as well?

*/
package verbose // import github.com/mpfos/autoflag-golang/verbose
