// Code generated by "stringer -type=ObjKind"; DO NOT EDIT.

package ast

import "strconv"

const _ObjKind_name = "BadObjTypObjVarObjFunObjFieldObj"

var _ObjKind_index = [...]uint8{0, 6, 12, 18, 24, 32}

func (i ObjKind) String() string {
	if i >= ObjKind(len(_ObjKind_index)-1) {
		return "ObjKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ObjKind_name[_ObjKind_index[i]:_ObjKind_index[i+1]]
}