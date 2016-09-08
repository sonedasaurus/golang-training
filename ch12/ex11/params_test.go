package params

import (
	"testing"
)

func TestPack(t *testing.T) {
	s := struct {
		Name string `http:"name"`
		Age  int    `http:"age"`
	}{"Arugula", 35}
	u, err := Pack(&s)
	if err != nil {
		t.Errorf("Pack(%#v): %s", s, err)
	}
	want := "age=35&name=Arugula"
	got := u.RawQuery
	if got != want {
		t.Errorf("Pack(%#v): got %q, want %q", s, got, want)
	}
}
