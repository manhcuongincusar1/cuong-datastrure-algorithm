package hashmultimaps

type IHashMultiMap[K, V comparable] interface {
	Put(key K, value V)
	PutAll(key K, value ...V)

	GetKeys() []K
	GetValues(key K) []V

	Contains(key K) bool
	ContainsAll(keys ...K) bool
	ContainsAny(keys ...K) bool
	IsEmpty() bool
	Size() int

	Clear()
	Remove(key K, value V)
	RemoveKey(keys ...K)

	Merge(maps ...*HashMultiMap[K, V])
}

type HashMultiMap[K, V comparable] struct {
	data map[K][]V
}

func New[K, V comparable]() IHashMultiMap[K, V] {
	_map := HashMultiMap[K, V]{
		data: make(map[K][]V),
	}
	return &_map
}

func (p *HashMultiMap[K, V]) Put(key K, value V) {
	values, ok := p.data[key]
	if !ok {
		values = make([]V, 0)
	}
	values = append(values, value)
	p.data[key] = values
}
func (p *HashMultiMap[K, V]) PutAll(key K, values ...V) {
	for _, v := range values {
		p.Put(key, v)
	}
}
func (p *HashMultiMap[K, V]) GetKeys() []K {
	keys := make([]K, 0, p.Size())
	for key := range p.data {
		keys = append(keys, key)
	}
	return keys
}
func (p *HashMultiMap[K, V]) GetValues(key K) []V {
	values, ok := p.data[key]
	if !ok {
		values = make([]V, 0)
	}
	return values
}
func (p *HashMultiMap[K, V]) Contains(key K) bool {
	_, ok := p.data[key]
	return ok
}
func (p *HashMultiMap[K, V]) ContainsAll(keys ...K) bool {
	for _, key := range keys {
		if !p.Contains(key) {
			return false
		}
	}
	return true
}
func (p *HashMultiMap[K, V]) ContainsAny(keys ...K) bool {
	for _, key := range keys {
		if p.Contains(key) {
			return true
		}
	}
	return false
}
func (p *HashMultiMap[K, V]) IsEmpty() bool {
	return p.Size() == 0
}
func (p *HashMultiMap[K, V]) Size() int {
	return len(p.data)
}
func (p *HashMultiMap[K, V]) Clear() {
	p.data = map[K][]V{}
}
func (p *HashMultiMap[K, V]) Remove(key K, value V) {
	values, ok := p.data[key]
	if !ok {
		return
	}

	index := getIndex(value, values...)
	if index == -1 {
		return
	}
	newValues := make([]V, 0)
	newValues = append(newValues, values[:index]...)
	p.data[key] = append(newValues, values[index+1:]...)
}
func (p *HashMultiMap[K, V]) RemoveKey(keys ...K) {
	for _, key := range keys {
		delete(p.data, key)
	}
}
func (p *HashMultiMap[K, V]) Merge(maps ...*HashMultiMap[K, V]) {
	for _, m := range maps {
		for key, values := range m.data {
			p.PutAll(key, values...)
		}
	}
}

func getIndex[V comparable](value V, values ...V) int {
	for i, v := range values {
		if v == value {
			return i
		}
	}
	return -1
}
