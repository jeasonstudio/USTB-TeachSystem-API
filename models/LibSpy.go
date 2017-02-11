package models

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"strconv"
	"strings"
)

// GifToInt gif 转化为 int 输出
func GifToInt(source io.Reader) int {
	aim := GifToSlice(source)
	var r int
	for i := range aim {
		r = r + (int(math.Pow10(len(aim)-i-1)) * aim[i])
	}
	return r
}

// GifToString gif 转换成字符串输出
func GifToString(source io.Reader) string {
	aim := GifToSlice(source)
	var newa []string

	for i := 0; i < len(aim); i++ {
		newa = append(newa, strconv.Itoa(aim[i]))
	}
	return strings.Join(newa, "")
}

// GifToSlice gif 转换成 []int 的切片输出
func GifToSlice(source io.Reader) []int {
	var r []int

	theNum := 0
	for nn := 0; nn < 5; nn++ {
		tagArr := readGifCap(source)
		arrNew, zs, ze := splitArrZ(tagArr, theNum)
		if ze > zs {
			final := splitArrH(arrNew)
			if final[4][0] == 1 && final[5][0] == 1 && final[4][1] == 1 && final[5][1] == 1 && final[4][6] == 1 && final[4][7] == 1 && final[5][6] == 1 && final[5][7] == 1 {
				r = append(r, 0)
			}
			if len(final[0]) == 6 {
				r = append(r, 1)
			}
			if len(final[0]) > 6 && final[9][0] == 1 && final[9][1] == 1 && final[9][2] == 1 && final[9][3] == 1 && final[9][4] == 1 && final[9][5] == 1 && final[9][6] == 1 && final[9][7] == 1 {
				r = append(r, 2)
			}
			if len(final[0]) > 6 && final[0][0] == 0 && final[0][1] == 1 && final[0][2] == 1 {
				r = append(r, 3)
			}
			if len(final[0]) > 6 && final[9][5] == 1 && final[9][6] == 1 && final[9][7] == 0 {
				r = append(r, 4)
			}
			if len(final[0]) > 6 && final[0][0] == 1 && final[1][0] == 1 {
				r = append(r, 5)
			}
			if len(final[0]) > 6 && final[5][0] == 1 && final[5][1] == 1 && final[5][2] == 1 {
				r = append(r, 6)
			}
			if len(final[0]) > 6 && final[8][0] == 1 && final[9][0] == 1 && final[8][1] == 1 && final[9][1] == 1 {
				r = append(r, 7)
			}
			if len(final[0]) > 6 && final[0][2] == 1 && final[0][3] == 1 && final[0][4] == 1 && final[4][2] == 1 && final[4][3] == 1 && final[4][4] == 1 && final[9][2] == 1 && final[9][3] == 1 && final[9][4] == 1 {
				r = append(r, 8)
			}
			if len(final[0]) > 6 && final[0][2] == 1 && final[0][3] == 1 && final[0][4] == 1 && final[5][2] == 1 && final[5][3] == 1 && final[5][4] == 1 && final[9][2] == 1 && final[9][3] == 1 && final[9][4] == 1 {
				r = append(r, 9)
			}
			// printArr(splitArrH(arrNew))
			theNum = ze + 5
		} else {
			break
		}
	}

	// fmt.Println(r)
	return r
}

// 纵切切片
func splitArrZ(splitZArr [][]int, startLineSplit int) ([][]int, int, int) {
	if startLineSplit > len(splitZArr[0]) {
		fmt.Println("Error:Out of range")
		return nil, 0, 0
	}
	var startLine, endLine int = startLineSplit, startLineSplit
	var thisLine, lastLine bool = false, false

	for k := startLineSplit; k < len(splitZArr[0]); k++ {
		lastLine = thisLine
		for l := 0; l < len(splitZArr); l++ {
			if splitZArr[l][k] == 1 {
				thisLine = true
				break
			} else {
				thisLine = false
			}
		}
		if thisLine && !lastLine {
			startLine = k
		} else if !thisLine && lastLine {
			endLine = k
			break
		}
	}
	dbA := splitZArr
	for i := range splitZArr {
		dbA[i] = dbA[i][startLine:endLine]
	}
	return splitZArr, startLine, endLine
	// fmt.Println(splitZArr[20][startLine[0]:endLine[0]])
}

// 横切切片
func splitArrH(splitArr [][]int) [][]int {
	var startLine, endLine int
	var thisLine, lastLine bool = false, false
	for i := 0; i < len(splitArr); i++ {
		lastLine = thisLine
		for j := 0; j < len(splitArr[i]); j++ {
			if splitArr[i][j] == 1 {
				thisLine = true
				break
			} else {
				thisLine = false
			}
		}
		if thisLine == true && lastLine == false {
			startLine = i
		} else if thisLine == false && lastLine == true {
			endLine = i
		}
	}
	// fmt.Println(startLine, endLine)
	// printArr(splitArr[startLine:endLine])
	// if startLine == 0 && endLine == 0 {
	// 	err := new error("No number")
	// }

	return splitArr[startLine:endLine]
}

func readGifCap(sourceGif io.Reader) [][]int {
	// fmt.Println("start captcha")

	// file, err := os.Open(sourceGif)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// s, _ := io.Reader.Read(sourceGif)

	cap, _ := gif.Decode(sourceGif)
	new := image.NewGray16(cap.Bounds())

	xWidth := cap.Bounds().Dx()
	yHeight := cap.Bounds().Dy()

	// arr := make([][]int, yHeight, xWidth)
	var arr [][]int
	for i := 0; i < yHeight; i++ {
		var eachArr []int
		for j := 0; j < xWidth; j++ {

			if judgeRed(cap.At(j, i)) {
				eachArr = append(eachArr, 1)
				var s color.Gray16
				s.Y = 65535
				new.SetGray16(j, i, s)
			} else {
				eachArr = append(eachArr, 0)
			}
		}
		arr = append(arr, eachArr)
	}
	return arr
}

// RGBAToGray 转灰度
func RGBAToGray(color color.Color) uint16 {
	thisR, thisG, thisB, _ := color.RGBA()
	return uint16((thisR*299 + thisG*587 + thisB*114 + 500) / 1000)
}

// 判断是否是红色
func judgeRed(color color.Color) bool {
	thisR, thisG, thisB, _ := color.RGBA()
	if thisR > 22500 && thisG < 22500 && thisB < 22500 {
		return true
	} else {
		return false
	}
}

func printArr(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
