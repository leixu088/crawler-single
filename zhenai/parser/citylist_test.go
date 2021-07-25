package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//log.Printf("%s",contents)
	result := ParseCityList(contents)

	const resultSize = 24
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/beijing", "http://www.zhenai.com/zhenghun/shanghai", "http://www.zhenai.com/zhenghun/guangzhou",
	}

	expectedCities := []string{
		"City 北京", "City 上海", "City 广州",
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
	for i, c := range expectedCities{
		if result.Items[i] != c {
			t.Errorf("url #%d: %s; but %s", i, c, result.Items[i])
		}
	}
}
