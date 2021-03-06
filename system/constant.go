package system

import "github.com/nongfah/go-hook/pkg/types"

var scanCode2VkCodeMap = map[Code]types.VKCode{
	1:  27,
	10: 57,
	11: 48,
	12: 189,
	13: 187,
	14: 8,
	15: 9,
	16: 81,
	17: 87,
	18: 69,
	19: 82,
	2:  49,
	20: 84,
	21: 89,
	22: 85,
	23: 73,
	24: 79,
	25: 80,
	26: 219,
	27: 221,
	28: 13,
	29: 162,
	3:  50,
	30: 65,
	31: 83,
	32: 68,
	33: 70,
	34: 71,
	35: 72,
	36: 74,
	37: 75,
	38: 76,
	39: 186,
	4:  51,
	40: 222,
	41: 192,
	42: 160,
	43: 220,
	44: 90,
	45: 88,
	46: 67,
	47: 86,
	48: 66,
	49: 78,
	5:  52,
	50: 77,
	51: 188,
	52: 190,
	53: 191,
	54: 161,
	55: 106,
	56: 164,
	57: 32,
	58: 20,
	59: 112,
	6:  53,
	60: 113,
	61: 114,
	62: 115,
	63: 116,
	64: 117,
	65: 118,
	66: 119,
	67: 120,
	68: 121,
	69: 144,
	7:  54,
	70: 145,
	71: 36,
	72: 38,
	73: 33,
	74: 109,
	75: 37,
	76: 12,
	77: 39,
	78: 107,
	79: 35,
	8:  55,
	80: 40,
	81: 34,
	82: 45,
	83: 46,
	87: 122,
	88: 123,
	9:  56,
	91: 241,
	93: 249,
}

var vkCode2ScanCodeMap = map[Code]types.VKCode{
	27:  1,
	57:  10,
	48:  11,
	189: 12,
	187: 13,
	8:   14,
	9:   15,
	81:  16,
	87:  17,
	69:  18,
	82:  19,
	49:  2,
	84:  20,
	89:  21,
	85:  22,
	73:  23,
	79:  24,
	80:  25,
	219: 26,
	221: 27,
	13:  28,
	162: 29,
	50:  3,
	65:  30,
	83:  31,
	68:  32,
	70:  33,
	71:  34,
	72:  35,
	74:  36,
	75:  37,
	76:  38,
	186: 39,
	51:  4,
	222: 40,
	192: 41,
	160: 42,
	220: 43,
	90:  44,
	88:  45,
	67:  46,
	86:  47,
	66:  48,
	78:  49,
	52:  5,
	77:  50,
	188: 51,
	190: 52,
	191: 53,
	161: 54,
	106: 55,
	164: 56,
	32:  57,
	20:  58,
	112: 59,
	53:  6,
	113: 60,
	114: 61,
	115: 62,
	116: 63,
	117: 64,
	118: 65,
	119: 66,
	120: 67,
	121: 68,
	144: 69,
	54:  7,
	145: 70,
	36:  71,
	38:  72,
	33:  73,
	109: 74,
	37:  75,
	12:  76,
	39:  77,
	107: 78,
	35:  79,
	55:  8,
	40:  80,
	34:  81,
	45:  82,
	46:  83,
	122: 87,
	123: 88,
	56:  9,
	241: 91,
	249: 93,
}
