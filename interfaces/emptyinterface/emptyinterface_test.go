package emptyinterface

import (
	"fmt"
	"testing"
	"time"
)

// go test -run TestCasting
func TestCasting(t *testing.T) {
	// Pass an empty interface and cast to non-pointer
	change := func(item interface{}, newValue string) {
		o := item.(Object)
		o.Value = newValue
	}

	o := Object{}
	change(o, "newval")
	// this won't work
	// change(&o, "newval")

	fmt.Println(o)
}

// go test -run TestPointerCasting
func TestPointerCasting(t *testing.T) {
	// Pass an empty interface and cast to pointer
	change := func(item interface{}, newValue string) {
		o := item.(*Object)
		o.Value = newValue
	}

	o := Object{}
	change(&o, "newval")
	// this won't work
	// change(o, "newval")

	fmt.Println(o)
}

// go test -run TestMapFlattening
func TestMapFlattening(t *testing.T) {
	jsonObj := map[string]interface{}{
		"sub1": map[string]interface{}{
			// Must be "map[string]interface{}" instead of "map[string]int"
			// otherwise "case map[string]interface{}" will fail
			"sub11": map[string]interface{}{
				"sub111": 111,
				"sub112": 112,
			},
			"sub12": map[string]interface{}{
				"sub121": []interface{}{"121_1", "121_2"},
			},
		},
		"sub2": map[string]interface{}{
			"sub21": time.Now(),
			"sub22": []interface{}{
				map[string]interface{}{
					"sub221": 10,
					"sub222": "222",
				},
				"223",
				nil,
			},
		},
	}

	fmt.Println(jsonObj)
	flattened := Flatten(jsonObj, ".")
	fmt.Println(flattened)
}
