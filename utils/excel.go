package utils

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateExcel(excelName string) {
	// 列标题
	titles := map[string][]string{
		"sheet1": {"姓名", "年龄"},
		"sheet2": {"姓名1", "年龄1", "身高1"},
	}
	//存放所有sheet的名称
	sheetList := []string{}
	for k, _ := range titles {
		sheetList = append(sheetList, k)
	}
	// 数据源
	data := [][]map[string]interface{}{
		//sheet1数据
		{
			map[string]interface{}{"name": "jack", "age": 18},
			map[string]interface{}{"name": "mary", "age": 28},
		},
		//sheet2数据
		{
			map[string]interface{}{"name": "jack1", "age1": 18, "sg1": 180},
			map[string]interface{}{"name": "mary1", "age1": 28, "sq1": 175},
		},
	}

	f := excelize.NewFile()
	// 根据名称创建sheet
	for k, _ := range titles {
		f.NewSheet(k)
	}

	for sheetName, sheetTitle := range titles {
		//遍历当前sheet的列
		for clumnNum, clumnVal := range sheetTitle {
			sheetPosition := Div(clumnNum+1) + "1"
			//fmt.Print(sheetPosition)
			f.SetCellValue(sheetName, sheetPosition, clumnVal)
			println(clumnVal)
		}

	}
	for sheetNum, sheetData := range data {
		for lineNum, v := range sheetData {
			clumnNum := 0
			for _, vv := range v {
				clumnNum++
				sheetPosition := Div(clumnNum) + strconv.Itoa(lineNum+2)
				switch vv.(type) {
				case string:

					f.SetCellValue(sheetList[sheetNum], sheetPosition, vv.(string))
					break
				case int:
					f.SetCellValue(sheetList[sheetNum], sheetPosition, vv.(int))
					break
				case float64:
					f.SetCellValue(sheetList[sheetNum], sheetPosition, vv.(float64))
					break
				}
			}
		}
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(1)
	// Save xlsx file by the given path.
	if err := f.SaveAs(excelName + ".xlsx"); err != nil {
		println(err.Error())
	}
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
