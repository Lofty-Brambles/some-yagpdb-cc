{{$admin := 944551206536245258}}
{{$mod := 944551072452730951}}
{{$scur := true}}{{$denom := 27000}}{{$c2name := "LMC"}}{{$emoji := "🍉"}}
{{$cd := 21600}}
{{if ge (len .CmdArgs) 1}}{{$c := (index .CmdArgs 0|lower)}}
  {{if eq $c "find"}}
    {{$args := parseArgs 2 "Usage: `!!find  <Item name>`" (carg "string" "") (carg "string" "regex")}}
    {{$dz := lower ($args.Get 1)}}{{$co := 1}}{{$it := cslice}}{{$pg := cslice}}{{$d := ""}}
  {{range $i, $e := (dbGet 0 "id-slice").Value}}
    {{- if or (eq $dz (reFind $dz (lower .))) (eq $dz "all")}}
      {{$it = $it.Append (joinStr "" "> `" (($co = add $co 1)|sub 1) ".` __" $e "__ (ID: `" (add $i 1) "`)\n")}}{{end}}{{end -}}{{$l := len $it}}
    {{if eq $l 0}}{{$d = "`No items matched your search`"}}
    {{else if le $l 20}}{{range $it}}{{- $d = joinStr "" $d .}}{{end -}}
    {{else}}{{range $i, $e := $it}}{{- if and (ne $i 0) (eq (mod $i 20) 0.0)}}{{$pg = $pg.Append $d}}{{$d = $e}}{{else}}{{$d = joinStr "" $d $e}}{{end}}{{if eq $i (sub $l 1)}}{{$pg = $pg.Append $d}}{{$d = index $pg 0}}{{end}}{{end -}}{{end}}
    {{$id := sendMessageRetID nil (cembed "title" "Item Menu" "description" (joinStr "" "Your search matched the following items:\n" $d) "color" 0x2e3137 "footer" (sdict "text" "The search was conducted" "icon_url" "https://bit.ly/3JKCscP") "timestamp" currentTime)}}
    {{if $pg}}{{dbSet 1 "pgtemp" (cslice $pg 1 $id)}}{{addMessageReactions nil $id "◀️" "▶️"}}{{end}}
    {{end}}
 
  {{if eq $c "add"}}
{{if not (hasRoleID $admin)}}{{sendMessage nil "❌ You aren't allowed to add items."}}{{else}}
    {{$args := parseArgs 4 "Usage: `!!add  <Cost> <Picture URL> <Name>`" (carg "string" "") (carg "int" "Price") (carg "string" "url") (carg "string" "Object Name")}}
    {{dbSet 0 "id-slice" ((dbGet 0 "id-slice").Value.Append ($args.Get 3))}}
    {{$last := dbIncr 1 "lastpos" 1}}
    {{$key := toString (toInt (roundFloor (div (sub $last 1) 50)))}}
    {{$entry := sdict "item" ($args.Get 3) "tn" ($args.Get 2) "cost" (cslice ($args.Get 1) ($args.Get 1) ($args.Get 1)) "update" currentTime "last" .User.ID "out" false}}
    {{$re := ($args.Get 1)}}
    {{if not (dbGet 420 $key)}}{{dbSet 420 $key (cslice $entry)}}{{dbSet 421 $key (cslice $re)}}
    {{else}}
      {{dbSet 420 $key ((dbGet 420 $key).Value.Append $entry)}}{{dbSet 421 $key ((dbGet 421 $key).Value.Append $re)}}
    {{end}}
    {{sendMessage nil (print "The item has been added!\nLook it up with the command: `!!price " (sub $last 1) "`")}}{{end}}{{end}}
 
  {{if eq $c "price"}}
    {{$args := parseArgs 2 "Usage: `!!price <ID>`" (carg "string" "") (carg "int" "ID")}}
    {{if ge (toFloat ($args.Get 1)) (dbGet 1 "lastpos").Value}}{{sendMessage nil "⚠️ This item couldn't be found. Please recheck the ID."}}{{else}}
    {{$key := toString (toInt (roundFloor (div ($args.Get 1) 50)))}}
    {{$place := toInt (mod ($args.Get 1) 50)}}
    {{$item := index (dbGet 420 $key).Value $place}}{{$pslice := $item.Get "cost"}}
    {{$price := div (round (mult (div (add (index $pslice 0) (index $pslice 1) (index $pslice 2)) 3) 100)) 100}}
    {{$disprice := humanizeThousands $price}}
    {{$lmc := div (round (mult (div $price $denom) 100)) 100}}
    {{$dislmc := slice (printf "%f" (mod $lmc 1) ) 1 4 | printf "%s%s" (humanizeThousands $lmc)}}{{$desc := ""}}
    {{$addin := ""}}{{if $scur}}{{$addin = joinStr "" "`\n> The Price of this item in " $c2name " is: " $emoji "`" $dislmc "`"}}{{end}}
    {{if ($item.Get "out")}}{{$desc = "> This Item is outdated. Sorry!"}}
    {{else}}{{$desc = (joinStr "" "> The Price of this item is: <:Monei:944250681357926430>`" $disprice $addin)}}{{end}}
    {{sendMessage nil (cembed "title" ($item.Get "item") "thumbnail" (sdict "url" ($item.Get "tn")) "author" (sdict "name" (joinStr "" "Last Updated by: " ((userArg ($item.Get "last")).Username)) "icon_url" ((userArg ($item.Get "last")).AvatarURL "256")) "color" 0x2e3137 "description" $desc "footer" (sdict "text" "This was last updated" "icon_url" "https://bit.ly/356aGsm") "timestamp" ($item.Get "update"))}}{{end}}{{end}}
 
  {{if eq $c "update"}}
{{if not (hasRoleID $mod)}}
  {{sendMessage nil "❌ You aren't allowed to update item prices."}}
{{else}}
  {{$args := parseArgs 3 "Usage: `!!update <ID> <price>`" (carg "string" "") (carg "int" "ID") (carg "int" "Price")}}{{$a := $args.Get 1}}{{$k := (dbGet .User.ID "update").Value}}
  {{if ge (toFloat $a) (dbGet 1 "lastpos").Value}}{{sendMessage nil "⚠️ This item couldn't be found. Please recheck the ID."}}
  {{else if and ($k.HasKey (toString $a)) (ge (toInt ($k.Get (toString $a))) currentTime.Unix)}}
  {{sendMessage nil "❌ You can only update an item every 6 hours."}}{{else}}
  {{$key := toString (toInt (roundFloor (div $a 50)))}}
  {{$place := toInt (mod $a 50)}}
  {{$item := (dbGet 420 $key).Value}}{{$res := (dbGet 421 $key).Value}}
  {{$count := (index $item $place).Get "cost"}}
  {{$temp := (slice $count 1).Append ($args.Get 2)}}{{$save := index $count 0}}
  {{$reres := $res.Set $place $save}}
  {{$s := (index $item $place).Set "cost" $temp}}{{$s := (index $item $place).Set "update" currentTime}}{{$s := (index $item $place).Set "last" .User.ID}}{{$p := (dbGet .User.ID "update").Value}}{{$s := $p.Set (toString $a) (add currentTime.Unix $cd)}}
  {{dbSet 420 $key $item}}{{dbSet 421 $key $res}}{{dbSet .User.ID "update" $p}}
  {{sendMessage nil (joinStr "" "The item has been updated!\nLook it up with the command: `!!price " ($args.Get 1) "`")}}{{end}}{{end}}{{end}}
 
  {{if eq $c "remprice"}}
{{if not (hasRoleID $admin)}}
  {{sendMessage nil "❌ You can't remove an item's price."}}
{{else}}
  {{$args := parseArgs 2 "Usage: `!!remprice <ID>`" (carg "string" "") (carg "int" "ID")}}
  {{if ge (toFloat ($args.Get 1)) (dbGet 1 "lastpos").Value}}{{sendMessage nil "⚠️ | This item couldn't be found. Please recheck the ID."}}{{else}}
  {{$key := toString (toInt (roundFloor (div ($args.Get 1) 50)))}}
  {{$place := toInt (mod ($args.Get 1) 50)}}
  {{$res := index (dbGet 421 $key).Value $place}}
  {{$db := (dbGet 420 $key).Value}}
  {{$item := (index $db $place).Get "cost"}}
  {{$silent := (index $db $place).Set "cost" (cslice $res (index $item 0) (index $item 1))}}
  {{dbSet 420 $key $db}}
  {{sendMessage nil (joinStr "" "The item's last price has been removed!\nLook it up with the command: `!!price " ($args.Get 1) "`")}}{{end}}{{end}}{{end}}
 
  {{if eq $c "remlastitem"}}
{{if not (hasRoleID $admin)}}
  {{sendMessage nil "❌ You can't remove an item."}}
{{else}}
  {{$idslice := (dbGet 0 "id-slice").Value}}
  {{dbSet 0 "id-slice" (slice $idslice 0 (sub (len $idslice) 1))}}
  {{$silent := add (dbIncr 1 "lastpos" (toFloat -1)) 1}}
  {{$key := toString (toInt (roundFloor (div $silent 50)))}}
  {{$main := (dbGet 420 $key).Value}}{{$rev := (dbGet 421 $key).Value}}
  {{dbSet 420 $key (slice $main 0 (sub (len $main) 1))}}
  {{dbSet 421 $key (slice $rev 0 (sub (len $rev) 1))}}
  {{sendMessage nil "Item removed"}}{{end}}{{end}}
 
  {{if eq $c "outdate"}}
{{if not (hasRoleID $mod)}}
  {{sendMessage nil "❌ You cannot mark an item as outdated."}}
{{else}}
  {{$args := parseArgs 3 "Usage: `!!outdate <ID> <true/false>`" (carg "string" "") (carg "int" "ID") (carg "string" "Bool")}}
  {{if ge (toFloat ($args.Get 1)) (dbGet 1 "lastpos").Value}}{{sendMessage nil "⚠️ | This item couldn't be found. Please recheck the ID."}}{{else}}
  {{$key := toString (toInt (roundFloor (div ($args.Get 1) 50)))}}
  {{$place := toInt (mod ($args.Get 1) 50)}}
  {{if eq "true" ($args.Get 2)}}
    {{$db := (dbGet 420 $key).Value}}{{$silent := (index $db $place).Set "out" true}}{{dbSet 420 $key $db}}{{sendMessage nil (print "The item has been marked as outdated!\nLook it up with the command: `!!price " ($args.Get 1) "`")}}
  {{else if eq "false" ($args.Get 2)}}
    {{$db := (dbGet 420 $key).Value}}{{$silent := (index $db $place).Set "out" false}}{{dbSet 420 $key $db}}{{sendMessage nil (print "The item is no longer outdated!\nLook it up with the command: `!!price " ($args.Get 1) "`")}}
  {{else}}{{sendMessage nil "Usage: `!!outdate <ID> <true/false>`"}}{{end}}
  {{end}}{{end}}{{end}}
  
{{if eq $c "help"}}
  {{$desc := "\n\n➼ `!!find <Item name>/<Regex>`\n <:SubMenu:944115521794764800>Use this command to look up the ID of the item whose price you want to know.\n\n➼ `!!price <Item ID>`\n <:SubMenu:944115521794764800>Use this command to look up an Item's price using it's ID"}}{{$foot := sdict "text" "Member Help Menu" "icon_url" "https://bit.ly/3h7fYWW"}}
  {{if (hasRoleID $mod)}}
  {{$desc = joinStr "" $desc "\n\n➼ `!!update <Item ID> <Price>`\n <:SubMenu:944115521794764800>Use this command to update the price of an Item.\n\n➼ `!!outdate <Item ID> <true/false>`\n <:SubMenu:944115521794764800>Use this command to mark an item as __outdated__ or remove it's outdated status"}}{{$foot = sdict "text" "Moderator Help Menu" "icon_url" "https://bit.ly/3BX4iQx"}}{{end}}
  {{if (hasRoleID $admin)}}
    {{$desc = joinStr "" $desc "\n\n➼ `!!add  <Cost> <Picture URL> <Name>`\n <:SubMenu:944115521794764800>Use this command to add an item to the price guide.\n\n➼ `!!remprice <Item ID>`\n <:SubMenu:944115521794764800>Use this command to remove the last price of an item.\n\n➼ `!!remlastitem`\n <:SubMenu:944115521794764800>Use this command to remove the last item added to the price guide."}}{{$foot = sdict "text" "Admin Help Menu" "icon_url" "https://bit.ly/3BMBU3d"}}
  {{end}}
  {{sendMessage nil (cembed "title" "Help Menu" "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256")) "footer" $foot "color" 0x2e3137 "description" $desc)}}{{end}}
{{end}}
