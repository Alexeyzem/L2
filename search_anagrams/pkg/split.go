package pkg

import (
	"sort"
	"strings"
)

func Split(StrSl []string) map[string][]string {
	if StrSl == nil {
		return nil
	}
	out := make(map[string][]string)
	for _, v := range StrSl {
		flag := false
		for k, _ := range out {
			if check(v, k) {
				out[strings.ToLower(k)] = append(out[strings.ToLower(k)], v)
				flag = true
			}
		}
		if !flag {
			out[v] = append(out[v], v)
		}
	}

	combineOtherWords(out)
	sortAndDelete(out)
	return out
}
func check(first, second string) bool {
	if len(second) != len(first) {
		return false
	}
	firstSl, secondSl := []rune(strings.ToLower(first)), []rune(strings.ToLower(second))
	sort.Slice(firstSl, func(i, j int) bool {
		return firstSl[i] < firstSl[j]
	})
	sort.Slice(secondSl, func(i, j int) bool {
		return secondSl[i] < secondSl[j]
	})
	for i := range firstSl {
		if firstSl[i] != secondSl[i] {
			return false
		}
	}
	return true
}
func combineOtherWords(m map[string][]string) {
	flag := true
	keyForOther := ""
	for k, v := range m {
		if len(v) == 1 && flag {
			flag = false
			keyForOther = k
		} else if len(v) == 1 {
			delete(m, k)
			m[keyForOther] = append(m[keyForOther], v[0])
		}
	}
}
func sortAndDelete(m map[string][]string) {
	for k, v := range m {
		sort.Strings(v)
		if len(v) == 1 {
			delete(m, k)
			continue
		}
		last := strings.ToLower(v[0])
		now := []string{last}
		for i := 1; i < len(v); i++ {
			if last != strings.ToLower(v[i]) {
				now = append(now, strings.ToLower(v[i]))
				last = strings.ToLower(v[0])
			}
		}
		if len(now) > 1 {
			m[k] = now
		} else {
			delete(m, k)
		}
	}
}
