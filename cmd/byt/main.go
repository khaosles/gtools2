package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"

	"github.com/khaosles/gtools2/components/g/sqlite"
	"golang.org/x/sys/unix"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/15 21:10
   @Desc:
*/

var a = sqlite.NewSqlite(nil)

const ONE = 209 * 2

func main() {
	filename := "/Users/yherguot/code/GolandProjects/gtools2/cmd/byt/data_15.bin"

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	// 获取文件大小
	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println("Failed to get file size:", err)
		return
	}
	// 将文件映射到内存
	mappedView, err := unix.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println("Failed to mmap file:", err)
		return
	}

	startTime := time.Now()
	//wg := sync.WaitGroup{}
	var i int64
	for i = 0; i < 1; i++ {
		//wg.Add(1)
		go func(i int64) {
			//defer wg.Done()
			offset := ONE * i
			var data = make([]byte, ONE)
			copy(data, mappedView[offset:offset+ONE])
			var val int16
			for j := 0; j < len(data); j += 2 {
				err = binary.Read(bytes.NewReader(data[j:j+2]), binary.LittleEndian, &val)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				//fmt.Println("并发数 -> ", runtime.NumGoroutine())
				fmt.Print(val, " ")
			}
		}(i)
	}
	//wg.Wait()
	fmt.Println()
	fmt.Println("耗时: ", time.Now().Sub(startTime).Microseconds())
	//// 刷新数据到硬盘
	//err = unix.Msync(mappedView, syscall.MS_SYNC)
	//if err != nil {
	//	fmt.Println("Failed to sync data to disk:", err)
	//}
	time.Sleep(time.Minute)
	func() {
		err = unix.Munmap(mappedView)
		if err != nil {
			fmt.Println(err)
		}
		mappedView, err = unix.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ, syscall.MAP_SHARED)
		fmt.Println("刷新数据完成")
		if err != nil {
			fmt.Println("Failed to mmap file:", err)
			return
		}
		func(i int64) {
			//defer wg.Done()
			offset := ONE * i
			var data = make([]byte, ONE)
			copy(data, mappedView[offset:offset+ONE])
			var val int16
			for j := 0; j < len(data); j += 2 {
				err = binary.Read(bytes.NewReader(data[j:j+2]), binary.LittleEndian, &val)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				//fmt.Println("并发数 -> ", runtime.NumGoroutine())
				fmt.Print(val, " ")
			}
		}(0)
	}()
	fmt.Println("Data refreshed successfully.")
}
