{{/*Trigger Type: Regex, Trigger: \A(?i);\s*b(?:ank)?(?: +|\z) */}}
 
{{$args := parseArgs 0 "Usage : `;bank <Optional: withdraw/deposit> <Optional: amount>" (carg "string" "") (carg "string" "")}}
 
{{$ci := (dbGet 123 "eco").Value.Get "icon"}}
{{$dash := (dbGet .User.ID "dash").Value}}
{{- if not $dash }}{{ $dash = sdict "bbal" 1 "bquota" 20000 "brate" 0.02}}{{ end -}}
{{$interest := $dash.Get "brate"}}
 
{{$embed := sdict 
    "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256"))
    "title" "The Gold Jar"
    "footer" (sdict "text" "Remember to work (;work) everyday, or the bank won't pay you! | Also check your account balance to prevent overfilling!")
    "thumbnail" (sdict "url" "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTG_d_Svyu2EAARNGTqwLlKfyilN4DJwwng1A&usqp=CAU")
    "color" 0x2e3137}}
 
{{$warning := print "<a:WarningNeon:947360557512654848> **" .User.Mention ", you have an alert** <a:WarningNeon:947360557512654848>\n\n> Your bank account is almost full. Please purchase a Locker Pass or withdraw in order to not lose money the next time you withdraw."}}
 
{{if le ($dash.Get "bbal") (div 1 (add 1 $interest)|mult ($dash.Get "bquota")|toInt)}}
    {{$warning = ""}}
{{end}}
 
{{if not .CmdArgs}}
    {{$embed.Set "description" (print "> Your current Bank balance is:" $ci "`" (humanizeThousands ($dash.Get "bbal")) "`\n> Your current Bank Account Limit is: " $ci "`" (humanizeThousands ($dash.Get "bquota")) "`")}}
    {{sendMessage nil (complexMessage "content" $warning "embed" (cembed $embed))}}
{{else}}
    {{if reFind `(?i)w(ith)?(draw)?` (index .CmdArgs 0)}}
        {{if $args.IsSet 1}}
            {{if $amt := reFind `\d+` ($args.Get 1)|toInt}}
                {{if gt $amt ($dash.Get "bbal")}}
                    {{$embed.Set "description" (print "You only have " $ci "`" (humanizeThousands ($dash.Get "bbal")) "` in your bank account. You can't transfer more than that.")}}
                {{else}}
                    {{$embed.Set "description" (print "An amount of " $ci "`" $amt "` was transferred from your bank account to your pocket balance.\n> **New Bank Balance**: " $ci "`" (humanizeThousands (sub ($dash.Get "bbal") $amt)) "` / `" (humanizeThousands ($dash.Get "bquota")) "`")}}
                    {{$sil := $dash.Set "bbal" (sub ($dash.Get "bbal") $amt)}}
                    {{dbSet .User.ID "bal" (add (dbGet .User.ID "bal").Value $amt)}}
                    {{dbSet .User.ID "dash" $dash}}
                {{end}}
            {{else if reFind `(?i)all` ($args.Get 1)}}
                {{$embed.Set "description" (print "An amount of " $ci "`" ($dash.Get "bbal") "` was transferred from your bank account to your pocket balance.\n> **New Bank Balance**: " $ci "`0` / `" (humanizeThousands ($dash.Get "bquota")) "`")}}
                {{dbSet .User.ID "bal" (add (dbGet .User.ID "bal").Value ($dash.Get "bbal"))}}
                {{$sil := $dash.Set "bbal" 0}}
                {{dbSet .User.ID "dash" $dash}}
            {{else}}
                {{$embed.Set "description" "Please provide a valid amount to withdraw."}}
            {{end}}
            {{sendMessage nil (complexMessage "content" $warning "embed" (cembed $embed))}}
        {{else}}
            {{$embed.Set "description" "Please provide a valid amount to withdraw."}}
            {{sendMessage nil (complexMessage "content" $warning "embed" (cembed $embed))}}
        {{end}}
    {{else if reFind `(?i)d(ep)?(osit)?` (index .CmdArgs 0)}}
        {{if $args.IsSet 1}}
            {{$bal := (dbGet .User.ID "bal").Value|toInt}}
            {{if $amt := reFind `\d+` ($args.Get 1)|toInt}}
                {{if gt $amt $bal}}
                    {{$embed.Set "description" (print "You only have " $ci "`" (humanizeThousands $bal) "` in your pocket. You can't transfer more than that.")}}
                {{else if gt (add $amt ($dash.Get "bbal")) ($dash.Get "bquota")}}
                    {{$embed.Set "description" (print "You only have an account limit of " $ci "`" (humanizeThousands ($dash.Get "bquota")) "`. You cannot transfer more than that.")}}
                {{else}}
                    {{$embed.Set "description" (print "An amount of " $ci "`" $amt "` was transferred from your pocket balance to your bank account.\n> **New Bank Balance**: " $ci "`" (humanizeThousands (add $amt ($dash.Get "bbal"))) "` / `" (humanizeThousands ($dash.Get "bquota")) "`")}}
                    {{$sil := $dash.Set "bbal" (add ($dash.Get "bbal") $amt)}}
                    {{dbSet .User.ID "bal" (sub $bal $amt)}}
                    {{dbSet .User.ID "dash" $dash}}
                {{end}}
            {{else if reFind `(?i)all` ($args.Get 1)}}
                {{$amt := $bal}}
                {{if le (add $bal ($dash.Get "bbal")) ($dash.Get "bquota")}}
                    {{$amt = toInt $bal}}
                    {{$embed.Set "description" (print "An amount of " $ci "`" $amt "` was transferred from your pocket balance to your bank account.\n> **New Bank Balance**: " $ci "`" (add $amt ($dash.Get "bbal")) "` / `" (humanizeThousands ($dash.Get "bquota")) "`")}}
                    {{dbSet .User.ID "bal" 0}}
                    {{$sil := $dash.Set "bbal" (add $amt ($dash.Get "bbal"))}}
                    {{dbSet .User.ID "dash" $dash}}
                {{else}}
                    {{$embed.Set "description" "Your pocket balance exceeds your bank quota. Please buy some Locker Passes or enter a valid amount to keep in your bank."}}
                {{end}}
            {{else}}
                {{$embed.Set "description" "Please provide a valid amount to deposit."}}
            {{end}}
            {{sendMessage nil (complexMessage "content" $warning "embed" (cembed $embed))}}
        {{else}}
            {{$embed.Set "description" "Please provide a valid amount to deposit."}}
            {{sendMessage nil (complexMessage "content" $warning "embed" (cembed $embed))}}
        {{end}}
    {{end}}
{{end}}
