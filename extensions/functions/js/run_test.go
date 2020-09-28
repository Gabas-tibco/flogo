package js

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var run = &Run{}

func TestRunNumber(t *testing.T) {
	script := "return 999+1;"
	result, _ := run.Eval(script)

	assert.True(t, toString(result) == "1000")
}

func TestRunWithParams(t *testing.T) {
	script := "return $flow.Number + $flow.Text;"
	result, _ := run.Eval(script, map[string]interface{}{"Number": 288, "Text": "HelloWorld", "other": "something"})

	assert.Equal(t, toString(result), "288HelloWorld")
}

func TestObject(t *testing.T) {
	script := `
		console.log("Number value before-> "+$flow.Number)
		$flow.Number = $flow.Number*2;
        console.log("Number value after-> "+$flow.Number)
		return $flow;
	`
	result, _ := run.Eval(script, map[string]interface{}{"Number": 42, "Text": "HelloWorld", "other": "something"} )
	objMap, _ := structToMap(result)

	assert.Equal(t, "84", toString(objMap["Number"]))
}

func toString(something interface{}) string {
	return fmt.Sprintf("%v", something)
}

func structToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}
