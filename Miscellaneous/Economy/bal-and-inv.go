{{/*Trigger Type: Regex, Trigger: \A(?i);\s*(bal(ance)?|inv(entory)?) */}}

{{$ci := (dbGet 123 "eco").Value.Get "icon"}}

{{$inp := reFind "(bal|inv)" (lower .Message.Content)}}
{{$a := parseArgs 0 (print "Usage: `;" $inp "<User:optional>`") (carg "userid" "")}}
{{$user := .User.ID}}{{if $a.IsSet 0}}{{$user = (userArg ($a.Get 0)).ID}}{{end}}

{{/*;bal*/}}
{{if eq $inp "bal"}}
 {{$dash := (dbGet $user.ID "dash").Value}}
 {{$perc := div (round (mult (div ($dash.Get "bbal") ($dash.Get "bquota")) 10)) 10}}
 {{sendMessage nil (cembed 
  "title" (print $user.Username " - Balance") 
  "color" 0x2e3137 
  "description" (print "**Pocket**: " $ci "`" (humanizeThousands (dbGet $user.ID "bal")) "`\n**Bank Balance**: " $ci "`" (humanizeThousands ($dash.Get "bbal")) "` / `" (humanizeThousands ($dash.Get "bquota")) "` (`" $perc "`)") 
  "footer" (sdict "icon_url" ($user.AvatarURL "256") 
  "text" "Balance checked") 
  "timestamp" currentTime)}}
{{end}}

{{/*;inv*/}}
{{if eq $inp "inv"}}
 {{$inv := (dbGet $user.ID "inv").Value}}{{$plrinv := cslice}}{{$pagination := sdict}}{{$desc := ""}}
 {{- range $k, $v := $inv -}}
  {{- $plrinv = $plrinv.Append (print ($v.Get "icon") " - **" ($v.Get "name") "**\n__ID__ `" $k "` **|** Left : `" ($v.Get "quantity") "`") -}}
 {{else}}
  {{- $plrinv = $plrinv.Append (print "You have __no items__ in your inventory.\nYou can look at the shop, using `;shop`.")}}
 {{- end -}}
 {{$totp := 1}}
 {{if gt ($x := len $plrinv) 10}}
  {{$rem := toInt (mod $x 10)}}{{$div := div $x 10}}{{$totp = $div}}{{$pages := cslice}}
  {{if ne $rem 0}}{{$totp := add $totp 1}}{{end}}
  {{range seq 1 $totp}}
   {{- $end := sub . 1|mult 10|add 9 -}}
   {{- if eq . $totp}}{{$end = sub $x 1}}{{end -}}
   {{- $pages := $pages.Append (joinStr "\n" (slice $plrinv (sub . 1|mult 10) $end) -}}
  {{end}}
  {{$pagination = sdict "pages" $pages "page" 1 "id" 0}}
 {{else}}
  {{$desc = joinStr "\n" $plrinv}}
 {{end}}
 {{$id := sendMessageRetID nil (cembed 
  "author" (sdict "name" (print $user.Username " - Inventory") "icon_url" ($user.AvatarURL "256"))
  "color" 0x2e3137
  "description" $desc
  "footer" (sdict "text" (print "You can use an item with - ;use <ID> | Page 1 of " $totp)))}}
 {{$silent := $pagination.Set "id" $id}}
 {{if $pagination.HasKey "pages"}}
  {{$temp := (dbGet 123 "pagination").Value}}{{$tempv := $temp.Set "inv" $pagination}}
  {{addReactions "◀️" "▶️"}}{{dbSet 123 "pagination" $temp}}
 {{end}}
