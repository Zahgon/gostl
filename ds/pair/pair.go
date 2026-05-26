package pair

type Pair struct {
	Front any
	Back  any
}

func MakePair(front any, back any) *Pair { _ = "STUB: not implemented"; return nil }

func (P *Pair) New(front any, back any) { _ = "STUB: not implemented"; return }

func (P *Pair) Equal(pair2 Pair) bool { _ = "STUB: not implemented"; return false }

func (P *Pair) Fronts() any { _ = "STUB: not implemented"; return *new(any) }

func (P *Pair) Backs() any { _ = "STUB: not implemented"; return *new(any) }
