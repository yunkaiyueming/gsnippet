package main

import "fmt"
import "time"

func main() {
	MyDate()
}

func MyDate() {
	p := fmt.Println
	now := time.Now()
	p(now)
	p(now.Date())
	p(now.Format("2006-01-02 15:04:05"))
}

func test() {
	now := time.Now()
	p := fmt.Println
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))
}
