package types

func PromoteAdd(a, b Type) Type {
	if a == String || b == String {
		return String
	}
	return PromoteNumbers(a, b)
}

func PromoteNumbers(a, b Type) Type {
	return PromoteNumeric(a, b, true)
}

func PromoteNumeric(a, b Type, decimal bool) Type {
	if isOne(Def, a, b) {
		return Def
	}

	if decimal {
		if isOne(Double, a, b) {
			return Double
		}
		if isOne(Float, a, b) {
			return Float
		}
	}

	if isOne(Long, a, b) {
		return Long
	}

	if a.Extends() == Numeric && b.Extends() == Numeric {
		return Int
	}
	return nil
}

func isOne(check, a, b Type) bool {
	return a == check || b == check
}
