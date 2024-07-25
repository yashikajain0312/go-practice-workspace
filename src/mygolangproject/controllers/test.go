// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"sort"
// 	"encoding/json"
// 	"sync"
// 	"time"
// 	"net/http"
// 	"io/ioutil"
// )
	
// func main() {
// 	// interfaceFunction()
// 	//questionOne()
// 	//questionTwo()
// 	//questionThree()
// 	//questionFour()
// 	//questionFive()
// 	//questionSix()
// 	//mapQestionOne()
// 	//mapQestionTwo()
// 	//mapQuestionThree()
// 	//specialQuestion()
// 	//mapQestionFour()
// 	//mapQestionFive()
// 	//mapQuestionSix()
// 	//mapQuestionSeven()
// 	//specialQuestionTwo()
// 	//specialQuestionThree()
// 	//GoRoutineQuestionOne()
// 	//channelBasicQuestion()
// 	//producerConsumerProblem()
// 	//removeElementFromSlice()
// 	//StringReverseQuestion()
// 	//factorialQuestion()
// 	Channel01()
// 	//fibonacciQuestion()
// 	//GRQuestionOne()
//  //printingNumbers()
// //generateAndPrintNumbers()
// //CharFrequncyQuestion()
// //websiteScrappingGoroutine()
// }

// // interface.
// func interfaceFunction() {
// 	val:= map[string]interface{}{
// 		"name":1,
// 		"ag3":"hh",
// 		"address": true,
// 	}
// 	fmt.Println("val", val)
// }

// // Write a function that takes a slice of integers and returns
// // a new slice containing only the unique elements in the original slice.
// func sliceContainingUniqueElements(originalSlice []int) []int  {
//     newUniqueSlice := []int{}
// 	uniqueMap := make(map[int]bool)

// 	for _, v := range originalSlice {
// 		if !uniqueMap[v] {
// 			uniqueMap[v] = true
// 			newUniqueSlice = append(newUniqueSlice, v)
// 		}	
// 	}
// 	return newUniqueSlice
// }

// func questionOne() {
// 	originalSlice := []int{1,2,3,4,5,6,7,3,2,10}
// 	fmt.Println("original slice", originalSlice)
// 	newUniqueSlice := sliceContainingUniqueElements(originalSlice) 
// 	fmt.Println("Slice containing unique values", newUniqueSlice)
// }

// // Write a function that calculates the sum of all elements in a slice of integers.
// func calculateSum(slice1 []int) int {
// sum := 0
// for _, v := range slice1 {
// 	sum += v
// }
// return sum
// }

// func questionTwo() {
// slice1 := []int{1,4,5,6,3,4,8}

// result := calculateSum(slice1)
// fmt.Println("sum of all elements of a slice", result)
// }

// //Implement a function that reverses the elements of a slice in-place. The original slice should be modified.
// func reverseOriginalSlice(originalSlice []int) []int {
// 	l := len(originalSlice)

// 	for i:=0; i< l/2; i++{
// 		originalSlice[i], originalSlice[l-i-1] = originalSlice[l-i-1], originalSlice[i]
// 	}
// 	return originalSlice
// }

// func questionThree() {
// 	originalslice := []int{1,2,3,4,5,6,7,8,9,10}

// 	result := reverseOriginalSlice(originalslice)
// 	fmt.Println("reversed original slice", result)
// }

// // Create a function that finds the maximum and minimum values in a slice of integers and returns them.
// func findMaxAndMin(mySlice []int) (int, int){
// 	max := mySlice[0]
// 	min := mySlice[0]

// 	for _, v := range mySlice {
// 		if v > max {
// 			max = v
// 		}
// 		if v < min {
// 			min = v
// 		}
// 	}

//   return max, min
// }

// func questionFour() {
// 	mySlice := []int{6,8,3,2,9,19,2,67}

// 	max, min := findMaxAndMin(mySlice)
// 	fmt.Println("max = ", max)
// 	fmt.Println("min = ", min)
// }

// // Write a function that takes two slices of integers and merges them into a single sorted slice.
// func mergeSlices(slice1 []int, slice2 []int) []int {

// 	newSlice := []int{}
// 	newSlice = append(slice1, slice2...)
// 	return newSlice

// }

// func questionFive() {

// 	slice1 := []int{1,2,3}
// 	slice2 := []int{4,5,6}

// 	result := mergeSlices(slice1, slice2)
// 	fmt.Println("merged slices", result)
// }

// // Create a function that checks if a given slice of 
// // integers is a palindrome (reads the same forwards and backward).
// func isPalindrome(slice []int) bool {
//     l := len(slice)

// 	for i := 0; i < l/2; i++ {
// 		if slice[i] != slice[l-i-1] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func questionSix()  {
// 	slice := []int{1,2,3,2,1}

// 	result := isPalindrome(slice)
// 	fmt.Println("is slice palindrome:", result)
// }

// // Questions related to map string interface

// //Write a function that takes a map with string keys and interface{} values
// // and counts the occurrences of each distinct value. Return the counts in 
// //a new map[string]int where the keys are the distinct values and the values are their counts.
// func checkOccurence(data map[string]interface{}) map[string]int{
//     newMap := make(map[string]int)

// 	for _, val := range data {
	 
// 		stringValue := fmt.Sprintf("%v", val)

// 		newMap[stringValue]++
// 	}

// 	return newMap
// }

// func mapQestionOne() {
// 	data := map[string]interface{}{
// 		"a": 1,
// 		"b": "hello",
// 		"c": "world",
// 		"d": 2,
// 		"e": "hello",
// 		"f": 1,
// 		"g": true,
// 	}

// 	result:= checkOccurence(data)
// 	fmt.Println("result",result)
// }

// // Write a function that takes a map with string keys and interface{} values and returns a new 
// //map containing only the key-value pairs where the value is of a specific type (e.g., int, string, float64).
// func findSpecifiedTypeKVPair(data map[string]interface{}) map[string]interface{} {

// 	newMap := make(map[string]interface{})

// 	for k, v := range data {

//         valueType := reflect.TypeOf(v)

// 		if valueType == reflect.TypeOf(int(0)) || valueType == reflect.TypeOf(float64(0.0)) || valueType == reflect.TypeOf("") {
// 			newMap[k] = v
// 		}
// 	}

// 	return newMap
// }

// func mapQestionTwo() {
// 	data := map[string]interface{}{
// 		"one": 1,
// 		"two": "twoo",
// 		"three": 3,
// 		"four": 4.0,
// 		"five": true,
// 		"six": "sixx",
// 	}

// 	result := findSpecifiedTypeKVPair(data)
// 	fmt.Println("result",result)
// }

// // Write a function that takes multiple maps with string keys and interface{} values and merges
// // them into a single map. If a key exists in multiple maps, prefer the value from the later maps.
// func mergedMaps(maps ...map[string]interface{}) map[string]interface{}{

// 	mergedMap := make(map[string]interface{})

// 	for _, mapp := range maps {
//        for k,v := range mapp {
// 		  mergedMap[k] = v
// 	   }
// 	}
// 	return mergedMap
// }

// func mapQuestionThree() {
// 	data1 := map[string]interface{}{
// 		"name": "yashi",
// 		"age": 10,
// 	}
// 	data2 := map[string]interface{}{
// 		"address": "kolar",
// 		"name": "soumya",
// 		"gender": "female",
// 	}
// 	data3 := map[string]interface{}{
// 		"gender": "male",
// 		"ok": true,
// 	}

// 	result := mergedMaps(data1, data2, data3)
// 	fmt.Println("result", result)
// }

// // check indexes of those students who are standing on wrong position according
// // to their height which should be in ascending order.
// func checkWrongHeightIndex(arr []int) []int{
//     oldArr := make([]int, len(arr))
// 	copy(oldArr, arr)

// 	resultArr := []int{}

// 	// sort.Ints sort the original array in ascending order.
// 	sort.Ints(arr)

// 	for i := 0; i<len(arr); i++ {
// 		if arr[i] != oldArr[i] {
// 			resultArr = append(resultArr, i)
// 		}
// 	}
	
// 	return resultArr
// }

// func specialQuestion() {

// 	arr := []int{8,10,6,7,12,5,3}
    
// 	result := checkWrongHeightIndex(arr)
// 	fmt.Println("result", result)
// }

// // Write a function that takes a map with string keys and interface{} values and converts 
// // it to a JSON string. Ensure that the resulting JSON includes proper type information for the values.
// func convertMapToJson(data map[string]interface{}) (string, error) {

// 	mapWithType := make(map[string]interface{})

// 	for key, val := range data {
// 		valType := reflect.TypeOf(val).String()

// 		valuemap := map[string]interface{}{
// 			"value": val,
// 			"type": valType,
// 		}

// 		mapWithType[key] = valuemap
// 	}

// 	// convert map into json
// 	jsonBytes, err := json.Marshal(mapWithType)
// 	if err != nil {
// 		return "", err
// 	}
	
// 	//convert json byte into json string
// 	jsonString := string(jsonBytes)

// 	return jsonString, nil

// }

// func mapQestionFour() error {
// 	data := map[string]interface{}{
// 		"name": "yashika",
// 		"age": 22,
// 		"gender": "female",
// 	}

// 	result, err := convertMapToJson(data)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("result", result)
// 	return nil
// }

// // Write a function that takes a map with string keys and interface{} values 
// // and converts all the values to strings. Return the modified map.
// func convertMapValuesToString(data map[string]interface{}) map[string]string {
//     modifiedMap := make(map[string]string)
    
// 	for key, val := range data {
// 			strVal := fmt.Sprintf("%v", val)
// 			modifiedMap[key] = strVal
// 	}

// 	return modifiedMap
// }

// func mapQestionFive() {
// 	data := map[string]interface{}{
// 		"name": "yashika",
// 		"age": 22,
// 		"gender": "female",
// 	}

// 	result :=  convertMapValuesToString(data)
// 	fmt.Println("result", result)
// }

// // Write a function that takes a map with string keys and interface{} 
// // values and returns a slice containing all the values from the map.
// func sliceFromMapValues(data map[string]interface{}) []interface{} {

// 	slice := []interface{}{}

// 	for _, val := range data {
// 			slice = append(slice, val)
// 	}

// 	return slice
// }

// func mapQuestionSix() {
// 	data := map[string]interface{}{
// 		"name": "yashika",
// 		"age": 22,
// 		"gender": "female",
// 	}

// 	slice :=  sliceFromMapValues(data)
// 	fmt.Println("result", slice)
// }

// // Write a function that takes a map with string keys and interface{} values and checks
// // if a specific key exists in the map. Return true if the key exists, otherwise return false.
// func isKeyExists(data map[string]interface{}, inputKey string) bool {

// 	for key, _ := range data {
// 		if key == inputKey {
// 			return true
// 		}
// 	}

// 	return false
// }

// func mapQuestionSeven() {
// 	data := map[string]interface{}{
// 		"name": "yashika",
// 		"age": 22,
// 		"gender": "female",
// 	}

// 	isExists := isKeyExists(data, "name")
// 	fmt.Println("result:", isExists)
// }


// // special question
// func specialQuestionTwo() {
// 	originalData := `
// 		{
// 			"boardInfo": {
// 				"data": "Random Data"
// 			},
// 			"things": [
// 				{
// 					"data": "Things Data 1"
// 				},
// 				{
// 					"data": "Things Data 2"
// 				},
// 				{
// 					"data": "Things Data 3"
// 				},
// 				{
// 					"data": "Things Data 4"
// 				}
// 			]
// 		}
// 	`

// 	newData := `
// 		{
// 			"data": "New Things Data"
// 		}
// 	`
// 	/*
// 		Objective 1: Check the length of "things" in originalData
// 		Objective 2: Append the "newData" in "things" of originalData
// 		Objective 3: Verify by checking the length of "things"

// 	*/

// 	var originalMapData map[string]interface{}
// 	var newMapData map[string]interface{}

// 	// unmarshal data(convert json data into go data type i.e map)
// 	_ = json.Unmarshal([]byte(originalData), &originalMapData)
// 	_ = json.Unmarshal([]byte(newData), &newMapData)
	
// 	things := originalMapData["things"].([]interface{})

// 	fmt.Println(len(things))
// 	things = append(things, newMapData)
// 	fmt.Println(len(things))
// }

// // nesting questions
// func specialQuestionThree() {
// 	nestedMap := map[string]interface{}{
// 		"key1": "value1",
// 		"key2": map[string]interface{}{
// 			"subkey1": "subvalue1",
// 			"subkey2": 123,
// 			"subkey3": true,
// 		},
// 		"key3": []string{"apple", "banana", "orange"},
// 	}

// 	// Accessing Top-Level Values:
// 	// How do you access the value associated with the key "key1"?
// 	// How do you access the value associated with the key "key3"?

// 	val1 := nestedMap["key1"]
// 	fmt.Println("val1", val1)
	
// 	val2 := nestedMap["key3"]
// 	fmt.Println("val1", val2)

// 	// Accessing Nested Values:
// 	//How do you access the value associated with the key "subkey1" 
// 	//inside the nested map associated with the key "key2"?
// 	// How do you access the value associated with the key "subkey2" 
// 	// inside the nested map associated with the key "key2"?

// 	val3 := nestedMap["key2"].(map[string]interface{})["subkey1"]
// 	fmt.Println("val3", val3)

// 	val4 := nestedMap["key2"].(map[string]interface{})["subkey2"]
// 	fmt.Println("val3", val4)

// 	// Accessing Values in Slices:
// 	// How do you access the first element of the slice associated with the key "key3"?
// 	// How do you access the second element of the slice associated with the key "key3"?

// 	val5 := nestedMap["key3"].([]string)[0]
// 	fmt.Println("val1", val5)

// 	val6 := nestedMap["key3"].([]string)[1]
// 	fmt.Println("val1", val6)

// 	// How would you append the string "grape" to the slice associated with the key "key3"?
// 	slice := nestedMap["key3"].([]string)
// 	nestedMap["key3"] = append(slice, "grape")
// 	fmt.Println("slice", slice)

// 	// Can you append a new key-value pair "subkey4": "newSubValue" to the nested
// 	// map associated with the key "key2"?
// 	nestedKVPair := nestedMap["key2"].(map[string]interface{})
// 	nestedKVPair["subkey4"] = "newSubValue"
// 	fmt.Println("nestedKVPair", nestedKVPair)

// 	// Create a new nested map with the key "newKey" and add a key-value 
// 	// pair "newSubkey": "newSubValue" to it.
// 	newNestedKVPair := map[string]interface{}{
// 		"newSubKey": "newSubValue",
// 	}
//     nestedMap["newKey"] = newNestedKVPair
// 	fmt.Println("nestedMap", nestedMap)
// }

// // GoRoutines and channel questions

// // Write a program that uses goroutines to print "Hello, World!" concurrently.

// func GoRoutineQuestionOne() {
// 	var wg sync.WaitGroup
// 	numGoroutines := 5
	
// 	wg.Add(numGoroutines)
	
// 	for i := 0; i < numGoroutines; i++ {
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("Hello, World!")
// 		}()
// 	}
	
// 	wg.Wait()
	
// }

// // Write a program that creates a channel, sends a message on the channel, 
// // and then receives and prints the message.
// func sendingMessage(ch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	msg := <-ch
// 	fmt.Println("meg", msg)
// }

// func channelBasicQuestion() {
// 	ch := make(chan string)
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go sendingMessage(ch, &wg)
// 	ch <- "hello"
// 	wg.Wait()
// }

// // Implement a simple producer-consumer problem using goroutines and channels.
// // The producer goroutine should generate numbers and send them to a channel, 
// //and the consumer goroutine should receive and print the numbers.
// func producerConsumerProblem() {
// 	numbers := make(chan int)

// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go producer(numbers, &wg)

// 	wg.Add(1)
// 	go consumer(numbers, &wg)

// 	wg.Wait()
// }

// func producer(numbers chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 5; i++ {
// 		numbers <- i 
// 	}
// 	close(numbers) 
// }

// func consumer(numbers <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range numbers {
// 		fmt.Println("Received:", num)
// 	}
// }

// func removeElementFromSlice() {
// 	slice := []int{1,5,6,3,9,23,8}
// 	indexToRemove := 2

// 	slice1 := slice[:indexToRemove]
// 	slice2 := slice[indexToRemove+1:]

// 	newSlice := append(slice1, slice2...)

	
// 	fmt.Println("new slice: ", newSlice)
// }


// // Reverse a String: Write a function to reverse a string in Go.
// func reverseString(str string) string {

// 	// Convert the string to a slice of runes
// 	runes := []rune(str)
// 	length := len(runes)

//     for i := 0; i < length/2; i++ {
//     runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
//     }

// 	reversedString := string(runes)

// 	return reversedString
// }

// func StringReverseQuestion() {

// 	stringVal := "Yashika"

// 	result := reverseString(stringVal)
// 	fmt.Println("result", result)
// }

// // Factorial of a Number: Write a function to calculate the factorial of a number in Go.
// func findFactorial(number int) int{

// 	factorial := 1

// 	for i := 1; i <= number; i++ {
// 		factorial = factorial*i
// 	}

// 	return factorial
// }

// func factorialQuestion() {
// 	number := 5

// 	result := findFactorial(number)
// 	fmt.Println("result", result)
// }

// // optimize time question
// func makeRange(min, max int) []int {
// 	a := make([]int, max-min+1)
// 	for i := range a {
// 		a[i] = min + i
// 	}
// 	return a
// }

// func apiCall(ch chan int) {
// 	fmt.Println("aaaa")
// 	for v := range ch {
// 		fmt.Println("API call for", v, "started")
// 		//time.Sleep(time.Millisecond)
// 	}

// }

// func Channel01() {
// 	numArray := makeRange(0, 500)
// 	ch := make(chan int)
// 	start := time.Now()
// 	go apiCall(ch)
	
// 	for _,v := range numArray {
// 		ch <- v
// 	}
//     close(ch)
	
// 	elapsed := time.Since(start)
// 	fmt.Printf("Time taken %s", elapsed)
// }

// // Fibonacci Series: Write a function to generate the Fibonacci series up to a given number of terms in Go.
// func fibonacciSeries(terms int) []int {
// 	// Initialize the Fibonacci series with the first two terms
// 	fib := []int{0, 1}

// 	// Generate additional terms if terms > 2
// 	if terms > 2 {
// 		for i := 2; i < terms; i++ {
// 			nextTerm := fib[i-1] + fib[i-2]
// 			fib = append(fib, nextTerm)
// 		}
// 	}

// 	return fib
// }

// func fibonacciQuestion() {
// 	numberOfTerms := 10
// 	fibSeries := fibonacciSeries(numberOfTerms)
// 	fmt.Println("Fibonacci Series up to", numberOfTerms, "terms:", fibSeries)
// }

// // Question: Write a program to calculate the sum of numbers from 1 to N concurrently
// // using goroutines and channels.
// // This program should accept an integer N as input and calculate the sum of numbers from 1 to N
// // concurrently using goroutines and channels. The main goroutine should wait for all goroutines
// // to finish and then print the final sum.
// func GRQuestionOne() {
	
// 	var wg sync.WaitGroup
	
// 	numbers := make(chan int)
// 	sum := make(chan int, 1)

// 	wg.Add(1)
// 	go generateNumbers(numbers, &wg)

// 	wg.Add(1)
// 	go sumNumbers(numbers, &wg, sum)

// 	wg.Wait()

// 	fmt.Println("sum is:", <-sum)
// }

// func generateNumbers(numbers chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i:=1; i<=10; i++ {
// 		numbers <- i
// 	}

// 	close(numbers)
// }


// func sumNumbers(numbers chan int, wg *sync.WaitGroup, sum chan int) {
// 	defer wg.Done()
// 	total := 0
// 	for num := range numbers {
// 		total += num
// 	}
	
// 	sum <- total
// }

// // bacany question interview
// func bacanyQuestion() {
// 	arr := []int{1,4,6,8,9,10,34,55}

// 	result := findSeconLargestElement(arr)
// 	fmt.Println("result", result)
	
// }

// func findSeconLargestElement(arr []int) int{

// 	l := len(arr)

// 	for i:=0; i<l-1; i++ {
// 		for j:=0; j<l-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}

// 	result := arr[l-2]
// 	return result
// }

// func callGetAPI() {
// 	url := "https://color.serialif.com/aquamarine"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	defer resp.Body.Close()

// 	dataBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	content := string(dataBytes)
// 	fmt.Println("api :", content)
// }

// // antino labs interview question
// func main() {

// 	var arr []int

// 	for i:=1; i<=100; i++ {
// 		arr = append(arr, i)
// 	}

// 	var wg sync.WaitGroup
// 	result := make(chan int, 1)

// 	sum := 0
	
// 	fmt.Println("working")
// 	for i:=0; i<=100; i=i+11{
// 		slice := arr[i:i+11]
// 		wg.Add(1)
// 		go SumOfElement(slice, &wg, result)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()

// 	for v := range result {
// 		sum = sum + v
// 	}
	
// 	fmt.Println("result", sum)
// }

// func SumOfElement(slice []int, wg *sync.WaitGroup, result chan int) {
// 	defer wg.Done()
//  	sum := 0

// 	for _, v := range slice {

// 		sum += v
		
// 	}
// 	result <- sum
// }


// interview questions.
// type sampleWeight struct {
//     id int
//     weight int
// }

// // [{2 30} {1 50} {3 60} {5 70}  {4 90}{6 140}] 

// func main() {

// 	ObjOne := []sampleWeight {
 
// 		sampleWeight {
// 			id: 1,
// 			weight: 50,
// 		},
	 
// 		sampleWeight {
// 			id: 3,
// 			weight: 60,
// 		},
	 
// 		sampleWeight {
// 			id: 5,	 
// 			weight: 70, 
// 		},
// 	}
	
// 	ObjTwo := [] sampleWeight {
	 
// 		sampleWeight {	 
// 			id: 2,	 
// 			weight: 30,	 
// 		},
	 
// 		sampleWeight {	 
// 			id: 4, 
// 			weight: 90,	 
// 		},
	 
// 		sampleWeight {	 
// 			id: 6,	 
// 			weight: 140,	 
// 		},
	 
// 	}


// 	combined := append(ObjOne, ObjTwo...) // Combine the slices

// 	// Define custom sorting
// 	sort.Slice(combined, func(i, j int) bool {
// 		return combined[i].weight < combined[j].weight
// 	})

// 	// Print the sorted slice
// 	fmt.Println("Merged and sorted by weight:")
// 	fmt.Println(combined)
// }

// // Write a function that counts the number of vowels and consonants in a given string.
// func practiceQues() {
// 	str := "Aashika"

// 	vCount, cCount := CountVowelsAndConsonants(str)
// 	fmt.Println("result", vCount, cCount)
// }

// func CountVowelsAndConsonants(str string) (int, int) {
//     vCount := 0
// 	cCount := 0

// 	str = strings.ToLower(str)
// 	for _, v := range str {
// 		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
// 			vCount++
// 		} else {
// 			cCount++
// 		}
// 	}

// 	return vCount, cCount
// }

// //Input: {0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
// //Output: {0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2} pace problem to sort the array in one iteration.

// //Write a program that creates two goroutines. One goroutine should print even numbers from 0 to 10,
// //and the other goroutine should print odd numbers from 1 to 9. Ensure that the main goroutine waits for both goroutines to finish before exiting.
// func printingNumbers() {
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go printEvenNumbers(&wg)
// 	go printOddNumbers(&wg)

// 	wg.Wait()
// }

// func printEvenNumbers(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i:=0; i<=10; i+=2 {
// 		fmt.Println("even no:", i)
// 	}
// }

// func printOddNumbers(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i:=1; i<10; i+=2 {
// 		fmt.Println("odd no:", i)
// 	}
// }

// //Create a program that simulates a message passing system between two goroutines using channels. 
// //One goroutine should generate a sequence of integers and send them to another goroutine through a channel. 
// //The receiving goroutine should print the received numbers.

// func generateAndPrintNumbers() {

// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go generateNumbers(&wg)

// 	wg.Wait()
// }

// func generateNumbers(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	ch := make(chan int)
// 	go printNumber(ch) 

// 	for i:=0; i<=20; i++ {
// 		ch <- i
// 	}

// 	close(ch)
// }

// func printNumber(ch chan int) {
// 	for num := range ch {
// 		fmt.Println("num:", num)
// 	}
// }

// // Write a Go program that calculates the frequency of each character in a
// // given string concurrently using goroutines. Print the frequency of each character.

// func CharFrequncyQuestion() {

// 	var wg sync.WaitGroup

// 	str := "yashika"
// 	charFrequency := make(map[rune]int)

// 	for _, runeChar := range str {
// 		wg.Add(1)
// 		go frequencyOfCharacter(&wg, runeChar, &charFrequency) 
// 	}

// 	wg.Wait()

// 	for runeCharac, v := range charFrequency {
// 		fmt.Println("char:", string(runeCharac) , "frequency:", v)
// 	}
// }

// func frequencyOfCharacter(wg *sync.WaitGroup, runeChar rune, charFrequency *map[rune]int) {
// 	defer wg.Done()
// 	(*charFrequency)[runeChar]++
// }

// //Create a Go program that concurrently scrapes data from multiple 
// // websites and aggregates the results. Print the combined data.

// func websiteScrappingGoroutine() {

// 	websites := []string{
// 		"https://example.com",
// 		"https://example.org",
// 		"https://example.net",
// 	}

// 	var wg sync.WaitGroup
// 	result := make(chan string)

// 	for _, url := range websites {
// 		wg.Add(1)
// 		go scrapeData(&wg, url, result)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()


// 	for data := range result {
// 		fmt.Println(data)
// 	}
// }

// func scrapeData(wg *sync.WaitGroup, url string, result chan string) {

// 	defer wg.Done()

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal("error", err)
// 	}

// 	defer resp.Body.Close()

// 	byteData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal("error", err)
// 	}

// 	result <- string(byteData)
// }

// // wednesday solution question
// func isSubstring(mainString, subString string) bool {
//     mainLen := len(mainString)
//     subLen := len(subString)

//     if subLen > mainLen {
//         return false // Substring cannot be longer than main string
//     }

//     for i := 0; i <= mainLen-subLen; i++ {
//         if mainString[i:i+subLen] == subString {
//             return true
//         }
//     }

//     return false
// }

// func findSubstringQues() {
//     mainString := "yashika"
//     subString := "ika"

//     if isSubstring(mainString, subString) {
//         fmt.Printf("'%s' is a part of '%s'\n", subString, mainString)
//     } else {
//         fmt.Printf("'%s' is not a part of '%s'\n", subString, mainString)
//     }
// }

// Write a function to count the number of words in a 
// given string, where words are separated by spaces.
// func wordCountQues() {
// 	str := "Hello my name is yashika "

// 	noOfWords := CountWords(str)
// 	fmt.Println("noOfWords", noOfWords)
// }

// func CountWords(str string) int {

// 	wordCount := 0
// 	inWord := true

// 	for _, v := range str {
// 		if v == ' ' {
// 			if inWord {
// 				wordCount++
// 				inWord = false
// 			}
// 		} else {
// 			inWord = true
// 		}
// 	}

//     // Increment the word count if the string doesn't end with a space
// 	if inWord {
// 		wordCount++
// 	}

// 	return wordCount
// }


// full code
// import (
// 	"database/sql"
// 	"log"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// )

// var db *sql.DB

// // post api which will take request contains book id, book name and author.
// func callPOSTApi() {
// 	connStr := "postgres://postgres:ryd12@localhost:5432/postgres?sslmode=disable"

// 	// creating database connection.
// 	var err error
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		 log.Fatal("error connecting to the database:", err)
// 	}

// 	defer db.Close()

// 	// creating table.
// 	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS book (
// 		id SERIAL PRIMARY KEY,
// 		name VARCHAR(255),
// 		author VARCHAR(255)
// 	)`)
// 	if err != nil {
// 		log.Fatal("Error creating book table: ", err)
// 	}

// 	route := gin.Default()

// 	route.POST("/book/add", BookHandler)

// 	route.Run(":3000")
// }

// type BookRequest struct {
// 	BookID   int  `json:"id"`
// 	BookName string  `json:"name"`
// 	Author   string  `json:"author"`
// }

// func BookHandler(c *gin.Context) {

// 	var req BookRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err := db.Exec("INSERT INTO book (id, name, author) VALUES ($1, $2, $3)", req.BookID, req.BookName, req.Author)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting book into database"})
// 		return
// 	}
	
// 	c.JSON(http.StatusOK, "book added successfully")
// }

// // coditation solution question (last round)
// // printPattern(3)
// // 1
// // 1 2
// // 1 2 3
// // 1 2
// // 1

// func coditationquestion() {
// 	number := 3

// 	printPattern(number)
// }

// func printPattern(num int) {

// 	for i:=1; i<=num; i++{
// 		for j:=1; j<=i; j++ {
// 			fmt.Print(j)
// 		}
// 		fmt.Println("")
// 	}

	
// 	for k:=1; k<=num-1; k++ {
// 		for l:=1; l<=num-k; l++ {
// 			fmt.Print(l)
// 		}

// 		fmt.Println("")

// 	}
// }

// func WiJUnlge() {

// 	key := 26

// 	matrix := [][]int{
//         {0, 1, 2, 3},
//         {4, 5, 6, 7},
// 		{56, 26, 78, 55},
// 		{22, 76, 66, 37},
//     }

// 	result1, result2 := findPositionOfKey(matrix, key)
// 	fmt.Println("result1", result1)
// 	fmt.Println("result2", result2)

// }

// func findPositionOfKey ( matrix [][]int, key int) (int, int) {
     
// 	for i, matr := range matrix {

// 	for k,v := range matr {
// 	 if v == key {
// 	 return i, k
// 	}
// 	}
// 	}
//    return -1, -1
// }

// print right triangle
// *
// * *
// * * *
// * * * *

// func PrintRightTriangle() {
// 	rows := 4
// 	printStar(rows)
// }

// func printStar(rows int) {
// 	for i:=0; i<rows; i++ {
// 		for j:=0; j<i+1; j++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println("")
// 	}
// }

// Print a pyramid star pattern.
// func pyramidStarPattern() {
// 	rows := 4
// 	printStar(rows)
// }

// func printStar(rows int) {

// 	var space int

//     space = rows - 1

//     for i := 1; i <= rows; i++ {
//         for j := 1; j <= space; j++ {
//             fmt.Print(" ")
//         }
//         space--

//         for k := 1; k <= 2*i-1; k++ {
//             fmt.Print("*")
//         }

//         fmt.Println()
//     }
// }

// Print an inverted pyramid star pattern.
// func invertedPyramidStarPattern() {
// 	rows := 4
// 	printStar(rows)
// }

// func printStar(rows int) {

// 	var space int

// 	for i:=0; i<rows; i++ {
		
// 		for j:=1; j<=space; j++ {
// 			fmt.Print(" ")
// 		}

// 		space++

// 		for k:=1; k<=rows*2-i-space; k++ {
// 			fmt.Print("*")
// 		}

// 		fmt.Println()
// 	}
// }

// Print a hollow rectangle star pattern.
// func hollowRectangleStarPattern() {
// 	rows := 4
// 	cols := 8
// 	printStar(rows, cols)
// }

// func printStar(rows int, cols int) {

// 	for i:=1; i<=rows; i++ {
// 			for j:=1; j<=cols; j++ {
// 				if i == 1 || i == rows || j == 1 || j == cols  {
// 					fmt.Print("* ")
// 				} else {
// 					fmt.Printf("  ")
// 				}
// 			}

// 		fmt.Println()
// 	}
// }

// find the second smallest or minimum number from the array.
// func CodeAegisQuestion() {
// 	arr := []int{-8, 3, -12, 7, 4, -2, 9, -12, 66, 89}

// 	secondMin := findSecondSmallestNumber(arr)
// 	fmt.Println("second min", secondMin)
// }

// func findSecondSmallestNumber(arr []int) int {

// 	min := arr[0]
// 	secondMin := arr[1]

// 	for i:=0; i<len(arr); i++ {

// 		if arr[i] < min {
// 			secondMin = min
// 			min = arr[i]
// 		} else if arr[i] != min && arr[i] < secondMin {
// 			secondMin = arr[i]
// 		} else {
// 			continue
// 		}
// 	}

// 	return secondMin

// }

// find the second largest number from the array.
// func SecondLargestNumQuestion() {
// 	arr := []int{60, 3, 12, 7, 4, 2, 9, 67, 66, 89}

// 	secondLargest := findSecondLargestNumber(arr)
// 	fmt.Println("second largest", secondLargest)
// }

// func findSecondLargestNumber(arr []int) int {

// 	max := arr[0]
// 	secondMax := arr[1]

// 	for i:=0; i<len(arr); i++ {

// 		if arr[i] > max {
// 			secondMax = max
// 			max = arr[i]
// 		} else if arr[i] != max && arr[i] > secondMax {
// 			secondMax = arr[i]
// 		} else {
// 			continue
// 		}
// 	}

// 	return secondMax

// }

// func SquareGoroutine() {
	
// 	var wg sync.WaitGroup
	
// 	numbers := make(chan int)

// 	wg.Add(1)
// 	go generateNumbers(numbers, &wg)

// 	wg.Add(1)
// 	go SquareAndPrintNum(numbers, &wg)

// 	wg.Wait()

// }

// func generateNumbers(numbers chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i:=1; i<=10; i++ {
// 		numbers <- i
// 	}
// 	close(numbers)
// }


// func SquareAndPrintNum(numbers chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range numbers {
// 		fmt.Printf("square of %d is %d", num, num*num)
// 		fmt.Println()
// 	}
// }

// //patelwar question
// //Q-1Given an integer array nums, move all 0's to the end of it while 
// //maintaining the relative order of the non-zero elements.
// //Note that you must do this in-place without making a copy of the array.
 
// //Example 1:
// //Input: nums = [0,1,0,3,12]
// //Output: [1,3,12,0,0]

// //Q-2Print hello word five times in go routine.
// func pantelwarGoroutineQuestion() {
// 	var wg sync.WaitGroup

// 	ch := make(chan string)


// 	for i:= 1; i<=5; i++ {
// 		wg.Add(1)
// 		go PrintHello(&wg, ch)
// 		ch <- "hello"
// 	}

	
// 	wg.Wait()
// 	close(ch)
// }

// func PrintHello(wg *sync.WaitGroup, ch chan string) {
// 	defer wg.Done()
// 	fmt.Println(<-ch)
// }

// // Print numbers from 1 to 10 using 2 goroutines. One printing even and the other printing odd numbers.

// // Task:
// // create a POST API in Go using the Gin framework. The API should accept a JSON payload with the following fields:
// // name: string
// // pan: string (PAN number)
// // Mobile: number (mobile number)
// // email: string (emailid)
// // The PAN number should follow the format of five letters, followed by four digits, followed by a letter (e.g., ABCDE1234F). The mobile number should be a 10-digit number and email.
// // Additionally, you need to implement middleware to validate the PAN number and mobile number before reaching the endpoint handler.
// // Requirements:
// // Create a Gin router with the POST endpoint /register.
// // Implement middleware to validate the PAN number and mobile number.
// // If the validation fails, the middleware should respond with an appropriate error message and status code.
// // If the validation passes, the endpoint should respond with a success message and status code.

// func CoditasQuestion() {


// 	r := gin.Default()

// 	r.Use(validationMiddleware())
// 	r.POST("/register", POSTApiHandler)

// 	r.Run(":5600")
// }

// func validationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var req POSTRequest
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			c.Abort()
// 			return
// 		}

// 		// Validate PAN number
// 		panRegex := regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]$`)
// 		if !panRegex.MatchString(req.Pan) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PAN number format."})
// 			c.Abort()
// 			return
// 		}

// 		// Validate mobile number
// 		mobileRegex := regexp.MustCompile(`^[0-9]{10}$`)
// 		if !mobileRegex.MatchString(req.Mobile) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number must be 10 digits."})
// 			c.Abort()
// 			return
// 		}

// 		// Pass on to the next middleware/handler
// 		c.Next()
// 	}
// }

// type POSTRequest struct {
// 	Name   string `json:"name" binding:"required"`
// 	Pan    string `json:"pan" binding:"required"`
// 	Mobile string `json:"mobile" binding:"required"`
// 	Email  string `json:"email" binding:"required,email"`
// }

// func POSTApiHandler(c *gin.Context) {

// 	var req POSTRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"errorrr": err.Error()})
// 		return
// 	}
	
// 	c.JSON(http.StatusOK, gin.H{"message": "Success"})
// }

// IBM question
// func IBM() {
// 	str1 := "net"
// 	str2 := "ten"

// 	isStringCharaterSame(str1, str2)
// }

// // meat team
// // note tone
// func isStringCharaterSame(str1 string, str2 string) {

// 	arr := []rune{}

// 	for _, v := range str1 {
// 		arr = append(arr, v)
// 	}


	
// 	fmt.Println("str2", str2)
	
	
// }

// Given an integer array nums and an integer k, return the k most frequent elements within the array.
//The test cases are generated such that the answer is always unique.
//You may return the output in any order.

// Example 1:

// Input: nums = [1,2,2,3,3,3], k = 2

// Output: [2,3]
// Example 2:

// Input: nums = [7,7], k = 1

// Output: [7]

// func AveshaQuestion() {
// 	nums := []int{1,1,1,-1,-1,2,2,2,2,2, 5,5,5,5,5,5,5,5}
// 	k :=  3

// 	FindMostFrequentNumber(nums, k)
// }

// func FindMostFrequentNumber(nums []int, ke int) {

// 	m := map[int]int{}

// 	for _,v := range nums {
// 		m[v]++
// 	}

// 	arr := []int{}

// 	for _, v := range m {
// 		arr = append(arr, v)
// 	} 


// 	for i:=0; i<len(arr)-1; i++ {
// 		for j:=0; j<len(arr)-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}

// 	l := len(arr)
// 	secondLargestVal := arr[l-ke]
//     key := 0

// 	for k, v := range m {

// 		if v == secondLargestVal {
// 			key = k
// 		}
// 	} 

// 	fmt.Println("finalOutput", key)
// }