package fetcher

import (
	"log"
	"testing"
)


func TestFetch(t *testing.T) {
	contents, err := Fetch("https://album.zhenai.com/u/1999489298")
	if err != nil{
		t.Errorf("存在错误: %v", err)
	}
	log.Printf("%s", contents)
}
