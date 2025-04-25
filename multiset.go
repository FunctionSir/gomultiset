package gomultiset

type Multiset[T comparable] map[T]int

// Get total elements count. Slow.
func (ms Multiset[T]) Size() int {
	size := 0
	for _, x := range ms {
		size += int(x)
	}
	return size
}

// Get distinct elements count. Faster.
func (ms Multiset[T]) Distinct() int {
	return len(ms)
}

// Check x is in multiset or not.
func (ms Multiset[T]) Has(x T) bool {
	_, has := ms[x]
	return has
}

// Count how many xs in the multiset.
func (ms Multiset[T]) Count(x T) int {
	if !ms.Has(x) {
		return 0
	}
	return ms[x]
}

// Insert count of x to multiset.
//
// If count <= 0, nothing will be done.
func (ms Multiset[T]) Insert(x T, count int) {
	if count <= 0 {
		return
	}
	if !ms.Has(x) {
		ms[x] = count
		return
	}
	ms[x] += count
}

// Remove count of x from multiset. Return removed count.
//
// If count is larger than Count(x), or <= 0, all x will be erased.
func (ms Multiset[T]) Erase(x T, count int) int {
	if !ms.Has(x) {
		return 0
	}
	if count <= 0 || count >= ms[x] {
		erased := ms[x]
		delete(ms, x)
		return erased
	}
	ms[x] -= count
	return count
}

// Clear the multiset.
func (ms Multiset[T]) Clear() {
	for x := range ms {
		ms.Erase(x, 0)
	}
}
