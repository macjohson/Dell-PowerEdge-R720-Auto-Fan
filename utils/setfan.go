package utils

import (
	"fmt"
	"os/exec"
)

func SetFan(percent int, account *Account) {
	a1 := account.Command([]string{"raw", "0x30", "0x30", "0x01", "0x00"})
	err := exec.Command(a1[0], a1[1:]...).Run()

	if err != nil {
		fmt.Println(err)
	}

	a2 := account.Command([]string{"raw", "0x30", "0x30", "0x02", "0xff", fmt.Sprintf("0x%X", percent)})
	err = exec.Command(a2[0], a2[1:]...).Run()

	if err != nil {
		fmt.Println(err)
	}
}

func SetManual(account *Account) {
	a1 := account.Command([]string{"raw", "raw", "0x30", "0x30", "0x01", "0x00"})
	err := exec.Command(a1[0], a1[1:]...).Run()

	if err != nil {
		fmt.Println(err)
	}

	a2 := account.Command([]string{"raw", "0x30", "0x30", "0x01", "0x01"})
	err = exec.Command(a2[0], a2[1:]...).Run()

	if err != nil {
		fmt.Println(err)
	}
}
