package parser

import (
	"crawler/fetcher"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun/beijing")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents)

	const resultSize = 20
	expectedUrls := []string{
		"http://album.zhenai.com/u/1313056057", "http://album.zhenai.com/u/1175302111", "http://album.zhenai.com/u/1029297350",
	}

	expectedUsers := []string{
		"User 上古龙裔", "User 粥粥", "User 鸽子",
	}

	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d request, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls{
		if result.Requests[i].Url != url {
			t.Errorf("url #%d: %s; but %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize{
		t.Errorf("result should have %d request, but had %d", resultSize, len(result.Items))
	}
	for i, c := range expectedUsers{
		if result.Items[i] != c {
			t.Errorf("user #%d: %s; but %s", i, c, result.Items[i])
		}
	}
}
