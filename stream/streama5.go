// package stream is responsible for solving a stream cipher
package main

var ciphertext2 = []int64{
	44, 225, 90, 111, 44, 209, 208, 235, 11, 91,
	63, 26, 41, 151, 131, 78, 146, 62, 63, 63,
	149, 181, 157, 103, 168, 218, 141, 203, 71,
	4, 61, 153, 180, 107, 52, 27, 121, 131, 26,
	161, 15, 31, 111, 70, 49, 1, 190, 178, 175,
	6, 179, 214, 32, 203, 232, 131, 211, 238, 60,
	99, 136, 206, 145, 174, 186, 212, 156, 113,
	153, 120, 40, 193, 84, 36, 181, 7, 196, 72,
	246, 157, 186, 127, 227, 34, 223, 48, 189,
	107, 94, 152, 164, 193, 67, 70, 38, 204, 135,
	4, 29, 164, 105, 129, 161, 11, 96, 239, 132,
	120, 71, 94, 134, 120, 225, 231, 250, 136, 12,
	201, 142, 201, 62, 12, 199, 34, 125, 82, 237,
	43, 8, 17, 28, 216, 225, 66, 233, 89, 152, 198,
	81, 139, 30, 110, 47, 178, 199, 245, 244, 77, 35,
	102, 32, 113, 236, 90, 76, 250, 239, 63, 175, 180,
	96, 33, 15, 98, 137, 23, 168, 10, 10, 125, 64, 42,
	21, 164, 163, 187, 1, 163, 208, 42, 214, 232, 135,
	213, 224, 57, 106, 141, 207, 142, 162, 190, 218,
	133, 106, 130, 116, 55, 208, 67, 47, 191, 4, 215,
	74, 228, 131, 172, 127, 229, 37, 215, 48, 178, 99,
	73, 133, 161, 206, 90, 70, 60, 222, 153, 14, 7, 170,
	121, 130, 167, 8, 97, 234, 143, 115,
}

func r51decipher() {
	tapped_bits1 := []int64{13, 16, 17, 18}
	tapped_bits2 := []int64{20, 21}
	tapped_bits3 := []int64{7, 20, 21, 22}
}

// func shift(c string, x string, howManytimes int) string {
// 	bt := 0
// 	c = c[0:7] + fmt.Sprint(1)
// 	for i := 0; i < howManytimes; i++ {
// 		bt = 0
// 		for i := 0; i < 8; i++ {
// 			bt += (int(x[i]) - 48) * (int(c[7-i]) - 48)
// 		}
// 		x = x[1:8] + fmt.Sprint(bt%2)
// 	}
//
// 	return x
// }
