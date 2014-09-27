function main()
    print("localhost:")
    if cookie_get("sid") == "" then print("POST Failed, sid cookie not set") return end

    print("«POST OK»")
    print("Host: " .. header_get("Host"))
    print("Referer: " .. header_get("Referer"))
    print("Origin: " .. header_get("Origin"))
end

main()
