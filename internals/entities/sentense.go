package entities

import (
	"fmt"
	"math"
	"strings"
	"xulu/internals/consts"
)

type Sentence struct {
	TopSentense *Sentence
	Verb        string
	Names       []string
	SubSentense []*Sentence
}

func (s *Sentence) String() (result string) {
	action := "+"
	switch s.Verb {
	case "bcde":
		action = "-"
	case "dede":
		action = "*"
	}

	if len(s.SubSentense) > 0 {
		subSentenseS := []string{}
		for _, subSentense := range s.SubSentense {
			subSentenseS = append(subSentenseS, subSentense.String())
		}
		result += fmt.Sprintf("(%s)", strings.Join(subSentenseS, action))
	} else {
		result += fmt.Sprintf("(%s)", strings.Join(s.Names, action))
	}

	return
}

func (s *Sentence) Calculate() (result int) {
	if len(s.SubSentense) > 0 {
		for _, subSentense := range s.SubSentense {
			sebValue := subSentense.Calculate()

			switch s.Verb {
			case "abcd":
				result += sebValue
			case "bcde":
				result -= sebValue
			case "dede":
				if result != 0 {
					result *= sebValue
				} else {
					result = sebValue
				}
			}
		}
	} else {
		for _, name := range s.Names {
			nameValue := CalcName(name)
			switch s.Verb {
			case "abcd":
				result += nameValue
			case "bcde":
				result -= nameValue
			case "dede":
				if result != 0 {
					result *= nameValue
				} else {
					result = nameValue
				}
			}
		}
	}
	return
}

func CalcName(name string) (result int) {
	alphaValues := map[rune]int{}
	for _, a := range name {
		alphaValues[a] += consts.Alpha[a]
	}
	for _, value := range alphaValues {
		result += int(math.Pow(float64(value%5), 2))
	}
	return
}
