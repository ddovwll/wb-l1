package main

func swapBits(num int64, bitPos int, bitToSet int) {
	if bitPos < 1 || bitPos > 64 {
		println("bitPos out of range")
		return
	}

	switch bitToSet {
	case 0:
		num = num & ^int64(1<<(bitPos-1))
		println(num)
	case 1:
		num = num | int64(1<<(bitPos-1))
		println(num)
	default:
		println("bitToSet out of range")
	}
}

func main() {
	var num int64 = 5
	bitPos := 1
	bitToSet := 0

	swapBits(num, bitPos, bitToSet)
}
