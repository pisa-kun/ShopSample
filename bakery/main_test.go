package main

import "testing"

// go testの対象
func TestProductShow(t *testing.T) {
	var sandwitch bread = bread{Name: "sandwitch", Value: 120, Number: 3}
	var list []bread = []bread{sandwitch}
	c := chashier{
		Product_list: list,
		Shop_name:    "hoge",
		Name:         "huge",
	}
	if err := c.productShow(); err != true {
		t.Fatal("c.ProductShow return false")
	}
}

func TestProductRecalculate(t *testing.T) {
	var sandwitch bread = bread{Name: "sandwitch", Value: 120, Number: 3}
	var list []bread = []bread{sandwitch}
	c := chashier{
		Product_list: list,
		Shop_name:    "hoge",
		Name:         "huge",
	}
	if err := c.productReCalculate(sandwitch.Name, 2); err != true {
		t.Fatal("error")
	}
}
