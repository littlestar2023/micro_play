package utils

import (
	"fmt"
)

func GetHostServiceAddress(serviceIP, servicePort, podIP string) string {

	if len(podIP) > 0 {
		return fmt.Sprintf("%v:%v", podIP, servicePort)
	} else {
		return fmt.Sprintf("%v:%v", serviceIP, servicePort)
	}
}
