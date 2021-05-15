package entry

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type newData struct {
	n      int
	arrayX []float32
	arrayY []float32
	x      float32
}

func GetData() *newData {
	fileOrConsole := "" //TODO
	for ; fileOrConsole != "file" && fileOrConsole != "console"; {
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

	line, err = reader.ReadString('\n')
	line = line[:len(line)-2]
	x, err := strconv.ParseFloat(line, 32)
	(*data).x = float32(x)
}

func (data *newData) readFromConsole() {

	tableOrFunction := "" //TODO
	for ; tableOrFunction != "func" && tableOrFunction != "table"; {
		fmt.Print("Задать данные набором или фукнуцией:")
		fmt.Scan(&tableOrFunction)
	}

	function := ""

	y, err := strconv.ParseInt(function, 10, 32)
	for ; err != nil; {
		fmt.Print("Введите количество точек:")
		fmt.Scan(&function)
		y, err = strconv.ParseInt(function, 10, 32)
	}
	(*data).n = int(y)

	if tableOrFunction == "func" {

		for ; function != "1" && function != "2" && function != "3"; {
			fmt.Print("Выберите функцию:\n\t\t sin(x) \t- введите 1 \n\t\t х \t\t\t- введите 2 \n\t\t x^2-x-3 \t - введите 3\n")
			fmt.Scan(&function)
		}

		to := "z"
		x1, err := strconv.ParseFloat(to, 32)
		for ; err != nil; {
			fmt.Print("Введите левую границу:")
			fmt.Scan(&to)
			x1, err = strconv.ParseFloat(to, 32)
		}

		x2, err := strconv.ParseFloat("z", 32)
		for ; err != nil; {
			fmt.Print("Введите правую границу:")
			fmt.Scan(&to)
			x2, err = strconv.ParseFloat(to, 32)
		}

		h := (x2 - x1) /float64((*data).n)
		for i := x1; i <= x2; i += h {
			(*data).arrayX = append((*data).arrayX, float32(i))
			if function == "1" {
				(*data).arrayY = append((*data).arrayY, float32(math.Sin(i)))
			}else if function == "2"{
				(*data).arrayY = append((*data).arrayY, float32(i))
			}else if function == "3"{
				(*data).arrayY = append((*data).arrayY, float32(i*i - i - 3))
			}
		}
	} else {

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
	z, err := strconv.ParseFloat("z", 32)
	for ; err != nil; {
		fmt.Print("Введите значение X:")
		fmt.Scan(&function)
		z, err = strconv.ParseFloat(function, 32)
	}
	(*data).x = float32(z)
}
