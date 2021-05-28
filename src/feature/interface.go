package main

import "fmt"

type (
	Usb interface {
		Start()
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
	fmt.Println("phone...")
}

func (camera Camera) Start() {
	fmt.Println("Camera...")
}

func (pc *Pc) Working(usb Usb) {
	usb.Start()
}

func main() {
	pc := Pc{}
	phone := Phone{}
	camera := Camera{}
	pc.Working(phone)
	pc.Working(camera)
}
