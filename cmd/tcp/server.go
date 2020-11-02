package tcp

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type User struct {
	ID int
	Addr string
	EnterAt time.Time
	MessageChannel chan string

}
func main()  {
	listener,err := net.Listen("tcp",":9999")

	if err != nil {
		panic(err)
	}

	go broadcaster()

	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	users := make(map[*User]struct{})
}

func handleConn(conn net.Conn) {
	user := &User{
		ID:             GenUserId(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string,8),
	}

	go sendMessage(conn,user.MessageChannel)

	user.MessageChannel <- "welcome," + user.String()
	messageChannel <- "user: '" + strconv.Itoa(user.ID) + "' has enter"
	enterChannel <- user

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChannel <- strconv.Itoa(user.ID) + ":" + input.Text()
	}

	if err := input.Err();err != nil {
		log.Println("读取错误",err)
	}

	leavingChannel <- user
	messgaeChannel <- "user: '" + strconv.Itoa(user.ID) + "' has left"




}

func sendMessage(conn net.Conn, ch <- chan string) {
	for msg := range ch{
		_, _ = fmt.Fprintln(conn, msg)
	}
}

func GenUserId() int {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	return x
}