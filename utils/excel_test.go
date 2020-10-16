package utils

import "testing"

func TestCreateExcel(t *testing.T) {
	exData := ExcelData{
		ExcelName:       "test001",
		SheetNameSlice:  []string{"sheet1", "sheet2"},
		SheetTitleSlice: [][]string{{"t1", "t2", "t3"}, {"tt1", "tt2"}},
		SheetDataSlice: [][][]string{
			{{"v1", "v2", "v3"}, {"v11", "t22", "v23"}},
			{{"p1", "p2"}, {"p11", "p22"}}},
	}
	CreateExcel(&exData)

}

func TestCreateExce2(t *testing.T) {
	exData2 := ExcelData{
		ExcelName:       "test002",
		SheetNameSlice:  []string{"sheet1"},
		SheetTitleSlice: [][]string{{"t1", "t2", "t3"}},
		SheetDataSlice: [][][]string{
			{{"v1", "v2", "v3"}, {"v11", "t22", "v23"}},
		},
	}
	CreateExcel(&exData2)
}
