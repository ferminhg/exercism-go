// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

type Checker func() (string, bool)
type AnswerChain struct{
	handlers []Checker
	question string
}

func (c *AnswerChain) GetAnswer() string {
	for _, checker := range c.handlers {
		answer, next := checker()
		if !next {
			return answer
		}
	}
	return "Whatever."
}

func (c *AnswerChain) IsYeildQuestion() (string, bool) {
	if c.question[(len(c.question)-1):] == "?" &&
		c.question == strings.ToUpper(c.question){
		return "Calm down, I know what I'm doing!", false
	}
	return "", true
}

func (c *AnswerChain) IsYield() (string, bool){
	re := regexp.MustCompile(`[a-z].`)
	hasLetter := re.Find([]byte(strings.ToLower(c.question))) != nil
	if c.question == strings.ToUpper(c.question) && hasLetter {
		return "Whoa, chill out!", false
	}
	return "", true
}

func (c *AnswerChain) IsQuestion() (string, bool) {
	if c.question[(len(c.question)-1):] == "?"{
		return "Sure.", false
	}
	return "", true
}

func NewAnswerChain(remark string) AnswerChain {
	chain := AnswerChain{
		handlers: []Checker{},
		question: remark,
	}
	chain.handlers = []Checker{
		chain.IsYeildQuestion,
		chain.IsYield,
		chain.IsQuestion,
	}
	return chain
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	chain := NewAnswerChain(remark)
	return chain.GetAnswer()
}
