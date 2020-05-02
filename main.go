/**
 	main package from here package get initialized.
	request is coming here
*/

package main

import "BFC/utilities"

// Main function to start a function
func main() {
	// initalize a global variable
	utilities.GetGlobalVariable()

	// initalize a route
	Init()
}
