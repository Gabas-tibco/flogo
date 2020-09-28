package js

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/robertkrimen/otto"
)

type Run struct {
}

func init() {
	function.Register(&Run{})
}

func (s *Run) Name() string {
	return "run"
}

func (s *Run) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeAny}, true
}

func (s *Run) Eval(params ...interface{}) (interface{}, error) {
	vm := otto.New()

	if len(params) > 1 && params[1] != nil {
		vm.Set("$param", params[1])
	} else {
		vm.Set("$param", "{}")
	}

	baseScript := `
		function runScript() {
          console.log("Executing JS script");
          var $flow = JSON.parse(JSON.stringify($param));
          // console.log(JSON.stringify($flow));
		  %s
		} 
        var $result = runScript();
    `
	completeScript := fmt.Sprintf(baseScript, params[0].(string))
	fmt.Println(completeScript)

	_, err := vm.Run(completeScript)
	result, _ := vm.Get("$result")
	export, _ := result.Export()

	if err != nil {
		fmt.Println("JS error -> "+err.Error())
	}

	return export, err
}