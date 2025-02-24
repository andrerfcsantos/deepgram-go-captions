package renderers

import "fmt"

func SecondsToTimestamp(seconds float64) string {
	hours := int(seconds / 3600)
	minutes := int((seconds - float64(hours*3600)) / 60)
	seconds = seconds - float64(hours*3600) - float64(minutes*60)
	return fmt.Sprintf("%02d:%02d:%02.3f", hours, minutes, seconds)
}
