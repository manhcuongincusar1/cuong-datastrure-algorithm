package hashsets

type IHashSet[T comparable] interface {
	Add(values ...T)
	Remove(values ...T)
	Contains(value T) bool
	ContainsAll(values ...T) bool
	ContainsAny(values ...T) bool
	Merge(sets ...*HashSet[T])
	Clear()
	GetValues() []T
	IsEmpty() bool
	Size() int
	Copy() *HashSet[T]
	Union(ss *HashSet[T]) *HashSet[T]
	Intersection(ss *HashSet[T]) *HashSet[T]
	SymmetricDifference(ss *HashSet[T]) *HashSet[T]
	Subtraction(ss *HashSet[T]) *HashSet[T]
}

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable](values ...T) IHashSet[T] {
	set := HashSet[T]{data: make(map[T]struct{}, len(values))}
	set.Add(values...)
	return &set
}
func (p *HashSet[T]) Add(values ...T)
func (p *HashSet[T]) Remove(values ...T)
func (p *HashSet[T]) Contains(value T) bool
func (p *HashSet[T]) ContainsAll(values ...T) bool
func (p *HashSet[T]) ContainsAny(values ...T) bool
func (p *HashSet[T]) Merge(sets ...*HashSet[T])
func (p *HashSet[T]) Clear()
func (p *HashSet[T]) GetValues() []T
func (p *HashSet[T]) IsEmpty() bool
func (p *HashSet[T]) Size() int
func (p *HashSet[T]) Copy() *HashSet[T]
func (p *HashSet[T]) Union(ss *HashSet[T]) *HashSet[T]
func (p *HashSet[T]) Intersection(ss *HashSet[T]) *HashSet[T]
func (p *HashSet[T]) SymmetricDifference(ss *HashSet[T]) *HashSet[T]
func (p *HashSet[T]) Subtraction(ss *HashSet[T]) *HashSet[T]
