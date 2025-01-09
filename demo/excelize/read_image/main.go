package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

// 读取 excel 中的图片，图片左上角位于哪个单元格就属于哪个单元格

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	pics, err := f.GetPictures("Sheet1", "A2")
	if err != nil {
		fmt.Println(err)
	}
	for idx, pic := range pics {
		fmt.Println(idx)
		name := fmt.Sprintf("image%d%s", idx+1, pic.Extension)
		if err := os.WriteFile(name, pic.File, 0644); err != nil {
			fmt.Println(err)
		}
	}

}
