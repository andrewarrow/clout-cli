# clout-cli

![image](https://user-images.githubusercontent.com/127054/119290208-f4527a80-bc00-11eb-9458-c29d828e4df0.png)

Welcome to the [cloutcli](https://bitclout.com/u/cloutcli) project.

There WAS a [bounty](https://stackoverflow.com/questions/67661276/how-do-i-properly-sign-a-bitclout-tx-in-golang-vs-typescript) for how to fix our tx signing to not need javascript. It was claimed by https://stackoverflow.com/users/589259/maarten-bodewes thanks!

To understand the code start with the big if else in main.go for each menu option. Many things are still in the package `main` but slowly moving into various packages.

[they own your coin](https://andrewarrow.substack.com/p/they-own-your-coin) is a blog article I wrote about bc in general.

# Entering your secret words

Everything cloutcli stores on your local drive it stores in `~/clout-cli-data` so when you want to remove everything just run:

```
$ rm -rf ~/clout-cli-data
```

And any secret you entered has been deleted.

# You can login with multiple accounts

```
$ ./clout login
~/clout-cli $ ./clout login
Enter mnenomic: lorem ipsum dolor sit amet consectetur adipiscing elit aenean ac mauris sit
Secret stored at: /Users/andrewarrow/clout-cli-data/secrets.txt

$ ./clout login
~/clout-cli $ ./clout login
Enter mnenomic: these are twelve different words from the words above this line done
Secret stored at: /Users/andrewarrow/clout-cli-data/secrets.txt
```

Each time you login the words are appened to that `secrets.txt` file.

```
$ ./clout accounts

andrewarrow
cloutcli

To select account, run `clout account [username]`
```

When I run the `accounts` command I see all my logged in accounts.

# Backup

`export CLOUT_PHRASE='these are some nice words and stuff.'`

If you set a `CLOUT_PHRASE` you can run `./clout backup` to
place your secrets.txt file with ALL your words into a encrypted file
that can only be read with your CLOUT_PHRASE.

Backup this one file, and it can get you into all your accounts.


# Building

Just run `go build`

# Examples

```
clout-cli $ ./clout help

  clout accounts               # list your various accounts
  clout diamond [username]     # award 1 diamond to last post
  clout follow [username]      # toggle follow
  clout followers              # who follows you
  clout following              # who you follow
  clout ls                     # list global posts
  clout ls --follow            # filter by follow
  clout ls --post=id           # show single post
  clout login                  # enter secret phrase
  clout logout                 # delete secret from drive
  clout messages               # list messages
  clout notifications          # list notifications
  clout post --reply=id        # post or reply
  clout reclout [username]     # reclout last post
  clout sync                   # fill local hard drive with data
  clout update [desc]          # update profile description
  clout whoami                 # base58 pubkey logged in
  clout [username]             # username's profile & posts
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

