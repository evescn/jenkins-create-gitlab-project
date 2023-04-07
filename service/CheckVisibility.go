package service

import (
	"fmt"
	"os"
)

func CheckVisibility(visibility string) {
	if visibility != "private" && visibility != "internal" && visibility != "public" {
		fmt.Println("ERROR: 可见度级别关键字错误，只接收 private、internal、public 级别 ")
		os.Exit(-3)
	}
}
