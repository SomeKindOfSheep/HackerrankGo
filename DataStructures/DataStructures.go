package main

import "fmt"

func main() {

	//ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	arr := [][]int32{
		{1},
		{1, 2, 3},
		{4, 5, 6},
		{4, 5},
		{4},
		{9, 8, 9},
	}

	diagonalDifference(arr)
	//fmt.Println(aVeryBigSum(ar))xx
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
