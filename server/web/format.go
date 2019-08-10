package web

import (
	"fmt"
	"math"
	"mlp/catalog/sim"
	"strings"
	"time"
)

type Format struct {

}




func (f *Format) BytesDecimal(b int64) string{
	return avoidWrapping(bytesDecimal(b))
}




// avoidWrapping adds non-breaking spaces where there previously were normal spaces.
func avoidWrapping(s string) string{
	const nbsp = "\u00A0"
	return strings.Replace(s, " ", nbsp, -1)
}

func bytesDecimal(b int64) string {

	if b == 0{
		return "0 bytes"
	}
	if b == 1 {
		return "1 byte"
	}

	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d bytes", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func (f *Format) Timestamp(t int64) string{
	now := sim.UTC()
	stamp := time.Unix(t, 0)

	delta := now.Sub(stamp)

	secs := delta.Seconds()
	if secs <= 1 && secs >= -1 {
		return avoidWrapping("just now")
	}

	if secs > 0 {
		return avoidWrapping(durationAbs(delta) + " ago")
	}
	return avoidWrapping(durationAbs(stamp.Sub(now)))
}


func durationAbs(d time.Duration) string {

	days := int(math.Floor( d.Hours() / 24))


	switch {
	case days > 365 * 2:
		return fmt.Sprintf("%d years", days/365)
	case days >= 365:
		return "a year ago"
	case days > 60:
		return fmt.Sprintf("%d months", days/30)

	case days >= 30:
		return "a month"
	case days >= 14:
		return fmt.Sprintf("%d weeks", days/7)
	case days >= 7:
		return "a week"
	case days > 1:
		return fmt.Sprintf("%d days", days)
	case days == 1:
		return "a day"

	}


	hours := int(math.Floor(d.Hours()))
	sec := int(d.Seconds())

	switch {
	case hours > 1:
		return fmt.Sprintf("%d hours", hours)
	case hours ==1:
		return "an hour"
	case sec >= 120:
		return fmt.Sprintf("%d minutes", sec/60)
	case sec >= 60:
		return "a minute"
	case sec > 1:
		return fmt.Sprintf("%d seconds", sec)
	}
	return ""
}
