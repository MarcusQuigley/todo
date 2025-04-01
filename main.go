package main

import(
    "fmt"
)

func main() {
    want := `[{"id":1,"description":"first todo","is_completed":true},{"id":2,"description":"second todo","is_completed":false}]`
    kant := "eded"
    fmt.Printf("%v %v\n", want, kant)

}