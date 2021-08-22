package example

type Calculator struct {
	A int
	B int
}

func (c *Calculator) Add1() int {
	return c.A + c.B
}

func (c *Calculator) Add2(option int) int {
	return c.Add1() + option
}
