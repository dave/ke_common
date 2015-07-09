package types

import (
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728841856 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "@localizer", Exists: true}, Description: "Automatically created basic rule for localizer", Rules: []system.Rule(nil)}
	ptr8728523808 := &system.Type{Base: ptr8728841856, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728842112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "simple", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728841984 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728660384 := &system.String_rule{Base: ptr8728841984, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728524144 := &system.Type{Base: ptr8728842112, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "localizer", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"string": ptr8728660384}, Rule: (*system.Type)(nil)}
	ptr8728841728 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "localizer", Exists: true}, Rules: []system.Rule(nil)}
	ptr8728523584 := &system.Type{Base: ptr8728841728, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728842240 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "@simple", Exists: true}, Description: "Automatically created basic rule for simple", Rules: []system.Rule(nil)}
	ptr8728524368 := &system.Type{Base: ptr8728842240, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728840704 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "translation", Exists: true}, Description: "This represents a translated string", Rules: []system.Rule(nil)}
	ptr8728840832 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The original English string", Rules: []system.Rule(nil)}
	ptr8728817664 := &system.RuleBase{Optional: true}
	ptr8728658096 := &system.String_rule{Base: ptr8728840832, RuleBase: ptr8728817664, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728840960 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "The translated strings", Rules: []system.Rule(nil)}
	ptr8728841088 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728659328 := &system.String_rule{Base: ptr8728841088, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728683840 := &system.Map_rule{Base: ptr8728840960, RuleBase: (*system.RuleBase)(nil), Items: ptr8728659328, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728518768 := &system.Type{Base: ptr8728840704, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "localizer", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"english": ptr8728658096, "translations": ptr8728683840}, Rule: (*system.Type)(nil)}
	ptr8728841216 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "@translation", Exists: true}, Description: "Automatically created basic rule for translation", Rules: []system.Rule(nil)}
	ptr8728522576 := &system.Type{Base: ptr8728841216, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_common/words", "localizer", ptr8728523584)
	system.RegisterType("github.com/davelondon/ke_common/words", "@simple", ptr8728524368)
	system.RegisterType("github.com/davelondon/ke_common/words", "translation", ptr8728518768)
	system.RegisterType("github.com/davelondon/ke_common/words", "@translation", ptr8728522576)
	system.RegisterType("github.com/davelondon/ke_common/words", "@localizer", ptr8728523808)
	system.RegisterType("github.com/davelondon/ke_common/words", "simple", ptr8728524144)
}
