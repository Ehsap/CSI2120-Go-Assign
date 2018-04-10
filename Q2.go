/* Justin Huynh
 * 7745112
 * CSI2120
 * Question 2
 */
 
package main

import (
	"fmt"
	"strconv"
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
	shoppingList(ingredientsHave map[string]Item)(ingredientsNeed map[string]Item, ingredientsLeft map[string]Item)
	printBakeInstructions()
	printBreadInfo()

}

func (b *Bread) shoppingList(ingredientsHave map[string]Item) (map[string]Item, map[string]Item){
	var ingredientsNeed = make(map[string]Item)
	var ingredientsLeft = make(map[string]Item)
	
	for key := range b.ingredients {
		if val, ok := ingredientsHave[key]; ok{
			if val.weight - b.ingredients[key].weight < 0 {
				ingredientsNeed[key] = Item{b.ingredients[key].weight - val.weight}
			}
		} else {
			ingredientsNeed[key] = b.ingredients[key]
		}
	}

	//Find the ingredients we will have left after we make the bread
	for key := range b.ingredients {
	//**TODO: Check if the required ingredient is in the pantry
		ingredientsLeft[key] = Item{ingredientsHave[key].weight -  b.ingredients[key].weight}
	}
	
	
	return ingredientsNeed, ingredientsLeft
}

func (b *Bread) printBakeInstructions(){
	fmt.Println("Bake at " + strconv.Itoa(b.baking.temperature) + " Celsius for " + strconv.Itoa(b.baking.bakeTime) + " minutes and let cool for " + strconv.Itoa(b.baking.coolTime) + " minutes")
}

func FloatToString(input_num float32) string {
    // to convert a float number to a string
    return strconv.FormatFloat(float64(input_num), 'f', 3, 32)
}

func (b *Bread) printBreadInfo(){
	fmt.Println(b.name + ":")
	
	ingredientList := "map["
	
	for key, value := range b.ingredients {
		ingredientList += key + ":{" + strconv.Itoa(value.weight) + "}" + " "
	}
	
	ingredientList += "]\n" + "Weight: " + FloatToString(b.weight/1000) + " kg"
	fmt.Println(ingredientList)
}

func NewBread()(*Bread){
	var bread Bread
	b := &bread

	b.name = "Whole Wheat Bread"
	b.ingredients = make(map[string]Item)
	
	b.ingredients["whole wheat flour"] = Item{500}
	b.ingredients["yeast"] = Item{25}
	b.ingredients["salt"] = Item{25}
	b.ingredients["sugar"] = Item{50}
	b.ingredients["butter"] = Item{50}
	b.ingredients["water"] = Item{350}
	b.baking.bakeTime = 120
	b.baking.temperature = 180
	b.baking.coolTime = 60
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
				b.ingredients[k] = val
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
				b.ingredients[k] = val
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
	//Create a standard whole wheat bread
	 wholeWheat := NewBread()
	 
	 //Initialize the ingredients for sesame bread
	 var sesameIngredients = make(map[string]Item)
	 sesameIngredients["sesame"] = Item{50}
	 sesameIngredients["white flour"] = Item{250}
	 
	 var sesameIngredientsLess = make(map[string]Item)
	 sesameIngredientsLess["whole wheat flour"] = Item{250}
	 
	 //Create a sesame bread
	 sesame := NewBreadVariation("Sesame bread", sesameIngredients, sesameIngredientsLess)
	
	//List of ingredients that we have currently
	var ingredientsHave = make(map[string]Item)
	ingredientsHave["whole wheat flour"] = Item{5000}
	ingredientsHave["salt"] = Item{500}
	ingredientsHave["sugar"] = Item{1000}
	
	//Get shopping list for whole wheat bread
	ingredientsNeed, _ := wholeWheat.shoppingList(ingredientsHave) 
	
	//Get shopping list for sesame bread
	ingredientsNeed2, _ := sesame.shoppingList(ingredientsHave)
	
	//Combine shopping lists for both whole wheat and sesame breads
	for k, v := range ingredientsNeed2 {
		if val, ok := ingredientsNeed[k]; ok{
			val.weight += ingredientsNeed2[k].weight
			ingredientsNeed[k] = val
		} else {
		ingredientsNeed[k] = v
		}
	}
		
	shopList := "map["
	
	//Build the shopping list string
	for key, value := range ingredientsNeed {
		shopList += key + ":{" + strconv.Itoa(value.weight) + "}" + " "
	}
	shopList += "]"
	
	//Print the breadInfo, shopping list, and bake instructions
	wholeWheat.printBreadInfo()
	fmt.Println()
	sesame.printBreadInfo()
	fmt.Println()
	
	fmt.Println("Shopping List:")
	fmt.Println(shopList)
	fmt.Println()
	
	fmt.Println("Baking Instructions:")
	wholeWheat.printBakeInstructions()
	sesame.printBakeInstructions()

}
 
