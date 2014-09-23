-- header_set("Content-Type", "text/html; charset=utf-8")
-- cookie_set("A", "b")
-- header_set("X-Header", "XYZ")
function main()
    print("localhost:")
    -- print("Cookie A=" .. cookie_get("A"))
    host = header_get("Host")
    print("Host: " .. host)

    refH = header_get("Referer")
    print("Referer: " .. refH)
    print("RefOrigin: " .. get_referer_host(refH))

    origin = origin_get("Origin")
    print("Real origin: " .. origin)
    origin = header_get("Origin")
    print("Browser origin: " .. origin)
    -- origin = string.gsub((string.gsub(origin, "https?://", "")), ":.*", "")
    origin = (string.gsub(origin, "https?://", ""))
    --print("Origin: " .. origin)
    if host == origin then
        print("«POST OK»")
    else
        print("«POST FAILED»")
    end
end

function get_referer_host(s)
    local start, finish = string.find(s, 'http://[^/]+/')
    local ret = ""
    if start ~= nil then
        ret = "concat: " .. start .. ":" .. finish
        ret = string.sub(s, start, finish - 1)
    end
    return ret
    -- return ""
    -- (table.concat(res, ", ") or " nil ")
end

main()

-- silently doesn't work
-- cookie_set("A", "c")
-- header_set("Content-Type", "text/plain; charset=utf-8")
