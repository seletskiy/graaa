package main

import (
	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
)

func start() {

}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
