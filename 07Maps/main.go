package main

import "fmt"

func main() {
  // creating a map of authors and one of their books
  m := map[string]string{
    "Miguel de Cervantes" : "Don Quixote" ,
    "Charles Dickens" : "A Tale of Two Cities" ,
    "J.R.R. Tolkien" : "The Lord of the Rings" ,
    "Antoine de Saint-Exupery" : "The Little Prince" ,
    "J.K. Rowling" : "Harry Potter and the Sorcerer's Stone" ,
    "Agatha Christie" : "And Then There Were None" ,
    "Cao Xueqin" : "The Dream of the Red Chamber" ,
    "H. Rider Haggard" : "She: A History of Adventure" ,
    "C.S. Lewis" : "The Lion, the Witch and the Wardrobe" ,
    "Dan Brown" : "The Da Vinci Code" ,
  }
 // Prints out a single entry from the map with a key of J.K. Rowling
  fmt.Println(m["J.K. Rowling"])
  // prints out the entire map
  fmt.Println(m)
}
