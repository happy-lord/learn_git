package main

import "fmt"

type Person struct {
	name string
	sex string
	height float64
	weight float64
	age int
	BMI,fatRate float64
	suggest string

}

func Calculator(sex string,height float64,weight float64,age int)(float64,float64,string){
	var (
		BMI, fatRate float64
		sexNum    int
		suggest string
	)
	if sex == "男" {
		sexNum = 1
	} else {
		sexNum = 0
	}
	BMI = weight / (height * height)
	fatRate = 1.2*BMI + 0.23*float64(age) - 5.4 - 10.8*float64(sexNum)
	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fatRate >= 10 && fatRate <= 10 {
				suggest = "偏瘦，请多吃肉"
			} else if fatRate >= 11 && fatRate <= 16 {
				suggest = "标准，太棒了，请继续保持"
			} else if fatRate >= 17 && fatRate <= 21 {
				suggest = "偏重，吃完饭多散散步"
			} else if fatRate >= 22 && fatRate <= 26 {
				suggest = "肥胖，少吃点，多运动运动"
			} else {
				suggest = "严重肥胖，赶紧健身游泳，跑步"
			}

		} else if age >= 40 && age <= 59 {
			if fatRate >= 0 && fatRate <= 11 {
				suggest = "偏瘦，请多吃肉"
			} else if fatRate >= 12 && fatRate <= 17 {
				suggest = "标准，太棒了，请继续保持"
			} else if fatRate >= 18 && fatRate <= 22 {
				suggest = "偏重，吃完饭多散散步"
			} else if fatRate >= 23 && fatRate <= 27 {
				suggest = "肥胖，少吃点，多运动运动"
			} else {
				suggest = "严重肥胖，赶紧健身游泳，跑步"
			}
		} else {
			if fatRate >= 0 && fatRate <= 13 {
				suggest = "偏瘦，请多吃肉"
			} else if fatRate >= 14 && fatRate <= 19 {
				suggest = "标准，太棒了，请继续保持"
			} else if fatRate >= 20 && fatRate <= 24 {
				suggest = "偏重，吃完饭多散散步"
			} else if fatRate >= 25 && fatRate <= 29 {
				suggest = "肥胖，少吃点，多运动运动"
			} else {
				suggest = "严重肥胖，赶紧健身游泳，跑步"
			}

		}

	} else {
		if age >= 18 && age <= 39 {
			if fatRate >= 0 && fatRate <= 20 {
				suggest = "偏瘦，请多吃肉"
			} else if fatRate >= 21 && fatRate <= 27 {
				suggest = "标准，太棒了，请继续保持"
			} else if fatRate >= 28 && fatRate <= 34 {
				suggest = "偏重，吃完饭多散散步"
			} else if fatRate >= 35 && fatRate <= 39 {
				suggest = "肥胖，少吃点，多运动运动"
			} else {
				suggest = "严重肥胖，赶紧健身游泳，跑步"
			}
		} else if age >= 40 && age <= 59 {
			if fatRate >= 0 && fatRate <= 21 {
				suggest = "偏瘦，请多吃肉"
			} else if fatRate >= 22 && fatRate <= 28 {
				suggest = "标准，太棒了，请继续保持"
			} else if fatRate >= 29 && fatRate <= 35 {
				suggest = "偏重，吃完饭多散散步"
			} else if fatRate >= 36 && fatRate <= 40 {
				suggest = "肥胖，少吃点，多运动运动"
			} else {
				suggest = "严重肥胖，赶紧健身游泳，跑步"
			}
		} else {
			if fatRate >= 0 && fatRate <= 22 {
				suggest = "偏瘦，请多吃肉"
			} else if fatRate >= 23 && fatRate <= 29 {
				suggest = "标准，太棒了，请继续保持"
			} else if fatRate >= 30 && fatRate <= 36 {
				suggest = "偏重，吃完饭多散散步"
			} else if fatRate >= 37 && fatRate <= 41 {
				suggest = "肥胖，少吃点，多运动运动"
			} else {
				suggest = "严重肥胖，赶紧健身游泳，跑步"
			}
		}

	}
	return BMI,fatRate,suggest
}

func main() {
	var flag string = "y"
	var count int =0
	var persons[5] Person
	var sumFatRate,aveFatRate float64

	for flag=="y"||flag=="Y"{
		//fmt.Printf("请输入第%d个人的信息，用逗号分隔（姓名、性别（男/女）、身高/米、体重/kg、年龄）：\n",count)
		//fmt.Scanf("%s,%s,%f,%f,%d",&persons[count].name,&persons[count].sex,&persons[count].height,&persons[count].weight,&persons[count].age)
		fmt.Printf("请输入第%d个人的姓名：\n",count)
		fmt.Scanf("%s",&persons[count].name)
		fmt.Printf("请输入第%d个人的性别（男/女）：\n",count)
		fmt.Scanf("%s",&persons[count].sex)
		fmt.Printf("请输入第%d个人的身高/米：\n",count)
		fmt.Scanf("%f",&persons[count].height)
		fmt.Printf("请输入第%d个人的体重/kg：\n",count)
		fmt.Scanf("%f",&persons[count].weight)
		fmt.Printf("请输入第%d个人的年龄：\n",count)
		fmt.Scanf("%d",&persons[count].age)
		count++
		fmt.Println("还继续输入吗？y/n:")
		fmt.Scanf("%s",&flag)
	}

	for i:=0;i<count;i++{
		persons[i].BMI, persons[i].fatRate,persons[i].suggest = Calculator(persons[i].sex,persons[i].height,persons[i].weight,persons[i].age)
		fmt.Printf("姓名：%s,BMI:%.3f,体脂率(%%)：%.3f,建议：%s\n",persons[i].name,persons[i].BMI,persons[i].fatRate,persons[i].suggest)
	    sumFatRate+=persons[i].fatRate
	}
	aveFatRate=sumFatRate/float64(count)
	fmt.Printf("总人数为：%d,平均体脂率为(%%)：%.3f",count,aveFatRate)



}
