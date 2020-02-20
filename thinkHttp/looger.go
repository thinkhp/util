package thinkHttp

import (
	"strconv"
	"time"
	"util/thinkString"
)

var (
	LogPrint = true
	LogPrintResponseBody = true
)
type logger struct {
	uuid string
	print bool
	responseBody bool
}

func (l *logger)Init() *logger{
	l.uuid = newRequestId()

	return l
}

func newRequestId() string {
	uuid := time.Now().Format("20060102150405")
	uuid += thinkString.UUID(8)
	return uuid
}

func (l *logger)SprintRequest(method string, url string, header map[string]string, body []byte) string{
	log := l.uuid
	log += "\n"
	log += "***************************** " +"request,m->>" + " ****************************S\n"
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

func (l *logger)SprintResponse(code int, url string, headers map[string][]string, body []byte) string{
	log := l.uuid
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

func (l *logger)SprintRequestReceive(method string, url string, header map[string][]string, body []byte) string{
	log := l.uuid
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

func (l *logger)SprintResponseSend(headers map[string]string, body []byte) string {
	log := l.uuid
	log += "\n"
	log += "***************************** " + "response,m->>" + " ****************************S\n"
	for key, value := range headers {
		log += key + ":" + value + "\n"
	}
	log += string(body) + "\n"
	log += "***************************** " + "response,m->>" + " ****************************E"

	return log
}


