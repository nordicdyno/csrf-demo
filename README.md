Demo of CSRF attack
---------------------

Just demonstrate how work CSRF and protection strategies

Also it's my play with Go&Lua interconnection


How to use
------------

without compilation

1. get release from github (Mac OS X only aviable)
2. fetch this repo in any dir
3. run binary from relase in this dir
4. add to /etc/hosts fake domains

    127.0.0.1 www.evil.ro
    127.0.0.1 www.nice.ro

5. open http://localhost:8080/login.lua for login
   (http://localhost:8080/logout.lua for logout)
6. open http://www.nice.ro:8080/ for CSRF-demos

Install on Linux
------------------

1. Install Go runtime: https://golang.org/dl/, http://golang.org/doc/install
2. Install lua dev libs apt-get liblua5.1-0-dev
3. Setup GOPATH: export GOPATH=~
4. go get github.com/nordicdyno/csrf-demo

other steps is the same as in `How to use' section


Demo Description or How it works
-------------

*TODO*
