package ws2tcp

import (
	"io"
	"net"
	"net/http"

	log "yangtizi/log/zaplog"

	"golang.org/x/net/websocket"
)

// TWebsocketToTcpsocket ws到tcps 的转换
type TWebsocketToTcpsocket struct {
	tcpAddress string
	binaryMode bool
	port       string
	certFile   string
	keyFile    string
}

// Start 开启
func (self *TWebsocketToTcpsocket) Start() {
	if len(self.tcpAddress) <= 0 {
		log.Errorf("没有选择地址")
		return
	}

	log.Infof("开始监听端口[%s]", self.port)

	http.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		self.relayHandler(ws)
	}))

	var err error
	if len(self.certFile) > 0 && len(self.keyFile) > 0 {
		err = http.ListenAndServeTLS(self.port, self.certFile, self.keyFile, nil)
	} else {
		err = http.ListenAndServe(self.port, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (self *TWebsocketToTcpsocket) copyWorker(dst io.Writer, src io.Reader, doneCh chan<- bool) {
	io.Copy(dst, src)
	doneCh <- true
}

func (self *TWebsocketToTcpsocket) relayHandler(ws *websocket.Conn) {
	conn, err := net.Dial("tcp", self.tcpAddress)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	if self.binaryMode {
		ws.PayloadType = websocket.BinaryFrame
	}

	doneCh := make(chan bool)

	go self.copyWorker(conn, ws, doneCh)
	go self.copyWorker(ws, conn, doneCh)

	<-doneCh
	conn.Close()
	ws.Close()
	<-doneCh
}

// func usage() {
// 	fmt.Fprintf(os.Stderr, "Usage: %s <tcpTargetAddress>\n", os.Args[0])
// 	flag.PrintDefaults()
// }

// func main() {

// 	flag.UintVar(&port, "p", 4223, "The port to listen on")
// 	flag.UintVar(&port, "port", 4223, "The port to listen on")
// 	flag.StringVar(&certFile, "tlscert", "", "TLS cert file path")
// 	flag.StringVar(&keyFile, "tlskey", "", "TLS key file path")
// 	flag.BoolVar(&binaryMode, "b", false, "Use binary frames instead of text frames")
// 	flag.BoolVar(&binaryMode, "binary", false, "Use binary frames instead of text frames")
// 	flag.Usage = usage
// 	flag.Parse()

// 	tcpAddress = flag.Arg(0)

// }
