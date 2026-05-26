package comparator

type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

// Comparator Should return a number:
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
type Comparator[T any] func(a, b T) int

func OrderedTypeCmp[T Ordered](a, b T) int { _ = "STUB: not implemented"; return 0 }

// Reverse returns a comparator reverse to cmp
func Reverse[T any](cmp Comparator[T]) Comparator[T] { _ = "STUB: not implemented"; return nil }

// IntComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func IntComparator(a, b int) int { _ = "STUB: not implemented"; return 0 }

// UintComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func UintComparator(a, b uint) int { _ = "STUB: not implemented"; return 0 }

// Int8Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int8Comparator(a, b int8) int { _ = "STUB: not implemented"; return 0 }

// Uint8Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint8Comparator(a, b uint8) int { _ = "STUB: not implemented"; return 0 }

// Int16Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int16Comparator(a, b int16) int { _ = "STUB: not implemented"; return 0 }

// Uint16Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint16Comparator(a, b uint16) int { _ = "STUB: not implemented"; return 0 }

// Int32Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int32Comparator(a, b int32) int { _ = "STUB: not implemented"; return 0 }

// Uint32Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint32Comparator(a, b uint32) int { _ = "STUB: not implemented"; return 0 }

// Int64Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Int64Comparator(a, b int64) int { _ = "STUB: not implemented"; return 0 }

// Uint64Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Uint64Comparator(a, b uint64) int { _ = "STUB: not implemented"; return 0 }

// Float32Comparator compare a with b
//
//	-1 , if a < b or a is NaN and b is not NaN
//	0  , if a == b or a is NaN and b is NaN
//	1  , if a > b or a is not NaN and b is NaN
func Float32Comparator(a, b float32) int { _ = "STUB: not implemented"; return 0 }

// Float64Comparator compare a with b
//
//	-1 , if a < b or a is NaN and b is not NaN
//	0  , if a == b or a is NaN and b is NaN
//	1  , if a > b or a is not NaN and b is NaN
func Float64Comparator(a, b float64) int { _ = "STUB: not implemented"; return 0 }

// StringComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func StringComparator(a, b string) int { _ = "STUB: not implemented"; return 0 }

// UintptrComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func UintptrComparator(a, b uintptr) int { _ = "STUB: not implemented"; return 0 }

// BoolComparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func BoolComparator(a, b bool) int { _ = "STUB: not implemented"; return 0 }

// Complex64Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Complex64Comparator(a, b complex64) int { _ = "STUB: not implemented"; return 0 }

// Complex128Comparator compare a with b
//
//	-1 , if a < b
//	0  , if a == b
//	1  , if a > b
func Complex128Comparator(a, b complex128) int { _ = "STUB: not implemented"; return 0 }
