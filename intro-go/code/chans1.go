var numbers = make(chan int)

numbers <- 1

number := <-numbers
fmt.Println(number)
