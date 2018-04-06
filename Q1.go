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

func AbsDiff(sliceA, sliceB []float32, version int) (res []float32, err error) {

	if version == 0{
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
	
	}else if version == -1 {
		//Find the shorter slice
		if len(sliceA) < len(sliceB){ //SliceA is shorter 
			//Add 0's to slice A so that both slices are same length
			for i:= len(sliceA); i < (len(sliceA) + len(sliceB) - len(sliceA)); i++{
				sliceA = append(sliceA, 0)
			}
			
			//Find the absolute difference
			for i:=0; i < len(sliceB); i++{
				num := float32(math.Abs(float64(sliceA[i]) - float64(sliceB[i])))
				res = append(res, num)
			}
		return res, nil
		
		}else { //SliceB is shorter
			//Add 0's to slice B so that both slices are same length
			for i:= len(sliceB); i < (len(sliceB) + len(sliceA) - len(sliceB)); i++{
				sliceB = append(sliceB, 0)
			}
			
			//Find the absolute difference
			for i:=0; i < len(sliceA); i++{
				num := float32(math.Abs(float64(sliceA[i]) - float64(sliceB[i])))
				res = append(res, num)
			}
		return res, nil
		}
		
	}else if version == 1 {
		//Find the shorter slice
		var lengthShortest int
		
		if len(sliceA) < len(sliceB){
			lengthShortest = len(sliceA)
		} else {
			lengthShortest = len(sliceB)
		}
		//Find the absolute difference
		for i:=0; i < lengthShortest; i++{
			num := float32(math.Abs(float64(sliceA[i]) - float64(sliceB[i])))
			res = append(res, num)
		}
		return res, nil	
	}
	
	return nil, err

}



func main() {
	
    var (
        next float32
        sliceA []float32
		sliceB []float32
		cont string
		version int
    )
	
	//Ask for the version
	fmt.Println("Enter the version. [-1, 0, 1]")
	fmt.Scanf("%d\n", &version)
	
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
		
		res, err := AbsDiff(sliceA, sliceB, version)
		
		if err == nil && res != nil{
			fmt.Printf("Result: ")
			fmt.Printf("%v", res)
			fmt.Println("q to quit (Enter to continue):")
			fmt.Scanf("%v", &cont)
			sliceA = sliceB
			
			if cont == "q" {
				os.Exit(3)
			}
		} else {
			fmt.Println(err)
			//fmt.Println("Slices are not the same length")
			fmt.Println("q to quit (Enter to continue):")
			fmt.Scanf("%v", &cont)
			
			if cont == "q" {
				os.Exit(3)
			}
		}
	}
}


