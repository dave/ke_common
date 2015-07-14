package images

import (
	"reflect"

	"github.com/davelondon/ke_common/units"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for icon
type Icon_rule struct {
	*system.Base
	*system.RuleBase
}

// Restriction rules for images
type Image_rule struct {
	*system.Base
	*system.RuleBase
	Secure system.Bool
}

// Automatically created basic rule for photo
type Photo_rule struct {
	*system.Base
	*system.RuleBase
}

// This is a type of image, which just contains the url of the image
type Icon struct {
	*system.Base
	Url system.String
}

// This represents an image, and contains path, server and protocol separately
type Photo struct {
	*system.Base
	// The path for the url - e.g. /foo/bar.jpg
	Path system.String
	// The protocol for the url - e.g. http or https
	Protocol system.String `kego:"{\"default\":{\"value\":\"http\"}}"`
	// The server for the url - e.g. www.google.com
	Server system.String
	Size   *units.Rectangle
}

func init() {
	json.RegisterType("github.com/davelondon/ke_common/images", "@icon", reflect.TypeOf(&Icon_rule{}), 0x31466dcde3ac4844)
	json.RegisterType("github.com/davelondon/ke_common/images", "@image", reflect.TypeOf(&Image_rule{}), 0xa3f1010f55717184)
	json.RegisterType("github.com/davelondon/ke_common/images", "@photo", reflect.TypeOf(&Photo_rule{}), 0xf540a935655f2a95)
	json.RegisterType("github.com/davelondon/ke_common/images", "icon", reflect.TypeOf(&Icon{}), 0x31466dcde3ac4844)
	json.RegisterType("github.com/davelondon/ke_common/images", "photo", reflect.TypeOf(&Photo{}), 0xf540a935655f2a95)
}
