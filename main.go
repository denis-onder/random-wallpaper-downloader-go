package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func generateFileName() string {
	// By default, save images to $HOME/Pictures
	var output bytes.Buffer
	val, _ := os.LookupEnv("HOME")
	output.WriteString(val)
	output.WriteString("/Pictures/")
	output.WriteString(randomString(12))
	output.WriteString(".jpeg")
	return output.String()
}

func download(url string) error {
	filepath := generateFileName()

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	// Defer => Execute as last action in function
	defer res.Body.Close()

	img, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer img.Close()

	io.Copy(img, res.Body)
	return err
}

func main() {
	const url = "https://source.unsplash.com/random/3840x2160/?wallpaper"
	if err := download(url); err != nil {
		fmt.Println("An error has occured!\n", err)
	}
}
