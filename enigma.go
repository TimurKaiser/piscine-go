package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	*c = ***a
	*d = *******c
	*b = ****d
	*a = *b

	Enigma(&a, &b, &c, &d) // switch les valeyurs par pointages
}
