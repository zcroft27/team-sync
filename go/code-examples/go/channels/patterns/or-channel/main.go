package main

func main() {
	var or func(channels ...<-chan any) <-chan any // a function that takes a variadic
	//												  slice of read-only channels of type any,
	// 												  returning a single channel that will close
	// 												  when the first value is received.

	or = func(channels ...<-chan any) <-chan any {
		switch len(channels) {
		// First two base cases: if we have 1 channel, just return it.
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		orDone := make(chan any)

		// Function returns when the switch case exits.
		// On return, the channel is closed in the defer statement.
		// Thus, mulitplexing n channels onto one notification-channel.
		go func() {
			defer close(orDone)
			switch len(channels) {
			// Final base case: if we have 2 channels, jsut return them.
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // ... is the spread operator, expanding the slice.
				}
			}
		}()
		return orDone
	}
}
