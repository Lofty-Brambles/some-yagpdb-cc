{{/* Trigger Type: Regex, Trigger: \A(?i)\s*;c(?:on)?fi?gadd(?: +|\z) */}}
{{/* Adds the replies to the ;con and ;cave commands */}}
{{$conGoodReplies := cslice
 "You managed to find someone who just finished a cave expedition and rob him at gunpoint."
 "You stole some fertile eggs and sold them at the black market."
 "You visited the admin's base while they were offline and yoinked some decent blueprints. You sell them for a lot."
 "Someone left some unattended metal and you took it as before they noticed."
 "Your tribe was stocking up on berries and brews, but you sell some off secretly."}}
{{$conBadReplies := cslice
 "You were caught trying to sneak some polymer out from the tribe storage."
 "On the way to your base's storage, a flying Titanoboa knocks you out and kills you, and your stuff despawns."
 "You were trading illegal goods, and the buyer left without paying."
 "You ate some random crops you stole and got food-poisoning. You had to go to the hospital."
 "You ran into a Giganotosaurus while trying to get a few eggs, on foot. You were eaten alive."}}
{{$caveGoodReplies := cslice
 "You were running the Lower-South Cave and found a dead player's loot crate."
 "While finding the Artifact of the Strong, you found a friendly Yeti who lent you a few bucks. Weird, huh?"
 "Deep in the Lava cave ruins, you found ancient archaeological pieces. You sold them to the scholars."
 "You found an underwater trench and a stuck sea monster. You avoid it and obtain some rare black pearls."
 "You explore the poisonous depths of the Swamp Cave and find some rare ancient loot."}}
{{$caveBadReplies := cslice
 "You got jump-scared by 5 hidden arthropleuras and lost all your equipment."
 "You decide to explore the North-Eastern Cave without training, and broke your ankle."
 "Enjoying a successful and profitable cave run, you got too drunk. Someone managed to take all you had while you were passed out."
 "You lost an expensive ascendant hatchet. Must have been the lag. Pain."
 "You were exploring underwater and got ambushed by Cnidarians. You lie helpless while your equipment floats away."}}
{{dbSet 123 "conpassreply" $conGoodReplies}}
{{dbSet 123 "confailreply" $conBadReplies}}
{{dbSet 123 "cavepassreply" $caveGoodReplies}}
{{dbSet 123 "cavefailreply" $caveBadReplies}}
{{$shop := sdict
"clover" (sdict "name" "Stonks Clover" "icon" "<a:rainbowclover:950383022069411860>" "price" 3000 "stock" 100 "desc" "The typical charm. Using one has a 50% chance to grant a huge, temporary multiplier on any one of the various income actions.")
"hammer" (sdict "name" "Thonk Hammer" "icon" "<:thonkhammer:950385819175223366>" "price" 20000 "stock" 20 "desc" "Has a random chance to crack up the return on the bank for a fortnight, to 10x the original.")
"locker" (sdict "name" "Locker Pass" "icon" "????" "price" 2500 "stock" 50 "desc" "Grants a random amount of locker space for your bank account.")
"map" (sdict "name" "Quester Map" "icon" "<a:RibbonWhite:950462581846474753>" "price" 2000 "stock" 20 "desc" "Unlocks a mystery quest. Can be completed to get rewards!")}}
{{$temp := (dbGet 123 "eco").Value}}
{{if not $temp}}
    {{$temp = sdict}}
{{end}}
{{$sil := $temp.Set "shop" $shop}}
{{$c := $temp.Set "icon" "<:Giant_Bee_Honey:947950040905822268>"}}
{{dbSet 123 "eco" $temp}}
{{sendMessage nil "Done."}}
