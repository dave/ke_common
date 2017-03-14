package main

import (
	_ "github.com/dave/ke_common/images"
	_ "github.com/dave/ke_common/images/types"
	_ "github.com/dave/ke_common/units"
	_ "github.com/dave/ke_common/units/types"
	"kego.io/editor/client"
	"kego.io/js/console"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func main() {
	if err := client.Start("github.com/dave/ke_common/images"); err != nil {
		console.Error(err.Error())
	}
}
