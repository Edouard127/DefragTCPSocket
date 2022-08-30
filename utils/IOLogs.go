package utils

import (
	"fmt"
	"os"
	"time"
)

func LogFile(con bool, message ...string) error {
	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		f, err := os.Create("log.txt")
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
	}
	if f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644); err != nil {
		return err
	} else {
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
		for _, v := range message {
			date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
			if s, err := f.WriteString(date.String() + ": " + v + " "); err != nil {
				return err
			} else if con {
				fmt.Println(s)
			}
		}
		return nil
	}
}
