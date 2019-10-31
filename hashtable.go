package ufap

// HashTable is HashTables
type HashTable struct {
	hashtable []*element
	length    int
}

type element struct {
	key int
	val interface{}
}

// NewHashTable generate HashTable
func NewHashTable() HashTable {
	length := 1046527
	return HashTable{make([]*element, length), length}
}

// Put new value
func (h *HashTable) Put(key int, val interface{}) {
	new := element{key, val}
	hash := key % h.length
	index := 0
	for {
		// double hash
		if h.hashtable[((1+hash%(h.length-1))*index+hash)%h.length] == nil {
			h.hashtable[((1+hash%(h.length-1))*index+hash)%h.length] = &new
			break
		} else {
			index++
		}
	}
}

// Find new value
func (h *HashTable) Find(key int) (val interface{}, ok bool) {
	hash := key % h.length
	index := 0
	for {
		target := h.hashtable[((1+hash%(h.length-1))*index+hash)%h.length]
		if target == nil {
			return nil, false
		}

		if target.key == key {
			return h.hashtable[((1+hash%(h.length-1))*index+hash)%h.length].val, true
		}

		index++
	}

}
