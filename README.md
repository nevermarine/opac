# opac
Automatic book renewal script for HSE OPAC.
# Why?
Because I got tired of renewing my books every 2 weeks.
# Building and installation
Grab the files:
```shell
git clone https://github.com/nevermarine/opac
```
Move to project directory:
```shell
cd opac
```
Build:
```shell
make
```
Install:
```shell
sudo make install
```
Alternatively, you can skip `make install` and run the binary by its' relative path. Also, you can specify your own destination directory:
```bash
make DESTDIR=~/.local/ install
```
# Usage
```shell
$ opac                             
  -n, --dry-run        Do nothing, just print to stdout.
  -h, --help           Print this help message.
  -l, --login string   User login for HSE OPAC. Required.
  -p, --pass string    Password for HSE OPAC. Required.
```
`--dry-run` flag just prints the links to stdout without actually renewing them.
```shell
$ opac -l 111111 -p TotallyRealName --dry-run
Renewal link: https://www.youtube.com/watch?v=dQw4w9WgXcQ
# ...
$ opac -l 111111 -p TotallyRealName 
Successfully renewed the book.
Successfully renewed the book.
Successfully renewed the book.
Successfully renewed the book.
```
Running this script by hand is rather inconvinient, so you probably want to add it to your crontab. Sample crontab string:
```shell
# every 1st and 15th day of the month at 12:10 PM
10 12 1,15 * * opac -l 111111 -p TotallyRealName
```
