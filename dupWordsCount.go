package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

/**
* Назначение: Программа подсчитывает количество повторяющихся слов в тексте
* Входные данные: текстовый файл
**/

/*
* Примеры использования:
* go run dupWordsCount test.txt (подсчитает число повторяющихся слов в файле test.txt)
* 
 */

const (
	defaultWordsBuffer = 1024 
)

func main() {
	flag.Parse()

	for _, fname := range flag.Args() {
		words, err := countFileWords(fname)
		if err != nil {
			fmt.Printf("Ошибка чтения файла %s\nОшибка:%v", fname, err)
			continue
		}
		dupWords := countDupWords(words)
		printDupWords(dupWords, fname)
	}
}

func countFileWords(filename string) (map[string]int, error) {
	words := make(map[string]int)
	f, err := os.Open(filename)

	if err != nil {
		return words, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := filterWord(scanner.Text())

		if word != "" {
			words[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return words, err
	}

	return words, nil
}

func filterWord(word string) string {
	reg, err := regexp.Compile(`[^0-9a-zA-Zа-яёА-ЯЁ]`)

	if err != nil {
		panic(err)
	}

	return reg.ReplaceAllString(word, "")
}

func countDupWords(words map[string]int) map[string]int {
	for i := range words {
		if words[i] == 1 {
			delete(words, i)
		}
	}

	return words
}

func printDupWords(words map[string]int, filename string) {
	const (
		mainHead = iota
		noneDupWords
		resHead
	)
	txt := [...]string{
		fmt.Sprintf("\n%sРезультаты анализа файла <%s>%[1]s\n\n", strings.Repeat("*", 10), filename),
		"Повторяющихся слов не обнаружено",
		"Слово\t\tЧисло повторений в тексте\n\n",
	}

	fmt.Printf(txt[mainHead])

	if len(words) == 0 {
		fmt.Println(txt[noneDupWords])
		return
	}
	/** Сортировка по возрастанию **/
	keys := make([]string, 0, len(words))
	for k := range words {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return words[keys[i]] > words[keys[j]]
	})

	/** Вывод на печать **/
	fmt.Printf(txt[resHead])
	for _, k := range keys {
		fmt.Printf("%-20s%d\n", k, words[k])
	}
}
