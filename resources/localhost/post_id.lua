function main()
    print("localhost:")
    formId = form_value('session')
    cookieId = cookie_get("sid")
    if cookieId == "" then print("POST Failed, sid cookie not set") return end

    if cookieId == formId then
        print("«POST OK»")
    else
        print("«POST FAILED»")
        print("session field '" .. cookieId .. "' doesn't match Cookie 'sid'='" .. formId .. "'")
    end
end

main()
