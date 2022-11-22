package main

import "fmt"

type Bird interface { // we have Bird interface that implements
	fly() //  fly and makeSound methods
	makeSound()
}

type RubberToy interface { // we also have RubberToy interface that implements
	squeak() // squeak method. For some reason we want birds to
} // squeak too.

type Duck struct{} // impelents Bird interface

func (this Duck) fly() {
	fmt.Println("*duck flies*")
}

func (this Duck) makeSound() {
	fmt.Println("GAGAGAGAGAGAGAGA")
}

type RubberDuck struct{} // implements RubberToy interface

func (this RubberDuck) squeak() {
	fmt.Println("SQUEAK")
}

type BirdAdapter struct { // We implement adapter that stores
	bird Bird // object that implements Bird interface
}

func (this BirdAdapter) squeak() { // and invokes Bird`s makeSound
	this.bird.makeSound() // method when squeak is called
}

func NewBirdAdapter(bird Bird) *BirdAdapter {
	return &BirdAdapter{bird}
}

func main() {
	duck := Duck{}
	rubberDuck := RubberDuck{}
	birdAdapter := NewBirdAdapter(duck) // this adapter will store duck
	rubberDuck.squeak()                 // and will tell it to make sound
	birdAdapter.squeak()                // when it is requested to squeak
}
