package goanda

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// func TestGoandaTest(t *testing.T) {

// 	testString := GoandaTest()

// 	if testString != "Test" {
// 		t.Errorf("GoandaTest didn't return 'Test'")
// 	}

// }

func TestGetAccounts(t *testing.T) {

	accts, err := GetAccounts()
	if err != nil {
		t.Errorf("TestGetAccounts error: ")
		fmt.Println(err)
	}
	spew.Dump(accts)

}

func TestGetAccount(t *testing.T) {
	// t.FailNow()

	accts, err := GetAccounts()
	if err != nil {
		t.Errorf("TestGetAccount error: ")
		fmt.Println(err)
	}

	acct, err := GetAccount(accts.Accounts[0].ID)

	if err != nil {
		t.Errorf("TestGetAccount error: ")
		fmt.Println(err)
	}

	if acct.CreatedByUserID <= 0 {
		t.Errorf("TestGetAccount didn't seem to get a valid account as 'created by' wasn't valid")
	}

	spew.Dump(acct)

}

func TestGetAcctInstruments(t *testing.T) {
	accts, err := GetAccounts()
	if err != nil {
		t.Errorf("TestGetAccount error: ")
		fmt.Println(err)
	}

	instruments, err := GetAcctInstruments(accts.Accounts[0].ID)
	if err != nil {
		t.Errorf("TestGetAccount error: ")
		fmt.Println(err)
	}

	if len(instruments) < 1 {
		t.Errorf("TestGetAcctInstruments failed due to instruments being null")
		t.FailNow()
	}

	if len(instruments[0].Name) < 1 {
		t.Errorf("TestGetAcctInstruments failed due to instrument name being null")
		t.FailNow()
	}

	spew.Dump(instruments[0])
}

func TestConfig(t *testing.T) {
	t.Log(TEST_API_KEY)
}
