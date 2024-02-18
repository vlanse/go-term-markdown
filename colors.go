package markdown

import (
	"fmt"
)

var (
	// we need a bunch of escape code for manual formatting
	boldOn = "\x1b[1m"
	// boldOff       = "\x1b[21m" --> use resetAll + snapshot with bold off instead
	italicOn      = "\x1b[3m"
	italicOff     = "\x1b[23m"
	crossedOutOn  = "\x1b[9m"
	crossedOutOff = "\x1b[29m"
	greenOn       = "\x1b[32m"

	//resetAll = "\x1b[0m"
	//colorOff = "\x1b[39m"

	resetAll = "[]"
	colorOff = "[]"

	//Green        = color.New(color.FgGreen).SprintFunc()
	//HiGreen      = color.New(color.FgHiGreen).SprintFunc()
	//GreenBold    = color.New(color.FgGreen, color.Bold).SprintFunc()
	//Blue = color.New(color.FgBlue).SprintFunc()
	//BlueBgItalic = color.New(color.BgBlue, color.Italic).SprintFunc()
	//Red          = color.New(color.FgRed).SprintFunc()

	Green = func(a ...interface{}) string {
		//return func(a ...interface{}) string {
		return fmt.Sprintf("[green]%s[]", fmt.Sprint(a...))
		//}
	}

	HiGreen = func(a ...interface{}) string {
		//return func(a ...interface{}) string {
		return fmt.Sprintf("[green]%s[]", fmt.Sprint(a...))
		//}
	}

	GreenBold = func(a ...interface{}) string {
		//return func(a ...interface{}) string {
		return fmt.Sprintf("[green]%s[]", fmt.Sprint(a...))
		//}
	}

	Blue = func(a ...interface{}) string {
		//return func(a ...interface{}) string {
		return fmt.Sprintf("[blue]%s[]", fmt.Sprint(a...))
		//}
	}

	BlueBgItalic = func(a ...interface{}) string {
		//return func(a ...interface{}) string {
		return fmt.Sprintf("[green]%s[]", fmt.Sprint(a...))
		//}
	}

	Red = func(a ...interface{}) string {
		//return func(a ...interface{}) string {
		return fmt.Sprintf("[red]%s[]", fmt.Sprint(a...))
		//}
	}
)
