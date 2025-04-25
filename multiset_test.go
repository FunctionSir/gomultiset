/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-25 21:11:26
 * @LastEditTime: 2025-04-25 22:48:35
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /gomultiset/multiset_test.go
 */

package gomultiset

import (
	"math/rand"
	"testing"
)

const TEST_SCALE int = 1000000

func TestMultisetSimple(t *testing.T) {
	ms := make(Multiset[int])
	size := ms.Size()
	distinct := ms.Distinct()
	if size != 0 {
		t.Errorf("Size should be 0, but got %d", size)
	}
	if distinct != 0 {
		t.Errorf("Distinct should be 0, but got %d", distinct)
	}
	ms.Insert(1, 0)
	size = ms.Size()
	distinct = ms.Distinct()
	if size != 0 {
		t.Errorf("Size should be 0, but got %d", size)
	}
	if distinct != 0 {
		t.Errorf("Distinct should be 0, but got %d", distinct)
	}
	ms.Insert(1, -1)
	size = ms.Size()
	distinct = ms.Distinct()
	if size != 0 {
		t.Errorf("Size should be 0, but got %d", size)
	}
	if distinct != 0 {
		t.Errorf("Distinct should be 0, but got %d", distinct)
	}
	for i := 1; i <= TEST_SCALE; i++ {
		ms.Insert(1, 1)
		count := ms.Count(1)
		size := ms.Size()
		distinct := ms.Distinct()
		if count != i {
			t.Errorf("Count should be %d, but got %d", i, count)
		}
		if size != i {
			t.Errorf("Size should be %d, but got %d", i, size)
		}
		if distinct != 1 {
			t.Errorf("Distinct should be %d, but got %d", i, distinct)
		}
	}
	for i := TEST_SCALE - 1; i >= 0; i-- {
		ms.Erase(1, 1)
		count := ms.Count(1)
		size := ms.Size()
		if count != i {
			t.Errorf("Count should be %d, but got %d", i, count)
		}
		if size != i {
			t.Errorf("Size should be %d, but got %d", i, size)
		}
	}
	if ms.Has(1) {
		t.Error("Should have no 1, but got have")
	}
	if ms.Erase(1, 1) != 0 {
		t.Error("Should erase 0 1, but got not 0")
	}
	ms.Insert(1, 10)
	size = ms.Size()
	if size != 10 {
		t.Errorf("Size should be 10, but got %d", size)
	}
	if ms.Erase(1, -1) != 10 {
		t.Error("Should erase 10 1, but got not 10.")
	}
}

func TestMultisetRandom(t *testing.T) {
	data := make([]int, TEST_SCALE)
	correct := make(map[int]int)
	ms := make(Multiset[int])
	for i := 0; i < TEST_SCALE; i++ {
		data[i] = rand.Intn(0x3f3f3f3f)
		correct[data[i]]++
		ms.Insert(data[i], 1)
	}
	for i := 0; i < TEST_SCALE; i++ {
		count := ms.Count(data[i])
		if count != correct[data[i]] {
			t.Errorf("Count should be %d, but got %d", correct[data[i]], count)
		}
	}
	size := ms.Size()
	if size != TEST_SCALE {
		t.Errorf("Size should be %d, but got %d", TEST_SCALE, size)
	}
	ms.Clear()
	size = ms.Size()
	if size != 0 {
		t.Errorf("Size should be 0, but got %d", size)
	}
}
