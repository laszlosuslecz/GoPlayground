// Package twofer : package comment
package twofer

// ShareWith : returns the given string or the given string with the passed argument
func ShareWith(x string) string {

	if len(x) == 0 {
		return "One for you, one for me."
	}
	return "One for " + x + ", one for me."
}
