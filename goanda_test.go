package goanda

import "testing"

func TestGoandaTest(t *testing.T) {

	testString := GoandaTest()

	if testString != "Test" {
		t.Errorf("GoandaTest didn't return 'Test'")
	}

}

func TestGetAccounts(t *testing.T) {

	GetAccounts()
	t.Log("accounts test")

}

func TestConfig(t *testing.T) {
	t.Log(TEST_API_KEY)
}
