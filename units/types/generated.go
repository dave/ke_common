package types

import (
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728832512 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728832128 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660704 := &system.Number_rule{Base: ptr8728832128, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728832384 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660800 := &system.Number_rule{Base: ptr8728832384, RuleBase: (*system.RuleBase)(nil), Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728514384 := &system.Type{Base: ptr8728832512, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"height": ptr8728660704, "width": ptr8728660800}, Rule: (*system.Type)(nil)}
	ptr8728832640 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "@rectangle", Exists: true}, Description: "Automatically created basic rule for rectangle", Rules: []system.Rule(nil)}
	ptr8728514608 := &system.Type{Base: ptr8728832640, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_common/units", "rectangle", ptr8728514384)
	system.RegisterType("github.com/davelondon/ke_common/units", "@rectangle", ptr8728514608)
}
