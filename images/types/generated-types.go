package types

import (
	"github.com/davelondon/ke_common/units"
	_ "github.com/davelondon/ke_common/units/types"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for icon", Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@icon", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Restriction rules for images", Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr4 := &system.RuleBase{Optional: true}
	ptr5 := &system.Bool_rule{Base: ptr3, RuleBase: ptr4, Default: system.Bool{}}
	ptr6 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"secure": ptr5}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr7 := &system.Base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr8 := &system.Type{Base: ptr7, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr9 := &system.Base{Description: "This is a type of image, which just contains the url of the image", Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "icon", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr10 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr11 := &system.String_rule{Base: ptr10, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr12 := &system.Type{Base: ptr9, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr11}, Is: []system.Reference{system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr13 := &system.Base{Description: "This interface type represents all images - icon and photo", Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr14 := &system.Type{Base: ptr13, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: ptr6}
	ptr15 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr16 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr17 := &system.String_rule{Base: ptr16, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr18 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr19 := &system.RuleBase{Optional: true}
	ptr20 := &system.String_rule{Base: ptr18, RuleBase: ptr19, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr21 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr22 := &system.String_rule{Base: ptr21, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "@rectangle", Exists: true}}
	ptr24 := &units.Rectangle_rule{Base: ptr23, RuleBase: (*system.RuleBase)(nil)}
	ptr25 := &system.Type{Base: ptr15, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"path": ptr17, "protocol": ptr20, "server": ptr22, "size": ptr24}, Is: []system.Reference{system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_common/images", "@icon", ptr1, 0x31466dcde3ac4844)
	system.RegisterType("github.com/davelondon/ke_common/images", "@image", ptr6, 0xa3f1010f55717184)
	system.RegisterType("github.com/davelondon/ke_common/images", "@photo", ptr8, 0xf540a935655f2a95)
	system.RegisterType("github.com/davelondon/ke_common/images", "icon", ptr12, 0x31466dcde3ac4844)
	system.RegisterType("github.com/davelondon/ke_common/images", "image", ptr14, 0xa3f1010f55717184)
	system.RegisterType("github.com/davelondon/ke_common/images", "photo", ptr25, 0xf540a935655f2a95)
}
