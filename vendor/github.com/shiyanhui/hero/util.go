package hero

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	escapedKeys   = []byte{'&', '\'', '<', '>', '"'}
	escapedValues = []string{"&amp;", "&#39;", "&lt;", "&gt;", "&#34;"}
)

// EscapeHTML escapes the html and then put it to the buffer.
func EscapeHTML(html string, buffer *bytes.Buffer) {
	var i, j, k int

	for i < len(html) {
		for j = i; j < len(html); j++ {
			k = bytes.IndexByte(escapedKeys, html[j])
			if k != -1 {
				break
			}
		}

		buffer.WriteString(html[i:j])
		if k != -1 {
			buffer.WriteString(escapedValues[k])
		}
		i = j + 1
	}
}

// FormatUint formats uint to string and put it to the buffer.
// It's part of go source:
// https://github.com/golang/go/blob/master/src/strconv/itoa.go#L60
func FormatUint(u uint64, buffer *bytes.Buffer) {
	var a [64 + 1]byte
	i := len(a)

	if ^uintptr(0)>>32 == 0 {
		for u > uint64(^uintptr(0)) {
			q := u / 1e9
			us := uintptr(u - q*1e9)
			for j := 9; j > 0; j-- {
				i--
				qs := us / 10
				a[i] = byte(us - qs*10 + '0')
				us = qs
			}
			u = q
		}
	}

	us := uintptr(u)
	for us >= 10 {
		i--
		q := us / 10
		a[i] = byte(us - q*10 + '0')
		us = q
	}

	i--
	a[i] = byte(us + '0')
	buffer.Write(a[i:])
}

// FormatInt format int to string and then put the result to the buffer.
func FormatInt(i int64, buffer *bytes.Buffer) {
	if i < 0 {
		buffer.WriteByte('-')
		i = -i
	}
	FormatUint(uint64(i), buffer)
}

// FormatFloat format float64 to string and then put the result to the buffer.
func FormatFloat(f float64, buffer *bytes.Buffer) {
	buffer.WriteString(strconv.FormatFloat(f, 'f', -1, 64))
}

// FormatBool format bool to string and then put the result to the buffer.
func FormatBool(b bool, buffer *bytes.Buffer) {
	if b {
		buffer.WriteString("true")
		return
	}
	buffer.WriteString("false")
}

// execCommand wraps exec.Command
func execCommand(command string) {
	parts := strings.Split(command, " ")
	if len(parts) == 0 {
		return
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
