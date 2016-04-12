package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestPanicAndRecover(t *testing.T) {
	descr := fmt.Sprintf("panicAndRecover")
	god := panicAndRecover()
	want := errors.New("recoverd!")
	if !reflect.DeepEqual(god, want) {
		t.Errorf("%s", descr)
		t.Errorf("got")
		t.Errorf("%s", god)
		t.Errorf("expect")
		t.Errorf("%s", want)
	}
}
