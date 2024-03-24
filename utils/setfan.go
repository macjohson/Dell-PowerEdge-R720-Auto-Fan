package utils

import (
	"fmt"
	"os/exec"
)

func SetFan(percent int) {
	err := exec.Command("ipmitool", "raw", "0x30", "0x30", "0x01", "0x00").Run()

	if err != nil {
		fmt.Println(err)
	}

	err = exec.Command("ipmitool", "raw", "0x30", "0x30", "0x02", "0xff", fmt.Sprintf("0x%X", percent)).Run()

	if err != nil {
		fmt.Println(err)
	}
}
