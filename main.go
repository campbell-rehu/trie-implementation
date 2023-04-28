package main

import (
	"fmt"
)


type Trie struct {
	RootNode *Node
	Words []string
}

func (t *Trie) AddWord(word string) {
	node := t.RootNode
	for i := 0; i < len(word); i++ {
		isLastChar := i == len(word) - 1
		char := string(word[i])
		_, ok := node.Children[char]
		if (!ok) {
			// Add new Node if it doesn't exist
			node.Children[char] = CreateNewNode(char, make(map[string]*Node), isLastChar)
		}
		node = node.Children[char]
	}
}

func (t *Trie) GetAllWords() {
	node := t.RootNode
	for char, childNode := range node.Children {
		t.GetWord(char, childNode)
	}
}

func (t *Trie) GetWord(word string, node *Node) {
	if node.CompleteWord {
		t.Words = append(t.Words, word)
	}
	for char, childNode := range node.Children {
		t.GetWord(word + char, childNode)
	}
}

type Node struct {
	Char string
	Children map[string]*Node
	CompleteWord bool
}

func CreateNewNode(char string, children map[string]*Node, completeWord bool) *Node {
	return &Node{
		Char:         char,
		Children:     children,
		CompleteWord: completeWord,
	}
}

func main()  {
	trie := new(Trie)
	trie.RootNode = CreateNewNode("*", make(map[string]*Node), true)
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

	trie.GetAllWords()
	// trie.GetWord("c", trie.RootNode.Children["c"])
	fmt.Println(trie.Words)

	// TODO: Add keyboard autocomplete for the above words.

	// keysEvents, err := keyboard.GetKeys(10)
	// if err != nil {
    //     panic(err)
    // }
    // defer func() {
    //     _ = keyboard.Close()
    // }()

	// fmt.Println("Press ESC to quit")
    // for {
    //     event := <-keysEvents
    //     if event.Err != nil {
    //         panic(event.Err)
    //     }
    //     fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
    //     if event.Key == keyboard.KeyEsc {
    //         break
    //     }
	// 	char := string(event.Rune)
	// 	hasChild := node.HasChild(char)
	// 	if (hasChild) {
	// 		result := GetWords(node.GetChildNode(char))
	// 		fmt.Printf("Result: %s", result)
	// 	}
    // }
}

