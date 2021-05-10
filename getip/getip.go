/*
* @Author: Try
* @Date:   2021/5/5 11:45
* 处理c段
 */
package getip

import (
	"fmt"
	"github.com/ozgio/strutil"
	"regexp"
	"strings"
)

func Getip(ipadd string, ipport string) []string { //将ip转换为c段数组
	var ips string
	allip := strings.Replace(ipadd, "/24", "999", -1)
	ipps := strutil.Words(ipport)
	//fmt.Println(allip)
	a := strutil.Words(allip)
	//fmt.Println(a[3])
	ips = strings.Replace(allip, a[3], "", 1)
	//fmt.Println(ips)
	var s []string
	for a := 1; a <= 255; a++ {
		for i := 0; i < len(ipps); i++ {
			//fmt.Printf("a 的值为: %d\n", a)
			ipp := ipps[i]
			b := fmt.Sprintf("%d", a)
			ipc := ips + b + ":" + ipp + "/"
			s = append(s, ipc)
		}
	}
	return s
}
func Getoneip(ipadd string, ipport string) []string {
	var s []string
	var ips string
	ips = ipadd
	ipps := strutil.Words(ipport)
	for i := 0; i < len(ipps); i++ {
		//fmt.Printf("a 的值为: %d\n", a)
		ipp := ipps[i]
		ipc := ips + ":" + ipp + "/"
		s = append(s, ipc)
	}
	return s
}
func IsIp(ip string) (b bool) { //验证ip合法性
	if m, _ := regexp.MatchString("^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
		return false
	}
	return true
}
