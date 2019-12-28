package opslicemake

func newSlice() []int {
	arr := [3]int{1, 2, 3}
	slice := arr[0:1]
	return slice
}
