// Code generated by "stringer -type=BranchKind"; DO NOT EDIT.

package ast

import "strconv"

const _BranchKind_name = "BranchInvalidBranchBreakBranchContinue"

var _BranchKind_index = [...]uint8{0, 13, 24, 38}

func (i BranchKind) String() string {
	if i >= BranchKind(len(_BranchKind_index)-1) {
		return "BranchKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BranchKind_name[_BranchKind_index[i]:_BranchKind_index[i+1]]
}
