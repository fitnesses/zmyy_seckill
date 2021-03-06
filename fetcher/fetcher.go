package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/util"
)

func Fetch(url string, headers map[string]string) ([]byte, error) {
	consts.RequestLimitRate.Limit()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
func FetchBigResp(url string, headers map[string]string, prefix string) error {
	consts.RequestLimitRate.Limit()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	b, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil || resp.StatusCode != http.StatusOK || b < 100 {
		fmt.Printf("获取验证码图片失败，请求可能被禁止！code： %d\n", resp.StatusCode)
		return errors.New("获取验证码图片失败，请求可能被禁止！")
	}
	path := util.GetCurrentPath()
	path = path + "/imgs/" + prefix
	f, _ := os.Create(path)
	defer f.Close()

	buf := make([]byte, 1024*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		f.Write(buf[:n])
	}
	return nil
}
