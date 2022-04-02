package main

import (
	"fmt"
	"strings"
	"xulu/internals/consts"
	"xulu/internals/entities"
	"xulu/internals/utils"
)

func main() {
	fmt.Println("Write your xulu sentense:")

	sen := utils.ReadLine()
	words := strings.Split(string(sen), " ")
	var root *entities.Sentence
	current := &entities.Sentence{}

	for inx, word := range words {
		if utils.IsWordVerb(word) {
			if inx == 0 {
				// this is root sentense
				newSentense := &entities.Sentence{
					TopSentense: nil,
					Verb:        word,
					SubSentense: []*entities.Sentence{},
				}
				current = newSentense
				root = newSentense
				continue
			}

			// create sub sentense and appned it to topLevel / current
			newSentense := &entities.Sentence{
				TopSentense: current,
				Verb:        word,
				SubSentense: []*entities.Sentence{},
			}
			if utils.IsWordVerb(words[inx-1]) {
				current.SubSentense = append(current.SubSentense, newSentense)
			} else {
				current.TopSentense.SubSentense = append(current.TopSentense.SubSentense, newSentense)
			}
			current = newSentense

			continue
		}
		if inx == 0 {
			// sentense started with nonVerb and its not valid sentense
			panic(consts.Invalid_Sentense)
		}
		current.Names = append(current.Names, word)
	}

	fmt.Println(consts.Sentence, root.String())
	fmt.Println(consts.Total, root.Calculate())
}
