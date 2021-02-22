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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// 查看自己浏览器中的User-Agent信息（检查元素->Network->User-Agent）
	req.Header.Add("Host", "album.zhenai.com")
	req.Header.Add("Cookie", "sid=718e1906-668d-4ff8-bcff-bb2385533726; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1613620141; FSSBBIl1UgzbN7NO=5LQbtj3q65GdOZ6UNdGO6v8f7rCh4FiukHsHK69VhhpCvC45VJqQiaTnIxxA3G8rde9IRv2G.rWUuqkyFHGobgq; ec=FVTrmidR-1613620286783-ca5a54324df1c557364444; spliterabparams=1613623253824%3A-9082433820229133149; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1613623572; FSSBBIl1UgzbN7NP=53c6SmbqzZpWqqqm676UchqJze15HZSPduxchLrxV30xPhLAdtQLoO4DHQBRbELn9CfsguQgooiTZcHm.mRCDsLQnSI2czdkrqA5S37VxSdKry8ycJm.w9qqQX00AXJGvejKTATpQFxGpK2s_pIof.ExGFjd2Bo3aIZk_lNTVo9moa5VWs0j5rYBYbgwhpqzk2tZF5WF5DZozBJNkkrsUu0D0Fz_NZO2MgD83l7lHcU68s3hP9h0VzLP0A1Etnn7h7; _efmdata=QXu7T9wveKjnnR%2FrSgla%2F4aicNBhKEnpPPoMOumutFJVBPXPGWPDCofuTQrm%2FyB37842iDHvGCTR5OK1PD4QjQ%2Fkctz3j9KZpb4%2BgQmHHlk%3D; _exid=gjjH7Dwdeamy6v59SLZ0E%2B7uhhjQCG5jDa1qM6meUHi8CiMFEK7jaDRvJj6jFVH1S6kSrBEDqqa7zs%2FSEWq25w%3D%3D; FSSBBIl1UgzbN7NO=5IwD.xEiZNGGechW9ZveRQnO8mskxojTAtSZnDSbWgHNrm4EAiZk0_GyXEpeSSq1fmC.c949EMAKX6.sYhEcZea")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"88\", \"Google Chrome\";v=\"88\", \";Not A Brand\";v=\"99\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("dnt", "1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("referer", "https://album.zhenai.com/u/1428069052")
	req.Header.Add("accept-language", "zh")

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
