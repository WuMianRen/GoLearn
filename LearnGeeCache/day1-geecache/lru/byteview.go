/*
@Time : 2020/9/11 14:33
@Author : wumian
@File : byteview
@Software: GoLand
*/
package geecache

// 1 ByteView 只有一个数据成员 b []byte，b 将会存储真实的缓存值 这样可以存放任意数据
// 2 实现 Len 方法 返回占用的内存大小
// 3 b 是只读的  是用 ByteSlice 进行一个拷贝 防止缓存值被外部程序拷贝
type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
