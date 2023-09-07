package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 定义图像的宽度和高度
	width := 2400
	height := 2400

	// 创建一个空白图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 设置毛玻璃颜色和透明度
	blurColor := color.RGBA{255, 21, 184, 225} // 紫色颜色，透明度5%
	blurSize := 10

	// 创建毛玻璃效果
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			avgR, avgG, avgB, avgA := 0, 0, 0, 0

			// 计算平均颜色
			count := 0
			for i := x - blurSize; i <= x+blurSize; i++ {
				for j := y - blurSize; j <= y+blurSize; j++ {
					if i >= 0 && i < width && j >= 0 && j < height {
						r, g, b, a := blurColor.RGBA()
						avgR += int(r)
						avgG += int(g)
						avgB += int(b)
						avgA += int(a)
						count++
					}
				}
			}

			// 计算平均值
			avgR /= count
			avgG /= count
			avgB /= count
			avgA /= count

			// 设置像素颜色
			c := color.RGBA{uint8(avgR), uint8(avgG), uint8(avgB), uint8(avgA)}
			img.Set(x, y, c)
		}
	}

	// 创建输出文件
	file, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	// 将图像保存为PNG格式
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error encoding PNG:", err)
		return
	}

	fmt.Println("Image saved as output.png")
}
