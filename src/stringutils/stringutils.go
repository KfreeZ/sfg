package stringutils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MyAtoi(str string) int {
	pos := 0
	startPos := 0
	endPos := 0
	length := len(str)
	negative := 1
	abort := false
	start := false
	finish := false

	for pos < length {

		fmt.Println(str[pos])

		if finish == true || abort == true {
			break
		}

		if str[pos] == '-' || str[pos] == '+' {
			fmt.Println("symbol")
			if start {
				fmt.Println("1")
				finish = true
				endPos = pos - 1
				break
			} else {
				fmt.Println("2")
				start = true
				if str[pos] == '-' {
					negative = -1
				} else {
					negative = 1
				}
				startPos = pos + 1
				pos++
				continue
			}
		} else if str[pos] == ' ' {
			fmt.Println("space")
			if start {
				fmt.Println("3")
				finish = true
				endPos = pos - 1
				break
			} else {
				fmt.Println("4")
				pos++
				continue
			}
		} else if str[pos] >= '0' && str[pos] <= '9' {
			fmt.Println("number")
			if !start {
				fmt.Println("5")
				start = true
				startPos = pos
			}
			if pos == length-1 {
				finish = true
				endPos = pos
			} else {
				pos++
				continue
			}
		} else {
			fmt.Println("other")
			if start {
				fmt.Println("6")
				finish = true
				endPos = pos - 1
				break
			} else {
				fmt.Println("7")
				abort = true
				break
			}
		}
	}

	fmt.Println(str[startPos:endPos+1], startPos, endPos, negative)

	if abort || !start {
		fmt.Println("abort or not start")
		return 0
	}

	absoluteN := 0
	for _, num := range str[startPos : endPos+1] {
		absoluteN = absoluteN*10 + int(num-'0')

		if negative == 1 && absoluteN > math.MaxInt32 {
			fmt.Println("int32 overflow")
			return math.MaxInt32
		}

		if negative == -1 && (absoluteN*-1) < math.MinInt32 {
			fmt.Println("int32 underflow")
			return math.MinInt32
		}
	}

	fmt.Println(str[startPos:endPos+1], startPos, endPos, negative, absoluteN)
	return absoluteN * negative

}

func StrLEn(S string) string {
	str := "名称Tom"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
	fmt.Println(len([]byte(str)))

	str = "123"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
	fmt.Println(len([]byte(str)))

	str = "abc"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
	fmt.Println(len([]byte(str)))

	return S[1:2]
}

func CompressString(S string) string {
	oldLen := len(S)

	if oldLen <= 2 {
		return S
	}

	ret := []byte{}
	pos := 0
	loopCnt := 1
	pre := S[0]
	for pos < oldLen {

		if pos == 0 {
			ret = append(ret, S[pos])
			pos++
			loopCnt = 1
			continue
		}

		if pre == S[pos] {
			loopCnt++
		} else {
			ret = append(ret, []byte(strconv.Itoa(loopCnt))...)
			ret = append(ret, S[pos])
			loopCnt = 1
			pre = S[pos]
		}

		pos++
	}

	ret = append(ret, []byte(strconv.Itoa(loopCnt))...)

	fmt.Print(string(ret))

	if len(ret) > oldLen {
		return S
	}
	return string(ret)

}
func CompressString2(S string) string {
	oldLen := len(S)

	if oldLen <= 2 {
		return S
	}

	ret := strings.Builder{}
	ret.Grow(50000)
	pos := 0
	loopCnt := 1
	pre := S[0]
	for pos < oldLen {

		if pos == 0 {
			ret.WriteString(string(pre))
			pos++
			loopCnt = 1
			continue
		}

		if pre == S[pos] {
			loopCnt++
		} else {
			ret.WriteString(strconv.Itoa(loopCnt))
			pre = S[pos]
			ret.WriteString(string(pre))
			loopCnt = 1
		}

		pos++
	}

	ret.WriteString(strconv.Itoa(loopCnt))

	fmt.Print()

	if ret.Len() > oldLen {
		return S
	}
	return ret.String()

}

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	wdLeft := 0
	wdRight := 1
	maxCnt := 1

	for wdRight <= len(s)-1 {
		if NoDuplica(s, wdRight, wdLeft) {
			if wdRight-wdLeft+1 > maxCnt {
				maxCnt = wdRight - wdLeft + 1
			}
			wdRight++
		} else {
			wdLeft++
		}
	}

	fmt.Println(wdLeft, wdRight)
	return maxCnt
}

func NoDuplica(s string, r int, l int) bool {
	for pos := l; pos < r; pos++ {
		if s[pos] == s[r] {
			return false
		}
	}
	return true
}

func FormAlpbet(s string) []int {
	retChars := make([]int, 26)
	for idx := 0; idx < len(s); idx++ {
		retChars[s[idx]-'a']++
	}

	return retChars
}

func CountCharacters(words []string, chars string) int {
	ret := 0

	aChars := FormAlpbet(chars)

	for _, w := range words {
		wChars := FormAlpbet(w)
		matchCnt := 0
		for idx := 0; idx < 26; idx++ {
			if wChars[idx] == 0 {
				continue
			}
			if aChars[idx] >= wChars[idx] {
				matchCnt += wChars[idx]
			} else {
				matchCnt = 0
				break
			}
		}

		ret += matchCnt
	}

	return ret

}
