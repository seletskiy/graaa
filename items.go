package main

import (
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

var (
	viewId int = 14
)

type Items []*Item
type Item struct {
	Content string
}

func NewItems() *Items {
	return &Items{}
}

func NewItem(content string) *Item {
	return &Item{Content: content}
}

func (item Item) Display() {
	viewId := viewId + 1

	textView := android.CreateView(
		strconv.Itoa(viewId), "android.widget.TextView").(sdk.TextView)
	textView.SetText1s(item.Content)

	android.AttachView(textView, mainLayoutId)
}

func (items Items) Display() {
	for _, item := range items {
		item.Display()
	}
}
