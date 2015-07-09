package types

import (
	units "github.com/davelondon/ke_common/units"
	_ "github.com/davelondon/ke_common/units/types"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728866688 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "image", Exists: true}, Description: "This interface type represents all images - icon and photo", Rules: []system.Rule(nil)}
	ptr8728866816 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}, Description: "Restriction rules for images", Rules: []system.Rule(nil)}
	ptr8728866944 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728853344 := &system.RuleBase{Optional: true}
	ptr8728853312 := &system.Bool_rule{Base: ptr8728866944, RuleBase: ptr8728853344, Default: system.Bool{}}
	ptr8728937008 := &system.Type{Base: ptr8728866816, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"secure": ptr8728853312}, Rule: (*system.Type)(nil)}
	ptr8728936896 := &system.Type{Base: ptr8728866688, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: ptr8728937008}
	ptr8728867584 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}, Description: "This represents an image, and contains path, server and protocol separately", Rules: []system.Rule(nil)}
	ptr8728867712 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The path for the url - e.g. /foo/bar.jpg", Rules: []system.Rule(nil)}
	ptr8728676416 := &system.String_rule{Base: ptr8728867712, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 1, Exists: true}, Equal: system.String{}, Pattern: system.String{Value: "^/.*$", Exists: true}, Format: system.String{}}
	ptr8728864896 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The protocol for the url - e.g. http or https", Rules: []system.Rule(nil)}
	ptr8729231776 := &system.RuleBase{Optional: true}
	ptr8728674304 := &system.String_rule{Base: ptr8728864896, RuleBase: ptr8729231776, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728865152 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "The server for the url - e.g. www.google.com", Rules: []system.Rule(nil)}
	ptr8728675712 := &system.String_rule{Base: ptr8728865152, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728865280 := &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "@rectangle", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728512416 := &units.Rectangle_rule{Base: ptr8728865280, RuleBase: (*system.RuleBase)(nil)}
	ptr8728937456 := &system.Type{Base: ptr8728867584, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"path": ptr8728676416, "protocol": ptr8728674304, "server": ptr8728675712, "size": ptr8728512416}, Rule: (*system.Type)(nil)}
	ptr8728866304 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@photo", Exists: true}, Description: "Automatically created basic rule for photo", Rules: []system.Rule(nil)}
	ptr8728937120 := &system.Type{Base: ptr8728866304, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728865920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "icon", Exists: true}, Description: "This is a type of image, which just contains the url of the image", Rules: []system.Rule(nil)}
	ptr8728866048 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728674480 := &system.String_rule{Base: ptr8728866048, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728936560 := &system.Type{Base: ptr8728865920, Embed: []system.Reference(nil), Is: []system.Reference{system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8728674480}, Rule: (*system.Type)(nil)}
	ptr8728866176 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@icon", Exists: true}, Description: "Automatically created basic rule for icon", Rules: []system.Rule(nil)}
	ptr8728936672 := &system.Type{Base: ptr8728866176, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_common/images", "image", ptr8728936896)
	system.RegisterType("github.com/davelondon/ke_common/images", "photo", ptr8728937456)
	system.RegisterType("github.com/davelondon/ke_common/images", "@photo", ptr8728937120)
	system.RegisterType("github.com/davelondon/ke_common/images", "icon", ptr8728936560)
	system.RegisterType("github.com/davelondon/ke_common/images", "@image", ptr8728937008)
	system.RegisterType("github.com/davelondon/ke_common/images", "@icon", ptr8728936672)
}
