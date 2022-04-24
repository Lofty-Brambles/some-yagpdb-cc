{{/* Config starts */}}

{{ $emojis := cslice "🚒"}}{{/* If you want to add emojis, add them. If not, keep empty. */}}
{{ $max := 1 }}{{/* Max attachments for default user */}}
{{ $limits := sdict "123" "1" "123" "3" "947954288909488190" "bypass" }}{{/* Add Role ID - limit pairs. "bypass" will remove limit.*/}}

{{/* Config ends */}}

{{/* Setting variables */}}
{{ $count := 0 }}{{ $messageObj := getMessage nil .Message.ID }}
{{ $del := false }}{{ $role := "0" }}

{{/* Fetching attachments */}}
{{ if $embs := $messageObj.Embeds }}
    {{ range $embs }}
       {{ if .Thumbnail }}{{ $count = add $count 1 }}{{ end }}
       {{ if .Image }}{{ $count = add $count 1 }}{{ end }}
    {{ end }}
{{ else if $att := $messageObj.Attachments }}
    {{ $count = len $att | add $count }}
{{ end }}

{{/* Comparing limits */}}
{{ range $k, $v := $limits }}
    {{ if and ( toInt $k | hasRoleID ) ( gt ( toInt $k ) $max ) }}
        {{ $temp := $v }}
        {{ if eq $temp "bypass" }}
            {{ $max = 999999999 }}
        {{ else }}
            {{ $max = toInt $temp }}
        {{ end }}
        {{ $role = $k }}
    {{ end }}
{{ end }}
{{ if gt $count $max }}
    {{ $del = true }}
{{ end }}

{{/* Sending out de outputs */}}
{{ $rolemessage := "" }}
{{ if ne $role "0" }}
    {{ $rolemessage = print ", because of you having the <@" $role "> role" }}
{{ end }}
{{ if eq $del true }}
    {{ deleteMessage nil .Message.ID 0 }}
    {{ $id := sendMessageRetID nil ( complexMessage "content" .User.Mention "embed" ( cembed 
        "title" "Warning"
        "description" ( print "> <:scross:963688138155388978> | Your message was __deleted!__\nYou have been allowed to post a maximum of " $max " attachments" $rolemessage "!" )
        "color" 0x2e3137
        "timestamp" currentTime ) ) }}
    {{ deleteMessage nil $id 10 }}
{{ else }}
    {{ range $emojis }}{{ addReactions . }}{{ end }}
{{ end }}
