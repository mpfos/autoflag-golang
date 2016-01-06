// autoflag-golang/help project doc.go
// Copyright 2015 Michael Fabio

/*

autoflag-golang/help.go Provides automated help flags for Golang programs

PURPOSE

Intended for command line utilities, and general Golang applications, this package
provides an easy way to add command line help flags to Golang programs. When
imported the package maintains a list of program options which can be updated by
the importing program.

When other packages use this package, they can add their own help entries at init
as they add their own command line flags to extend the application.

TYPES

	Entry - a "help entry" in the list of program options

FUNCTIONS

	SetPreamble - sets the leading text of the help page, to be displayed prior
to the list of help entries

	Add - adds a "help entry" of type Entry to the list of program options

	AddEntry - adds a pair of strings to the list of program options

	FormatEntries - returns a string from the list of program options

	PrintHelp - prints the preamble text and the formatted entries

INITIALIZATION

	On import and initialization, help.go adds a "-help" boolean flag to the
	flags the program will accept, and preps the list of help entries with a
	help entry about the help flag.

TODO

	Provide a short usage message routine in addition to the full manpage.

*/
package help // import github.com/mpfos/autoflag-golang/help
