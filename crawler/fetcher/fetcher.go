package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(30 * time.Millisecond)

func generateCookies() (ret []*http.Cookie) {
	const cookieValues = "sid=89f9f5b8-fba1-4b9e-abd0-e79298128a24; ec=lhpvKfOg-1601368795488-cdd2fc0facd5a-115353789; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601368810; notificationPreAuthorizeSwitch=157728; __channelId=905821%2C0; _exid=bQ3XNFXNaDAYg67K4SJS5RLdJ%2BzG7q5BzsH8jJRCTlElxrk%2Bo%2Fhy9Ua1fHQRlE4yZx6IjpEA3j5P2fFg2D2Phg%3D%3D; _efmdata=ScTRp5N13XaQrrDgzGwiG9T2Ds8BhOi4AONN3fJO61C6S1wSw52FCYhdD8EGKWK6kS4S44nZ8ZElJYWMF4qyX59dbj0MJjk240RMKP%2BorcI%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601696107"

	pairs := strings.Split(cookieValues, "; ")
	for _, pair := range pairs {
		arr := strings.Split(pair, "=")
		cookie := &http.Cookie{
			Name:  arr[0],
			Value: arr[1],
		}
		ret = append(ret, cookie)
	}
	return ret
}

var cookies = generateCookies()

// 获取网页信息
// input: url
// output:
func Fetch(url string) ([]byte, error) {

	// 流量限制：执行到这里，需要隔10毫秒才继续往下执行
	<-rateLimiter
	//resp, err := http.Get(url)
	// 直接用http.Get(url)进行获取信息会报错：Error: status code 403
	client := &http.Client{}
	strings.Replace(url, "http://album.zhenai.com", "http://127.0.0.1:8000", 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// 查看自己浏览器中的User-Agent信息（检查元素->Network->User-Agent）
	req.Header.Add("authority", "album.zhenai.com")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("dnt", "1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("referer", "https://album.zhenai.com/u/1491731990")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("cookie", "sid=98196e18-d8f8-4969-8554-0eb8041930d1; FSSBBIl1UgzbN7NO=5r23wzGfoRmbLhnImzKHcjJC.4vmLdma2gQbjSoCxK2.1TrG09Gpkb_CQnG7BXV1k7MVAYugbUSmmYq1gLb05za; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621236454; ec=lFVERHrc-1621236455008-d58b6d5c17ba5-2057112252; _exid=%2Bb18ix%2FUYCoJ6IFb8CZHMk%2FTBkNjHowZshO5ZeNGexxKNGA216EdFXMqDZrd%2FXcsSrYiNy1jNXW98s%2B8spDyKQ%3D%3D; _efmdata=NHo%2Bq8CZ9cYBeisjQT5Lk4XIgxOCz4CMndxXVxENnUvmwiN8rk0c6ajQwjkFaYzPeTjZVPXwx%2BxgKbz2kDjUMh5NWTn%2Fs44gjCGTy2dYSr4%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621236919; FSSBBIl1UgzbN7NP=53tEQXbcPoh3qqqmy8Vq_Pq7HbCTSAN_Z_MMrp5CqVj.16zcqgewGivQJoVbCN3tzLwvTYDy72ygUJ7oNRW_mhSKY7l7nrt5XMvOt9v3goW.T7F8w2xyUgggvuMnWVO3JqfYLiIBgA.jT6LK._6Ltx7IKeeR03WWjXpPadNbMbJj6.WkZ33tfndcMYz7cM.x0oDC3Nv6YIvi_Ii5Qs.3w0L.F2T2P.UkbGcVGQtz1a88e9b24wvAfX5Au8qrWaCpgg; FSSBBIl1UgzbN7NO=5Pfq1jqSPKHdqVVj7Gq2y2zMXEFllmI3U.Ufjrfnbczgf_YVSFyyiibIjqeyxH8dqzjrQuJPsz3Y7oVCSRzhNfq")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	// 自动识别网页html编码，并转换为utf-8
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之前是有效的
	// 如果引用的数据长度小于 n，则返回一个错误信息；如果 n 大于缓存的总大小，则返回 ErrBufferFull
	// 通过 Peek 的返回值，可以修改缓存中的数据，但是不能修改底层 io.Reader 中的数据
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
