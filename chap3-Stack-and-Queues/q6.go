package chap3

type Animal struct {
	name  string
	isDog bool
	next  *Animal
}
type ASQueue struct {
	last *Animal
	head *Animal
}

func (as *ASQueue) Enqueue(name string, isDog bool) {
	a := &Animal{name, isDog, nil}
	if as.last != nil {
		as.last.next = a
		as.last = a
		return
	}
	as.last = a
	as.head = a
}

func (as *ASQueue) DequeueAny() *Animal {
	if as.head == nil {
		return nil
	}

	t := as.head
	as.head = as.head.next
	t.next = nil
	return t
}

func (as *ASQueue) DequeueCat() *Animal {
	if as.head == nil {
		return nil
	}

	t := as.head
	var prev *Animal

	for t.isDog {
		prev = t
		t = t.next
	}
	if prev != nil {
		prev.next = t.next
	} else {
		as.head = t.next
	}
	t.next = nil
	return t
}
func (as *ASQueue) DequeueDog() *Animal {
	if as.head == nil {
		return nil
	}

	t := as.head
	var prev *Animal

	for !t.isDog {
		prev = t
		t = t.next
	}
	if prev != nil {
		prev.next = t.next
	} else {
		as.head = t.next
	}
	t.next = nil
	return t
}
