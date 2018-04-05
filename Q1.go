/* Justin Huynh
 * 7745112
 * CSI2120
 */
 
package main

import (
	"fmt"
	"errors"
	"math"
	"os"
)

func AbsDiff(sliceA, sliceB []float32) (res []float32, err error) {
	//Confirm that slice lengths are the same
	if len(sliceA) != len(sliceB){
		err = errors.New("Slices are not the same length")
		return nil, err
	}
	
	//Find the absolute difference
	for i:= 0; i < len(sliceA); i++ {
		num := float32(math.Abs(float64(sliceA[i]) - float64(sliceB[i])))
		res = append(res, num)
	}
	
	
	return res, nil
	}


func main() {
	
    var (
        next float32
        sliceA []float32
		sliceB []float32
		cont string
    )
	
	//Get the first slice
	fmt.Println("Please enter a slice. (Press enter to end slice)")
    for {
        n, _:= fmt.Scanf("%f", &next)
        if n == 0 {
            break
        }
        sliceA = append(sliceA, next)
    }
	
	for {
		fmt.Printf("Previous slice: ")
		fmt.Printf("%v", sliceA)
		fmt.Println("\nEnter another slice of floating point numbers (Press enter to end slice)")
		sliceB = nil //Reset the slice values
		
		for {
			n, _:= fmt.Scanf("%f", &next)
			if n == 0 {
				break
			}
			sliceB = append(sliceB, next)
		}
		
		res, err := AbsDiff(sliceA, sliceB)
		
		if err == nil && res != nil{
			fmt.Printf("Result: ")
			fmt.Printf("%v", res)
			fmt.Println("q to quit (Anything else to continue):")
			fmt.Scanf("%v", &cont)
			sliceA = sliceB
			
			if cont == "q" {
				os.Exit(3)
			}
		} else {
			fmt.Println(err)
			//fmt.Println("Slices are not the same length")
			fmt.Println("q to quit (Anything else to continue):")
			fmt.Scanf("%v", &cont)
			
			if cont == "q" {
				os.Exit(3)
			}
		}
	}
}


