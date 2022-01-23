/*
 * Author: Faheem Kha
 * Date: 22-01-2021
 * Description: ANSI codes for colors & formatting
 */

package main

type Color string
type Format string

// Formatting
const (
	FormatBold      Format = "\u001b[1m"
	FormatUnderline        = "\u001b[4m"
	FormatReversed         = "\u001b[7m"
)

// 8bit Colors
const (
	ColorReset  Color = "\u001b[0m"
	ColorBlack        = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorPurple       = "\u001b[35m"
	ColorCyan         = "\u001b[36m"
	ColorWhite        = "\u001b[37m"
)
