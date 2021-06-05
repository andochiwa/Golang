package main

import "fmt"

func main() {
	// 接收用户的选择
	key := 0
	for loop := true; loop; {
		fmt.Println("--------------欢迎登陆多人聊天系统----------------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择 (1-3):")

		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			fmt.Println("输入有误 请重新输入")
			continue
		}
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入不正确，请重新输入")
			continue
		}
	}

	// 执行逻辑
	switch key {
	case 1:
		for {
			var id int
			var password string
			fmt.Println("输入用户id")
			_, err := fmt.Scanf("%d\n", &id)
			if err != nil {
				fmt.Println("输入格式有误 请重新输入")
				continue
			}
			fmt.Println("输入用户密码")
			_, err = fmt.Scanf("%s\n", &password)
			if err != nil {
				fmt.Println("输入格式有误 请重新输入")
				continue
			}

			err = login(id, password)
			if err != nil {
				fmt.Println("用户名或密码有误，请重新输入")
				continue
			}
		}
	case 2:
		fmt.Println("注册用户")
	case 3:
		fmt.Println("退出系统")
	default:
		fmt.Println("输入不正确，请重新输入")
	}
}
