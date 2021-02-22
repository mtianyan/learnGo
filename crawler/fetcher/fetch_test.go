package fetcher

import (
	"testing"
)

func TestFetcher(t *testing.T) {
	url := "https://album.zhenai.com/u/1428069052"
	content, err := Fetch(url)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%s\n", content)
	//req, err := http.Get(url)
	//if err != nil {
	//	t.Log(req, err)
	//}
	//t.Log(req)
	//
	//t.Log(req.Request.Header)

	//fmt.Println(getUrlBase("http://www.zhenai.com/zhenghun/henan"))
	//json.Marshal()
}
