package bob

import (
	"regexp"
	"strings"
)

type Sentence string

func (s Sentence) isYelling() bool {
	containsSmallLetter, _ := regexp.Match(`[a-z]`, []byte(s))
	containsCapitalLetter, _ := regexp.Match(`[A-Z]`, []byte(s))
	return !containsSmallLetter && containsCapitalLetter
}

func (s Sentence) isQuestion() bool {
	return strings.LastIndex(string(s), `?`) == len(s)-1
}

func (s Sentence) isYellingQuestion() bool {
	return s.isYelling() && s.isQuestion()
}

func Hey(remark string) string {
	s := Sentence(strings.TrimSpace(remark))
	if len(s) == 0 {
		return "Fine. Be that way!"
	}
	if s.isYellingQuestion() {
		return "Calm down, I know what I'm doing!"
	}
	if s.isQuestion() {
		return "Sure."
	}
	if s.isYelling() {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
