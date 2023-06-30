package datetime

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func DateTime(year, month, day, hour, minute, second int) string {
	var moStr, dStr, hStr, mnStr, sStr string

	moStr = fmt.Sprintf("%02d", month)
	dStr = fmt.Sprintf("%02d", day)
	hStr = fmt.Sprintf("%02d", hour)
	mnStr = fmt.Sprintf("%02d", minute)
	sStr = fmt.Sprintf("%02d", second)

	return fmt.Sprintf(`%d-%s-%sT%s:%s:%s`, year, moStr, dStr, hStr, mnStr, sStr)
}

func DateTimeToTime(date string, loc *time.Location) (time.Time, error) {
	var data []int
	sample := []rune(date + "-")

	var cur string
	for i := 0; i < len(sample); i++ {
		c := string(sample[i])

		if c == "-" || c == "T" || c == ":" {
			val, err := strconv.Atoi(cur)
			if err != nil {
				return time.Now(), err
			} else {
				data = append(data, val)
			}
			cur = ""
		} else {
			cur += c
		}
	}

	switch len(data) {
	case 6:
		return time.Date(data[0], time.Month(data[1]), data[2], data[3], data[4], data[5], 0, loc), nil
	case 3:
		return time.Date(data[0], time.Month(data[1]), data[2], 0, 0, 0, 0, loc), nil
	default:
		return time.Now(), errors.New("Could not parse the data!")
	}
}
