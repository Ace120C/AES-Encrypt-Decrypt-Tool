package main

import (
  "crypto/aes"
  "crypto/cipher"
  "fmt"
  "io"
  "os"
)

var key []byte
var ciphertext []byte

func decryptAES(key, ciphertext []byte)  ([]byte, error){
  block, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }

  iv := ciphertext[:aes.BlockSize]
  ciphertext = ciphertext[aes.BlockSize:]

  stream := cipher.NewCFBDecrypter(block, iv)
  stream.XORKeyStream(ciphertext, ciphertext)

  return ciphertext, nil
}


func main()  {
  var keypath string

  fmt.Printf("what's the key path: ")
  fmt.Scanln(&keypath)

  p, err := os.Open(keypath)
  if err != nil {
    fmt.Println("Error: couldn't open keypath ", err)
  }
  defer p.Close()

  key, err := io.ReadAll(p)
  if err != nil {
    fmt.Println("Error: couldnt read ", err)
  }


  f, err := os.Open("../encrypt/encrypted.gpg")
  if err != nil {
    fmt.Println("Error Opening the file: ", err)
  }
  defer f.Close()

  FileRead, err := io.ReadAll(f)
  if err != nil {
    fmt.Println("Failed to read: ", err)
  }

  ciphertext := FileRead

  plainText, err := decryptAES(key, ciphertext)
  if err != nil {
    fmt.Println("couldn't decrypt: ", err)
    return
  }
  fmt.Println("file decrypted successfully!")

  fileDE := "output.txt"
  content := plainText

  fileOUT, err := os.Create(fileDE)
  if err != nil {
    fmt.Println("couldn't create a file: ")
  }

  defer fileOUT.Close()

  _, err = fileOUT.WriteString(string(content))
  if err != nil {
    fmt.Println("couldn't write: ", err)
  }

  fmt.Println("Finished!")
  // fmt.Printf("deciphered text: %s\n", plainText)
}
