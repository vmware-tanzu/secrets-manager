package data

import (
	"reflect"
	"testing"
)

func TestConvertValueNoTemplate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		values := []string{"{\"key\":\"value\"}"}
		expected := map[string][]byte{"key": []byte("value")}
		actual := convertValueNoTemplate(values)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		values := []string{"key:value"}
		expected := map[string][]byte{"VALUE": []byte("key:value")}
		actual := convertValueNoTemplate(values)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	})
}
