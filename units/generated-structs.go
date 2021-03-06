package units

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for rectangle
type Rectangle_rule struct {
	*system.Base
	*system.RuleBase
}
type Rectangle struct {
	*system.Base
	Height system.Int `json:"height"`
	Width  system.Int `json:"width"`
}

func init() {
	json.RegisterType("github.com/dave/ke_common/units", "@rectangle", reflect.TypeOf(&Rectangle_rule{}), 0x1dec0b16d337111a)
	json.RegisterType("github.com/dave/ke_common/units", "rectangle", reflect.TypeOf(&Rectangle{}), 0x1dec0b16d337111a)
}
