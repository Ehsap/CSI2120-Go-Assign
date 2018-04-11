/* Justin Huynh
 * 7745112
 * CSI2120
 * Go assignment
 * Question 3
 */
 
package main

import (
	"fmt"
	"math"
	"math/rand"
	//"os"
)

func RandomArray(len int) []float32 {
	array := make([]float32, len)
	for i := range array {
		array[i] = rand.Float32()
	}
	return array
}

func Process(arr []float32, c chan float32){
	//Split each array into two equal-sized sub-arrays
	
	right := arr[len(arr)/2:]
	left := arr[:len(arr)/2]

	res := AbsDiff(left, right)
	
	var sum float32
	for i := range res{
		sum += res[i]
	}
	
	c <- sum

}

func main() {    
	rand.Seed(100) // use this seed value  
	out := make(chan float32, 1000) //Increase buffer capacity
	defer close(out) 

	var sumTemp float32
	var res []float32
 
	for i := 0; i<1000 ; i++ {      
		a:= RandomArray(2*(50+rand.Intn(50)))       
		go Process(a,out)
		// read here the results of the processing
		sumTemp = <- out
		res = append(res, sumTemp)
	} 
 
   // and sum these results 
   	var sum float32
   for i := range res{
		fmt.Println(res[i])
		sum += res[i]
	}
   
   fmt.Println("THE SUM OF SUMS IS:")
   fmt.Println(sum)   
} 

func AbsDiff(sliceA, sliceB []float32) (res []float32) {
	//Find the absolute difference
	for i:= 0; i < len(sliceA); i++ {
		num := float32(math.Abs(float64(sliceA[i]) - float64(sliceB[i])))
		res = append(res, num)
	}	
	return res

}

