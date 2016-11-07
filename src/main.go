package main

import (
	"io/ioutil"
	"os"
	"github.com/robertkrimen/otto"
	"time"
	"strconv"
	"os/signal"
	"syscall"
	"github.com/Unknwon/goconfig"
	"github.com/op/go-logging"
)

func timer_tick(ttl time.Duration, fun string) {
	timer1 := time.NewTicker(time.Millisecond * ttl)
	for {
		select {
		case <-timer1.C:
			if _, err := vm.Run(fun); err != nil {
				log.Error(err)
			}
		}
	}
}

func timer_after(ttl time.Duration, fun string) {
	time.Sleep(time.Millisecond * ttl)
	if _, err := vm.Run(fun); err != nil {
		log.Error(err)
	}
}

var vm *otto.Otto;

var server *Server
var hall Room

var log = logging.MustGetLogger("GameFramwork")
var log_format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {

	config, _ := goconfig.LoadConfigFile("conf/game.ini")
	ip, _ := config.GetValue("tcp", "ip")
	port, _ := config.GetValue("tcp", "port")
	port_int, _ := strconv.Atoi(port)

	log_file_path, _ := config.GetValue("log", "file")

	log_file, _ := os.OpenFile(log_file_path, os.O_WRONLY, 0666)
	backend1 := logging.NewLogBackend(log_file, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	logging.NewBackendFormatter(backend1, log_format)
	backend1Formatter := logging.NewBackendFormatter(backend2, log_format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend1Formatter)

	server = &Server{};
	go server.Run(ip, uint32(port_int))
	go signalListen();

	init_js();

	if _, err := vm.Run(`onStart();`); err != nil {
		log.Error(err)
		os.Exit(-2)
	}

	log.Info(strconv.Itoa(os.Getpid()) + " start finished!")
	time.Sleep(time.Hour * 1)
}

func init_js() {
	log.Info("restart js vm")
	main_js := ""

	if data, err := ioutil.ReadFile("./scripts/main.js"); err != nil {
		log.Info("not find main.js")
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
		log.Info(msg)
		return otto.Value{}
	})
	vm.Set("exit", func(call otto.FunctionCall) otto.Value {
		os.Exit(0)
		return otto.Value{}
	})
	vm.Set("sendToFd", func(call otto.FunctionCall) otto.Value {
		fd, _ := strconv.Atoi(call.Argument(0).String())
		msg := call.Argument(0).String()

		log.Info("SendTo:", fd, msg)
		hall.SendToFd(fd, msg)

		return otto.Value{}
	})
	vm.Set("fdToIp", func(call otto.FunctionCall) otto.Value {
		msg := call.Argument(0).String()
		log.Info(msg)
		return otto.Value{}
	})
	if _, err := vm.Run(main_js); err != nil {
		log.Error(err.Error())
	};
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for {
		s := <-c
		log.Info("on signal:" + s.String())
		if s == syscall.SIGUSR2 {
			init_js();
		} else {
			if _, err := vm.Run(`onShutdown();`); err != nil {
				log.Error(err)
			}
		}
	}
}
