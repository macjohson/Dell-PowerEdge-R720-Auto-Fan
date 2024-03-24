package main

import (
	"auto-fan/utils"
	"fmt"
	"slices"
	"time"
)

var last int = 0

func main() {
	for {
		// 获取温度数据
		sensorData, err := utils.GetTemp()
		if err != nil {
			fmt.Println("获取温度数据失败:", err)
			continue
		}

		fmt.Print("\033[H\033[2J")
		temps := make([]int, 0)
		// 打印结果
		for _, data := range sensorData {
			fmt.Printf("%s: %d°C\n", data.Label, data.Temp)
			temps = append(temps, data.Temp)
		}

		fmt.Printf("当前转速%d%%", last)

		maxTemp := slices.Max(temps)

		fanSpeed := 10

		limitTemp := 40

		if maxTemp > limitTemp {
			fanSpeed = (maxTemp-limitTemp)*2 + 15
			if fanSpeed > 100 {
				fanSpeed = 100
			}
		}

		if last != fanSpeed {
			utils.SetFan(fanSpeed)
		}

		last = fanSpeed

		time.Sleep(time.Second * 1)
	}
}
