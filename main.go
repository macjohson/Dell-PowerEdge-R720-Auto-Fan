package main

import (
	"auto-fan/utils"
	"fmt"
	"os"
	"slices"
	"time"

	"gopkg.in/yaml.v3"
)

var last int = 0

var manual = false

type (
	Config struct {
		Limit  int  `yaml:"limit"`
		Base   int  `yaml:"base"`
		Offset int  `yaml:"offset"`
		Auto   bool `yaml:"auto"`
	}
)

func main() {

	for {
		file, err := os.ReadFile("/etc/auto-fan/config.yaml")

		if err != nil {
			panic(err)
		}

		config := &Config{}

		err = yaml.Unmarshal(file, config)

		if err != nil {
			panic(err)
		}

		if !config.Auto {

			if !manual {
				utils.SetManual()
				manual = true

				fmt.Println("NOW IS MANUAL")
			}

			time.Sleep(time.Second * 2)
			continue
		}

		manual = false

		// 获取温度数据
		sensorData, err := utils.GetTemp()
		if err != nil {
			fmt.Println("get temp fail:", err)
			continue
		}

		temps := make([]int, 0)
		// 打印结果
		for _, data := range sensorData {
			fmt.Printf("%s: %d°C\n", data.Label, data.Temp)
			temps = append(temps, data.Temp)
		}

		maxTemp := slices.Max(temps)

		fanSpeed := config.Base

		limitTemp := config.Limit

		if maxTemp > limitTemp {
			fanSpeed = (maxTemp-limitTemp)*config.Offset + config.Base
			if fanSpeed > 100 {
				fanSpeed = 100
			}
		}

		if last != fanSpeed {
			utils.SetFan(fanSpeed)
		}

		fmt.Printf("Fan Speed: %d%%\n\n\n", fanSpeed)

		last = fanSpeed

		time.Sleep(time.Second * 2)
	}
}
