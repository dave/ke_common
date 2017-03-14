package main

import (
	_ "github.com/dave/ke_common/words"
	_ "github.com/dave/ke_common/words/types"
	"kego.io/editor/client"
	"kego.io/js/console"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := client.Start("github.com/dave/ke_common/words"); err != nil {
		console.Error(err.Error())
	}
}
