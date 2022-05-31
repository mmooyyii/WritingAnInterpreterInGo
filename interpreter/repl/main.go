package main

import (
	"WritingAnInterpreterInGo/interpreter/lex"
	"bufio"
	"fmt"
	"os"
)

type Repl struct {
}

const Prompt = ">> "

func (r *Repl) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	var cin string
	for {
		fmt.Printf(Prompt)
		if scanner.Scan() {
			cin = scanner.Text()
			fmt.Printf("%+v\n", lex.New(cin).Scan())
		} else {
			return
		}
	}
}

func main() {
	r := Repl{}
	fmt.Println("Monkey 0.0.1 on linux")
	r.Start()
}
