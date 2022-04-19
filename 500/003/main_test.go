package main

import (
	"ddd/domain"
	"ddd/domain/item"
	"fmt"
	"testing"
)

func TestItem(t *testing.T) {
	itermInit()
	var it item.Item
	it.ID = "1"
	it.Name = "1"
	it.Description = "1"
	it.Price = 1
	domain.ItemsAdd(it)
	ins, err := domain.ItemsGet("1")
	fmt.Printf("ins:%+v\n", ins)
	fmt.Printf("err:%+v\n", err)
}
