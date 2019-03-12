// 简单的http协议

package main

import (
	"fmt"
	"net/http"
)

func main() {
	request, _ := http.NewRequest("HEAD", "http://www.baidu.com/", nil)
	client := http.Client{}
	resp, _ := client.Do(request)
	fmt.Println(resp.Header)
}
