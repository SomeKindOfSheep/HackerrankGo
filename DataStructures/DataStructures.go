package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	/*	arr := [][]int32{
		{1},
		{1, 2, 3},
		{4, 5, 6},
		{4, 5},
		{4},
		{9, 8, 9},
	}*/

	//diagonalDifference(arr)
	//fmt.Println(aVeryBigSum(ar))xx

	//ar := []int32{1, 1, 0, -1, -2}
	//plusMinus(ar)

	//staircase(6)

	//ar := []int32{4, 4, 1, 3}
	//miniMaxSum(ar)
	//birthdayCakeCandles(ar)

	time := "12:01:00AM"
	fmt.Println(timeConversion(time))
}

func timeConversion(s string) string {
	var sb strings.Builder
	// check if am or pm
	// convert
	if strings.Contains(s, "PM") {
		firstSplit := strings.Split(s, "PM")
		secondSplit := strings.Split(firstSplit[0], ":")
		convertedStringSplit, _ := strconv.Atoi(secondSplit[0])
		var newValue = convertedStringSplit + 12

		if newValue == 24 {
			newValue = 12
		}

		sb.WriteString(strconv.Itoa(newValue))
		sb.WriteString(":")
		sb.WriteString(secondSplit[1])
		sb.WriteString(":")
		sb.WriteString(secondSplit[2])
		return sb.String()
	}

	if strings.Contains(s, "AM") {
		firstSplit := strings.Split(s, "AM")
		secondSplit := strings.Split(firstSplit[0], ":")
		convertedStringSplit, _ := strconv.Atoi(secondSplit[0])
		var newValue = convertedStringSplit
		if newValue == 12 {
			newValue = 0
		}
		if newValue < 10 {
			sb.WriteString("0")
		}
		sb.WriteString(strconv.Itoa(newValue))
		sb.WriteString(":")
		sb.WriteString(secondSplit[1])
		sb.WriteString(":")
		sb.WriteString(secondSplit[2])
		return sb.String()
	}

	return ""
}

func birthdayCakeCandles(candles []int32) int32 {
	m := make(map[int32]int32)

	for i := 0; i < len(candles); i++ {

		// check if key exists
		// if yes, increase value
		if m[candles[i]] == candles[i] {
		}

		if val, ok := m[candles[i]]; ok {
			val++
			mapValue := m[candles[i]]
			m[candles[i]] = mapValue + 1
		} else {
			m[candles[i]] = 1
		}
	}

	var maxKey int32 = 0
	for key := range m {
		if key > maxKey {
			maxKey = key
		}
	}

	return m[maxKey]
}

func miniMaxSum(arr []int32) {

	// have to sort
	var i = 1
	for i < len(arr) {
		var j = i
		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		i++
	}

	var sumMin int64 = 0
	var sumMax int64 = 0

	for i := 0; i < len(arr); i++ {
		// sum max
		if i > 0 {
			sumMax = sumMax + int64(arr[i])
		}
		// sum min
		if i+1 < len(arr) {
			sumMin = sumMin + int64(arr[i])
		}
	}
	fmt.Println(sumMin, sumMax)
}

func staircase(n int32) {

	var stairs strings.Builder
	for i := 1; i < int(n)+1; i++ {
		var lineSb strings.Builder
		diff := int(n) - i

		for j := 0; j < int(n); j++ {

			if j >= diff {
				lineSb.WriteString("#")
			}
			if j < diff {
				lineSb.WriteRune(' ')
			}
		}
		stairs.WriteString(lineSb.String() + "\n")
	}
	fmt.Println(stairs.String())
}

func plusMinus(arr []int32) {

	arrSize := len(arr)
	pos := 0
	neg := 0
	zero := 0

	for i := 0; i < len(arr); i++ {

		if arr[i] > 0 {
			pos = pos + 1
		}
		if arr[i] < 0 {
			neg = neg + 1
		}
		if arr[i] == 0 {
			zero = zero + 1
		}
	}

	x := float32(pos) / float32(arrSize)
	y := float32(neg) / float32(arrSize)
	z := float32(zero) / float32(arrSize)
	fmt.Printf("%.6f\n", x)
	fmt.Printf("%.6f\n", y)
	fmt.Printf("%.6f\n", z)

}

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	var counter int64
	j := 0
	for range ar {
		counter = counter + ar[j]
		j++
	}
	return counter
}

// ugh ugly code
func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var validArrPos []int

	for i := 0; i < len(arr); i++ {
		//fmt.Println(len(arr[i]))
		if len(arr[i]) == 3 {
			validArrPos = append(validArrPos, i)
		}
	}
	var numOne int
	var numTwo int
	var numThree int
	var reNumOne int
	var reNumThree int
	for j := 0; j < len(validArrPos); j++ {
		fmt.Println(arr[validArrPos[j]])

		for k := 0; k < len(validArrPos); k++ {
			if j == 0 {
				fmt.Println(arr[validArrPos[j]][0])
				numOne = int(arr[validArrPos[j]][0])
				reNumOne = int(arr[validArrPos[j]][2])
			}
			if j == 1 {
				fmt.Println(arr[validArrPos[j]][1])
				numTwo = int(arr[validArrPos[j]][1])

			}
			if j == 2 {
				fmt.Println(arr[validArrPos[j]][2])
				numThree = int(arr[validArrPos[j]][2])
				reNumThree = int(arr[validArrPos[j]][0])

			}

		}

	}
	var forwardAns int = numOne + numTwo + numThree
	var reverseAns int = reNumOne + numTwo + reNumThree
	var totalAns int32 = int32((forwardAns - reverseAns) * -1)

	fmt.Println(totalAns)
	return totalAns
}
