package thinkHttp

import (
	"strconv"
)

func SprintRequest(method string, url string, header map[string]string, body []byte) string{
	log := ""
	log += "\n"
	log += "***************************** " + "request" + " ****************************\n"
	log += method + " " + url + "\n"
	for key, value := range header {
		log += key + ":" + value + "\n"
	}
	if len(body) != 0 {
		log += string(body) + "\n"
	}
	log += "***************************** " + "request" + " ****************************\n"

	return log
}

func SprintResponse(code int, url string, headers map[string][]string, body []byte) string{
	log := ""
	log += "\n"
	log += "***************************** " + "response" + " ****************************\n"
	log += strconv.Itoa(code)+ " " + url + "\n"
	for hKey, hValue := range headers {
		log += hKey + ":"
		for k, v := range hValue {
			log += v
			if k != len(hValue)-1 {
				log += ","
			}
		}
		log += "\n"

	}
	if len(body) != 0 {
		log += string(body) + "\n"
	}
	log += "***************************** " + "response" + " ****************************\n"

	return log
}

