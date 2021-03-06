{{/*Trigger Type: Regex; Trigger: \A(?i);\s*g(?:i[vb]e?)?(?: +|\z)*/}}
 
{{$ci := (dbGet 123 "eco").Value.Get "icon"}}
    {{if .ExecData}}
        {{sendMessage nil (print (.ExecData.user | userArg).Mention ", your transaction has been timed out for failing verification.")}}
    {{else}}
        {{if $info := (dbGet .User.ID "giveResp").Value}}
            {{if reFind (print `(?i)` ((dbGet .User.ID "giveResp").Value.Get "code")) .Message.Content}}
                {{dbSet .User.ID "bal" ($info.Get "amt" | sub (dbGet .User.ID "bal").Value)}}
                {{dbSet ($info.Get "given" | toInt) "bal" ($info.Get "amt" | add (dbGet ($info.Get "given"|toInt) "bal").Value)}}
                {{dbDel .User.ID "giveResp"}}
                {{cancelScheduledUniqueCC .CCID (print "give " .User.ID)}}
                {{sendMessage nil (print .User.Mention ", you have successfully given " $ci "`" ($info.Get "amt") "` to " ($info.Get "given" | userArg).Mention ".")}}
            {{end}}
        {{else}}
            {{$args := parseArgs 2 "Usage: `;give <@User/UserID> <Amount>`" (carg "userid" "") (carg "int" "")}}
            {{$amt := $args.Get 1}}
            {{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256")) "color" 0x2e3137}}
            {{if gt $amt ((dbGet .User.ID "bal").Value | toInt)}}
                {{$embed.Set "description" "You do not have enough balance to give this amount."}}
            {{else if eq .User.ID ($args.Get 0)}}
                {{$embed.Set "description" "You can't give yourself money, no cheating!"}}
            {{else}}
                {{$list := cslice `A` `B` `C` `D` `E` `F` `G` `H` `I` `J` `K` `L` `M` `N` `O` `P` `Q` `R` `S` `T` `U` `V` `W` `X` `Y` `Z`}}
                {{$code := ``}}
                {{- range seq 0 4 }}{{ $code = print $code (index (shuffle $list) 1) }}{{ end -}}
                {{$embed.Set "description" (print .User.Mention ", you are about to send " $ci "`" $amt "` to " ($args.Get 0 | userArg).Mention ".\n__If you want to proceed, type the command below.__")}}
                {{$embed.Set "footer" (sdict "text" (print ";give " $code))}}
                {{dbSetExpire .User.ID "giveResp" (sdict "code" $code "given" ($args.Get 0 | toString) "amt" $amt) 15}}
                {{scheduleUniqueCC .CCID nil 15 (print "give " .User.ID) (sdict "user" .User.ID)}}
            {{end}}
            {{sendMessage nil (cembed $embed)}}
        {{end}}
    {{end}}
