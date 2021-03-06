{{/*Trigger Type: Regex; Trigger: \A(?i);\s*u(?:se)(?: +|\z) */}}
{{if not .ExecData}}
 
{{$args := parseArgs 1 "Usage: `;use <Item ID>`" (carg "string" "")}}
{{$inv := (dbGet .User.ID "inv").Value}}{{$key := ($args.Get 0|lower)}}
{{if not ($inv.HasKey $key)}}
  {{sendMessage nil "<a:Cross:947957187928526868> **|** That is either an invalid Item ID or you do not own this item."}}
{{else}}
  {{$item := $inv.Get $key}}{{$quantity := sub ($item.Get "quantity") 1}}
  {{if $quantity}}{{$sike := $item.Set "quantity" $quantity}}{{$sikeu2 := $inv.Set $key $item}}
  {{else}}{{$nosike := $inv.Del $key}}{{end}}
  {{dbSet .User.ID "inv" $inv}}
  {{if eq $key "clover"}}
    {{$multi := div (randInt 10 100|toFloat) 100.0}}{{$sel := print (index (shuffle (cslice "work" "con" "cave")) 0) "multi"}}
    {{if lt (randInt 99) 50}}
      {{$misc := (dbGet .User.ID "misc").Value}}{{if not $misc}}{{$misc := sdict}}{{end}}
      {{if $misc.HasKey $sel}}
        {{$misc.Set $sel ($misc.Get $sel | add $multi)}}
      {{else}}
        {{$misc.Set $sel $multi}}
      {{end}}{{dbSet .User.ID "misc" $misc}}
      {{sendMessage nil "<a:rainbow_hype:951524938282573885> | The clover glows! You now have a multiplier on one of your actions for 2 days!"}}
      {{scheduleUniqueCC .CCID nil 86400 (print "clover" .User.ID) (sdict "key" "clover" "entry" $sel "user" .User.ID)}}
    {{else}}{{sendMessage nil "<a:CH_CatCry:951482071858151454> | The Clover fooled ye, it wasn't lucky..."}}{{end}}
  {{end}}
  {{if eq $key "hammer"}}
    {{if lt (randInt 69) 50}}
      {{$dash := (dbGet .User.ID "dash").Value}}{{if not $dash}}{{$dash = sdict "bbal" 1 "bquota" 20000 "brate" 0.02}}{{end}}
      {{$dash.Set "brate" (($dash.Get "brate")|mult 10.0)}}{{dbSet .User.ID "dash" $dash}}
      {{sendMessage nil "<a:rainbow_hype:951524938282573885> | The hammer struck! You now have 10x the bank return! Get that stonks!"}}
      {{scheduleUniqueCC .CCID nil 1209600 (print "hammer" .User.ID) (sdict "key" "hammer" "user" .User.ID)}}
    {{else}}{{sendMessage nil "<a:CH_CatCry:951482071858151454> | Poor you couldn't lift the hammer and got wasted..."}}{{end}}
  {{end}}
  {{if eq $key "locker"}}
    {{$dash := (dbGet .User.ID "dash").Value}}{{if not $dash}}{{$dash = sdict "bbal" 1 "bquota" 20000 "brate" 0.02}}{{end}}
    {{$rand := index (shuffle (cslice 1000 2000 3000 4000 5000)) 0}}
    {{$sikee := $dash.Set "bquota" ($dash.Get "bquota" | add $rand)}}
    {{dbSet .User.ID "dash" $dash}}
    {{sendMessage nil "<a:Tick:947957018675781662> | Your bank quota was increased! Check it with `;bank`"}}
  {{end}}
  {{if eq $key "map"}}
    {{sendMessage nil "???? | This item is still under construction for a later feature. Your balance has been refunded, meanwhile."}}
    {{dbSet .User.ID "bal" (add (dbGet .User.ID "bal").Value 2000)}}
  {{end}}
{{end}}
 
{{else}}
 
  {{if eq .ExecData.key "clover"}}
    {{$misc := (dbGet .ExecData.user "misc").Value}}{{$silent := $misc.Del .ExecData.entry}}{{dbSet .ExecData.user "misc" $misc}}
  {{end}}
  {{if eq .ExecData.key "hammer"}}
    {{$dash := (dbGet .ExecData.user "dash").Value}}{{$set := $dash.Set "brate" (div ($dash.Get "brate") 10)}}{{dbSet .ExecData.user "dash" $dash}}
  {{end}}
 
{{end}}
