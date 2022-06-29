package main

import (
	"fmt"

	"github.com/ajay-ajith/protobuf_tutorial/articlepb"
	"google.golang.org/protobuf/proto"
)

func main() {
	var article articlepb.Article = articlepb.Article{
		Id:      1,
		Title:   "A sample Article",
		Content: "A very simple article",
		Authors: []string{"Author_1", "Auhtor_2"},
		Type:    articlepb.ArticleType_news,
	}

	encoded_article, _ := proto.Marshal(&article)

	var decoded_article articlepb.Article
	proto.Unmarshal(encoded_article, &decoded_article)

	fmt.Println(encoded_article, decoded_article)
}
