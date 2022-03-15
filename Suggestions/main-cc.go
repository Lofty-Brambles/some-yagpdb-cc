{{/*Trigger Type: Regex; Trigger: \A*/}}
{{/*Configurable Values*/}} 
{{$anon := true}}
{{$neutral := true}}
{{$up := "a:Tick:947957018675781662"}}{{$down := "a:Cross:947957187928526868"}}{{$neutral := "neutral:780024034468691968"}}
{{$agenda := "Something"}}
{{$modRoles := cslice 123 456}}
{{$log := sdict "approve" "935519890138361896" "implement" "935519890138361896" "deny" "935519890138361896" "delete" "935519890138361896"}}
{{$cooldown := 10}}
{{$userperm := sdict "edit" false "delete" false}}
{{$prefix := `\-`}}{{/*Must be escaped*/}}
{{/*End of Configurable Values*/}}
 
{{deleteTrigger 0}}
{{$mod := false}}
{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") "Administrator")}}
	{{$mod = true}}
{{else}}
	{{range $modRoles}}
		{{- if (hasRoleID .)}}}
		{{- $mod = true}}
		{{- end}}
	{{- end}}
{{end}}
{{$er := ""}}
{{if (reFind (print `(?i)\A` $prefix `(approve|implement|deny|comment|dupe|edit|delete)`) (index .CmdArgs 0))}}
	{{if $mess := getMessage nil (index .CmdArgs 1)}}
		{{$id := $mess.ID}}
		{{$data := sdict "em" $mess}}
		{{template "s2s" $data}}
		{{$newem := $data.Get "em"}}
		{{$reason := ""}}
		{{if gt (len .CmdArgs) 2}}
			{{range (slice .CmdArgs 2)}}{{- $reason = print $reason .}}{{end}}
		{{end}}
		{{$act := (reFind `(approve|implement|deny|comment|dupe|edit|delete)` (index .CmdArgs 0) | lower)}}
		{{if eq $act "approve"}}
			{{if $mod}}
				{{if $reason}}
					{{$newem.Set "Description" (print "__✅ **| Approved |** ✅__\n" ($newem.Get "Description") "\n\n__**Approved by:**__ " .User.Username "\n__Reason:__\n" $reason)}}
					{{$newem.Set "Color" 43127}}{{$newem.Set "Timestamp" currentTime}}
					{{editMessage nil $id (cembed $newem)}}
					{{if ($c := $log.Get "approve")}}{{sendMessage $c (cembed $newem)}}{{end}}
				{{else}}
					{{$er = "Please provide an reason approve it!"}}
				{{end}}
			{{else}}
				{{$er = "You cannot accept suggestions!"}}
			{{end}}
		{{else if eq $act "implement"}}
			{{if $mod}}	
				{{if $reason}}
					{{$emojiList := ""}}
					{{range $mess.Reactions}}
						{{- $emojiList = print $emojiList .Emoji.MessageFormat .Count " | "}}
					{{end}}
					{{$emojiList := slice $emojiList 0 (sub (len $emojiList) 3)}}
					{{$newem.Set "Description" (print "__❕ **| Implemented |** ❕__\n" ($newem.Get "Description") "\n\n__Implemented by:__ " .User.Username "\n__Reason:__\n" $reason "\n\n**Final Vote:**\n" $emojiList)}}
					{{$newem.Set "Color" 43127}}{{$newem.Set "Timestamp" currentTime}}
					{{deleteAllMessageReactions nil $id}}
					{{editMessage nil $id (cembed $newem)}}
					{{if ($c := $log.Get "implement")}}{{sendMessage $c (cembed $newem)}}{{end}}
                {{else}}
                    {{$er = "Please provide a reason to implement it!"}}
                {{end}}
			{{else}}
				{{$er = "You cannot implement suggestions!"}}
			{{end}}
		{{else if eq $act "deny"}}
			{{if $mod}}
				{{if $reason}}
					{{$emojiList := ""}}
					{{range $mess.Reactions}}
						{{- $emojiList = print $emojiList .Emoji.MessageFormat .Count " | "}}
					{{end}}
					{{$emojiList := slice $emojiList 0 (sub (len $emojiList) 3)}}
					{{$newem.Set "Description" (print "__❌ **| Denied |** ❌__\n" ($newem.Get "Description") "\n\n__Denied by:__ " .User.Username "\n__Reason:__\n" $reason "\n\n**Final Vote:**\n" $emojiList)}}
					{{$newem.Set "Color" 15605837}}{{$newem.Set "Timestamp" currentTime}}
					{{deleteAllMessageReactions nil $id}}
					{{editMessage nil $id (cembed $newem)}}
					{{if ($c := $log.Get "deny")}}{{sendMessage $c (cembed $newem)}}{{end}}
                {{else}}
					{{$er = "Please provide a reason to deny it!"}}
                {{end}}
			{{else}}
				{{$er = "You cannot deny suggestions!"}}
			{{end}}
		{{else if eq $act "comment"}}
			{{if $mod}}
				{{if $reason}}
					{{$newem.Set "Description" (print ($newem.Get "Description") "\n\n**__Comment:__**\n__By: __" .User.Username "\n" $reason)}}
					{{editMessage nil $id (cembed $newem)}}
				{{else if not $reason}}
					{{$er = "Please provide a proper comment!"}}
				{{end}}
			{{else}}
				{{$er = "You cannot comment on suggestions!"}}	
			{{end}}
        {{else if eq $act "dupe"}}
			{{if $mod}}
				{{if $reason}}
					{{$newem.Set "Description" (print ($newem.Get "Description") "\n\n**This Suggestion was marked as a Dupe!**\n__Reason:__\n" $reason)}}
					{{sendDM (complexMessage "content" (print "") "embed" (cembed $newem))}}
					{{deleteMessage nil $id}}
				{{else}}
					{{$er = "Please provide a proper reason to mark this as dupe!"}}
				{{end}}
			{{else}}
			{{$er = "You cannot mark suggestions as dupe!"}}
            {{end}}
		{{else if eq $act "edit"}}
			{{if or $mod (eq .User.ID (reFind `\d{17,19}` $newem.Author.IconURL | toInt))}}
				{{if $reason}}
					{{$newem.Set "Description" (print (index (split ($newem.Get "Description") "\n") 0) "\n" $reason)}}
					{{editMessage nil $id (cembed $newem)}}
				{{else}}
					{{$er = "Please provide the entire edited suggestion."}}
				{{end}}
			{{else}}
				{{$er = "You cannot edit someone else's suggestion!"}}
			{{end}}
		{{end}}
		{{if eq $act "delete"}}
			{{if or $mod (eq .User.ID (reFind `\d{17,19}` $newem.Author.IconURL | toInt))}}
				{{if ($c := $log.Get "delete")}}{{sendMessage $c (cembed $newem)}}{{end}}
				{{deleteMessage nil $id}}
			{{else}}
				{{$er = "You cannot delete someone else's suggestion!"}}
			{{end}}
		{{end}}
    {{else}}
        {{$er = "Please enter the valid Message ID of the Suggestion!"}}
    {{end}}
    {{if $er}}
        {{$id := sendMessageRetID nil (print .User.Mention " **—** " $er)}}
	{{deleteMessage nil $id 10}}
    {{end}}
{{else}}
	{{$precd := 1}}
	{{if ($gcd := toInt (dbGet .User.ID "suggestion").Value)}}{{$precd = $gcd}}{{end}}
	{{if or (not $cooldown) (ge currentTime.Unix $precd)}}
		{{$em := sdict "description" "<a:Loading:952832375954477097> Loading" "color" 0x2e3137}}{{$cont := ""}}
		{{$id := sendMessageRetID nil (cembed $em)}}
		{{addMessageReactions nil $id $up}}
		{{if $neutral}}{{addMessageReactions nil $id $neutral}}{{end}}
		{{addMessageReactions nil $id $down}}
		{{$hid := reFind `-a(?:nonymous)` .Message.Content}}
		{{if $hid}}
			{{$cont = reReplace `-a(?:nonymous)?` .Message.Content ""}}
		{{else}}
			{{$cont = .Message.Content}}
		{{end}}
		{{$em.Set "description" (print "> **" $agenda " | Suggestion #" (dbIncr 619 "suggest-count" 1) "**\n" $cont)}}
		{{if and $anon $hid}}
			{{$em.Set "author" (sdict "name" "Anonymous" "icon_url" "https://static.thenounproject.com/png/302770-200.png")}}
		{{else}}
			{{$em.Set "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "256"))}}
		{{end}}
		{{if .Message.Attachments}}
			{{$em.Set "image" (sdict "url" (index .Message.Attachments 0).ProxyURL)}}
		{{end}}
		{{$em.Set "timestamp" currentTime}}
		{{editMessage nil $id (cembed $em)}}
		{{if or ($userperm.Get "edit") ($userperm.Get "delete")}}
			{{$mess := print "__Your Suggestion from **" .Guild.Name "** was noted!__"}}
			{{if ($userperm.Get "edit")}}
				{{$mess = print $mess "\n● You can edit the suggestion by typing in `edit " $id " <Edited Suggestion>` in the suggestion's channel."}}
			{{end}}
			{{if ($userperm.Get "delete")}}
				{{$mess = print $mess "\n● You can also delete your suggestion by typing in `delete " $id "` in the suggestion's channel."}}
			{{end}}
			{{sendDM (cembed "color" 0x2e3137 "description" $mess)}}
		{{end}}
		{{if $cooldown}}
			{{dbSet .User.ID "suggestion" (add currentTime.Unix $cooldown)}}
		{{end}}
	{{else}}
			{{$id := sendMessageRetID nil (print .User.Mention " **—** You have to wait before you can suggest again. You can suggest again <t:" $precd ":R>")}}
			{{deleteMessage nil $id 10}}
	{{end}}
{{end}}
{{define "s2s"}}
	{{$embin := structToSdict (index .em.Embeds 0)}}
	{{range $k, $v := $embin}}
		{{- if eq (kindOf $v true) "struct"}}
			{{- $embin.Set $k (structToSdict $v)}}
		{{- end}}
	{{- end}}
	{{if $embin.Author}}{{$embin.Author.Set "Icon_URL" $embin.Author.IconURL}}{{end}}
	{{if $embin.Footer}}{{$embin.Footer.Set "Icon_URL" $embin.Footer.IconURL}}{{end}}
	{{.Set "em" $embin}}
{{end}}
