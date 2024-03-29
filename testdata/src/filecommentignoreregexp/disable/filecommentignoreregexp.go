// ignore this file: no

package filecommentignoreregexp

type foo interface {
	bar(a, b int) // want "declare the type of function type parameters explicitly"
}
