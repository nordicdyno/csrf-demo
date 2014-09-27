Demo of CSRF attack
---------------------

Just demonstrate how work CSRF and protection strategies

Also it's my play with Go&Lua interconnection


How to use
------------

without compilation

1. get release from github (Windows, Linux)
2. fetch this repo in any dir
3. run binary from relase in this dir
4. add to /etc/hosts fake domains

    127.0.0.1 www.evil.ro
    127.0.0.1 www.nice.ro

5. open http://localhost:8080/login.lua for login
   (http://localhost:8080/logout.lua for logout)
6. open http://www.nice.ro:8080/ for CSRF-demos


Demo Description or How it works
-------------

*TODO*
