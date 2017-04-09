package router

import "sort"

type Data map[string]interface{}

type List struct {
  Length int
  Items map[int]Handler
}

func CreateList() *List {
  return &List{Length:0, Items: make(map[int]Handler)}
}

func (this *List) Push(h Handler) {
  this.Items[this.Length] = h
  this.Length++
}

func (this *List) Get(index int) Handler {
  return this.Items[index]
}

// need to make sure that the handlers are sorted
func (this *List) ToArray() []Handler {
  var keys []int
  for k := range this.Items {
    keys = append(keys, k)
  }
  sort.Ints(keys)

  output := make([]Handler, this.Length)
  for _, index := range keys {
    output[index] = this.Items[index]
  }

  return output
}


