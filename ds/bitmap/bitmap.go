package bitmap

// Bitmap is a mapping from some domain (for example, a range of integers) to bits. It is also called a bit array or bitmap index
type Bitmap struct {
	data []byte
	size uint64 //bitmap's size in bit, is the multiple of 8
}

// New creates a new bitmap
func New(size uint64) *Bitmap { _ = "STUB: not implemented"; return nil }

// NewFromData creates a bitmap from the exported data
func NewFromData(data []byte) *Bitmap { _ = "STUB: not implemented"; return nil }

// Set sets 1 at position pos
func (b *Bitmap) Set(pos uint64) bool { _ = "STUB: not implemented"; return false }

// Unset sets 0 at position pos
func (b *Bitmap) Unset(pos uint64) bool { _ = "STUB: not implemented"; return false }

// IsSet returns true if the position pos is 1
func (b *Bitmap) IsSet(pos uint64) bool { _ = "STUB: not implemented"; return false }

// Resize resizes the bitmap with the passed size
func (b *Bitmap) Resize(size uint64) { _ = "STUB: not implemented"; return }

// Size returns the bitmap's size in bit
func (b *Bitmap) Size() uint64 {
	_ = "STUB: not implemented"

	// Clear clear the bitmap's data
	return 0
}

func (b *Bitmap) Clear() { _ = "STUB: not implemented"; return }

// Data returns the bitmap's internal data slice
func (b *Bitmap) Data() []byte { _ = "STUB: not implemented"; return nil }
