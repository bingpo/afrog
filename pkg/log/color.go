package log

import (
	"strings"

	"github.com/fatih/color"
	"github.com/zan8in/afrog/pkg/utils"
)

func GetColor(level string, log string) string {
	var result string
	level = strings.ToLower(level)
	switch utils.SeverityMap[level] {
	case utils.INFO:
		result = color.BlueString(log)
	case utils.LOW:
		result = color.GreenString(log)
	case utils.MEDIUM:
		result = color.YellowString(log)
	case utils.HIGH:
		result = color.RedString(log)
	case utils.CRITICAL:
		result = color.RedString(log)
	default:
		result = log
	}
	return result
}
