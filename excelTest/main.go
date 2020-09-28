package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

func main() {
	const (
		fileName  = "附件三：285号和313号样本原始数据.xlsx"
		sheetName = "操作变量"
		sheetNum  = 4  // sheet的脚标
		d1s       = 3  // 第一份的起始行
		d1e       = 43 // 第一份的终止行
		d2s       = 44 // 第二份的起始行
		d2e       = 84 // 第二份的终止行
	)
	// 读取文件
	file, err := xlsx.OpenFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %s", err)
		return
	}
	fmt.Printf("读取%s文件成功\n", fileName)
	sheet := file.Sheet[sheetName]
	// 计算出一共有多少列
	w := sheet.MaxCol
	all, _ := file.ToSlice()
	s4 := all[sheetNum]
	fmt.Printf("读取 %s 表内容成功\n", sheetName)
	// 列
	for i := 1; i < w; i++ {
		// 第一份数据的处理
		var d []float64
		for j := d1s; j < d1e; j++ {
			in, _ := strconv.ParseFloat(s4[j][i], 10)
			d = append(d, in)
		}
		f, z := getAvg(d)
		if f != 0 {
			for _, t := range z {
				h := t + d1s
				fmt.Printf("发现第%d行的%d列值为0，已经更改为 %f\n", h, i, f)
				s4[h][i] = strconv.FormatFloat(f, 'E', -1, 64)
			}
		}
		// 第二份数据的处理
		var d2 []float64
		for j := d2s; j < d2e; j++ {
			in, _ := strconv.ParseFloat(s4[j][i], 10)
			d2 = append(d2, in)
		}
		f2, z2 := getAvg(d2)
		if f2 != 0 {
			for _, t := range z2 {
				h := t + d2s
				fmt.Printf("发现第%d行的%d列值为0，已经更改为 %f\n", h, i, f2)
				s4[h][i] = strconv.FormatFloat(f2, 'E', -1, 64)
			}
		}
	}
	// 重新生成文件
	writingXlsx(s4)
	fmt.Println("重新生成文件成功")
	fmt.Println("按Enter退出")
	fmt.Scanln()
}

// 求取平均值
func getAvg(s []float64) (f float64, z []int) {
	var sum, count float64
	for i := 0; i < len(s); i++ {
		// 记录下当前值为0 的位置，且跳过
		if s[i] == 0 {
			z = append(z, i)
		}
		count++
		sum += s[i]
	}
	f = sum / count
	return
}

// 把结果写入文件
func writingXlsx(oraList [][]string) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var cell *xlsx.Cell
	var err error
	// 声明一个新文件
	file = xlsx.NewFile()
	// 添加一个sheet
	sheet, err = file.AddSheet("结果数据")
	if err != nil {
		fmt.Printf(err.Error())
	}
	// 循环输入数据
	for _, i := range oraList {
		var row1 *xlsx.Row
		row1 = sheet.AddRow()
		for _, j := range i {
			cell = row1.AddCell()
			cell.Value = j
		}
	}
	// 保存结果
	err = file.Save("结果数据.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
