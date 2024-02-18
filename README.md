# godukto
https://github.com/aler9/howto-udp-broadcast-golang
https://www.digitalocean.com/community/tutorials/how-to-use-the-cobra-package-in-go
Dukto commandline app written in Go!

Explore gum, charm and others to imporve cli

## TODO
- Sending a file  

```sh
$ godukto sendfile README.md
```

- Sending a directory

```sh
$ godukto folder pictures/
```

- Receive a directory or file

```sh
$ godukto receive 
```

- Add Progress bar when sending a file to other dukto apps
- Show waiting animation when waiting for other dukto apps to show up
- support sending to multiple dukto clients using channels and goroutines
- Enable sending multiple files
