package main

import(
	"fmt"
	"net"
	"time"
	"bufio"
	"os"
)

var arrConn []ConnName

type ConnName struct {
	Connection net.Conn
	Name string
	Joined bool
	Online bool
}

type ConnMessage struct {
	Conn net.Conn
	Message string
	Time time.Time
}

var prev []ConnMessage

func CreateServer(port string) {
	if port == "" {
		port = "2525"
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on the port :" + port)

	messages := make(chan ConnMessage)

	go SendMail(messages)
	for {
		c, cerr := ln.Accept()
		if cerr != nil {
			fmt.Println(cerr)
			continue
		}
		go handlerSreverConnection(c, messages)
	}
}

func handlerSreverConnection(c net.Conn, messages chan ConnMessage) {
	var newConn ConnName
	initial := true
	first := true
	var connM ConnMessage
	connM.Conn = c
	left := false
	for {
		if initial {
			linuxfile, er := os.Open("linux.txt")

			if er != nil {
				fmt.Println("error opening file : " + er.Error())
				return
			}
		
			scanner := bufio.NewScanner(linuxfile)
			fmt.Fprintf(c, "Welcome to TCP-Chat!\n")
			for scanner.Scan() {
				fmt.Fprintf(c, scanner.Text()+"\n")
				time.Sleep(time.Millisecond)
			}
			fmt.Fprintf(c, "[ENTER YOUR NAME]: ")
			time.Sleep(time.Millisecond)
			initial = false
			continue
		}
		if left {
			break
		}
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			for i, conn := range arrConn {
				if conn.Connection == c {
					arrConn[i].Joined = false
					arrConn[i].Online = false
					connM.Message = newConn.Name + " has left our chat...\n"
					messages<-connM
					left = true
				}
			}
			break
		}
		if first {
			if msg[:len(msg)-1] == "" {
				fmt.Fprintf(c, "[INVALID NAME ENTER AGAIN!] ")
			} else {
				newConn.Connection = c
				newConn.Name = msg[:len(msg)-1]
				newConn.Joined = false
				newConn.Online = true
				first = false
				arrConn = append(arrConn, newConn)
				go SendAllToConn(c)
				connM.Message = newConn.Name + " has joined our chat...\n"
				messages<-connM
			}
			
		} else {
			if msg == "-l\n" {
				fmt.Fprintf(c, "Currently in chat:\n")
				amount := 0
				for i, _ := range arrConn {
					if arrConn[i].Online == true {
						amount++
						fmt.Fprintf(c, arrConn[i].Name+"\n")
						time.Sleep(time.Millisecond)
					}
				}
				a := fmt.Sprintf("%d", amount)
				fmt.Fprintf(c, "total amount of users: "+a+"\n")
				fmt.Fprintf(c, PrintPreMsg(newConn.Name))
				
			} else {
				connM.Message = "["+newConn.Name+"]:"+msg
				t := time.Now()
				form := t.Format("2006 01 02")
				hour := fmt.Sprintf("%d", t.Hour())
				minute := fmt.Sprintf("%d", t.Minute())
				second := fmt.Sprintf("%d", t.Second())
				connM.Message = "["+form+" "+hour+":"+minute+":"+second+"]"+connM.Message
	
				messages<-connM
			}

		}
	}
}

func SendMail(messages <-chan ConnMessage) {
	for m := range messages {
		for i, c := range arrConn {
			if m.Conn == c.Connection {
				if c.Joined == false {
					arrConn[i].Joined = true
				} else {
					fmt.Fprintf(c.Connection, PrintPreMsg(c.Name))
				}
				continue
			}
			if c.Joined {
				fmt.Fprintf(c.Connection, "\n"+m.Message)
				time.Sleep(time.Millisecond)
				fmt.Fprintf(c.Connection, PrintPreMsg(c.Name))
			}
		}
		prev = append(prev, m)
	}
}


func SendAllToConn(c net.Conn) {
	name := ""
	for _, m := range prev {
		if m.Conn == c {
			continue
		}
		fmt.Fprintf(c, m.Message)
		time.Sleep(time.Millisecond)
	}

	for _, f := range arrConn {
		if f.Connection == c {
			name = f.Name
		}
	}
	fmt.Fprintf(c, PrintPreMsg(name))
	time.Sleep(time.Millisecond)
}

func PrintPreMsg(name string) string {
	t := time.Now()
	form := t.Format("2006 01 02")
	hour := fmt.Sprintf("%d", t.Hour())
	minute := fmt.Sprintf("%d", t.Minute())
	second := fmt.Sprintf("%d", t.Second())
	return "["+form+" "+hour+":"+minute+":"+second+"]"+"["+name+"]:"
}