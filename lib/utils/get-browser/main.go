// Package main ...
package main

import (
	"fmt"

	"github.com/PrismAIO/rod/lib/launcher"
	"github.com/PrismAIO/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
