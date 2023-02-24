package keyboard

type Key struct {
	Address []int
	Name    string
	Red     int
	Green   int
	Blue    int
}

func NewKey(name string, firstByte int) *Key {
	return &Key{
		Address: []int{
			firstByte,
			firstByte + 1,
			firstByte + 2,
		},
		Name: name,
	}
}

func (k *Key) Reset() {
	k.Red = 0
	k.Green = 0
	k.Blue = 0
}

func (k *Key) GetRedIndex() int {
	return k.Address[0] - 1
}

func (k *Key) GetGreenIndex() int {
	return k.Address[1] - 1
}

func (k *Key) GetBlueIndex() int {
	return k.Address[2] - 1
}

func (k *Key) GetRed() byte {
	return byte(k.Red)
}

func (k *Key) GetGreen() byte {
	return byte(k.Green)
}

func (k *Key) GetBlue() byte {
	return byte(k.Blue)
}
