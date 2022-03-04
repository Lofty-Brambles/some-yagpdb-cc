{{$pag := (dbGet 1 "pgtemp").Value}}
{{if eq .Message.ID (index $pag 2)}}
  {{if and (eq $.Reaction.Emoji.Name "◀️") (ge (index $pag 1) 2)}}
    {{$i := index $pag 1}}
    {{dbSet 1 "pgtemp" (cslice (index $pag 0) (sub $i 1) (index $pag 2))}}
    {{editMessage nil (index $pag 2) (cembed
      "title" "Item Menu"
      "description" (joinStr "" "Your search matched the following items:\n" (index (index $pag 0) (sub $i 2)))
      "color" 0x2e3137
      "footer" (sdict 
        "text" "The search was conducted on "
        "icon_url" "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTtIcUVPAg3-QU2ptdXHNQ96Kxv0qKmIs1GSw&usqp=CAU")
      "timestamp" currentTime)}}
  {{end}}
  {{if and (eq $.Reaction.Emoji.Name "▶️") (lt (index $pag 1) (len (index $pag 0)))}}
    {{$i := index $pag 1}}
    {{dbSet 1 "pgtemp" (cslice (index $pag 0) (add $i 1) (index $pag 2))}}
    {{editMessage nil (index $pag 2) (cembed
      "title" "Item Menu"
      "description" (joinStr "" "Your search matched the following items:\n" (index (index $pag 0) $i))
      "color" 0x2e3137
      "footer" (sdict 
        "text" "The search was conducted on "
        "icon_url" "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTtIcUVPAg3-QU2ptdXHNQ96Kxv0qKmIs1GSw&usqp=CAU")
      "timestamp" currentTime)}} 
  {{end}}
{{end}}
