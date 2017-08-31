package main

import "fmt"

const running = 1
const stopped = 2
const killed = 3

type allowedSChanges map[string]([]string)

func (hm *MyHash) UpdateState(hash string, newstate string, as allowedSChanges) bool {
	allowed := false
	for _, s := range as[hm.state] {
		if newstate == s {
			allowed = true
		}
	}
	if allowed {
		hm.state = newstate
		return true
	} else {
		return false
	}
}
func (hm *MyHash) AddHash(hash string) {
	hm.hmm[hash] = true
}

type MyHash struct {
	hmm   map[string]bool
	state string
}

func (hm *MyHash) GetHash(hash string) bool {
	_, ok := hm.hmm[hash]
	return ok
}

func main() {
	hm := MyHash{
		hmm: make(map[string]bool),
	}

	allowedStateChanges := allowedSChanges{}

	allowedStateChanges["running"] = []string{"stopped"}
	allowedStateChanges["stopped"] = []string{"killed"}
	allowedStateChanges["killed"] = []string{}

	/*	hm.AddHash("12345")
		fmt.Println(hm.GetHash("111111"))
		fmt.Println(hm.GetHash("12345"))
	*/
	hm.state = "running"
	fmt.Println(hm.UpdateState("12345", "killed", allowedStateChanges))
	fmt.Println(hm.UpdateState("12345", "stopped", allowedStateChanges))
}
