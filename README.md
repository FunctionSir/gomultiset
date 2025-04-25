# gomultiset

Unordered multiset in Golang.

Well tested.

## Usage

``` go
ms := make(gomultiset.Multiset[int])
ms.Insert(1,2)
count := ms.Count(1)
erasedCnt := ms.Erase(1,1)
erasedCnt += ms.Erase(1,0)
distinct := ms.Distinct()
total := ms.Size()
has := ms.Has(1)
ms.Clear()
```
