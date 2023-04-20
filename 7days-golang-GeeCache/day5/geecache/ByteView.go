// 缓存值的抽象与封装
package geecache

type ByteView struct {
	b []byte //缓存值,byte是为了通用性
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string { // 返回一个字符串的拷贝
	return string(v.b)
}

func cloneBytes(b []byte) []byte { // 返回一个byte的拷贝
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
