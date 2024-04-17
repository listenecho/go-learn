package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 获取所有网络接口的名称
	ifaceNames, err := netInterfaceNames()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 定义一个map来存储每个网络接口的上一次的上传和下载字节数
	prevBytes := make(map[string]uint64)

	// 无限循环，实时监测上传下载流量
	for {
		for _, name := range ifaceNames {
			bytes, err := getInterfaceBytes(name)
			if err != nil {
				fmt.Printf("Error getting bytes for interface %s: %v\n", name, err)
				continue
			}

			prevByte, ok := prevBytes[name]
			if !ok {
				prevBytes[name] = bytes
				continue
			}

			// 计算增量字节数
			txBytes := bytes - prevByte
			prevBytes[name] = bytes

			// 输出结果
			fmt.Printf("Interface: %s, Upload: %d bytes/sec\n", name, txBytes)
		}

		// 等待一秒钟
		time.Sleep(time.Second)
	}
}

// 获取所有网络接口的名称
func netInterfaceNames() ([]string, error) {
	files, err := ioutil.ReadDir("/sys/class/net")
	if err != nil {
		return nil, err
	}

	var names []string
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		names = append(names, file.Name())
	}
	return names, nil
}

// 获取指定网络接口的字节数
func getInterfaceBytes(ifaceName string) (uint64, error) {
	contents, err := ioutil.ReadFile(fmt.Sprintf("/sys/class/net/%s/statistics/%s", ifaceName, "tx_bytes"))
	if err != nil {
		return 0, err
	}
	txBytesStr := strings.TrimSpace(string(contents))
	txBytes, err := strconv.ParseUint(txBytesStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return txBytes, nil
}
