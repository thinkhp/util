package thinkHttp

import (
	"strconv"
)

func SprintRequest(method string, url string, header map[string]string, body []byte) string{
	log := ""
	log += "\n"
	log += "***************************** " + "request,m->>" + " ****************************S\n"
	log += method + " " + url + "\n"
	for key, value := range header {
		log += key + ":" + value + "\n"
	}
	if len(body) != 0 {
		log += string(body) + "\n"
	}
	log += "***************************** " + "request,m->>" + " ****************************E"

	return log
}

func SprintResponse(code int, url string, headers map[string][]string, body []byte) string{
	log := ""
	log += "\n"
	log += "***************************** " + "response,->>m" + " ****************************S\n"
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
	log += "***************************** " + "response,->>m" + " ****************************E"

	return log
}

func SprintRequestReceive(method string, url string, header map[string][]string, body []byte) string{
	log := ""
	log += "\n"
	log += "***************************** " + "request,->>m" + " ****************************S\n"
	log += method + " " + url + "\n"
	for hKey, hValue := range header {
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
	log += "***************************** " + "request,->>m" + " ****************************E"

	return log
}

func SprintResponseSend(headers map[string]string, body []byte) string {
	log := ""
	log += "\n"
	log += "***************************** " + "response,m->>" + " ****************************S\n"
	//log += strconv.Itoa(code)+ "\n"
	for key, value := range headers {
		log += key + ":" + value + "\n"
	}
	log += string(body) + "\n"
	log += "***************************** " + "response,m->>" + " ****************************E"

	return log
}


