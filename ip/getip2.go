package ip

// import (
// 	"fmt"
// 	"net"
// 	"strings"
// )
//
// func main() {
// 	conn, err := net.Dial("udp", "google.com:80")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer conn.Close()
// 	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
// }
