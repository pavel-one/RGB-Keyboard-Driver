package keyboard

type Key struct {
	Address []int
	Name    string
	Red     int
	Green   int
	Blue    int
}

func NewKey(name string, address ...int) *Key {
	return &Key{
		Address: address,
		Name:    name,
	}
}
