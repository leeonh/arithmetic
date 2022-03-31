package utils

type ByteSlice []byte

func (x ByteSlice) Len() int           { return len(x) }
func (x ByteSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x ByteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
