
```
package main

import
	(
		"syscall"
		"fmt"
	)

func main() {
	fd,_:= syscall.Open("foo.q",syscall.O_RDONLY|syscall.O_NONBLOCK, 0644)
	buf := make([]byte, 32)
	n, _ := syscall.Read(fd, buf)
	fmt.Println(string(buf[:n]))
	syscall.Close(fd)
}
```
