package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type SensorData struct {
	Label  string
	ID     string
	Status string
	Value  string
	Temp   int
}

func GetTemp(account *Account) ([]SensorData, error) {
	args := account.Command([]string{"sdr", "type", "temperature"})
	// 执行 ipmitool 命令
	cmd := exec.Command(args[0], args[1:]...)

	fmt.Println(cmd.String())

	// 捕获命令输出
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("执行命令失败: %w", err)
	}

	// 解析命令输出
	sensorData := parseOutput(string(output))

	return sensorData, nil
}

func parseOutput(output string) []SensorData {
	lines := strings.Split(output, "\n")
	var sensorData []SensorData

	cpu := 0

	for _, line := range lines {
		// 忽略空行
		if line == "" {
			continue
		}

		fields := strings.Split(line, "|")

		if len(fields) != 5 {
			continue
		}

		temp, err := strconv.Atoi(strings.TrimSpace(strings.Replace(fields[4], "degrees C", "", 1)))
		if err != nil {
			// 温度转换失败，忽略当前条目
			fmt.Println(err)
			continue
		}

		label := strings.TrimSpace(fields[0])

		if label == "Temp" {
			cpu += 1
			label = fmt.Sprintf("CPU%d", cpu)
		}

		data := SensorData{
			Label:  label,
			ID:     strings.TrimSpace(fields[1]),
			Status: strings.TrimSpace(fields[2]),
			Value:  strings.TrimSpace(fields[3]),
			Temp:   temp,
		}
		sensorData = append(sensorData, data)
	}

	return sensorData
}
