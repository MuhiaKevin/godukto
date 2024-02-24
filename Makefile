file:
	go build -o bin/godukto && bin/godukto sendfiles ~/Downloads/ssstwitter.com_1707503817955.mp4

files:
	go build -o bin/godukto && bin/godukto sendfiles  ~/Downloads/ssstwitter.com_1707503817955.mp4 ~/Downloads/GFT-Liverpool.mp3 ~/Downloads/Videos/*.jpg

build:
	go build -o bin/godukto
