package main

import (
	"github.com/robertkrimen/otto"
	"strconv"
	"time"
	"os"
	"io/ioutil"
)

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

	vm.Set("redis_get", func(call otto.FunctionCall) otto.Value {
		key := call.Argument(0).String()
		rst, _ := my_redis.get(key)
		result, _ := vm.ToValue(rst)
		return result
	})

	if err := vm.Set("redis_set", func(call otto.FunctionCall) otto.Value {
		key := call.Argument(0).String()
		value := call.Argument(1).String()
		rst, _ := my_redis.set(key, value)
		result, _ := vm.ToValue(rst)
		return result
	}); err != nil {
		log.Error(err.Error())
	}

	if _, err := vm.Run(main_js); err != nil {
		log.Error(err.Error())
	};
}
