package utils

// MinInt return the min of int
func MinInt(a int, args ...int) int {
	for i := range args {
		if args[i] < a {
			a = args[i]
		}
	}
	return a
}

// MaxInt return the max of int
func MaxInt(a int, args ...int) int {
	for i := range args {
		if args[i] > a {
			a = args[i]
		}
	}
	return a
}

// MinByte return the min of byte
func MinByte(a byte, args ...byte) byte {
	for i := range args {
		if args[i] < a {
			a = args[i]
		}
	}
	return a
}

// MaxByte return the max of Byte
func MaxByte(a byte, args ...byte) byte {
	for i := range args {
		if args[i] > a {
			a = args[i]
		}
	}
	return a
}
