#!/bin/bash

go build .
sudo mv ./random-wallpaper-downloader-go /usr/local/bin

echo "Moved the binary to /usr/local/bin"
echo "You can now use the random-wallpaper-downloader from anywhere in your terminal!"
echo "Try running rwd"