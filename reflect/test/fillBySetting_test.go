package Reflect

import (
	fillbysetting "common/reflect/fillBySetting"
	"testing"
)

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}
type Lee struct {
	Name string
}

func TestFillBySetting(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employee{}
	if err := fillbysetting.FillBySetting(&e, settings); err != nil {
		t.Error(err)
	}
	t.Log(e)
	setting := map[string]interface{}{"Name": "lee"}
	a := &Lee{}
	if err := fillbysetting.FillBySetting(a, setting); err != nil {
		t.Error(err)
	}
	t.Log(a)

	d := 1
	b := &d
	if err := fillbysetting.FillBySetting(b, setting); err != nil {
		t.Error(err)
	}
	t.Log(b)
	// c := new(Customer)
	// if err := fillBySetting(c, settings); err != nil {
	// 	t.Error(err)
	// }
	// t.Log(*c)
}
