-- RESOURCES:
-- http://www.keplerproject.org/en/Talk:WSAPI
-- http://keplerproject.github.io/wsapi/
-- https://github.com/keplerproject/wsapi/blob/master/tests/test_request.lua
-- https://github.com/keplerproject/wsapi/blob/master/src/wsapi/request.lua
-- http://localhost:8080/wsapi
-- print "WSAPI"
-- local wsapi = require "wsapi"
local request = require "wsapi.request"

function run(wsapi_env)
  local r = request.new(wsapi_env, {delay_post=true})
  print "Cookies : -----------"
  for k,v in pairs(r.cookies) do print(key,value) end
  print("_panamax-ui_session -> " .. r.cookies["_panamax-ui_session"])
  print "---------------------"
  -- local res = wsapi.response.new()
  local headers = {
      ["Content-type"] = "text/html",
      ["Vary"] = "Accept-Encoding, Cookie, User-Agent"
  }

  for key,value in pairs(wsapi_env) do print(key,value) end

  local function hello_text()
    coroutine.yield("<html><body>")
    coroutine.yield("<p>Hello Wsapi!</p>")
    coroutine.yield("</body></html>")
  end

  return 200, headers, coroutine.wrap(hello_text)
end

-- print "OK"
-- run(t)
