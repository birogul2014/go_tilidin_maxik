// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import "time"
import "fmt"

func main() {

   
    timer1 := time.NewTimer(time.Second * 5)

    // The `<-timer1.C` blocks on the timer's channel `C`
    // until it sends a value indicating that the timer
    // expired.
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // If you just wanted to wait, you could have used
    // `time.Sleep`. One reason a timer may be useful is
    // that you can cancel the timer before it expires.
    // Here's an example of that.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
