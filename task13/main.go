package main

func mathSwap(a, b *int) {
	*a -= *b
	*b += *a
	*a = *b - *a
}

func xorSwap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func main() {
	a, b := 5, 3
	println("before swap:", a, b)
	mathSwap(&a, &b)
	println("after math swap:", a, b)
	xorSwap(&a, &b)
	println("after XOR swap:", a, b)
}
