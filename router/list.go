package router

import "sort"

type List struct {
  Length int
  Items map[int]interface{}
}

func CreateList() *List {
  return &List{Length:0, Items: make(map[int]interface{})}
}

func (this *List) Push(handler interface{}) {
  this.Items[this.Length] = handler
  this.Length++
}

func (this *List) Get(index int) interface{} {
  return this.Items[index]
}

// need to make sure that the handlers are sorted
func (this *List) ToArray() []interface{} {
  var keys []int
  for k := range this.Items {
    keys = append(keys, k)
  }
  sort.Ints(keys)

  output := make([]interface{}, this.Length)
  for _, index := range keys {
    output[index] = this.Items[index]
  }

  return output
}


