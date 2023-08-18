package list_test

import (
	"testing"

	"github.com/hymkor/go-generics-list"
)

func makeData(base string) list.List[rune] { // use PushBack
	L := list.New[rune]()
	for _, r := range base {
		L.PushBack(r)
	}
	return L
}

func TestFront(t *testing.T) {
	L := makeData("123")
	p := L.Front()
	if p.IsNone() || p.Value() != '1' {
		t.Fail()
	}
}

func TestBack(t *testing.T) {
	L := makeData("123")
	p := L.Back()
	if p.IsNone() || p.Value() != '3' {
		t.Fail()
	}
}

func TestNext(t *testing.T) {
	L := makeData("123")

	p := L.Front()
	if p.IsNone() {
		t.Fail()
	}
	p = p.Next()
	if p.IsNone() || p.Value() != '2' {
		t.Fail()
	}
}

func TestPrev(t *testing.T) {
	L := makeData("123")

	p := L.Back()
	if p.IsNone() {
		t.Fail()
	}
	p = p.Prev()
	if p.IsNone() || p.Value() != '2' {
		t.Fail()
	}
}

func compareData(L list.List[rune], expect string) bool {
	p := L.Front()
	for _, r := range expect {
		if p.IsNone() || p.Value() != r {
			return false
		}
		p = p.Next()
	}
	return p.IsNone()
}

func TestInsertBeforeAndAfter(t *testing.T) {
	L := list.New[rune]()
	m := L.PushBack('M')
	L.InsertAfter('N', m)
	L.InsertBefore('L', m)

	if !compareData(L, "LMN") {
		t.Fail()
	}
}

func TestLen(t *testing.T) {
	L := makeData("123")
	if L.Len() != 3 {
		t.Fail()
	}
}

func TestMoveAfter(t *testing.T) {
	L := makeData("XYZ")
	z := L.Back()
	x := L.Front()
	L.MoveAfter(x, z)
	if !compareData(L, "YZX") {
		t.Fail()
	}
}

func TestMoveBefore(t *testing.T) {
	L := makeData("XYZ")
	z := L.Back()
	x := L.Front()
	L.MoveBefore(x, z)
	if !compareData(L, "YXZ") {
		t.Fail()
	}
}

func TestMoveToBack(t *testing.T) {
	L := makeData("ABC")
	L.MoveToBack(L.Front())
	if !compareData(L, "BCA") {
		t.Fail()
	}
}

func TestMoveToFront(t *testing.T) {
	L := makeData("ABC")
	L.MoveToFront(L.Back())
	if !compareData(L, "CAB") {
		t.Fail()
	}
}

func TestPushBackList(t *testing.T) {
	L := makeData("ABC")
	M := makeData("DEF")
	L.PushBackList(M)
	if !compareData(L, "ABCDEF") {
		t.Fail()
	}
}

func TestPushFront(t *testing.T) {
	L := makeData("ABC")
	L.PushFront('@')
	if !compareData(L, "@ABC") {
		t.Fail()
	}
}

func TestPushFrontList(t *testing.T) {
	L := makeData("ABC")
	M := makeData("DEF")
	L.PushFrontList(M)
	if !compareData(L, "DEFABC") {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	L := makeData("ABC")
	L.Remove(L.Front().Next())
	if !compareData(L, "AC") {
		t.Fail()
	}
}