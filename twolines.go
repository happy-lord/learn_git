package main

import (
	"fmt"
	"math"
)

type Line struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}
func main()  {
	var line1,line2 Line
	var k1,k2 float64
	fmt.Println("请输入第一条直线的x1坐标：")
	fmt.Scanf("%f",&line1.x1)
	fmt.Println("请输入第一条直线的y1坐标：")
	fmt.Scanf("%f",&line1.y1)
	fmt.Println("请输入第一条直线的x2坐标：")
	fmt.Scanf("%f",&line1.x2)
	fmt.Println("请输入第一条直线的y2坐标：")
	fmt.Scanf("%f",&line1.y2)
	fmt.Println("请输入第二条直线的x1坐标：")
	fmt.Scanf("%f",&line2.x1)
	fmt.Println("请输入第二条直线的y1坐标：")
	fmt.Scanf("%f",&line2.y1)
	fmt.Println("请输入第二条直线的x2坐标：")
	fmt.Scanf("%f",&line2.x2)
	fmt.Println("请输入第二条直线的y2坐标：")
	fmt.Scanf("%f",&line2.y2)
	if ((line1.x1 == line1.x2) &&(line2.x1 == line2.x2))||((line1.y1 == line1.y2) &&(line2.y1 == line2.y2)) {
		fmt.Println("两条直线是和坐标轴垂直的，他们是平行的")
	}else if (line1.x1 == line1.x2) &&(line2.x1 != line2.x2){
		fmt.Println("两条直线不平行")
	}else if (line1.x1 != line1.x2) &&(line2.x1 == line2.x2){
		fmt.Println("两条直线不平行")
	}else {
		k1 = (line1.y2-line1.y1)/(line1.x2-line1.x1)
		k2 = (line2.y2-line2.y1)/(line2.x2-line2.x1)
		if (k1*k2>0)&&(math.Abs(k1-k2)<1.0e-8) {
			fmt.Println("两条直线平行")
		}else {
			fmt.Println("两条直线不平行，是相交的")
		}
	}

}