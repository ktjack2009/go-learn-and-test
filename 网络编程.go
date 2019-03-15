// 简单的http协议

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/olivere/elastic"
)

type DSL struct {
	Query map[string]interface{} `json:"query"`
	From  int                    `json:"from"`
	Size  int                    `json:"size"`
}

func main() {
	// a := fmt.Sprintf("http://localhost:9200/%s/%s/%s", "news", "article", "1")
	// resp, _ := http.Get(a)
	// b, _ := ioutil.ReadAll(resp.Body)
	// var res map[string]interface{}
	// json.Unmarshal(b, &res)
	// fmt.Println(res)
	// source := res["_source"]
	// source_new, ok := source.(map[string]interface{})
	// fmt.Println(ok)

	a := fmt.Sprintf("http://localhost:9200/%s/%s/_search", "news", "article")
	dsl := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]string{"content": "测试"},
		},
		"from": 0,
		"size": 2,
	}
	data, _ := json.Marshal(dsl)
	reader := bytes.NewBuffer(data)
	request, _ := http.NewRequest("POST", a, reader)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	response, _ := client.Do(request)
	respBytes, _ := ioutil.ReadAll(response.Body)

	var searchResult elastic.SearchResult
	json.Unmarshal(respBytes, &searchResult)
	hits := searchResult.Hits.Hits
	var res []map[string]interface{}
	for _, v := range hits {
		article := make(map[string]interface{})
		json.Unmarshal(*v.Source, &article)
		article["score"] = *v.Score
		article["id"] = v.Id // 文章在elastic中的id
		switch article["table"] {
		case "newspapers_article":
			article["table"] = "报纸"
			article["refer"] = article["platform"]
		case "web_article":
			article["table"] = "网站"
		case "weixin_article":
			article["table"] = "微信"
		default:
		}
		res = append(res, article)
	}
}
