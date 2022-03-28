{{/*Trigger Type: Regex, Trigger: \A */}}
 
{{$ci := (dbGet 123 "eco").Value.Get "icon"}}
{{$start := `;work`}}
{{$misc := (dbGet .User.ID "misc").Value}}
{{if not $misc}}{{$misc = sdict}}{{end}}
{{$dbcool := $misc.Get "workcd"}}
{{$dbResp := (dbGet .User.ID "workResp").Value}}
 
{{$maincd := 14400}}
{{$cd1 := 10}}
{{$cd2 := 25}}
{{$cd3 := 10}}
{{$cd4 := 10}}
 
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256")) "color" 0x2e3137}}
 
{{/*Bank CC updater*/}}
{{$dash := (dbGet .User.ID "dash").Value}}
{{if not $dash}}{{$dash = sdict "bbal" 0 "bquota" 20000 "brate" 0.02 "bcd" currentTime.Unix}}{{end}}
{{if ge currentTime.Unix ($dash.Get "bcd")}}
  {{$sil1 := $dash.Set "bbal" ($dash.Get "bbal"| mult ($dash.Get "brate"| add 1.0)| toInt)}}
  {{$sil2 := $dash.Set "bcd" ($dash.Get "bcd" | add 86400)}}
  {{dbSet .User.ID "dash" $dash}}
{{end}}
 
{{if .ExecData}}
 
 {{$embed.Set "title" "Failed!"}}
 {{$embed.Set "description" (print "<a:Cross:947957187928526868> You did not complete your work on time.\nYou were paid " $ci "`150` for late work.")}}
 {{$embed.Set "color" 16711680}}
 {{sendMessage nil (cembed $embed)}}
 {{dbSet .User.ID "bal" ((dbGet .User.ID "bal").Value|add 150)}}
 
{{else}}
 
 {{if not $dbResp}}
 
  {{$tempor := (print `\A(?i)` $start)}}
  {{if reFind $tempor (lower .Message.Content)}}
  
   {{if lt currentTime.Unix (toInt $dbcool)}}
    {{$embed.Set "description" (print "<a:Cross:947957187928526868> You are still on cooldown for this command.\nYou will be able to use this command in <t:" $dbcool ":R>.")}}
    {{$embed.Set "color" 16711680}}
    {{sendMessage nil (cembed $embed)}}
   {{else}}
  
   {{$task := index (cslice "Passing the code" "Word Wants" "Number me right" "Introvert's Dilemma") (randInt 4)}}{{$task = "Word Wants"}}
   {{$gtemp := cslice "Achatina" "Angler" "Araneo" "Ammonite" "Beelzebufo" "Castoroides" "Carnotaurus" "Carbonemy" "Cnidaria" "Compy" "Daeodon" "Direwolf" "Diplodocus" "Eurypterid" "Gallimimus" "Ickthyosaurus" "Iguanadon" "Jerboa" "Kairuku" "Lystrosaurus" "Mammoth" "Megaloceros" "Megalodon" "Meganeura" "Megatherium" "Moschop" "Otter" "Oviraptor" "Parasaur" "Pelagornis" "Phiomia" "Piranha" "Purlovia" "Quetzal" "Raptor" "Sabertooth" "Tapejara" "Titanoboa" "Triceratops" "Trilobite" "Troodon" "Unicorn"}}
   {{$embed.Set "title" (print "Work - " $task)}}{{$resp := ""}}
   
   {{if eq $task "Passing the code"}}
    {{$temp := cslice "A" "B" "C" "D" "E" "F" "G" "H" "I" "J" "K" "L" "M" "N" "O" "P" "Q" "R" "S" "T" "U" "V" "W" "X" "Y" "Z" "1" "2" "3" "4" "5" "6" "7" "8" "9" "0" }}{{$code := ""}}
    {{range seq 0 4}}{{- $code = print $code (index (shuffle $temp) 0)}}{{end -}}
    {{$embed.Set "description" (print "You have exactly " $cd1 " seconds to type out the four digit code given.")}}
    {{$randBG := print "http://www.gstatic.com/webp/gallery/" (randInt 1 6) ".png"}}
    {{$coderxt := print "https://api.memegen.link/images/custom/_/" $code ".png?background=" $randBG}}
    {{$embed.Set "image" (sdict "url" $coderxt)}}{{$resp = lower $code}}
    {{sendMessage nil (cembed $embed)}}
    {{dbSetExpire .User.ID "workResp" $resp $cd1}}
    {{scheduleUniqueCC .CCID nil $cd1 (print "timed" .User.ID) 1}}
    
   {{else if eq $task "Word Wants"}}
    {{$shuf := shuffle $gtemp}}{{$code := ""}}{{$resp := `(?i)`}}
    {{range seq 0 4}}{{- $code = print $code "\n" (index $shuf .)}}{{$resp = print $resp `\s*` (index $shuf .|lower)}}{{end -}}
    {{$resp = print $resp `\s*`}}
    {{$embed.Set "description" (print "You have exactly " $cd2 " seconds to type out the four following words, __with spaces in between, in one message.__ ```" $code "```")}}
    {{sendMessage nil (cembed $embed)}}
    {{dbSetExpire .User.ID "workResp" $resp $cd2}}
    {{scheduleUniqueCC .CCID nil $cd2 (print "timed" .User.ID) 1}}
    
   {{else if eq $task "Number me right"}}
    {{$shuf := shuffle $gtemp}}{{$num := randInt 1 4}}
    {{$di := sdict "1" "<:bfd1:947955364362604544>" "2" "<:bfd2:947955424617979984>" "3" "<:bfd3:947955471694827530>"}}
    {{$code := print "<:bfd1:947955364362604544> `" (index $shuf 0) "`\n<:bfd2:947955424617979984> `" (index $shuf 1) "`\n<:bfd3:947955471694827530> `" (index $shuf 2) "`"}}
    {{$embed.Set "description" (print "Carefully read the following words. You will have exactly " $cd3 " seconds to type out the word next to the given number.\n" $code)}}
    {{$id := sendMessageRetID nil (cembed $embed)}}
    {{$embed.Set "description" (print "What was the word next to " ($di.Get (toString $num)) "? You have " $cd3 " seconds to answer.")}}
    {{sleep 10}}{{editMessage nil $id (cembed $embed)}}
    {{$resp = index $shuf (sub $num 1)|lower}}
    {{dbSetExpire .User.ID "workResp" $resp (add $cd3 5)}}
    {{scheduleUniqueCC .CCID nil (add $cd3 5) (print "timed" .User.ID) 1}}
    
   {{else if eq $task "Introvert's Dilemma"}}
    {{$shuf := cslice "😊" "🤣" "😍" "😎" "🤪" "😒" "🥺" "😠" "😱" "😐" "😲" "🤢"|shuffle}}{{$num := randInt 1 4}}
    {{$di := sdict "1" "<:bfd1:947955364362604544>" "2" "<:bfd2:947955424617979984>" "3" "<:bfd3:947955471694827530>"}}
    {{$code := print "<:bfd1:947955364362604544> `" (index $shuf 0) "`\n<:bfd2:947955424617979984> `" (index $shuf 1) "`\n<:bfd3:947955471694827530> `" (index $shuf 2) "`"}}
    {{$embed.Set "description" (print "Carefully note the following emojis. You will have exactly " $cd4 " seconds to send the emoji next to the given number.\n" $code)}}
    {{$id := sendMessageRetID nil (cembed $embed)}}
    {{$embed.Set "description" (print "What was the emoji next to " ($di.Get (toString $num)) "? You have " $cd4 " seconds to answer.")}}
    {{sleep 10}}{{editMessage nil $id (cembed $embed)}}
    {{$resp = index $shuf (sub $num 1)}}
    {{dbSetExpire .User.ID "workResp" $resp (add $cd4 5)}}
    {{scheduleUniqueCC .CCID nil (add $cd4 5) (print "timed" .User.ID) 1}}
   {{end}}
   
   {{$sil := $misc.Set "workcd" (add $maincd currentTime.Unix)}}
   {{dbSet .User.ID "misc" $misc}}
   {{end}}
  {{end}}
  
 {{else}}
  {{if reFind (joinStr `` `` $dbResp ``) .Message.Content}}
   {{$embed.Set "title" "Work well done!"}}
   {{$embed.Set "description" (print "<a:Tick:947957018675781662> You did completed your work on time.\nYou were paid " $ci "`1500` for some good ol' work.")}}
   {{$embed.Set "color" 65480}}
   {{sendMessage nil (cembed $embed)}}{{dbDel .User.ID "workResp"}}
   {{dbSet .User.ID "bal" ((dbGet .User.ID "bal").Value|add 1500)}}
  {{else}}
   {{$embed.Set "title" "Failed!"}}
   {{$embed.Set "description" (print "<a:Cross:947957187928526868> You did not do your work properly.\nYou were paid " $ci "`150` for bad work.")}}
   {{$embed.Set "color" 16711680}}
   {{sendMessage nil (cembed $embed)}}{{dbDel .User.ID "workResp"}}
   {{dbSet .User.ID "bal" ((dbGet .User.ID "bal").Value|add 150)}}
  {{end}}
   {{cancelScheduledUniqueCC .CCID (print "timed" .User.ID)}}
   
 {{end}}
{{end}}
