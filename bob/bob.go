/* Bob is a lackadaisical teenager. In conversation, his responses are very limited.
Bob answers 'Sure.' if you ask him a question.
He answers 'Whoa, chill out!' if you yell at him.
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying anything.
He answers 'Whatever.' to anything else.
*/

package bob

import (
	"regexp"
	"strings"
)

const (
	sure     = "Sure."
	comedown = "Calm down, I know what I'm doing!"
	chill    = "Whoa, chill out!"
	fine     = "Fine. Be that way!"
	whatever = "Whatever."
)

// Hey converts remark message to type of response
func Hey(remark string) string {
	remark = strings.Trim(remark, "\t\n\r ")
	isRemark := len(remark) == 0
	isWord := regexp.MustCompile(`[a-zA-Z]+`).MatchString(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isShouting := isWord && (remark == strings.ToUpper(remark))
	isYell := isShouting && isQuestion

	switch {
	case isRemark:
		return fine
	case isYell:
		return comedown
	case isQuestion:
		return sure
	case isShouting:
		return chill
	default:
		return whatever
	}
}
