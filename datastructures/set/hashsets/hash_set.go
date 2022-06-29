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
func (p *HashSet[T]) Add(values ...T) {
	for _, value := range values {
		p.data[value] = struct{}{}
	}
}
func (p *HashSet[T]) Remove(values ...T) {
	for _, value := range values {
		delete(p.data, value)
	}

}
func (p *HashSet[T]) Contains(value T) bool {
	_, ok := p.data[value]
	return ok
}
func (p *HashSet[T]) ContainsAll(values ...T) bool {
	for _, value := range values {
		if !p.Contains(value) {
			return false
		}
	}
	return true
}
func (p *HashSet[T]) ContainsAny(values ...T) bool {
	for _, value := range values {
		if p.Contains(value) {
			return true
		}
	}
	return false
}

func (p *HashSet[T]) GetValues() []T {
	values := []T{}
	for key := range p.data {
		values = append(values, key)
	}
	return values
}

func (p *HashSet[T]) Merge(sets ...*HashSet[T]) {
	for _, set := range sets {
		for _, value := range set.GetValues() {
			p.Add(value)
		}
	}
}

func (p *HashSet[T]) Copy() *HashSet[T] {
	return New(p.GetValues()...).(*HashSet[T])
}

func (p *HashSet[T]) Clear() {
	p.data = make(map[T]struct{})
}
func (p *HashSet[T]) IsEmpty() bool {
	return p.Size() == 0
}
func (p *HashSet[T]) Size() int {
	return len(p.data)
}

func (p *HashSet[T]) Union(ss *HashSet[T]) *HashSet[T] {
	s := p.Copy()
	s.Merge(ss)
	return s
}
func (p *HashSet[T]) Intersection(ss *HashSet[T]) *HashSet[T] {
	s := New[T]()
	for _, value := range p.GetValues() {
		if ss.Contains(value) {
			s.Add(value)
		}
	}
	return s.(*HashSet[T])
}
func (p *HashSet[T]) SymmetricDifference(ss *HashSet[T]) *HashSet[T] {
	s := New[T]()

	for _, v := range p.GetValues() {
		if !ss.Contains(v) {
			s.Add(v)
		}
	}

	for _, v := range ss.GetValues() {
		if !p.Contains(v) {
			s.Add(v)
		}
	}

	return s.(*HashSet[T])
}
func (p *HashSet[T]) Subtraction(ss *HashSet[T]) *HashSet[T] {
	s := p.Copy()
	for _, v := range ss.GetValues() {
		if p.Contains(v) {
			s.Remove(v)
		}
	}
	return s
}
