package set_test

import (
	"encoding/json"
	"testing"

	"github.com/Mungrel/set"
)

func TestAdd(t *testing.T) {
	s := set.New(1, 2, 3)
	s.Add(3, 3, 3, 4)

	if len(s) != 4 {
		t.Fatalf("expected set to have len 4, got %d", len(s))
	}

	if !s.ContainsAll(1, 2, 3, 4) {
		t.Fatalf("expected set to contain [1,2,3,4]\ngot: %v", s)
	}
}

func TestRemove(t *testing.T) {
	s := set.New(1, 2, 3)
	s.Remove(3)

	if len(s) != 2 {
		t.Fatalf("expected set to have len 2, got %d", len(s))
	}

	if !s.ContainsAll(1, 2) {
		t.Fatalf("expected set to contain [1,2]\ngot: %v", s)
	}
}

func TestContainsAny(t *testing.T) {
	s := set.New(1, 2, 3)

	if !s.ContainsAny(3, 4, 5) {
		t.Fatal("expected set to contain any of [3,4,5]")
	}

	if s.ContainsAny(999) {
		t.Fatal("did not expect set to contain 999")
	}
}

func TestEquals(t *testing.T) {
	s1 := set.New(1, 2, 3)
	s2 := set.New(1, 2, 3)
	s3 := set.New(4, 5, 6)

	if !s1.Equals(s2) {
		t.Fatalf("expected %v to equal %v", s1, s2)
	}

	if s1.Equals(s3) {
		t.Fatalf("did not expect %v to equal %v", s1, s3)
	}

	s4 := set.Set[int](nil)
	s5 := set.Set[int](nil)

	if s1.Equals(s4) {
		t.Fatalf("did not expect %v to equal %v", s1, s4)
	}

	if !s4.Equals(s5) {
		t.Fatalf("nil sets should be considered equal")
	}
}

func TestJSONMarshal(t *testing.T) {
	s := set.New(1, 2)

	data, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

	actual := string(data)
	if actual != "[1,2]" && actual != "[2,1]" {
		t.Fatalf("unexpected JSON: %s", actual)
	}
}

func TestJSONUnmarshal(t *testing.T) {
	const data = "[1,2]"

	var s set.Set[int]
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		t.Fatal(err)
	}

	if len(s) != 2 || !s.ContainsAll(1, 2) {
		t.Fatalf("expected set to contain only [1,2]\nactual: %v", s)
	}
}
