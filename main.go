package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math" // mod
	"os"   // Работа с файлами
	"regexp"
	"strconv"
	"strings" // Работа со строками
	"time"
)

func checkdate(date string) bool {
	t, err := time.Parse("02.01.2006", date)
	if err != nil {
		return false
	} else {
		return t.Format("02.01.2006") == date
	}
}

type print interface {
	PrintContainer() // интерфейс
}

type Fish struct {
	name  string
	areal string
}

func (a Fish) PrintContainer() {
	fmt.Println("Рыба->", "Имя:", a.name, "Место обитания:", a.areal) // Вывод для рыбы
}

type Bird struct {
	name  string
	speed float64
}

func (a Bird) PrintContainer() {
	fmt.Println("Птица->", "Имя:", a.name, "Скорость:", a.speed) // Вывод для птицы
}

type Insects struct {
	name string
	size float64
	date string
}

func (a Insects) PrintContainer() {
	fmt.Println("Насекомое->", "Имя:", a.name, "Размер:", a.size, "Дата обнаружения:", a.date) // Вывод для насекомого
}

func splitValues(values string) []string {
	re := regexp.MustCompile(`\s*,\s*`)
	blocks := re.Split(values, -1)
	return blocks
}

func ADD(line string, container *list.List) {
	re := regexp.MustCompile(`\(([^)]+)\)`) // Регулярное выражение
	matches := re.FindStringSubmatch(line)
	if strings.Contains(line, "Fish") {
		fmt.Println("Fish")
		if len(matches) > 1 {
			values := matches[1]
			blocks := splitValues(values) // Разбиваем значения на блоки
			if len(blocks) == 2 {
				var a Fish // структура типа Fish
				a.name = blocks[0]
				a.areal = blocks[1]
				container.PushBack(a)
			} else {
				fmt.Println("Not enough parameters for", line) // Если внутри скобок не заданы параметры
				return
			}
		} else {
			fmt.Println("No parameters found", line)
			return
		}
	} else if strings.Contains(line, "Bird") {
		fmt.Println("Bird")
		if len(matches) > 1 {
			values := matches[1]
			blocks := splitValues(values) // Разбиваем значения на блоки
			if len(blocks) == 2 {
				var a Bird // структура типа Bird
				a.name = blocks[0]
				if StrToFloat, err := strconv.ParseFloat(blocks[1], 64); err == nil {
					a.speed = StrToFloat
				} else {
					panic(err) // если string не конвертировался в float64
				}
				container.PushBack(a)
			} else {
				fmt.Println("Not enough parameters for", line) // Если внутри скобок не заданы параметры
				return
			}
		} else {
			fmt.Println("No parameters found", line)
		}
	} else if strings.Contains(line, "Insects") {
		fmt.Println("Insects")
		if len(matches) > 1 {
			values := matches[1]
			blocks := splitValues(values) // Разбиваем значения на блоки
			if len(blocks) == 3 {
				var a Insects // структура типа Insects
				a.name = blocks[0]
				if StrToFloat, err := strconv.ParseFloat(blocks[1], 64); err == nil {
					a.size = StrToFloat
				} else {
					panic(err) // если string не конвертировался в float64
				}
				if checkdate(blocks[2]) {
					a.date = blocks[2]
				} else {
					fmt.Println("Invalid date")
					return
				}
				container.PushBack(a)
			} else {
				fmt.Println("Not enough parameters for", line) // Если внутри скобок не заданы параметры
				return
			}
		} else {
			fmt.Println("No parameters found", line)
			return
		}
	}
}

func REM(container *list.List, line string) {
	words := strings.SplitN(line, " ", 2) // разбивает строку на подстроки, используя в качестве разделителя пробел (" "). Второй аргумент, 2, указывает на то, что она должна разделить строку не более чем на две подстроки.
	if len(words) > 1 {
		result := words[1]
		fmt.Println(result)
		for e := container.Front(); e != nil; e = e.Next() {
			if fish, ok := e.Value.(Fish); ok {
				if fish.name == result {
					container.Remove(e)
					fmt.Println("Удалил")
				}
			} else if bird, ok := e.Value.(Bird); ok {
				if bird.name == result {
					container.Remove(e)
					fmt.Println("Удалил")
				}
			} else if insects, ok := e.Value.(Insects); ok {
				if insects.name == result {
					container.Remove(e)
					fmt.Println("Удалил")
				}
			}
		}
	} else {
		fmt.Println("Not found word")
		return
	}
}

func PRINT(container list.List) {
	for e := container.Front(); e != nil; e = e.Next() {
		if fish, ok := e.Value.(Fish); ok {
			fish.PrintContainer()
		} else if bird, ok := e.Value.(Bird); ok {
			bird.PrintContainer()
		} else if insects, ok := e.Value.(Insects); ok {
			insects.PrintContainer()
		}
	}
}

const nvar int16 = 9 // Номер варианта

func ntask(nvar int16) float64 { // Номер условия задачи

	fmt.Println("Номер условия задачи: ")
	return math.Mod(float64(nvar)-1, 11) + 1

}

func ncont(nvar int16) float64 { // Номер контейнера

	fmt.Println("Номер контейнера: ")
	return math.Mod(float64((nvar-1)%12), 3)

}

func main() {

	fmt.Println(ntask(nvar)) // Выводим условие задачи
	fmt.Println(ncont(nvar)) // Выводим номер контейнера

	container := list.New() // Двунаправленный линейный список

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "ADD") {
			ADD(line, container)
		} else if strings.HasPrefix(line, "REM") {
			REM(container, line)
		} else if strings.HasPrefix(line, "PRINT") {
			PRINT(*container)
		}
	}
}

/* fmt.Println("Fish")
re := regexp.MustCompile(`\(([^)]+)\)`) // Регулярное выражение
matches := re.FindStringSubmatch(line)
if len(matches) > 1 {
	values := matches[1]
	blocks := splitValues(values) // Разбиваем значения на блоки
	if len(blocks) == 2 {
		fmt.Println(len(blocks), "Размер")
		for _, block := range blocks {
			fmt.Println(block)
		}
	} else {
		fmt.Println("Not enough parameters for", line) // Если внутри скобок не заданы параметры
	}
} else {
	fmt.Println("No parameters found", line)
}
*/
