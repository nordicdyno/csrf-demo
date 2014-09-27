header_set("Content-Type", "text/html; charset=utf-8")

function main()
    print "send to <a href='/post_id.lua'>/post_id.lua</a></br>"
    print "<form action='/post_id.lua' method='POST'>"
    -- print "<form action='/post_id.lua' method='POST' enctype='multipart/form-data'>"
    print "Text: <input type='text' name='text' value='text'><br>"
    print("Session: <input type='text' name='session' value='" .. cookie_get("sid") .. "'><br>")
    print "<input type='submit'>" 
    print "</form>"
end

main()
