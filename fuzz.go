// +build gofuzz

package split

func Fuzz(data []byte) int {

	if _, err := SplitSpace(data); err != nil {
		return 0
	} else {
		return 1
	}

}
