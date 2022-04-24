package parser

import (
	"001_go_env/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil{
		panic(err)
	}

	result := ParseProfile(contents,"Lily")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element " +
			"but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name: "Lily",
		Age: 27,
		Height: 157,
		Weight: 51,
		Income: "3001-5000",
		Marriage: "未婚",
		Hokou: "阿坝",
	}

	if profile != expected{
		t.Errorf("expected %v; but was %v",
			expected, profile)
	}

}
