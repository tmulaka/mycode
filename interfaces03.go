/* Alta3 Research | RZFeeser
   Interface - Abstract class
   
   GoLang tries to distance itself from OOP terms, however, most of what we can do in OOP, we can do in GoLang. Our solution mimics an abstract class type. As with Java, our abstractAnimal type provides only a partial implementation of the Animal interface. A specialized animal type like Chicken inherits partial implementation from embedded abstract type and provides the remaining methods required by the animal interface. This is pretty much how abstract classes work in other languages that support OOP.
     */
   
package main
   
import (
    "fmt"
)

// define the interface
type Animal interface{
    Sound() string
    MakeSound()
}

// partially define the 
type abstractAnimal struct {Animal}

func (a abstractAnimal) MakeSound() {
    fmt.Printf("%v", a.Sound())
}


type Chicken struct {abstractAnimal}
func (c Chicken) Sound() string {
    return "bwaaak bwak bwak bwak\n"
}

func NewChicken() *Chicken {
    chicken := Chicken{abstractAnimal{}}
    chicken.abstractAnimal.Animal = chicken
    return &chicken
}

type Lion struct {abstractAnimal}
func (d Lion) Sound() string {
    return "RAAAAWWWWWRRRR\n"
}

func NewLion() *Lion{
    lion := Lion{abstractAnimal{}}
    lion.abstractAnimal.Animal = lion
    return &lion
}

func main(){
    c := NewChicken()
    c.MakeSound()
    d := NewLion()
    d.MakeSound()
}  

