-- header_set("Content-Type", "text/html; charset=utf-8")
-- cookie_set("A", "b")
-- header_set("X-Header", "XYZ")
print("localhost:")
print("«POST OK»")
print("Host: " .. header_get("Host"))
print("Origin: " .. header_get("Origin"))
-- print("Cookie A=" .. cookie_get("A"))
-- print("Cache-Control: " .. header_get("Cache-Control"))
-- silently doesn't work
-- cookie_set("A", "c")
-- header_set("Content-Type", "text/plain; charset=utf-8")
