package main

import (
	"fmt"
	"rsc.io/quote"
	"example/greetings"
	"log"
	_ "golang.org/x/example/hello/reverse"
)
func practice_16(){
	log.SetPrefix("greetings: ")
    log.SetFlags(0)
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
	greetingOfme, err := greetings.Hello("byBit")
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println(greetingOfme)
	}
	names := []string{
		"Amy","Tom","Jack","Musk",
	}
	grts, err := greetings.Hellos(names)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(grts)
	greetingOferr, err := greetings.Hello("")
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println(greetingOferr)
	}
}

const (
	a = iota
	b = iota
	c = iota
)
const (
	e = iota
	f
)
const (
	BIG = 1 << 50
	SMALL = BIG >> 49
)
type Car struct{
	price float32
	weight float32
}

func (car *Car) setCarPrice(p float32){
	car.price = p
}
func main(){
	// fmt.Println(reverse.String("Hello!"))
	// fmt.Println(reverse.Int(349))
	// fmt.Println(BIG, SMALL)
	// if n := 1; n > 0{
	// 	fmt.Println("Big")
	// }
	car := Car{1.2, 9.2}
	fmt.Println(car.price)
	car.setCarPrice(9.9)
	fmt.Println(car.price)
}