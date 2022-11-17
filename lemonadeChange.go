package main

import (
	"fmt"
	"sort"
)

/**
Напишите функцию, которая посчитает, сможет ли продавец в киоске обслужить всех покупателей.
В киоске каждый лимонад стоит пять долларов. Клиенты стоят в очереди, чтобы купить у вас, и заказывают по одному лимонаду.
Каждый покупатель может купить только один лимонад и заплатить купюрами номиналом 5, 10 или 20 долларов.
Вы должны дать каждому покупателю сдачу с его купюры.

Обратите внимание, что сначала у вас нет сдачи.
**/

const (
	LEMONADE_COST = 5 // стоимость лимонада
)

func main() {
	bills := []int{5, 10, 5, 10, 15}      // суммы покупателей
	cashbox := make([]int, 0, len(bills)) // касса продавца

	for k, v := range bills {
		if v == LEMONADE_COST { // сдача не требуется
			cashbox = append(cashbox, v)
			continue
		}

		changeNeed := v - LEMONADE_COST
		fmt.Printf("Требуется сдача %d покупателю %d | ", changeNeed, k+1)

		ok := false
		if cashbox, ok = getChange(cashbox, changeNeed); ok {
			fmt.Println("Сдача выдана")
			cashbox = append(cashbox, v)
		} else {
			fmt.Println("Нет денег для сдачи")
		}
		printCashboxSum(cashbox)
	}
}

func getChange(cashbox []int, need int) ([]int, bool) {
	var sum int
	var delCashboxIndex = make([]int, 0, len(cashbox))

	sort.Slice(cashbox, func(i, j int) bool {
		return (cashbox)[i] > (cashbox)[j]
	})

	for k, v := range cashbox {
		if v <= need && (sum+v) <= need {
			sum += v
			delCashboxIndex = append(delCashboxIndex, k)
		}

		if sum == need {
			cashbox = removeElementsFromSlice(cashbox, delCashboxIndex)
			return cashbox, true
		}
	}

	return cashbox, false
}

func printCashboxSum(cashbox []int) {
	var sum int

	for i := 0; i < len(cashbox); i, sum = i+1, sum+cashbox[i] {
	}
	fmt.Println("В кассе осталось:", sum)
}

// удалить в слайсе элементы с индексами в []indices
func removeElementsFromSlice(sl []int, indices []int) []int {
	newSlice := make([]int, 0, cap(sl))

outLoop:
	for i, v := range sl {
		for ii, vv := range indices {
			if i == vv {
				indices = append(indices[0:ii], indices[ii+1:]...)
				continue outLoop
			}
		}
		newSlice = append(newSlice, v)
	}
	return newSlice
}

