package chap3

import (
	"fmt"
	"testing"
)

func TestAnimalShelter(t *testing.T) {
	as := ASQueue{}
	as.Enqueue("Dog1", true)
	as.Enqueue("Cat1", false)
	as.Enqueue("Dog2", true)
	as.Enqueue("Dog3", true)
	as.Enqueue("Cat2", false)
	as.Enqueue("Cat3", false)
	as.Enqueue("Dog4", true)
	as.Enqueue("Cat4", false)
	as.Enqueue("Dog5", true)
	as.Enqueue("Dog6", true)
	as.Enqueue("Cat5", false)
	as.Enqueue("Cat6", false)
	fmt.Println("Get Any:", as.DequeueAny().name)
	fmt.Println("Get Any:", as.DequeueAny().name)
	fmt.Println("Get Dog:", as.DequeueDog().name)
	fmt.Println("Get Cat:", as.DequeueCat().name)
	fmt.Println("Get Dog:", as.DequeueDog().name)
	fmt.Println("Get Cat:", as.DequeueCat().name)
	fmt.Println("Get Dog:", as.DequeueDog().name)
	fmt.Println("Get Dog:", as.DequeueDog().name)
	fmt.Println("Get Cat:", as.DequeueCat().name)
	fmt.Println("Get Any:", as.DequeueAny().name)
	fmt.Println("Get Any:", as.DequeueAny().name)
	fmt.Println("Get Cat:", as.DequeueCat().name)
}
