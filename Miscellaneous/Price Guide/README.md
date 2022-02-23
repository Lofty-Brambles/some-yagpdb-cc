# An Item Price Guide
This is a basic 'fetch price' 'Item to price' guide that supports a singular currency, as well as a second currency.
## Features
- Stores Item-Price values in a condensed database.
- Has the ability to add new items with an initial price and a picture.
- Has a search to find items and return their ID.
- Displays Prices by Item ID's, with 1-2 different prices, rounded to two decimals.
- Allows Managers to update prices
  - Prices that get displayed are the average of the last three updates.
  - On updating a particular item's price once, it sets a cooldown for when it can be updated by that manager again.
- Allows Managers to mark items as "Outdated" and remove that "Outdated" status.
  - While an item is "Outdated", it will not display any prices. 
  - However, it's prices can still be updated by managers.
- Has the ability to remove the last item added to the guide, in case of errors while adding.
- Has the ability to remove the last price added to an item and resolve the previous price, in case of any errors while updating.
- Has seperate Help Menus for Members, Managers and Admins.
## Setup
- Add the [Prime](prime.go) file. It has all the triggers and trigger types for the custom commands in it.
- Run `-prime` once and delete the custom command.
- Add the other three files and configure them.
  - Configuration in [addmod command](addmod.go): Add the Role ID of the manager/moderator role. They can update/outdate items.
  - Configuration in [Main command](main-command.go): 
    - Add the Role ID's of Manager and Admins.
    - Keep `$scur` as `true` if you want to have prices in a second currency as well. Else change it to `false` and skip the next point. [__1 Second Currency = X Main currency__]
    - If you have a second currency, add the value of X in `$denom`, The name of the second currency in `$c2name` and an emoji relating to the second currency in `$emoji`.
    - Add the cooldown to update prices for each manager in seconds in `$cd`. Default provided is 6 hours = 21600 seconds.
  - The [Reaction Handler](reaction-handler.go) has no configuration necessity.
- You can now start adding items and using the price guide.
## Gallery
- __The `!!find all` command__

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985109377556540/IMG_2301.png)
- __The `!!find <item>` command__

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985210191855626/IMG_2302.png)
- __The `!!price <id>` command__

![image](https://cdn.discordapp.com/attachments/899657506173882461/945986401218985994/IMG_2304.png)
- __The `!!update <id> <price>` command__

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985873646866442/IMG_2307.png)
- __The `!!outdate <id> <true/false>` command__

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985884803727471/IMG_2308.png)
- __The `!!add <cost> <picture_url> <Name>` command__

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985891648802816/IMG_2310.png)
- __The `!!remprice <id>` command__

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985896610680852/IMG_2311.png)
- __The `!!help` commands__ (for different roles)

 For everyone:

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985913723437086/IMG_2314.png)
 For Manager/Moderators:

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985907482308608/IMG_2313.png)
 For Admins:

![image](https://cdn.discordapp.com/attachments/899657506173882461/945985900834344980/IMG_2312.png)
