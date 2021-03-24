package Slist

type Snode struct {
	item interface{}
	next *Snode
}

func NilSnode() *Snode {
	return nil
}

type Slist struct {
	head *Snode
	tail *Snode
}

// Swap in O(1) two sequences
func (seq *Slist) Swap(rhs interface{}) interface{} {

	other := rhs.(*Slist)
	seq.head, other.head = other.head, seq.head
	seq.tail, other.tail = other.tail, seq.tail

	return seq
}

// Create a new sequence with the received items
func New(items ...interface{}) *Slist {

	seq := new(Slist)
	for _, item := range items {
		seq.Append(item)
	}
	return seq
}

// Return true if sequence is empty
func (seq *Slist) IsEmpty() bool {

	return seq.head == nil && seq.tail == nil // double check!
}

// Return a slice with the elements of the list
func (seq *Slist) ToSlice() []interface{} {

	ret := make([]interface{}, 0, 4)
	for it := NewIterator(seq); it.HasCurr(); it.Next() {
		ret = append(ret, it.GetCurr())
	}

	return ret
}

func (seq *Slist) __append(item interface{}) *Slist {

	ptr := new(Snode)
	ptr.item = item
	if seq.IsEmpty() {
		seq.head = ptr
		seq.tail = ptr
		return seq
	}

	seq.tail.next = ptr
	seq.tail = ptr
	return seq
}

// Append the received items at the end of the sequence
func (seq *Slist) Append(item interface{}, items ...interface{}) interface{} {

	result := seq.__append(item)
	for _, i := range items {
		result.__append(i)
	}
	return result
}

func (seq *Slist) __insert(item interface{}) *Slist {

	ptr := new(Snode)
	ptr.item = item
	if seq.IsEmpty() {
		seq.head = ptr
		seq.tail = ptr
		return seq
	}

	ptr.next = seq.head
	seq.head = ptr
	return seq
}

// Insert at the beginning of the sequence all the received items (in the given order)
func (seq *Slist) Insert(item interface{}, items ...interface{}) *Slist {

	result := seq.__insert(item)
	for _, i := range items {
		result.__insert(i)
	}
	return result
}

// Remove the first item of the sequence
func (seq *Slist) RemoveFirst() interface{} {

	if seq.IsEmpty() {
		return nil
	}

	ret := seq.head.item
	seq.head = seq.head.next
	if seq.head == nil { // list became empty?
		seq.tail = nil
	}
	return ret
}

// Return the first element of the list
func (seq *Slist) First() interface{} {

	if seq == nil {
		return nil
	}
	return seq.head.item
}

// Return the last element of the list
func (seq *Slist) Last() interface{} {
	if seq == nil {
		return nil
	}
	return seq.tail.item
}

// Return true if the list is empty
func (seq *Slist) Empty() *Slist {
	seq.head = nil
	seq.tail = nil
	return seq
}

// Append in O(1) l to list_ptr and destroys l
func (seq *Slist) staticAppendList(l *Slist) *Slist {

	if l.IsEmpty() {
		return seq
	}

	if seq.IsEmpty() {
		return seq.Swap(l).(*Slist)
	}

	seq.tail.next = l.head
	seq.tail = l.tail
	l.head = nil
	l.tail = nil

	return seq
}

// Append to seq the received lists. Complexity O(n) where n is the number of received list
func (seq *Slist) AppendList(l *Slist, ln ...*Slist) *Slist {
	seq.staticAppendList(l)
	for _, ll := range ln {
		seq.staticAppendList(ll)
	}
	return seq
}

type Iterator struct {
	listPtr *Slist
	curr    *Snode
}

// Return an iterator to the list
func NewIterator(seq *Slist) *Iterator {
	it := new(Iterator)
	it.listPtr = seq
	it.curr = seq.head
	return it
}

// Return an iterator to the list
func (seq *Slist) CreateIterator() interface{} {

	return NewIterator(seq)
}

// Reset the iterator to the first element of the list
func (it *Iterator) ResetFirst() interface{} {

	it.curr = it.listPtr.head
	return it
}

// Return true if the iterator is positioned on a valid element
func (it *Iterator) HasCurr() bool {
	return it.curr != nil
}

// Return true if the current element of the list is the last of the list
func (it *Iterator) IsLast() bool {
	return it.curr == it.listPtr.tail
}

// Return the current element of the list
func (it *Iterator) GetCurr() interface{} {
	if it.curr == nil {
		return nil
	}
	return it.curr.item
}

// Advance the iterator to the next element of the list
func (it *Iterator) Next() interface{} {
	it.curr = it.curr.next
	if it.curr == nil {
		return nil
	}
	return it
}

// Return the number of elements of the list
func (seq *Slist) Size() int {

	n := 0
	for it := NewIterator(seq); it.HasCurr(); it.Next() {
		n++
	}
	return n
}

// Traverse the list and execute operation. It stops if the operation return false. Return true if
// all the elements of the list were traversed
func (seq *Slist) Traverse(operation func(key interface{}) bool) bool {

	for it := NewIterator(seq); it.HasCurr(); it.Next() {
		if !operation(it.GetCurr()) {
			return false
		}
	}

	return true
}

func (seq *Slist) clone() *Slist {
	ret := New()
	for it := NewIterator(seq); it.HasCurr(); it.Next() {
		ret.Append(it.GetCurr())
	}
	return ret
}

// Reverse the list in place
func (seq *Slist) ReverseInPlace() *Slist {

	tmp := New()

	for !seq.IsEmpty() {
		tmp.Insert(seq.RemoveFirst())
	}

	return seq.Swap(tmp).(*Slist)
}

// Return a reversed copy of seq
func (seq *Slist) Reverse() *Slist {
	return seq.clone().ReverseInPlace()
}

// Rotate in place n positions to left
func (seq *Slist) RotateLeftInPlace(n int) *Slist {

	if seq.IsEmpty() || n == 0 {
		return seq
	}

	for i := 0; i < n; i++ {
		seq.Append(seq.RemoveFirst())
	}

	return seq
}

// Return a copy of seq rotated n positions to left
func (seq *Slist) RotateLeft(n int) *Slist {
	return seq.clone().RotateLeftInPlace(n)
}
