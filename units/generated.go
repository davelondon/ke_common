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
	Height system.Number
	Width  system.Number
}

func init() {
	json.RegisterType("github.com/davelondon/ke_common/units", "@rectangle", reflect.TypeOf(&Rectangle_rule{}))
	json.RegisterType("github.com/davelondon/ke_common/units", "rectangle", reflect.TypeOf(&Rectangle{}))
}
