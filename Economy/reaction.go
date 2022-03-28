{{if .ReactionAdded}}
 
  {{$pagin := (dbGet 123 "pagination").Value}}
  {{$fun := sdict}}{{$key := ""}}
  {{$eco := (dbGet 123 "eco").Value}}{{$ci := $eco.Get "icon"}}
 
  {{range $k, $v := $pagin}}
    {{- if eq (.Get "id") .Message.ID -}}
      {{- $fun = $v }}{{- $key = $k -}}
    {{- end -}}
  {{- end -}}
 
  {{if $key}}
  
    {{if eq .Reaction.Emoji.Name "left"}}
    
      {{if or (eq $key "shop") (eq $key "inv")}}
        {{if ge ($fun.Get "page") 2}}
        
          {{$embed := structToSdict (index .Message.Embeds 0)}}
          {{range $k, $v := $embed}}
            {{- if eq (kindOf $v) "struct" -}}
            {{- $embed.Set $k (structToSdict $v) -}}
            {{- end -}}
          {{- end -}}
          {{if $embed.Author}}
            {{$embed.Author.Set "icon_url" $embed.Author.IconURL}}
          {{end}}
          {{if $embed.Footer}}
            {{$embed.Footer.Set "icon_url" $embed.Footer.IconURL}}
          {{end}}
          {{$page := sub ($fun.Get "page") 1}}{{$sil := $fun.Set "page" $page}}
          {{$embed.Set "description" (index ($fun.Get "pages") ($fun.Get "page"| add -1))}}
          {{if eq $key "inv"}}
            {{$embed.Footer.Set "text" (print "You can use an item with ;use <ID> | Page " $page " of " ($fun.Get "pages"|len))}}
          {{else if eq $key "shop"}}
            {{$embed.Footer.Set "text" (print "You can buy an item with ;buy <ID> | Page " $page " of " ($fun.Get "pages"|len))}}
          {{end}}
          {{editMessage nil ($fun.Get "id") (cembed $embed)}}
          {{$sil := $pagin.Set $key $fun}}
          {{dbSet 123 "pagination" $pagin}}
          
        {{end}}
      {{end}}
      
      {{if eq $key "lb"}}
        {{if ge ($fun.Get "page") 2}}
        
          {{$page := sub ($fun.Get "page") 1}}{{$sil := $fun.Set "page" $page}}
          {{$totpages := $fun.Get "total"}}
          {{$desc := ""}}
          {{$skip := mult (sub $page 1) 10}}
          {{$poiple := dbTopEntries "bal" 10 $skip}}
          {{$serial := $skip}}
          {{- range $poiple -}}
            {{- $serial = add $serial 1 -}}
            {{- $desc = print $desc "\n➼  **" $serial ".** [" .User.Username "](" (.User.AvatarURL "256") ") **|** Balance: " $ci "`" (toInt .Value) "`" -}}
          {{- end -}}
          {{editMessage nil ($fun.Get "id") (complexMessageEdit "embed" (cembed 
            "title" "Leaderboards"
            "thumbnail" (sdict "url" "https://cdn-icons-png.flaticon.com/512/4489/4489655.png")
            "color" 0x2e3137
            "timestamp" currentTime
            "footer" (sdict "text" (print "Page: " $page " of " $totpages " | Searched ")
                                "icon_url" "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsB0ZctJzyh8UWoLg8yRV7COvF3slJY4bjh2H36GPQcBWKfOMlJKahmNI&s=10")
            "description" $desc))}}
          {{$rest := $pagin.Set $key $fun}}
          {{dbSet 123 "pagination" $pagin}}
        
        {{end}}
      {{end}}
      
    {{end}}
    
    {{if eq .Reaction.Emoji.Name "right"}}
    
      {{if or (eq $key "shop") (eq $key "inv")}}
        {{if lt ($fun.Get "page") ($fun.Get "pages" | len)}}
        
          {{$embed := structToSdict (index .Message.Embeds 0)}}
          {{range $k, $v := $embed}}
            {{- if eq (kindOf $v) "struct" -}}
            {{- $embed.Set $k (structToSdict $v) -}}
            {{- end -}}
          {{- end -}}
          {{if $embed.Author}}
            {{$embed.Author.Set "icon_url" $embed.Author.IconURL}}
          {{end}}
          {{if $embed.Footer}}
            {{$embed.Footer.Set "icon_url" $embed.Footer.IconURL}}
          {{end}}
          {{$page := add ($fun.Get "page") 1}}{{$sil := $fun.Set "page" $page}}
          {{$embed.Set "description" (index ($fun.Get "pages") ($fun.Get "page"| add 1))}}
          {{if eq $key "inv"}}
            {{$embed.Footer.Set "text" (print "You can use an item with ;use <ID> | Page " $page " of " ($fun.Get "pages"|len))}}
          {{else if eq $key "shop"}}
            {{$embed.Footer.Set "text" (print "You can buy an item with ;buy <ID> | Page " $page " of " ($fun.Get "pages"|len))}}
          {{end}}
          {{editMessage nil ($fun.Get "id") (cembed $embed)}}
          {{$sil := $pagin.Set $key $fun}}
          {{dbSet 123 "pagination" $pagin}}
          
        {{end}}
      {{end}}
      
      {{if eq $key "lb"}}
        {{if lt ($fun.Get "page") ($fun.Get "total")}}
        
          {{$page := add ($fun.Get "page") 1}}{{$sil := $fun.Set "page" $page}}
          {{$totpages := $fun.Get "total"}}
          {{$desc := ""}}
          {{$skip := mult (sub $page 1) 10}}
          {{$poiple := dbTopEntries "bal" 10 $skip}}
          {{$serial := $skip}}
          {{- range $poiple -}}
            {{- $serial = add $serial 1 -}}
            {{- $desc = print $desc "\n➼  **" $serial ".** [" .User.Username "](" (.User.AvatarURL "256") ") **|** Balance: " $ci "`" (toInt .Value) "`" -}}
          {{- end -}}
          {{editMessage nil ($fun.Get "id") (complexMessageEdit "embed" (cembed 
            "title" "Leaderboards"
            "thumbnail" (sdict "url" "https://cdn-icons-png.flaticon.com/512/4489/4489655.png")
            "color" 0x2e3137
            "timestamp" currentTime
            "footer" (sdict "text" (print "Page: " $page " of " $totpages " | Searched ")
                                "icon_url" "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsB0ZctJzyh8UWoLg8yRV7COvF3slJY4bjh2H36GPQcBWKfOMlJKahmNI&s=10")
            "description" $desc))}}
          {{$rest := $pagin.Set $key $fun}}
          {{dbSet 123 "pagination" $pagin}}
          
        {{end}}
      {{end}}
      
    {{end}}
    
  {{end}}
  
{{end}}
