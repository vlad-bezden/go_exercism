/*Package twofer
Two-fer or 2-fer is short for two for one. One for you and one for me.

"One for X, one for me."
When X is a name or "you".

If the given name is "Alice", the result should be "One for Alice,
one for me." If no name is given, the result should be
"One for you, one for me."
package twofer
*/
import "fmt"

// ShareWith return X is a name or "you" for "One for X, one for me."
func ShareWith(name string) string {
	var p string
	if 0 < len(name) {
		p = name
	} else {
		p = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", p)
}
