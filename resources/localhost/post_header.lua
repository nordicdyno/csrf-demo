function main()
    print("localhost:")
    if cookie_get("sid") == "" then print("POST Failed, sid cookie not set") return end

    host = header_get("Host")
    -- print("Host: " .. host)

    refH = header_get("Referer")
    -- print("Referer: " .. refH)
    refH = get_referer_host(refH)
    -- print("RefOrigin: " .. refH)

    -- get origin from request
    hostOrigin = origin_get("Origin")
    -- print("Host Real origin: " .. hostOrigin)

    check_field, check_value = "Referer", refH
    origin = header_get("Origin")
    -- print("Browser origin: " .. origin)
    -- origin = string.gsub((string.gsub(origin, "https?://", "")), ":.*", "")
    -- origin = (string.gsub(origin, "https?://", ""))
    --print("Origin: " .. origin)
    if string.len(origin) > 0 then
        check_field, check_value = "Origin", origin
    end
    if hostOrigin == check_value then
        print("«POST OK»")
        print("Real origin '" .. hostOrigin .. "' match header " .. check_field .. " '" .. check_value .. "'")
    else
        print("«POST FAILED»")
        print("Real origin '" .. hostOrigin .. "' not match header " .. check_field .. " '" .. check_value .. "'")
    end
end

function get_referer_host(s)
    local start, finish = string.find(s, 'http://[^/]+/')
    local ret = ""
    if start ~= nil then
        ret = string.sub(s, start, finish - 1)
    end
    return ret
end

main()
