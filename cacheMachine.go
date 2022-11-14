package main

import (
  "fmt"
  "sort"
)

/*****
Постановка задачи:

Программное обеспечение банкоматов постоянно решает задачу, как имеющимися купюрами сформировать сумму, введённую пользователем. Попробуйте решить похожую задачу и определить, сможет ли пользователь заплатить за товар без сдачи или нет. Для этого он будет вводить стоимость товара и номиналы трёх монет.
*****/

const N = 5 // количество номиналов монет

// структура хранит возможные варианты выдачи денег монетами различных номиналов
type coinsGiver struct {
  giveWays []giveWay
}

// структура хранит возможный вариант выдачи суммы
type giveWay struct {
  m map[int]int
}

func main() {
  var (
    success bool = false  // выдача денег возможна
    sum     int
    cg      coinsGiver = coinsGiver{}   
    coins   []int      = make([]int, N) // номиналы монет
  )

  fmt.Println("Пожалуйста, введите желаемую сумму для снятия: ")
  fmt.Scan(&sum)

  fmt.Printf("Пожалуйста, %d раз(а) введите номиналы монет: ", N)
  for i := range coins {
    fmt.Scan(&coins[i])
  }

  // сортировка по возрастанию
  sort.Slice(coins, func(i, j int) bool {
    return coins[i] < coins[j]
  })

  // определение максимального количества переборов для каждой монеты
  // например, для суммы 100 максимальное число монет с номиналом 5 будет = 20
  // для монет с номиналом 3 будет = 33
  enumsMap := getEnumsMap(coins, sum)

  // перебор всевозможных комбинаций монет
  for _, v := range combinations(enumsMap) {
    // вычисление суммы очередного варианта
    calc := 0
    for _, k := range v {
      calc += k
    }
    // если найдена подходящая комбинация монет
    if calc == sum {
      // подготовка структуры для записи
      newGiveWay := formatGiverStruct(v, enumsMap)

      cg.giveWays = append(cg.giveWays, newGiveWay)

      success = true
    }
  }

  //Вывод результата
  if success {
    cg.giveWaysPrint()
  } else {
    fmt.Println("К сожалению, желаемую сумму не получить имеющимися номиналами монет")
  }
}

func getEnumsMap(coins []int, sum int) map[int][]int {
  res := make(map[int][]int)

  for _, v := range coins {
    maxEn := sum / v
    enums := make([]int, 0, maxEn)

    for i := 0; i <= maxEn; i++ {
      enums = append(enums, i*v)
    }

    res[v] = enums
    enums = nil
  }

  return res
}

// печать возможных вариантов выдачи денег
func (cg coinsGiver) giveWaysPrint() {
  for i, j := range cg.giveWays {
    fmt.Printf("\n--------\n\nВариант № %d\n", (i + 1))
    for k, v := range j.m {
      if v != 0 {
        fmt.Printf("Выдать %d шт. номиналом %d\n", v, k)
      }
    }
  }
}

// замена значений мапы m на значение равное числу монет для выдачи
func formatGiverStruct(m map[int]int, enums map[int][]int) giveWay {
  for k, v := range m {
    for ek, ev := range enums[k] {
      if v == ev {
        m[k] = ek
      }
    }
  }
  return giveWay{m}
}

// вернуть всевозможные комбинации значений карты m
func combinations(m map[int][]int) []map[int]int {
  result := make([]map[int]int, 1)

  for property, property_values := range m {
    tmp := make([]map[int]int, 1)
    for _, result_item := range result {
      for _, property_value := range property_values {
        newMap := map[int]int{
          property: property_value,
        }
        merged := mergeMaps(newMap, result_item)

        tmp = append(tmp, merged)
      }
    }
    result = tmp
  }

  // удалить лишние слайсы
  for i, mp := range result {
    if len(mp) != len(m) {
      removeElementByIndex(result, i)
    }
  }

  return result
}

func mergeMaps(maps ...map[int]int) (result map[int]int) {
  result = make(map[int]int)
  for _, m := range maps {
    for k, v := range m {
      result[k] = v
    }
  }
  return result
}

func removeElementByIndex(slice []map[int]int, index int) []map[int]int {
  sliceLen := len(slice)
  sliceLastIndex := sliceLen - 1

  if index != sliceLastIndex {
    slice[index] = slice[sliceLastIndex]
  }

  return slice[:sliceLastIndex]
}
