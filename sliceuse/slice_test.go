package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// slice init way
	b := []int{}
	fmt.Println(b)
	c := []int{1, 2, 4}
	fmt.Println(c)
	// two ele , cap 3
	d := c[:2]
	fmt.Println(d)
	// two ele , cap is three
	e := c[0:2:3]
	fmt.Println(e)
	// cap is three, 0 cap, 0ele
	f := c[:0]
	fmt.Println(f)
	// TODO ------ create slice two way ,first direct 	assignment , two way use make
	// create null slice
	// 3 len, 3cap , 0ele
	g := make([]int, 0, 3)
	fmt.Println(g)
	// 3 ele, len and cap is 3
	h := make([]int, 3)
	fmt.Println(h)
	// cap is three , len is 2
	i := make([]int, 2, 3)
	fmt.Println(i)
	// range is same as array
	var slice1 []int
	slice1 = append(slice1, 2)
	// add slice need to unpackage
	slice1 = append(slice1, []int{1, 2, 4}...)
	// slice 可以在头部， 尾巴， 中间插入
	// 头部插入， 那么slice放在尾巴就行。 如果在中间插入， 需要  将指定位置向后移动， 然后填充指定位置的值。

	// 不仅可以在尾部， 也可以 在头部 add
	slice1 = append([]int{1, 2}, slice1...)
	fmt.Println(slice1)
	a := []int{1, 2, 4}
	a = append(a, 0)
	fmt.Println("a add a ele:", a)
	var x int32 = 1
	fmt.Println("a[x+1:]: ", a[x+1:])
	fmt.Println("a[x:]:", a[x:])

	// 使用copy 将 将数据移动, 前面目标数组，后面源数组。 copy 将后面源数组的内容复制到目标数组。 两个坐标的差异决定了到底是左移还是一边移动
	// 将x: 开始的数据向后移动一位
	copy(a[x+1:], a[x:])
	fmt.Println(a)
	a[x] = 100
	fmt.Println(a)

	// 删除元素的三种办法
	// 删除尾部
	a = a[:(len(a) - 1)]
	// 删除开头 移动指针 删除开头的第1个、n个元素
	a = a[1:]

	// 也可以使用append
}

func Test_addslice(t *testing.T) {
	var a []int
	a = append(a, 9)

	var b = []int{10}
	a = append(a, b...)
	fmt.Println("尾部增加元素", a)

	a = append([]int{1}, a...)
	fmt.Println("头部增加元素", a)

	fmt.Println("在中间添加元素使用append 扩大空间, 使用copy 移动位置")
	a = append(a, 0)
	fmt.Println("z 扩容后的元素", a)
	// 增加元素的索引
	var i = 1
	copy(a[i+1:], a[i:])
	fmt.Println("向后移动一个元素", a)
	a[i] = 9898
	fmt.Println("插入一个元素后", a)
	fmt.Println("删除尾部的元素是最快的。")
}

// 截取分片一般cap 保存下来了， 但是
func Test_deleteSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := a[:1]
	fmt.Println(b)
	fmt.Printf("a %v \n", a)
	var n = 1
	a = a[:len(a)-n]
	fmt.Printf("删除1个尾部 %v \n", a)
	a = a[n:]
	fmt.Printf("删除头部后 %v \n", a)
	// 删除头部也可以使用append 向前移动, append 还是在原来的空间里面完成
	a = append(a[:0], a[n:]...)
	fmt.Printf("使用append 删除头部 %v \n", a)
	a = a[:copy(a, a[1:])]
	fmt.Printf("使用copy 覆盖前面一个元素 %v \n", a)

	// 删除元素有2种方法一种是，  使用append(截取两边的数据） 然后加起来。
	//  另外一种使用copy ,

	// delete   medum and delete head is speeccal way
	// append 删除中间的元素就是， 截取左边的元素，然后截取右边的元素， 然后append 起来
	h := []int{1, 2, 3, 4, 5, 6}
	i := 1
	var deleteL = 2
	// 删除索引2开始2个元素
	h = append(h[:i], h[i+deleteL:]...) // 删除中间 N 个元素
	fmt.Println("使用append 删除索引为2开始的2个元素", h)

	// 使用copy 将末尾的元素从要删除的元素开始覆盖即可。 然后从删除的index的位置，加上被删除的元素的个数。
	i2 := copy(h[i:], h[i+deleteL:])
	fmt.Println(i2)
	h = h[:i+i2]

	fmt.Printf("索引为2删除一个元素 %v", h)
}

func Test_space(t *testing.T) {
	bytes := []byte{12, ' ', 123}
	println(string(TrimSpace(bytes)))

}

func TrimSpace(b []byte) []byte {
	s := b[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

func TestPrint(t *testing.T) {
	var a = []interface{}{1, "abd"}
	// 解包调用的话， 只有 1,adb
	fmt.Println(a...)
	fmt.Println("--------------------")
	// 没有解包调用的话会有外部类型
	fmt.Println(a)
}

func If(condition bool, a any, b any) any {
	if condition {
		return a
	}
	return b
}

func IfNil(cond any, a any, b any) any {
	if cond == nil {
		return a
	}
	return b
}
func IfEmpty(obj interface{}, a any, b any) any {
	if obj == nil {
		return a
	}
	switch obj.(type) {
	case string:
		b2 := len(obj.(string)) > 0
		return If(b2, a, b)
	}
}
func Test_if(t *testing.T) {
	ifNil := IfNil("asdf", "is nil", "not nil")
	fmt.Println(ifNil)
}
