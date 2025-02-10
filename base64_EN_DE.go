package main

import (
	"bufio"
	"fmt"
	"github.com/dromara/dongle"
	"os"
)

// 清除输入缓冲区
func clearInputBuffer() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // 读取并丢弃缓冲区中的内容
}

func main() {
	var choice int
	var input string

	for {
		// 主菜单
		fmt.Println("请选择功能：")
		fmt.Println("1. 编码")
		fmt.Println("2. 解码")
		fmt.Println("3. 退出程序")
		fmt.Print("请输入选项：")
		fmt.Scanf("%d", &choice)
		clearInputBuffer() // 清除输入缓冲区

		switch choice {
		case 1:
			// 编码功能
			for {
				fmt.Println("当前功能：Base64 编码")
				fmt.Print("请输入要编码的字符串（输入 'exit' 退出编码功能）：")
				fmt.Scanf("%s", &input)
				clearInputBuffer() // 清除输入缓冲区

				if input == "exit" {
					fmt.Println("退出编码功能。")
					break
				}

				// 使用 dongle 进行 Base64 编码
				encoded := dongle.Encode.FromString(input).ByBase64().ToString()
				fmt.Printf("编码后的字符串：%s\n", encoded)
			}

		case 2:
			// 解码功能
			for {
				fmt.Println("当前功能：Base64 解码")
				fmt.Print("请输入要解码的字符串（输入 'exit' 退出解码功能）：")
				fmt.Scanf("%s", &input)
				clearInputBuffer() // 清除输入缓冲区

				if input == "exit" {
					fmt.Println("退出解码功能。")
					break
				}

				// 使用 dongle 进行 Base64 解码
				decoded := dongle.Decode.FromString(input).ByBase64().ToString()
				fmt.Printf("解码后的字符串：%s\n", decoded)
			}

		case 3:
			// 退出程序
			fmt.Println("程序已退出。")
			return

		default:
			fmt.Println("无效的选择，请重新输入！")
		}
	}
}
