package words

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for localizer
type Localizer_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for simple
type Simple_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for translation
type Translation_rule struct {
	*system.Base
	*system.RuleBase
}
type Simple struct {
	*system.Base
	String system.String `json:"string"`
}

// This represents a translated string
type Translation struct {
	*system.Base
	// The original English string
	English system.String `kego:"{\"default\":{\"value\":\"http\"}}" json:"english"`
	// The translated strings
	Translations map[string]system.String `json:"translations"`
}

func init() {
	json.RegisterType("github.com/dave/ke_common/words", "@localizer", reflect.TypeOf(&Localizer_rule{}), 0x8266937074db9a35)
	json.RegisterType("github.com/dave/ke_common/words", "@simple", reflect.TypeOf(&Simple_rule{}), 0xdd35f294171a4fee)
	json.RegisterType("github.com/dave/ke_common/words", "@translation", reflect.TypeOf(&Translation_rule{}), 0xf341b6b3891dc7c9)
	json.RegisterType("github.com/dave/ke_common/words", "simple", reflect.TypeOf(&Simple{}), 0xdd35f294171a4fee)
	json.RegisterType("github.com/dave/ke_common/words", "translation", reflect.TypeOf(&Translation{}), 0xf341b6b3891dc7c9)
}
