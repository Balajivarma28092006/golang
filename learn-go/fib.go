package main

import (
	"errors"
	"fmt"
)

type Vegetable struct {
	Name     string
	Quantity int
	price    float32
	discount float32
}

type List []Vegetable

func (veggies *List) Add(veg Vegetable) {
	*veggies = append(*veggies, veg)
}

func (veggies *List) Find(Name string) (*Vegetable, error) {
	for _, veg := range *veggies {
		if veg.Name == Name {
			return &veg, nil
		}
	}
	return nil, errors.New("Unable to find the veggie")
}

func Print(vegetable Vegetable) {
	fmt.Printf("Name: %s\n", vegetable.Name)
	fmt.Printf("Quantity: %d\n", vegetable.Quantity)
	fmt.Printf("Price: ₹%.2f\n", vegetable.price)
	fmt.Printf("Discount: %.2f%%\n", vegetable.discount)
}

func printAll(veggies List) {
	for _, veg := range veggies {
		Print(veg)
		fmt.Println("------------------")
	}
}

func main() {
	var myList List
	myList.Add(Vegetable{"Tomato", 10, 20.5, 5})
	myList.Add(Vegetable{"Carrot", 5, 15.0, 10})
	myList.Add(Vegetable{"Spinach", 2, 12.0, 0})

	veg, err := myList.Find("Carrot")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		Print(*veg)
	}

	printAll(myList)
}
