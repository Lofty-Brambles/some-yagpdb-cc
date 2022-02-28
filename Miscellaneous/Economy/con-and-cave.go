{{/*Trigger Type: Regex, Trigger: \A(?i);\s*(con+|caves?) */}}

{{$ci := (dbGet 123 "eco").Value.Get "icon"}}{{$fail := sdict}}{{$inc := sdict}}

{{$inp := reFind "(con|cave)" (lower .Message.Content)}}

{{$confailchance := $fail.Set "con" 0.40}}{{/*Fail chances for ;con*/}}
{{$cavefailchance := $fail.Set "cave" 0.25}}{{/*Fail chances for ;cave*/}}
{{$conpass := $inc.Set "conpass" (randInt 300 500)}}{{/*;con income range*/}}
{{$cavepass := $inc.Set "cavepass" (randInt 150 300)}}{{/*;cave income range*/}}
{{$confail := $inc.Set "confail" (randInt 50 150)}}{{/*;con income range*/}}
{{$cavefail := $inc.Set "cavefail" (randInt 25 75)}}{{/*;cave income range*/}}

{{$bin := ""}}{{$col := 0x2e3137}}{{$earn := ""}}
{{if ge (randInt 99|mult 0.01) ($fail.Get $inp)}}
 {{$bin = "pass"}}{{$col = 65480}}{{$val := $inc.Get (print $inc $bin)}}
 {{if eq $inp "con"}}{{$earn = print "You make " $ci $val "."}}
 {{else}}{{$earn = print "You scavenge " $ci $val "."}}{{end}}
 {{dbSet .User.ID "bal" (add ($inc.Get (print $inc $bin)) (dbGet .User.ID "bal").Value)}}
{{else}}
 {{$bin = "fail"}}{{$col = 16711680}}{{$val := $inc.Get (print $inc $bin)}}
 {{$earn = print "You lose " $ci $val "."}}{{$reset := sub $val (dbGet .User.ID "bal").Value}}
 {{if lt $reset 0}}{{$reset ￼= 0}}{{end}}
 {{dbSet .User.ID "bal" $reset}}
{{end}}
{{$replies := (dbGet 123 (print $inp $bin "reply")).Value}}

{{sendMessage nil (cembed 
 "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256"))
 "description" (print (index $replies (randInt (sub (len $replies) 1))) "\n" $earn)
 "footer" (sdict "text" "Check your new balance with ;bal")
 "color" $col)}}
