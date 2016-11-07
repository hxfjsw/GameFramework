package main

import (
	"log"
	"io/ioutil"
	"os"
	"github.com/robertkrimen/otto"
	"time"
	"strconv"
	"os/signal"
	"syscall"
	"github.com/Unknwon/goconfig"
)

func timer_tick(ttl time.Duration, fun string) {
	timer1 := time.NewTicker(time.Millisecond * ttl)
	for {
		select {
		case <-timer1.C:
			vm.Run(fun)
		}
	}
}

func timer_after(ttl time.Duration, fun string) {
	time.Sleep(time.Millisecond * ttl)
	vm.Run(fun)
}

var vm *otto.Otto;

var server *Server
var hall Room

func main() {

	config, _ := goconfig.LoadConfigFile("conf/game.ini")
	ip, _ := config.GetValue("net", "ip")
	port, _ := config.GetValue("net", "port")
	port_int, _ := strconv.Atoi(port)

	server = &Server{};
	go server.Run(ip, uint32(port_int))
	go signalListen();

	init_js();
	vm.Run(`onStart();`)

	log.Println(strconv.Itoa(os.Getpid()) + " start finished!")
	time.Sleep(time.Hour * 1)
}

func init_js() {
	log.Println("restart js vm")
	main_js := ""

	if data, err := ioutil.ReadFile("./scripts/main.js"); err != nil {
		log.Println("not find main.js")
		os.Exit(-1);
	} else {
		main_js = string(data);
	}

	vm = otto.New()
	vm.Set("timer_tick", func(call otto.FunctionCall) otto.Value {
		ttl, _ := strconv.Atoi(call.Argument(0).String())
		fun := call.Argument(1).String()
		go timer_tick(time.Duration(ttl), fun);
		return otto.Value{}
	})
	vm.Set("timer_after", func(call otto.FunctionCall) otto.Value {
		ttl, _ := strconv.Atoi(call.Argument(0).String())
		fun := call.Argument(1).String()
		go timer_after(time.Duration(ttl), fun);
		return otto.Value{}
	})
	vm.Set("log", func(call otto.FunctionCall) otto.Value {
		msg := call.Argument(0).String()
		log.Println(msg)
		return otto.Value{}
	})
	vm.Set("exit", func(call otto.FunctionCall) otto.Value {
		os.Exit(0)
		return otto.Value{}
	})
	vm.Set("sendToFd", func(call otto.FunctionCall) otto.Value {
		fd, _ := strconv.Atoi(call.Argument(0).String())
		msg := call.Argument(0).String()

		log.Println("SendTo:", fd, msg)
		hall.SendToFd(fd, msg)

		return otto.Value{}
	})
	vm.Set("fdToIp", func(call otto.FunctionCall) otto.Value {
		msg := call.Argument(0).String()
		log.Println(msg)
		return otto.Value{}
	})
	vm.Run(main_js);
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for {
		s := <-c
		log.Println("on signal:" + s.String())
		if s == syscall.SIGUSR2 {
			init_js();
		} else {
			vm.Run(`onShutdown();`)
		}
	}
}
