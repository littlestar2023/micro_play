package util

import (
	"fmt"
	"os"
)

func GetHostServiceAddress(serviceIP, servicePort string) string {

	envIP := os.Getenv("POD_IP")
	if len(envIP) > 0 {
		return fmt.Sprintf("%v:%v", envIP, servicePort)
	} else {
		return fmt.Sprintf("%v:%v", serviceIP, servicePort)
	}
}
