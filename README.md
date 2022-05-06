# GoServe

GoServe is tool to quickly start webserver in specific folder.
Typical usage is for web development.

> It's recommended to NOT use it in production!

##  Quick usage
```shell
goserve # will run webserver in current working directory at random port (49152-65535)
goserve . # is effectivelly same as `goserve`
goserve -port <port> # will run webserver at specific port (1-1024 may need privileged permission)
goserve <path> # will run webserver at specific (relative or absolute) path
goserve -public # will run webserver at 0.0.0.0 instead of localhost
```
