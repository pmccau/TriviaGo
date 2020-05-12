package processor

import (
	"github.com/pmccau/TriviaGo/server/data"
	//"regexp"
	"strings"
)

// getShortAns returns a shortened version which is the original
// less any preceding words (an, the, a)
func getShortAns(ans string) string {
	if strings.Contains(ans, "the") ||
		strings.Contains(ans, "a") ||
		strings.Contains(ans, "an") {
		temp := strings.Split(ans, " ")

		if len(temp) > 1 {
			return strings.Join(temp[1:], " ")
		}
	}
	return ans
}

// CheckResponse checks if a response to a question is a match. This
// needs to be refactored to do some more serious regex, but at the moment
// will clean extra spaces and check for preceding words that don't impact meaning
func CheckResponse(Response string, q data.Question) bool {
	ans := strings.ToLower(q.Answer)
	shortAns := getShortAns(ans)
	Response = strings.ToLower(strings.TrimSpace(Response))

	return Response == ans || Response == shortAns
}
