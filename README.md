# domz1

Implementation of data caching. For using doing

go get -u github.com/Hrukem/domz1


Then:

    package main
    
    import (
        "fmt"
        "github.com/Hrukem/domz1"
    )

    func main() {
       c := domz1.NewCache()

       c.Set("1", 12345)
       c.Set("2", 54321)
       c.Set("3", "Golang")
     
       fmt.Println(c.Get("1"))
       fmt.Println(c.Get("2"))
       fmt.Println(c.Get("3"))
     
       fmt.Println("")

       c.Delete("2")

       fmt.Println(c.Get("1"))
       fmt.Println(c.Get("2"))
       fmt.Println(c.Get("3"))
    }
