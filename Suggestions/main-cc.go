{{/* Trigger Type: Regex; Trigger: \A
	Copyright: Lofty | Brambles, 2021
	License: MIT
	Repository: https://github.com/Lofty-Brambles/some-yagpdb-cc 
	This script uses Regex safe prefixes. It's idea was borrowed from:
	Script: https://bit.ly/3wK6pX0
	Created by Wolveric Catkin (https://github.com/Wolveric), 2021, with assistance from DZ (https://github.com/DZ-TM) */}}
 
{{/*Configurable Values*/}}
{{$anon := false}}
{{$up := "a:Tick:947957018675781662"}}{{$down := "a:Cross:947957187928526868"}}{{$neutral := "neutral:780024034468691968"}}
{{$agenda := "Tavern's"}}{{/*This has to be unique in different suggestions*/}}
{{$modRoles := cslice 953682919568846958}}
{{$log := sdict "approve" "935519890138361896" "implement" "935519890138361896" "deny" "935519890138361896" "delete" "935519890138361896" "dupe" "935519890138361896"}}
{{$cooldown := 0}}
{{$userperm := sdict "edit" true "delete" true}}
{{$mainChannelExists := true}}{{/*If you added a lot of suggestion channel, you want to turn this on, and add the main Channel Handler.*/}}
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
{{$cmd := exec "prefix"}}{{$excess := toString .Guild.ID | len | add 15}}
{{$disprefix := len $cmd | add -1 | slice $cmd $excess}}
{{$prefix := split $disprefix `\E`|joinStr `\E\\E\Q`|printf "\\A\\Q%s\\E"}}
{{if (reFind (print `(?i)\A` $prefix `(approve|implement|deny|comment|dupe|edit|remove|delete)`) (index .CmdArgs 0))}}
{{if lt (len .CmdArgs) 2}}
	{{$er = "Please enter the valid Message ID of the Suggestion!"}}
{{else if $mess := getMessage nil (index .CmdArgs 1)}}
	{{$id := $mess.ID}}
	{{$data := sdict "em" $mess}}
	{{template "s2s" $data}}
	{{$newem := $data.Get "em"}}
	{{$reason := ""}}
	{{if gt (len .CmdArgs) 2}}
		{{range (slice .CmdArgs 2)}}{{- $reason = print $reason . " "}}{{end}}
	{{end}}
	{{$act := (reFind `(approve|implement|deny|comment|dupe|edit|delete|remove)` (index .CmdArgs 0) | lower)}}
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
				{{$emojiList = slice $emojiList 0 (sub (len $emojiList) 3)}}
				{{$newem.Set "Description" (print "__❕ **| Implemented |** ❕__\n" ($newem.Get "Description") "\n\n__Implemented by:__ " .User.Username "\n__Reason:__\n" $reason "\n\n**Final Vote:**\n" $emojiList)}}
				{{$newem.Set "Color" 43127}}{{$newem.Set "Timestamp" currentTime}}
				{{deleteAllMessageReactions nil $id}}
				{{editMessage nil $id (cembed $newem)}}
				{{if ($c := $log.Get "implement")}}{{sendMessage $c (cembed $newem)}}{{end}}
			{{else}}
				{{$er = "Please provide a reason to implement it	!"}}
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
				{{$emojiList = slice $emojiList 0 (sub (len $emojiList) 3)}}
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
			{{else}}
				{{$er = "Please provide a proper comment!"}}
			{{end}}
		{{else}}
			{{$er = "You cannot comment on suggestions!"}}	
		{{end}}
	{{else if eq $act "remove"}}
		{{if $mod}}
			{{if reFind `\n\n\*\*__Comment:__\*\*\n__By: __` ($newem.Get "Description")}}
				{{$new := index (reSplit `\n\n**__Comment:__**\n__By: __` ($newem.Get "Description")) 0}}
				{{$newem.Set "Description" $new}}
				{{editMessage nil $id (cembed $newem)}}
			{{else}}
				{{$er = "You cannot remove a comment as there is none."}}
			{{end}}
		{{else}}
			{{$er = "You cannot remove moderation actions on suggestions!"}}
		{{end}}
	{{else if eq $act "dupe"}}
		{{if $mod}}
			{{$reg := `https://(?:\w+\.)?discord(?:app)?\.com/channels\/(\d+)\/(\d+)\/(\d+)`}}
			{{if ($link := reFind $reg $reason)}}
				{{$newem.Set "Description" (print ($newem.Get "Description") "\n\n**This Suggestion was marked as a Dupe of [This](" $link ")!**\n__Reason:__\n" (reReplace (print $reg `\s*`) $reason ""))}}
				{{deleteMessage nil $id}}
				{{if ($c := $log.Get "dupe")}}{{sendMessage $c (cembed $newem)}}{{end}}
			{{else}}
				{{$er = "Please provide a proper reason to mark this as dupe! The reason must have the link to the Original Suggestion."}}
			{{end}}
		{{else}}
		{{$er = "You cannot mark suggestions as dupe!"}}
		{{end}}
	{{else if and ($userperm.Get "edit") (eq $act "edit")}}
		{{if or $mod (eq .User.ID (reFind `\d{17,19}` $newem.Author.IconURL | toInt))}}
			{{if $reason}}
				{{$newem.Set "Description" (print "⚠ | **Edited**\n" (index (split ($newem.Get "Description") "\n") 0) "\n" $reason)}}
				{{editMessage nil $id (cembed $newem)}}
			{{else}}
				{{$er = "Please provide the entire edited suggestion."}}
			{{end}}
		{{else}}
			{{$er = "You cannot edit someone else's suggestion!"}}
		{{end}}
	{{end}} 
	{{if and ($userperm.Get "delete") (eq $act "delete")}}
		{{if or $mod (eq .User.ID (reFind `\d{17,19}` $newem.Author.IconURL | toInt))}}
							{{$newem.Set "Description" (print "🚫 | **Deleted**\n" ($newem.Get "Description"))}}
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
	{{$sugg := (dbGet 619 .Channel.ID).Value}}
	{{if not $sugg}}
		{{$sugg = $log}}{{$sugg.Set "count" 0}}
	{{end}}
	{{$sugg.Set "count" ($sugg.Get "count" | add 1)}}
	{{$count := $sugg.Get "count"}}
	{{dbSet 619 .Channel.ID $sugg}}
	{{$em := sdict "description" "<a:Loading:952832375954477097> Loading" "color" 0x2e3137}}{{$cont := ""}}
	{{$id := sendMessageRetID nil (cembed $em)}}
	{{addMessageReactions nil $id $up}}
	{{if $neutral}}{{addMessageReactions nil $id $neutral}}{{end}}
	{{addMessageReactions nil $id $down}}
	{{$hid := reFind `\b-anonymous\b` .Message.Content}}
	{{if $hid}}
		{{$cont = reReplace `\b-anonymous\b` .Message.Content ""}}
	{{else}}
		{{$cont = .Message.Content}}
	{{end}}
	{{$em.Set "description" (print "> **" $agenda " | Suggestion #" $count "**\n" $cont)}}
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
			{{$mess = print $mess "\n● You can edit the suggestion by typing in `" $disprefix "edit " $id " <Edited Suggestion>` in the suggestion's channel."}}
		{{end}}
		{{if ($userperm.Get "delete")}}
			{{$mess = print $mess "\n● You can also delete your suggestion by typing in `" $disprefix "delete " $id "` in the suggestion's channel."}}
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
