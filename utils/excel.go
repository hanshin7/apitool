package utils

import (
	"apitool/config"
	"apitool/logging"
	"errors"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//excel数据体
type ExcelData struct {
	//excel文件名称
	ExcelName string
	//sheet名称的切片
	SheetNameSlice []string
	//sheet的title
	SheetTitleSlice [][]string
	//第一层[]对应每个sheet，第二层[]对应该sheet的数据（多行）,第三层[]对应每一行数据
	SheetDataSlice [][][]string
}

//返回文件名、错误标识码
func CreateExcel(exData *ExcelData) bool {
	f := excelize.NewFile()
	// 根据名称创建sheet
	for index, sheetName := range exData.SheetNameSlice {
		f.NewSheet(sheetName)
		//创建当前sheet的表头
		for clumnNum, clumnVal := range exData.SheetTitleSlice[index] {
			sheetPosition := Div(clumnNum+1) + "1"
			f.SetCellValue(sheetName, sheetPosition, clumnVal)
		}
		if exData.SheetDataSlice == nil || len(exData.SheetDataSlice) == 0 {
			continue
		}
		//设置当前sheet的表数据
		for lineNum, lineVal := range exData.SheetDataSlice[index] {
			for clumnNum, clumnVal := range lineVal {
				sheetPosition := Div(clumnNum+1) + strconv.Itoa(lineNum+2)
				f.SetCellValue(sheetName, sheetPosition, clumnVal)
			}
		}
	}
	filepath := config.Conf.Section("sys").Key("file_path").Value() + "/download/" + exData.ExcelName
	if err := f.SaveAs(filepath); err != nil {
		err = errors.New(err.Error())
		return false
	}
	logging.LogD("接口响应报文结果文件写文件完成[%s]\n", filepath)
	return true

}

// Div 数字转字母
func Div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}
