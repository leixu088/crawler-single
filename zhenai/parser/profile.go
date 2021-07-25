package parser

import (
	"crawler/engine"
	"crawler/model"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)


//解析json数据
func parseJson(json []byte) model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()
	//infos是一个切片，里面的类型是interface{}

	//fmt.Printf("infos:%v,  %T\n", infos, infos) //infos:[离异 47岁 射手座(11.22-12.21) 157cm 55kg 工作地:阿坝汶川 月收入:3-5千 教育/科研 大学本科],  []interface {}

	var profile model.Profile
	//所以我们遍历这个切片，里面使用断言来判断类型
	for k, v := range infos {
		//fmt.Printf("k:%v,%T\n", k, k)
		//fmt.Printf("v:%v,%T\n", v, v)

		/*
		    "basicInfo":[
		       "未婚",
		       "25岁",
		       "魔羯座(12.22-01.19)",
		       "152cm",
		       "42kg",
		       "工作地:阿坝茂县",
		       "月收入:3-5千",
		       "医生",
		       "大专"
		   ],
		*/
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				//年龄:47岁，我们可以设置int类型，所以可以通过另一个json字段来获取
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}

	}
	return profile
}

func ParseProfile(contents []byte, name string) engine.ParseResult {
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		profile := parseJson(json)
		profile.Name = name
		result.Items = append(result.Items, profile)
	}
	// profile : {坚果 离异 37岁  178cm 71kg 月收入:1.2-2万 大专 自由职业  魔羯座(12.22-01.19)  }
	return result
}
