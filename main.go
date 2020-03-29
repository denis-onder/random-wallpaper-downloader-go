package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func randomString() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x%x%x%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func generateFileName() string {
	// By default, save images to $HOME/Pictures/Wallpapers
	var output bytes.Buffer
	val, _ := os.LookupEnv("HOME")
	output.WriteString(val)
	output.WriteString("/Pictures/Wallpapers/")
	output.WriteString(randomString())
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
