package main

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

//type Config struct {
//	Header map[string]string `json:"header"`
//	Type   string            `json:"type"`
//	Url    string            `json:"url"`
//}
//
//func main() {
//
//	fmt.Println("hello")
//	jsonData, err := ioutil.ReadFile("config.json")
//	var config Config
//	err = json.Unmarshal(jsonData, &config)
//	client := &http.Client{}
//	data, err := ioutil.ReadFile("data.json")
//
//	req, _ := http.NewRequest(config.Type, config.Url, bytes.NewBuffer(data))
//	for k, v := range config.Header {
//		req.Header.Add(k, v)
//	}
//	resp, err := client.Do(req)
//
//	if err != nil {
//		fmt.Println("读取响应失败:", err)
//	}
//	filePath := "response.txt"
//	responseBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("读取响应失败:", err)
//		return
//	}
//	err = ioutil.WriteFile(filePath, responseBody, 0644)
//	if err != nil {
//		fmt.Println("写入文件失败:", err)
//		return
//	}
//
//	fmt.Println("响应结果已写入到文件:", filePath)
//	defer resp.Body.Close()
//}

func main() {

	fmt.Println(utils.Call("HelmRpc", "Test", system.SysComponent{Host: "1w345"}))
}
