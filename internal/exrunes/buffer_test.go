package exrunes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewBuffer(t *testing.T) {
	b := NewBuffer(100)
	assert.Equal(t, 0, len(b.buf))
	assert.Equal(t, 100, cap(b.buf))
	assert.Equal(t, 0, b.index)
	assert.Equal(t, 0, len(b.items))
}

func Test_Buffer_Add(t *testing.T) {
	b := NewBuffer(100)

	// add first element
	index := b.Add([]rune{'1', '2', '3'})
	assert.Equal(t, 0, index)
	assert.Equal(t, 1, b.index)
	assert.Equal(t, []rune{'1', '2', '3', 0, 0, 0}, b.buf)
	assert.Equal(t, 0, b.items[index].from)
	assert.Equal(t, 3, b.items[index].to)
	assert.Equal(t, 3, b.items[index].space)

	// add second element
	index = b.Add([]rune{'4', '5', '6'})
	assert.Equal(t, 1, index)
	assert.Equal(t, 2, b.index)
	assert.Equal(t, []rune{'1', '2', '3', 0, 0, 0, '4', '5', '6', 0, 0, 0}, b.buf)
	assert.Equal(t, 6, b.items[index].from)
	assert.Equal(t, 9, b.items[index].to)
	assert.Equal(t, 3, b.items[index].space)
}

func Test_Buffer_AddWithSpace(t *testing.T) {
	b := NewBuffer(100)

	// add first element
	index := b.AddWithSpace([]rune{'1', '2', '3'}, 1)
	assert.Equal(t, 0, index)
	assert.Equal(t, 1, b.index)
	assert.Equal(t, []rune{'1', '2', '3', 0}, b.buf)
	assert.Equal(t, 0, b.items[index].from)
	assert.Equal(t, 3, b.items[index].to)
	assert.Equal(t, 1, b.items[index].space)

	// add second element
	index = b.AddWithSpace([]rune{'4', '5', '6'}, 5)
	assert.Equal(t, 1, index)
	assert.Equal(t, 2, b.index)
	assert.Equal(t, []rune{'1', '2', '3', 0, '4', '5', '6', 0, 0, 0, 0, 0}, b.buf)
	assert.Equal(t, 4, b.items[index].from)
	assert.Equal(t, 7, b.items[index].to)
	assert.Equal(t, 5, b.items[index].space)
}

func Test_Buffer_Copy(t *testing.T) {
	b := NewBuffer(100)

	index := b.Add([]rune{'1', '2', '3'})

	index2 := b.Copy(index)
	assert.Equal(t, 1, index2)
	assert.Equal(t, 2, b.index)
	assert.Equal(t, []rune{'1', '2', '3', 0, 0, 0, '1', '2', '3', 0, 0, 0}, b.buf)
	assert.Equal(t, 6, b.items[index2].from)
	assert.Equal(t, 9, b.items[index2].to)
	assert.Equal(t, b.items[index].space, b.items[index2].space)
}

func Test_Buffer_Get(t *testing.T) {
	b := NewBuffer(100)

	data := []rune{'1', '2', '3'}
	id := b.Add(data)

	require.Equal(t, data, b.Get(id))
	require.Nil(t, b.Get(id+1))
	require.Nil(t, b.Get(id-1))

	data2 := []rune{'4', '5', '6'}
	id2 := b.Add(data2)
	require.Equal(t, data2, b.Get(id2))
	require.Equal(t, data, b.Get(id))
}

func Test_Buffer_Append(t *testing.T) {
	t.Run("must add runes to the buffer", func(t *testing.T) {
		b := NewBuffer(100)

		id := b.Add([]rune{'1', '2', '3'})
		b.Append(id, []rune{'4', '5', '6'})

		require.Equal(t, []rune{'1', '2', '3', '4', '5', '6'}, b.buf)
	})

	t.Run("must move rune batch to the end if cannot append", func(t *testing.T) {
		b := NewBuffer(100)

		index := b.Add([]rune{'1', '2', '3'})
		_ = b.Add([]rune{'4', '5', '6'})
		b.Append(index, []rune{'7', '8', '9', '0'})

		require.Equal(t, []rune{'1', '2', '3', 0, 0, 0, '4', '5', '6', 0, 0, 0, '1', '2', '3', '7', '8', '9', '0', 0, 0, 0, 0, 0, 0, 0}, b.buf)
		assert.Equal(t, 12, b.items[index].from)
		assert.Equal(t, 19, b.items[index].to)
		assert.Equal(t, 7, b.items[index].space)
	})
}
