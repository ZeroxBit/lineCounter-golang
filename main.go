package lineCounter

import (
	"bytes"
	"io"
)

// cuenta la cantidad de lineas
func lineCounter(r io.Reader) (count int, err error) {
	var readSize int

	buf := make([]byte, 1024)

	for {
		readSize, err = r.Read(buf)
		if err != nil {
			break
		}
	}

	var buffPosition int
	for {
		i := bytes.IndexByte(buf[buffPosition:], '\n')
		if i == -1 || readSize == buffPosition {
			break
		}
		buffPosition += i + 1
		count++
	}

	if readSize > 0 && count == 0 || count > 0 {
		count++
	}

	if err == io.EOF {
		err = nil
	}

	return
}
