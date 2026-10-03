package main

import "testing"

func TestSayHello(t *testing.T) {
	d := Dog{}
	c := Cat{}

	if got := SayHello(d); got != "ワンワン" {
		t.Errorf("SayHello(d) = %q; want \"ワンワン\"", got)
	}

	if got := SayHello(c); got != "ニャー" {
		t.Errorf("SayHello(c) = %q; want \"ニャー\"", got)
	}
}
