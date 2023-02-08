function string:split(delimiter)
  if not delimiter or not self then
    return
  end
  
  local result = { }
  local from  = 1
  local delim_from, delim_to = string.find( self, delimiter, from  )
  while delim_from do
    table.insert( result, string.sub( self, from , delim_from-1 ) )
    from  = delim_to + 1
    delim_from, delim_to = string.find( self, delimiter, from  )
  end
  table.insert( result, string.sub( self, from  ) )
  return result
end

--仅兑换码
function string:trim (s)
  s = (string.gsub(s, "^%s*(.-)%s*$", "%1"))
  return (string.gsub(s, "[^%w]", ""))
end

function string:filter_spec_chars(s,more)
  local ss = {}
  local k = 1
  while true do
    if k > #s then break end
    local c = string.byte(s,k)
    if not c then break end
    if c<192 then
        if (c>=48 and c<=57) or (c>= 65 and c<=90) or (c>=97 and c<=122) then
            table.insert(ss, string.char(c))
        end
        k = k + 1
    elseif c<224 then
        local c1 = string.byte(s,k+1)
        table.insert(ss, string.char(c,c1))
        k = k + 2
    elseif c<240 then
        if c>=228 and c<=233 and not more then
            local c1 = string.byte(s,k+1)
            local c2 = string.byte(s,k+2)
            if c1 and c2 then
                local a1,a2,a3,a4 = 128,191,128,191
                if c == 228 then
                    a1 = 184
                elseif c == 233 then
                    a2,a4 = 190,c1 ~= 190 and 191 or 165
                end
                if c1>=a1 and c1<=a2 and c2>=a3 and c2<=a4 then
                  table.insert(ss, string.char(c,c1,c2))
                end
            end
        elseif more then
          local c1 = string.byte(s,k+1)
          local c2 = string.byte(s,k+2)
          table.insert(ss, string.char(c,c1,c2))
        end
        k = k + 3
    elseif c<248 then
        k = k + 4
    elseif c<252 then
        k = k + 5
    elseif c<254 then
        k = k + 6
    end
  end
  return table.concat(ss)
end