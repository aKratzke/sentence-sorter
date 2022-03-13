package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// 1. Get text from place.
	// 		Phase 1. static text.
	// 		Phase 2. Dynamic retrieval
	input := "Eat from dog's food. Toilet paper attack claws fluff everywhere meow miao french ciao litterbox making bread on the bathrobe so attack curtains yet lick arm hair. Rub against owner because nose is wet get suspicious of own shadow then go play with toilette paper. Leave fur on owners clothes vommit food and eat it again so take a deep sniff of sock then walk around with mouth half open destroy couch. Is good you understand your place in my world. Rub whiskers on bare skin act innocent stuff and things but immediately regret falling into bathtub but ooooh feather moving feather!        . Catch small lizards, bring them into house, then unable to find them on carpet. Oooo!  "

	//input := "zumberg Begged for amazing tests"
	printStringWithNote("input", input)

	// 2. create list of sentences per paragraph.
	//		Phase 1, use 1 paragraph.
	//		Phase 2, make sure to be able to order by paragraphs as well.

	words := strings.FieldsFunc(input, split)

	printArrayWithNote("List of words", words)

	// 3. Sort words.
	// 		Phase skip.
	// 		Phase 2 implement
	for k, _ := range words {
		sortWordsInSentence(&words[k])
	}

	printArrayWithNote("Sorted Sentences", words)

	paragraph := strings.Join(words, " ")

	printStringWithNote("Final output paragraph", paragraph)
}

func split(r rune) bool {
	return r == '.' || r == '?' || r == '!'
}

func sortWordsInSentence(s *string) {
	theseWords := strings.Split(*s, " ")

	sort.Slice(theseWords, func(i, j int) bool { return strings.ToLower(theseWords[i]) < strings.ToLower(theseWords[j]) })

	*s = strings.Join(theseWords, " ")
}

func printStringWithNote(note string, value string) {
	toPrint := fmt.Sprintf("%s: %s", note, value)
	fmt.Println(toPrint)
}

func printArrayWithNote(note string, value []string) {

	fmt.Println(note)
	printArray(value)
}

func printArray(s []string) {
	for k, v := range s {

		fmt.Println(k, v)
	}
}
