package main

import "plugin"

func main()  {
	p, err := plugin.Open("say_hi_plugin.so")
	if err != nil {
		panic(err)
	}

	h,err := p.Lookup("Hello")

	if err != nil {
		panic(err)
	}
	f,err := p.Lookup("SayFunc")

	if err != nil {
		panic(err)
	}

	*h.(*string) = "Hello"

	f.(func())()
}