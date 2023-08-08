package goanda

import "testing"

func TestGoandaTest(t *testing.T) {

	// got := Add(4, 6)
	// want := 10

	// if got != want {
	//     t.Errorf("got %q, wanted %q", got, want)
	// }

	testString := GoandaTest()

	if testString != "Test" {
		t.Errorf("GoandaTest didn't return 'Test'")
	}

}
