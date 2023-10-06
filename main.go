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

type Creature interface {
	PrintElement() // интерфейс
}

type Fish struct {
	name  string
	areal string
}

func (a Fish) PrintElement() {
	fmt.Println("Рыба->", "Имя:", a.name, "Место обитания:", a.areal) // Вывод для рыбы
}

type Bird struct {
	name  string
	speed float64
}

func (a Bird) PrintElement() {
	fmt.Println("Птица->", "Имя:", a.name, "Скорость:", a.speed) // Вывод для птицы
}

type Insects struct {
	name string
	size float64
	date string
}

func (a Insects) PrintElement() {
	fmt.Println("Насекомое->", "Имя:", a.name, "Размер:", a.size, "Дата обнаружения:", a.date) // Вывод для насекомого
}

func Sound(a Creature) {
	a.PrintElement() // Вывод интерфейса
}

func ParametersToBlocks(values string) []string { // разбиваем строку (*1*.*2*.*3*) на блоки block[0] = *1*, block[1] = *2*, block[2] = *3*
	re := regexp.MustCompile(`\s*,\s*`) // Используем регулярное выражение
	blocks := re.Split(values, -1)
	return blocks
}

// Добавление рыбы
func AddFish(container *list.List, matches []string, line string) {
	if len(matches) > 1 {
		values := matches[1]                 // параметры переданные в функция для структуры Fish
		blocks := ParametersToBlocks(values) // Разбиваем значения на блоки
		if len(blocks) == 2 {                // 1 блок - имя, 2-ой блок - ареал
			var a Fish          // структура типа Fish
			a.name = blocks[0]  // Записываем 1-ый блок в имя
			a.areal = blocks[1] // Записываем 2-ой блок в ариал
			container.PushBack(a)
		} else {
			fmt.Println("Not enough parameters for", line) // Если внутри скобок не заданы параметры
		}
	}
}

// Добавление птицы
func AddBird(container *list.List, matches []string, line string) (err bool) {
	if len(matches) > 1 {
		values := matches[1]                 // параметры переданные в функция для структуры Bird
		blocks := ParametersToBlocks(values) // Разбиваем значения на блоки
		if len(blocks) == 2 {                // 1 блок - имя, 2-ой блок - скорость
			var a Bird                                                            // структура типа Bird
			a.name = blocks[0]                                                    // Записывается имя из блока
			if StrToFloat, err := strconv.ParseFloat(blocks[1], 64); err == nil { // Конвертация string в float64
				a.speed = StrToFloat // Записываем 2-ой блок в скорость
			} else {
				panic(err) // если string не конвертировался в float64
			}
			container.PushBack(a) // Добавление в контейнер
			return true           // Успешно добавился
		} else {
			fmt.Println("Not enough parameters for", line) // Если внутри скобок не заданы параметры
			return false                                   // Ошибка
		}
	} else {
		// fmt.Println("No parameters found", line) // Не найден параметр
		return false
	}
}

// Добавление насекомого
func AddInsects(container *list.List, matches []string, line string) (err bool) {
	if len(matches) > 1 {
		values := matches[1]                 // параметры переданные в функция для структуры Insects
		blocks := ParametersToBlocks(values) // Разбиваем значения на блоки
		if len(blocks) == 3 {                // 1 блок - имя, 2-ой блок - размер, 3 блок - дата обнаружения
			var a Insects                                                         // структура типа Insects
			a.name = blocks[0]                                                    // // Записываем 1-ый блок в имя
			if StrToFloat, err := strconv.ParseFloat(blocks[1], 64); err == nil { // Конвертация string в float64
				a.size = StrToFloat // Записываем 2-ой блок в размеры
			} else {
				panic(err) // если string не конвертировался в float64
			}
			if checkdate(blocks[2]) { // Проверка даты на валидность
				a.date = blocks[2] // Проверка выдала true
			} else {
				fmt.Println("Invalid date") // Проверка выдала false
				return false
			}
			container.PushBack(a) // Добавление в контейнер
			return true
		} else {
			fmt.Println("Not enough parameters for", line) // Если внутри скобок не заданы параметры
			return false
		}
	} else {
		fmt.Println("No parameters found", line)
		return false
	}
}

// Функция добавления
func add(line string, container *list.List) {
	re := regexp.MustCompile(`\(([^)]+)\)`) // Регулярное выражение
	matches := re.FindStringSubmatch(line)
	if strings.Contains(line, "Fish") { // В строке есть слово "Fish"
		AddFish(container, matches, line)
	} else if strings.Contains(line, "Bird") { // В строке есть слово "Bird"
		AddBird(container, matches, line)
	} else if strings.Contains(line, "Insects") { // В строке есть слово "Insects"
		AddInsects(container, matches, line)
	}
}

// Функция удаления
func rem(container *list.List, line string) {
	words := strings.SplitN(line, " ", 2) // разбивает строку на подстроки, используя в качестве разделителя пробел (" "). Второй аргумент, 2, указывает на то, что она должна разделить строку не более чем на две подстроки.
	if len(words) > 1 {
		result := words[1]
		for e := container.Front(); e != nil; e = e.Next() { // Удаление по типу структуры
			if fish, ok := e.Value.(Fish); ok { // Если рыба
				if fish.name == result {
					container.Remove(e)
					e = container.Front()
				}
			} else if bird, ok := e.Value.(Bird); ok { // Если птица
				if bird.name == result {
					container.Remove(e)
					e = container.Front()
				}
			} else if insects, ok := e.Value.(Insects); ok { // Если насекомое
				if insects.name == result {
					container.Remove(e)
					e = container.Front()
				}
			}
		}
	} else {
		fmt.Println("Not found word") // Если не найдено слово
	}
}

// Функция вывода
func print(container list.List) {
	for e := container.Front(); e != nil; e = e.Next() {
		if creature, ok := e.Value.(Creature); ok { // сравниваем является ли значением допустимым значением Fish, Bird, Insects
			Sound(creature) // В зависимости от типа структуры Fish, Bird, Insects - будет "звук", которые выводит его параметры
		}
	}
	fmt.Println("________________________")
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
		if strings.HasPrefix(line, "ADD") { // Проверяет, первое слово в строке "ADD"
			add(line, container)
		} else if strings.HasPrefix(line, "REM") { // Проверяет, первое слово в строке "REM"
			rem(container, line)
		} else if strings.HasPrefix(line, "PRINT") { // Проверяет, первое слово в строке "PRINT"
			print(*container)
		}
	}
}
