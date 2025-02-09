package binarytree

import (
	"errors"

	"github.com/untref-ayp2/data-structures/stack"
	"github.com/untref-ayp2/data-structures/types"
)

// BinarySearchTreeIterator es un iterador para recorrer un BinarySearchTree.
type binarySearchTreeIterator[T types.Ordered] struct {
	internalStack *stack.Stack[*BinaryNode[T]]
}

// pushLeftNodes apila los nodos izquierdos desde un nodo.
//
// Parámetros:
//   - `node` un puntero a un BinaryNode.
func (it *binarySearchTreeIterator[T]) pushLeftNodes(node *BinaryNode[T]) {
	for node != nil {
		it.internalStack.Push(node)
		node = node.left
	}
}

// HasNext indica si hay un siguiente dato.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	// ...
//	it := bst.Iterator()
//	for it.HasNext() {
//		fmt.Println(it.Next())
//	}
//
// Retorna:
//   - true si hay un siguiente nodo, false en caso contrario.
func (it *binarySearchTreeIterator[T]) HasNext() bool {
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido InOrder.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	// ...
//	it := bst.Iterator()
//	for it.HasNext() {
//		fmt.Println(it.Next())
//	}
//
// Retorna:
//   - el dato del siguiente nodo.
func (it *binarySearchTreeIterator[T]) Next() (T, error) {
	if it.internalStack.IsEmpty() {
		var emptyValue T

		return emptyValue, errors.New("árbol vacío")
	}
	node, _ := it.internalStack.Pop()
	if node.right != nil {
		it.pushLeftNodes(node.right)
	}

	return node.data, nil
}
