// Code generated from PainlessParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PainlessParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 86, 512,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 3, 2, 7, 2, 72, 10, 2, 12, 2, 14, 2, 75, 11, 2, 3,
	2, 7, 2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2, 3, 2, 5, 2, 84, 10, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 100, 10, 4, 12, 4, 14, 4, 103, 11, 4, 5, 4, 105, 10, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 113, 10, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 123, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 131, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 136, 10, 6, 3, 6,
	3, 6, 5, 6, 140, 10, 6, 3, 6, 3, 6, 5, 6, 144, 10, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 149, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6,
	171, 10, 6, 13, 6, 14, 6, 172, 5, 6, 175, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	192, 10, 7, 3, 8, 3, 8, 5, 8, 196, 10, 8, 3, 9, 3, 9, 7, 9, 200, 10, 9,
	12, 9, 14, 9, 203, 11, 9, 3, 9, 5, 9, 206, 10, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 5, 11, 214, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 7, 13, 222, 10, 13, 12, 13, 14, 13, 225, 11, 13, 3, 14, 3, 14, 3,
	14, 7, 14, 230, 10, 14, 12, 14, 14, 14, 233, 11, 14, 3, 15, 3, 15, 3, 15,
	5, 15, 238, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 298,
	10, 17, 12, 17, 14, 17, 301, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 316, 10,
	18, 3, 19, 3, 19, 7, 19, 320, 10, 19, 12, 19, 14, 19, 323, 11, 19, 3, 19,
	3, 19, 3, 19, 7, 19, 328, 10, 19, 12, 19, 14, 19, 331, 11, 19, 3, 19, 3,
	19, 3, 19, 7, 19, 336, 10, 19, 12, 19, 14, 19, 339, 11, 19, 5, 19, 341,
	10, 19, 3, 19, 3, 19, 7, 19, 345, 10, 19, 12, 19, 14, 19, 348, 11, 19,
	5, 19, 350, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	5, 20, 370, 10, 20, 3, 21, 3, 21, 3, 21, 5, 21, 375, 10, 21, 3, 22, 3,
	22, 5, 22, 379, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 6,
	26, 398, 10, 26, 13, 26, 14, 26, 399, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 7, 27, 410, 10, 27, 12, 27, 14, 27, 413, 11, 27, 5,
	27, 415, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 423,
	10, 28, 12, 28, 14, 28, 426, 11, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28,
	432, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 438, 10, 29, 12, 29, 14,
	29, 441, 11, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 448, 10, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 458, 10,
	31, 12, 31, 14, 31, 461, 11, 31, 5, 31, 463, 10, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 5, 32, 470, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7,
	33, 477, 10, 33, 12, 33, 14, 33, 480, 11, 33, 5, 33, 482, 10, 33, 3, 33,
	5, 33, 485, 10, 33, 3, 33, 3, 33, 3, 33, 5, 33, 490, 10, 33, 3, 34, 5,
	34, 493, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 510, 10, 35, 3,
	35, 2, 3, 32, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 2, 14, 3, 2, 32, 34, 3, 2, 35, 36, 3, 2, 58, 59, 3, 2, 37, 39, 3, 2,
	40, 43, 3, 2, 44, 47, 3, 2, 62, 73, 3, 2, 60, 61, 4, 2, 30, 31, 35, 36,
	3, 2, 74, 77, 3, 2, 11, 12, 3, 2, 85, 86, 2, 568, 2, 73, 3, 2, 2, 2, 4,
	87, 3, 2, 2, 2, 6, 92, 3, 2, 2, 2, 8, 112, 3, 2, 2, 2, 10, 174, 3, 2, 2,
	2, 12, 191, 3, 2, 2, 2, 14, 195, 3, 2, 2, 2, 16, 197, 3, 2, 2, 2, 18, 209,
	3, 2, 2, 2, 20, 213, 3, 2, 2, 2, 22, 215, 3, 2, 2, 2, 24, 217, 3, 2, 2,
	2, 26, 226, 3, 2, 2, 2, 28, 234, 3, 2, 2, 2, 30, 239, 3, 2, 2, 2, 32, 246,
	3, 2, 2, 2, 34, 315, 3, 2, 2, 2, 36, 349, 3, 2, 2, 2, 38, 369, 3, 2, 2,
	2, 40, 374, 3, 2, 2, 2, 42, 378, 3, 2, 2, 2, 44, 380, 3, 2, 2, 2, 46, 384,
	3, 2, 2, 2, 48, 387, 3, 2, 2, 2, 50, 391, 3, 2, 2, 2, 52, 401, 3, 2, 2,
	2, 54, 431, 3, 2, 2, 2, 56, 447, 3, 2, 2, 2, 58, 449, 3, 2, 2, 2, 60, 453,
	3, 2, 2, 2, 62, 469, 3, 2, 2, 2, 64, 484, 3, 2, 2, 2, 66, 492, 3, 2, 2,
	2, 68, 509, 3, 2, 2, 2, 70, 72, 5, 4, 3, 2, 71, 70, 3, 2, 2, 2, 72, 75,
	3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 79, 3, 2, 2, 2,
	75, 73, 3, 2, 2, 2, 76, 78, 5, 8, 5, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3,
	2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81,
	79, 3, 2, 2, 2, 82, 84, 5, 12, 7, 2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2,
	2, 2, 84, 85, 3, 2, 2, 2, 85, 86, 7, 2, 2, 3, 86, 3, 3, 2, 2, 2, 87, 88,
	5, 26, 14, 2, 88, 89, 7, 84, 2, 2, 89, 90, 5, 6, 4, 2, 90, 91, 5, 16, 9,
	2, 91, 5, 3, 2, 2, 2, 92, 104, 7, 9, 2, 2, 93, 94, 5, 26, 14, 2, 94, 101,
	7, 84, 2, 2, 95, 96, 7, 13, 2, 2, 96, 97, 5, 26, 14, 2, 97, 98, 7, 84,
	2, 2, 98, 100, 3, 2, 2, 2, 99, 95, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101,
	99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 105, 3, 2, 2, 2, 103, 101, 3,
	2, 2, 2, 104, 93, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106, 3, 2, 2,
	2, 106, 107, 7, 10, 2, 2, 107, 7, 3, 2, 2, 2, 108, 113, 5, 10, 6, 2, 109,
	110, 5, 12, 7, 2, 110, 111, 7, 14, 2, 2, 111, 113, 3, 2, 2, 2, 112, 108,
	3, 2, 2, 2, 112, 109, 3, 2, 2, 2, 113, 9, 3, 2, 2, 2, 114, 115, 7, 15,
	2, 2, 115, 116, 7, 9, 2, 2, 116, 117, 5, 32, 17, 2, 117, 118, 7, 10, 2,
	2, 118, 122, 5, 14, 8, 2, 119, 120, 7, 17, 2, 2, 120, 123, 5, 14, 8, 2,
	121, 123, 6, 6, 2, 2, 122, 119, 3, 2, 2, 2, 122, 121, 3, 2, 2, 2, 123,
	175, 3, 2, 2, 2, 124, 125, 7, 18, 2, 2, 125, 126, 7, 9, 2, 2, 126, 127,
	5, 32, 17, 2, 127, 130, 7, 10, 2, 2, 128, 131, 5, 14, 8, 2, 129, 131, 5,
	18, 10, 2, 130, 128, 3, 2, 2, 2, 130, 129, 3, 2, 2, 2, 131, 175, 3, 2,
	2, 2, 132, 133, 7, 20, 2, 2, 133, 135, 7, 9, 2, 2, 134, 136, 5, 20, 11,
	2, 135, 134, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137,
	139, 7, 14, 2, 2, 138, 140, 5, 32, 17, 2, 139, 138, 3, 2, 2, 2, 139, 140,
	3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 7, 14, 2, 2, 142, 144, 5, 22,
	12, 2, 143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2,
	145, 148, 7, 10, 2, 2, 146, 149, 5, 14, 8, 2, 147, 149, 5, 18, 10, 2, 148,
	146, 3, 2, 2, 2, 148, 147, 3, 2, 2, 2, 149, 175, 3, 2, 2, 2, 150, 151,
	7, 20, 2, 2, 151, 152, 7, 9, 2, 2, 152, 153, 5, 26, 14, 2, 153, 154, 7,
	84, 2, 2, 154, 155, 7, 54, 2, 2, 155, 156, 5, 32, 17, 2, 156, 157, 7, 10,
	2, 2, 157, 158, 5, 14, 8, 2, 158, 175, 3, 2, 2, 2, 159, 160, 7, 20, 2,
	2, 160, 161, 7, 9, 2, 2, 161, 162, 7, 84, 2, 2, 162, 163, 7, 16, 2, 2,
	163, 164, 5, 32, 17, 2, 164, 165, 7, 10, 2, 2, 165, 166, 5, 14, 8, 2, 166,
	175, 3, 2, 2, 2, 167, 168, 7, 25, 2, 2, 168, 170, 5, 16, 9, 2, 169, 171,
	5, 30, 16, 2, 170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3,
	2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 175, 3, 2, 2, 2, 174, 114, 3, 2, 2,
	2, 174, 124, 3, 2, 2, 2, 174, 132, 3, 2, 2, 2, 174, 150, 3, 2, 2, 2, 174,
	159, 3, 2, 2, 2, 174, 167, 3, 2, 2, 2, 175, 11, 3, 2, 2, 2, 176, 177, 7,
	19, 2, 2, 177, 178, 5, 16, 9, 2, 178, 179, 7, 18, 2, 2, 179, 180, 7, 9,
	2, 2, 180, 181, 5, 32, 17, 2, 181, 182, 7, 10, 2, 2, 182, 192, 3, 2, 2,
	2, 183, 192, 5, 24, 13, 2, 184, 192, 7, 21, 2, 2, 185, 192, 7, 22, 2, 2,
	186, 187, 7, 23, 2, 2, 187, 192, 5, 32, 17, 2, 188, 189, 7, 27, 2, 2, 189,
	192, 5, 32, 17, 2, 190, 192, 5, 32, 17, 2, 191, 176, 3, 2, 2, 2, 191, 183,
	3, 2, 2, 2, 191, 184, 3, 2, 2, 2, 191, 185, 3, 2, 2, 2, 191, 186, 3, 2,
	2, 2, 191, 188, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 13, 3, 2, 2, 2,
	193, 196, 5, 16, 9, 2, 194, 196, 5, 8, 5, 2, 195, 193, 3, 2, 2, 2, 195,
	194, 3, 2, 2, 2, 196, 15, 3, 2, 2, 2, 197, 201, 7, 5, 2, 2, 198, 200, 5,
	8, 5, 2, 199, 198, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199, 3, 2, 2,
	2, 201, 202, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 204,
	206, 5, 12, 7, 2, 205, 204, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207,
	3, 2, 2, 2, 207, 208, 7, 6, 2, 2, 208, 17, 3, 2, 2, 2, 209, 210, 7, 14,
	2, 2, 210, 19, 3, 2, 2, 2, 211, 214, 5, 24, 13, 2, 212, 214, 5, 32, 17,
	2, 213, 211, 3, 2, 2, 2, 213, 212, 3, 2, 2, 2, 214, 21, 3, 2, 2, 2, 215,
	216, 5, 32, 17, 2, 216, 23, 3, 2, 2, 2, 217, 218, 5, 26, 14, 2, 218, 223,
	5, 28, 15, 2, 219, 220, 7, 13, 2, 2, 220, 222, 5, 28, 15, 2, 221, 219,
	3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2,
	2, 2, 224, 25, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 226, 231, 7, 83, 2, 2,
	227, 228, 7, 7, 2, 2, 228, 230, 7, 8, 2, 2, 229, 227, 3, 2, 2, 2, 230,
	233, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 27, 3,
	2, 2, 2, 233, 231, 3, 2, 2, 2, 234, 237, 7, 84, 2, 2, 235, 236, 7, 62,
	2, 2, 236, 238, 5, 32, 17, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2,
	2, 238, 29, 3, 2, 2, 2, 239, 240, 7, 26, 2, 2, 240, 241, 7, 9, 2, 2, 241,
	242, 7, 83, 2, 2, 242, 243, 7, 84, 2, 2, 243, 244, 7, 10, 2, 2, 244, 245,
	5, 16, 9, 2, 245, 31, 3, 2, 2, 2, 246, 247, 8, 17, 1, 2, 247, 248, 5, 34,
	18, 2, 248, 299, 3, 2, 2, 2, 249, 250, 12, 17, 2, 2, 250, 251, 9, 2, 2,
	2, 251, 298, 5, 32, 17, 18, 252, 253, 12, 16, 2, 2, 253, 254, 9, 3, 2,
	2, 254, 298, 5, 32, 17, 17, 255, 256, 12, 15, 2, 2, 256, 257, 9, 4, 2,
	2, 257, 298, 5, 32, 17, 16, 258, 259, 12, 14, 2, 2, 259, 260, 9, 5, 2,
	2, 260, 298, 5, 32, 17, 15, 261, 262, 12, 13, 2, 2, 262, 263, 9, 6, 2,
	2, 263, 298, 5, 32, 17, 14, 264, 265, 12, 11, 2, 2, 265, 266, 9, 7, 2,
	2, 266, 298, 5, 32, 17, 12, 267, 268, 12, 10, 2, 2, 268, 269, 7, 48, 2,
	2, 269, 298, 5, 32, 17, 11, 270, 271, 12, 9, 2, 2, 271, 272, 7, 49, 2,
	2, 272, 298, 5, 32, 17, 10, 273, 274, 12, 8, 2, 2, 274, 275, 7, 50, 2,
	2, 275, 298, 5, 32, 17, 9, 276, 277, 12, 7, 2, 2, 277, 278, 7, 51, 2, 2,
	278, 298, 5, 32, 17, 8, 279, 280, 12, 6, 2, 2, 280, 281, 7, 52, 2, 2, 281,
	298, 5, 32, 17, 7, 282, 283, 12, 5, 2, 2, 283, 284, 7, 53, 2, 2, 284, 285,
	5, 32, 17, 2, 285, 286, 7, 54, 2, 2, 286, 287, 5, 32, 17, 5, 287, 298,
	3, 2, 2, 2, 288, 289, 12, 4, 2, 2, 289, 290, 7, 55, 2, 2, 290, 298, 5,
	32, 17, 4, 291, 292, 12, 3, 2, 2, 292, 293, 9, 8, 2, 2, 293, 298, 5, 32,
	17, 3, 294, 295, 12, 12, 2, 2, 295, 296, 7, 29, 2, 2, 296, 298, 5, 26,
	14, 2, 297, 249, 3, 2, 2, 2, 297, 252, 3, 2, 2, 2, 297, 255, 3, 2, 2, 2,
	297, 258, 3, 2, 2, 2, 297, 261, 3, 2, 2, 2, 297, 264, 3, 2, 2, 2, 297,
	267, 3, 2, 2, 2, 297, 270, 3, 2, 2, 2, 297, 273, 3, 2, 2, 2, 297, 276,
	3, 2, 2, 2, 297, 279, 3, 2, 2, 2, 297, 282, 3, 2, 2, 2, 297, 288, 3, 2,
	2, 2, 297, 291, 3, 2, 2, 2, 297, 294, 3, 2, 2, 2, 298, 301, 3, 2, 2, 2,
	299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 33, 3, 2, 2, 2, 301, 299,
	3, 2, 2, 2, 302, 303, 9, 9, 2, 2, 303, 316, 5, 36, 19, 2, 304, 305, 5,
	36, 19, 2, 305, 306, 9, 9, 2, 2, 306, 316, 3, 2, 2, 2, 307, 316, 5, 36,
	19, 2, 308, 309, 9, 10, 2, 2, 309, 316, 5, 34, 18, 2, 310, 311, 7, 9, 2,
	2, 311, 312, 5, 26, 14, 2, 312, 313, 7, 10, 2, 2, 313, 314, 5, 34, 18,
	2, 314, 316, 3, 2, 2, 2, 315, 302, 3, 2, 2, 2, 315, 304, 3, 2, 2, 2, 315,
	307, 3, 2, 2, 2, 315, 308, 3, 2, 2, 2, 315, 310, 3, 2, 2, 2, 316, 35, 3,
	2, 2, 2, 317, 321, 5, 38, 20, 2, 318, 320, 5, 40, 21, 2, 319, 318, 3, 2,
	2, 2, 320, 323, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2,
	322, 350, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 324, 325, 5, 26, 14, 2, 325,
	329, 5, 42, 22, 2, 326, 328, 5, 40, 21, 2, 327, 326, 3, 2, 2, 2, 328, 331,
	3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 350, 3, 2,
	2, 2, 331, 329, 3, 2, 2, 2, 332, 340, 5, 50, 26, 2, 333, 337, 5, 42, 22,
	2, 334, 336, 5, 40, 21, 2, 335, 334, 3, 2, 2, 2, 336, 339, 3, 2, 2, 2,
	337, 335, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339,
	337, 3, 2, 2, 2, 340, 333, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 350,
	3, 2, 2, 2, 342, 346, 5, 52, 27, 2, 343, 345, 5, 40, 21, 2, 344, 343, 3,
	2, 2, 2, 345, 348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2, 2,
	2, 347, 350, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 317, 3, 2, 2, 2, 349,
	324, 3, 2, 2, 2, 349, 332, 3, 2, 2, 2, 349, 342, 3, 2, 2, 2, 350, 37, 3,
	2, 2, 2, 351, 352, 7, 9, 2, 2, 352, 353, 5, 32, 17, 2, 353, 354, 7, 10,
	2, 2, 354, 370, 3, 2, 2, 2, 355, 370, 9, 11, 2, 2, 356, 370, 7, 80, 2,
	2, 357, 370, 7, 81, 2, 2, 358, 370, 7, 82, 2, 2, 359, 370, 7, 78, 2, 2,
	360, 370, 7, 79, 2, 2, 361, 370, 5, 54, 28, 2, 362, 370, 5, 56, 29, 2,
	363, 370, 7, 84, 2, 2, 364, 365, 7, 84, 2, 2, 365, 370, 5, 60, 31, 2, 366,
	367, 7, 24, 2, 2, 367, 368, 7, 83, 2, 2, 368, 370, 5, 60, 31, 2, 369, 351,
	3, 2, 2, 2, 369, 355, 3, 2, 2, 2, 369, 356, 3, 2, 2, 2, 369, 357, 3, 2,
	2, 2, 369, 358, 3, 2, 2, 2, 369, 359, 3, 2, 2, 2, 369, 360, 3, 2, 2, 2,
	369, 361, 3, 2, 2, 2, 369, 362, 3, 2, 2, 2, 369, 363, 3, 2, 2, 2, 369,
	364, 3, 2, 2, 2, 369, 366, 3, 2, 2, 2, 370, 39, 3, 2, 2, 2, 371, 375, 5,
	44, 23, 2, 372, 375, 5, 46, 24, 2, 373, 375, 5, 48, 25, 2, 374, 371, 3,
	2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 373, 3, 2, 2, 2, 375, 41, 3, 2, 2,
	2, 376, 379, 5, 44, 23, 2, 377, 379, 5, 46, 24, 2, 378, 376, 3, 2, 2, 2,
	378, 377, 3, 2, 2, 2, 379, 43, 3, 2, 2, 2, 380, 381, 9, 12, 2, 2, 381,
	382, 7, 86, 2, 2, 382, 383, 5, 60, 31, 2, 383, 45, 3, 2, 2, 2, 384, 385,
	9, 12, 2, 2, 385, 386, 9, 13, 2, 2, 386, 47, 3, 2, 2, 2, 387, 388, 7, 7,
	2, 2, 388, 389, 5, 32, 17, 2, 389, 390, 7, 8, 2, 2, 390, 49, 3, 2, 2, 2,
	391, 392, 7, 24, 2, 2, 392, 397, 7, 83, 2, 2, 393, 394, 7, 7, 2, 2, 394,
	395, 5, 32, 17, 2, 395, 396, 7, 8, 2, 2, 396, 398, 3, 2, 2, 2, 397, 393,
	3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 399, 400, 3, 2,
	2, 2, 400, 51, 3, 2, 2, 2, 401, 402, 7, 24, 2, 2, 402, 403, 7, 83, 2, 2,
	403, 404, 7, 7, 2, 2, 404, 405, 7, 8, 2, 2, 405, 414, 7, 5, 2, 2, 406,
	411, 5, 32, 17, 2, 407, 408, 7, 13, 2, 2, 408, 410, 5, 32, 17, 2, 409,
	407, 3, 2, 2, 2, 410, 413, 3, 2, 2, 2, 411, 409, 3, 2, 2, 2, 411, 412,
	3, 2, 2, 2, 412, 415, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2, 414, 406, 3, 2,
	2, 2, 414, 415, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 7, 6, 2, 2,
	417, 53, 3, 2, 2, 2, 418, 419, 7, 7, 2, 2, 419, 424, 5, 32, 17, 2, 420,
	421, 7, 13, 2, 2, 421, 423, 5, 32, 17, 2, 422, 420, 3, 2, 2, 2, 423, 426,
	3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 427, 3, 2,
	2, 2, 426, 424, 3, 2, 2, 2, 427, 428, 7, 8, 2, 2, 428, 432, 3, 2, 2, 2,
	429, 430, 7, 7, 2, 2, 430, 432, 7, 8, 2, 2, 431, 418, 3, 2, 2, 2, 431,
	429, 3, 2, 2, 2, 432, 55, 3, 2, 2, 2, 433, 434, 7, 7, 2, 2, 434, 439, 5,
	58, 30, 2, 435, 436, 7, 13, 2, 2, 436, 438, 5, 58, 30, 2, 437, 435, 3,
	2, 2, 2, 438, 441, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 439, 440, 3, 2, 2,
	2, 440, 442, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 442, 443, 7, 8, 2, 2, 443,
	448, 3, 2, 2, 2, 444, 445, 7, 7, 2, 2, 445, 446, 7, 54, 2, 2, 446, 448,
	7, 8, 2, 2, 447, 433, 3, 2, 2, 2, 447, 444, 3, 2, 2, 2, 448, 57, 3, 2,
	2, 2, 449, 450, 5, 32, 17, 2, 450, 451, 7, 54, 2, 2, 451, 452, 5, 32, 17,
	2, 452, 59, 3, 2, 2, 2, 453, 462, 7, 9, 2, 2, 454, 459, 5, 62, 32, 2, 455,
	456, 7, 13, 2, 2, 456, 458, 5, 62, 32, 2, 457, 455, 3, 2, 2, 2, 458, 461,
	3, 2, 2, 2, 459, 457, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 463, 3, 2,
	2, 2, 461, 459, 3, 2, 2, 2, 462, 454, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2,
	463, 464, 3, 2, 2, 2, 464, 465, 7, 10, 2, 2, 465, 61, 3, 2, 2, 2, 466,
	470, 5, 32, 17, 2, 467, 470, 5, 64, 33, 2, 468, 470, 5, 68, 35, 2, 469,
	466, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2, 469, 468, 3, 2, 2, 2, 470, 63, 3,
	2, 2, 2, 471, 485, 5, 66, 34, 2, 472, 481, 7, 9, 2, 2, 473, 478, 5, 66,
	34, 2, 474, 475, 7, 13, 2, 2, 475, 477, 5, 66, 34, 2, 476, 474, 3, 2, 2,
	2, 477, 480, 3, 2, 2, 2, 478, 476, 3, 2, 2, 2, 478, 479, 3, 2, 2, 2, 479,
	482, 3, 2, 2, 2, 480, 478, 3, 2, 2, 2, 481, 473, 3, 2, 2, 2, 481, 482,
	3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 485, 7, 10, 2, 2, 484, 471, 3, 2,
	2, 2, 484, 472, 3, 2, 2, 2, 485, 486, 3, 2, 2, 2, 486, 489, 7, 57, 2, 2,
	487, 490, 5, 16, 9, 2, 488, 490, 5, 32, 17, 2, 489, 487, 3, 2, 2, 2, 489,
	488, 3, 2, 2, 2, 490, 65, 3, 2, 2, 2, 491, 493, 5, 26, 14, 2, 492, 491,
	3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 495, 7, 84,
	2, 2, 495, 67, 3, 2, 2, 2, 496, 497, 7, 83, 2, 2, 497, 498, 7, 56, 2, 2,
	498, 510, 7, 84, 2, 2, 499, 500, 5, 26, 14, 2, 500, 501, 7, 56, 2, 2, 501,
	502, 7, 24, 2, 2, 502, 510, 3, 2, 2, 2, 503, 504, 7, 84, 2, 2, 504, 505,
	7, 56, 2, 2, 505, 510, 7, 84, 2, 2, 506, 507, 7, 28, 2, 2, 507, 508, 7,
	56, 2, 2, 508, 510, 7, 84, 2, 2, 509, 496, 3, 2, 2, 2, 509, 499, 3, 2,
	2, 2, 509, 503, 3, 2, 2, 2, 509, 506, 3, 2, 2, 2, 510, 69, 3, 2, 2, 2,
	52, 73, 79, 83, 101, 104, 112, 122, 130, 135, 139, 143, 148, 172, 174,
	191, 195, 201, 205, 213, 223, 231, 237, 297, 299, 315, 321, 329, 337, 340,
	346, 349, 369, 374, 378, 399, 411, 414, 424, 431, 439, 447, 459, 462, 469,
	478, 481, 484, 489, 492, 509,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'{'", "'}'", "'['", "']'", "'('", "')'", "'.'", "'?.'", "','",
	"';'", "'if'", "'in'", "'else'", "'while'", "'do'", "'for'", "'continue'",
	"'break'", "'return'", "'new'", "'try'", "'catch'", "'throw'", "'this'",
	"'instanceof'", "'!'", "'~'", "'*'", "'/'", "'%'", "'+'", "'-'", "'<<'",
	"'>>'", "'>>>'", "'<'", "'<='", "'>'", "'>='", "'=='", "'==='", "'!='",
	"'!=='", "'&'", "'^'", "'|'", "'&&'", "'||'", "'?'", "':'", "'?:'", "'::'",
	"'->'", "'=~'", "'==~'", "'++'", "'--'", "'='", "'+='", "'-='", "'*='",
	"'/='", "'%='", "'&='", "'^='", "'|='", "'<<='", "'>>='", "'>>>='", "",
	"", "", "", "", "", "'true'", "'false'", "'null'",
}
var symbolicNames = []string{
	"", "WS", "COMMENT", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "LP", "RP",
	"DOT", "NSDOT", "COMMA", "SEMICOLON", "IF", "IN", "ELSE", "WHILE", "DO",
	"FOR", "CONTINUE", "BREAK", "RETURN", "NEW", "TRY", "CATCH", "THROW", "THIS",
	"INSTANCEOF", "BOOLNOT", "BWNOT", "MUL", "DIV", "REM", "ADD", "SUB", "LSH",
	"RSH", "USH", "LT", "LTE", "GT", "GTE", "EQ", "EQR", "NE", "NER", "BWAND",
	"XOR", "BWOR", "BOOLAND", "BOOLOR", "COND", "COLON", "ELVIS", "REF", "ARROW",
	"FIND", "MATCH", "INCR", "DECR", "ASSIGN", "AADD", "ASUB", "AMUL", "ADIV",
	"AREM", "AAND", "AXOR", "AOR", "ALSH", "ARSH", "AUSH", "OCTAL", "HEX",
	"INTEGER", "DECIMAL", "STRING", "REGEX", "TRUE", "FALSE", "NULL", "TYPE",
	"ID", "DOTINTEGER", "DOTID",
}

var ruleNames = []string{
	"source", "function", "parameters", "statement", "rstatement", "dstatement",
	"trailer", "block", "empty", "initializer", "afterthought", "declaration",
	"decltype", "declvar", "trap", "expression", "unary", "chain", "primary",
	"postfix", "postdot", "callinvoke", "fieldaccess", "braceaccess", "stdarray",
	"initializedarray", "listinitializer", "mapinitializer", "maptoken", "arguments",
	"argument", "lambda", "lamtype", "funcref",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PainlessParser struct {
	*antlr.BaseParser
}

func NewPainlessParser(input antlr.TokenStream) *PainlessParser {
	this := new(PainlessParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PainlessParser.g4"

	return this
}

// PainlessParser tokens.
const (
	PainlessParserEOF        = antlr.TokenEOF
	PainlessParserWS         = 1
	PainlessParserCOMMENT    = 2
	PainlessParserLBRACK     = 3
	PainlessParserRBRACK     = 4
	PainlessParserLBRACE     = 5
	PainlessParserRBRACE     = 6
	PainlessParserLP         = 7
	PainlessParserRP         = 8
	PainlessParserDOT        = 9
	PainlessParserNSDOT      = 10
	PainlessParserCOMMA      = 11
	PainlessParserSEMICOLON  = 12
	PainlessParserIF         = 13
	PainlessParserIN         = 14
	PainlessParserELSE       = 15
	PainlessParserWHILE      = 16
	PainlessParserDO         = 17
	PainlessParserFOR        = 18
	PainlessParserCONTINUE   = 19
	PainlessParserBREAK      = 20
	PainlessParserRETURN     = 21
	PainlessParserNEW        = 22
	PainlessParserTRY        = 23
	PainlessParserCATCH      = 24
	PainlessParserTHROW      = 25
	PainlessParserTHIS       = 26
	PainlessParserINSTANCEOF = 27
	PainlessParserBOOLNOT    = 28
	PainlessParserBWNOT      = 29
	PainlessParserMUL        = 30
	PainlessParserDIV        = 31
	PainlessParserREM        = 32
	PainlessParserADD        = 33
	PainlessParserSUB        = 34
	PainlessParserLSH        = 35
	PainlessParserRSH        = 36
	PainlessParserUSH        = 37
	PainlessParserLT         = 38
	PainlessParserLTE        = 39
	PainlessParserGT         = 40
	PainlessParserGTE        = 41
	PainlessParserEQ         = 42
	PainlessParserEQR        = 43
	PainlessParserNE         = 44
	PainlessParserNER        = 45
	PainlessParserBWAND      = 46
	PainlessParserXOR        = 47
	PainlessParserBWOR       = 48
	PainlessParserBOOLAND    = 49
	PainlessParserBOOLOR     = 50
	PainlessParserCOND       = 51
	PainlessParserCOLON      = 52
	PainlessParserELVIS      = 53
	PainlessParserREF        = 54
	PainlessParserARROW      = 55
	PainlessParserFIND       = 56
	PainlessParserMATCH      = 57
	PainlessParserINCR       = 58
	PainlessParserDECR       = 59
	PainlessParserASSIGN     = 60
	PainlessParserAADD       = 61
	PainlessParserASUB       = 62
	PainlessParserAMUL       = 63
	PainlessParserADIV       = 64
	PainlessParserAREM       = 65
	PainlessParserAAND       = 66
	PainlessParserAXOR       = 67
	PainlessParserAOR        = 68
	PainlessParserALSH       = 69
	PainlessParserARSH       = 70
	PainlessParserAUSH       = 71
	PainlessParserOCTAL      = 72
	PainlessParserHEX        = 73
	PainlessParserINTEGER    = 74
	PainlessParserDECIMAL    = 75
	PainlessParserSTRING     = 76
	PainlessParserREGEX      = 77
	PainlessParserTRUE       = 78
	PainlessParserFALSE      = 79
	PainlessParserNULL       = 80
	PainlessParserTYPE       = 81
	PainlessParserID         = 82
	PainlessParserDOTINTEGER = 83
	PainlessParserDOTID      = 84
)

// PainlessParser rules.
const (
	PainlessParserRULE_source           = 0
	PainlessParserRULE_function         = 1
	PainlessParserRULE_parameters       = 2
	PainlessParserRULE_statement        = 3
	PainlessParserRULE_rstatement       = 4
	PainlessParserRULE_dstatement       = 5
	PainlessParserRULE_trailer          = 6
	PainlessParserRULE_block            = 7
	PainlessParserRULE_empty            = 8
	PainlessParserRULE_initializer      = 9
	PainlessParserRULE_afterthought     = 10
	PainlessParserRULE_declaration      = 11
	PainlessParserRULE_decltype         = 12
	PainlessParserRULE_declvar          = 13
	PainlessParserRULE_trap             = 14
	PainlessParserRULE_expression       = 15
	PainlessParserRULE_unary            = 16
	PainlessParserRULE_chain            = 17
	PainlessParserRULE_primary          = 18
	PainlessParserRULE_postfix          = 19
	PainlessParserRULE_postdot          = 20
	PainlessParserRULE_callinvoke       = 21
	PainlessParserRULE_fieldaccess      = 22
	PainlessParserRULE_braceaccess      = 23
	PainlessParserRULE_stdarray         = 24
	PainlessParserRULE_initializedarray = 25
	PainlessParserRULE_listinitializer  = 26
	PainlessParserRULE_mapinitializer   = 27
	PainlessParserRULE_maptoken         = 28
	PainlessParserRULE_arguments        = 29
	PainlessParserRULE_argument         = 30
	PainlessParserRULE_lambda           = 31
	PainlessParserRULE_lamtype          = 32
	PainlessParserRULE_funcref          = 33
)

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_source
	return p
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) EOF() antlr.TerminalNode {
	return s.GetToken(PainlessParserEOF, 0)
}

func (s *SourceContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *SourceContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *SourceContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SourceContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SourceContext) Dstatement() IDstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDstatementContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterSource(s)
	}
}

func (s *SourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitSource(s)
	}
}

func (p *PainlessParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PainlessParserRULE_source)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(68)
				p.Function()
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(74)
				p.Statement()
			}

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(PainlessParserLBRACE-5))|(1<<(PainlessParserLP-5))|(1<<(PainlessParserDO-5))|(1<<(PainlessParserCONTINUE-5))|(1<<(PainlessParserBREAK-5))|(1<<(PainlessParserRETURN-5))|(1<<(PainlessParserNEW-5))|(1<<(PainlessParserTHROW-5))|(1<<(PainlessParserBOOLNOT-5))|(1<<(PainlessParserBWNOT-5))|(1<<(PainlessParserADD-5))|(1<<(PainlessParserSUB-5)))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(PainlessParserINCR-58))|(1<<(PainlessParserDECR-58))|(1<<(PainlessParserOCTAL-58))|(1<<(PainlessParserHEX-58))|(1<<(PainlessParserINTEGER-58))|(1<<(PainlessParserDECIMAL-58))|(1<<(PainlessParserSTRING-58))|(1<<(PainlessParserREGEX-58))|(1<<(PainlessParserTRUE-58))|(1<<(PainlessParserFALSE-58))|(1<<(PainlessParserNULL-58))|(1<<(PainlessParserTYPE-58))|(1<<(PainlessParserID-58)))) != 0) {
		{
			p.SetState(80)
			p.Dstatement()
		}

	}
	{
		p.SetState(83)
		p.Match(PainlessParserEOF)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *PainlessParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PainlessParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Decltype()
	}
	{
		p.SetState(86)
		p.Match(PainlessParserID)
	}
	{
		p.SetState(87)
		p.Parameters()
	}
	{
		p.SetState(88)
		p.Block()
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *ParametersContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *ParametersContext) AllDecltype() []IDecltypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecltypeContext)(nil)).Elem())
	var tst = make([]IDecltypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecltypeContext)
		}
	}

	return tst
}

func (s *ParametersContext) Decltype(i int) IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *ParametersContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserID)
}

func (s *ParametersContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserID, i)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *PainlessParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PainlessParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(PainlessParserLP)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PainlessParserTYPE {
		{
			p.SetState(91)
			p.Decltype()
		}
		{
			p.SetState(92)
			p.Match(PainlessParserID)
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PainlessParserCOMMA {
			{
				p.SetState(93)
				p.Match(PainlessParserCOMMA)
			}
			{
				p.SetState(94)
				p.Decltype()
			}
			{
				p.SetState(95)
				p.Match(PainlessParserID)
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(104)
		p.Match(PainlessParserRP)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Rstatement() IRstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRstatementContext)
}

func (s *StatementContext) Dstatement() IDstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDstatementContext)
}

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PainlessParserSEMICOLON, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *PainlessParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PainlessParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PainlessParserIF, PainlessParserWHILE, PainlessParserFOR, PainlessParserTRY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Rstatement()
		}

	case PainlessParserLBRACE, PainlessParserLP, PainlessParserDO, PainlessParserCONTINUE, PainlessParserBREAK, PainlessParserRETURN, PainlessParserNEW, PainlessParserTHROW, PainlessParserBOOLNOT, PainlessParserBWNOT, PainlessParserADD, PainlessParserSUB, PainlessParserINCR, PainlessParserDECR, PainlessParserOCTAL, PainlessParserHEX, PainlessParserINTEGER, PainlessParserDECIMAL, PainlessParserSTRING, PainlessParserREGEX, PainlessParserTRUE, PainlessParserFALSE, PainlessParserNULL, PainlessParserTYPE, PainlessParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Dstatement()
		}
		{
			p.SetState(108)
			p.Match(PainlessParserSEMICOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRstatementContext is an interface to support dynamic dispatch.
type IRstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRstatementContext differentiates from other interfaces.
	IsRstatementContext()
}

type RstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRstatementContext() *RstatementContext {
	var p = new(RstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_rstatement
	return p
}

func (*RstatementContext) IsRstatementContext() {}

func NewRstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RstatementContext {
	var p = new(RstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_rstatement

	return p
}

func (s *RstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RstatementContext) CopyFrom(ctx *RstatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForContext struct {
	*RstatementContext
}

func NewForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForContext {
	var p = new(ForContext)

	p.RstatementContext = NewEmptyRstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RstatementContext))

	return p
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserFOR, 0)
}

func (s *ForContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *ForContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserSEMICOLON)
}

func (s *ForContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserSEMICOLON, i)
}

func (s *ForContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *ForContext) Trailer() ITrailerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrailerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrailerContext)
}

func (s *ForContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *ForContext) Initializer() IInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitializerContext)
}

func (s *ForContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForContext) Afterthought() IAfterthoughtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAfterthoughtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAfterthoughtContext)
}

func (s *ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterFor(s)
	}
}

func (s *ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitFor(s)
	}
}

type TryContext struct {
	*RstatementContext
}

func NewTryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TryContext {
	var p = new(TryContext)

	p.RstatementContext = NewEmptyRstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RstatementContext))

	return p
}

func (s *TryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryContext) TRY() antlr.TerminalNode {
	return s.GetToken(PainlessParserTRY, 0)
}

func (s *TryContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TryContext) AllTrap() []ITrapContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITrapContext)(nil)).Elem())
	var tst = make([]ITrapContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITrapContext)
		}
	}

	return tst
}

func (s *TryContext) Trap(i int) ITrapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrapContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITrapContext)
}

func (s *TryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterTry(s)
	}
}

func (s *TryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitTry(s)
	}
}

type WhileContext struct {
	*RstatementContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	p.RstatementContext = NewEmptyRstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RstatementContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(PainlessParserWHILE, 0)
}

func (s *WhileContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *WhileContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *WhileContext) Trailer() ITrailerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrailerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrailerContext)
}

func (s *WhileContext) Empty() IEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyContext)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitWhile(s)
	}
}

type IneachContext struct {
	*RstatementContext
}

func NewIneachContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IneachContext {
	var p = new(IneachContext)

	p.RstatementContext = NewEmptyRstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RstatementContext))

	return p
}

func (s *IneachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IneachContext) FOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserFOR, 0)
}

func (s *IneachContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *IneachContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *IneachContext) IN() antlr.TerminalNode {
	return s.GetToken(PainlessParserIN, 0)
}

func (s *IneachContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IneachContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *IneachContext) Trailer() ITrailerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrailerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrailerContext)
}

func (s *IneachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterIneach(s)
	}
}

func (s *IneachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitIneach(s)
	}
}

type IfContext struct {
	*RstatementContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.RstatementContext = NewEmptyRstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RstatementContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(PainlessParserIF, 0)
}

func (s *IfContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *IfContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *IfContext) AllTrailer() []ITrailerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITrailerContext)(nil)).Elem())
	var tst = make([]ITrailerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITrailerContext)
		}
	}

	return tst
}

func (s *IfContext) Trailer(i int) ITrailerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrailerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITrailerContext)
}

func (s *IfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PainlessParserELSE, 0)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitIf(s)
	}
}

type EachContext struct {
	*RstatementContext
}

func NewEachContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EachContext {
	var p = new(EachContext)

	p.RstatementContext = NewEmptyRstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RstatementContext))

	return p
}

func (s *EachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EachContext) FOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserFOR, 0)
}

func (s *EachContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *EachContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *EachContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *EachContext) COLON() antlr.TerminalNode {
	return s.GetToken(PainlessParserCOLON, 0)
}

func (s *EachContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EachContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *EachContext) Trailer() ITrailerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrailerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrailerContext)
}

func (s *EachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterEach(s)
	}
}

func (s *EachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitEach(s)
	}
}

func (p *PainlessParser) Rstatement() (localctx IRstatementContext) {
	localctx = NewRstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PainlessParserRULE_rstatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(PainlessParserIF)
		}
		{
			p.SetState(113)
			p.Match(PainlessParserLP)
		}
		{
			p.SetState(114)
			p.expression(0)
		}
		{
			p.SetState(115)
			p.Match(PainlessParserRP)
		}
		{
			p.SetState(116)
			p.Trailer()
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(117)
				p.Match(PainlessParserELSE)
			}
			{
				p.SetState(118)
				p.Trailer()
			}

		case 2:
			p.SetState(119)

			if !(p.GetTokenStream().LA(1) != PainlessParserELSE) {
				panic(antlr.NewFailedPredicateException(p, " p.GetTokenStream().LA(1) != PainlessParserELSE ", ""))
			}

		}

	case 2:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(PainlessParserWHILE)
		}
		{
			p.SetState(123)
			p.Match(PainlessParserLP)
		}
		{
			p.SetState(124)
			p.expression(0)
		}
		{
			p.SetState(125)
			p.Match(PainlessParserRP)
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PainlessParserLBRACK, PainlessParserLBRACE, PainlessParserLP, PainlessParserIF, PainlessParserWHILE, PainlessParserDO, PainlessParserFOR, PainlessParserCONTINUE, PainlessParserBREAK, PainlessParserRETURN, PainlessParserNEW, PainlessParserTRY, PainlessParserTHROW, PainlessParserBOOLNOT, PainlessParserBWNOT, PainlessParserADD, PainlessParserSUB, PainlessParserINCR, PainlessParserDECR, PainlessParserOCTAL, PainlessParserHEX, PainlessParserINTEGER, PainlessParserDECIMAL, PainlessParserSTRING, PainlessParserREGEX, PainlessParserTRUE, PainlessParserFALSE, PainlessParserNULL, PainlessParserTYPE, PainlessParserID:
			{
				p.SetState(126)
				p.Trailer()
			}

		case PainlessParserSEMICOLON:
			{
				p.SetState(127)
				p.Empty()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 3:
		localctx = NewForContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(PainlessParserFOR)
		}
		{
			p.SetState(131)
			p.Match(PainlessParserLP)
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(PainlessParserLBRACE-5))|(1<<(PainlessParserLP-5))|(1<<(PainlessParserNEW-5))|(1<<(PainlessParserBOOLNOT-5))|(1<<(PainlessParserBWNOT-5))|(1<<(PainlessParserADD-5))|(1<<(PainlessParserSUB-5)))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(PainlessParserINCR-58))|(1<<(PainlessParserDECR-58))|(1<<(PainlessParserOCTAL-58))|(1<<(PainlessParserHEX-58))|(1<<(PainlessParserINTEGER-58))|(1<<(PainlessParserDECIMAL-58))|(1<<(PainlessParserSTRING-58))|(1<<(PainlessParserREGEX-58))|(1<<(PainlessParserTRUE-58))|(1<<(PainlessParserFALSE-58))|(1<<(PainlessParserNULL-58))|(1<<(PainlessParserTYPE-58))|(1<<(PainlessParserID-58)))) != 0) {
			{
				p.SetState(132)
				p.Initializer()
			}

		}
		{
			p.SetState(135)
			p.Match(PainlessParserSEMICOLON)
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(PainlessParserLBRACE-5))|(1<<(PainlessParserLP-5))|(1<<(PainlessParserNEW-5))|(1<<(PainlessParserBOOLNOT-5))|(1<<(PainlessParserBWNOT-5))|(1<<(PainlessParserADD-5))|(1<<(PainlessParserSUB-5)))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(PainlessParserINCR-58))|(1<<(PainlessParserDECR-58))|(1<<(PainlessParserOCTAL-58))|(1<<(PainlessParserHEX-58))|(1<<(PainlessParserINTEGER-58))|(1<<(PainlessParserDECIMAL-58))|(1<<(PainlessParserSTRING-58))|(1<<(PainlessParserREGEX-58))|(1<<(PainlessParserTRUE-58))|(1<<(PainlessParserFALSE-58))|(1<<(PainlessParserNULL-58))|(1<<(PainlessParserTYPE-58))|(1<<(PainlessParserID-58)))) != 0) {
			{
				p.SetState(136)
				p.expression(0)
			}

		}
		{
			p.SetState(139)
			p.Match(PainlessParserSEMICOLON)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(PainlessParserLBRACE-5))|(1<<(PainlessParserLP-5))|(1<<(PainlessParserNEW-5))|(1<<(PainlessParserBOOLNOT-5))|(1<<(PainlessParserBWNOT-5))|(1<<(PainlessParserADD-5))|(1<<(PainlessParserSUB-5)))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(PainlessParserINCR-58))|(1<<(PainlessParserDECR-58))|(1<<(PainlessParserOCTAL-58))|(1<<(PainlessParserHEX-58))|(1<<(PainlessParserINTEGER-58))|(1<<(PainlessParserDECIMAL-58))|(1<<(PainlessParserSTRING-58))|(1<<(PainlessParserREGEX-58))|(1<<(PainlessParserTRUE-58))|(1<<(PainlessParserFALSE-58))|(1<<(PainlessParserNULL-58))|(1<<(PainlessParserTYPE-58))|(1<<(PainlessParserID-58)))) != 0) {
			{
				p.SetState(140)
				p.Afterthought()
			}

		}
		{
			p.SetState(143)
			p.Match(PainlessParserRP)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PainlessParserLBRACK, PainlessParserLBRACE, PainlessParserLP, PainlessParserIF, PainlessParserWHILE, PainlessParserDO, PainlessParserFOR, PainlessParserCONTINUE, PainlessParserBREAK, PainlessParserRETURN, PainlessParserNEW, PainlessParserTRY, PainlessParserTHROW, PainlessParserBOOLNOT, PainlessParserBWNOT, PainlessParserADD, PainlessParserSUB, PainlessParserINCR, PainlessParserDECR, PainlessParserOCTAL, PainlessParserHEX, PainlessParserINTEGER, PainlessParserDECIMAL, PainlessParserSTRING, PainlessParserREGEX, PainlessParserTRUE, PainlessParserFALSE, PainlessParserNULL, PainlessParserTYPE, PainlessParserID:
			{
				p.SetState(144)
				p.Trailer()
			}

		case PainlessParserSEMICOLON:
			{
				p.SetState(145)
				p.Empty()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 4:
		localctx = NewEachContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.Match(PainlessParserFOR)
		}
		{
			p.SetState(149)
			p.Match(PainlessParserLP)
		}
		{
			p.SetState(150)
			p.Decltype()
		}
		{
			p.SetState(151)
			p.Match(PainlessParserID)
		}
		{
			p.SetState(152)
			p.Match(PainlessParserCOLON)
		}
		{
			p.SetState(153)
			p.expression(0)
		}
		{
			p.SetState(154)
			p.Match(PainlessParserRP)
		}
		{
			p.SetState(155)
			p.Trailer()
		}

	case 5:
		localctx = NewIneachContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.Match(PainlessParserFOR)
		}
		{
			p.SetState(158)
			p.Match(PainlessParserLP)
		}
		{
			p.SetState(159)
			p.Match(PainlessParserID)
		}
		{
			p.SetState(160)
			p.Match(PainlessParserIN)
		}
		{
			p.SetState(161)
			p.expression(0)
		}
		{
			p.SetState(162)
			p.Match(PainlessParserRP)
		}
		{
			p.SetState(163)
			p.Trailer()
		}

	case 6:
		localctx = NewTryContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(165)
			p.Match(PainlessParserTRY)
		}
		{
			p.SetState(166)
			p.Block()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(167)
					p.Trap()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IDstatementContext is an interface to support dynamic dispatch.
type IDstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDstatementContext differentiates from other interfaces.
	IsDstatementContext()
}

type DstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDstatementContext() *DstatementContext {
	var p = new(DstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_dstatement
	return p
}

func (*DstatementContext) IsDstatementContext() {}

func NewDstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DstatementContext {
	var p = new(DstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_dstatement

	return p
}

func (s *DstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DstatementContext) CopyFrom(ctx *DstatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclContext struct {
	*DstatementContext
}

func NewDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclContext {
	var p = new(DeclContext)

	p.DstatementContext = NewEmptyDstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DstatementContext))

	return p
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

type BreakContext struct {
	*DstatementContext
}

func NewBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakContext {
	var p = new(BreakContext)

	p.DstatementContext = NewEmptyDstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DstatementContext))

	return p
}

func (s *BreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(PainlessParserBREAK, 0)
}

func (s *BreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterBreak(s)
	}
}

func (s *BreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitBreak(s)
	}
}

type ThrowContext struct {
	*DstatementContext
}

func NewThrowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThrowContext {
	var p = new(ThrowContext)

	p.DstatementContext = NewEmptyDstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DstatementContext))

	return p
}

func (s *ThrowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowContext) THROW() antlr.TerminalNode {
	return s.GetToken(PainlessParserTHROW, 0)
}

func (s *ThrowContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ThrowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterThrow(s)
	}
}

func (s *ThrowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitThrow(s)
	}
}

type ContinueContext struct {
	*DstatementContext
}

func NewContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueContext {
	var p = new(ContinueContext)

	p.DstatementContext = NewEmptyDstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DstatementContext))

	return p
}

func (s *ContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(PainlessParserCONTINUE, 0)
}

func (s *ContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterContinue(s)
	}
}

func (s *ContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitContinue(s)
	}
}

type ExprContext struct {
	*DstatementContext
}

func NewExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprContext {
	var p = new(ExprContext)

	p.DstatementContext = NewEmptyDstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DstatementContext))

	return p
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

type DoContext struct {
	*DstatementContext
}

func NewDoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoContext {
	var p = new(DoContext)

	p.DstatementContext = NewEmptyDstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DstatementContext))

	return p
}

func (s *DoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoContext) DO() antlr.TerminalNode {
	return s.GetToken(PainlessParserDO, 0)
}

func (s *DoContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DoContext) WHILE() antlr.TerminalNode {
	return s.GetToken(PainlessParserWHILE, 0)
}

func (s *DoContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *DoContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DoContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *DoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterDo(s)
	}
}

func (s *DoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitDo(s)
	}
}

type ReturnContext struct {
	*DstatementContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	p.DstatementContext = NewEmptyDstatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DstatementContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(PainlessParserRETURN, 0)
}

func (s *ReturnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *PainlessParser) Dstatement() (localctx IDstatementContext) {
	localctx = NewDstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PainlessParserRULE_dstatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.Match(PainlessParserDO)
		}
		{
			p.SetState(175)
			p.Block()
		}
		{
			p.SetState(176)
			p.Match(PainlessParserWHILE)
		}
		{
			p.SetState(177)
			p.Match(PainlessParserLP)
		}
		{
			p.SetState(178)
			p.expression(0)
		}
		{
			p.SetState(179)
			p.Match(PainlessParserRP)
		}

	case 2:
		localctx = NewDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Declaration()
		}

	case 3:
		localctx = NewContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Match(PainlessParserCONTINUE)
		}

	case 4:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Match(PainlessParserBREAK)
		}

	case 5:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Match(PainlessParserRETURN)
		}
		{
			p.SetState(185)
			p.expression(0)
		}

	case 6:
		localctx = NewThrowContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(186)
			p.Match(PainlessParserTHROW)
		}
		{
			p.SetState(187)
			p.expression(0)
		}

	case 7:
		localctx = NewExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(188)
			p.expression(0)
		}

	}

	return localctx
}

// ITrailerContext is an interface to support dynamic dispatch.
type ITrailerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrailerContext differentiates from other interfaces.
	IsTrailerContext()
}

type TrailerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrailerContext() *TrailerContext {
	var p = new(TrailerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_trailer
	return p
}

func (*TrailerContext) IsTrailerContext() {}

func NewTrailerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrailerContext {
	var p = new(TrailerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_trailer

	return p
}

func (s *TrailerContext) GetParser() antlr.Parser { return s.parser }

func (s *TrailerContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TrailerContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *TrailerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrailerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrailerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterTrailer(s)
	}
}

func (s *TrailerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitTrailer(s)
	}
}

func (p *PainlessParser) Trailer() (localctx ITrailerContext) {
	localctx = NewTrailerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PainlessParserRULE_trailer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(193)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PainlessParserLBRACK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.Block()
		}

	case PainlessParserLBRACE, PainlessParserLP, PainlessParserIF, PainlessParserWHILE, PainlessParserDO, PainlessParserFOR, PainlessParserCONTINUE, PainlessParserBREAK, PainlessParserRETURN, PainlessParserNEW, PainlessParserTRY, PainlessParserTHROW, PainlessParserBOOLNOT, PainlessParserBWNOT, PainlessParserADD, PainlessParserSUB, PainlessParserINCR, PainlessParserDECR, PainlessParserOCTAL, PainlessParserHEX, PainlessParserINTEGER, PainlessParserDECIMAL, PainlessParserSTRING, PainlessParserREGEX, PainlessParserTRUE, PainlessParserFALSE, PainlessParserNULL, PainlessParserTYPE, PainlessParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACK, 0)
}

func (s *BlockContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACK, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) Dstatement() IDstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDstatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *PainlessParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PainlessParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(PainlessParserLBRACK)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(196)
				p.Statement()
			}

		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(PainlessParserLBRACE-5))|(1<<(PainlessParserLP-5))|(1<<(PainlessParserDO-5))|(1<<(PainlessParserCONTINUE-5))|(1<<(PainlessParserBREAK-5))|(1<<(PainlessParserRETURN-5))|(1<<(PainlessParserNEW-5))|(1<<(PainlessParserTHROW-5))|(1<<(PainlessParserBOOLNOT-5))|(1<<(PainlessParserBWNOT-5))|(1<<(PainlessParserADD-5))|(1<<(PainlessParserSUB-5)))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(PainlessParserINCR-58))|(1<<(PainlessParserDECR-58))|(1<<(PainlessParserOCTAL-58))|(1<<(PainlessParserHEX-58))|(1<<(PainlessParserINTEGER-58))|(1<<(PainlessParserDECIMAL-58))|(1<<(PainlessParserSTRING-58))|(1<<(PainlessParserREGEX-58))|(1<<(PainlessParserTRUE-58))|(1<<(PainlessParserFALSE-58))|(1<<(PainlessParserNULL-58))|(1<<(PainlessParserTYPE-58))|(1<<(PainlessParserID-58)))) != 0) {
		{
			p.SetState(202)
			p.Dstatement()
		}

	}
	{
		p.SetState(205)
		p.Match(PainlessParserRBRACK)
	}

	return localctx
}

// IEmptyContext is an interface to support dynamic dispatch.
type IEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyContext differentiates from other interfaces.
	IsEmptyContext()
}

type EmptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyContext() *EmptyContext {
	var p = new(EmptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_empty
	return p
}

func (*EmptyContext) IsEmptyContext() {}

func NewEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyContext {
	var p = new(EmptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_empty

	return p
}

func (s *EmptyContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PainlessParserSEMICOLON, 0)
}

func (s *EmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterEmpty(s)
	}
}

func (s *EmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitEmpty(s)
	}
}

func (p *PainlessParser) Empty() (localctx IEmptyContext) {
	localctx = NewEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PainlessParserRULE_empty)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(PainlessParserSEMICOLON)
	}

	return localctx
}

// IInitializerContext is an interface to support dynamic dispatch.
type IInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitializerContext differentiates from other interfaces.
	IsInitializerContext()
}

type InitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializerContext() *InitializerContext {
	var p = new(InitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_initializer
	return p
}

func (*InitializerContext) IsInitializerContext() {}

func NewInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializerContext {
	var p = new(InitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_initializer

	return p
}

func (s *InitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializerContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *InitializerContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterInitializer(s)
	}
}

func (s *InitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitInitializer(s)
	}
}

func (p *PainlessParser) Initializer() (localctx IInitializerContext) {
	localctx = NewInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PainlessParserRULE_initializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.expression(0)
		}

	}

	return localctx
}

// IAfterthoughtContext is an interface to support dynamic dispatch.
type IAfterthoughtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAfterthoughtContext differentiates from other interfaces.
	IsAfterthoughtContext()
}

type AfterthoughtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAfterthoughtContext() *AfterthoughtContext {
	var p = new(AfterthoughtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_afterthought
	return p
}

func (*AfterthoughtContext) IsAfterthoughtContext() {}

func NewAfterthoughtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AfterthoughtContext {
	var p = new(AfterthoughtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_afterthought

	return p
}

func (s *AfterthoughtContext) GetParser() antlr.Parser { return s.parser }

func (s *AfterthoughtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AfterthoughtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterthoughtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AfterthoughtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterAfterthought(s)
	}
}

func (s *AfterthoughtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitAfterthought(s)
	}
}

func (p *PainlessParser) Afterthought() (localctx IAfterthoughtContext) {
	localctx = NewAfterthoughtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PainlessParserRULE_afterthought)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.expression(0)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *DeclarationContext) AllDeclvar() []IDeclvarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclvarContext)(nil)).Elem())
	var tst = make([]IDeclvarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclvarContext)
		}
	}

	return tst
}

func (s *DeclarationContext) Declvar(i int) IDeclvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclvarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclvarContext)
}

func (s *DeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserCOMMA)
}

func (s *DeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserCOMMA, i)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *PainlessParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PainlessParserRULE_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Decltype()
	}
	{
		p.SetState(216)
		p.Declvar()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PainlessParserCOMMA {
		{
			p.SetState(217)
			p.Match(PainlessParserCOMMA)
		}
		{
			p.SetState(218)
			p.Declvar()
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDecltypeContext is an interface to support dynamic dispatch.
type IDecltypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecltypeContext differentiates from other interfaces.
	IsDecltypeContext()
}

type DecltypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecltypeContext() *DecltypeContext {
	var p = new(DecltypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_decltype
	return p
}

func (*DecltypeContext) IsDecltypeContext() {}

func NewDecltypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecltypeContext {
	var p = new(DecltypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_decltype

	return p
}

func (s *DecltypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DecltypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PainlessParserTYPE, 0)
}

func (s *DecltypeContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserLBRACE)
}

func (s *DecltypeContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACE, i)
}

func (s *DecltypeContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserRBRACE)
}

func (s *DecltypeContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACE, i)
}

func (s *DecltypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecltypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecltypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterDecltype(s)
	}
}

func (s *DecltypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitDecltype(s)
	}
}

func (p *PainlessParser) Decltype() (localctx IDecltypeContext) {
	localctx = NewDecltypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PainlessParserRULE_decltype)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(PainlessParserTYPE)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(225)
				p.Match(PainlessParserLBRACE)
			}
			{
				p.SetState(226)
				p.Match(PainlessParserRBRACE)
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclvarContext is an interface to support dynamic dispatch.
type IDeclvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclvarContext differentiates from other interfaces.
	IsDeclvarContext()
}

type DeclvarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclvarContext() *DeclvarContext {
	var p = new(DeclvarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_declvar
	return p
}

func (*DeclvarContext) IsDeclvarContext() {}

func NewDeclvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclvarContext {
	var p = new(DeclvarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_declvar

	return p
}

func (s *DeclvarContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclvarContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *DeclvarContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PainlessParserASSIGN, 0)
}

func (s *DeclvarContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterDeclvar(s)
	}
}

func (s *DeclvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitDeclvar(s)
	}
}

func (p *PainlessParser) Declvar() (localctx IDeclvarContext) {
	localctx = NewDeclvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PainlessParserRULE_declvar)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(PainlessParserID)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PainlessParserASSIGN {
		{
			p.SetState(233)
			p.Match(PainlessParserASSIGN)
		}
		{
			p.SetState(234)
			p.expression(0)
		}

	}

	return localctx
}

// ITrapContext is an interface to support dynamic dispatch.
type ITrapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrapContext differentiates from other interfaces.
	IsTrapContext()
}

type TrapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrapContext() *TrapContext {
	var p = new(TrapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_trap
	return p
}

func (*TrapContext) IsTrapContext() {}

func NewTrapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrapContext {
	var p = new(TrapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_trap

	return p
}

func (s *TrapContext) GetParser() antlr.Parser { return s.parser }

func (s *TrapContext) CATCH() antlr.TerminalNode {
	return s.GetToken(PainlessParserCATCH, 0)
}

func (s *TrapContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *TrapContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PainlessParserTYPE, 0)
}

func (s *TrapContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *TrapContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *TrapContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TrapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterTrap(s)
	}
}

func (s *TrapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitTrap(s)
	}
}

func (p *PainlessParser) Trap() (localctx ITrapContext) {
	localctx = NewTrapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PainlessParserRULE_trap)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(PainlessParserCATCH)
	}
	{
		p.SetState(238)
		p.Match(PainlessParserLP)
	}
	{
		p.SetState(239)
		p.Match(PainlessParserTYPE)
	}
	{
		p.SetState(240)
		p.Match(PainlessParserID)
	}
	{
		p.SetState(241)
		p.Match(PainlessParserRP)
	}
	{
		p.SetState(242)
		p.Block()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleContext struct {
	*ExpressionContext
}

func NewSingleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleContext {
	var p = new(SingleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleContext) Unary() IUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *SingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterSingle(s)
	}
}

func (s *SingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitSingle(s)
	}
}

type CompContext struct {
	*ExpressionContext
}

func NewCompContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompContext {
	var p = new(CompContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompContext) LT() antlr.TerminalNode {
	return s.GetToken(PainlessParserLT, 0)
}

func (s *CompContext) LTE() antlr.TerminalNode {
	return s.GetToken(PainlessParserLTE, 0)
}

func (s *CompContext) GT() antlr.TerminalNode {
	return s.GetToken(PainlessParserGT, 0)
}

func (s *CompContext) GTE() antlr.TerminalNode {
	return s.GetToken(PainlessParserGTE, 0)
}

func (s *CompContext) EQ() antlr.TerminalNode {
	return s.GetToken(PainlessParserEQ, 0)
}

func (s *CompContext) EQR() antlr.TerminalNode {
	return s.GetToken(PainlessParserEQR, 0)
}

func (s *CompContext) NE() antlr.TerminalNode {
	return s.GetToken(PainlessParserNE, 0)
}

func (s *CompContext) NER() antlr.TerminalNode {
	return s.GetToken(PainlessParserNER, 0)
}

func (s *CompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterComp(s)
	}
}

func (s *CompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitComp(s)
	}
}

type BoolContext struct {
	*ExpressionContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BoolContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BoolContext) BOOLAND() antlr.TerminalNode {
	return s.GetToken(PainlessParserBOOLAND, 0)
}

func (s *BoolContext) BOOLOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserBOOLOR, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitBool(s)
	}
}

type ConditionalContext struct {
	*ExpressionContext
}

func NewConditionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalContext {
	var p = new(ConditionalContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalContext) COND() antlr.TerminalNode {
	return s.GetToken(PainlessParserCOND, 0)
}

func (s *ConditionalContext) COLON() antlr.TerminalNode {
	return s.GetToken(PainlessParserCOLON, 0)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

type AssignmentContext struct {
	*ExpressionContext
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PainlessParserASSIGN, 0)
}

func (s *AssignmentContext) AADD() antlr.TerminalNode {
	return s.GetToken(PainlessParserAADD, 0)
}

func (s *AssignmentContext) ASUB() antlr.TerminalNode {
	return s.GetToken(PainlessParserASUB, 0)
}

func (s *AssignmentContext) AMUL() antlr.TerminalNode {
	return s.GetToken(PainlessParserAMUL, 0)
}

func (s *AssignmentContext) ADIV() antlr.TerminalNode {
	return s.GetToken(PainlessParserADIV, 0)
}

func (s *AssignmentContext) AREM() antlr.TerminalNode {
	return s.GetToken(PainlessParserAREM, 0)
}

func (s *AssignmentContext) AAND() antlr.TerminalNode {
	return s.GetToken(PainlessParserAAND, 0)
}

func (s *AssignmentContext) AXOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserAXOR, 0)
}

func (s *AssignmentContext) AOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserAOR, 0)
}

func (s *AssignmentContext) ALSH() antlr.TerminalNode {
	return s.GetToken(PainlessParserALSH, 0)
}

func (s *AssignmentContext) ARSH() antlr.TerminalNode {
	return s.GetToken(PainlessParserARSH, 0)
}

func (s *AssignmentContext) AUSH() antlr.TerminalNode {
	return s.GetToken(PainlessParserAUSH, 0)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

type BinaryContext struct {
	*ExpressionContext
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryContext) MUL() antlr.TerminalNode {
	return s.GetToken(PainlessParserMUL, 0)
}

func (s *BinaryContext) DIV() antlr.TerminalNode {
	return s.GetToken(PainlessParserDIV, 0)
}

func (s *BinaryContext) REM() antlr.TerminalNode {
	return s.GetToken(PainlessParserREM, 0)
}

func (s *BinaryContext) ADD() antlr.TerminalNode {
	return s.GetToken(PainlessParserADD, 0)
}

func (s *BinaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(PainlessParserSUB, 0)
}

func (s *BinaryContext) FIND() antlr.TerminalNode {
	return s.GetToken(PainlessParserFIND, 0)
}

func (s *BinaryContext) MATCH() antlr.TerminalNode {
	return s.GetToken(PainlessParserMATCH, 0)
}

func (s *BinaryContext) LSH() antlr.TerminalNode {
	return s.GetToken(PainlessParserLSH, 0)
}

func (s *BinaryContext) RSH() antlr.TerminalNode {
	return s.GetToken(PainlessParserRSH, 0)
}

func (s *BinaryContext) USH() antlr.TerminalNode {
	return s.GetToken(PainlessParserUSH, 0)
}

func (s *BinaryContext) BWAND() antlr.TerminalNode {
	return s.GetToken(PainlessParserBWAND, 0)
}

func (s *BinaryContext) XOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserXOR, 0)
}

func (s *BinaryContext) BWOR() antlr.TerminalNode {
	return s.GetToken(PainlessParserBWOR, 0)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitBinary(s)
	}
}

type ElvisContext struct {
	*ExpressionContext
}

func NewElvisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElvisContext {
	var p = new(ElvisContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ElvisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElvisContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ElvisContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElvisContext) ELVIS() antlr.TerminalNode {
	return s.GetToken(PainlessParserELVIS, 0)
}

func (s *ElvisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterElvis(s)
	}
}

func (s *ElvisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitElvis(s)
	}
}

type InstanceofContext struct {
	*ExpressionContext
}

func NewInstanceofContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanceofContext {
	var p = new(InstanceofContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InstanceofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceofContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InstanceofContext) INSTANCEOF() antlr.TerminalNode {
	return s.GetToken(PainlessParserINSTANCEOF, 0)
}

func (s *InstanceofContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *InstanceofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterInstanceof(s)
	}
}

func (s *InstanceofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitInstanceof(s)
	}
}

func (p *PainlessParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *PainlessParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, PainlessParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewSingleContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(245)
		p.Unary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(248)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(PainlessParserMUL-30))|(1<<(PainlessParserDIV-30))|(1<<(PainlessParserREM-30)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(249)
					p.expression(16)
				}

			case 2:
				localctx = NewBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(251)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PainlessParserADD || _la == PainlessParserSUB) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(252)
					p.expression(15)
				}

			case 3:
				localctx = NewBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(253)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(254)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PainlessParserFIND || _la == PainlessParserMATCH) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(255)
					p.expression(14)
				}

			case 4:
				localctx = NewBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(257)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(PainlessParserLSH-35))|(1<<(PainlessParserRSH-35))|(1<<(PainlessParserUSH-35)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(258)
					p.expression(13)
				}

			case 5:
				localctx = NewCompContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(259)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(260)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(PainlessParserLT-38))|(1<<(PainlessParserLTE-38))|(1<<(PainlessParserGT-38))|(1<<(PainlessParserGTE-38)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(261)
					p.expression(12)
				}

			case 6:
				localctx = NewCompContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(263)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(PainlessParserEQ-42))|(1<<(PainlessParserEQR-42))|(1<<(PainlessParserNE-42))|(1<<(PainlessParserNER-42)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(264)
					p.expression(10)
				}

			case 7:
				localctx = NewBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(266)
					p.Match(PainlessParserBWAND)
				}
				{
					p.SetState(267)
					p.expression(9)
				}

			case 8:
				localctx = NewBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(269)
					p.Match(PainlessParserXOR)
				}
				{
					p.SetState(270)
					p.expression(8)
				}

			case 9:
				localctx = NewBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(272)
					p.Match(PainlessParserBWOR)
				}
				{
					p.SetState(273)
					p.expression(7)
				}

			case 10:
				localctx = NewBoolContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(275)
					p.Match(PainlessParserBOOLAND)
				}
				{
					p.SetState(276)
					p.expression(6)
				}

			case 11:
				localctx = NewBoolContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(278)
					p.Match(PainlessParserBOOLOR)
				}
				{
					p.SetState(279)
					p.expression(5)
				}

			case 12:
				localctx = NewConditionalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(281)
					p.Match(PainlessParserCOND)
				}
				{
					p.SetState(282)
					p.expression(0)
				}
				{
					p.SetState(283)
					p.Match(PainlessParserCOLON)
				}
				{
					p.SetState(284)
					p.expression(3)
				}

			case 13:
				localctx = NewElvisContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(287)
					p.Match(PainlessParserELVIS)
				}
				{
					p.SetState(288)
					p.expression(2)
				}

			case 14:
				localctx = NewAssignmentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(290)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(PainlessParserASSIGN-60))|(1<<(PainlessParserAADD-60))|(1<<(PainlessParserASUB-60))|(1<<(PainlessParserAMUL-60))|(1<<(PainlessParserADIV-60))|(1<<(PainlessParserAREM-60))|(1<<(PainlessParserAAND-60))|(1<<(PainlessParserAXOR-60))|(1<<(PainlessParserAOR-60))|(1<<(PainlessParserALSH-60))|(1<<(PainlessParserARSH-60))|(1<<(PainlessParserAUSH-60)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(291)
					p.expression(1)
				}

			case 15:
				localctx = NewInstanceofContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PainlessParserRULE_expression)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(293)
					p.Match(PainlessParserINSTANCEOF)
				}
				{
					p.SetState(294)
					p.Decltype()
				}

			}

		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_unary
	return p
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) CopyFrom(ctx *UnaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CastContext struct {
	*UnaryContext
}

func NewCastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastContext {
	var p = new(CastContext)

	p.UnaryContext = NewEmptyUnaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryContext))

	return p
}

func (s *CastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *CastContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *CastContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *CastContext) Unary() IUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *CastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterCast(s)
	}
}

func (s *CastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitCast(s)
	}
}

type PreContext struct {
	*UnaryContext
}

func NewPreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreContext {
	var p = new(PreContext)

	p.UnaryContext = NewEmptyUnaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryContext))

	return p
}

func (s *PreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *PreContext) INCR() antlr.TerminalNode {
	return s.GetToken(PainlessParserINCR, 0)
}

func (s *PreContext) DECR() antlr.TerminalNode {
	return s.GetToken(PainlessParserDECR, 0)
}

func (s *PreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterPre(s)
	}
}

func (s *PreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitPre(s)
	}
}

type ReadContext struct {
	*UnaryContext
}

func NewReadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReadContext {
	var p = new(ReadContext)

	p.UnaryContext = NewEmptyUnaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryContext))

	return p
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitRead(s)
	}
}

type PostContext struct {
	*UnaryContext
}

func NewPostContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostContext {
	var p = new(PostContext)

	p.UnaryContext = NewEmptyUnaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryContext))

	return p
}

func (s *PostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *PostContext) INCR() antlr.TerminalNode {
	return s.GetToken(PainlessParserINCR, 0)
}

func (s *PostContext) DECR() antlr.TerminalNode {
	return s.GetToken(PainlessParserDECR, 0)
}

func (s *PostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterPost(s)
	}
}

func (s *PostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitPost(s)
	}
}

type OperatorContext struct {
	*UnaryContext
}

func NewOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperatorContext {
	var p = new(OperatorContext)

	p.UnaryContext = NewEmptyUnaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryContext))

	return p
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) Unary() IUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *OperatorContext) BOOLNOT() antlr.TerminalNode {
	return s.GetToken(PainlessParserBOOLNOT, 0)
}

func (s *OperatorContext) BWNOT() antlr.TerminalNode {
	return s.GetToken(PainlessParserBWNOT, 0)
}

func (s *OperatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(PainlessParserADD, 0)
}

func (s *OperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(PainlessParserSUB, 0)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *PainlessParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PainlessParserRULE_unary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPreContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PainlessParserINCR || _la == PainlessParserDECR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(301)
			p.Chain()
		}

	case 2:
		localctx = NewPostContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.Chain()
		}
		{
			p.SetState(303)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PainlessParserINCR || _la == PainlessParserDECR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewReadContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(305)
			p.Chain()
		}

	case 4:
		localctx = NewOperatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(306)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(PainlessParserBOOLNOT-28))|(1<<(PainlessParserBWNOT-28))|(1<<(PainlessParserADD-28))|(1<<(PainlessParserSUB-28)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(307)
			p.Unary()
		}

	case 5:
		localctx = NewCastContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(308)
			p.Match(PainlessParserLP)
		}
		{
			p.SetState(309)
			p.Decltype()
		}
		{
			p.SetState(310)
			p.Match(PainlessParserRP)
		}
		{
			p.SetState(311)
			p.Unary()
		}

	}

	return localctx
}

// IChainContext is an interface to support dynamic dispatch.
type IChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChainContext differentiates from other interfaces.
	IsChainContext()
}

type ChainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainContext() *ChainContext {
	var p = new(ChainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_chain
	return p
}

func (*ChainContext) IsChainContext() {}

func NewChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainContext {
	var p = new(ChainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_chain

	return p
}

func (s *ChainContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ChainContext) AllPostfix() []IPostfixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPostfixContext)(nil)).Elem())
	var tst = make([]IPostfixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPostfixContext)
		}
	}

	return tst
}

func (s *ChainContext) Postfix(i int) IPostfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPostfixContext)
}

func (s *ChainContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *ChainContext) Postdot() IPostdotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostdotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostdotContext)
}

func (s *ChainContext) Stdarray() IStdarrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStdarrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStdarrayContext)
}

func (s *ChainContext) Initializedarray() IInitializedarrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializedarrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitializedarrayContext)
}

func (s *ChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterChain(s)
	}
}

func (s *ChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitChain(s)
	}
}

func (p *PainlessParser) Chain() (localctx IChainContext) {
	localctx = NewChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PainlessParserRULE_chain)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.Primary()
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(316)
					p.Postfix()
				}

			}
			p.SetState(321)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Decltype()
		}
		{
			p.SetState(323)
			p.Postdot()
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(324)
					p.Postfix()
				}

			}
			p.SetState(329)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)
			p.Stdarray()
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(331)
				p.Postdot()
			}
			p.SetState(335)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(332)
						p.Postfix()
					}

				}
				p.SetState(337)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(340)
			p.Initializedarray()
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(341)
					p.Postfix()
				}

			}
			p.SetState(346)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyFrom(ctx *PrimaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListinitContext struct {
	*PrimaryContext
}

func NewListinitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListinitContext {
	var p = new(ListinitContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *ListinitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListinitContext) Listinitializer() IListinitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListinitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListinitializerContext)
}

func (s *ListinitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterListinit(s)
	}
}

func (s *ListinitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitListinit(s)
	}
}

type RegexContext struct {
	*PrimaryContext
}

func NewRegexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexContext {
	var p = new(RegexContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *RegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexContext) REGEX() antlr.TerminalNode {
	return s.GetToken(PainlessParserREGEX, 0)
}

func (s *RegexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterRegex(s)
	}
}

func (s *RegexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitRegex(s)
	}
}

type NullContext struct {
	*PrimaryContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(PainlessParserNULL, 0)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitNull(s)
	}
}

type StringContext struct {
	*PrimaryContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(PainlessParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitString(s)
	}
}

type MapinitContext struct {
	*PrimaryContext
}

func NewMapinitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapinitContext {
	var p = new(MapinitContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *MapinitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapinitContext) Mapinitializer() IMapinitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapinitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapinitializerContext)
}

func (s *MapinitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterMapinit(s)
	}
}

func (s *MapinitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitMapinit(s)
	}
}

type CalllocalContext struct {
	*PrimaryContext
}

func NewCalllocalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CalllocalContext {
	var p = new(CalllocalContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *CalllocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalllocalContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *CalllocalContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CalllocalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterCalllocal(s)
	}
}

func (s *CalllocalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitCalllocal(s)
	}
}

type TrueContext struct {
	*PrimaryContext
}

func NewTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueContext {
	var p = new(TrueContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(PainlessParserTRUE, 0)
}

func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitTrue(s)
	}
}

type FalseContext struct {
	*PrimaryContext
}

func NewFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseContext {
	var p = new(FalseContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(PainlessParserFALSE, 0)
}

func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitFalse(s)
	}
}

type VariableContext struct {
	*PrimaryContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitVariable(s)
	}
}

type NumericContext struct {
	*PrimaryContext
}

func NewNumericContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericContext {
	var p = new(NumericContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *NumericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericContext) OCTAL() antlr.TerminalNode {
	return s.GetToken(PainlessParserOCTAL, 0)
}

func (s *NumericContext) HEX() antlr.TerminalNode {
	return s.GetToken(PainlessParserHEX, 0)
}

func (s *NumericContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PainlessParserINTEGER, 0)
}

func (s *NumericContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(PainlessParserDECIMAL, 0)
}

func (s *NumericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterNumeric(s)
	}
}

func (s *NumericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitNumeric(s)
	}
}

type NewobjectContext struct {
	*PrimaryContext
}

func NewNewobjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewobjectContext {
	var p = new(NewobjectContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *NewobjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewobjectContext) NEW() antlr.TerminalNode {
	return s.GetToken(PainlessParserNEW, 0)
}

func (s *NewobjectContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PainlessParserTYPE, 0)
}

func (s *NewobjectContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *NewobjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterNewobject(s)
	}
}

func (s *NewobjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitNewobject(s)
	}
}

type PrecedenceContext struct {
	*PrimaryContext
}

func NewPrecedenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrecedenceContext {
	var p = new(PrecedenceContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *PrecedenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecedenceContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *PrecedenceContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrecedenceContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *PrecedenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterPrecedence(s)
	}
}

func (s *PrecedenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitPrecedence(s)
	}
}

func (p *PainlessParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PainlessParserRULE_primary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrecedenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Match(PainlessParserLP)
		}
		{
			p.SetState(350)
			p.expression(0)
		}
		{
			p.SetState(351)
			p.Match(PainlessParserRP)
		}

	case 2:
		localctx = NewNumericContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(353)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(PainlessParserOCTAL-72))|(1<<(PainlessParserHEX-72))|(1<<(PainlessParserINTEGER-72))|(1<<(PainlessParserDECIMAL-72)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(354)
			p.Match(PainlessParserTRUE)
		}

	case 4:
		localctx = NewFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(355)
			p.Match(PainlessParserFALSE)
		}

	case 5:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(356)
			p.Match(PainlessParserNULL)
		}

	case 6:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(357)
			p.Match(PainlessParserSTRING)
		}

	case 7:
		localctx = NewRegexContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(358)
			p.Match(PainlessParserREGEX)
		}

	case 8:
		localctx = NewListinitContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(359)
			p.Listinitializer()
		}

	case 9:
		localctx = NewMapinitContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(360)
			p.Mapinitializer()
		}

	case 10:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(361)
			p.Match(PainlessParserID)
		}

	case 11:
		localctx = NewCalllocalContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(362)
			p.Match(PainlessParserID)
		}
		{
			p.SetState(363)
			p.Arguments()
		}

	case 12:
		localctx = NewNewobjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(364)
			p.Match(PainlessParserNEW)
		}
		{
			p.SetState(365)
			p.Match(PainlessParserTYPE)
		}
		{
			p.SetState(366)
			p.Arguments()
		}

	}

	return localctx
}

// IPostfixContext is an interface to support dynamic dispatch.
type IPostfixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostfixContext differentiates from other interfaces.
	IsPostfixContext()
}

type PostfixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixContext() *PostfixContext {
	var p = new(PostfixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_postfix
	return p
}

func (*PostfixContext) IsPostfixContext() {}

func NewPostfixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixContext {
	var p = new(PostfixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_postfix

	return p
}

func (s *PostfixContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixContext) Callinvoke() ICallinvokeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallinvokeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallinvokeContext)
}

func (s *PostfixContext) Fieldaccess() IFieldaccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldaccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldaccessContext)
}

func (s *PostfixContext) Braceaccess() IBraceaccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBraceaccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBraceaccessContext)
}

func (s *PostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterPostfix(s)
	}
}

func (s *PostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitPostfix(s)
	}
}

func (p *PainlessParser) Postfix() (localctx IPostfixContext) {
	localctx = NewPostfixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PainlessParserRULE_postfix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)
			p.Callinvoke()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Fieldaccess()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(371)
			p.Braceaccess()
		}

	}

	return localctx
}

// IPostdotContext is an interface to support dynamic dispatch.
type IPostdotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostdotContext differentiates from other interfaces.
	IsPostdotContext()
}

type PostdotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostdotContext() *PostdotContext {
	var p = new(PostdotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_postdot
	return p
}

func (*PostdotContext) IsPostdotContext() {}

func NewPostdotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostdotContext {
	var p = new(PostdotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_postdot

	return p
}

func (s *PostdotContext) GetParser() antlr.Parser { return s.parser }

func (s *PostdotContext) Callinvoke() ICallinvokeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallinvokeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallinvokeContext)
}

func (s *PostdotContext) Fieldaccess() IFieldaccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldaccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldaccessContext)
}

func (s *PostdotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostdotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostdotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterPostdot(s)
	}
}

func (s *PostdotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitPostdot(s)
	}
}

func (p *PainlessParser) Postdot() (localctx IPostdotContext) {
	localctx = NewPostdotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PainlessParserRULE_postdot)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Callinvoke()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Fieldaccess()
		}

	}

	return localctx
}

// ICallinvokeContext is an interface to support dynamic dispatch.
type ICallinvokeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallinvokeContext differentiates from other interfaces.
	IsCallinvokeContext()
}

type CallinvokeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallinvokeContext() *CallinvokeContext {
	var p = new(CallinvokeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_callinvoke
	return p
}

func (*CallinvokeContext) IsCallinvokeContext() {}

func NewCallinvokeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallinvokeContext {
	var p = new(CallinvokeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_callinvoke

	return p
}

func (s *CallinvokeContext) GetParser() antlr.Parser { return s.parser }

func (s *CallinvokeContext) DOTID() antlr.TerminalNode {
	return s.GetToken(PainlessParserDOTID, 0)
}

func (s *CallinvokeContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CallinvokeContext) DOT() antlr.TerminalNode {
	return s.GetToken(PainlessParserDOT, 0)
}

func (s *CallinvokeContext) NSDOT() antlr.TerminalNode {
	return s.GetToken(PainlessParserNSDOT, 0)
}

func (s *CallinvokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallinvokeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallinvokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterCallinvoke(s)
	}
}

func (s *CallinvokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitCallinvoke(s)
	}
}

func (p *PainlessParser) Callinvoke() (localctx ICallinvokeContext) {
	localctx = NewCallinvokeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PainlessParserRULE_callinvoke)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PainlessParserDOT || _la == PainlessParserNSDOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(379)
		p.Match(PainlessParserDOTID)
	}
	{
		p.SetState(380)
		p.Arguments()
	}

	return localctx
}

// IFieldaccessContext is an interface to support dynamic dispatch.
type IFieldaccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldaccessContext differentiates from other interfaces.
	IsFieldaccessContext()
}

type FieldaccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldaccessContext() *FieldaccessContext {
	var p = new(FieldaccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_fieldaccess
	return p
}

func (*FieldaccessContext) IsFieldaccessContext() {}

func NewFieldaccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldaccessContext {
	var p = new(FieldaccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_fieldaccess

	return p
}

func (s *FieldaccessContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldaccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(PainlessParserDOT, 0)
}

func (s *FieldaccessContext) NSDOT() antlr.TerminalNode {
	return s.GetToken(PainlessParserNSDOT, 0)
}

func (s *FieldaccessContext) DOTID() antlr.TerminalNode {
	return s.GetToken(PainlessParserDOTID, 0)
}

func (s *FieldaccessContext) DOTINTEGER() antlr.TerminalNode {
	return s.GetToken(PainlessParserDOTINTEGER, 0)
}

func (s *FieldaccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldaccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldaccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterFieldaccess(s)
	}
}

func (s *FieldaccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitFieldaccess(s)
	}
}

func (p *PainlessParser) Fieldaccess() (localctx IFieldaccessContext) {
	localctx = NewFieldaccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PainlessParserRULE_fieldaccess)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PainlessParserDOT || _la == PainlessParserNSDOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(383)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PainlessParserDOTINTEGER || _la == PainlessParserDOTID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBraceaccessContext is an interface to support dynamic dispatch.
type IBraceaccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBraceaccessContext differentiates from other interfaces.
	IsBraceaccessContext()
}

type BraceaccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBraceaccessContext() *BraceaccessContext {
	var p = new(BraceaccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_braceaccess
	return p
}

func (*BraceaccessContext) IsBraceaccessContext() {}

func NewBraceaccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BraceaccessContext {
	var p = new(BraceaccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_braceaccess

	return p
}

func (s *BraceaccessContext) GetParser() antlr.Parser { return s.parser }

func (s *BraceaccessContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACE, 0)
}

func (s *BraceaccessContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BraceaccessContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACE, 0)
}

func (s *BraceaccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BraceaccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BraceaccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterBraceaccess(s)
	}
}

func (s *BraceaccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitBraceaccess(s)
	}
}

func (p *PainlessParser) Braceaccess() (localctx IBraceaccessContext) {
	localctx = NewBraceaccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PainlessParserRULE_braceaccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(PainlessParserLBRACE)
	}
	{
		p.SetState(386)
		p.expression(0)
	}
	{
		p.SetState(387)
		p.Match(PainlessParserRBRACE)
	}

	return localctx
}

// IStdarrayContext is an interface to support dynamic dispatch.
type IStdarrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStdarrayContext differentiates from other interfaces.
	IsStdarrayContext()
}

type StdarrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStdarrayContext() *StdarrayContext {
	var p = new(StdarrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_stdarray
	return p
}

func (*StdarrayContext) IsStdarrayContext() {}

func NewStdarrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StdarrayContext {
	var p = new(StdarrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_stdarray

	return p
}

func (s *StdarrayContext) GetParser() antlr.Parser { return s.parser }

func (s *StdarrayContext) CopyFrom(ctx *StdarrayContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StdarrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StdarrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NewstandardarrayContext struct {
	*StdarrayContext
}

func NewNewstandardarrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewstandardarrayContext {
	var p = new(NewstandardarrayContext)

	p.StdarrayContext = NewEmptyStdarrayContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StdarrayContext))

	return p
}

func (s *NewstandardarrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewstandardarrayContext) NEW() antlr.TerminalNode {
	return s.GetToken(PainlessParserNEW, 0)
}

func (s *NewstandardarrayContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PainlessParserTYPE, 0)
}

func (s *NewstandardarrayContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserLBRACE)
}

func (s *NewstandardarrayContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACE, i)
}

func (s *NewstandardarrayContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *NewstandardarrayContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NewstandardarrayContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserRBRACE)
}

func (s *NewstandardarrayContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACE, i)
}

func (s *NewstandardarrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterNewstandardarray(s)
	}
}

func (s *NewstandardarrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitNewstandardarray(s)
	}
}

func (p *PainlessParser) Stdarray() (localctx IStdarrayContext) {
	localctx = NewStdarrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PainlessParserRULE_stdarray)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	localctx = NewNewstandardarrayContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(PainlessParserNEW)
	}
	{
		p.SetState(390)
		p.Match(PainlessParserTYPE)
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(391)
				p.Match(PainlessParserLBRACE)
			}
			{
				p.SetState(392)
				p.expression(0)
			}
			{
				p.SetState(393)
				p.Match(PainlessParserRBRACE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IInitializedarrayContext is an interface to support dynamic dispatch.
type IInitializedarrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitializedarrayContext differentiates from other interfaces.
	IsInitializedarrayContext()
}

type InitializedarrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializedarrayContext() *InitializedarrayContext {
	var p = new(InitializedarrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_initializedarray
	return p
}

func (*InitializedarrayContext) IsInitializedarrayContext() {}

func NewInitializedarrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializedarrayContext {
	var p = new(InitializedarrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_initializedarray

	return p
}

func (s *InitializedarrayContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializedarrayContext) CopyFrom(ctx *InitializedarrayContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InitializedarrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializedarrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NewinitializedarrayContext struct {
	*InitializedarrayContext
}

func NewNewinitializedarrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewinitializedarrayContext {
	var p = new(NewinitializedarrayContext)

	p.InitializedarrayContext = NewEmptyInitializedarrayContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InitializedarrayContext))

	return p
}

func (s *NewinitializedarrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewinitializedarrayContext) NEW() antlr.TerminalNode {
	return s.GetToken(PainlessParserNEW, 0)
}

func (s *NewinitializedarrayContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PainlessParserTYPE, 0)
}

func (s *NewinitializedarrayContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACE, 0)
}

func (s *NewinitializedarrayContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACE, 0)
}

func (s *NewinitializedarrayContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACK, 0)
}

func (s *NewinitializedarrayContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACK, 0)
}

func (s *NewinitializedarrayContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *NewinitializedarrayContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NewinitializedarrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserCOMMA)
}

func (s *NewinitializedarrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserCOMMA, i)
}

func (s *NewinitializedarrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterNewinitializedarray(s)
	}
}

func (s *NewinitializedarrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitNewinitializedarray(s)
	}
}

func (p *PainlessParser) Initializedarray() (localctx IInitializedarrayContext) {
	localctx = NewInitializedarrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PainlessParserRULE_initializedarray)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewNewinitializedarrayContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(PainlessParserNEW)
	}
	{
		p.SetState(400)
		p.Match(PainlessParserTYPE)
	}
	{
		p.SetState(401)
		p.Match(PainlessParserLBRACE)
	}
	{
		p.SetState(402)
		p.Match(PainlessParserRBRACE)
	}
	{
		p.SetState(403)
		p.Match(PainlessParserLBRACK)
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(PainlessParserLBRACE-5))|(1<<(PainlessParserLP-5))|(1<<(PainlessParserNEW-5))|(1<<(PainlessParserBOOLNOT-5))|(1<<(PainlessParserBWNOT-5))|(1<<(PainlessParserADD-5))|(1<<(PainlessParserSUB-5)))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(PainlessParserINCR-58))|(1<<(PainlessParserDECR-58))|(1<<(PainlessParserOCTAL-58))|(1<<(PainlessParserHEX-58))|(1<<(PainlessParserINTEGER-58))|(1<<(PainlessParserDECIMAL-58))|(1<<(PainlessParserSTRING-58))|(1<<(PainlessParserREGEX-58))|(1<<(PainlessParserTRUE-58))|(1<<(PainlessParserFALSE-58))|(1<<(PainlessParserNULL-58))|(1<<(PainlessParserTYPE-58))|(1<<(PainlessParserID-58)))) != 0) {
		{
			p.SetState(404)
			p.expression(0)
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PainlessParserCOMMA {
			{
				p.SetState(405)
				p.Match(PainlessParserCOMMA)
			}
			{
				p.SetState(406)
				p.expression(0)
			}

			p.SetState(411)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(414)
		p.Match(PainlessParserRBRACK)
	}

	return localctx
}

// IListinitializerContext is an interface to support dynamic dispatch.
type IListinitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListinitializerContext differentiates from other interfaces.
	IsListinitializerContext()
}

type ListinitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListinitializerContext() *ListinitializerContext {
	var p = new(ListinitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_listinitializer
	return p
}

func (*ListinitializerContext) IsListinitializerContext() {}

func NewListinitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListinitializerContext {
	var p = new(ListinitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_listinitializer

	return p
}

func (s *ListinitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ListinitializerContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACE, 0)
}

func (s *ListinitializerContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ListinitializerContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListinitializerContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACE, 0)
}

func (s *ListinitializerContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserCOMMA)
}

func (s *ListinitializerContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserCOMMA, i)
}

func (s *ListinitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListinitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListinitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterListinitializer(s)
	}
}

func (s *ListinitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitListinitializer(s)
	}
}

func (p *PainlessParser) Listinitializer() (localctx IListinitializerContext) {
	localctx = NewListinitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PainlessParserRULE_listinitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Match(PainlessParserLBRACE)
		}
		{
			p.SetState(417)
			p.expression(0)
		}
		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PainlessParserCOMMA {
			{
				p.SetState(418)
				p.Match(PainlessParserCOMMA)
			}
			{
				p.SetState(419)
				p.expression(0)
			}

			p.SetState(424)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(425)
			p.Match(PainlessParserRBRACE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(427)
			p.Match(PainlessParserLBRACE)
		}
		{
			p.SetState(428)
			p.Match(PainlessParserRBRACE)
		}

	}

	return localctx
}

// IMapinitializerContext is an interface to support dynamic dispatch.
type IMapinitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapinitializerContext differentiates from other interfaces.
	IsMapinitializerContext()
}

type MapinitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapinitializerContext() *MapinitializerContext {
	var p = new(MapinitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_mapinitializer
	return p
}

func (*MapinitializerContext) IsMapinitializerContext() {}

func NewMapinitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapinitializerContext {
	var p = new(MapinitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_mapinitializer

	return p
}

func (s *MapinitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *MapinitializerContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserLBRACE, 0)
}

func (s *MapinitializerContext) AllMaptoken() []IMaptokenContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMaptokenContext)(nil)).Elem())
	var tst = make([]IMaptokenContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMaptokenContext)
		}
	}

	return tst
}

func (s *MapinitializerContext) Maptoken(i int) IMaptokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaptokenContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMaptokenContext)
}

func (s *MapinitializerContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PainlessParserRBRACE, 0)
}

func (s *MapinitializerContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserCOMMA)
}

func (s *MapinitializerContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserCOMMA, i)
}

func (s *MapinitializerContext) COLON() antlr.TerminalNode {
	return s.GetToken(PainlessParserCOLON, 0)
}

func (s *MapinitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapinitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapinitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterMapinitializer(s)
	}
}

func (s *MapinitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitMapinitializer(s)
	}
}

func (p *PainlessParser) Mapinitializer() (localctx IMapinitializerContext) {
	localctx = NewMapinitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PainlessParserRULE_mapinitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Match(PainlessParserLBRACE)
		}
		{
			p.SetState(432)
			p.Maptoken()
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PainlessParserCOMMA {
			{
				p.SetState(433)
				p.Match(PainlessParserCOMMA)
			}
			{
				p.SetState(434)
				p.Maptoken()
			}

			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(440)
			p.Match(PainlessParserRBRACE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Match(PainlessParserLBRACE)
		}
		{
			p.SetState(443)
			p.Match(PainlessParserCOLON)
		}
		{
			p.SetState(444)
			p.Match(PainlessParserRBRACE)
		}

	}

	return localctx
}

// IMaptokenContext is an interface to support dynamic dispatch.
type IMaptokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaptokenContext differentiates from other interfaces.
	IsMaptokenContext()
}

type MaptokenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaptokenContext() *MaptokenContext {
	var p = new(MaptokenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_maptoken
	return p
}

func (*MaptokenContext) IsMaptokenContext() {}

func NewMaptokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaptokenContext {
	var p = new(MaptokenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_maptoken

	return p
}

func (s *MaptokenContext) GetParser() antlr.Parser { return s.parser }

func (s *MaptokenContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MaptokenContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MaptokenContext) COLON() antlr.TerminalNode {
	return s.GetToken(PainlessParserCOLON, 0)
}

func (s *MaptokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaptokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaptokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterMaptoken(s)
	}
}

func (s *MaptokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitMaptoken(s)
	}
}

func (p *PainlessParser) Maptoken() (localctx IMaptokenContext) {
	localctx = NewMaptokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PainlessParserRULE_maptoken)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.expression(0)
	}
	{
		p.SetState(448)
		p.Match(PainlessParserCOLON)
	}
	{
		p.SetState(449)
		p.expression(0)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *ArgumentsContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *PainlessParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PainlessParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(PainlessParserLP)
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(PainlessParserLBRACE-5))|(1<<(PainlessParserLP-5))|(1<<(PainlessParserNEW-5))|(1<<(PainlessParserTHIS-5))|(1<<(PainlessParserBOOLNOT-5))|(1<<(PainlessParserBWNOT-5))|(1<<(PainlessParserADD-5))|(1<<(PainlessParserSUB-5)))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(PainlessParserINCR-58))|(1<<(PainlessParserDECR-58))|(1<<(PainlessParserOCTAL-58))|(1<<(PainlessParserHEX-58))|(1<<(PainlessParserINTEGER-58))|(1<<(PainlessParserDECIMAL-58))|(1<<(PainlessParserSTRING-58))|(1<<(PainlessParserREGEX-58))|(1<<(PainlessParserTRUE-58))|(1<<(PainlessParserFALSE-58))|(1<<(PainlessParserNULL-58))|(1<<(PainlessParserTYPE-58))|(1<<(PainlessParserID-58)))) != 0) {
		{
			p.SetState(452)
			p.Argument()
		}
		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PainlessParserCOMMA {
			{
				p.SetState(453)
				p.Match(PainlessParserCOMMA)
			}
			{
				p.SetState(454)
				p.Argument()
			}

			p.SetState(459)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(462)
		p.Match(PainlessParserRP)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) Lambda() ILambdaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *ArgumentContext) Funcref() IFuncrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncrefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncrefContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *PainlessParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PainlessParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.Lambda()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(466)
			p.Funcref()
		}

	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) ARROW() antlr.TerminalNode {
	return s.GetToken(PainlessParserARROW, 0)
}

func (s *LambdaContext) AllLamtype() []ILamtypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILamtypeContext)(nil)).Elem())
	var tst = make([]ILamtypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILamtypeContext)
		}
	}

	return tst
}

func (s *LambdaContext) Lamtype(i int) ILamtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILamtypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILamtypeContext)
}

func (s *LambdaContext) LP() antlr.TerminalNode {
	return s.GetToken(PainlessParserLP, 0)
}

func (s *LambdaContext) RP() antlr.TerminalNode {
	return s.GetToken(PainlessParserRP, 0)
}

func (s *LambdaContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LambdaContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LambdaContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserCOMMA)
}

func (s *LambdaContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserCOMMA, i)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (p *PainlessParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PainlessParserRULE_lambda)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(482)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PainlessParserTYPE, PainlessParserID:
		{
			p.SetState(469)
			p.Lamtype()
		}

	case PainlessParserLP:
		{
			p.SetState(470)
			p.Match(PainlessParserLP)
		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PainlessParserTYPE || _la == PainlessParserID {
			{
				p.SetState(471)
				p.Lamtype()
			}
			p.SetState(476)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == PainlessParserCOMMA {
				{
					p.SetState(472)
					p.Match(PainlessParserCOMMA)
				}
				{
					p.SetState(473)
					p.Lamtype()
				}

				p.SetState(478)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(481)
			p.Match(PainlessParserRP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(484)
		p.Match(PainlessParserARROW)
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PainlessParserLBRACK:
		{
			p.SetState(485)
			p.Block()
		}

	case PainlessParserLBRACE, PainlessParserLP, PainlessParserNEW, PainlessParserBOOLNOT, PainlessParserBWNOT, PainlessParserADD, PainlessParserSUB, PainlessParserINCR, PainlessParserDECR, PainlessParserOCTAL, PainlessParserHEX, PainlessParserINTEGER, PainlessParserDECIMAL, PainlessParserSTRING, PainlessParserREGEX, PainlessParserTRUE, PainlessParserFALSE, PainlessParserNULL, PainlessParserTYPE, PainlessParserID:
		{
			p.SetState(486)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILamtypeContext is an interface to support dynamic dispatch.
type ILamtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLamtypeContext differentiates from other interfaces.
	IsLamtypeContext()
}

type LamtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLamtypeContext() *LamtypeContext {
	var p = new(LamtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_lamtype
	return p
}

func (*LamtypeContext) IsLamtypeContext() {}

func NewLamtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LamtypeContext {
	var p = new(LamtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_lamtype

	return p
}

func (s *LamtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *LamtypeContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *LamtypeContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *LamtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LamtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LamtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterLamtype(s)
	}
}

func (s *LamtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitLamtype(s)
	}
}

func (p *PainlessParser) Lamtype() (localctx ILamtypeContext) {
	localctx = NewLamtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PainlessParserRULE_lamtype)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PainlessParserTYPE {
		{
			p.SetState(489)
			p.Decltype()
		}

	}
	{
		p.SetState(492)
		p.Match(PainlessParserID)
	}

	return localctx
}

// IFuncrefContext is an interface to support dynamic dispatch.
type IFuncrefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncrefContext differentiates from other interfaces.
	IsFuncrefContext()
}

type FuncrefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncrefContext() *FuncrefContext {
	var p = new(FuncrefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PainlessParserRULE_funcref
	return p
}

func (*FuncrefContext) IsFuncrefContext() {}

func NewFuncrefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncrefContext {
	var p = new(FuncrefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PainlessParserRULE_funcref

	return p
}

func (s *FuncrefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncrefContext) CopyFrom(ctx *FuncrefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FuncrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncrefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ClassfuncrefContext struct {
	*FuncrefContext
}

func NewClassfuncrefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassfuncrefContext {
	var p = new(ClassfuncrefContext)

	p.FuncrefContext = NewEmptyFuncrefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncrefContext))

	return p
}

func (s *ClassfuncrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassfuncrefContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PainlessParserTYPE, 0)
}

func (s *ClassfuncrefContext) REF() antlr.TerminalNode {
	return s.GetToken(PainlessParserREF, 0)
}

func (s *ClassfuncrefContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *ClassfuncrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterClassfuncref(s)
	}
}

func (s *ClassfuncrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitClassfuncref(s)
	}
}

type CapturingfuncrefContext struct {
	*FuncrefContext
}

func NewCapturingfuncrefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapturingfuncrefContext {
	var p = new(CapturingfuncrefContext)

	p.FuncrefContext = NewEmptyFuncrefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncrefContext))

	return p
}

func (s *CapturingfuncrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapturingfuncrefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PainlessParserID)
}

func (s *CapturingfuncrefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PainlessParserID, i)
}

func (s *CapturingfuncrefContext) REF() antlr.TerminalNode {
	return s.GetToken(PainlessParserREF, 0)
}

func (s *CapturingfuncrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterCapturingfuncref(s)
	}
}

func (s *CapturingfuncrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitCapturingfuncref(s)
	}
}

type ConstructorfuncrefContext struct {
	*FuncrefContext
}

func NewConstructorfuncrefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstructorfuncrefContext {
	var p = new(ConstructorfuncrefContext)

	p.FuncrefContext = NewEmptyFuncrefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncrefContext))

	return p
}

func (s *ConstructorfuncrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructorfuncrefContext) Decltype() IDecltypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecltypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecltypeContext)
}

func (s *ConstructorfuncrefContext) REF() antlr.TerminalNode {
	return s.GetToken(PainlessParserREF, 0)
}

func (s *ConstructorfuncrefContext) NEW() antlr.TerminalNode {
	return s.GetToken(PainlessParserNEW, 0)
}

func (s *ConstructorfuncrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterConstructorfuncref(s)
	}
}

func (s *ConstructorfuncrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitConstructorfuncref(s)
	}
}

type LocalfuncrefContext struct {
	*FuncrefContext
}

func NewLocalfuncrefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalfuncrefContext {
	var p = new(LocalfuncrefContext)

	p.FuncrefContext = NewEmptyFuncrefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncrefContext))

	return p
}

func (s *LocalfuncrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalfuncrefContext) THIS() antlr.TerminalNode {
	return s.GetToken(PainlessParserTHIS, 0)
}

func (s *LocalfuncrefContext) REF() antlr.TerminalNode {
	return s.GetToken(PainlessParserREF, 0)
}

func (s *LocalfuncrefContext) ID() antlr.TerminalNode {
	return s.GetToken(PainlessParserID, 0)
}

func (s *LocalfuncrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.EnterLocalfuncref(s)
	}
}

func (s *LocalfuncrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PainlessParserListener); ok {
		listenerT.ExitLocalfuncref(s)
	}
}

func (p *PainlessParser) Funcref() (localctx IFuncrefContext) {
	localctx = NewFuncrefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PainlessParserRULE_funcref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		localctx = NewClassfuncrefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(494)
			p.Match(PainlessParserTYPE)
		}
		{
			p.SetState(495)
			p.Match(PainlessParserREF)
		}
		{
			p.SetState(496)
			p.Match(PainlessParserID)
		}

	case 2:
		localctx = NewConstructorfuncrefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(497)
			p.Decltype()
		}
		{
			p.SetState(498)
			p.Match(PainlessParserREF)
		}
		{
			p.SetState(499)
			p.Match(PainlessParserNEW)
		}

	case 3:
		localctx = NewCapturingfuncrefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(501)
			p.Match(PainlessParserID)
		}
		{
			p.SetState(502)
			p.Match(PainlessParserREF)
		}
		{
			p.SetState(503)
			p.Match(PainlessParserID)
		}

	case 4:
		localctx = NewLocalfuncrefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(504)
			p.Match(PainlessParserTHIS)
		}
		{
			p.SetState(505)
			p.Match(PainlessParserREF)
		}
		{
			p.SetState(506)
			p.Match(PainlessParserID)
		}

	}

	return localctx
}

func (p *PainlessParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *RstatementContext = nil
		if localctx != nil {
			t = localctx.(*RstatementContext)
		}
		return p.Rstatement_Sempred(t, predIndex)

	case 15:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PainlessParser) Rstatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.GetTokenStream().LA(1) != PainlessParserELSE

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PainlessParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
