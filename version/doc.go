// autoflag-golang/version project doc.go
// Copyright 2016 Michael Fabio

/*

autoflag-golang/version Provides automated version flag for golang programs

PURPOSE

Intended for command line utilities, and general Golang applications, this package
provides an easy way to add command line version flags to golang programs. When
imported the package adds to the list of program options stored by
autoflag-golang/help and sets up the flag and help handler.

CONSTANTS

	Limit_Major - maximum limit for Major component
	Limit_Minor - maximum limit for Minor component
	Limit_Subminor - maximum limit for Subminor component

TYPES

	V - a version number consisting of Major, Minor, and Subminor components

FUNCTIONS

	New - returns a new version number of type V

METHODS

	Next - returns the next version number from an existing one
	Previous - returns the previous version number from an existing one
	Print - prints a version number with a title string
	String - returns a string from a version number

INITIALIZATION

	On import and initialization, version.go installs a flag handler for the
	-version flag, installs a help entry in package help, and provides the
	public routines listed above.

TODO

	Test the increment and decrement ops (Next, Previous)
	Provide a Max function
	What's left todo? Shorter -v option as well?
	Additional version number/string handling routines
	Comparison, equality, etc...

*/
package version // import github.com/mpfos/autoflag-golang/version
