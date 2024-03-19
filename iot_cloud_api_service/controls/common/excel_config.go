package common

import "github.com/tealeg/xlsx"

// excel导出Header部分样式
func ExcelHeaderStyle() *xlsx.Style {
	// 标题样式
	style := xlsx.NewStyle()
	style.Font.Bold = true
	style.ApplyFill = true
	border := xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Border = *border
	style.ApplyBorder = true
	style.Fill = *xlsx.NewFill("solid", "a4c2f4", "FFFFFF")
	return style
}

// excel导出内容部分样式
func ExcelContentStyle() *xlsx.Style {
	// 标题样式
	style := xlsx.NewStyle()
	style.ApplyFill = true
	border := xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Border = *border
	style.ApplyBorder = true
	style.Alignment = xlsx.Alignment{
		//ShrinkToFit: true,//内容根据宽度自适应文字大小
		//WrapText: true, //自动换行
	}
	return style
}
