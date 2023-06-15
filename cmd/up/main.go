package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/15 22:16
   @Desc:
*/

func main() {

	// 打开文件
	file, err := os.OpenFile("/Users/yherguot/code/GolandProjects/gtools2/cmd/byt/data_15.bin", os.O_RDWR, 0644)
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
	mappedView, err := unix.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println("Failed to mmap file:", err)
		return
	}

	file1, err := os.Open("/Users/yherguot/code/GolandProjects/gtools2/cmd/byt/data.bin")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}
	size := stat.Size()
	length := size / 209 / 2
	fmt.Println(length)
	sl := SplitArray(length, 100000)
	fmt.Println(sl)
	wg := sync.WaitGroup{}
	for _, item := range sl {
		wg.Add(1)
		go func(item [2]int64) {
			defer wg.Done()
			var i int64
			for i = item[0]; i < item[1]; i++ {
				var d = make([]byte, 209*2)
				off := 209 * 2 * (item[1] - i)
				_, err = file1.ReadAt(d, off)
				copy(mappedView[i*209*2:i*209*2+209*2], d)
			}
		}(item)
	}
	wg.Wait()
	err = unix.Msync(mappedView, syscall.MS_SYNC)
	if err != nil {
		fmt.Println("Failed to sync data to disk:", err)
	}

	fmt.Println("更新完成")
}

func SplitArray(length int64, size int64) [][2]int64 {
	var result [][2]int64
	// length := len(arr)
	var i int64
	for i = 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		result = append(result, [2]int64{i, end})
	}
	return result
}
