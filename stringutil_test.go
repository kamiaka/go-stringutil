package stringutil

import "testing"

func TestInsert(t *testing.T) {

	sampleParam := struct {
		Name string
	}{
		Name: "World",
	}

	actual, _ := Insert("Hello {{.Name}}!", sampleParam)
	expected := "Hello World!"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
