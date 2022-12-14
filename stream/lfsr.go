// package stream is responsible for solving a stream cipher
package main

import (
	"fmt"
	"os"
	"strconv"
)

var ciphertext = []int64{
	3, 20, 22, 214, 229, 94, 245, 87, 151, 195, 93, 159, 0, 114, 50, 183,
	213, 255, 235, 80, 39, 102, 59, 115, 253, 87, 91, 251, 236, 47, 170, 188, 117, 52, 5,
	119, 159, 25, 171, 7, 0, 117, 82, 40, 10, 184, 188, 169, 13, 184, 213, 32, 218, 244,
	137, 213, 225, 43, 111, 144, 196, 149, 172, 181, 198, 148, 109, 129, 126, 53, 222, 89,
	45, 190, 6, 196, 86, 251, 134, 188, 96, 235, 56, 192, 39, 187, 113, 79, 130, 169, 223,
	79, 84, 59, 208, 150, 22, 29, 168, 120, 132, 160, 13, 103, 243, 136, 100, 95, 66, 134,
	108, 244, 239, 228, 135, 13, 196, 158, 205, 39, 6, 203, 39, 105, 78, 238, 51, 8, 14, 7,
	205, 248, 86, 246, 68, 153, 197, 77, 136, 28, 122, 47, 170, 211, 241, 230, 74, 51, 100,
	59, 99, 248, 77, 64, 227, 245, 36, 184, 179, 110, 52, 7, 96, 152, 31, 160, 2, 29, 117,
	95, 53, 9, 163, 182, 187, 30, 185, 213, 54, 204, 242, 130, 199, 255, 49, 114, 131, 209,
	142, 180, 167, 193, 148, 116, 132, 120, 41, 216, 94, 37, 181, 4, 194, 74, 246, 134, 185,
	98, 248, 58, 218, 43, 182, 105, 83, 130,
}

func decipher(a, b, c int64) {
	var currLetter string
	letterNumber := a
	file, _ := os.Create(fmt.Sprintf("%d%d%d", a, b, c))

	for i := 0; i < len(ciphertext); i++ {
		currLetter = shift(
			fmt.Sprintf("%08s", strconv.FormatInt(c, 2)),
			fmt.Sprintf("%08s", strconv.FormatInt(letterNumber, 2)),
			int(b),
		)
		letterNumber, _ = strconv.ParseInt(currLetter, 2, 16)
		file.WriteString(fmt.Sprintf("%s", string(rune(letterNumber^ciphertext[i]))))
	}
}

func shift(c string, x string, howManytimes int) string {
	bt := 0
	c = c[0:7] + fmt.Sprint(1)
	for i := 0; i < howManytimes; i++ {
		bt = 0
		for i := 0; i < 8; i++ {
			bt += (int(x[i]) - 48) * (int(c[7-i]) - 48)
		}
		x = x[1:8] + fmt.Sprint(bt%2)
	}

	return x
}

func toNumber(x string) (n int64) {
	n, _ = strconv.ParseInt(x, 2, 16)
	return
}
