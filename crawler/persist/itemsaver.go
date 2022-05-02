package persist

import (
	"001_go_env/crawler/engine"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error){

	client, err := elastic.NewClient(
		// Must turn off sniff in the docker
		elastic.SetSniff(false),
	)

	if err != nil{
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for{
			item := <-out
			log.Printf("Item Saver: got item " + "#%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("item saver: got item " +
					"#%v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// Response id String, err error
func save(client *elastic.Client, index string, item engine.Item) (err error){


	if item.Type == ""{
		return errors.New("must supply Type")
	}

	indexService:= client.Index().
		Index(index).
		Type(item.Type). // table name from parse
		BodyJson(item)

	if item.Id != ""{
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())

	//resp, err := client.Index().
	//	Index("dating_profile").
	//	Type(item.Type). // table name from parse
	//	Id(item.Id). // from parse
	//	BodyJson(item).
	//	Do(context.Background())

	if err != nil {
		return err
	}

	//fmt.Printf("%+v", resp)

	return nil

}