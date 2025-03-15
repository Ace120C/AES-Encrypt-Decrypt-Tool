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
  key := []byte("1234567890123456")
  fmt.Print("Msg: ")
  fmt.Scanln(&plainText)
  ciphertext, err := encryptAES(key, plainText)
  if err != nil {
    fmt.Println("error: ", err)
  }

  file := "encryptedfile.txt"
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
