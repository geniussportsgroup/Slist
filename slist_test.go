package Slist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	l := New()

	assert.Equal(t, l.Size(), 0)

	l1 := New(1, 2, 3)

	assert.Equal(t, l1.Size(), 3)
}

func TestSlist_Swap(t *testing.T) {

	l0 := New()
	lpos := New(1, 2, 3, 4, 5)
	lneg := New(-5, -4, -3, -2, -1)

	result := l0.Swap(lpos) // test empty swap non-empty
	assert.Equal(t, l0.Size(), 5)
	assert.Equal(t, result, l0)
	assert.Equal(t, lpos.Size(), 0)

	result = l0.Swap(lpos) // test non-empty swap empty
	assert.Equal(t, l0.Size(), 0)
	assert.Equal(t, lpos.Size(), 5)
	assert.Equal(t, result, l0)

	result = lpos.Swap(lneg)
	assert.Equal(t, lpos.Size(), 5)
	assert.Equal(t, lneg.Size(), 5)
	assert.Equal(t, result, lpos)

	assert.True(t, lpos.Traverse(func(key interface{}) bool {
		return key.(int) >= -5 && key.(int) <= -1
	}))

	assert.True(t, lneg.Traverse(func(key interface{}) bool {
		return key.(int) >= 1 && key.(int) <= 5
	}))
}

func Test_misc(t *testing.T) {

	l := New(1, 2, 3)

	assert.Equal(t, l.First().(int), 1)
	assert.Equal(t, l.Last().(int), 3)
}

func TestSlist_RemoveFirst(t *testing.T) {

	assert.Nil(t, New().RemoveFirst())

	l := New(1, 2, 3)

	for i := 1; !l.IsEmpty(); i++ {
		item := l.RemoveFirst()
		assert.NotNil(t, item)
		assert.Equal(t, item.(int), i)
		assert.LessOrEqual(t, i, 3)
	}
}

func TestSlist_ToSlice(t *testing.T) {

	l := New(1, 2, 3)
	s := l.ToSlice()

	for i, it := 0, NewIterator(l); it.HasCurr(); it.Next() {
		assert.Equal(t, it.GetCurr().(int), s[i].(int))
		fmt.Printf("%d == %d\n", it.GetCurr().(int), s[i].(int))
		i++
	}
}

func TestSlist_clone(t *testing.T) {

	l1 := New(1, 2, 3)
	l2 := l1.clone()
	for it1, it2 := NewIterator(l1), NewIterator(l2); it1.HasCurr() && it2.HasCurr(); {
		assert.Equal(t, it1.GetCurr().(int), it2.GetCurr().(int))
		it1.Next()
		it2.Next()
	}
}

func TestSlist_ReverseInPlace(t *testing.T) {

	l := New(1, 2, 3)

	l.ReverseInPlace()
	assert.Equal(t, l.Size(), 3)
	assert.Equal(t, l.First(), 3)
	assert.Equal(t, l.Last(), 1)
}

func TestSlist_Reverse(t *testing.T) {

	l := New(1, 2, 3)
	r := l.Reverse()

	assert.Equal(t, l.Size(), 3)
	assert.Equal(t, l.First(), 1)
	assert.Equal(t, l.Last(), 3)

	assert.Equal(t, r.Size(), 3)
	assert.Equal(t, r.First(), 3)
	assert.Equal(t, r.Last(), 1)
}

func TestSlist_RotateLeftInPlace(t *testing.T) {

	l := New(1, 2, 3, 4, 5)
	lp := l.RotateLeft(2)
	l.RotateLeftInPlace(2)

	for it1, it2 := NewIterator(l), NewIterator(lp); it1.HasCurr() && it2.HasCurr(); {
		assert.Equal(t, it1.GetCurr().(int), it2.GetCurr().(int))
		it1.Next()
		it2.Next()
	}
}
