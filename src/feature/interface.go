package main

import "fmt"

type (
	Usb interface {
		Start()
		Stop()
	}

	// Phone 实现Usb接口
	Phone struct {
	}

	// Camera 实现Usb接口
	Camera struct {
	}

	Pc struct {
	}
)

//只需要分别实现start()方法，即可实现Usb接口

func (phone Phone) Start() {
	fmt.Println("phone...Start")
}

func (phone Phone) Stop() {
	fmt.Println("phone...Stop")
}

func (phone Phone) Call() {
	fmt.Println("phone...Call")
}

func (camera Camera) Start() {
	fmt.Println("Camera...Start")
}

func (camera Camera) Stop() {
	fmt.Println("Camera...Stop")
}

func (pc *Pc) Working(usb Usb) {
	usb.Start()
	// *********类型断言，如果是Phone类型就执行Call()************
	if phone, phoneOk := usb.(Phone); phoneOk {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	pc := Pc{}
	phone := Phone{}
	camera := Camera{}
	pc.Working(phone)
	pc.Working(camera)
}
