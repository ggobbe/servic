# servic
**servic** is a tiny web server written in Go and using the net/http package.

It can only serve static files by pointing it at the folder containing the files you want to serve.

Usage: `servic [dir] ([port])`

Example: `servic my_static_folder 8080`
