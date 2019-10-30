package main

func generatePotentials(beginWord string, words map[string]struct{}) []string {
	potentials := make([]string, 0, 1)

	chars := []rune(beginWord)
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		for j := 'a'; j <= 'z'; j++ {
			if j == c {
				continue
			}

			chars[i] = j
			newWord := string(chars)
			if _, exist := words[newWord]; exist {
				potentials = append(potentials, newWord)
			}
		}
		chars[i] = c
	}

	return potentials
}

func findLadders(beginWord, endWord string, wordList []string) [][]string {
	words := make(map[string]struct{}, len(wordList))
	for _, w := range wordList {
		words[w] = struct{}{}
	}
	start := make(map[string]struct{}, 1)
	start[beginWord] = struct{}{}
	delete(words, beginWord)

	potentials := make(map[string][]string)
	var ans [][]string

	for len(start) > 0 {
		if _, exist := start[endWord]; exist {
			break
		}

		s := make(map[string]struct{})
		for w := range start {
			p := generatePotentials(w, words)
			if len(p) > 0 {
				potentials[w] = p

				for _, pword := range p {
					s[pword] = struct{}{}
				}
			}
		}
		for w := range s {
			delete(words, w)
		}
		start = s
	}

	path := make([]string, 0, 1)
	path = append(path, beginWord)
	dfs(beginWord, endWord, potentials, &ans, path)

	return ans
}

func copySlice(s []string) []string {
	r := make([]string, len(s), len(s))
	for i, v := range s {
		r[i] = v
	}

	return r
}

func dfs(beginWord, endWord string, potentials map[string][]string, ans *[][]string, path []string) {
	if beginWord == endWord {
		*ans = append(*ans, copySlice(path))
		return
	}

	if p := potentials[beginWord]; p != nil {
		for _, w := range p {
			path = append(path, w)
			dfs(w, endWord, potentials, ans, path)
			path = path[:len(path)-1]
		}
	}
}
