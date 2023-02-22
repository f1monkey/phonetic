package exrunes

import (
	"sync"

	"golang.org/x/exp/slices"
)

type Buffer struct {
	buf   []rune
	index int
	items map[int]bufferItem
}

func NewBuffer(length int) *Buffer {
	return &Buffer{
		buf:   make([]rune, 0, length),
		items: make(map[int]bufferItem, length/10),
	}
}

func (b *Buffer) Add(data []rune) int {
	curIndex := b.index
	b.rawAdd(data, len(data), curIndex)
	b.index++

	return curIndex
}

func (b *Buffer) Copy(id int) int {
	item := b.items[id]
	curIndex := b.index
	b.rawAdd(b.buf[item.from:item.to], item.space, curIndex)
	b.index++

	return curIndex
}

func (b *Buffer) AddWithSpace(data []rune, space int) int {
	if space < 0 {
		panic("space must be >= 0")
	}

	curIndex := b.index
	b.rawAdd(data, space, curIndex)
	b.index++

	return curIndex
}

func (b *Buffer) Append(id int, data []rune) {
	item := b.items[id]
	if item.space >= len(data) {
		written := b.rawWrite(data, id)
		item.to += written
		item.space -= written
		b.items[id] = item
		return
	}

	oldData := b.buf[item.from:item.to]
	dataLen := len(oldData) + len(data)
	from := len(b.buf)
	b.rawAppend(oldData, 0)
	b.rawAppend(data, dataLen)
	item.from = from
	item.to = from + dataLen
	item.space = dataLen
	b.items[id] = item
}

func (b *Buffer) Get(id int) []rune {
	item, ok := b.items[id]
	if !ok {
		return nil
	}

	return b.buf[item.from:item.to]
}

func (b *Buffer) rawAdd(data []rune, space int, id int) {
	l := len(b.buf)
	b.rawAppend(data, space)

	b.items[id] = bufferItem{
		from:  l,
		to:    l + len(data),
		space: space,
	}
}

func (b *Buffer) rawAppend(data []rune, space int) {
	b.buf = append(b.buf, data...)
	for i := 0; i < space; i++ {
		b.buf = append(b.buf, 0)
	}
}

func (b *Buffer) rawWrite(data []rune, id int) int {
	item := b.items[id]
	idx := item.to
	for i := idx; i < idx+len(data); i++ {
		b.buf[i] = data[i-idx]
	}

	return len(data)
}

type bufferItem struct {
	from  int
	to    int
	space int
}

var bufStorage = []*Buffer{}
var bufMtx sync.Mutex

// BufferFree destroy buffer contents and return it to storage slice
func BufferFree(buf *Buffer) {
	for k := range buf.items {
		delete(buf.items, k)
	}
	buf.buf = buf.buf[:0]

	bufMtx.Lock()
	bufStorage = append(bufStorage, buf)
	bufMtx.Unlock()
}

// BufferGet get from storage or create buffer
func BufferGet(length int) *Buffer {
	bufMtx.Lock()
	defer bufMtx.Unlock()

	if len(bufStorage) > 0 {
		result := bufStorage[0]
		bufStorage[0] = nil // prevent memory leak
		bufStorage = slices.Delete(bufStorage, 0, 1)
		return result
	}

	return NewBuffer(length)
}
