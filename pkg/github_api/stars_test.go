package github_api

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetStarts(t *testing.T) {
	starred, err := GetStarts("fs714")
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.FailNow()
	}

	starredJson, err := json.MarshalIndent(starred, "", "  ")
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.FailNow()
	}

	fmt.Println(string(starredJson))
}
