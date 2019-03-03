/*
Package twofer implements a single function for returning "One for X, one for me."
*/
package twofer

import "fmt"

/*
ShareWith accepts an input string and returns One for X, one for me.
*/
func ShareWith(name string) string {
	shareWith := ""
	if name != "" {
		shareWith = fmt.Sprintf("One for %s, one for me.", name)
	} else {
		shareWith = "One for you, one for me."
	}
	return shareWith
}
