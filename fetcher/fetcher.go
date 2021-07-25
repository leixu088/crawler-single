package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte, error) {

	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	cookie1 := "sid=b3eca7c3-634c-4906-9a04-d2b365eab0af; ec=Ak9Ys2dB-1625316317044-9a9c482ff40a11294748650; FSSBBIl1UgzbN7NO=5ms7j5n0N.MfUKEPeypbOGmUy57_pgKzyoTV5Q8lg0YePISgLcwhk_eabUruKHNUNFpBg.eITqhMUXyXs8G0Xaa; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1625316326,1627116568; _exid=MS%2FEH3F3ED9BOKR2zhOLjlfGVDCWDM%2FsIUOjJoI7XvdlhpCUk0w%2FT9Et8Z0PkzUlzAtb1wO78saT2ZcPrUwLPw%3D%3D; _efmdata=paTQSVpqadWdRXFLXV3pCO7O0ntOTxhDLGnUEJPX7FzXv54CZ3kwnmlTupgUR1kedAgjb55RQB5N9xdur9F2evnqZAhabm64W7lfVoGnggM%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1627210166; FSSBBIl1UgzbN7NP=53HTcYKkeA2lqqqm_6V78Aa1JUa553hpsfiRDdrB7kFhlMvsgr6wRRHxunBtUFfevv7QuDigvTEY5khdkzUAvWcDNp_Gg_Vk1MaCmOGNqjTRm8zYmdoGYG6enPMpLr4QDuh3i0Qq7.XPwY2pq1aVyBZho4225KOeoDGkOmSaEuwmXHSqvwtcBjwuC3k0HVnS22IRvuZqBZsDtEyrT2HQXZ9pM9xDM7gPlbHZ9RcPWV9xGNoW50jrWEpG_JqAIRaRHV"
	req.Header.Add("cookie", cookie1)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// 把网页转为utf-8编码
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
