package entry

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type newData struct {
	n      int
	arrayX []float32
	arrayY []float32
}

func GetData() *newData {
	fileOrConsole := "file" //TODO
	for ;fileOrConsole != "file" && fileOrConsole != "console"; {
		fmt.Print("Ввод данных из file, console: ")
		fmt.Scan(&fileOrConsole)
	}
	data := newData{}
	if fileOrConsole == "file" {
		data.readFromFile()
	} else {
		data.readFromConsole()
	}
	return &data
}

func (data *newData) readFromFile() {
	file, err := os.Open("resources/input.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл.")
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	line = line[:len(line)-2]
	s := strings.Split(line, " ")
	(*data).n = len(s)
	for i := 0; i < len(s); i++ {
		x, err := strconv.ParseFloat(s[i], 32)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Некорректно введено значение x%d", i)
			os.Exit(1)
		}
		(*data).arrayX = append((*data).arrayX, float32(x))
	}

	line, err = reader.ReadString('\n')
	line = line[:len(line)-2]
	s = strings.Split(line, " ")
	for i := 0; i < len(s); i++ {
		y, err := strconv.ParseFloat(s[i], 32)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Некорректно введено значение y%d", i)
			os.Exit(1)
		}
		(*data).arrayY = append((*data).arrayY, float32(y))
	}
}

func (data *newData) readFromConsole() {
	tableOrFunction := "table" //TODO
	for ; tableOrFunction != "func" && tableOrFunction != "table"; {
		fmt.Print("Задать данные набором или фукнуцией:")
		fmt.Scan(&tableOrFunction)
	}
	function := ""
	if tableOrFunction == "func" {
		for ; function != "1" && function != "2"; {
			fmt.Print("Выберите функцию:\n\t\t sin(x) - введите 1 \n\t\t х \t\t- введите 2")
			fmt.Scan(&function)
		}
		//TODO
	} else {
		y, err := strconv.ParseInt(function, 10, 32)
		for ; err != nil; {
			fmt.Print("Введите количество точек:")
			fmt.Scan(&function)
			y, err = strconv.ParseInt(function, 10, 32)
		}
		(*data).n = int(y)

		var inp string
		for i := 0; i < data.n; i++ {

			for ; ; {
				fmt.Fprintf(os.Stdout, "Введите значение х%d:", i)
				fmt.Scan(&inp)
				if x, err := strconv.ParseFloat(inp, 32)
					err == nil {
					(*data).arrayX = append((*data).arrayX, float32(x))
					break
				}
			}

			for ; ; {
				fmt.Fprintf(os.Stdout, "Введите значение y%d:", i)
				fmt.Scan(&inp)
				if y, err := strconv.ParseFloat(inp, 32)
					err == nil {
					(*data).arrayY = append((*data).arrayY, float32(y))
					break
				}
			}
		}
	}
}
