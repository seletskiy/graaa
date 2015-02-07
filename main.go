package main

import (
	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
)

func start() {
	items := NewItems()
	item := NewItem()
	items = append(items, item)
	showItems(items)

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
