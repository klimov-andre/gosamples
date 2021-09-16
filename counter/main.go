package main

import (
	"fmt"
	"sort"
	"strings"
)

func Counter(s string) string {
	// сначала подсчитываем число вхождений каждого символа
	counter := make(map[rune]int)
	for _, c := range s {
		counter[c]++
	}

	// выделяем все присуствующие символы в строке (ключи в мапе)
	allSymbols := []rune{}
	for k := range counter {
		allSymbols = append(allSymbols, k)
	}

	// эффективнее сортировать не всю строку целиком, а только всчтречающиеся в строке символы отдельно,
	// т.к. число ключей в мапе может быть не больше количества символов в языке,
	// при этом длина строки вообще не ограничена
	sort.Slice(allSymbols, func(i, j int) bool {
		return allSymbols[i] < allSymbols[j]
	})

	var resBuilder strings.Builder
	for _, c := range allSymbols {
		fmt.Fprintf(&resBuilder, "%c%d", c, counter[c])
	}

	return resBuilder.String()
}
