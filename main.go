package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	payload := url.Values{
		"userId":          {""},
		"password":        {""},
		"service":         {""},
		"queryString":     {"wlanuserip%3D1a19c729090174d9e042d333a6a948b7%26wlanacname%3Dc395c55fab377b7ea4bc1ebb3ce73975%26ssid%3D%26nasip%3Dc54e10c9b9c8d2cf77499cb55b7197d6%26snmpagentip%3D%26mac%3Dd7ccc973a8695c7c053bacb8fa0fe0b9%26t%3Dwireless-v2%26url%3Dbce0e6cfc64b38a70de132171bb75a3c9c9fd810bbc35379df1f2cd44001cac3edab421ca50cdf5e%26apmac%3D%26nasid%3Dc395c55fab377b7ea4bc1ebb3ce73975%26vid%3D1a2c39896d439e83%26port%3Dfd0a75785d9447bd%26nasportid%3Dd4fc8695056bf48b4ed4e589f32650aa6383fe40e201dbb405798cd4a6722e96ac8379dccd7907a0"},
		"operatorPwd":     {""},
		"operatorUserId":  {""},
		"validcode":       {""},
		"passwordEncrypt": {"false"},
	}
	resp, err := http.PostForm("http://172.26.156.158/eportal/InterFace.do?method=login", payload)
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}
