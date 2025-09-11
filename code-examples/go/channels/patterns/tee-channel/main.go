package main

import "fmt"

func main() {
	tee := func(
		in <-chan any,
	) (<-chan any, <-chan any) {
		out1 := make(chan any)
		out2 := make(chan any)
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range in {
				var out1, out2 = out1, out2
				// Shadow these variables so we can temporarily set one to nil
				// so we can force a send down the other channel.
				for i := 0; i < 2; i++ {
					select {
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	in := make(chan any)
	out1, out2 := tee(in)
	go func() { in <- "hello" }()
	fmt.Println(<-out1)
	fmt.Println(<-out2)
}
