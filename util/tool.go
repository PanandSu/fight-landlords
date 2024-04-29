package util

import (
	"sort"
	"strconv"
)

func Hash(b rune) int {
	switch b {
	case '3', '4', '5', '6', '7', '8', '9':
		n, _ := strconv.Atoi(string(b))
		return n
	case 'X':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	case '2':
		return 15
	case 'L':
		return 16
	case 'B':
		return 17
	}
	panic("invalid character")
}

func Unhash(n int) string {
	switch n {
	case 3, 4, 5, 6, 7, 8, 9:
		return strconv.Itoa(n)
	case 10:
		return "X"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	case 15:
		return "2"
	case 16:
		return "L"
	case 17:
		return "B"
	default:
		panic("Invalid num")
	}
}

// 顺序:A23456789XJQKLB
func SortedStr(s string) string {
	arr := make([]int, 0)
	for _, v := range s {
		n := Hash(v)
		arr = append(arr, n)
	}
	sort.Ints(arr)
	res := ""
	for _, v := range arr {
		b := Unhash(v)
		res += b
	}
	return res
}

// s1 in s2 ?
func ContainRunes(s1, s2 []rune) bool {
	s1 = []rune(SortedStr(string(s1)))
	s2 = []rune(SortedStr(string(s2)))
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	return i == len(s1)
}

// s1-s2,s1,s2皆升序
// len(res)=len(s1)-len(s2)
func RemoveRunes(s1, s2 []rune) []rune {
	s1 = []rune(SortedStr(string(s1)))
	s2 = []rune(SortedStr(string(s2)))
	res := make([]rune, 0)
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			j++
		} else {
			res = append(res, s1[i])
		}
		i++
	}
	res = append(res, s1[i:]...)
	return res
}
