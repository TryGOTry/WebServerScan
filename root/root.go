/*
* @Author: Try
* @Date:   2021/5/5 11:18
 */
package root

import (
	"fmt"
	"github.com/gookit/color"
	"sync"
	"syscall"
	"time"
	"title-scan/getip"
	"title-scan/golimit"
	"title-scan/scan"
	"unsafe"
)

func setTitle(title string) {
	kernel32, _ := syscall.LoadLibrary(`kernel32.dll`)
	sct, _ := syscall.GetProcAddress(kernel32, `SetConsoleTitleW`)
	syscall.Syscall(sct, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	syscall.FreeLibrary(kernel32)
}
func GoWebScan(ipadd string, ipport string, ishttp string, num int, timeout int64) {
	dicall := getip.Getip(ipadd, ipport)
	color.Red.Println("[Info] Web-Server-Scan|Try| By T00ls.Net; ")
	color.Red.Println("[Info] 一款内网web服务探测的工具（ps:当然你要扫公网也没问题）")
	color.Red.Println("[Info] 开始扫描中.当前线程:", num)
	color.Red.Println("---------------------------------------")
	//fmt.Println(dicall)
	g := golimit.NewG(num) //设置线程数量
	wg := &sync.WaitGroup{}
	beg := time.Now()
	for i := 0; i < len(dicall); i++ {
		wg.Add(1)
		task := dicall[i]
		g.Run(func() {
			setTitle("当前目标:" + task + " 任务:" + fmt.Sprintf("%d", i) + "/" + fmt.Sprintf("%d", len(dicall)))
			respBody, err := scan.Goscan(task, ishttp, timeout)
			if err != nil {
				//color.Warn.Println("目标访问错误，可能被ban了！")
				wg.Done()
				return
			}
			if respBody.StatusCode == 200 {
				color.Info.Println("[200] ", respBody.Res+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
				//writefile.Write(url, "[200] "+respBody.Res+"\n")
			} else if respBody.StatusCode == 403 {
				color.Warn.Println("[403] ", respBody.Res+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
				//writefile.Write(url, "[403] "+respBody.Res+"\n")
			} else if respBody.StatusCode == 302 {
				color.Warn.Println("[302] ", respBody.Res+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
				//writefile.Write(url, "[302] "+respBody.Res+"\n")
			}
			wg.Done()
		})
	}
	wg.Wait()
	color.Red.Printf("[info] 扫描完成！当前用时: %fs", time.Now().Sub(beg).Seconds())
}
func GoOneScan(oneip string, ipport string, ishttp string, timeout int64) {
	if getip.IsIp(oneip) {
		a := getip.Getoneip(oneip, ipport)
		color.Red.Println("[Info] Web-Server-Scan|Try| By T00ls.Net; ")
		color.Red.Println("[Info] 一款内网web服务探测的工具（ps:当然你要扫公网也没问题）")
		color.Red.Println("[Info] 开始扫描中.单ip扫描固定2线程")
		color.Red.Println("---------------------------------------")
		g := golimit.NewG(5) //设置线程数量
		wg := &sync.WaitGroup{}
		beg := time.Now()
		for i := 0; i < len(a); i++ {
			wg.Add(1)
			ip := a[i]
			g.Run(func() {
				respBody, err := scan.Goscan(ip, ishttp, timeout)
				if err != nil {
					//color.Warn.Println("目标访问错误，可能被ban了！")
					wg.Done()
					return
				}
				if respBody.StatusCode == 200 {
					color.Info.Println("[200] ", respBody.Res+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
					//writefile.Write(url, "[200] "+respBody.Res+"\n")
				} else if respBody.StatusCode == 403 {
					color.Warn.Println("[403] ", respBody.Res+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
					//writefile.Write(url, "[403] "+respBody.Res+"\n")
				} else if respBody.StatusCode == 302 {
					color.Warn.Println("[302] ", respBody.Res+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
					//writefile.Write(url, "[302] "+respBody.Res+"\n")
				}
				wg.Done()
			})
		}
		wg.Wait()
		color.Red.Printf("[info] 扫描完成！当前用时: %fs", time.Now().Sub(beg).Seconds())
	} else {
		color.Red.Println("[info] 请输入争取的ip地址.")
	}
}
