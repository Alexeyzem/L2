package mytime

import (
	"fmt"

	"github.com/beevik/ntp"
	"github.com/sirupsen/logrus"
)

func WriteTime(path string) {
	t, err := ntp.Time(path)
	if err != nil {
		logrus.Fatalf("Internal error: %v", err)
	}
	fmt.Println(t)
}
