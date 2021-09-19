package jsonparse

import "testing"

func Test_Parse(t *testing.T) {
	var test = `
		{
			"a": 32,
			"b": 55,
			"c": [{
					"f": 23
			}]
		}
	`

	j, err := New([]byte(test))
	if err != nil {
		t.Fatal(err)
	}

	a, err := j.Object().Get("a")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)

	b, err := j.Object().Get("b")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)

	cf, err := j.Object().Select("c").Array().Select(0).Object().Get("f")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cf)
}
