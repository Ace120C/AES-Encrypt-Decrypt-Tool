package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)


func keygen() (string, error) {
  key := make([]byte, 16)
  
  _, err := rand.Read(key)
  if err != nil {
    return "", nil
  }

  return hex.EncodeToString(key), nil
}


func main()  {
  key, err := keygen()
  if err != nil {
    fmt.Println("error generating a key: ", err)
  }

  var filename string = "key.txt"
  content := key
  
  f, err := os.Create(filename)
  if err != nil {
    fmt.Println("error creating file: ", err)
  }

  defer f.Close()

  _, err = f.WriteString(string(content))
  if err != nil {
    fmt.Println("error writing file: ", err)
  }

  fmt.Println("Key generated successfully, store the file somewhere safe!")

}
