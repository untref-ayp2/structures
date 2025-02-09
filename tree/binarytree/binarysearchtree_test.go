package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTNuevo(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.Equal(t, 0, bstree.Size())
	assert.Nil(t, bstree.GetRoot())

	minValue, errMin := bstree.FindMin()
	assert.Zero(t, minValue)
	assert.Error(t, errMin)

	maxValue, errMax := bstree.FindMax()
	assert.Zero(t, maxValue)
	assert.Error(t, errMax)
}

func TestBSTInsertarUnElemento(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())
}

func TestBSTObtenerRaiz(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 4, bstree.GetRoot().GetData())
}

func TestBSTBuscarElementoExistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.True(t, bstree.Search(2))
	assert.False(t, bstree.Search(6))
}

func TestBSTBuscarElementoInexistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.False(t, bstree.Search(6))
}

func TestBSTBuscarMinMax(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	minValue, errMin := bstree.FindMin()
	assert.Equal(t, 1, minValue)
	assert.NoError(t, errMin)

	maxValue, errMax := bstree.FindMax()
	assert.Equal(t, 5, maxValue)
	assert.NoError(t, errMax)
}

func TestBSTRecorrerInOrder(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.Equal(t, "1 2 3 4 5 ", bstree.InOrder())
}

func TestBSTRepresentacionEnString(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.Equal(t, "1 2 3 4 5 ", bstree.String())
}

func TestBSTBorrarRaiz(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.True(t, bstree.Search(4))
	assert.Equal(t, 5, bstree.Size())

	bstree.Remove(4)

	assert.False(t, bstree.Search(4))
	assert.Equal(t, 4, bstree.Size())
}

func TestBSTBorrarTodosDeAUno(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.Equal(t, 5, bstree.Size())

	bstree.Remove(1)
	assert.False(t, bstree.Search(1))
	assert.Equal(t, 4, bstree.Size())

	bstree.Remove(2)
	assert.False(t, bstree.Search(2))
	assert.Equal(t, 3, bstree.Size())

	bstree.Remove(3)
	assert.False(t, bstree.Search(3))
	assert.Equal(t, 2, bstree.Size())

	bstree.Remove(5)
	assert.False(t, bstree.Search(5))
	assert.Equal(t, 1, bstree.Size())

	bstree.Remove(4)
	assert.False(t, bstree.Search(4))
	assert.Equal(t, 0, bstree.Size())
}

func TestBSTBorrarInexistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.Equal(t, 0, bstree.Size())

	bstree.Remove(1)

	assert.Equal(t, 0, bstree.Size())
}

func TestBSTBorrarUnicoElemento(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())

	bstree.Remove(4)

	assert.Equal(t, 0, bstree.Size())
}

func TestBSTBorrarRaizConUnHijo(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(3)

	assert.Equal(t, 2, bstree.Size())

	bstree.Remove(4)

	assert.False(t, bstree.Search(4))
	assert.Equal(t, 1, bstree.Size())
}

func TestBSTClear(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())

	bstree.Clear()

	assert.Equal(t, 0, bstree.Size())
	assert.Nil(t, bstree.GetRoot())
}

func TestBSTIsEmpty(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.True(t, bstree.IsEmpty())

	bstree.Insert(4)

	assert.False(t, bstree.IsEmpty())
}

func TestBSTIteratorUnoPorUno(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(3)
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)

	it := bstree.Iterator()

	expectedValues := []int{1, 2, 3, 4, 5}

	for _, expected := range expectedValues {
		assert.True(t, it.HasNext())
		val, err := it.Next()
		assert.Equal(t, expected, val)
		assert.NoError(t, err)
	}
	assert.False(t, it.HasNext())
}

func TestBSTIteratorWhenEmpty(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	it := bstree.Iterator()

	assert.False(t, it.HasNext())

	val, err := it.Next()

	assert.Zero(t, val)
	assert.Error(t, err)
}
