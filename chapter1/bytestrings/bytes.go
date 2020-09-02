package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer will make use of the buffer created by the
// Buffer function
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array ❤️"
	reader := strings.NewReader(rawString)
	// we can also plug it into a scanner that allows buffered
	// reading and tokenzation
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// iterate over all of the scan events
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return scanner.Err()
}
