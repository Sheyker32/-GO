package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = []string{"+", "-", "/", "*"}

var romans = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
}

func getRezult(a, b int, operator string) (num int) {
	if a > 10 || b > 10 {
		panic(fmt.Errorf("Вводимое число не может превышать 10"))
	}
	switch operator {
	case "+":
		num = a + b
	case "-":
		num = a - b
	case "/":
		num = a / b
	case "*":
		num = a * b
	}
	return
}

func isRoman(num string) bool {
	if _, err := romans[strings.Split(num, "")[0]]; err {
		return true
	}

	return false
}

func romanToInt(num string) int {
	sum := 0
	n := len(num)

	for i := 0; i < n; i++ {
		if i != n-1 && romans[string(num[i])] < romans[string(num[i+1])] {
			sum += romans[string(num[i+1])] - romans[string(num[i])]
			i++
			continue
		}
		sum += romans[string(num[i])]
	}

	return sum
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func calculateRomansNums(num1, num2, op string) {
	a := romanToInt(num1)
	b := romanToInt(num2)

	if (op == "-" && a <= b) || (op == "/" && a < b) {
		panic(fmt.Errorf("Результатом работы калькулятора с римскими числами могут быть только положительные числа."))
	}

	rez := getRezult(a, b, op)
	fmt.Println(intToRoman(rez))
}

func main() {

	arrString := []string{}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	for _, value := range operators {
		arr := strings.Split(text, value)
		if len(arr) == 2 {
			arrString = append(arrString, arr[0], arr[1], value)
		}
	}

	if len(arrString) == 0 || len(arrString) > 3 {
		panic(fmt.Errorf("Операция возможна только между двумя целыми числами в диапозоне от 1 до 10"))
	}

	if isRoman(arrString[0]) && isRoman(arrString[1]) {
		calculateRomansNums(arrString[0], arrString[1], arrString[2])
	} else {
		num1, err := strconv.Atoi(arrString[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(arrString[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(getRezult(num1, num2, arrString[2]))
	}
}
