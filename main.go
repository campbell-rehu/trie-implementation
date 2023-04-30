package main

import (
	"fmt"
	"regexp"
	"strings"

	triepkg "dev/trie-implementation/trie"

	"github.com/eiannone/keyboard"
)

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func main()  {
	trie := new(triepkg.Trie)
	trie.RootNode = triepkg.CreateNewNode("*", make(map[string]*triepkg.Node), true)
	trie.Words = make([]string, 2)
	trie.AddWord("car")
	trie.AddWord("card")
	trie.AddWord("cards")
	trie.AddWord("cot")
	trie.AddWord("cots")
	trie.AddWord("trie")
	trie.AddWord("tried")
	trie.AddWord("tries")
	trie.AddWord("try")
	trie.AddWord("magic")
	trie.AddWord("magician")
	trie.AddWord("monster")
	trie.AddWord("minefield")

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
        panic(err)
    }
    defer func() {
        _ = keyboard.Close()
    }()

	fmt.Println("Press ESC to quit")
	trie.GetAllWords()
	fmt.Println("All options: ", strings.Join(trie.Words, ", "))
	chars := make([]string, 0, 100)
    for {
        event := <-keysEvents
        if event.Err != nil {
            panic(event.Err)
        }
        if event.Key == keyboard.KeyEsc {
            break
        }
		if event.Key == keyboard.KeyBackspace2 {
			if len(chars) > 1 {
				chars = chars[:len(chars)-1]
				trie.ResetWords()
			}
			// TODO: handle backspacing until there are no letters
		}
		if IsLetter(string(event.Rune)) {
			chars = append(chars, string(event.Rune))
			trie.ResetWords()
		}
		hasAllChars := trie.HasAllChars(chars)
		if hasAllChars {
			word := strings.Join(chars, "")
			node := trie.RootNode.Children[chars[0]].GetChildNode(word)
			trie.GetWord(word, node)
		}
		// TODO: figure out way to more elegantly print the autocomplete options (maybe with colours)
		fmt.Println(strings.Join(chars, ""), "\n Autocomplete options: ", strings.Join(trie.Words, ", "))
    }
}



