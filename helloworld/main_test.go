package main

import "testing"

func Test_hello(t *testing.T)  {
	expected := "hello"

	if actual := getHello(); actual != expected {
		t.Errorf("Expect '%s' but got '%s'", expected, actual)
	}
	
}