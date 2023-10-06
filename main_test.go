package main

import (
	"container/list"
	"regexp"
	"testing"
)

/*func TestStringSplitBlocksTest(t *testing.T) {

	testTable := []struct {
		a        string
		expected []string
	}{{
		a:        "Akaka, 325, 25.08.2002",
		expected: []string{"Akaka", "325", "25.08.2002"},
	},
		{
			a:        "asdafga, hdhfhdfh, hdfhgdfg, dhdfhfdh",
			expected: []string{"asdafga", "hdhfhdfh", "hdfhgdfg", "dhdfhfdh"},
		},
		{
			a:        "",
			expected: []string{},
		},
	}

	result := StringSplitBlocks(testTable.expected)

	/*for i, testCase := range testTable {
		result := StringSplitBlocks(testCase.a)
		if result[i] != testCase.expected[i] {
			t.Logf("Incorrect result")
		}
	}
}*/

func TestAddBird(t *testing.T) {
	// Создаем список для хранения птиц
	birdsList := list.New()

	// Тест 1: добавление птицы с корректными параметрами
	line1 := "Bird (Sparrow, 10.5)"
	matches1 := regexp.MustCompile(`\(([^)]+)\)`).FindStringSubmatch(line1)
	if !AddBird(birdsList, matches1, line1) {
		t.Errorf("Test 1: Expected true, got false")
	}

	// Проверяем, что птица успешно добавлена в список
	if birdsList.Len() != 1 {
		t.Errorf("Test 1: Expected list length 1, got %d", birdsList.Len())
	}

	// Проверяем, что добавленная птица имеет правильные параметры
	bird1 := birdsList.Front().Value.(Bird)
	if bird1.name != "Sparrow" {
		t.Errorf("Test 1: Expected bird name Sparrow, got %s", bird1.name)
	}
	if bird1.speed != 10.5 {
		t.Errorf("Test 1: Expected bird speed 10.5, got %f", bird1.speed)
	}

	// Тест 2: добавление птицы с некорректными параметрами
	line2 := "bird Pigeon"
	matches2 := regexp.MustCompile(`\(([^)]+)\)`).FindStringSubmatch(line2)
	err := AddBird(birdsList, matches2, line2)
	if err {
		t.Errorf("Test 2: Expected false, got true")
	}

	// Проверяем, что список птиц не изменился
	if birdsList.Len() != 1 {
		t.Errorf("Test 2: Expected list length 1, got %d", birdsList.Len())
	}
}

func TestAdd(t *testing.T) {
	container := list.New()

	// Test adding Fish
	line1 := "Fish (Andrey, AAA)"
	add(line1, container)
	if container.Len() != 1 {
		t.Errorf("Expected container length to be 1, got %d", container.Len())
	}

	// Test adding Bird
	line2 := "Bird (kaka Akak, 325)"
	add(line2, container)
	if container.Len() != 2 {
		t.Errorf("Expected container length to be 2, got %d", container.Len())
	}

	// Test adding Insects
	line3 := "Insects (Akaka, 325, 25.08.2002)"
	add(line3, container)
	if container.Len() != 3 {
		t.Errorf("Expected container length to be 3, got %d", container.Len())
	}

	// Test adding invalid line
	line4 := "Invalid line"
	add(line4, container)
	if container.Len() != 3 {
		t.Errorf("Expected container length to still be 3, got %d", container.Len())
	}
}

/*func addBirdTest()

func addFishTest()

func addInsectsTest()

func addtest()

func remtest()

func printest()*/
