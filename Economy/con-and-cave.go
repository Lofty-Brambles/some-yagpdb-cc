{{/*Trigger Type: Regex, Trigger: \A(?i);\s*(con+|caves?)(?: +|\z) */}}
 
{{$ci := (dbGet 123 "eco").Value.Get "icon"}}{{$fail := sdict}}{{$inc := sdict}}{{$con := sdict}}{{$cd := sdict}}
 
{{$inp := reFind "(con|cave)" (lower .Message.Content)}}
 
{{$confailchance := $fail.Set "con" 0.40}}{{/*Fail chances for ;con*/}}
{{$cavefailchance := $fail.Set "cave" 0.25}}{{/*Fail chances for ;cave*/}}
{{$conpass := $inc.Set "conpass" (randInt 300 500 | toFloat)}}{{/*;con income range*/}}
{{$cavepass := $inc.Set "cavepass" (randInt 150 300 | toFloat)}}{{/*;cave income range*/}}
{{$confail := $inc.Set "confail" (randInt 50 150 | toFloat)}}{{/*;con income range*/}}
{{$cavefail := $inc.Set "cavefail" (randInt 25 75 | toFloat)}}{{/*;cave income range*/}}
{{$concd := $cd.Set "con" 900}}{{/*Cooldown on ;con*/}}
{{$cavecd := $cd.Set "cave" 600}}{{/*Cooldown on ;cave*/}}
 
{{$misc := (dbGet .User.ID "misc").Value}}{{if not $misc}}{{$misc = sdict}}{{end}}
 
{{if lt currentTime.Unix (toInt ($misc.Get (print $inp "cd")))}}
 
{{sendMessage nil (cembed 
 "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256"))
 "description" (print "<a:no:830591739126611969> You are still on cooldown for this command.\nYou will be able to use this command <t:" ($misc.Get (print $inp "cd")) ":R>.")
 "color" 16711680)}}
 
{{else}}
 
    {{$bin := ""}}
    {{$col := 0x2e3137}}
    {{$earn := ""}}
    {{$foot := ""}}
    
    {{if ge (randInt 99|mult 0.01) ($fail.Get $inp)}}
    
        {{$bin = "pass"}}
        {{$col = 65480}}
        {{$val := $inc.Get (print $inp $bin)}}
        
        {{$reval := $val}}
        
        {{if $misc.HasKey ($m := print $inp "multi")}}
        
            {{$reval = add $m 1.0|mult $val|toInt}}
            {{$foot = print " | Multiplier: " ($misc.Get $m|mult 100.0|round) "% (+" ($misc.Get $m|mult $val|round) ")"}}
            
        {{end}}
        
        {{if eq $inp "con"}}
            {{$earn = print "You make " $ci (toInt $reval) "."}}
        {{else}}
            {{$earn = print "You scavenge " $ci (toInt $reval) "."}}
        {{end}}
        
        {{dbSet .User.ID "bal" (add (round $reval) (dbGet .User.ID "bal").Value)}}
        
    {{else}}
    
        {{$bin = "fail"}}
        {{$col = 16711680}}
        {{$val := $inc.Get (print $inp $bin)}}
 
        {{$earn = print "You lose " $ci (toInt $val)}}
        {{$reset := sub (dbGet .User.ID "bal").Value (toInt $val) | toFloat}}
        
        {{if lt $reset 0.00}}
            {{$reset = 0}}
        {{end}}
        
        {{dbSet .User.ID "bal" (toInt $reset)}}
        
    {{end}}
    
    {{$replies := (dbGet 123 (print $inp $bin "reply")).Value}}
    {{$sil := $misc.Set (print $inp "cd") (add ($cd.Get $inp) currentTime.Unix)}}
    {{dbSet .User.ID "misc" $misc}}
    
    {{sendMessage nil (cembed 
       "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256"))
       "description" (print (index $replies (randInt (sub (len $replies) 1))) "\n" $earn)
       "footer" (sdict "text" (print "Check your new balance with ;bal" $foot))
       "color" $col)}}
       
{{end}}
