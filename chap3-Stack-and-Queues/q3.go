package chap3

type SetofStacks struct {
	Stacks      []*Stack
	numOfStacks int
}

func (ss *SetofStacks) Push(d int) {
	if ss.numOfStacks == 0 {
		ss.Stacks = append(ss.Stacks, &Stack{})
		ss.numOfStacks++
	}
	ss.Stacks[ss.numOfStacks-1].Push(d)
	if ss.Stacks[ss.numOfStacks-1].len == 10 {
		ss.Stacks = append(ss.Stacks, &Stack{})
		ss.numOfStacks++
	}
}

func (ss *SetofStacks) Pop() int {
	x := ss.Stacks[ss.numOfStacks-1].Pop()
	if ss.Stacks[ss.numOfStacks-1].len == 0 {
		ss.numOfStacks--
		ss.Stacks = ss.Stacks[:ss.numOfStacks]
	}
	return x
}
