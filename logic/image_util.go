package logic

import "golang.org/x/image/font"

func SplitByMeasureWidth(text string, maxWidth int, dr *font.Drawer) []string {
	var (
		lines []string
		line  string
	)
	for _, v := range text {
		vs := string(v)
		w := dr.MeasureString(line + vs).Round()
		switch {
		case maxWidth <= w:
			lines = append(lines, line)
			line = vs
		default:
			line = line + vs
		}
	}
	lines = append(lines, line)
	return lines
}
