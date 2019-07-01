package thinkHttp

import (
	"regexp"
	"util/think"
)

func GetPublicIp() string {
	ip := ""
	var err error

	urlsP := make([]string, 0)
	urlsP = append(urlsP, "https://tool.lu/ip/ajax.html")
	for i := 0; i < len(urlsP); i++ {
		url := urlsP[i]
		ip, err = getPublicIpPost(url)
		if err == nil {
			return ip
		}
	}

	urlsG := make([]string, 0)
	urlsG = append(urlsG, "http://www.ip138.com")
	urlsG = append(urlsG, "http://www.net.cn/static/customercare/yourip.asp")
	for i := 0; i < len(urlsG); i++ {
		url := urlsG[i]
		ip, err = getPublicIpGet(url)
		if err == nil {
			return ip
		}
	}

	return ip


}

func getPublicIpPost(url string) (string, error){
	body, err := SendPost(url, map[string]string{}, nil)
	think.IsNil(err)

	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)

	return reg.FindString(string(body)), nil
}

func getPublicIpGet(url string) (string, error){
	body, err := SendGet(url, map[string]string{})
	think.IsNil(err)

	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)

	return reg.FindString(string(body)), nil
}