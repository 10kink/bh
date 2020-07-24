package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
//	"time"
)

func main() {
	fmt.Println("Hello, world")
	fp, _ := os.OpenFile("./123.txt", os.O_CREATE|os.O_APPEND, 6)
	defer fp.Close()
	for i := 7337; i<26000; i++ {

		resp, err := http.Get("http://chaziwang.com/show-" + strconv.Itoa(i) + ".html")
		if err != nil {
			fmt.Println("http get error", err)
			break
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read error", err)
			return
		}
		fmt.Print(charbh(string(body)))
		fmt.Print(i)
		fp.WriteString(charbh(string(body)))

	//	time.Sleep(time.Duration(1) * time.Second)

	}

}

func charbh(str string) string {
	reg1 := regexp.MustCompile(`[\p{Han}]{1}`)
	if reg1.FindString(str) == "查" {return("")}
	reg2 := regexp.MustCompile(`笔顺编号：(\d+)`)
	return (string(reg1.FindString(str)+" "+string([]byte(reg2.FindString(str))[15:])) + "\n")
}
