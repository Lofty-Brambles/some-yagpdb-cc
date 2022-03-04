{{/*Use this command to prepare before starting to use the price guide or my crappy code will create more issues than a 3 y/o*/}}
{{/*1. Pass this command to prepare*/}}
{{/*2. Add the rest of the cc's*/}}
{{/*3. Add the config values in the Main command and !!addmod command.*/}}
{{/*4. Remember to assign the manager role to people who will update prices, with the !!addmod command and not manually give them the role. Admins, you can do you and add the role manually.*/}}
{{/*If this runs and returns "Done!", proceed with steps 2,3,4.*/}}

{{/*Main command needs RoleID's of admin and manager, true/false if a second currency exists, Conversion rate (1 <Second_currency> = $denom <main_currency>), Name of second currency, an emoji as an icon of the second currency, and a cooldown on price updates (in seconds), as the configuration values.*/}}

{{/*Properties:*}}
{{/*1. This command - Trigger: prime, Trigger Type: command*/}}
{{/*2. "!!addmod" command - Trigger: !!addmod, Trigger Type: Starts with*/}}
{{/*3. Main command (The big one) - Trigger: !!, Trigger Type: Starts with*/}}
{{/*4. Reaction handler - Trigger Type: Reactions Added*/}}

{{dbSet 0 "id-slice" (cslice)}}
{{dbSet 1 "lastpos" 1}}

{{/* Make sure, that the UserID's of 420 and 421 have a all key's empty in the database*/}}
{{/*It also uses 1 key off the UserID of the user, so make sure to allocate that.*/}}

{{sendMessage nil "Done."}}

{{/*Oh and delete this once you are done so you don't accidentally messup your price guide*/}}
