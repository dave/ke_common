package words

import (
	reflect "reflect"

	json "kego.io/json"
	system "kego.io/system"
)

// Automatically created basic rule for translation
type Translation_rule struct {
	*system.Base
	*system.RuleBase
}

// This represents a translated string
type Translation struct {
	*system.Base
	// The translated strings
	Translations map[string]system.String
	// The original English string
	English system.String `kego:"{\"default\":{\"value\":\"http\"}}"`
}
type Simple struct {
	*system.Base
	String system.String
}

// Automatically created basic rule for simple
type Simple_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for localizer
type Localizer_rule struct {
	*system.Base
	*system.RuleBase
}

func init() {
	json.RegisterType("github.com/davelondon/ke_common/words", "@translation", reflect.TypeOf(&Translation_rule{}))
	json.RegisterType("github.com/davelondon/ke_common/words", "translation", reflect.TypeOf(&Translation{}))
	json.RegisterType("github.com/davelondon/ke_common/words", "simple", reflect.TypeOf(&Simple{}))
	json.RegisterType("github.com/davelondon/ke_common/words", "@simple", reflect.TypeOf(&Simple_rule{}))
	json.RegisterType("github.com/davelondon/ke_common/words", "@localizer", reflect.TypeOf(&Localizer_rule{}))
}
