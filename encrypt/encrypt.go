package main

import (
  "fmt"
  "crypto/aes"
  "crypto/rand"
  "crypto/cipher"
  "io"
  "os"
)

var plainText []byte 

func encryptAES(key, plaintext []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }

  cipherText := make([]byte, aes.BlockSize+len(plaintext))
  iv := cipherText[:aes.BlockSize]

  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    return nil, err
  }

  stream := cipher.NewCFBEncrypter(block, iv)
  stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

  return cipherText, nil
}


func main()  {
  var key []byte
  var path string
  fmt.Print("Type your passkey (has to be 16 char long): ")
  fmt.Scanln(&key)
  fmt.Print("Path: ")
  fmt.Scanln(&path)

  fp, err := os.Open(path)
  if err != nil {
    fmt.Println("Couldn't open path: ", err)
  }
  defer fp.Close()

  c, err := io.ReadAll(fp)
  if err != nil {
    fmt.Println("couldn't read file: ", err)
  }

  ciphertext, err := encryptAES(key, c)
  if err != nil {
    fmt.Println("error: ", err)
  }

  file := "encrypted.gpg"
  content := ciphertext
  
  f, err := os.Create(file)
  if err != nil {
    fmt.Println("Error creating the file: ", err)
  }

  defer f.Close()

  _, err = f.WriteString(string(content))
  if err != nil {
    fmt.Println("error: ", err)
  }

  fmt.Println("File written successfully!")
}
