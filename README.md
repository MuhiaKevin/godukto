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
- Error when sending a folder. It crashes when sending some folders.
  - The dukto client receiving the file closes connection when sending some folder


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
- https://medium.com/trendyol-tech/golang-what-is-broken-pipe-error-tcp-http-connections-and-pools-3988b79f28e5
