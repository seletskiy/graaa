package main

import (
	"github.com/seletskiy/go-android-rpc/android"
	"strconv"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
)

var mainLayoutId string

func start() {
	layoutResponse := android.GetLayoutById("main_layout")
	mainLayoutId = layoutResponse["layout_id"].(string)

	items := NewItems()
	for i := 0; i < 10; i++ {
		item := NewItem("aaa " + strconv.Itoa(i))
		*items = append(*items, item)
	}

	items.Display()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
