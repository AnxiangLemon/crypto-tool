package main

import (
	"bufio"
	"crypto-tool/diff"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var (
	// 存储上一步操作的结果
	previousResult []byte
)

func main() {

	fmt.Println("阿银的小工具~")

	for {
		// 打印菜单
		fmt.Println("\n请选择一个操作:")
		fmt.Println("1. 计算文件或字符串的 MD5")
		fmt.Println("...")
		fmt.Println("4. TextDiff")
		fmt.Println("5. 退出")

		// 获取用户输入
		var choice int
		fmt.Print("输入序号: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("输入无效，请输入一个数字")
			continue
		}
		//win 换行是\r\n
		if runtime.GOOS == "windows" {
			bufio.NewReader(os.Stdin).ReadString('\n')
		}
		switch choice {
		case 1:
			calculateMD5()
		case 4:
			callDiff()
		case 5:
			fmt.Println("退出程序")
			return
		default:
			fmt.Println("无效的选择，请重新输入")
		}
	}
}

func calculateMD5() {
	for {
		fmt.Println("\nMD5计算方式 请选择输入类型:")
		fmt.Println("1. 文件")
		fmt.Println("2. 十六进制")
		fmt.Println("3. 字符串")
		fmt.Println("4. 返回主菜单")

		var choice int
		fmt.Print("输入序号: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("输入无效，请输入一个数字")
			continue
		}

		if choice == 4 {
			return
		}
		var input string
		var data []byte
		switch choice {
		case 1:
			// 处理文件
			fmt.Print("请输入文件路径: ")
			fmt.Scan(&input)
			//linux 拉取会有个这字符
			input = strings.ReplaceAll(input, "'", "")
			data, err = os.ReadFile(input)
			if err != nil {
				fmt.Println("读取文件时出错:", err)
				continue
			}
		case 2:
			// 处理十六进制字符串
			fmt.Print("请输入十六进制字符串: ")
			fmt.Scan(&input)
			input = strings.ReplaceAll(input, " ", "") // 去除所有空格
			data, err = hex.DecodeString(input)
			if err != nil {
				fmt.Println("无效的十六进制字符串:", err)
				continue
			}
		case 3:
			// 处理普通字符串
			fmt.Print("请输入字符串: ")
			fmt.Scan(&input)
			data = []byte(input)

		default:
			fmt.Println("无效的选择，请重新输入")
			continue
		}

		// 计算 MD5 哈希
		hash := md5.Sum(data)
		previousResult = hash[:]
		fmt.Printf("MD5: %x\n", previousResult)
	}
}

func callDiff() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入第①个Hex: ")
	text1, _ := reader.ReadString('\n')
	fmt.Print("请输入第②个Hex: ")
	text2, _ := reader.ReadString('\n')
	diff.DiffText(text1, text2)
}
