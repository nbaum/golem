package main

import (
	"bufio"
	"fmt"
	"github.com/nbaum/golem/core"
	"io"
	"os"
)

func main() {
	s := core.NewScanner(bufio.NewReader(os.Stdin))
	env := core.NewEnv(core.NewDefaultEnv())
	for {
		form, err := s.ReadForm()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = env.Eval(form)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
