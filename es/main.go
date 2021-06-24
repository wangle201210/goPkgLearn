package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"
)

const (
	index = "per"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"properties":{
			"id":{
				"type":"long"
			},
			"name":{
				"type":"text"
			},
			"age":{
				"type":"long"
			},
			"created":{
				"type":"date"
			},
			"cellphone":{
				"type":"text"
			}
		}
	}
}
`

type Doctor struct {
	Id        int64
	Name      string
	Age       int64
	Cellphone string
	CreatedAt int64
}

var Client *elastic.Client

func main() {
	//for i := 0; i < 10; i++ {
	//	id := rand.Int63n(1000)
	//	println("=============== ", id)
	//}
	//return
	newClient()
	//createIndexIfNotExit()
	//insertData()
	updateData()
	deleteData()
	searchData()
}

func deleteData()  {
	_, err := Client.Delete().Index(index).Id("1").Do(context.TODO())
	if err != nil{
		panic(err)
	}
}

func updateData()  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		id := rand.Int63n(10)
		updateResponse, err := Client.UpdateByQuery().Index(index).
			Query(elastic.NewBoolQuery().Must(elastic.NewTermsQuery("Id", id))).
			Script(elastic.NewScript("ctx._source.Name = params.Name").Params(map[string]interface{}{"Name": "lalalal"})).
			Do(context.TODO())
		if err != nil {
			panic(err)
		}
		log.Printf("update (%d) updateResponse is (%+v)\n", id, updateResponse)
	}
}
func searchData()  {
	result, err := Client.Search().Index(index).From(0).Size(10).
		//Query(elastic.NewRangeQuery("Age").From(3).To(99)).
		//Query(elastic.NewMoreLikeThisQuery().Field("Name").LikeText("wanna1")).
		//Query(elastic.NewBoolQuery().Must(elastic.NewMatchQuery("Name", "wanna1"))).
		Query(elastic.NewPrefixQuery("Name", "wanna")).
		//Query(elastic.NewPrefixQuery("Cellphone", "1822759999")).
		Do(context.TODO())
	if err != nil {
		log.Printf("BodyJson err is (%+v)\n", err)
	}
	log.Println("searchData start echo")
	for i, d := range result.Each(reflect.TypeOf(Doctor{})) {
		data := d.(Doctor)
		log.Printf("get %d data %+v", i + 1, data)
	}
	log.Printf("searchData success total (%d)\n", result.Hits.TotalHits.Value)
}

func insertData() {
	for i := 1; i <= 10; i++ {
		var data = Doctor{
			Id:        int64(i),
			Name:      "wanna" + fmt.Sprintf("%d", i) +
				//"The quick brown fox jumps over the quick dog" +
				"",
			Age:       rand.Int63n(100),
			Cellphone: "1822759999" + fmt.Sprintf("%d", i),
			CreatedAt: time.Now().Unix(),
		}
		res, err := Client.Index().
			Index(index).Id(fmt.Sprintf("%d", data.Id)).BodyJson(data).Do(context.TODO())
		if err != nil {
			log.Printf("BodyJson err is (%+v)\n", err)
		}
		log.Printf("response is (%+v)\n", res)
	}

}

func createIndexIfNotExit() {
	if exists, _ := Client.IndexExists(index).Do(context.TODO()); exists {
		log.Printf("%s is exists\n", index)
		return
	}
	createIndex, err := Client.CreateIndex(index).BodyString(mapping).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged ,创建失败
		log.Println("Not acknowledged")
	}
	return
}

func newClient() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		panic(err)
	}
	Client = client
}