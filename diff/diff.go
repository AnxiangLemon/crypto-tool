package diff

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func DiffText(text1, text2 string) {

	totalWidth := 120 // 控制台的总宽度（可调整以适应不同的控制台）
	differences := textDifference(text1, text2, totalWidth)
	separator := strings.Repeat("-", totalWidth)
	fmt.Println(separator)
	for _, diff := range differences {
		fmt.Println(diff)
	}
	fmt.Println(separator)
}

func formatHexWithSpaces(hexString string) []string {
	var formatted []string
	for i := 0; i < len(hexString); i += 32 { // 32 个字符表示 16 个字节
		end := i + 32
		if end > len(hexString) {
			end = len(hexString)
		}
		line := hexString[i:end]
		formattedLine := strings.Join(splitIntoPairs(line), " ")
		formatted = append(formatted, formattedLine)
	}
	return formatted
}

func splitIntoPairs(s string) []string {
	var pairs []string
	for i := 0; i < len(s); i += 2 {
		pairs = append(pairs, s[i:i+2])
	}
	return pairs
}

func colorizeDifference(line1, line2 string) (string, string, int) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	var result1, result2 string
	diffCount := 0
	maxLength := len(line1)
	if len(line2) > maxLength {
		maxLength = len(line2)
	}

	for i := 0; i < maxLength; i++ {
		var char1, char2 string
		if i < len(line1) {
			char1 = string(line1[i])
		} else {
			char1 = " "
		}
		if i < len(line2) {
			char2 = string(line2[i])
		} else {
			char2 = " "
		}

		if char1 != char2 {
			result1 += red(char1)
			result2 += green(char2)
			diffCount++
		} else {
			result1 += char1
			result2 += char2
		}
	}

	return result1, result2, diffCount
}

// RemoveNewlines 移除字符串中的所有换行符
func RemoveNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "") // 处理 \r\n
	s = strings.ReplaceAll(s, "\n", "")   // 处理 \n
	return strings.ToUpper(s)
}

func textDifference(text1, text2 string, totalWidth int) []string {
	// lines1 := strings.Split(text1, "\n")
	// lines2 := strings.Split(text2, "\n")
	text1 = RemoveNewlines(text1)
	text2 = RemoveNewlines(text2)
	lines1 := formatHexWithSpaces(text1)
	lines2 := formatHexWithSpaces(text2)

	var differences []string
	var totalDiffCount int
	maxLines := len(lines1)
	if len(lines2) > maxLines {
		maxLines = len(lines2)
	}

	maxLength := 0
	for i := 0; i < maxLines; i++ {
		var line1, line2 string
		if i < len(lines1) {
			line1 = lines1[i]
		}
		if i < len(lines2) {
			line2 = lines2[i]
		}

		if len(line1) > maxLength {
			maxLength = len(line1)
		}
		if len(line2) > maxLength {
			maxLength = len(line2)
		}
	}

	// combinedWidth := maxLength*2 + 3
	// if combinedWidth > totalWidth {
	// 	combinedWidth = totalWidth
	// }
	// paddingSize := (totalWidth - combinedWidth) / 2
	paddingSize := 5

	for i := 0; i < maxLines; i++ {
		var line1, line2 string
		if i < len(lines1) {
			line1 = lines1[i]
		} else {
			line1 = ""
		}
		if i < len(lines2) {
			line2 = lines2[i]
		} else {
			line2 = ""
		}

		paddedLine1 := fmt.Sprintf("%-*s", maxLength, line1)
		paddedLine2 := fmt.Sprintf("%-*s", maxLength, line2)

		coloredLine1, coloredLine2, diffCount := colorizeDifference(paddedLine1, paddedLine2)

		totalDiffCount += diffCount

		padding := strings.Repeat(" ", paddingSize)
		differences = append(differences, fmt.Sprintf("%s%s | %s", padding, coloredLine1, coloredLine2))
	}

	if totalDiffCount == 0 {
		differences = append(differences, "提示: 数据完全相同")
	}
	//  else {
	// 	differences = append(differences, fmt.Sprintf("提示: %d 处不同", totalDiffCount))
	// }

	return differences
}
