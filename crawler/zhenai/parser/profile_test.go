package parser

import (
	"001_go_env/crawler/engine"
	"001_go_env/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil{
		panic(err)
	}

	result := ParseProfile(contents,"http://album.zhenai.com/u/1998372165",
		"Lily")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element " +
			"but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url: "http://album.zhenai.com/u/1998372165",
		Type: "zhenai",
		Id: "1998372165",
		Payload: 					model.Profile{
			Name: "Lily",
			Age: 27,
			Height: 157,
			Weight: 51,
			Income: "3001-5000",
			Marriage: "未婚",
			Hokou: "阿坝",
		},
	}

	if actual != expected{
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}

}
