package main

import (
	"fmt"
	_ "errors"
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

type greet interface{
	hello() string
}
type Chinese struct{

}
type American struct{

}
func (ch *Chinese) hello() string{
	return "你好"
}
func (am *American) hello() string{
	return "Hello"
}

func modify(s *string){
	*s = "j"
}
func main(){
	var s string = "hello"
	modify(&s)
	fmt.Print(s)
	// b := []byte{}
	// for i:=0;i<500;i++{
	// 	b = append(b, byte(i))
	// 	fmt.Printf("%c\n", b[i])
	// }
	// var p greet
	// p = &Chinese{}
	// fmt.Println(p.hello())
	// p = &American{}
	// fmt.Println(p.hello())	
	// s,ok := p.(*Chinese)
	// if ok{
	// 	fmt.Println(s.hello())
	// }
	// w := [4]byte{1,2,3,4}
	// for _, c := range w{
	// 	fmt.Println(string(c))
	// }
	// str := "hello"
	// fmt.Print(str[1:3])
	// str += "hel"
	// for _, c := range str{
	// 	fmt.Print()
	// }
	// errors.New("he")
}