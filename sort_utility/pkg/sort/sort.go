package sort

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func SortFile(args ...string) {
	if len(args) == 0 {
		in := bufio.NewReader(os.Stdin)
		str, err := in.ReadString('\n')
		if err != nil {
			logrus.Fatalf("input error:%v", err)
		}
		fmt.Println(standardSort(str))
		return
	}
	str := ""
	col := -1

	fs := flag.NewFlagSet("Flags for sort", flag.ContinueOnError)
	fs.IntVar(&col, "k", -1, "number of col, where we sort")
	reverse := fs.Bool("r", false, "sorting order")
	uniq := fs.Bool("u", false, "uniqueness")
	byInt := fs.Bool("n", false, "sorted by int")
	fs.Parse(args)

	file, err := os.Open(fs.Arg(0))
	if err != nil {
		logrus.Errorf("error with open file: %v", err)
	}
	scanner := bufio.NewScanner(file)

	if col == -1 {
		for scanner.Scan() {
			line := scanner.Text()
			str += line
		}
		if err := scanner.Err(); err != nil {
			logrus.Fatal(err)
		}
		if !(*byInt) {
			str = standardSort(str)
		} else {
			str = sortByInt(str)
		}
	} else {
		var strSl []string
		for scanner.Scan() {
			line := scanner.Text()
			strSl = append(strSl, line)
		}
		if !(*byInt) {
			str = sortByCol(strSl, col-1)
		} else {
			str = sortByColByInt(strSl, col-1)
		}
	}

	if *uniq {
		str = uniqStr(str)
	}
	if *reverse {
		str = reverseStr(str)
	}

	file.Close()
	err = os.Remove(fs.Arg(0))
	if err != nil {
		logrus.Error(err)
	}
	file, err = os.Create(fs.Arg(0))
	if err != nil {
		logrus.Fatal(err)
	}
	file.Write([]byte(str))
}

type typeForSortByColByInt struct {
	Num int
	str []string
}

func sortByColByInt(str []string, col int) string {
	var work [][]string
	var first [][]string
	for _, s := range str {
		tmp := strings.Split(s, " ")
		if len(tmp) <= col {
			first = append(first, tmp)
		} else {
			work = append(work, tmp)
		}
	}
	var withInt, onlyWords [][]string
	for _, w := range work {
		if strings.ContainsAny(w[col], "1234567890") {
			withInt = append(withInt, w)
		} else {
			onlyWords = append(onlyWords, w)
		}
	}

	sortColInt(withInt, col)
	work = append(withInt, onlyWords...)
	var strSl []string
	work = append(first, work...)
	for _, s := range work {
		tmp := strings.Join(s, " ")
		strSl = append(strSl, tmp)
	}
	return strings.Join(strSl, "\n")
}
func sortColInt(strSl [][]string, col int) {
	var tmpSl []typeForSortByColByInt
	for _, str := range strSl {
		num := 0
		fl := true
		for i, r := range str[col] {
			if r >= '0' && r <= '9' {
				n, _ := strconv.Atoi(string(r))
				num = num*10 + n
				fl = false
			} else if !fl {
				tmpSl = append(tmpSl, typeForSortByColByInt{Num: num, str: str})
				break
			}
			if i == len(str[col])-1 {
				tmpSl = append(tmpSl, typeForSortByColByInt{Num: num, str: str})
			}
		}
	}
	sort.Slice(tmpSl, func(i, j int) bool {
		return tmpSl[i].Num < tmpSl[j].Num
	})
	for i, v := range tmpSl {
		strSl[i] = v.str
	}
}
func sortByCol(str []string, col int) string {
	var work [][]string
	var first [][]string
	for _, s := range str {
		tmp := strings.Split(s, " ")
		if len(tmp) <= col {
			first = append(first, tmp)
		} else {
			work = append(work, tmp)
		}
	}
	sort.Slice(work, func(i, j int) bool {
		return work[i][col] < work[j][col]
	})
	var strSl []string
	work = append(first, work...)
	for _, s := range work {
		tmp := strings.Join(s, " ")
		strSl = append(strSl, tmp)
	}
	return strings.Join(strSl, "\n")
}
func sortByInt(str string) string {
	var withInt []string
	var onlyWords []string
	str = strings.ReplaceAll(str, "  ", " ")
	str = strings.Trim(str, "\n")
	strSl := strings.Split(str, " ")
	for _, val := range strSl {
		if strings.ContainsAny(val, "1234567890") {
			withInt = append(withInt, val)
		} else {
			onlyWords = append(onlyWords, val)
		}
	}
	sortInt(withInt)
	out := append(withInt, onlyWords...)
	return strings.Join(out, " ")
}

type myType struct {
	Num int
	str string
}

func sortInt(strSl []string) {
	var sl []myType
	for _, str := range strSl {
		num := 0
		fl := true
		for i, r := range str {
			if r >= '0' && r <= '9' {
				n, _ := strconv.Atoi(string(r))
				num = num*10 + n
				fl = false
			} else if !fl {
				sl = append(sl, myType{Num: num, str: str})
				break
			}
			if i == len(str)-1 {
				sl = append(sl, myType{Num: num, str: str})
			}
		}
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].Num < sl[j].Num
	})
	for i, v := range sl {
		strSl[i] = v.str
	}
}
func reverseStr(str string) string {
	var out []string
	strSl := strings.Split(str, " ")
	for i := len(strSl) - 1; i >= 0; i-- {
		out = append(out, strSl[i])
	}
	return strings.Join(out, " ")
}

func uniqStr(str string) string {
	var out []string
	strSl := strings.Split(str, " ")
	last := strSl[0]
	out = append(out, last)
	for _, v := range strSl {
		if last != v {
			out = append(out, v)
			last = v
		}
	}
	return strings.Join(out, " ")
}

func standardSort(str string) string {
	str = strings.ReplaceAll(str, "  ", " ")
	str = strings.Trim(str, "\n")
	strSl := strings.Split(str, " ")
	sort.Strings(strSl)
	str = strings.Join(strSl, " ")
	return str
}
