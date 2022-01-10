package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//https://www.codewars.com/users/azirsold
	x := "cake"
	fmt.Println(&x)
}

func duplicate_count(s1 string) int { //6kyu
	dup := strings.Split(strings.ToLower(s1), "")
	sort.Strings(dup)
	result := []string{}

	for i, value := range dup {

		if i+1 < len(dup) {

			if dup[i+1] == value {

				if len(result) == 0 {

					result = append(result, value)

				} else if value != result[len(result)-1] {
					result = append(result, value)

				}
			}
		}
	}
	return len(result)
}

// que codigo massa, pena que não é meu. basicamente ele coloca +1 dentro
// de um map ultilizando o valor da rune de cada char da string
// então digamos que 'a' que tem a rune 97, ele coloca nesse array na posição
// runeMap[97]  += 1
// no final ele loop em cima do runeMap e todo valor acima de 1 significa que
// o char se repetiu, visto que dois 'a' adicionariam runeMap[97] += 1
// finissimo
func duplicate_countRefactor(duplicate string) (result int) { //6kyu refactor, not my code.
	runeMap := map[rune]int{}

	for _, dValue := range strings.ToLower(duplicate) {
		runeMap[dValue]++
		fmt.Println(dValue)
	}
	for _, rValue := range runeMap {
		if rValue > 1 {
			result++
		}
	}
	fmt.Println(runeMap)
	return
}

func SpinWords(str string) string { //6kyu
	var result string
	spli := strings.Split(str, " ")
	for i, _ := range spli {
		if len(spli[i]) >= 5 {
			for _, value := range spli[i] {
				result = string(value) + result
			}
			spli[i] = result
			result = ""
		}
	}
	result = strings.Join(spli, " ")
	return result
}

//any positive number
func Multiple3And5Refactor(number int) int { //refactor looking at someone else code //6Kyu /
	total := 0
	if number <= 0 {
		return 0
	}
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}

//any positive number
func Multiple3And5(number int) int { //6Kyu
	five := 5
	three := 3
	total := 0
	if number <= 0 {
		return 0
	}
	for i := 1; five < number; i++ {
		total += 5 * i
		five += 5
	}
	for i := 1; three < number; i++ {
		if (3 * i % 5) != 0 {
			total += 3 * i
			three += 3
		} else {
			three += 3
		}
	}
	return total
}

//any positive number
func RowSumOddNumbersRefactor(n int) int { //refactor looking at someone else code //7Kyu
	return n * n * n //this is more a math problem than an code problem -.-'
}

//any positive number
func RowSumOddNumbers(n int) int { //7Kyu
	if n == 1 {
		return 1
	}
	rowNumb := (n - 1) * 2
	start := 1
	result := 0

	for rowNumb > 0 {
		start += rowNumb
		rowNumb -= 2
	}
	for i := 0; i <= n-1; i++ {
		result += start
		start += 2
	}
	return result
}

//a year, a percentage, an int value to increase, a final value
func NbYearRefactor(p0 int, percent float64, aug int, p int) int { //7kyu no copy on this refactor
	years := 0

	for p0 < p {
		p0 = int(((float64(p0) * (100.0 + percent)) / 100)) + aug
		years++
	}

	return years
}

//a year, a percentage, an int value to increase, a final value
func NbYear(p0 int, percent float64, aug int, p int) int { //7kyu
	years := 0
	p1 := []float64{float64(p0), float64(aug), float64(p)}

	for p1[0] < p1[2] {
		p1[0] = ((p1[0] * (100.0 + percent)) / 100) + p1[1]
		years++
	}

	return years
}

// two strings to sort and remove duplicate chars
func TwoToOne(s1 string, s2 string) string { //7kyu
	two := s1 + s2
	twoTo := strings.Split(two, "")
	sort.Strings(twoTo)
	result := []string{twoTo[0]}

	for _, value := range twoTo {
		if result[len(result)-1] != value {
			result = append(result, value)
		}
	}
	return strings.Join(result, "")
}

//a string to count aeiou, dunno why this count tho, 0 default
func GetCount(str string) (count int) { //7kyu
	for _, value := range str {
		if value == 'a' || value == 'e' || value == 'i' || value == 'o' || value == 'u' {
			count += 1
		}
	}
	return count
}

//any number, it will invert to negative/positive  (yeah, you can just retur -x and will do the same work)
func MakeNegative(x int) int { //8kyu
	if x > 0 {
		return ((x ^ -1) + 1)
	} else {
		return x
	}
}

// array of positive numbers
func PositiveSumRefactor(numbers []int) int { //8kyu refactor, no copy
	total := 0
	for _, value := range numbers {
		if value > 0 {
			total += value
		}
	}
	return total
}

// array of positive numbers
func PositiveSum(numbers []int) int { //8kyu
	positives := []int{}
	total := 0
	for _, value := range numbers {
		if value > 0 {
			positives = append(positives, value)
		}
	}
	for _, value := range positives {
		total += value
	}
	return total
}

//any positive number
func EvenOrOdd(number int) string { //8kyu
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
