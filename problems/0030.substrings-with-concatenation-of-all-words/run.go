package words

import (
	"fmt"
	"sort"
	"strings"
)

// 想先找到words中的所有子串，然后寻找s中字串的第一个位置
func findSubstring1(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	var (
		a      array
		length []int
	)
	a.permutation(words, 0, len(words))
	max := len(a.eles[0])
	for i := 0; i < len(s); i++ {
		fmt.Println(s)

		for _, v := range a.eles {
			l := strings.Index(s, v)
			if l > -1 {
				length = append(length, l)
			}
		}
		if len(s) == max {
			break
		}
		s = s[max:]
		fmt.Println(length)
	}
	sort.Ints(length)
	for i := len(length) - 1; i > 0; i-- {
		if length[i] == length[i-1] {
			length = append(length[:i-1], length[i:]...)
		}
	}
	return length
}

type array struct {
	eles []string
}

func (a *array) permutation(words []string, l, h int) {

	if l == h-1 {
		a.eles = append(a.eles, strings.Join(words, ""))
	} else {
		for i := l; i < h; i++ {
			words[i], words[l] = words[l], words[i]
			a.permutation(words, l+1, h)
			words[i], words[l] = words[l], words[i]

		}
	}
}

func findSubstring(s string, words []string) []int {
	var r []int
	if len(s) == 0 || len(words) == 0 {
		return r
	}
	n := len(words)
	m := len(words[0])
	tmap := make(map[string]int)
	for _, word := range words {
		tmap[word]++
	}

	ls := len(s)
	for i := 0; i <= ls-n*m; i++ {

		cmap := make(map[string]int, len(tmap))
		for t, v := range tmap {
			cmap[t] = v
		}

		// cmap := tmap
		k, j := n, i
		for k > 0 {
			str := s[j : j+m]
			if i == 9 {
				fmt.Println("s: ", str, cmap)
			}
			if v, ok := cmap[str]; !ok || v < 1 {
				break
			}

			cmap[str]--
			k--
			j += m
		}
		if k == 0 {
			r = append(r, i)
		}
	}
	return r
}

func findSubstring2(s string, words []string) []int {
	lens := len(s)
	res := make([]int, 0, lens)

	lenws := len(words)
	if lens == 0 || lenws == 0 || lens < lenws*len(words[0]) {
		return res
	}

	lenw := len(words[0])

	// record 记录 words 中每个单词总共出现的次数
	record := make(map[string]int, lenws)
	for _, w := range words {
		if len(w) != lenw {
			// 新添加的 example 2 出现了 words 中单词长度不一致的情况。
			// 这个违反了假设
			// 直接返回 res
			return res
		}
		record[w]++
	}

	// remain 记录 words 中每个单词还能出现的次数
	remain := make(map[string]int, lenws)
	// count 记录符合要求的单词的连续出现次数
	count := 1 // count 的初始值只要不为 0 就可以
	left, right := 0, 0

	/**
	 * s[left:right] 作为一个窗口存在
	 * 假设 word:= s[right:right+lenw]
	 * 如果 remain[word]>0 ，那么移动窗口 右边，同时更新移动后 s[left:right] 的统计信息
	 * 		remain[word]--
	 * 		right += lenw
	 * 		count++
	 * 否则，移动窗口 左边，同时更新移动后 s[left:right] 的统计信息
	 * 		remain[s[left:left+lenw]]++
	 * 		count--
	 * 		left += lenw
	 *
	 * 每次移动窗口 右边 后，如果 count = lenws ，那么
	 * 		说明找到了一个符合条件的解
	 * 		append(res, left)，然后
	 * 		移动窗口 左边
	 */

	// reset 重置 remain 和 count
	reset := func() {
		if count == 0 {
			// 统计记录没有被修改，不需要重置
			// 因为有这个预判，所以需要第一次执行 reset 时
			// count 的值不能为 0
			// 即 count 的初始值不能为 0
			return
		}
		for k, v := range record {
			remain[k] = v
		}
		count = 0
	}

	// moveLeft 让 left 指向下一个单词
	moveLeft := func() {
		// left 指向下一个单词前，需要修改统计记录
		// 增加 left 指向的 word 可出现次数一次，
		remain[s[left:left+lenw]]++
		// 连续符合条件的单词数减少一个
		count--
		// left 后移一个 word 的长度
		left += lenw
	}

	// left 需要分别从{0,1,...,lenw-1}这些位置开始检查，才能不遗漏
	for i := 0; i < lenw; i++ {
		left, right = i, i
		reset()

		// s[left:] 的长度 >= words 中所有 word 组成的字符串的长度时，
		// s[left:] 中才有可能存在要找的字符串
		for lens-left >= lenws*lenw {
			word := s[right : right+lenw]
			remainTimes, ok := remain[word]

			switch {
			case !ok:
				// word 不在 words 中
				// 从 right+lenw 处，作为新窗口，重新开始统计
				left, right = right+lenw, right+lenw
				reset()
			case remainTimes == 0:
				// word 的出现次数上次就用完了
				// 说明 word 在 s[left:right] 中出现次数过多
				moveLeft()
				// 这个case会连续出现
				// 直到 s[left:right] 中的统计结果是 remain[word] == 1
				// 这个过程中，right 一直不变
			default:
				// ok && remainTimes > 0，word 符合出现的条件
				// moveRight
				remain[word]--
				count++
				right += lenw
				// 检查 words 能否排列组合成 s[left:right]
				if count == lenws {
					res = append(res, left)
					// moveLeft 可以避免重复统计 s[left+lenw:right] 中的信息
					moveLeft()
				}
			}
		}
	}

	return res
}
