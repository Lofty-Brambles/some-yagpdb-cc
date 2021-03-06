{{/*Trigger Type: Regex, Trigger: \A(?i);\s*(shop|buy|item)(?: +|\z) */}}
 
{{/*IMPORTANT*/}}
{{/*Add the CCID of your ;give CC here! Both use the same timeout reply!*/}}
{{$gibccid := 17}}
{{/*End of configuration*/}}
 
{{$eco := (dbGet 123 "eco").Value}}{{$inv := (dbGet .User.ID "inv").Value}}
{{$ci := $eco.Get "icon"}}{{$items := $eco.Get "shop"}}
{{if not $inv}}{{$inv := sdict}}{{end}}
 
{{if reFind `(?i)shop` .Cmd}}
    {{if reFind `(?i)stock` .Message.Content}}
        {{if not (in (split (index (split (exec "viewperms") "\n") 2) ", ") "Administrator")}}
            {{sendMessage nil "<a:Cross:947957187928526868> **|** You do not have permissions to restock items in the shop."}}
        {{else}}
            {{$args := parseArgs 3 "Usage: `;shop stock <ItemID> <Stock>`" (carg "string" "") (carg "string" "") (carg "int" "")}}
            {{if not ($items.HasKey (lower (index .CmdArgs 1)))}}
                {{sendMessage nil "<a:Cross:947957187928526868> **|** You need a valid Item ID."}}
            {{else if ( index .CmdArgs 2 | toInt | gt 0 )}}
                {{sendMessage nil "<a:Cross:947957187928526868> **|** You need a valid amount of stock"}}
            {{else}}
                {{$dtls := $items.Get ( index .CmdArgs 1 | lower )}}
                {{$sil := $dtls.Set "stock" ( index .CmdArgs 2 | toInt )}}
                {{$sil := $items.Set ( index .CmdArgs 1 | lower ) $dtls}}
                {{$sil := $eco.Set "shop" $items}}
                {{dbSet 123 "eco" $eco}}
                {{sendMessage nil "<a:Tick:947957018675781662> **|** This item has been restocked!"}}
            {{end}}
        {{end}}
    {{else}}
        {{$plrinv := cslice}}{{$pagination := sdict}}{{$desc := ""}}{{$totp := 1}}
        {{range $k, $v := $items}}
            {{- if $inv.HasKey $k -}}
                {{- $plrinv = $plrinv.Append (print (.Get "icon") " **" (.Get "name") "** [" $ci "`" (.Get "price") "`] — Left in stock: " (.Get "stock") "\n*ID* " $k "\n" (.Get "desc") "\n__You have__: " (($inv.Get $k).Get "quantity")) -}}
            {{- else -}}
                {{- $plrinv = $plrinv.Append (print (.Get "icon") " **" (.Get "name") "** [" $ci "`" (.Get "price") "`] — Left in stock: " (.Get "stock") "\n*ID* " $k "\n" (.Get "desc") "\n__You have__: 0") -}}
            {{- end -}}
        {{- end -}}
        {{if not $items}}
            {{$plrinv = $plrinv.Append (print "There are no items in the shop!")}}
        {{end}}
        {{if gt ($x := len $plrinv) 5}}
            {{$rem := toInt (mod $x 5)}}{{$div := div $x 5}}{{$totp = $div}}{{$pages := cslice}}
            {{if ne $rem 0}}{{$totp := add $totp 1}}{{end}}
            {{range seq 1 $totp}}
                {{- $end := sub . 1|mult 5|add 4 -}}
                {{- if eq . $totp}}{{$end = sub $x 1}}{{end -}}
                {{- $pages := $pages.Append (joinStr "\n\n" (slice $plrinv (sub . 1|mult 5) $end)) -}}
            {{end}}
            {{$desc = index $pages 0}}
            {{$pagination = sdict "pages" $pages "page" 1 "id" 0}}
        {{else}}
            {{range $plrinv}}{{- $desc = joinStr "\n\n" $desc .}}{{end}}
        {{end}}
        {{$id := sendMessageRetID nil (cembed 
            "title" "🏬 | Shop | 🏬"
            "color" 0x2e3137
            "description" $desc
            "footer" (sdict "text" (print "You can buy an item with ;buy <ID> | Page 1 of " $totp)))}}
        {{$silent := $pagination.Set "id" $id}}
        {{if $pagination.HasKey "pages"}}
            {{$temp := (dbGet 123 "pagination").Value}}{{$tempv := $temp.Set "shop" $pagination}}
            {{addReactions "<a:left:949006430730584084>" "<a:right:949008076319637534>"}}
            {{dbSet 123 "pagination" $temp}}
        {{end}}
    {{end}}
{{else if reFind `(?i)b(uy)?` .Cmd}}
    {{if $info := (dbGet .User.ID "buyResp").Value}}
        {{if reFind (print `(?i)` ($info.Get "code")) .Message.Content}}
            {{$item := $items.Get ($info.Get "key")}}{{$price := $item.Get "price"}}{{$amt := $info.Get "amt"}}
            {{if $inv}}
                {{if $inv.HasKey ($info.Get "key")}}
                    {{$store := $inv.Get ($info.Get "key")}}{{$sil := $store.Set "quantity" ($store.Get "quantity" | add 1)}}{{$sil := $inv.Set ($info.Get "key") $store}}
                {{else}}
                    {{$itemInv := sdict "name" ($item.Get "name") "icon" ($item.Get "icon") "quantity" $amt}}{{$sil := $inv.Set ($info.Get "key") $itemInv}}
                {{end}}
            {{else}}
                {{$inv = sdict ($info.Get "key") (sdict "name" ($item.Get "name") "icon" ($item.Get "icon") "quantity" $amt)}}
            {{end}}
            {{dbSet .User.ID "inv" $inv}}
            {{$sil := $item.Set "stock" (sub ($item.Get "stock") $amt)}}
            {{$sil := $items.Set ($info.Get "key") $item}}
            {{$sil := $eco.Set "shop" $items}}
            {{dbSet 123 "eco" $eco}}
            {{dbSet .User.ID "bal" (sub (dbGet .User.ID "bal").Value (mult $amt $price))}}
            {{dbDel .User.ID "buyResp"}}
            {{cancelScheduledUniqueCC $gibccid (print "buy " .User.ID)}}
            {{sendMessage nil "<a:Tick:947957018675781662> **|** Your purchase was successful!"}}
        {{end}}
    {{else}}
        {{$args := parseArgs 1 "Usage: `;buy <ItemID> <Optional: Amount>`" (carg "string" "") (carg "int" "")}}
        {{$amt := 1}}
        {{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256")) "color" 0x2e3137}}
        {{if ($args.IsSet 1)}}
            {{if (gt ($args.Get 1) 0)}}{{$amt = $args.Get 1}}{{end}}
        {{end}}
        {{if ($items.HasKey ($x := lower (index .CmdArgs 0)))}}
            {{$item := $items.Get $x}}
            {{if le ($item.Get "price" | mult $amt | toFloat) (dbGet .User.ID "bal").Value}}
                {{if ge ($item.Get "stock") $amt}}
                    {{$list := cslice `A` `B` `C` `D` `E` `F` `G` `H` `I` `J` `K` `L` `M` `N` `O` `P` `Q` `R` `S` `T` `U` `V` `W` `X` `Y` `Z`}}
                    {{$code := ``}}
                    {{- range seq 0 4 }}{{ $code = print $code (index (shuffle $list) 1) }}{{ end -}}
                    {{$embed.Set "description" (print .User.Mention ", you are about to buy " $amt " " ($item.Get "icon") " **" ($item.Get "name") "**.\nTo confirm this purchase, please type:")}}
                    {{$embed.Set "footer" (sdict "text" (print ";buy " $code))}}
                    {{dbSetExpire .User.ID "buyResp" (sdict "code" $code "amt" $amt "key" $x) 15}}
                    {{scheduleUniqueCC $gibccid nil 15 (print "buy " .User.ID) (sdict "user" .User.ID)}}
                {{else}}
                    {{$embed.Set "description" "<a:Cross:947957187928526868> **|** There is not enough stock left for this item."}}
                {{end}}
            {{else}}
                {{$embed.Set "description" "<a:Cross:947957187928526868> **|** You do not have enough balance to buy this."}}
            {{end}}
        {{else}}
            {{$embed.Set "description" "<a:Cross:947957187928526868> **|** You need a valid Item ID."}}
        {{end}}
        {{sendMessage nil (cembed $embed)}}
    {{end}}
{{else if reFind `(?i)item` .Cmd}}
    {{$args := parseArgs 0 "Usage: `;item <ItemID>`" (carg "string" "")}}
    {{if ($items.HasKey ($x := ($args.Get 0 | lower)))}}
        {{$item := $items.Get $x}}
        {{sendMessage nil (cembed "title" ($item.Get "name") "description" (print "> " ($item.Get "desc") "\n**Item ID**: " $x "\n**Price**: " $ci "`" ($item.Get "price"|humanizeThousands ) "`\n**Stock**: " ($item.Get "stock")) "color" 0x2e3137 "footer" (sdict "text" "This was searched" "icon_url" "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQHNJUBpKMFTo1GJdp9m0k-6fCKhgJ7g1Wjcw&usqp=CAU") "timestamp" currentTime)}}
    {{else}}
        {{sendMessage nil "<a:Cross:947957187928526868> **|** You need a valid Item ID."}}
    {{end}}
{{end}}
