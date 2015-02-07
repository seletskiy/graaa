package main

type Items []*Item

type Item struct {
	Content string
}

func NewItems() *Items {
	return &Items{}
}

func NewItem() *Item {
	return &Item{}
}

func showItems(items *Items) {

}
