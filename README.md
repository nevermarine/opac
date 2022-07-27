# opac
Automatic book renewal script for HSE OPAC.
# Why?
Because I got tired of renewing my books every 2 weeks.
# Building and installation
Grab the files:
```
git clone https://github.com/nevermarine/opac
```
Move to project directory:
```
cd opac
```
Build:
```
make
```
Install:
```
sudo make install
```
Alternatively, you can skip `make install` and run the binary by its' relative path. Also, you can specify your own destination directory:
```
make DESTDIR=~/.local/ install
```
# Usage
```
$ opac                             
  -n, --dry-run        Do nothing, just print to stdout.
  -h, --help           Print this help message.
  -l, --login string   User login for HSE OPAC. Required.
  -p, --pass string    Password for HSE OPAC. Required.
```
`--dry-run` flag just prints the links to stdout without actually renewing them.
```
$ opac -l 111111 -p TotallyRealName --dry-run
Renewal link: https://www.youtube.com/watch?v=dQw4w9WgXcQ
...
$ opac -l 111111 -p TotallyRealName 
Successfully renewed the book.
Successfully renewed the book.
Successfully renewed the book.
Successfully renewed the book.
```
