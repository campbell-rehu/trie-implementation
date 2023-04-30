package trie

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

func (t *Trie) ResetWords() {
	t.Words = make([]string, 0, 10)
}

func (t *Trie) HasAllChars(chars []string) bool {
	node := t.RootNode
	valid := true
	for i, v := range chars {
		nextNode := t.HasChar(v, node)
		if (nextNode == nil) {
			if (i != len(chars)-1) {
				valid = false
			}
			break
		}
		node = nextNode

	}
	return valid
}

func (t *Trie) HasChar(char string, node *Node) *Node {
	v := node.Children[char]
	return v
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

func (n *Node) GetChildNode(word string) *Node {
	if len(word) == 1 {
		return n
	}
	return n.Children[string(word[1])].GetChildNode(word[1:])
}