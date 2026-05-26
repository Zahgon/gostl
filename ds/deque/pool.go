package deque

// Pool is a memory pool for holding Segments
type Pool[T any] struct {
	segs []*Segment[T]
}

func newPool[T any]() *Pool[T] { _ = "STUB: not implemented"; return nil }

func (p *Pool[T]) get() *Segment[T] { _ = "STUB: not implemented"; return nil }

func (p *Pool[T]) put(s *Segment[T]) { _ = "STUB: not implemented"; return }

func (p *Pool[T]) shrinkToSize(size int) { _ = "STUB: not implemented"; return }

func (p *Pool[T]) size() int { _ = "STUB: not implemented"; return 0 }
