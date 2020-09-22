package cos

import (
	"fmt"
	"reflect"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type UniquizeByProperty struct {
}

func init() {
	function.Register(&UniquizeByProperty{})
}

func (s *UniquizeByProperty) Name() string {
	return "uniquizeByProperty"
}

func (s *UniquizeByProperty) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

func (s *UniquizeByProperty) Eval(params ...interface{}) (interface{}, error) {
	items := s.InterfaceToArray(params[0])
	prop := params[1].(string)
	keys := map[string]int {}
	var out []interface{}
	var value string
	for _, item := range items {
		value = item.(map[string]interface{})[prop].(string)
		_, exist := keys[value]
		if !exist {
			keys[value] = 1
			out = append(out, item)
		}
	}

	return out, nil
}

func (u *UniquizeByProperty) InterfaceToArray(something interface{}) []interface{} {
	var items []interface{}
	value := reflect.ValueOf(something)
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i).Interface()
		items = append(items, item)
	}

	return items
}

func (u *UniquizeByProperty) ToString(something interface{}) string {
	return fmt.Sprintf("%v", something)
}
