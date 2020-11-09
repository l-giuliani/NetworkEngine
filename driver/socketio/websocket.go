package socketio

import (
	"errors"
	"fmt"
	"github.com/googollee/go-socket.io"
	"net/http"
	"github.com/l-giuliani/networkEngine/driver/drivermodel"
	"github.com/l-giuliani/networkEngine/libs/communicationLibs/commInterfaces"
	"github.com/l-giuliani/networkEngine/config"
	"strconv"
)

type SocketIo struct {
	initialized bool
	serverPort uint16
	subscribers []commInterfaces.CommSubscriber
}

func NewSocketIo() drivermodel.CommunicationDriver{
	socketIo := new(SocketIo)
	socketIo.subscribers = make([]commInterfaces.CommSubscriber, 0)
	socketIo.initialized = false
	return socketIo
}

func (w *SocketIo) Init(config *config.Config) {
	w.serverPort = config.CommDriver.Params.ServerPort

	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	w.initialized = true

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	fmt.Println("Serving at localhost:"+strconv.Itoa(int(w.serverPort)))
	go http.ListenAndServe(":"+strconv.Itoa(int(w.serverPort)), nil)
}

func (w *SocketIo) Write() {fmt.Println("write")}
func (w *SocketIo) Read() {}

func (w *SocketIo) Subscribe(subscriber commInterfaces.CommSubscriber) error{
	if !w.initialized {
		return errors.New("CommDriver not initialized")
	}
	w.subscribers = append(w.subscribers, subscriber)
	return nil
}

