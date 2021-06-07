package main

import "fmt"
import s "strings"
var p = fmt.Println

func main() {

	p("Contains: ", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("a-b-c-d", "-"))
	p("ToUpper:", s.ToUpper("abceDD"))
	p("ToLower", s.ToLower("abceDD"))
	p()

}
