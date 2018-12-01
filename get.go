package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/bitly/go-simplejson"
)

func main() {
	spider("https://api.map.baidu.com/?qt=s&c=233&wd=%E7%BE%8E%E9%A3%9F&rn=10&ie=utf-8&oue=1&res=api")
	//rn 每页包含的项目数
	//pn 页数
}

func spider(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()
	var data string
	buf := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if 0 == n {
			break
		}

		data = data + string(buf[:n])
	}
	//fmt.Println(data)


	//解析json
	js, err := simplejson.NewJson([]byte(data))
	if err != nil {
		log.Fatal(err)
	}

	content,_ := js.Get("content").Array()
	//fmt.Println(content)
	for i, _ := range content {
		point := js.Get("content").GetIndex(i)
		//fmt.Println(point)
		x := point.Get("x").MustInt()
		y := point.Get("y").MustInt()
		name := point.Get("name").MustString()
		comnum := point.Get("ext").Get("detail_info").Get("comment_num").MustInt()
		price := point.Get("ext").Get("detail_info").Get("price").MustString()
	    fmt.Printf("%d,%d,%s,%d,%s\n",x,y,name,comnum,price)
	    //fmt.Println(err)
	}
}
