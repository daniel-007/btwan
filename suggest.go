package btwan

import (
	"github.com/fvbock/trie"
)

var tree = trie.NewTrie()
var _suggestChan = make(chan string, 100)

func initSuggest() {
	loadSuggest()
	go loopSuggest()
}

func prefixSuggest(prefix string) []string {
	return tree.PrefixMembersList(prefix)
}

func loopSuggest() {
	var count = 0
	for q := range _suggestChan {
		if count >= 1000 {
			dumpSuggest()
			count = 0
		}
		tree.Add(q)
		count++
	}
}

func dumpSuggest() {
	err := tree.DumpToFile(workdir + "/suggest/suggest.db")
	if err != nil {
		fatal(err)
	}
}

func loadSuggest() {
	err := tree.MergeFromFile(workdir + "/suggest/suggest.db")
	if err != nil {
		fatal(err)
	}
}
