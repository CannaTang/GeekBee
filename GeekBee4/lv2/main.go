package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	println("提醒功能如下:\n1.单次日程提醒功能\n2.重复日程提醒功能\n请输入序号:\n")
	var n int
	fmt.Scanln(&n)
	var Time1, Time2, NextTime string
	fmt.Scanln(&Time1, &Time2)

	var build strings.Builder
	build.WriteString(Time1)
	build.WriteString(" ")
	build.WriteString(Time2)
	NextTime = build.String()

	location, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006/01/02 15:04:05", NextTime, location)

	d := time.Until(t)

	fmt.Println("距离计划还有", d)
	println(time.Duration(d.Seconds()))
	//2022/11/05 17:20:50

	switch n {
	case 1:
		{
			timer := time.NewTimer(time.Duration(d.Seconds()) * time.Second)

			for {
				select {
				case <-timer.C:
					fmt.Println("时间到!")
					break
				}
				break
			}
		}
	case 2:
		{
			timer := time.NewTimer(time.Duration(d.Seconds()) * time.Second)

			for {
				select {
				case <-timer.C:
					fmt.Println("时间到!")
					break
				}
				break
			}

			timer1 := time.NewTicker(7 * 24 * time.Hour)
			for range timer1.C {
				fmt.Println("时间到!")
			}
			break
		}
	}
}
