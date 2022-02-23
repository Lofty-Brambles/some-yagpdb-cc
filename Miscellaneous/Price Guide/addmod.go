{{/*Use this to add the "Manager/Moderator roles"*/}}
{{/*Admins can be added by just giving them the role manually*/}}
{{/*Config Values*/}}
 
{{$m := 944551072452730951}}{{/*Manager Role ID*/}}
 
{{/*Bad code below*/}}
{{$a := parseArgs 1 "Usage `!!addmod @Member`" (carg "userid" "user")}}
{{dbSet ($a.Get 0) "update" (sdict)}}
{{giveRoleID ($a.Get 0) $m}}
