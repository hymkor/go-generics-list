package list_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hymkor/go-generics-list"
)

func makeData(base string) *list.List[rune] { // use PushBack
	L := list.New[rune]()
	for _, r := range base {
		L.PushBack(r)
	}
	return L
}

func showData(t *testing.T, L *list.List[rune]) {
	var buffer strings.Builder
	buffer.WriteString("List={")
	for p := L.Front(); p != nil; p = p.Next() {
		fmt.Fprintf(&buffer, " '%c'", p.Value)
	}
	buffer.WriteString(" }")
	t.Log(buffer.String())
}

func TestFront(t *testing.T) {
	L := makeData("123")
	p := L.Front()
	if p == nil || p.Value != '1' {
		t.Fail()
	}
}

func TestBack(t *testing.T) {
	L := makeData("123")
	p := L.Back()
	if p == nil || p.Value != '3' {
		t.Fail()
	}
}

func TestNext(t *testing.T) {
	L := makeData("123")

	p := L.Front()
	if p == nil {
		t.Fail()
	}
	p = p.Next()
	if p == nil || p.Value != '2' {
		t.Fail()
	}
}

func TestPrev(t *testing.T) {
	L := makeData("123")

	p := L.Back()
	if p == nil {
		t.Fail()
	}
	p = p.Prev()
	if p == nil || p.Value != '2' {
		t.Fail()
	}
}

func compareData(L *list.List[rune], expect string) bool {
	p := L.Front()
	for _, r := range expect {
		if p == nil || p.Value != r {
			return false
		}
		p = p.Next()
	}
	if p != nil {
		return false
	}

	p = L.Back()
	for i := len(expect) - 1; i >= 0; i-- {
		if p == nil || p.Value != rune(expect[i]) {
			return false
		}
		p = p.Prev()
	}
	if p != nil {
		return false
	}
	return L.Len() == len(expect)
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
		t.Fatal("case 1")
	}

	L = makeData("123")
	L.MoveAfter(L.Front(), L.Front())
	if !compareData(L, "123") {
		showData(t, L)
		t.Fatal("case 2")
	}
	L.MoveAfter(L.Back(), L.Back())
	if !compareData(L, "123") {
		showData(t, L)
		t.Fatal("case 3")
	}
}

func TestMoveBefore(t *testing.T) {
	L := makeData("XYZ")
	z := L.Back()
	x := L.Front()
	L.MoveBefore(x, z)
	if !compareData(L, "YXZ") {
		t.Fail()
		t.Fatal("case 1")
	}

	L = makeData("123")
	L.MoveBefore(L.Front(), L.Front())
	if !compareData(L, "123") {
		showData(t, L)
		t.Fatal("case 2")
	}
	L.MoveBefore(L.Back(), L.Back())
	if !compareData(L, "123") {
		showData(t, L)
		t.Fatal("case 3")
	}

}

func TestMoveToBack(t *testing.T) {
	L := makeData("ABC")
	L.MoveToBack(L.Front())
	if !compareData(L, "BCA") {
		t.Fail()
	}

	L = makeData("A")
	L.MoveToBack(L.Front())
	if !compareData(L, "A") {
		t.Fail()
	}
}

func TestMoveToFront(t *testing.T) {
	L := makeData("ABC")
	L.MoveToFront(L.Back())
	if !compareData(L, "CAB") {
		t.Fail()
	}

	L = makeData("A")
	L.MoveToFront(L.Back())
	if !compareData(L, "A") {
		t.Fail()
	}
}

func TestPushBackList(t *testing.T) {
	L := makeData("ABC")
	M := makeData("DEF")
	L.PushBackList(M)
	if !compareData(L, "ABCDEF") {
		t.Fatal("case 1")
	}

	L = makeData("ABC")
	L.PushBackList(L)
	if !compareData(L, "ABCABC") {
		t.Fatal("case 2")
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
