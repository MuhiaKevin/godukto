# Godukto
Dukto commandline app written in Go!

Explore gum, charm and others to imporve cli

## TODO
- [x] Sending a file  
- [x] support sending to multiple dukto clients using channels and goroutines
- [x] Enable sending multiple files
- [ ] Add Progress bar when sending a file to other dukto apps
- [ ] Show waiting animation when waiting for other dukto apps to show up

# Bugs
- Error when sending a folder. It crashes when sending some files
- Crashes when sending multiple files in a folder that has some folders


```sh
$ godukto sendfiles README.md
```

- Sending a directory

```sh
$ godukto sendfolder pictures/
```

- Receive a directory or file

```sh
$ godukto receive 
```

### Some Resources
- https://github.com/aler9/howto-udp-broadcast-golang
- https://www.digitalocean.com/community/tutorials/how-to-use-the-cobra-package-in-go
