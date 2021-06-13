# clout-cli

![image](https://user-images.githubusercontent.com/127054/119290208-f4527a80-bc00-11eb-9458-c29d828e4df0.png)

Welcome to the [cloutcli](https://bitclout.com/u/cloutcli) project.

[7 Minute Video Talk About This Project](https://vimeo.com/559521458)

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

If you don't have gcc installed [install it first](https://www.guru99.com/c-gcc-install.html).

Then run `go mod vendor` and then `go build`

See this [blog post](https://andrewarrow.substack.com/p/how-to-clone-build-and-run-clout).

# Linux (Fedora)

sudo dnf install bzr
sudo dnf install gtk3-devel
sudo dnf install webkit2gtk3-devel.x86_64
sudo dnf install ImageMagick

# Examples

```
~/clout-cli $ ./clout

  clout accounts               # list your various accounts
  clout backup                 # encrypt and copy secrets
  clout balances               # list available $bitclout
  clout boards                 # list boards you are on
  clout buy                    # buy creator coin
  clout diamond [username]     # award 1 diamond to last post
  clout follow [username]      # toggle follow
  clout followers              # who follows you
  clout following              # who you follow
  clout help                   # this menu
  clout like --hash=x          # like a post
  clout ls                     # list global posts
  clout ls --follow            # filter by follow
  clout ls --hash=x            # show single post
  clout login                  # enter secret phrase
  clout logout                 # delete secret from drive
  clout messages               # list messages
  clout notifications          # list notifications
  clout post --reply=x         # post or reply
  clout reclout [username]     # reclout last post
  clout sync                   # fill local hard drive with data
  clout update                 # update profile description
  clout wallet                 # list what you own
  clout whoami                 # base58 pubkey logged in
  clout [username]             # username's profile & posts
```

```
~/clout-cli $ ./clout ls --body
hash       username             body
----       --------             ----
b7cded5    MemeGod              Where are those people now? Probably buying BitClo
0d10064    RajLahoti            We just on-boarded the photographer here at #Clout
98da010    InURfeelz2           How are you going to celebrate the deflation bomb
28a5805    HIKIMBERLY           Deflation Bomb countdown
76af117    VishalGulia          Lmao 😂🤣
6b51d65    thorsten             🧐 I foresee a huge increase of sightings of the
34c30e2    BitCloutBuffett      @diamondhands walking around Miami tonight..
6bf58b3    DeflationBomb        Top 20 before 💣 detonation? We have until appro
dc9873c    YasminBcreative      I can’t get over how clear the water in Portugue
3ef8b62    Yellowredsparks      Is it just me or are your parents using more emoji
ec2430c    MemeGod              This is the moment we've all been waiting for! Get
9a9de37    pamelaanderson       It's amazing that we get to witness this first han
c36f4db    NotInMiami           Something for those of us #NotInMiami
8d3d91a    Klesh                Keep Calm and Carry On!
0539682    tijn                 Bomb Block: 33783 Current Block: 31113 To go: 2670
08659a4    connormitchell       wow you get on a plane to Miami and suddenly there
067d4f5    Tetono               Bitclout Deflation 💣  The total supply currentl
a21b9e9    Abhiandnow           This is going to bring so many changes to the scen
447e613    nigeleccles          Release the bomb!
84f51d1    Tetono               I see this as 📈
2ba1aac    CloutStreetBets      We knew this was all going to happen. Time to 🚀
d3f5bd3    Mirina               Change is coming!  I don't quite understand what.
```

```
~/clout-cli $ ./clout ls
username             ago             likes  replies  reclouts  cap        hash
--------             ---             -----  -------  --------  -------    -------
MemeGod              13 minutes      12     5        0         77.82      b7cded5
RajLahoti            14 minutes      33     5        2         5058.26    0d10064
InURfeelz2           15 minutes      5      3        0         34.50      98da010
HIKIMBERLY           16 minutes      5      1        1         123.22     28a5805
VishalGulia          17 minutes      8      2        1         14.39      76af117
thorsten             17 minutes      6      1        0         80.09      6b51d65
BitCloutBuffett      19 minutes      15     5        0         1290.88    34c30e2
DeflationBomb        28 minutes      10     6        1         163.06     6bf58b3
YasminBcreative      28 minutes      16     5        1         20.46      dc9873c
Yellowredsparks      31 minutes      14     6        0         37.78      3ef8b62
MemeGod              35 minutes      24     10       14        77.82      ec2430c
pamelaanderson       36 minutes      33     10       4         3361.05    9a9de37
NotInMiami           37 minutes      9      0        0         2.56       c36f4db
Klesh                42 minutes      17     3        2         48.31      8d3d91a
tijn                 42 minutes      15     6        0         572.18     0539682
connormitchell       42 minutes      14     2        0         981.33     08659a4
Tetono               about 1 hour    14     5        5         120.07     067d4f5
Abhiandnow           about 1 hour    11     2        1         9.23       a21b9e9
nigeleccles          about 1 hour    22     6        2         595.68     447e613
Tetono               about 1 hour    15     5        0         120.07     84f51d1
CloutStreetBets      about 1 hour    20     4        0         54.23      2ba1aac
Mirina               about 1 hour    20     9        1         25.00      d3f5bd3
```

```
~/clout-cli $ ./clout accounts

01. andrewarrow
02. cloutcli
03. cloutfactory
04. wolfschedule
05. imitate

To select account, run `clout account [username or i]`
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

