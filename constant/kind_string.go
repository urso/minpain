// Code generated by "stringer -type=Kind"; DO NOT EDIT.

package constant

import "strconv"

const _Kind_name = "BoolNumericDecimalString"

var _Kind_index = [...]uint8{0, 4, 11, 18, 24}

func (i Kind) String() string {
	if i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}