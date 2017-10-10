package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

func main() {
	res := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	fmt.Println(res)
	heap.Init()
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	for i, oldWord := range wordList {
		if oldWord == string(endWord) {
			break
		}
		if i == len(wordList)-1 {
			return 0
		}
	}
	visited := map[string]struct{}{}
	wordPip := list.New()
	wordPip.PushBack(beginWord)
	steps := list.New()
	steps.PushBack(1)
	for wordPip.Len() != 0 {
		step := steps.Remove(steps.Front()).(int)
		word := wordPip.Remove(wordPip.Front()).(string)
		visited[word] = struct{}{}
		for i := 0; i < len(word); i++ {
			for c := 'a'; c < 'z'; c++ {
				byteWord := []byte(word)
				byteWord[i] = byte(c)
				if string(byteWord) == endWord {
					return step + 1
				}
				for _, oldWord := range wordList {
					if oldWord == string(byteWord) {
						if _, ok := visited[string(byteWord)]; !ok {
							steps.PushBack(step + 1)
							wordPip.PushBack(string(byteWord))
						}
					}
				}
			}
		}
	}
	return 0
}
