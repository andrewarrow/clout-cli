# clout-cli

![image](https://user-images.githubusercontent.com/127054/119290208-f4527a80-bc00-11eb-9458-c29d828e4df0.png)

Welcome to the [cloutcli](https://bitclout.com/u/cloutcli) project.

There is a [bounty](https://stackoverflow.com/questions/67661276/how-do-i-properly-sign-a-bitclout-tx-in-golang-vs-typescript) for how to fix our tx signing to not need javascript. 

To understand the code start with the big if else in main.go for each menu option. Many things are still in the package `main` but slowly moving into various packages.

[they own your coin](https://andrewarrow.substack.com/p/they-own-your-coin) is a blog article I wrote about bc in general.

# Building

Just run `go build`

# Examples

```
clout-cli $ ./clout help

  clout ls                     # list global posts
  clout ls --follow            # filter by follow
  clout [username]             # username's profile & posts
  clout login                  # enter secret phrase
  clout logout                 # delete secret from drive
  clout like [postHash]        # like/unlike a post
  clout diamond [postHash]     # send 1 diamond
  clout post --reply=postHash  # post or reply
  clout reclout [postHash]     # reclout specific post
  clout follow [username]      # toggle follow
  clout notifications          # list notifications
  clout followers [username]   # who follows username
  clout following              # who you follow
  clout whoami                 # base58 pubkey logged in
```

```
clout-cli $ ./clout ls
uditsonkhiya                   16                   less than a minute
         This so real😁😁😁😂😂😂🤣

LeighannBrindley               1                    less than a minute
         “Freedom begins with owning your flaws

ArtHero                        4                    less than a minute
         can i share my talent here?

Alexa_kim                      3                    less than a minute
         Thank you so much @teddybear! 😄

HenriA                         3                    less than a minute
         It`s all energy!

munny                          48                   less than a minute
         Joe Scarborough steadily warping into a

Phylanit                       4                    1 minute
         Who is your least favorite actor?

jefferydavid                   63                   1 minute
         I literally have "Boom Boom Boom let's g
```

```
clout-cli $ ./clout diamondhands
     0 Sounds awesome. Quintuple 💎💎💎💎💎 from me for the first person who does this.
3 days

     0 Today, we take the decentralization of social media further than any other project has in the past.
     2 Today, BitClout does to social media what Bitcoin is doing to the traditional financial system.
     4 Today, 100% of the BitClout code goes public.
     6 https://github.com/bitclout/core
4 days

     0 The more feedback you give us, the better the product gets. And the shipping never stops, not even on a beautiful sunny Saturday.
     2 Click to learn about some important upgrades to your Wallet and Creator Coin pages that we hope will promote HODLing and minimize scams. 👇
13 days
```

