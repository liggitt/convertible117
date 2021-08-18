package convertible117

import (
	"reflect"
	"testing"

	a "github.com/liggitt/convertible117/a"
	b "github.com/liggitt/convertible117/b"
)

func TestConvertible(t *testing.T) {
	actual := []a.Reference{}
	expected := []b.Reference{}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		t.Fatal("nil")
	}
	expectedValue := reflect.ValueOf(expected)
	if !expectedValue.IsValid() {
		t.Fatal("!valid")
	}
	if !expectedValue.Type().ConvertibleTo(actualType) {
		t.Fatal("!convertible")
	}
	if !reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual) {
		t.Fatal("!equal")
	}
}
