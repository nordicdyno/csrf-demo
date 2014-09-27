header_set("Content-Type", "text/html; charset=utf-8")
cookie_set("sid", "AUTH_OK")
print "<b>Logged in!</b>"
print "<br><a href='/logout.lua'>Log Out</a>"
