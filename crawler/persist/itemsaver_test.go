package persist

import (
	"001_go_env/crawler/engine"
	"001_go_env/crawler/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {

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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = save(client, index, expected)


	if err != nil{
		panic(err)
	}


	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil{
		panic(err)
	}

	//t.Logf("%+v", resp)

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil{
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile
	// Verify result
	if actual != expected{
		t.Errorf("got %v; expected %v", actual, expected)
	}


}