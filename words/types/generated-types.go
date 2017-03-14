package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for localizer", Id: system.Reference{Package: "github.com/dave/ke_common/words", Name: "@localizer", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Automatically created basic rule for simple", Id: system.Reference{Package: "github.com/dave/ke_common/words", Name: "@simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Base{Description: "Automatically created basic rule for translation", Id: system.Reference{Package: "github.com/dave/ke_common/words", Name: "@translation", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Base: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Base{Id: system.Reference{Package: "github.com/dave/ke_common/words", Name: "localizer", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Base: ptr6, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Base{Id: system.Reference{Package: "github.com/dave/ke_common/words", Name: "simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr10 := &system.RuleBase{}
	ptr11 := &system.String_rule{Base: ptr9, RuleBase: ptr10, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr12 := &system.Type{Base: ptr8, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"string": ptr11}, Is: []system.Reference{system.Reference{Package: "github.com/dave/ke_common/words", Name: "localizer", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr13 := &system.Base{Description: "This represents a translated string", Id: system.Reference{Package: "github.com/dave/ke_common/words", Name: "translation", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr14 := &system.Base{Description: "The original English string", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr15 := &system.RuleBase{Optional: true}
	ptr16 := &system.String_rule{Base: ptr14, RuleBase: ptr15, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr17 := &system.Base{Description: "The translated strings", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr18 := &system.RuleBase{}
	ptr19 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr20 := &system.RuleBase{}
	ptr21 := &system.String_rule{Base: ptr19, RuleBase: ptr20, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr22 := &system.Map_rule{Base: ptr17, RuleBase: ptr18, Items: ptr21, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr23 := &system.Type{Base: ptr13, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"english": ptr16, "translations": ptr22}, Is: []system.Reference{system.Reference{Package: "github.com/dave/ke_common/words", Name: "localizer", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/dave/ke_common/words", "@localizer", ptr1, 0x8266937074db9a35)
	system.RegisterType("github.com/dave/ke_common/words", "@simple", ptr3, 0xdd35f294171a4fee)
	system.RegisterType("github.com/dave/ke_common/words", "@translation", ptr5, 0xf341b6b3891dc7c9)
	system.RegisterType("github.com/dave/ke_common/words", "localizer", ptr7, 0x8266937074db9a35)
	system.RegisterType("github.com/dave/ke_common/words", "simple", ptr12, 0xdd35f294171a4fee)
	system.RegisterType("github.com/dave/ke_common/words", "translation", ptr23, 0xf341b6b3891dc7c9)
}
