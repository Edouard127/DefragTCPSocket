package utils

import (
	"fmt"
	"kamigen/socket/enums"
	"os"
	"strings"
	"time"
)

var ConsoleHistory []*string

func LogFile(con bool, logger enums.Level, message ...string) error {
	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		f, err := os.Create("log.txt")
		if err != nil {
			return err
		}
		defer func(f *os.File) error {
			err := f.Close()
			if err != nil {
				return err
			}
			return nil
		}(f)
	}
	if f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644); err != nil {
		return err
	} else {
		date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Local)
		defer func(f *os.File) error {
			err := f.Close()
			if err != nil {
				return err
			}
			return nil
		}(f)
		s := fmt.Sprintf("%s: [%s] %s\n", date.String(), logger, strings.Join(message[:], " "))
		if _, err := f.WriteString(s); err != nil {
			return err
		}
		if con == true {
			ConsoleHistory = append(ConsoleHistory, &s)
		}
		return nil
	}
}
