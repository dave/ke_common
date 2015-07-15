package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for rectangle", Id: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "@rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Id: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr4 := &system.Int_rule{Base: ptr3, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr5 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr6 := &system.Int_rule{Base: ptr5, RuleBase: (*system.RuleBase)(nil), Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr7 := &system.Type{Base: ptr2, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"height": ptr4, "width": ptr6}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_common/units", "@rectangle", ptr1, 0x1dec0b16d337111a)
	system.RegisterType("github.com/davelondon/ke_common/units", "rectangle", ptr7, 0x1dec0b16d337111a)
}
