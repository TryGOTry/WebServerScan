/*
* @Author: Try
* @Date:   2021/5/5 9:59
 */
package main

import (
	"flag"
	"github.com/gookit/color"
	"title-scan/root"
)

func main() {
	ipadd := flag.String("h", "", "ip地址(c段)")
	oneip := flag.String("u", "", "单个地址扫描")
	ipport := flag.String("p", "80,8080", "扫描端口")
	ishttp := flag.String("type", "http", "扫描类型(http,https)默认http")
	num := flag.Int("s", 5, "线程")
	timeout := flag.Int64("t", 2, "超时时间")
	flag.Parse()
	if *ipadd != "" && *oneip == "" {
		t := *ipadd
		a := t[len(t)-3:]
		if a == "/24" {
			if *ishttp != "http" || *ishttp != "https" {
				root.GoWebScan(*ipadd, *ipport, *ishttp, *num, *timeout)
			} else {
				flag.Usage()
			}
		} else {
			color.Red.Println("[info] 请输入正确的ip地址,如:127.0.0.1/24")
		}
	} else if *ipadd == "" && *oneip != "" {
		root.GoOneScan(*oneip, *ipport, *ishttp, *timeout)
	} else {
		flag.Usage()
	}
}
