package goanda

import (
	"fmt"
	"testing"
)

func TestGoandaTest(t *testing.T) {

	testString := GoandaTest()

	if testString != "Test" {
		t.Errorf("GoandaTest didn't return 'Test'")
	}

}

func TestGetAccounts(t *testing.T) {

	accts, err := GetAccounts()
	if err != nil {
		t.Errorf("TestGetAccounts error: ")
		fmt.Println(err)
	}
	fmt.Println(accts)

}

func TestGetAccount(t *testing.T) {

	accts, err := GetAccounts()
	if err != nil {
		t.Errorf("TestGetAccount error: ")
		fmt.Println(err)
	}

	fmt.Println(accts.Accounts[0])

	acct, err := GetAccount(accts.Accounts[0].ID)

	if err != nil {
		t.Errorf("TestGetAccount error: ")
		fmt.Println(err)
	}
	fmt.Println(acct)

}

func TestConfig(t *testing.T) {
	t.Log(TEST_API_KEY)
}
