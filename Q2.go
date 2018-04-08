/* Justin Huynh
 * 7745112
 * CSI2120
 * Question 2
 */
 
package main

import (
	"fmt"
//	"errors"
//	"math"
//	"os"
)

type Baking struct{
	bakeTime int
	coolTime int
	temperature int

}

type Item struct {
	weight int
}

type Bread struct {
	name string
	weight float32
	ingredients map[string]Item
	baking Baking
}

type Baker interface{
	shoppingList(ingredients map[string]Item)(shopList map[string]Item)
	printBakeInstructions()
	printBreadInfo()

}

func printBakeInstructions(){

}

func printBreadInfo(){

}

func NewBread()(*Bread){
	var bread Bread
	b := &bread

	b.name = "Whole Wheat"
	b.ingredients = make(map[string]Item)
	
	b.ingredients["whole wheat flour"] = Item{500}
	b.ingredients["yeast"] = Item{25}
	b.ingredients["salt"] = Item{25}
	b.ingredients["sugar"] = Item{50}
	b.ingredients["butter"] = Item{50}
	b.ingredients["water"] = Item{350}
	b.baking.bakeTime = 2
	b.baking.temperature = 180
	b.baking.coolTime = 1
	b.weight = 0
	
	for k := range b.ingredients {
		b.weight += float32(b.ingredients[k].weight)
	}
	return b
	
}

func NewBreadVariation(newName string, ingredientsPlus (map[string]Item), 
						ingredientsLess (map[string]Item)) (*Bread){
	
	b := NewBread()
	b.name = newName
	
	//Add ingredients
	if len(ingredientsPlus) > 0{
		for k, v := range ingredientsPlus{
			//If the ingredient already exists, just add the weight of the item
			if val, ok := b.ingredients[k]; ok{
				val.weight += ingredientsPlus[k].weight
			} else {
			//Ingredient doesn't exist, add to the ingredient list
			b.ingredients[k] = v
			}
		}
	}
	
	//Remove ingredients
	if len(ingredientsLess) > 0{
		for k := range ingredientsLess{
			//If the ingredient already exists, just subtract the weight of the item
			if val, ok := b.ingredients[k]; ok{
				val.weight -= ingredientsLess[k].weight
			} 
		}
	}
	
	//Recalculate the weight of all ingredients
	b.weight = 0
	for k := range b.ingredients {
		b.weight += float32(b.ingredients[k].weight)
	}
	return b
}

func main() {
	
	var ingMap = make(map[string]Item)
	ingMap["apples"] = Item{99}
	
	var ingMap2 = make(map[string]Item)
	ingMap["sugar"] = Item{50}

	b := NewBreadVariation("JustinToast", ingMap, ingMap2)

	fmt.Println(b.name)
	fmt.Println(b.ingredients["apples"].weight)
	fmt.Println(b.weight)
	//fmt.Println(b.ingredients["whole wheat flour"])

}
 
