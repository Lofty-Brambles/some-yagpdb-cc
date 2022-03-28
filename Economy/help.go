{{/*Trigger Type: Regex; Trigger: \A(?i);\s*h(?:elp)?(?: +|\z) */}}
 
{{$genhelp := "> **Welcome Down-under!**\nIn our server, we have an integrated Ark-Discord Economy. To know more about the specifics, check out <#>.\n**For more specific information on a section, type `;help <Section Name>`!**\n> __Sections__\n<a:FLAMEviolet:952285076123701268> | How to Earn\n<a:FLAMEblue:952285256466190426> | Banking\n<a:FLAMEcherry:952284880132276264> | Shops\n<a:FLAMEcray:952285399382888478> | Items\n<a:FLAMEgreen:952284957793988730> | Others\n> __Note__```The HONEY JARS used as In-Discord Currency ARE NOT the IN-GAME HONEY JARS. They cannot be interchanged or traded.```"}}
{{$earn := "> **How to Earn**\n<a:FLAMEviolet:952285076123701268> | `;work` — High pay, needs you to do a task, and pays even on fails. But, has a huge cooldown,\n<a:FLAMEviolet:952285076123701268> | `;con` — Has a 60% chance to pay, with a cooldown of 15 minutes, but you can lose money too,\n<a:FLAMEviolet:952285076123701268> | `;cave` — Has a 75% chance to pay, with a cooldown of 10 minutes, but, again, you can lose money,\n> These are the main income sources."}}
{{$bank := "> **Banks**\n<a:FLAMEblue:952285256466190426> | `;bank` — Check how much money is in your bank, and how much it can store,\n<a:FLAMEblue:952285256466190426> | `;bank deposit X` — Deposits X amount of money back into your bank,\n<a:FLAMEblue:952285256466190426> | `;bank withdraw X` — Withdraws X amount of money from your bank account,\n> Remember, banks are a grest passive income source, cos money kept in the bank __increments 2% everyday__. Might not sound much at first, but it is compounded interest."}}
{{$shops := "> **Shop**\n<a:FLAMEcherry:952284880132276264> | `;shop` — Shows the server's item shop,\n<a:FLAMEcherry:952284880132276264> | `;buy <item ID> X` — Buys an X amount of an item. If X isn't mentioned, it defaults to 1,\n> This is __only the shop for custom items__. For player trades, there are other shops, as well as admin shops for materials."}}
{{$item := "> **Items**\n<a:FLAMEcray:952285399382888478> | `;item <item ID>` — Displays the description of any item. Item ID's can be found from the shop,\n<a:FLAMEcray:952285399382888478> | `;use <item ID>` — Uses an item, which can be found/bought for,\n> These refer to the __custom items__ in the economy."}}
{{$other := "> **Others**\n<a:FLAMEgreen:952284957793988730> | `;give <@User/UserID> X` — Used to give an user, a certain amount of money, X,\n<a:FLAMEgreen:952284957793988730> | `;balance <@User/UserID>` — Displays the balance of an user, if user is not mentioned, shows the balance of the person who ran the command,\n<a:FLAMEgreen:952284957793988730> | `;inventory <@User/UserID>` — Displays the inventory of an user, if user is not mentioned, shows the inventory of the person who ran the command,\n<a:FLAMEgreen:952284957793988730> | `;leaderboard` — Displays the standings of the people, based on their pocket money,\n> These are the utility commands to the economy."}}
 
{{$embed := sdict
    "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256"))
    "color" 0x2e3137
    "footer" (sdict "text" "Use ;help to know more!")
    "timestamp" currentTime
    "title" " 🛰 | Help | 🛰 "
    "thumbnail" (sdict "url" "")}}
 
{{$args := parseArgs 0 "Usage: `;help <Optional: Section>`" (carg "string" "")}}
{{if ($args.IsSet 0)}}
 
    {{if ($args.Get 0 | reFind `(?i)earn\z`)}}
        {{$embed.Set "description" $earn}}{{$embed.Set "color" 8323327}}
    {{else if ($args.Get 0 | reFind `(?i)banks?(ing)?\z`)}}
        {{$embed.Set "description" $bank}}{{$embed.Set "color" 255}}
    {{else if ($args.Get 0 | reFind `(?i)shops?\z`)}}
        {{$embed.Set "description" $shops}}{{$embed.Set "color" 16758725}}
    {{else if ($args.Get 0 | reFind `(?i)items?\z`)}}
        {{$embed.Set "description" $item}}{{$embed.Set "color" 16747520}}
    {{else if ($args.Get 0 | reFind `(?i)others?\z`)}}
        {{$embed.Set "description" $other}}{{$embed.Set "color" 43127}}
    {{else}}
        {{$embed.Set "description" "Please mention a proper section to view it's commands!"}}
    {{end}}
    
{{else}}
    {{$embed.Set "description" $genhelp}}
{{end}}
 
{{sendMessage nil (cembed $embed)}}
