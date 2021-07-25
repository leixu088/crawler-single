package parser

import (
	"crawler/fetcher"
	"log"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch("https://album.zhenai.com/u/1999489298")
	log.Printf("%s", contents)
	if err != nil{
		panic(err)
	}
	result := ParseProfile(contents, "坚果")


	const resultSize = 1
	if len(result.Items) != resultSize{
		t.Errorf("result should have %d request, but had %d", resultSize, len(result.Items))
	}
	//for
	//xx := result.Items[0]
	//log.Printf("%#v", xx)
}
