package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

/**
*
* Author: Marek
* Date: 2024-02-15 20:12
* Email: 364021318@qq.com
*
 */

/*
TCP 扫描器

Go 在网络应用编程方面堪称完美。它自带标准库也很优秀，在开发过程中可以给予我们很多帮助。
在本文中，我们将会用Go写一个简单的TCP扫描器。整个程序的代码在50行以内。在我们开始动手之前，先
介绍一些理论知识：
不得不说，TCP是比我们介绍的要复杂的多的，但是我们只介绍一点基础知识。TCP的握手有三个过程。首
先，客户端发送一个syn 的包，表示建立回话的开始。如果客户端收到超时，说明端口可能在防业墙后面，
第二，如果服务端应答 syn-ack 包， 意味着这个端口是打开的，否则会返回rst包。最后，客户端需要另外
发送一个ack 包。 从这时起，连接就已经建立。
我们TCP扫描器第一步实现单个端口的测试。使用标准库中的 net.Dial 函数， 该函数接收两个参数： 协议
和测试地址（带端口号）。

 */

//func main() {
//	_, err := net.Dial("tcp", "google.com:80")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Connection successful")
//}


/*
为了不一个一个地测试每一个端口，我们将添加一个简单的循环来简化整个测试过程。
 */

//func main() {
//	for port := 80; port < 100; port++ {
//		conn, err := net.Dial("tcp", fmt.Sprintf("google.com:%d", port))
//		if err == nil {
//			conn.Close()
//			fmt.Println("Connection successful")
//		} else {
//			fmt.Println(err)
//		}
//	}
//}
// 这种处理方式有个很大的问， 极度的慢。我们可以通过两个操作来处理一下：并行的执行及为每个连接添加超时控制。

// 我们来看下如何实现并行。第一步先把扫描功能拆分为一个独立函数。这样会使我们代码看起来清晰。

//func isOpen(host string, port int) bool {
//	time.Sleep(time.Millisecond * 1)
//	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
//	if err == nil {
//		_ = conn.Close()
//		return true
//	}
//
//	return false
//}
// 我们会引入一个新的方法 WaitGroup, 详细用法信息可以参考标准库文档。 在主函数中， 我们可以拆分为协程支执行，然后等待执行结束。
// 使用 Groutuine 的方式可以减少  遍历执行时的等待时间， 使用GPU的并发调度

//func main() {
//	ports := []int{}
//
//	wg := &sync.WaitGroup{}
//
//	for port := 1; port < 100; port++ {
//		wg.Add(1)
//		go func() {
//			opened := isOpen("google.com", port)
//			if opened {
//				ports = append(ports, port)
//			}
//
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//	fmt.Printf("opened ports: %v\n", ports)
//}

// 我们的代码已经执行的很快了，但是由于超时的原因，我们需要等待很久才能收到返回的错误信息。我们可
// 以假设如果我们200毫秒内没有收到服务器的回应，就不再继续等待。

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		_ = conn.Close()
		return true
	}

	return false
}

func main() {

	// 最后添加
	hostname := flag.String("hostname",  "","hostname to test")
	startPort := flag.Int("start-port", 80, "theport on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port from which the scaning ends")
	timeout:= flag.Duration("timeout", time.Millisecond * 200, "timeout")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	// 现在唯一的问题就是，现在这个程序会有竟争条件。在只扫描少数端口时，速度比较慢，可能不会出现，但确实
	// 存在这个问题。所以我们需要使用 mutex 来修复它。
	mutex := &sync.Mutex{}
	for port := *startPort; port < *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}

			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("opened ports : %v\n", ports)
}

// 至此， 我们得到了一个简单的端口扫描器。但有些不好的是， 不能很方便的修改域名地址以及端口号范转围，
// 我们必须要重新编译代码才可以。 Go 还有一个很不错的包叫做 flag .
// flag 包 可以帮助我们编写命令行程序。 我们可以配置每个字符串或数字。 我们为主机名及要测试的端口范围
// 和连接超时添加参数
