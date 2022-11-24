package main

import (
	"bufio"
	"encoding/json"
	"errors"
	f "fmt"
	"os"
	"strconv"
	"strings"
)

/**
Что нужно сделать:
Напишите программу, которая считывает ввод с stdin,
создаёт структуру student и записывает указатель на структуру в хранилище map[studentName] *Student.

Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent,
далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d)
вывести на экран имена всех студентов из хранилища. Также необходимо реализовать методы put, get.
----
при получении одной строки (например, «имяСтудента 24 1») программа создаёт студента и сохраняет его, далее ожидает следующую строку или сигнал EOF (Сtrl + Z);
при получении сигнала EOF программа должна вывести имена всех студентов из map.
**/

const (
	PROMPT          = "Пожалуйста, введите имя, возраст и ранг студента:"
	EXIT_WORD       = "EOF"
	STUD_FIELDS_LEN = 3
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

// students storage
type StudentsList map[uint64]*Student

type IStudentsList interface {
	put(uint64, *Student)
	get(uint64) *Student
}

func (stl StudentsList) put(idx uint64, s *Student) {
	stl[idx] = s
}
func (stl StudentsList) get(idx uint64) *Student {
	if s, ok := stl[idx]; ok {
		return s
	}
	return nil
}
func (stl StudentsList) Print() {
	f.Println("Список студентов:")
	f.Printf("\nid\tname\tage\tgrade\n%s\n", strings.Repeat("-", 30))
	for idx, s := range stl {
		f.Printf("%d\t%s\t%d\t%d\n", idx, s.Name, s.Age, s.Grade)
	}
}

func newStudent(inp map[string]interface{}) *Student {
	s := &Student{}

	// fill student struct by values from map
	dbByte, _ := json.Marshal(inp)
	json.Unmarshal(dbByte, s)

	return s
}

func main() {
	var idx uint64 = 1 //student id
	studList := make(StudentsList)
	sc := bufio.NewScanner(os.Stdin)

	f.Println(PROMPT)
	for sc.Scan() {
		f.Println(PROMPT)
		inp := sc.Text()
		if inp == EXIT_WORD {
			break
		}

		words := strings.Fields(inp)

		// check input
		inpMap, err := parseInp(words)
		if err != nil {
			f.Printf("%v\n", err)
			continue
		}

		// на выходе получить переменные
		newStudent := newStudent(inpMap)
		studList.put(idx, newStudent)
		idx++
	}

	//output all students
	studList.Print()
}

func parseInp(inp []string) (map[string]interface{}, error) {
	var err error

	if len(inp) != STUD_FIELDS_LEN {
		err = errors.New("ошибка. Введите в формате: имя (string), возраст (int), ранг(int)")
		return nil, err
	}

	res := make(map[string]interface{}, STUD_FIELDS_LEN)

	res["age"], err = strconv.Atoi(inp[1])
	if err != nil {
		return nil, f.Errorf("не удалось считать возраст: %w", err)
	}

	res["grade"], err = strconv.Atoi(inp[2])
	if err != nil {
		return nil, f.Errorf("не удалось считать ранг: %w", err)
	}

	res["name"] = inp[0]

	return res, nil
}

