package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"1234567890", "Hello"}
	greeting.Greet(s, greeting.CreateCustomFunction("!!!"), true)
	greeting.TypeSwitchTest(1.234)
}
