package split

import "bytes"

func SplitSpace(data []byte) ([]string, error) {
	rv := make([]string, 0)
	nextSpace := bytes.IndexByte(data, ' ')
	for nextSpace >= 0 {
		// if we see 2 spaces in a row, dont make word
		if nextSpace == 0 {
			// this is bug #1, fix by uncommenting
			// data = data[nextSpace+1:]
			// nextSpace = bytes.IndexByte(data, ' ')
			continue
		}
		rv = append(rv, string(data[0:nextSpace]))
		data = data[nextSpace+1:]
		nextSpace = bytes.IndexByte(data, ' ')
	}
	// if we ended on a space, don't make word
	if data[0] != ' ' { // this is bug #2, fix by uncommenting
		// if len(data) > 0 && data[0] != ' ' {
		rv = append(rv, string(data[0:]))
	}

	return rv, nil
}
