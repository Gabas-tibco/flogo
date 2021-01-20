package java

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type HashCode struct {
}

func init() {
	function.Register(&HashCode{})
}

func (hc *HashCode) Name() string {
	return "hashCode"
}

func (hc *HashCode) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

func (hc *HashCode) Eval(params ...interface{}) (interface{}, error) {
	item := params[0]

	hash := hc.hashCode([]byte(hc.ToString(item)))

	return int(hash), nil
}

func (hc *HashCode) hashCode(bytes []byte) uint32 {
	var hash uint32
	for _, b := range bytes {
		hash = 31*hash + uint32(b)
	}

	return hash
}

func (hc *HashCode) ToString(something interface{}) string {
	return fmt.Sprintf("%v", something)
}