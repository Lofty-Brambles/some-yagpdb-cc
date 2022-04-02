# A Hassle-less Multi-Suggestion System
I've rarely seen a multi-channel suggestion system where you can have more than one parallel suggestion channels for more than one branch of a server.
This CC aims to cover those issues.
## Further Plans
- [ ] Add a Main Channel CC to fetch and manage suggestions across all channels.
- [ ] Add a Reaction Handler to manage highly upvoted suggstions and so on.
## Database Usage
- `One Universal Key` [User-ID: 619, Key: Channel ID] for every Suggestion channel.
- `One User Key per user`, per Suggestion channel. This can be avoided, and is explained later, [during setup](#setup).
## Features
- Multiple and seperate suggetion channels, with no hassle of prefixes for the general user.
- Anonymous user suggestion abililties. If disabled, doesn't allow user to edit/delete suggestions.
- Custom Upvote, Downvote and Neutral emojis.
- Allows management to approve, implement, deny, comment, edit, delete and mark suggestions as duplicates.
- Allows logging of approved, implemented, denied or deleted comments.
- Allows custom cooldowns on suggestions.
## Setup
- For one suggestion channel, make a Custom Command, and select the correct trigger type and trigger.
- Turn `Output Errors as command response` off, and restrict the CC - `Only run in the following channels` : You suggestion channel.
- Head to the configuration portion in the Custom Command.
  - First, decide if you want the user to be able to anonymously add suggestions. If you decide, then they will not be able to edit/delete suggestions.
  - If you want them to be able to, edit/delete suggestions, keep `{{$anon := false}}`, else, make it `{{$anon := true}}`.
  - Add the emojis for Upvote `{{$up}}`, Downvote `{{$down}}` and Neutral `{{$neutral}}`. Neutral is optional, only add it if you want to.
  - Add the topic to your suggestions in `{{$agenda}}`. It can be anything, even your Server name, if you are out of ideas.
  - Add the ID's of the roles you want to allow to manage suggestions in `{{$modRoles := cslice 123 456}}`. Roles with Administrative permissions do not have to be added here. 
  - For logging of suggestions, add the channel ID where you want the suggestions to be sent when they are approved/implemented/denied/deleted/dupe-marked. If you don't want them to be logged, keep that area blank.
  - If you want there to be a cooldown for users to be able to suggest, add it in `{{$cooldown}}` in __seconds__. This can also be kept as 0, and a cooldown can be added in the channel settings. No cooldown is not recommended, to avoid spammers/trolls.
  - If you wanted anonymous suggestions, make it `{{$userperm := sdict "edit" false "delete" false}}`. Otherwise, you can let users edit/delete suggestions by keeping them as true.
  - Add your wanted prefix in `{{$prefix}}`, keeping in mind, to escape it, if it is a special character. Remember, this prefix is only necessary for management of suggestions and not the suggestions themself.
  - Add the Channel ID you have in `{{$count}}`.
- You should be ready to go!
