// http://www.code2succeed.com/golang-insert-and-search-trie/

package main

import (
	"fmt"
)

const ALPHABET_SIZE = 26

type TrieNode struct {
	children   [ALPHABET_SIZE]*TrieNode
	endOfWords bool
	synonyms   []string
}

func getNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false

	for i := 0; i < ALPHABET_SIZE; i++ {
		node.children[i] = nil
	}

	return node
}

func insert(root *TrieNode, key string) {
	temp := root // En primer lugar, se utiliza la raiz para verificar si alguno de sus hijos coincide con la primer letra del KEY

	for i := 0; i < len(key); i++ { // Se itera sobre todos los elementos del KEY
		index := key[i] - 'a'
		if temp.children[index] == nil { // Si no hay un hijo asignado a esa letra, se crea un nuevo nodo para ese hijo
			temp.children[index] = getNode() //Esto permite que se vayan cargando las palabras nuevas, letra por letra
		}
		temp = temp.children[index] // Se actualiza el nodo actual para referenciar al hijo en la siguiente iteracion
	}

	temp.endOfWords = true //Se deja un flag al final de la palabra
}

func search(root *TrieNode, key string) bool {
	temp := root

	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			return false
		}
	}
	return (temp != nil && temp.endOfWords)
}

func insertSynonimous(root *TrieNode, key string, syn string) bool {
	temp := root
	r := false
	i := 0
	for !r && i < len(key) {
		index := key[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			r = true
		}
		i++
	}
	if temp.endOfWords {
		temp.synonyms = append(temp.synonyms, syn) //sinonimo insertado
	}
	return r
}

func getSynonimous(root *TrieNode, key string) []string {
	temp := root
	flag := false
	r := []string{}
	i := 0
	for !flag && i < len(key) {
		index := key[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			flag = true
		}
		i++
	}
	if temp.endOfWords {
		r = temp.synonyms
	}
	return r
}

func main() {
	words := []string{"a", "and", "go", "golang", "man", "mango"} //strings a ingresar
	root := getNode()

	for i := 0; i < len(words); i++ {
		insert(root, words[i]) //Insercion de strings
	}

	// Consulta por los strings
	fmt.Println("contains words [a]: ", search(root, "a"))
	fmt.Println("contains words [mango]: ", search(root, "mango"))
	fmt.Println("contains words [lang]: ", search(root, "lang"))
	fmt.Println("contains words [an]: ", search(root, "an"))

	insertSynonimous(root, "mango", "durazno")
	fmt.Println("Mango sinonymous:", getSynonimous(root, "mango"))
}
