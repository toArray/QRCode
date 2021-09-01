package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
)

func main() {
	content := "这是二维码内容"     //二维码内容
	fileName := "output.png" //输出文件名
	fileSize := 256          //二维码大小

	//生成二维码
	err := CreatePNG(content, fileSize, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}

/*
CreatePNG
@Desc	生成二维码
@Param	content 	string	二维码内容
@Param	fileSize 	int		二维码大小
@Param	fileName 	string	文件昵称
*/
func CreatePNG(content string, fileSize int, fileName string) (err error) {
	//生成一个二维码对象
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return
	}

	//q.DisableBorder = true          //去除边框
	q.BackgroundColor = color.White //二维码背景颜色
	q.ForegroundColor = color.Black //二维码条纹颜色

	//写入文件
	err = q.WriteFile(fileSize, fileName)
	if err != nil {
		return
	}

	return
}
