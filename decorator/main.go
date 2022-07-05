// The objective is to add functionalities to an existing object without rewrite or alter exiting code and keeping these
// functionalities separate.
// These functionalities should be able to interact with existing structures.
// The solution is to embed the decorator object and provide additional functionalities

package main

import "fmt"

/*
type Bird struct {
  Age int
}

func (b *Bird) Fly() {
  if b.Age >= 10 {
    fmt.Println("Flying!")
  }
}

type Lizard struct {
  Age int
}

func (l *Lizard) Crawl() {
  if l.Age < 10 {
    fmt.Println("Crawling!")
  }
}

type Dragon struct {
  Bird
  Lizard
}
*/

type Aged interface {
  Age() int
  SetAge(age int)
}

type Bird struct {
  age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
  if b.age >= 10 {
    fmt.Println("Flying!")
  }
}

type Lizard struct {
  age int
}

func (l *Lizard) Age() int       { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
  if l.age < 10 {
    fmt.Println("Crawling!")
  }
}

type Dragon struct {
  bird   Bird
  lizard Lizard
}

func (d *Dragon) Age() int {
  return d.bird.age
}

func (d *Dragon) SetAge(age int) {
  d.bird.SetAge(age)
  d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
  d.bird.Fly()
}

func (d *Dragon) Crawl() {
  d.lizard.Crawl()
}

func NewDragon() *Dragon {
  return &Dragon{Bird{}, Lizard{}}
}

func main() {
  //d := Dragon{}
  //d.Bird.Age = 10
  //fmt.Println(d.Lizard.Age)
  //d.Fly()
  //d.Crawl()

  d := &Dragon{Bird{}, Lizard{}}
  d.SetAge(10)
  d.Fly()
  fmt.Println(d.Age())
  d.Crawl()
}
