package controller

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/roblesoft/stockapi/internal/helper"
)

func (s Server) Ping(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, H{"status": "OK"})
}

func (s Server) Square(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("number"))

	if err != nil {
		fmt.Fprintf(w, "parameter is not a number")
	}

	c := gen(number)
	out := sq(c)
	result := <-out

	fmt.Fprintf(w, strconv.Itoa(result))
}

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}

		close(out)
	}()
	return out
}
