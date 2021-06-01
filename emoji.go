package main

import (
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"html"
	"math/rand"
	"strconv"
	"strings"
)

func HandleClown() {
	items := ParseEmojiAsList()
	item1 := items[rand.Intn(len(items))]
	item2 := items[rand.Intn(len(items))]
	item3 := items[rand.Intn(len(items))]

	val1, _ := strconv.ParseInt(item1, 16, 64)
	val2, _ := strconv.ParseInt(item2, 16, 64)
	val3, _ := strconv.ParseInt(item3, 16, 64)
	str1 := html.UnescapeString(string(val1))
	str2 := html.UnescapeString(string(val2))
	str3 := html.UnescapeString(string(val3))
	text := fmt.Sprintf("%s%s%s = $%d", str1, str2, str3, SumIt(item1)+SumIt(item2)+SumIt(item3))
	mnemonic := session.ReadLoggedInWords()
	if mnemonic == "" {
		return
	}
	pub58, priv := keys.ComputeKeysFromSeed(session.SeedBytes(mnemonic))

	longHash := "a167e616c33047f73ce386bb877b0044b275ca59aa12af1a5a0312b10c3a756b"
	bigString := network.SubmitPost(pub58, text, longHash, "")

	var tx models.TxReady
	json.Unmarshal([]byte(bigString), &tx)

	jsonString := network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		fmt.Println("Success.")
	}
}

func SumIt(item string) int64 {
	sum := int64(0)
	for i, _ := range item {
		thing := item[i : i+1]
		val, _ := strconv.ParseInt(thing, 16, 64)
		sum += val
	}
	return sum
}

func ParseEmojiAsList() []string {
	list := []string{}
	for k, _ := range ParseEmoji() {
		list = append(list, k)
	}
	return list
}
func ParseEmoji() map[string]bool {
	m := map[string]bool{}
	tokens := strings.Split(emojiTableString, "<tr>")
	for _, token := range tokens {
		tokens = strings.Split(token, "<td>")
		if len(tokens) <= 4 {
			continue
		}

		thing := tokens[4]
		thing = thing[6 : len(thing)-7]
		tokens = strings.Split(thing, ";")
		for _, token = range tokens {
			if !strings.HasPrefix(token, "#x") {
				continue
			}
			thing = token[2:]
			m[thing] = true
		}
	}
	return m
}

var emojiTableString = `<tr><td>&#x1F600;<td><td>GRINNING FACE<td><code>&amp;#x1F600;</code><td><code>&amp;#128512;</code><td>
<tr><td>&#x1F601;<td><td>GRINNING FACE WITH SMILING EYES<td><code>&amp;#x1F601;</code><td><code>&amp;#128513;</code><td>
<tr><td>&#x1F602;<td><td>FACE WITH TEARS OF JOY<td><code>&amp;#x1F602;</code><td><code>&amp;#128514;</code><td>
<tr><td>&#x1F923;<td><td>ROLLING ON THE FLOOR LAUGHING<td><code>&amp;#x1F923;</code><td><code>&amp;#129315;</code><td>
<tr><td>&#x1F603;<td><td>SMILING FACE WITH OPEN MOUTH<td><code>&amp;#x1F603;</code><td><code>&amp;#128515;</code><td>
<tr><td>&#x1F604;<td><td>SMILING FACE WITH OPEN MOUTH AND SMILING EYES &#x224A; smiling face with open mouth &amp; smiling eyes<td><code>&amp;#x1F604;</code><td><code>&amp;#128516;</code><td>
<tr><td>&#x1F605;<td><td>SMILING FACE WITH OPEN MOUTH AND COLD SWEAT &#x224A; smiling face with open mouth &amp; cold sweat<td><code>&amp;#x1F605;</code><td><code>&amp;#128517;</code><td>
<tr><td>&#x1F606;<td><td>SMILING FACE WITH OPEN MOUTH AND TIGHTLY-CLOSED EYES &#x224A; smiling face with open mouth &amp; closed eyes<td><code>&amp;#x1F606;</code><td><code>&amp;#128518;</code><td>
<tr><td>&#x1F609;<td><td>WINKING FACE<td><code>&amp;#x1F609;</code><td><code>&amp;#128521;</code><td>
<tr><td>&#x1F60A;<td><td>SMILING FACE WITH SMILING EYES<td><code>&amp;#x1F60A;</code><td><code>&amp;#128522;</code><td>
<tr><td>&#x1F60B;<td><td>FACE SAVOURING DELICIOUS FOOD<td><code>&amp;#x1F60B;</code><td><code>&amp;#128523;</code><td>
<tr><td>&#x1F60E;<td><td>SMILING FACE WITH SUNGLASSES<td><code>&amp;#x1F60E;</code><td><code>&amp;#128526;</code><td>
<tr><td>&#x1F60D;<td><td>SMILING FACE WITH HEART-SHAPED EYES &#x224A; smiling face with heart-eyes<td><code>&amp;#x1F60D;</code><td><code>&amp;#128525;</code><td>
<tr><td>&#x1F618;<td><td>FACE THROWING A KISS<td><code>&amp;#x1F618;</code><td><code>&amp;#128536;</code><td>
<tr><td>&#x1F617;<td><td>KISSING FACE<td><code>&amp;#x1F617;</code><td><code>&amp;#128535;</code><td>
<tr><td>&#x1F619;<td><td>KISSING FACE WITH SMILING EYES<td><code>&amp;#x1F619;</code><td><code>&amp;#128537;</code><td>
<tr><td>&#x1F61A;<td><td>KISSING FACE WITH CLOSED EYES<td><code>&amp;#x1F61A;</code><td><code>&amp;#128538;</code><td>
<tr><td>&#x263A;<td><td>WHITE SMILING FACE &#x224A; smiling face<td><code>&amp;#x263A;</code><td><code>&amp;#9786;</code><td>
<tr><td>&#x1F642;<td><td>SLIGHTLY SMILING FACE<td><code>&amp;#x1F642;</code><td><code>&amp;#128578;</code><td>
<tr><td>&#x1F917;<td><td>HUGGING FACE<td><code>&amp;#x1F917;</code><td><code>&amp;#129303;</code><td>
<tr><td>&#x1F914;<td><td>THINKING FACE<td><code>&amp;#x1F914;</code><td><code>&amp;#129300;</code><td>
<tr><td>&#x1F610;<td><td>NEUTRAL FACE<td><code>&amp;#x1F610;</code><td><code>&amp;#128528;</code><td>
<tr><td>&#x1F611;<td><td>EXPRESSIONLESS FACE<td><code>&amp;#x1F611;</code><td><code>&amp;#128529;</code><td>
<tr><td>&#x1F636;<td><td>FACE WITHOUT MOUTH<td><code>&amp;#x1F636;</code><td><code>&amp;#128566;</code><td>
<tr><td>&#x1F644;<td><td>FACE WITH ROLLING EYES<td><code>&amp;#x1F644;</code><td><code>&amp;#128580;</code><td>
<tr><td>&#x1F60F;<td><td>SMIRKING FACE<td><code>&amp;#x1F60F;</code><td><code>&amp;#128527;</code><td>
<tr><td>&#x1F623;<td><td>PERSEVERING FACE<td><code>&amp;#x1F623;</code><td><code>&amp;#128547;</code><td>
<tr><td>&#x1F625;<td><td>DISAPPOINTED BUT RELIEVED FACE<td><code>&amp;#x1F625;</code><td><code>&amp;#128549;</code><td>
<tr><td>&#x1F62E;<td><td>FACE WITH OPEN MOUTH<td><code>&amp;#x1F62E;</code><td><code>&amp;#128558;</code><td>
<tr><td>&#x1F910;<td><td>ZIPPER-MOUTH FACE<td><code>&amp;#x1F910;</code><td><code>&amp;#129296;</code><td>
<tr><td>&#x1F62F;<td><td>HUSHED FACE<td><code>&amp;#x1F62F;</code><td><code>&amp;#128559;</code><td>
<tr><td>&#x1F62A;<td><td>SLEEPY FACE<td><code>&amp;#x1F62A;</code><td><code>&amp;#128554;</code><td>
<tr><td>&#x1F62B;<td><td>TIRED FACE<td><code>&amp;#x1F62B;</code><td><code>&amp;#128555;</code><td>
<tr><td>&#x1F634;<td><td>SLEEPING FACE<td><code>&amp;#x1F634;</code><td><code>&amp;#128564;</code><td>
<tr><td>&#x1F60C;<td><td>RELIEVED FACE<td><code>&amp;#x1F60C;</code><td><code>&amp;#128524;</code><td>
<tr><td>&#x1F913;<td><td>NERD FACE<td><code>&amp;#x1F913;</code><td><code>&amp;#129299;</code><td>
<tr><td>&#x1F61B;<td><td>FACE WITH STUCK-OUT TONGUE<td><code>&amp;#x1F61B;</code><td><code>&amp;#128539;</code><td>
<tr><td>&#x1F61C;<td><td>FACE WITH STUCK-OUT TONGUE AND WINKING EYE &#x224A; face with stuck-out tongue &amp; winking eye<td><code>&amp;#x1F61C;</code><td><code>&amp;#128540;</code><td>
<tr><td>&#x1F61D;<td><td>FACE WITH STUCK-OUT TONGUE AND TIGHTLY-CLOSED EYES &#x224A; face with stuck-out tongue &amp; closed eyes<td><code>&amp;#x1F61D;</code><td><code>&amp;#128541;</code><td>
<tr><td>&#x1F924;<td><td>DROOLING FACE<td><code>&amp;#x1F924;</code><td><code>&amp;#129316;</code><td>
<tr><td>&#x1F612;<td><td>UNAMUSED FACE<td><code>&amp;#x1F612;</code><td><code>&amp;#128530;</code><td>
<tr><td>&#x1F613;<td><td>FACE WITH COLD SWEAT<td><code>&amp;#x1F613;</code><td><code>&amp;#128531;</code><td>
<tr><td>&#x1F614;<td><td>PENSIVE FACE<td><code>&amp;#x1F614;</code><td><code>&amp;#128532;</code><td>
<tr><td>&#x1F615;<td><td>CONFUSED FACE<td><code>&amp;#x1F615;</code><td><code>&amp;#128533;</code><td>
<tr><td>&#x1F643;<td><td>UPSIDE-DOWN FACE<td><code>&amp;#x1F643;</code><td><code>&amp;#128579;</code><td>
<tr><td>&#x1F911;<td><td>MONEY-MOUTH FACE<td><code>&amp;#x1F911;</code><td><code>&amp;#129297;</code><td>
<tr><td>&#x1F632;<td><td>ASTONISHED FACE<td><code>&amp;#x1F632;</code><td><code>&amp;#128562;</code><td>
<tr><td>&#x2639;<td><td>WHITE FROWNING FACE &#x224A; frowning face<td><code>&amp;#x2639;</code><td><code>&amp;#9785;</code><td>
<tr><td>&#x1F641;<td><td>SLIGHTLY FROWNING FACE<td><code>&amp;#x1F641;</code><td><code>&amp;#128577;</code><td>
<tr><td>&#x1F616;<td><td>CONFOUNDED FACE<td><code>&amp;#x1F616;</code><td><code>&amp;#128534;</code><td>
<tr><td>&#x1F61E;<td><td>DISAPPOINTED FACE<td><code>&amp;#x1F61E;</code><td><code>&amp;#128542;</code><td>
<tr><td>&#x1F61F;<td><td>WORRIED FACE<td><code>&amp;#x1F61F;</code><td><code>&amp;#128543;</code><td>
<tr><td>&#x1F624;<td><td>FACE WITH LOOK OF TRIUMPH &#x224A; face with steam from nose<td><code>&amp;#x1F624;</code><td><code>&amp;#128548;</code><td>
<tr><td>&#x1F622;<td><td>CRYING FACE<td><code>&amp;#x1F622;</code><td><code>&amp;#128546;</code><td>
<tr><td>&#x1F62D;<td><td>LOUDLY CRYING FACE<td><code>&amp;#x1F62D;</code><td><code>&amp;#128557;</code><td>
<tr><td>&#x1F626;<td><td>FROWNING FACE WITH OPEN MOUTH<td><code>&amp;#x1F626;</code><td><code>&amp;#128550;</code><td>
<tr><td>&#x1F627;<td><td>ANGUISHED FACE<td><code>&amp;#x1F627;</code><td><code>&amp;#128551;</code><td>
<tr><td>&#x1F628;<td><td>FEARFUL FACE<td><code>&amp;#x1F628;</code><td><code>&amp;#128552;</code><td>
<tr><td>&#x1F629;<td><td>WEARY FACE<td><code>&amp;#x1F629;</code><td><code>&amp;#128553;</code><td>
<tr><td>&#x1F62C;<td><td>GRIMACING FACE<td><code>&amp;#x1F62C;</code><td><code>&amp;#128556;</code><td>
<tr><td>&#x1F630;<td><td>FACE WITH OPEN MOUTH AND COLD SWEAT &#x224A; face with open mouth &amp; cold sweat<td><code>&amp;#x1F630;</code><td><code>&amp;#128560;</code><td>
<tr><td>&#x1F631;<td><td>FACE SCREAMING IN FEAR<td><code>&amp;#x1F631;</code><td><code>&amp;#128561;</code><td>
<tr><td>&#x1F633;<td><td>FLUSHED FACE<td><code>&amp;#x1F633;</code><td><code>&amp;#128563;</code><td>
<tr><td>&#x1F635;<td><td>DIZZY FACE<td><code>&amp;#x1F635;</code><td><code>&amp;#128565;</code><td>
<tr><td>&#x1F621;<td><td>POUTING FACE<td><code>&amp;#x1F621;</code><td><code>&amp;#128545;</code><td>
<tr><td>&#x1F620;<td><td>ANGRY FACE<td><code>&amp;#x1F620;</code><td><code>&amp;#128544;</code><td>
<tr><td>&#x1F607;<td><td>SMILING FACE WITH HALO<td><code>&amp;#x1F607;</code><td><code>&amp;#128519;</code><td>
<tr><td>&#x1F920;<td><td>FACE WITH COWBOY HAT &#x224A; cowboy hat face<td><code>&amp;#x1F920;</code><td><code>&amp;#129312;</code><td>
<tr><td>&#x1F921;<td><td>CLOWN FACE<td><code>&amp;#x1F921;</code><td><code>&amp;#129313;</code><td>
<tr><td>&#x1F925;<td><td>LYING FACE<td><code>&amp;#x1F925;</code><td><code>&amp;#129317;</code><td>
<tr><td>&#x1F637;<td><td>FACE WITH MEDICAL MASK<td><code>&amp;#x1F637;</code><td><code>&amp;#128567;</code><td>
<tr><td>&#x1F912;<td><td>FACE WITH THERMOMETER<td><code>&amp;#x1F912;</code><td><code>&amp;#129298;</code><td>
<tr><td>&#x1F915;<td><td>FACE WITH HEAD-BANDAGE<td><code>&amp;#x1F915;</code><td><code>&amp;#129301;</code><td>
<tr><td>&#x1F922;<td><td>NAUSEATED FACE<td><code>&amp;#x1F922;</code><td><code>&amp;#129314;</code><td>
<tr><td>&#x1F927;<td><td>SNEEZING FACE<td><code>&amp;#x1F927;</code><td><code>&amp;#129319;</code><td>
<tr><td>&#x1F608;<td><td>SMILING FACE WITH HORNS<td><code>&amp;#x1F608;</code><td><code>&amp;#128520;</code><td>
<tr><td>&#x1F47F;<td><td>IMP<td><code>&amp;#x1F47F;</code><td><code>&amp;#128127;</code><td>
<tr><td>&#x1F479;<td><td>JAPANESE OGRE &#x224A; ogre<td><code>&amp;#x1F479;</code><td><code>&amp;#128121;</code><td>
<tr><td>&#x1F47A;<td><td>JAPANESE GOBLIN &#x224A; goblin<td><code>&amp;#x1F47A;</code><td><code>&amp;#128122;</code><td>
<tr><td>&#x1F480;<td><td>SKULL<td><code>&amp;#x1F480;</code><td><code>&amp;#128128;</code><td>
<tr><td>&#x2620;<td><td>SKULL AND CROSSBONES<td><code>&amp;#x2620;</code><td><code>&amp;#9760;</code><td>
<tr><td>&#x1F47B;<td><td>GHOST<td><code>&amp;#x1F47B;</code><td><code>&amp;#128123;</code><td>
<tr><td>&#x1F47D;<td><td>EXTRATERRESTRIAL ALIEN &#x224A; alien<td><code>&amp;#x1F47D;</code><td><code>&amp;#128125;</code><td>
<tr><td>&#x1F47E;<td><td>ALIEN MONSTER<td><code>&amp;#x1F47E;</code><td><code>&amp;#128126;</code><td>
<tr><td>&#x1F916;<td><td>ROBOT FACE<td><code>&amp;#x1F916;</code><td><code>&amp;#129302;</code><td>
<tr><td>&#x1F4A9;<td><td>PILE OF POO<td><code>&amp;#x1F4A9;</code><td><code>&amp;#128169;</code><td>
<tr><td>&#x1F63A;<td><td>SMILING CAT FACE WITH OPEN MOUTH<td><code>&amp;#x1F63A;</code><td><code>&amp;#128570;</code><td>
<tr><td>&#x1F638;<td><td>GRINNING CAT FACE WITH SMILING EYES<td><code>&amp;#x1F638;</code><td><code>&amp;#128568;</code><td>
<tr><td>&#x1F639;<td><td>CAT FACE WITH TEARS OF JOY<td><code>&amp;#x1F639;</code><td><code>&amp;#128569;</code><td>
<tr><td>&#x1F63B;<td><td>SMILING CAT FACE WITH HEART-SHAPED EYES &#x224A; smiling cat face with heart-eyes<td><code>&amp;#x1F63B;</code><td><code>&amp;#128571;</code><td>
<tr><td>&#x1F63C;<td><td>CAT FACE WITH WRY SMILE<td><code>&amp;#x1F63C;</code><td><code>&amp;#128572;</code><td>
<tr><td>&#x1F63D;<td><td>KISSING CAT FACE WITH CLOSED EYES<td><code>&amp;#x1F63D;</code><td><code>&amp;#128573;</code><td>
<tr><td>&#x1F640;<td><td>WEARY CAT FACE<td><code>&amp;#x1F640;</code><td><code>&amp;#128576;</code><td>
<tr><td>&#x1F63F;<td><td>CRYING CAT FACE<td><code>&amp;#x1F63F;</code><td><code>&amp;#128575;</code><td>
<tr><td>&#x1F63E;<td><td>POUTING CAT FACE<td><code>&amp;#x1F63E;</code><td><code>&amp;#128574;</code><td>
<tr><td>&#x1F648;<td><td>SEE-NO-EVIL MONKEY &#x224A; see-no-evil<td><code>&amp;#x1F648;</code><td><code>&amp;#128584;</code><td>
<tr><td>&#x1F649;<td><td>HEAR-NO-EVIL MONKEY &#x224A; hear-no-evil<td><code>&amp;#x1F649;</code><td><code>&amp;#128585;</code><td>
<tr><td>&#x1F64A;<td><td>SPEAK-NO-EVIL MONKEY &#x224A; speak-no-evil<td><code>&amp;#x1F64A;</code><td><code>&amp;#128586;</code><td>
<tr><td>&#x1F466;<td><td>BOY<td><code>&amp;#x1F466;</code><td><code>&amp;#128102;</code><td>
<tr><td>&#x1F466;&#x1F3FB;<td><td>boy, type-1-2<td><code>&amp;#x1F466; &amp;#x1F3FB;</code><td><code>&amp;#128102; &amp;#127995;</code><td>
<tr><td>&#x1F466;&#x1F3FC;<td><td>boy, type-3<td><code>&amp;#x1F466; &amp;#x1F3FC;</code><td><code>&amp;#128102; &amp;#127996;</code><td>
<tr><td>&#x1F466;&#x1F3FD;<td><td>boy, type-4<td><code>&amp;#x1F466; &amp;#x1F3FD;</code><td><code>&amp;#128102; &amp;#127997;</code><td>
<tr><td>&#x1F466;&#x1F3FE;<td><td>boy, type-5<td><code>&amp;#x1F466; &amp;#x1F3FE;</code><td><code>&amp;#128102; &amp;#127998;</code><td>
<tr><td>&#x1F466;&#x1F3FF;<td><td>boy, type-6<td><code>&amp;#x1F466; &amp;#x1F3FF;</code><td><code>&amp;#128102; &amp;#127999;</code><td>
<tr><td>&#x1F467;<td><td>GIRL<td><code>&amp;#x1F467;</code><td><code>&amp;#128103;</code><td>
<tr><td>&#x1F467;&#x1F3FB;<td><td>girl, type-1-2<td><code>&amp;#x1F467; &amp;#x1F3FB;</code><td><code>&amp;#128103; &amp;#127995;</code><td>
<tr><td>&#x1F467;&#x1F3FC;<td><td>girl, type-3<td><code>&amp;#x1F467; &amp;#x1F3FC;</code><td><code>&amp;#128103; &amp;#127996;</code><td>
<tr><td>&#x1F467;&#x1F3FD;<td><td>girl, type-4<td><code>&amp;#x1F467; &amp;#x1F3FD;</code><td><code>&amp;#128103; &amp;#127997;</code><td>
<tr><td>&#x1F467;&#x1F3FE;<td><td>girl, type-5<td><code>&amp;#x1F467; &amp;#x1F3FE;</code><td><code>&amp;#128103; &amp;#127998;</code><td>
<tr><td>&#x1F467;&#x1F3FF;<td><td>girl, type-6<td><code>&amp;#x1F467; &amp;#x1F3FF;</code><td><code>&amp;#128103; &amp;#127999;</code><td>
<tr><td>&#x1F468;<td><td>MAN<td><code>&amp;#x1F468;</code><td><code>&amp;#128104;</code><td>
<tr><td>&#x1F468;&#x1F3FB;<td><td>man, type-1-2<td><code>&amp;#x1F468; &amp;#x1F3FB;</code><td><code>&amp;#128104; &amp;#127995;</code><td>
<tr><td>&#x1F468;&#x1F3FC;<td><td>man, type-3<td><code>&amp;#x1F468; &amp;#x1F3FC;</code><td><code>&amp;#128104; &amp;#127996;</code><td>
<tr><td>&#x1F468;&#x1F3FD;<td><td>man, type-4<td><code>&amp;#x1F468; &amp;#x1F3FD;</code><td><code>&amp;#128104; &amp;#127997;</code><td>
<tr><td>&#x1F468;&#x1F3FE;<td><td>man, type-5<td><code>&amp;#x1F468; &amp;#x1F3FE;</code><td><code>&amp;#128104; &amp;#127998;</code><td>
<tr><td>&#x1F468;&#x1F3FF;<td><td>man, type-6<td><code>&amp;#x1F468; &amp;#x1F3FF;</code><td><code>&amp;#128104; &amp;#127999;</code><td>
<tr><td>&#x1F469;<td><td>WOMAN<td><code>&amp;#x1F469;</code><td><code>&amp;#128105;</code><td>
<tr><td>&#x1F469;&#x1F3FB;<td><td>woman, type-1-2<td><code>&amp;#x1F469; &amp;#x1F3FB;</code><td><code>&amp;#128105; &amp;#127995;</code><td>
<tr><td>&#x1F469;&#x1F3FC;<td><td>woman, type-3<td><code>&amp;#x1F469; &amp;#x1F3FC;</code><td><code>&amp;#128105; &amp;#127996;</code><td>
<tr><td>&#x1F469;&#x1F3FD;<td><td>woman, type-4<td><code>&amp;#x1F469; &amp;#x1F3FD;</code><td><code>&amp;#128105; &amp;#127997;</code><td>
<tr><td>&#x1F469;&#x1F3FE;<td><td>woman, type-5<td><code>&amp;#x1F469; &amp;#x1F3FE;</code><td><code>&amp;#128105; &amp;#127998;</code><td>
<tr><td>&#x1F469;&#x1F3FF;<td><td>woman, type-6<td><code>&amp;#x1F469; &amp;#x1F3FF;</code><td><code>&amp;#128105; &amp;#127999;</code><td>
<tr><td>&#x1F474;<td><td>OLDER MAN &#x224A; old man<td><code>&amp;#x1F474;</code><td><code>&amp;#128116;</code><td>
<tr><td>&#x1F474;&#x1F3FB;<td><td>old man, type-1-2<td><code>&amp;#x1F474; &amp;#x1F3FB;</code><td><code>&amp;#128116; &amp;#127995;</code><td>
<tr><td>&#x1F474;&#x1F3FC;<td><td>old man, type-3<td><code>&amp;#x1F474; &amp;#x1F3FC;</code><td><code>&amp;#128116; &amp;#127996;</code><td>
<tr><td>&#x1F474;&#x1F3FD;<td><td>old man, type-4<td><code>&amp;#x1F474; &amp;#x1F3FD;</code><td><code>&amp;#128116; &amp;#127997;</code><td>
<tr><td>&#x1F474;&#x1F3FE;<td><td>old man, type-5<td><code>&amp;#x1F474; &amp;#x1F3FE;</code><td><code>&amp;#128116; &amp;#127998;</code><td>
<tr><td>&#x1F474;&#x1F3FF;<td><td>old man, type-6<td><code>&amp;#x1F474; &amp;#x1F3FF;</code><td><code>&amp;#128116; &amp;#127999;</code><td>
<tr><td>&#x1F475;<td><td>OLDER WOMAN &#x224A; old woman<td><code>&amp;#x1F475;</code><td><code>&amp;#128117;</code><td>
<tr><td>&#x1F475;&#x1F3FB;<td><td>old woman, type-1-2<td><code>&amp;#x1F475; &amp;#x1F3FB;</code><td><code>&amp;#128117; &amp;#127995;</code><td>
<tr><td>&#x1F475;&#x1F3FC;<td><td>old woman, type-3<td><code>&amp;#x1F475; &amp;#x1F3FC;</code><td><code>&amp;#128117; &amp;#127996;</code><td>
<tr><td>&#x1F475;&#x1F3FD;<td><td>old woman, type-4<td><code>&amp;#x1F475; &amp;#x1F3FD;</code><td><code>&amp;#128117; &amp;#127997;</code><td>
<tr><td>&#x1F475;&#x1F3FE;<td><td>old woman, type-5<td><code>&amp;#x1F475; &amp;#x1F3FE;</code><td><code>&amp;#128117; &amp;#127998;</code><td>
<tr><td>&#x1F475;&#x1F3FF;<td><td>old woman, type-6<td><code>&amp;#x1F475; &amp;#x1F3FF;</code><td><code>&amp;#128117; &amp;#127999;</code><td>
<tr><td>&#x1F476;<td><td>BABY<td><code>&amp;#x1F476;</code><td><code>&amp;#128118;</code><td>
<tr><td>&#x1F476;&#x1F3FB;<td><td>baby, type-1-2<td><code>&amp;#x1F476; &amp;#x1F3FB;</code><td><code>&amp;#128118; &amp;#127995;</code><td>
<tr><td>&#x1F476;&#x1F3FC;<td><td>baby, type-3<td><code>&amp;#x1F476; &amp;#x1F3FC;</code><td><code>&amp;#128118; &amp;#127996;</code><td>
<tr><td>&#x1F476;&#x1F3FD;<td><td>baby, type-4<td><code>&amp;#x1F476; &amp;#x1F3FD;</code><td><code>&amp;#128118; &amp;#127997;</code><td>
<tr><td>&#x1F476;&#x1F3FE;<td><td>baby, type-5<td><code>&amp;#x1F476; &amp;#x1F3FE;</code><td><code>&amp;#128118; &amp;#127998;</code><td>
<tr><td>&#x1F476;&#x1F3FF;<td><td>baby, type-6<td><code>&amp;#x1F476; &amp;#x1F3FF;</code><td><code>&amp;#128118; &amp;#127999;</code><td>
<tr><td>&#x1F47C;<td><td>BABY ANGEL<td><code>&amp;#x1F47C;</code><td><code>&amp;#128124;</code><td>
<tr><td>&#x1F47C;&#x1F3FB;<td><td>baby angel, type-1-2<td><code>&amp;#x1F47C; &amp;#x1F3FB;</code><td><code>&amp;#128124; &amp;#127995;</code><td>
<tr><td>&#x1F47C;&#x1F3FC;<td><td>baby angel, type-3<td><code>&amp;#x1F47C; &amp;#x1F3FC;</code><td><code>&amp;#128124; &amp;#127996;</code><td>
<tr><td>&#x1F47C;&#x1F3FD;<td><td>baby angel, type-4<td><code>&amp;#x1F47C; &amp;#x1F3FD;</code><td><code>&amp;#128124; &amp;#127997;</code><td>
<tr><td>&#x1F47C;&#x1F3FE;<td><td>baby angel, type-5<td><code>&amp;#x1F47C; &amp;#x1F3FE;</code><td><code>&amp;#128124; &amp;#127998;</code><td>
<tr><td>&#x1F47C;&#x1F3FF;<td><td>baby angel, type-6<td><code>&amp;#x1F47C; &amp;#x1F3FF;</code><td><code>&amp;#128124; &amp;#127999;</code><td>
<tr><td>&#x1F46E;<td><td>POLICE OFFICER<td><code>&amp;#x1F46E;</code><td><code>&amp;#128110;</code><td>
<tr><td>&#x1F46E;&#x1F3FB;<td><td>police officer, type-1-2<td><code>&amp;#x1F46E; &amp;#x1F3FB;</code><td><code>&amp;#128110; &amp;#127995;</code><td>
<tr><td>&#x1F46E;&#x1F3FC;<td><td>police officer, type-3<td><code>&amp;#x1F46E; &amp;#x1F3FC;</code><td><code>&amp;#128110; &amp;#127996;</code><td>
<tr><td>&#x1F46E;&#x1F3FD;<td><td>police officer, type-4<td><code>&amp;#x1F46E; &amp;#x1F3FD;</code><td><code>&amp;#128110; &amp;#127997;</code><td>
<tr><td>&#x1F46E;&#x1F3FE;<td><td>police officer, type-5<td><code>&amp;#x1F46E; &amp;#x1F3FE;</code><td><code>&amp;#128110; &amp;#127998;</code><td>
<tr><td>&#x1F46E;&#x1F3FF;<td><td>police officer, type-6<td><code>&amp;#x1F46E; &amp;#x1F3FF;</code><td><code>&amp;#128110; &amp;#127999;</code><td>
<tr><td>&#x1F575;<td><td>SLEUTH OR SPY &#x224A; detective<td><code>&amp;#x1F575;</code><td><code>&amp;#128373;</code><td>
<tr><td>&#x1F575;&#x1F3FB;<td><td>detective, type-1-2<td><code>&amp;#x1F575; &amp;#x1F3FB;</code><td><code>&amp;#128373; &amp;#127995;</code><td>
<tr><td>&#x1F575;&#x1F3FC;<td><td>detective, type-3<td><code>&amp;#x1F575; &amp;#x1F3FC;</code><td><code>&amp;#128373; &amp;#127996;</code><td>
<tr><td>&#x1F575;&#x1F3FD;<td><td>detective, type-4<td><code>&amp;#x1F575; &amp;#x1F3FD;</code><td><code>&amp;#128373; &amp;#127997;</code><td>
<tr><td>&#x1F575;&#x1F3FE;<td><td>detective, type-5<td><code>&amp;#x1F575; &amp;#x1F3FE;</code><td><code>&amp;#128373; &amp;#127998;</code><td>
<tr><td>&#x1F575;&#x1F3FF;<td><td>detective, type-6<td><code>&amp;#x1F575; &amp;#x1F3FF;</code><td><code>&amp;#128373; &amp;#127999;</code><td>
<tr><td>&#x1F482;<td><td>GUARDSMAN<td><code>&amp;#x1F482;</code><td><code>&amp;#128130;</code><td>
<tr><td>&#x1F482;&#x1F3FB;<td><td>guardsman, type-1-2<td><code>&amp;#x1F482; &amp;#x1F3FB;</code><td><code>&amp;#128130; &amp;#127995;</code><td>
<tr><td>&#x1F482;&#x1F3FC;<td><td>guardsman, type-3<td><code>&amp;#x1F482; &amp;#x1F3FC;</code><td><code>&amp;#128130; &amp;#127996;</code><td>
<tr><td>&#x1F482;&#x1F3FD;<td><td>guardsman, type-4<td><code>&amp;#x1F482; &amp;#x1F3FD;</code><td><code>&amp;#128130; &amp;#127997;</code><td>
<tr><td>&#x1F482;&#x1F3FE;<td><td>guardsman, type-5<td><code>&amp;#x1F482; &amp;#x1F3FE;</code><td><code>&amp;#128130; &amp;#127998;</code><td>
<tr><td>&#x1F482;&#x1F3FF;<td><td>guardsman, type-6<td><code>&amp;#x1F482; &amp;#x1F3FF;</code><td><code>&amp;#128130; &amp;#127999;</code><td>
<tr><td>&#x1F477;<td><td>CONSTRUCTION WORKER<td><code>&amp;#x1F477;</code><td><code>&amp;#128119;</code><td>
<tr><td>&#x1F477;&#x1F3FB;<td><td>construction worker, type-1-2<td><code>&amp;#x1F477; &amp;#x1F3FB;</code><td><code>&amp;#128119; &amp;#127995;</code><td>
<tr><td>&#x1F477;&#x1F3FC;<td><td>construction worker, type-3<td><code>&amp;#x1F477; &amp;#x1F3FC;</code><td><code>&amp;#128119; &amp;#127996;</code><td>
<tr><td>&#x1F477;&#x1F3FD;<td><td>construction worker, type-4<td><code>&amp;#x1F477; &amp;#x1F3FD;</code><td><code>&amp;#128119; &amp;#127997;</code><td>
<tr><td>&#x1F477;&#x1F3FE;<td><td>construction worker, type-5<td><code>&amp;#x1F477; &amp;#x1F3FE;</code><td><code>&amp;#128119; &amp;#127998;</code><td>
<tr><td>&#x1F477;&#x1F3FF;<td><td>construction worker, type-6<td><code>&amp;#x1F477; &amp;#x1F3FF;</code><td><code>&amp;#128119; &amp;#127999;</code><td>
<tr><td>&#x1F473;<td><td>MAN WITH TURBAN &#x224A; person with turban<td><code>&amp;#x1F473;</code><td><code>&amp;#128115;</code><td>
<tr><td>&#x1F473;&#x1F3FB;<td><td>person with turban, type-1-2<td><code>&amp;#x1F473; &amp;#x1F3FB;</code><td><code>&amp;#128115; &amp;#127995;</code><td>
<tr><td>&#x1F473;&#x1F3FC;<td><td>person with turban, type-3<td><code>&amp;#x1F473; &amp;#x1F3FC;</code><td><code>&amp;#128115; &amp;#127996;</code><td>
<tr><td>&#x1F473;&#x1F3FD;<td><td>person with turban, type-4<td><code>&amp;#x1F473; &amp;#x1F3FD;</code><td><code>&amp;#128115; &amp;#127997;</code><td>
<tr><td>&#x1F473;&#x1F3FE;<td><td>person with turban, type-5<td><code>&amp;#x1F473; &amp;#x1F3FE;</code><td><code>&amp;#128115; &amp;#127998;</code><td>
<tr><td>&#x1F473;&#x1F3FF;<td><td>person with turban, type-6<td><code>&amp;#x1F473; &amp;#x1F3FF;</code><td><code>&amp;#128115; &amp;#127999;</code><td>
<tr><td>&#x1F471;<td><td>PERSON WITH BLOND HAIR &#x224A; blond person<td><code>&amp;#x1F471;</code><td><code>&amp;#128113;</code><td>
<tr><td>&#x1F471;&#x1F3FB;<td><td>blond person, type-1-2<td><code>&amp;#x1F471; &amp;#x1F3FB;</code><td><code>&amp;#128113; &amp;#127995;</code><td>
<tr><td>&#x1F471;&#x1F3FC;<td><td>blond person, type-3<td><code>&amp;#x1F471; &amp;#x1F3FC;</code><td><code>&amp;#128113; &amp;#127996;</code><td>
<tr><td>&#x1F471;&#x1F3FD;<td><td>blond person, type-4<td><code>&amp;#x1F471; &amp;#x1F3FD;</code><td><code>&amp;#128113; &amp;#127997;</code><td>
<tr><td>&#x1F471;&#x1F3FE;<td><td>blond person, type-5<td><code>&amp;#x1F471; &amp;#x1F3FE;</code><td><code>&amp;#128113; &amp;#127998;</code><td>
<tr><td>&#x1F471;&#x1F3FF;<td><td>blond person, type-6<td><code>&amp;#x1F471; &amp;#x1F3FF;</code><td><code>&amp;#128113; &amp;#127999;</code><td>
<tr><td>&#x1F385;<td><td>FATHER CHRISTMAS &#x224A; santa claus<td><code>&amp;#x1F385;</code><td><code>&amp;#127877;</code><td>
<tr><td>&#x1F385;&#x1F3FB;<td><td>santa claus, type-1-2<td><code>&amp;#x1F385; &amp;#x1F3FB;</code><td><code>&amp;#127877; &amp;#127995;</code><td>
<tr><td>&#x1F385;&#x1F3FC;<td><td>santa claus, type-3<td><code>&amp;#x1F385; &amp;#x1F3FC;</code><td><code>&amp;#127877; &amp;#127996;</code><td>
<tr><td>&#x1F385;&#x1F3FD;<td><td>santa claus, type-4<td><code>&amp;#x1F385; &amp;#x1F3FD;</code><td><code>&amp;#127877; &amp;#127997;</code><td>
<tr><td>&#x1F385;&#x1F3FE;<td><td>santa claus, type-5<td><code>&amp;#x1F385; &amp;#x1F3FE;</code><td><code>&amp;#127877; &amp;#127998;</code><td>
<tr><td>&#x1F385;&#x1F3FF;<td><td>santa claus, type-6<td><code>&amp;#x1F385; &amp;#x1F3FF;</code><td><code>&amp;#127877; &amp;#127999;</code><td>
<tr><td>&#x1F936;<td><td>MOTHER CHRISTMAS<td><code>&amp;#x1F936;</code><td><code>&amp;#129334;</code><td>
<tr><td>&#x1F936;&#x1F3FB;<td><td>mother christmas, type-1-2<td><code>&amp;#x1F936; &amp;#x1F3FB;</code><td><code>&amp;#129334; &amp;#127995;</code><td>
<tr><td>&#x1F936;&#x1F3FC;<td><td>mother christmas, type-3<td><code>&amp;#x1F936; &amp;#x1F3FC;</code><td><code>&amp;#129334; &amp;#127996;</code><td>
<tr><td>&#x1F936;&#x1F3FD;<td><td>mother christmas, type-4<td><code>&amp;#x1F936; &amp;#x1F3FD;</code><td><code>&amp;#129334; &amp;#127997;</code><td>
<tr><td>&#x1F936;&#x1F3FE;<td><td>mother christmas, type-5<td><code>&amp;#x1F936; &amp;#x1F3FE;</code><td><code>&amp;#129334; &amp;#127998;</code><td>
<tr><td>&#x1F936;&#x1F3FF;<td><td>mother christmas, type-6<td><code>&amp;#x1F936; &amp;#x1F3FF;</code><td><code>&amp;#129334; &amp;#127999;</code><td>
<tr><td>&#x1F478;<td><td>PRINCESS<td><code>&amp;#x1F478;</code><td><code>&amp;#128120;</code><td>
<tr><td>&#x1F478;&#x1F3FB;<td><td>princess, type-1-2<td><code>&amp;#x1F478; &amp;#x1F3FB;</code><td><code>&amp;#128120; &amp;#127995;</code><td>
<tr><td>&#x1F478;&#x1F3FC;<td><td>princess, type-3<td><code>&amp;#x1F478; &amp;#x1F3FC;</code><td><code>&amp;#128120; &amp;#127996;</code><td>
<tr><td>&#x1F478;&#x1F3FD;<td><td>princess, type-4<td><code>&amp;#x1F478; &amp;#x1F3FD;</code><td><code>&amp;#128120; &amp;#127997;</code><td>
<tr><td>&#x1F478;&#x1F3FE;<td><td>princess, type-5<td><code>&amp;#x1F478; &amp;#x1F3FE;</code><td><code>&amp;#128120; &amp;#127998;</code><td>
<tr><td>&#x1F478;&#x1F3FF;<td><td>princess, type-6<td><code>&amp;#x1F478; &amp;#x1F3FF;</code><td><code>&amp;#128120; &amp;#127999;</code><td>
<tr><td>&#x1F934;<td><td>PRINCE<td><code>&amp;#x1F934;</code><td><code>&amp;#129332;</code><td>
<tr><td>&#x1F934;&#x1F3FB;<td><td>prince, type-1-2<td><code>&amp;#x1F934; &amp;#x1F3FB;</code><td><code>&amp;#129332; &amp;#127995;</code><td>
<tr><td>&#x1F934;&#x1F3FC;<td><td>prince, type-3<td><code>&amp;#x1F934; &amp;#x1F3FC;</code><td><code>&amp;#129332; &amp;#127996;</code><td>
<tr><td>&#x1F934;&#x1F3FD;<td><td>prince, type-4<td><code>&amp;#x1F934; &amp;#x1F3FD;</code><td><code>&amp;#129332; &amp;#127997;</code><td>
<tr><td>&#x1F934;&#x1F3FE;<td><td>prince, type-5<td><code>&amp;#x1F934; &amp;#x1F3FE;</code><td><code>&amp;#129332; &amp;#127998;</code><td>
<tr><td>&#x1F934;&#x1F3FF;<td><td>prince, type-6<td><code>&amp;#x1F934; &amp;#x1F3FF;</code><td><code>&amp;#129332; &amp;#127999;</code><td>
<tr><td>&#x1F470;<td><td>BRIDE WITH VEIL<td><code>&amp;#x1F470;</code><td><code>&amp;#128112;</code><td>
<tr><td>&#x1F470;&#x1F3FB;<td><td>bride with veil, type-1-2<td><code>&amp;#x1F470; &amp;#x1F3FB;</code><td><code>&amp;#128112; &amp;#127995;</code><td>
<tr><td>&#x1F470;&#x1F3FC;<td><td>bride with veil, type-3<td><code>&amp;#x1F470; &amp;#x1F3FC;</code><td><code>&amp;#128112; &amp;#127996;</code><td>
<tr><td>&#x1F470;&#x1F3FD;<td><td>bride with veil, type-4<td><code>&amp;#x1F470; &amp;#x1F3FD;</code><td><code>&amp;#128112; &amp;#127997;</code><td>
<tr><td>&#x1F470;&#x1F3FE;<td><td>bride with veil, type-5<td><code>&amp;#x1F470; &amp;#x1F3FE;</code><td><code>&amp;#128112; &amp;#127998;</code><td>
<tr><td>&#x1F470;&#x1F3FF;<td><td>bride with veil, type-6<td><code>&amp;#x1F470; &amp;#x1F3FF;</code><td><code>&amp;#128112; &amp;#127999;</code><td>
<tr><td>&#x1F935;<td><td>MAN IN TUXEDO<td><code>&amp;#x1F935;</code><td><code>&amp;#129333;</code><td>
<tr><td>&#x1F935;&#x1F3FB;<td><td>man in tuxedo, type-1-2<td><code>&amp;#x1F935; &amp;#x1F3FB;</code><td><code>&amp;#129333; &amp;#127995;</code><td>
<tr><td>&#x1F935;&#x1F3FC;<td><td>man in tuxedo, type-3<td><code>&amp;#x1F935; &amp;#x1F3FC;</code><td><code>&amp;#129333; &amp;#127996;</code><td>
<tr><td>&#x1F935;&#x1F3FD;<td><td>man in tuxedo, type-4<td><code>&amp;#x1F935; &amp;#x1F3FD;</code><td><code>&amp;#129333; &amp;#127997;</code><td>
<tr><td>&#x1F935;&#x1F3FE;<td><td>man in tuxedo, type-5<td><code>&amp;#x1F935; &amp;#x1F3FE;</code><td><code>&amp;#129333; &amp;#127998;</code><td>
<tr><td>&#x1F935;&#x1F3FF;<td><td>man in tuxedo, type-6<td><code>&amp;#x1F935; &amp;#x1F3FF;</code><td><code>&amp;#129333; &amp;#127999;</code><td>
<tr><td>&#x1F930;<td><td>PREGNANT WOMAN<td><code>&amp;#x1F930;</code><td><code>&amp;#129328;</code><td>
<tr><td>&#x1F930;&#x1F3FB;<td><td>pregnant woman, type-1-2<td><code>&amp;#x1F930; &amp;#x1F3FB;</code><td><code>&amp;#129328; &amp;#127995;</code><td>
<tr><td>&#x1F930;&#x1F3FC;<td><td>pregnant woman, type-3<td><code>&amp;#x1F930; &amp;#x1F3FC;</code><td><code>&amp;#129328; &amp;#127996;</code><td>
<tr><td>&#x1F930;&#x1F3FD;<td><td>pregnant woman, type-4<td><code>&amp;#x1F930; &amp;#x1F3FD;</code><td><code>&amp;#129328; &amp;#127997;</code><td>
<tr><td>&#x1F930;&#x1F3FE;<td><td>pregnant woman, type-5<td><code>&amp;#x1F930; &amp;#x1F3FE;</code><td><code>&amp;#129328; &amp;#127998;</code><td>
<tr><td>&#x1F930;&#x1F3FF;<td><td>pregnant woman, type-6<td><code>&amp;#x1F930; &amp;#x1F3FF;</code><td><code>&amp;#129328; &amp;#127999;</code><td>
<tr><td>&#x1F472;<td><td>MAN WITH GUA PI MAO &#x224A; man with chinese cap<td><code>&amp;#x1F472;</code><td><code>&amp;#128114;</code><td>
<tr><td>&#x1F472;&#x1F3FB;<td><td>man with chinese cap, type-1-2<td><code>&amp;#x1F472; &amp;#x1F3FB;</code><td><code>&amp;#128114; &amp;#127995;</code><td>
<tr><td>&#x1F472;&#x1F3FC;<td><td>man with chinese cap, type-3<td><code>&amp;#x1F472; &amp;#x1F3FC;</code><td><code>&amp;#128114; &amp;#127996;</code><td>
<tr><td>&#x1F472;&#x1F3FD;<td><td>man with chinese cap, type-4<td><code>&amp;#x1F472; &amp;#x1F3FD;</code><td><code>&amp;#128114; &amp;#127997;</code><td>
<tr><td>&#x1F472;&#x1F3FE;<td><td>man with chinese cap, type-5<td><code>&amp;#x1F472; &amp;#x1F3FE;</code><td><code>&amp;#128114; &amp;#127998;</code><td>
<tr><td>&#x1F472;&#x1F3FF;<td><td>man with chinese cap, type-6<td><code>&amp;#x1F472; &amp;#x1F3FF;</code><td><code>&amp;#128114; &amp;#127999;</code><td>
<tr><td>&#x1F64D;<td><td>PERSON FROWNING<td><code>&amp;#x1F64D;</code><td><code>&amp;#128589;</code><td>
<tr><td>&#x1F64D;&#x1F3FB;<td><td>person frowning, type-1-2<td><code>&amp;#x1F64D; &amp;#x1F3FB;</code><td><code>&amp;#128589; &amp;#127995;</code><td>
<tr><td>&#x1F64D;&#x1F3FC;<td><td>person frowning, type-3<td><code>&amp;#x1F64D; &amp;#x1F3FC;</code><td><code>&amp;#128589; &amp;#127996;</code><td>
<tr><td>&#x1F64D;&#x1F3FD;<td><td>person frowning, type-4<td><code>&amp;#x1F64D; &amp;#x1F3FD;</code><td><code>&amp;#128589; &amp;#127997;</code><td>
<tr><td>&#x1F64D;&#x1F3FE;<td><td>person frowning, type-5<td><code>&amp;#x1F64D; &amp;#x1F3FE;</code><td><code>&amp;#128589; &amp;#127998;</code><td>
<tr><td>&#x1F64D;&#x1F3FF;<td><td>person frowning, type-6<td><code>&amp;#x1F64D; &amp;#x1F3FF;</code><td><code>&amp;#128589; &amp;#127999;</code><td>
<tr><td>&#x1F64E;<td><td>PERSON WITH POUTING FACE &#x224A; person pouting<td><code>&amp;#x1F64E;</code><td><code>&amp;#128590;</code><td>
<tr><td>&#x1F64E;&#x1F3FB;<td><td>person pouting, type-1-2<td><code>&amp;#x1F64E; &amp;#x1F3FB;</code><td><code>&amp;#128590; &amp;#127995;</code><td>
<tr><td>&#x1F64E;&#x1F3FC;<td><td>person pouting, type-3<td><code>&amp;#x1F64E; &amp;#x1F3FC;</code><td><code>&amp;#128590; &amp;#127996;</code><td>
<tr><td>&#x1F64E;&#x1F3FD;<td><td>person pouting, type-4<td><code>&amp;#x1F64E; &amp;#x1F3FD;</code><td><code>&amp;#128590; &amp;#127997;</code><td>
<tr><td>&#x1F64E;&#x1F3FE;<td><td>person pouting, type-5<td><code>&amp;#x1F64E; &amp;#x1F3FE;</code><td><code>&amp;#128590; &amp;#127998;</code><td>
<tr><td>&#x1F64E;&#x1F3FF;<td><td>person pouting, type-6<td><code>&amp;#x1F64E; &amp;#x1F3FF;</code><td><code>&amp;#128590; &amp;#127999;</code><td>
<tr><td>&#x1F645;<td><td>FACE WITH NO GOOD GESTURE &#x224A; person gesturing not ok<td><code>&amp;#x1F645;</code><td><code>&amp;#128581;</code><td>
<tr><td>&#x1F645;&#x1F3FB;<td><td>person gesturing not ok, type-1-2<td><code>&amp;#x1F645; &amp;#x1F3FB;</code><td><code>&amp;#128581; &amp;#127995;</code><td>
<tr><td>&#x1F645;&#x1F3FC;<td><td>person gesturing not ok, type-3<td><code>&amp;#x1F645; &amp;#x1F3FC;</code><td><code>&amp;#128581; &amp;#127996;</code><td>
<tr><td>&#x1F645;&#x1F3FD;<td><td>person gesturing not ok, type-4<td><code>&amp;#x1F645; &amp;#x1F3FD;</code><td><code>&amp;#128581; &amp;#127997;</code><td>
<tr><td>&#x1F645;&#x1F3FE;<td><td>person gesturing not ok, type-5<td><code>&amp;#x1F645; &amp;#x1F3FE;</code><td><code>&amp;#128581; &amp;#127998;</code><td>
<tr><td>&#x1F645;&#x1F3FF;<td><td>person gesturing not ok, type-6<td><code>&amp;#x1F645; &amp;#x1F3FF;</code><td><code>&amp;#128581; &amp;#127999;</code><td>
<tr><td>&#x1F646;<td><td>FACE WITH OK GESTURE &#x224A; person gesturing ok<td><code>&amp;#x1F646;</code><td><code>&amp;#128582;</code><td>
<tr><td>&#x1F646;&#x1F3FB;<td><td>person gesturing ok, type-1-2<td><code>&amp;#x1F646; &amp;#x1F3FB;</code><td><code>&amp;#128582; &amp;#127995;</code><td>
<tr><td>&#x1F646;&#x1F3FC;<td><td>person gesturing ok, type-3<td><code>&amp;#x1F646; &amp;#x1F3FC;</code><td><code>&amp;#128582; &amp;#127996;</code><td>
<tr><td>&#x1F646;&#x1F3FD;<td><td>person gesturing ok, type-4<td><code>&amp;#x1F646; &amp;#x1F3FD;</code><td><code>&amp;#128582; &amp;#127997;</code><td>
<tr><td>&#x1F646;&#x1F3FE;<td><td>person gesturing ok, type-5<td><code>&amp;#x1F646; &amp;#x1F3FE;</code><td><code>&amp;#128582; &amp;#127998;</code><td>
<tr><td>&#x1F646;&#x1F3FF;<td><td>person gesturing ok, type-6<td><code>&amp;#x1F646; &amp;#x1F3FF;</code><td><code>&amp;#128582; &amp;#127999;</code><td>
<tr><td>&#x1F481;<td><td>INFORMATION DESK PERSON &#x224A; person tipping hand<td><code>&amp;#x1F481;</code><td><code>&amp;#128129;</code><td>
<tr><td>&#x1F481;&#x1F3FB;<td><td>person tipping hand, type-1-2<td><code>&amp;#x1F481; &amp;#x1F3FB;</code><td><code>&amp;#128129; &amp;#127995;</code><td>
<tr><td>&#x1F481;&#x1F3FC;<td><td>person tipping hand, type-3<td><code>&amp;#x1F481; &amp;#x1F3FC;</code><td><code>&amp;#128129; &amp;#127996;</code><td>
<tr><td>&#x1F481;&#x1F3FD;<td><td>person tipping hand, type-4<td><code>&amp;#x1F481; &amp;#x1F3FD;</code><td><code>&amp;#128129; &amp;#127997;</code><td>
<tr><td>&#x1F481;&#x1F3FE;<td><td>person tipping hand, type-5<td><code>&amp;#x1F481; &amp;#x1F3FE;</code><td><code>&amp;#128129; &amp;#127998;</code><td>
<tr><td>&#x1F481;&#x1F3FF;<td><td>person tipping hand, type-6<td><code>&amp;#x1F481; &amp;#x1F3FF;</code><td><code>&amp;#128129; &amp;#127999;</code><td>
<tr><td>&#x1F64B;<td><td>HAPPY PERSON RAISING ONE HAND &#x224A; person raising hand<td><code>&amp;#x1F64B;</code><td><code>&amp;#128587;</code><td>
<tr><td>&#x1F64B;&#x1F3FB;<td><td>person raising hand, type-1-2<td><code>&amp;#x1F64B; &amp;#x1F3FB;</code><td><code>&amp;#128587; &amp;#127995;</code><td>
<tr><td>&#x1F64B;&#x1F3FC;<td><td>person raising hand, type-3<td><code>&amp;#x1F64B; &amp;#x1F3FC;</code><td><code>&amp;#128587; &amp;#127996;</code><td>
<tr><td>&#x1F64B;&#x1F3FD;<td><td>person raising hand, type-4<td><code>&amp;#x1F64B; &amp;#x1F3FD;</code><td><code>&amp;#128587; &amp;#127997;</code><td>
<tr><td>&#x1F64B;&#x1F3FE;<td><td>person raising hand, type-5<td><code>&amp;#x1F64B; &amp;#x1F3FE;</code><td><code>&amp;#128587; &amp;#127998;</code><td>
<tr><td>&#x1F64B;&#x1F3FF;<td><td>person raising hand, type-6<td><code>&amp;#x1F64B; &amp;#x1F3FF;</code><td><code>&amp;#128587; &amp;#127999;</code><td>
<tr><td>&#x1F647;<td><td>PERSON BOWING DEEPLY &#x224A; person bowing<td><code>&amp;#x1F647;</code><td><code>&amp;#128583;</code><td>
<tr><td>&#x1F647;&#x1F3FB;<td><td>person bowing, type-1-2<td><code>&amp;#x1F647; &amp;#x1F3FB;</code><td><code>&amp;#128583; &amp;#127995;</code><td>
<tr><td>&#x1F647;&#x1F3FC;<td><td>person bowing, type-3<td><code>&amp;#x1F647; &amp;#x1F3FC;</code><td><code>&amp;#128583; &amp;#127996;</code><td>
<tr><td>&#x1F647;&#x1F3FD;<td><td>person bowing, type-4<td><code>&amp;#x1F647; &amp;#x1F3FD;</code><td><code>&amp;#128583; &amp;#127997;</code><td>
<tr><td>&#x1F647;&#x1F3FE;<td><td>person bowing, type-5<td><code>&amp;#x1F647; &amp;#x1F3FE;</code><td><code>&amp;#128583; &amp;#127998;</code><td>
<tr><td>&#x1F647;&#x1F3FF;<td><td>person bowing, type-6<td><code>&amp;#x1F647; &amp;#x1F3FF;</code><td><code>&amp;#128583; &amp;#127999;</code><td>
<tr><td>&#x1F926;<td><td>FACE PALM &#x224A; person facepalming<td><code>&amp;#x1F926;</code><td><code>&amp;#129318;</code><td>
<tr><td>&#x1F926;&#x1F3FB;<td><td>person facepalming, type-1-2<td><code>&amp;#x1F926; &amp;#x1F3FB;</code><td><code>&amp;#129318; &amp;#127995;</code><td>
<tr><td>&#x1F926;&#x1F3FC;<td><td>person facepalming, type-3<td><code>&amp;#x1F926; &amp;#x1F3FC;</code><td><code>&amp;#129318; &amp;#127996;</code><td>
<tr><td>&#x1F926;&#x1F3FD;<td><td>person facepalming, type-4<td><code>&amp;#x1F926; &amp;#x1F3FD;</code><td><code>&amp;#129318; &amp;#127997;</code><td>
<tr><td>&#x1F926;&#x1F3FE;<td><td>person facepalming, type-5<td><code>&amp;#x1F926; &amp;#x1F3FE;</code><td><code>&amp;#129318; &amp;#127998;</code><td>
<tr><td>&#x1F926;&#x1F3FF;<td><td>person facepalming, type-6<td><code>&amp;#x1F926; &amp;#x1F3FF;</code><td><code>&amp;#129318; &amp;#127999;</code><td>
<tr><td>&#x1F937;<td><td>SHRUG &#x224A; person shrugging<td><code>&amp;#x1F937;</code><td><code>&amp;#129335;</code><td>
<tr><td>&#x1F937;&#x1F3FB;<td><td>person shrugging, type-1-2<td><code>&amp;#x1F937; &amp;#x1F3FB;</code><td><code>&amp;#129335; &amp;#127995;</code><td>
<tr><td>&#x1F937;&#x1F3FC;<td><td>person shrugging, type-3<td><code>&amp;#x1F937; &amp;#x1F3FC;</code><td><code>&amp;#129335; &amp;#127996;</code><td>
<tr><td>&#x1F937;&#x1F3FD;<td><td>person shrugging, type-4<td><code>&amp;#x1F937; &amp;#x1F3FD;</code><td><code>&amp;#129335; &amp;#127997;</code><td>
<tr><td>&#x1F937;&#x1F3FE;<td><td>person shrugging, type-5<td><code>&amp;#x1F937; &amp;#x1F3FE;</code><td><code>&amp;#129335; &amp;#127998;</code><td>
<tr><td>&#x1F937;&#x1F3FF;<td><td>person shrugging, type-6<td><code>&amp;#x1F937; &amp;#x1F3FF;</code><td><code>&amp;#129335; &amp;#127999;</code><td>
<tr><td>&#x1F486;<td><td>FACE MASSAGE &#x224A; person getting massage<td><code>&amp;#x1F486;</code><td><code>&amp;#128134;</code><td>
<tr><td>&#x1F486;&#x1F3FB;<td><td>person getting massage, type-1-2<td><code>&amp;#x1F486; &amp;#x1F3FB;</code><td><code>&amp;#128134; &amp;#127995;</code><td>
<tr><td>&#x1F486;&#x1F3FC;<td><td>person getting massage, type-3<td><code>&amp;#x1F486; &amp;#x1F3FC;</code><td><code>&amp;#128134; &amp;#127996;</code><td>
<tr><td>&#x1F486;&#x1F3FD;<td><td>person getting massage, type-4<td><code>&amp;#x1F486; &amp;#x1F3FD;</code><td><code>&amp;#128134; &amp;#127997;</code><td>
<tr><td>&#x1F486;&#x1F3FE;<td><td>person getting massage, type-5<td><code>&amp;#x1F486; &amp;#x1F3FE;</code><td><code>&amp;#128134; &amp;#127998;</code><td>
<tr><td>&#x1F486;&#x1F3FF;<td><td>person getting massage, type-6<td><code>&amp;#x1F486; &amp;#x1F3FF;</code><td><code>&amp;#128134; &amp;#127999;</code><td>
<tr><td>&#x1F487;<td><td>HAIRCUT &#x224A; person getting haircut<td><code>&amp;#x1F487;</code><td><code>&amp;#128135;</code><td>
<tr><td>&#x1F487;&#x1F3FB;<td><td>person getting haircut, type-1-2<td><code>&amp;#x1F487; &amp;#x1F3FB;</code><td><code>&amp;#128135; &amp;#127995;</code><td>
<tr><td>&#x1F487;&#x1F3FC;<td><td>person getting haircut, type-3<td><code>&amp;#x1F487; &amp;#x1F3FC;</code><td><code>&amp;#128135; &amp;#127996;</code><td>
<tr><td>&#x1F487;&#x1F3FD;<td><td>person getting haircut, type-4<td><code>&amp;#x1F487; &amp;#x1F3FD;</code><td><code>&amp;#128135; &amp;#127997;</code><td>
<tr><td>&#x1F487;&#x1F3FE;<td><td>person getting haircut, type-5<td><code>&amp;#x1F487; &amp;#x1F3FE;</code><td><code>&amp;#128135; &amp;#127998;</code><td>
<tr><td>&#x1F487;&#x1F3FF;<td><td>person getting haircut, type-6<td><code>&amp;#x1F487; &amp;#x1F3FF;</code><td><code>&amp;#128135; &amp;#127999;</code><td>
<tr><td>&#x1F6B6;<td><td>PEDESTRIAN &#x224A; person walking<td><code>&amp;#x1F6B6;</code><td><code>&amp;#128694;</code><td>
<tr><td>&#x1F6B6;&#x1F3FB;<td><td>person walking, type-1-2<td><code>&amp;#x1F6B6; &amp;#x1F3FB;</code><td><code>&amp;#128694; &amp;#127995;</code><td>
<tr><td>&#x1F6B6;&#x1F3FC;<td><td>person walking, type-3<td><code>&amp;#x1F6B6; &amp;#x1F3FC;</code><td><code>&amp;#128694; &amp;#127996;</code><td>
<tr><td>&#x1F6B6;&#x1F3FD;<td><td>person walking, type-4<td><code>&amp;#x1F6B6; &amp;#x1F3FD;</code><td><code>&amp;#128694; &amp;#127997;</code><td>
<tr><td>&#x1F6B6;&#x1F3FE;<td><td>person walking, type-5<td><code>&amp;#x1F6B6; &amp;#x1F3FE;</code><td><code>&amp;#128694; &amp;#127998;</code><td>
<tr><td>&#x1F6B6;&#x1F3FF;<td><td>person walking, type-6<td><code>&amp;#x1F6B6; &amp;#x1F3FF;</code><td><code>&amp;#128694; &amp;#127999;</code><td>
<tr><td>&#x1F3C3;<td><td>RUNNER &#x224A; person running<td><code>&amp;#x1F3C3;</code><td><code>&amp;#127939;</code><td>
<tr><td>&#x1F3C3;&#x1F3FB;<td><td>person running, type-1-2<td><code>&amp;#x1F3C3; &amp;#x1F3FB;</code><td><code>&amp;#127939; &amp;#127995;</code><td>
<tr><td>&#x1F3C3;&#x1F3FC;<td><td>person running, type-3<td><code>&amp;#x1F3C3; &amp;#x1F3FC;</code><td><code>&amp;#127939; &amp;#127996;</code><td>
<tr><td>&#x1F3C3;&#x1F3FD;<td><td>person running, type-4<td><code>&amp;#x1F3C3; &amp;#x1F3FD;</code><td><code>&amp;#127939; &amp;#127997;</code><td>
<tr><td>&#x1F3C3;&#x1F3FE;<td><td>person running, type-5<td><code>&amp;#x1F3C3; &amp;#x1F3FE;</code><td><code>&amp;#127939; &amp;#127998;</code><td>
<tr><td>&#x1F3C3;&#x1F3FF;<td><td>person running, type-6<td><code>&amp;#x1F3C3; &amp;#x1F3FF;</code><td><code>&amp;#127939; &amp;#127999;</code><td>
<tr><td>&#x1F483;<td><td>DANCER &#x224A; woman dancing<td><code>&amp;#x1F483;</code><td><code>&amp;#128131;</code><td>
<tr><td>&#x1F483;&#x1F3FB;<td><td>woman dancing, type-1-2<td><code>&amp;#x1F483; &amp;#x1F3FB;</code><td><code>&amp;#128131; &amp;#127995;</code><td>
<tr><td>&#x1F483;&#x1F3FC;<td><td>woman dancing, type-3<td><code>&amp;#x1F483; &amp;#x1F3FC;</code><td><code>&amp;#128131; &amp;#127996;</code><td>
<tr><td>&#x1F483;&#x1F3FD;<td><td>woman dancing, type-4<td><code>&amp;#x1F483; &amp;#x1F3FD;</code><td><code>&amp;#128131; &amp;#127997;</code><td>
<tr><td>&#x1F483;&#x1F3FE;<td><td>woman dancing, type-5<td><code>&amp;#x1F483; &amp;#x1F3FE;</code><td><code>&amp;#128131; &amp;#127998;</code><td>
<tr><td>&#x1F483;&#x1F3FF;<td><td>woman dancing, type-6<td><code>&amp;#x1F483; &amp;#x1F3FF;</code><td><code>&amp;#128131; &amp;#127999;</code><td>
<tr><td>&#x1F57A;<td><td>MAN DANCING<td><code>&amp;#x1F57A;</code><td><code>&amp;#128378;</code><td>
<tr><td>&#x1F57A;&#x1F3FB;<td><td>man dancing, type-1-2<td><code>&amp;#x1F57A; &amp;#x1F3FB;</code><td><code>&amp;#128378; &amp;#127995;</code><td>
<tr><td>&#x1F57A;&#x1F3FC;<td><td>man dancing, type-3<td><code>&amp;#x1F57A; &amp;#x1F3FC;</code><td><code>&amp;#128378; &amp;#127996;</code><td>
<tr><td>&#x1F57A;&#x1F3FD;<td><td>man dancing, type-4<td><code>&amp;#x1F57A; &amp;#x1F3FD;</code><td><code>&amp;#128378; &amp;#127997;</code><td>
<tr><td>&#x1F57A;&#x1F3FE;<td><td>man dancing, type-5<td><code>&amp;#x1F57A; &amp;#x1F3FE;</code><td><code>&amp;#128378; &amp;#127998;</code><td>
<tr><td>&#x1F57A;&#x1F3FF;<td><td>man dancing, type-6<td><code>&amp;#x1F57A; &amp;#x1F3FF;</code><td><code>&amp;#128378; &amp;#127999;</code><td>
<tr><td>&#x1F46F;<td><td>WOMAN WITH BUNNY EARS &#x224A; people partying<td><code>&amp;#x1F46F;</code><td><code>&amp;#128111;</code><td>
<tr><td>&#x1F574;<td><td>MAN IN BUSINESS SUIT LEVITATING<td><code>&amp;#x1F574;</code><td><code>&amp;#128372;</code><td>
<tr><td>&#x1F5E3;<td><td>SPEAKING HEAD IN SILHOUETTE &#x224A; speaking head<td><code>&amp;#x1F5E3;</code><td><code>&amp;#128483;</code><td>
<tr><td>&#x1F464;<td><td>BUST IN SILHOUETTE<td><code>&amp;#x1F464;</code><td><code>&amp;#128100;</code><td>
<tr><td>&#x1F465;<td><td>BUSTS IN SILHOUETTE<td><code>&amp;#x1F465;</code><td><code>&amp;#128101;</code><td>
<tr><td>&#x1F46B;<td><td>MAN AND WOMAN HOLDING HANDS<td><code>&amp;#x1F46B;</code><td><code>&amp;#128107;</code><td>
<tr><td>&#x1F46C;<td><td>TWO MEN HOLDING HANDS<td><code>&amp;#x1F46C;</code><td><code>&amp;#128108;</code><td>
<tr><td>&#x1F46D;<td><td>TWO WOMEN HOLDING HANDS<td><code>&amp;#x1F46D;</code><td><code>&amp;#128109;</code><td>
<tr><td>&#x1F48F;<td><td>KISS<td><code>&amp;#x1F48F;</code><td><code>&amp;#128143;</code><td>
<tr><td>&#x1F469;&#x200D;&#x2764;&#xFE0F;&#x200D;&#x1F48B;&#x200D;&#x1F468;<td><td>kiss, woman, man<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x2764; &amp;#xFE0F; &amp;#x200D; &amp;#x1F48B; &amp;#x200D; &amp;#x1F468;</code><td><code>&amp;#128105; &amp;#8205; &amp;#10084; &amp;#65039; &amp;#8205; &amp;#128139; &amp;#8205; &amp;#128104;</code><td>
<tr><td>&#x1F468;&#x200D;&#x2764;&#xFE0F;&#x200D;&#x1F48B;&#x200D;&#x1F468;<td><td>kiss, man, man<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x2764; &amp;#xFE0F; &amp;#x200D; &amp;#x1F48B; &amp;#x200D; &amp;#x1F468;</code><td><code>&amp;#128104; &amp;#8205; &amp;#10084; &amp;#65039; &amp;#8205; &amp;#128139; &amp;#8205; &amp;#128104;</code><td>
<tr><td>&#x1F469;&#x200D;&#x2764;&#xFE0F;&#x200D;&#x1F48B;&#x200D;&#x1F469;<td><td>kiss, woman, woman<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x2764; &amp;#xFE0F; &amp;#x200D; &amp;#x1F48B; &amp;#x200D; &amp;#x1F469;</code><td><code>&amp;#128105; &amp;#8205; &amp;#10084; &amp;#65039; &amp;#8205; &amp;#128139; &amp;#8205; &amp;#128105;</code><td>
<tr><td>&#x1F491;<td><td>COUPLE WITH HEART<td><code>&amp;#x1F491;</code><td><code>&amp;#128145;</code><td>
<tr><td>&#x1F469;&#x200D;&#x2764;&#xFE0F;&#x200D;&#x1F468;  <td><td>couple with heart, woman, man<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x2764; &amp;#xFE0F; &amp;#x200D; &amp;#x1F468;  </code><td><code>&amp;#128105; &amp;#8205; &amp;#10084; &amp;#65039; &amp;#8205; &amp;#128104;  </code><td>
<tr><td>&#x1F468;&#x200D;&#x2764;&#xFE0F;&#x200D;&#x1F468;  <td><td>couple with heart, man, man<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x2764; &amp;#xFE0F; &amp;#x200D; &amp;#x1F468;  </code><td><code>&amp;#128104; &amp;#8205; &amp;#10084; &amp;#65039; &amp;#8205; &amp;#128104;  </code><td>
<tr><td>&#x1F469;&#x200D;&#x2764;&#xFE0F;&#x200D;&#x1F469;  <td><td>couple with heart, woman, woman<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x2764; &amp;#xFE0F; &amp;#x200D; &amp;#x1F469;  </code><td><code>&amp;#128105; &amp;#8205; &amp;#10084; &amp;#65039; &amp;#8205; &amp;#128105;  </code><td>
<tr><td>&#x1F46A;<td><td>FAMILY<td><code>&amp;#x1F46A;</code><td><code>&amp;#128106;</code><td>
<tr><td>&#x1F468;&#x200D;&#x1F469;&#x200D;&#x1F466;   <td><td>family, man, woman, boy<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F466;   </code><td><code>&amp;#128104; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128102;   </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F469;&#x200D;&#x1F467;   <td><td>family, man, woman, girl<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F467;   </code><td><code>&amp;#128104; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128103;   </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F469;&#x200D;&#x1F467;&#x200D;&#x1F466; <td><td>family, man, woman, girl, boy<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F467; &amp;#x200D; &amp;#x1F466; </code><td><code>&amp;#128104; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128103; &amp;#8205; &amp;#128102; </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F469;&#x200D;&#x1F466;&#x200D;&#x1F466; <td><td>family, man, woman, boy, boy<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F466; &amp;#x200D; &amp;#x1F466; </code><td><code>&amp;#128104; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128102; &amp;#8205; &amp;#128102; </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F469;&#x200D;&#x1F467;&#x200D;&#x1F467; <td><td>family, man, woman, girl, girl<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F467; &amp;#x200D; &amp;#x1F467; </code><td><code>&amp;#128104; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128103; &amp;#8205; &amp;#128103; </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F468;&#x200D;&#x1F466;   <td><td>family, man, man, boy<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F468; &amp;#x200D; &amp;#x1F466;   </code><td><code>&amp;#128104; &amp;#8205; &amp;#128104; &amp;#8205; &amp;#128102;   </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F468;&#x200D;&#x1F467;   <td><td>family, man, man, girl<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F468; &amp;#x200D; &amp;#x1F467;   </code><td><code>&amp;#128104; &amp;#8205; &amp;#128104; &amp;#8205; &amp;#128103;   </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F468;&#x200D;&#x1F467;&#x200D;&#x1F466; <td><td>family, man, man, girl, boy<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F468; &amp;#x200D; &amp;#x1F467; &amp;#x200D; &amp;#x1F466; </code><td><code>&amp;#128104; &amp;#8205; &amp;#128104; &amp;#8205; &amp;#128103; &amp;#8205; &amp;#128102; </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F468;&#x200D;&#x1F466;&#x200D;&#x1F466; <td><td>family, man, man, boy, boy<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F468; &amp;#x200D; &amp;#x1F466; &amp;#x200D; &amp;#x1F466; </code><td><code>&amp;#128104; &amp;#8205; &amp;#128104; &amp;#8205; &amp;#128102; &amp;#8205; &amp;#128102; </code><td>
<tr><td>&#x1F468;&#x200D;&#x1F468;&#x200D;&#x1F467;&#x200D;&#x1F467; <td><td>family, man, man, girl, girl<td><code>&amp;#x1F468; &amp;#x200D; &amp;#x1F468; &amp;#x200D; &amp;#x1F467; &amp;#x200D; &amp;#x1F467; </code><td><code>&amp;#128104; &amp;#8205; &amp;#128104; &amp;#8205; &amp;#128103; &amp;#8205; &amp;#128103; </code><td>
<tr><td>&#x1F469;&#x200D;&#x1F469;&#x200D;&#x1F466;   <td><td>family, woman, woman, boy<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F466;   </code><td><code>&amp;#128105; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128102;   </code><td>
<tr><td>&#x1F469;&#x200D;&#x1F469;&#x200D;&#x1F467;   <td><td>family, woman, woman, girl<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F467;   </code><td><code>&amp;#128105; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128103;   </code><td>
<tr><td>&#x1F469;&#x200D;&#x1F469;&#x200D;&#x1F467;&#x200D;&#x1F466; <td><td>family, woman, woman, girl, boy<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F467; &amp;#x200D; &amp;#x1F466; </code><td><code>&amp;#128105; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128103; &amp;#8205; &amp;#128102; </code><td>
<tr><td>&#x1F469;&#x200D;&#x1F469;&#x200D;&#x1F466;&#x200D;&#x1F466; <td><td>family, woman, woman, boy, boy<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F466; &amp;#x200D; &amp;#x1F466; </code><td><code>&amp;#128105; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128102; &amp;#8205; &amp;#128102; </code><td>
<tr><td>&#x1F469;&#x200D;&#x1F469;&#x200D;&#x1F467;&#x200D;&#x1F467; <td><td>family, woman, woman, girl, girl<td><code>&amp;#x1F469; &amp;#x200D; &amp;#x1F469; &amp;#x200D; &amp;#x1F467; &amp;#x200D; &amp;#x1F467; </code><td><code>&amp;#128105; &amp;#8205; &amp;#128105; &amp;#8205; &amp;#128103; &amp;#8205; &amp;#128103; </code><td>
<tr><td>&#x1F3FB;<td><td>EMOJI MODIFIER FITZPATRICK TYPE-1-2 &#x224A; type-1-2<td><code>&amp;#x1F3FB;</code><td><code>&amp;#127995;</code><td>
<tr><td>&#x1F3FC;<td><td>EMOJI MODIFIER FITZPATRICK TYPE-3 &#x224A; type-3<td><code>&amp;#x1F3FC;</code><td><code>&amp;#127996;</code><td>
<tr><td>&#x1F3FD;<td><td>EMOJI MODIFIER FITZPATRICK TYPE-4 &#x224A; type-4<td><code>&amp;#x1F3FD;</code><td><code>&amp;#127997;</code><td>
<tr><td>&#x1F3FE;<td><td>EMOJI MODIFIER FITZPATRICK TYPE-5 &#x224A; type-5<td><code>&amp;#x1F3FE;</code><td><code>&amp;#127998;</code><td>
<tr><td>&#x1F3FF;<td><td>EMOJI MODIFIER FITZPATRICK TYPE-6 &#x224A; type-6<td><code>&amp;#x1F3FF;</code><td><code>&amp;#127999;</code><td>
<tr><td>&#x1F4AA;<td><td>FLEXED BICEPS<td><code>&amp;#x1F4AA;</code><td><code>&amp;#128170;</code><td>
<tr><td>&#x1F4AA;&#x1F3FB;<td><td>flexed biceps, type-1-2<td><code>&amp;#x1F4AA; &amp;#x1F3FB;</code><td><code>&amp;#128170; &amp;#127995;</code><td>
<tr><td>&#x1F4AA;&#x1F3FC;<td><td>flexed biceps, type-3<td><code>&amp;#x1F4AA; &amp;#x1F3FC;</code><td><code>&amp;#128170; &amp;#127996;</code><td>
<tr><td>&#x1F4AA;&#x1F3FD;<td><td>flexed biceps, type-4<td><code>&amp;#x1F4AA; &amp;#x1F3FD;</code><td><code>&amp;#128170; &amp;#127997;</code><td>
<tr><td>&#x1F4AA;&#x1F3FE;<td><td>flexed biceps, type-5<td><code>&amp;#x1F4AA; &amp;#x1F3FE;</code><td><code>&amp;#128170; &amp;#127998;</code><td>
<tr><td>&#x1F4AA;&#x1F3FF;<td><td>flexed biceps, type-6<td><code>&amp;#x1F4AA; &amp;#x1F3FF;</code><td><code>&amp;#128170; &amp;#127999;</code><td>
<tr><td>&#x1F933;<td><td>SELFIE<td><code>&amp;#x1F933;</code><td><code>&amp;#129331;</code><td>
<tr><td>&#x1F933;&#x1F3FB;<td><td>selfie, type-1-2<td><code>&amp;#x1F933; &amp;#x1F3FB;</code><td><code>&amp;#129331; &amp;#127995;</code><td>
<tr><td>&#x1F933;&#x1F3FC;<td><td>selfie, type-3<td><code>&amp;#x1F933; &amp;#x1F3FC;</code><td><code>&amp;#129331; &amp;#127996;</code><td>
<tr><td>&#x1F933;&#x1F3FD;<td><td>selfie, type-4<td><code>&amp;#x1F933; &amp;#x1F3FD;</code><td><code>&amp;#129331; &amp;#127997;</code><td>
<tr><td>&#x1F933;&#x1F3FE;<td><td>selfie, type-5<td><code>&amp;#x1F933; &amp;#x1F3FE;</code><td><code>&amp;#129331; &amp;#127998;</code><td>
<tr><td>&#x1F933;&#x1F3FF;<td><td>selfie, type-6<td><code>&amp;#x1F933; &amp;#x1F3FF;</code><td><code>&amp;#129331; &amp;#127999;</code><td>
<tr><td>&#x1F448;<td><td>WHITE LEFT POINTING BACKHAND INDEX &#x224A; backhand index pointing left<td><code>&amp;#x1F448;</code><td><code>&amp;#128072;</code><td>
<tr><td>&#x1F448;&#x1F3FB;<td><td>backhand index pointing left, type-1-2<td><code>&amp;#x1F448; &amp;#x1F3FB;</code><td><code>&amp;#128072; &amp;#127995;</code><td>
<tr><td>&#x1F448;&#x1F3FC;<td><td>backhand index pointing left, type-3<td><code>&amp;#x1F448; &amp;#x1F3FC;</code><td><code>&amp;#128072; &amp;#127996;</code><td>
<tr><td>&#x1F448;&#x1F3FD;<td><td>backhand index pointing left, type-4<td><code>&amp;#x1F448; &amp;#x1F3FD;</code><td><code>&amp;#128072; &amp;#127997;</code><td>
<tr><td>&#x1F448;&#x1F3FE;<td><td>backhand index pointing left, type-5<td><code>&amp;#x1F448; &amp;#x1F3FE;</code><td><code>&amp;#128072; &amp;#127998;</code><td>
<tr><td>&#x1F448;&#x1F3FF;<td><td>backhand index pointing left, type-6<td><code>&amp;#x1F448; &amp;#x1F3FF;</code><td><code>&amp;#128072; &amp;#127999;</code><td>
<tr><td>&#x1F449;<td><td>WHITE RIGHT POINTING BACKHAND INDEX &#x224A; backhand index pointing right<td><code>&amp;#x1F449;</code><td><code>&amp;#128073;</code><td>
<tr><td>&#x1F449;&#x1F3FB;<td><td>backhand index pointing right, type-1-2<td><code>&amp;#x1F449; &amp;#x1F3FB;</code><td><code>&amp;#128073; &amp;#127995;</code><td>
<tr><td>&#x1F449;&#x1F3FC;<td><td>backhand index pointing right, type-3<td><code>&amp;#x1F449; &amp;#x1F3FC;</code><td><code>&amp;#128073; &amp;#127996;</code><td>
<tr><td>&#x1F449;&#x1F3FD;<td><td>backhand index pointing right, type-4<td><code>&amp;#x1F449; &amp;#x1F3FD;</code><td><code>&amp;#128073; &amp;#127997;</code><td>
<tr><td>&#x1F449;&#x1F3FE;<td><td>backhand index pointing right, type-5<td><code>&amp;#x1F449; &amp;#x1F3FE;</code><td><code>&amp;#128073; &amp;#127998;</code><td>
<tr><td>&#x1F449;&#x1F3FF;<td><td>backhand index pointing right, type-6<td><code>&amp;#x1F449; &amp;#x1F3FF;</code><td><code>&amp;#128073; &amp;#127999;</code><td>
<tr><td>&#x261D;<td><td>WHITE UP POINTING INDEX &#x224A; index pointing up<td><code>&amp;#x261D;</code><td><code>&amp;#9757;</code><td>
<tr><td>&#x261D;&#x1F3FB;<td><td>index pointing up, type-1-2<td><code>&amp;#x261D; &amp;#x1F3FB;</code><td><code>&amp;#9757; &amp;#127995;</code><td>
<tr><td>&#x261D;&#x1F3FC;<td><td>index pointing up, type-3<td><code>&amp;#x261D; &amp;#x1F3FC;</code><td><code>&amp;#9757; &amp;#127996;</code><td>
<tr><td>&#x261D;&#x1F3FD;<td><td>index pointing up, type-4<td><code>&amp;#x261D; &amp;#x1F3FD;</code><td><code>&amp;#9757; &amp;#127997;</code><td>
<tr><td>&#x261D;&#x1F3FE;<td><td>index pointing up, type-5<td><code>&amp;#x261D; &amp;#x1F3FE;</code><td><code>&amp;#9757; &amp;#127998;</code><td>
<tr><td>&#x261D;&#x1F3FF;<td><td>index pointing up, type-6<td><code>&amp;#x261D; &amp;#x1F3FF;</code><td><code>&amp;#9757; &amp;#127999;</code><td>
<tr><td>&#x1F446;<td><td>WHITE UP POINTING BACKHAND INDEX &#x224A; backhand index pointing up<td><code>&amp;#x1F446;</code><td><code>&amp;#128070;</code><td>
<tr><td>&#x1F446;&#x1F3FB;<td><td>backhand index pointing up, type-1-2<td><code>&amp;#x1F446; &amp;#x1F3FB;</code><td><code>&amp;#128070; &amp;#127995;</code><td>
<tr><td>&#x1F446;&#x1F3FC;<td><td>backhand index pointing up, type-3<td><code>&amp;#x1F446; &amp;#x1F3FC;</code><td><code>&amp;#128070; &amp;#127996;</code><td>
<tr><td>&#x1F446;&#x1F3FD;<td><td>backhand index pointing up, type-4<td><code>&amp;#x1F446; &amp;#x1F3FD;</code><td><code>&amp;#128070; &amp;#127997;</code><td>
<tr><td>&#x1F446;&#x1F3FE;<td><td>backhand index pointing up, type-5<td><code>&amp;#x1F446; &amp;#x1F3FE;</code><td><code>&amp;#128070; &amp;#127998;</code><td>
<tr><td>&#x1F446;&#x1F3FF;<td><td>backhand index pointing up, type-6<td><code>&amp;#x1F446; &amp;#x1F3FF;</code><td><code>&amp;#128070; &amp;#127999;</code><td>
<tr><td>&#x1F595;<td><td>REVERSED HAND WITH MIDDLE FINGER EXTENDED &#x224A; middle finger<td><code>&amp;#x1F595;</code><td><code>&amp;#128405;</code><td>
<tr><td>&#x1F595;&#x1F3FB;<td><td>middle finger, type-1-2<td><code>&amp;#x1F595; &amp;#x1F3FB;</code><td><code>&amp;#128405; &amp;#127995;</code><td>
<tr><td>&#x1F595;&#x1F3FC;<td><td>middle finger, type-3<td><code>&amp;#x1F595; &amp;#x1F3FC;</code><td><code>&amp;#128405; &amp;#127996;</code><td>
<tr><td>&#x1F595;&#x1F3FD;<td><td>middle finger, type-4<td><code>&amp;#x1F595; &amp;#x1F3FD;</code><td><code>&amp;#128405; &amp;#127997;</code><td>
<tr><td>&#x1F595;&#x1F3FE;<td><td>middle finger, type-5<td><code>&amp;#x1F595; &amp;#x1F3FE;</code><td><code>&amp;#128405; &amp;#127998;</code><td>
<tr><td>&#x1F595;&#x1F3FF;<td><td>middle finger, type-6<td><code>&amp;#x1F595; &amp;#x1F3FF;</code><td><code>&amp;#128405; &amp;#127999;</code><td>
<tr><td>&#x1F447;<td><td>WHITE DOWN POINTING BACKHAND INDEX &#x224A; backhand index pointing down<td><code>&amp;#x1F447;</code><td><code>&amp;#128071;</code><td>
<tr><td>&#x1F447;&#x1F3FB;<td><td>backhand index pointing down, type-1-2<td><code>&amp;#x1F447; &amp;#x1F3FB;</code><td><code>&amp;#128071; &amp;#127995;</code><td>
<tr><td>&#x1F447;&#x1F3FC;<td><td>backhand index pointing down, type-3<td><code>&amp;#x1F447; &amp;#x1F3FC;</code><td><code>&amp;#128071; &amp;#127996;</code><td>
<tr><td>&#x1F447;&#x1F3FD;<td><td>backhand index pointing down, type-4<td><code>&amp;#x1F447; &amp;#x1F3FD;</code><td><code>&amp;#128071; &amp;#127997;</code><td>
<tr><td>&#x1F447;&#x1F3FE;<td><td>backhand index pointing down, type-5<td><code>&amp;#x1F447; &amp;#x1F3FE;</code><td><code>&amp;#128071; &amp;#127998;</code><td>
<tr><td>&#x1F447;&#x1F3FF;<td><td>backhand index pointing down, type-6<td><code>&amp;#x1F447; &amp;#x1F3FF;</code><td><code>&amp;#128071; &amp;#127999;</code><td>
<tr><td>&#x270C;<td><td>VICTORY HAND<td><code>&amp;#x270C;</code><td><code>&amp;#9996;</code><td>
<tr><td>&#x270C;&#x1F3FB;<td><td>victory hand, type-1-2<td><code>&amp;#x270C; &amp;#x1F3FB;</code><td><code>&amp;#9996; &amp;#127995;</code><td>
<tr><td>&#x270C;&#x1F3FC;<td><td>victory hand, type-3<td><code>&amp;#x270C; &amp;#x1F3FC;</code><td><code>&amp;#9996; &amp;#127996;</code><td>
<tr><td>&#x270C;&#x1F3FD;<td><td>victory hand, type-4<td><code>&amp;#x270C; &amp;#x1F3FD;</code><td><code>&amp;#9996; &amp;#127997;</code><td>
<tr><td>&#x270C;&#x1F3FE;<td><td>victory hand, type-5<td><code>&amp;#x270C; &amp;#x1F3FE;</code><td><code>&amp;#9996; &amp;#127998;</code><td>
<tr><td>&#x270C;&#x1F3FF;<td><td>victory hand, type-6<td><code>&amp;#x270C; &amp;#x1F3FF;</code><td><code>&amp;#9996; &amp;#127999;</code><td>
<tr><td>&#x1F91E;<td><td>HAND WITH INDEX AND MIDDLE FINGERS CROSSED &#x224A; crossed fingers<td><code>&amp;#x1F91E;</code><td><code>&amp;#129310;</code><td>
<tr><td>&#x1F91E;&#x1F3FB;<td><td>crossed fingers, type-1-2<td><code>&amp;#x1F91E; &amp;#x1F3FB;</code><td><code>&amp;#129310; &amp;#127995;</code><td>
<tr><td>&#x1F91E;&#x1F3FC;<td><td>crossed fingers, type-3<td><code>&amp;#x1F91E; &amp;#x1F3FC;</code><td><code>&amp;#129310; &amp;#127996;</code><td>
<tr><td>&#x1F91E;&#x1F3FD;<td><td>crossed fingers, type-4<td><code>&amp;#x1F91E; &amp;#x1F3FD;</code><td><code>&amp;#129310; &amp;#127997;</code><td>
<tr><td>&#x1F91E;&#x1F3FE;<td><td>crossed fingers, type-5<td><code>&amp;#x1F91E; &amp;#x1F3FE;</code><td><code>&amp;#129310; &amp;#127998;</code><td>
<tr><td>&#x1F91E;&#x1F3FF;<td><td>crossed fingers, type-6<td><code>&amp;#x1F91E; &amp;#x1F3FF;</code><td><code>&amp;#129310; &amp;#127999;</code><td>
<tr><td>&#x1F596;<td><td>RAISED HAND WITH PART BETWEEN MIDDLE AND RING FINGERS &#x224A; vulcan salute<td><code>&amp;#x1F596;</code><td><code>&amp;#128406;</code><td>
<tr><td>&#x1F596;&#x1F3FB;<td><td>vulcan salute, type-1-2<td><code>&amp;#x1F596; &amp;#x1F3FB;</code><td><code>&amp;#128406; &amp;#127995;</code><td>
<tr><td>&#x1F596;&#x1F3FC;<td><td>vulcan salute, type-3<td><code>&amp;#x1F596; &amp;#x1F3FC;</code><td><code>&amp;#128406; &amp;#127996;</code><td>
<tr><td>&#x1F596;&#x1F3FD;<td><td>vulcan salute, type-4<td><code>&amp;#x1F596; &amp;#x1F3FD;</code><td><code>&amp;#128406; &amp;#127997;</code><td>
<tr><td>&#x1F596;&#x1F3FE;<td><td>vulcan salute, type-5<td><code>&amp;#x1F596; &amp;#x1F3FE;</code><td><code>&amp;#128406; &amp;#127998;</code><td>
<tr><td>&#x1F596;&#x1F3FF;<td><td>vulcan salute, type-6<td><code>&amp;#x1F596; &amp;#x1F3FF;</code><td><code>&amp;#128406; &amp;#127999;</code><td>
<tr><td>&#x1F918;<td><td>SIGN OF THE HORNS<td><code>&amp;#x1F918;</code><td><code>&amp;#129304;</code><td>
<tr><td>&#x1F918;&#x1F3FB;<td><td>sign of the horns, type-1-2<td><code>&amp;#x1F918; &amp;#x1F3FB;</code><td><code>&amp;#129304; &amp;#127995;</code><td>
<tr><td>&#x1F918;&#x1F3FC;<td><td>sign of the horns, type-3<td><code>&amp;#x1F918; &amp;#x1F3FC;</code><td><code>&amp;#129304; &amp;#127996;</code><td>
<tr><td>&#x1F918;&#x1F3FD;<td><td>sign of the horns, type-4<td><code>&amp;#x1F918; &amp;#x1F3FD;</code><td><code>&amp;#129304; &amp;#127997;</code><td>
<tr><td>&#x1F918;&#x1F3FE;<td><td>sign of the horns, type-5<td><code>&amp;#x1F918; &amp;#x1F3FE;</code><td><code>&amp;#129304; &amp;#127998;</code><td>
<tr><td>&#x1F918;&#x1F3FF;<td><td>sign of the horns, type-6<td><code>&amp;#x1F918; &amp;#x1F3FF;</code><td><code>&amp;#129304; &amp;#127999;</code><td>
<tr><td>&#x1F919;<td><td>CALL ME HAND<td><code>&amp;#x1F919;</code><td><code>&amp;#129305;</code><td>
<tr><td>&#x1F919;&#x1F3FB;<td><td>call me hand, type-1-2<td><code>&amp;#x1F919; &amp;#x1F3FB;</code><td><code>&amp;#129305; &amp;#127995;</code><td>
<tr><td>&#x1F919;&#x1F3FC;<td><td>call me hand, type-3<td><code>&amp;#x1F919; &amp;#x1F3FC;</code><td><code>&amp;#129305; &amp;#127996;</code><td>
<tr><td>&#x1F919;&#x1F3FD;<td><td>call me hand, type-4<td><code>&amp;#x1F919; &amp;#x1F3FD;</code><td><code>&amp;#129305; &amp;#127997;</code><td>
<tr><td>&#x1F919;&#x1F3FE;<td><td>call me hand, type-5<td><code>&amp;#x1F919; &amp;#x1F3FE;</code><td><code>&amp;#129305; &amp;#127998;</code><td>
<tr><td>&#x1F919;&#x1F3FF;<td><td>call me hand, type-6<td><code>&amp;#x1F919; &amp;#x1F3FF;</code><td><code>&amp;#129305; &amp;#127999;</code><td>
<tr><td>&#x1F590;<td><td>RAISED HAND WITH FINGERS SPLAYED<td><code>&amp;#x1F590;</code><td><code>&amp;#128400;</code><td>
<tr><td>&#x1F590;&#x1F3FB;<td><td>raised hand with fingers splayed, type-1-2<td><code>&amp;#x1F590; &amp;#x1F3FB;</code><td><code>&amp;#128400; &amp;#127995;</code><td>
<tr><td>&#x1F590;&#x1F3FC;<td><td>raised hand with fingers splayed, type-3<td><code>&amp;#x1F590; &amp;#x1F3FC;</code><td><code>&amp;#128400; &amp;#127996;</code><td>
<tr><td>&#x1F590;&#x1F3FD;<td><td>raised hand with fingers splayed, type-4<td><code>&amp;#x1F590; &amp;#x1F3FD;</code><td><code>&amp;#128400; &amp;#127997;</code><td>
<tr><td>&#x1F590;&#x1F3FE;<td><td>raised hand with fingers splayed, type-5<td><code>&amp;#x1F590; &amp;#x1F3FE;</code><td><code>&amp;#128400; &amp;#127998;</code><td>
<tr><td>&#x1F590;&#x1F3FF;<td><td>raised hand with fingers splayed, type-6<td><code>&amp;#x1F590; &amp;#x1F3FF;</code><td><code>&amp;#128400; &amp;#127999;</code><td>
<tr><td>&#x270B;<td><td>RAISED HAND<td><code>&amp;#x270B;</code><td><code>&amp;#9995;</code><td>
<tr><td>&#x270B;&#x1F3FB;<td><td>raised hand, type-1-2<td><code>&amp;#x270B; &amp;#x1F3FB;</code><td><code>&amp;#9995; &amp;#127995;</code><td>
<tr><td>&#x270B;&#x1F3FC;<td><td>raised hand, type-3<td><code>&amp;#x270B; &amp;#x1F3FC;</code><td><code>&amp;#9995; &amp;#127996;</code><td>
<tr><td>&#x270B;&#x1F3FD;<td><td>raised hand, type-4<td><code>&amp;#x270B; &amp;#x1F3FD;</code><td><code>&amp;#9995; &amp;#127997;</code><td>
<tr><td>&#x270B;&#x1F3FE;<td><td>raised hand, type-5<td><code>&amp;#x270B; &amp;#x1F3FE;</code><td><code>&amp;#9995; &amp;#127998;</code><td>
<tr><td>&#x270B;&#x1F3FF;<td><td>raised hand, type-6<td><code>&amp;#x270B; &amp;#x1F3FF;</code><td><code>&amp;#9995; &amp;#127999;</code><td>
<tr><td>&#x1F44C;<td><td>OK HAND SIGN &#x224A; ok hand<td><code>&amp;#x1F44C;</code><td><code>&amp;#128076;</code><td>
<tr><td>&#x1F44C;&#x1F3FB;<td><td>ok hand, type-1-2<td><code>&amp;#x1F44C; &amp;#x1F3FB;</code><td><code>&amp;#128076; &amp;#127995;</code><td>
<tr><td>&#x1F44C;&#x1F3FC;<td><td>ok hand, type-3<td><code>&amp;#x1F44C; &amp;#x1F3FC;</code><td><code>&amp;#128076; &amp;#127996;</code><td>
<tr><td>&#x1F44C;&#x1F3FD;<td><td>ok hand, type-4<td><code>&amp;#x1F44C; &amp;#x1F3FD;</code><td><code>&amp;#128076; &amp;#127997;</code><td>
<tr><td>&#x1F44C;&#x1F3FE;<td><td>ok hand, type-5<td><code>&amp;#x1F44C; &amp;#x1F3FE;</code><td><code>&amp;#128076; &amp;#127998;</code><td>
<tr><td>&#x1F44C;&#x1F3FF;<td><td>ok hand, type-6<td><code>&amp;#x1F44C; &amp;#x1F3FF;</code><td><code>&amp;#128076; &amp;#127999;</code><td>
<tr><td>&#x1F44D;<td><td>THUMBS UP SIGN &#x224A; thumbs up<td><code>&amp;#x1F44D;</code><td><code>&amp;#128077;</code><td>
<tr><td>&#x1F44D;&#x1F3FB;<td><td>thumbs up, type-1-2<td><code>&amp;#x1F44D; &amp;#x1F3FB;</code><td><code>&amp;#128077; &amp;#127995;</code><td>
<tr><td>&#x1F44D;&#x1F3FC;<td><td>thumbs up, type-3<td><code>&amp;#x1F44D; &amp;#x1F3FC;</code><td><code>&amp;#128077; &amp;#127996;</code><td>
<tr><td>&#x1F44D;&#x1F3FD;<td><td>thumbs up, type-4<td><code>&amp;#x1F44D; &amp;#x1F3FD;</code><td><code>&amp;#128077; &amp;#127997;</code><td>
<tr><td>&#x1F44D;&#x1F3FE;<td><td>thumbs up, type-5<td><code>&amp;#x1F44D; &amp;#x1F3FE;</code><td><code>&amp;#128077; &amp;#127998;</code><td>
<tr><td>&#x1F44D;&#x1F3FF;<td><td>thumbs up, type-6<td><code>&amp;#x1F44D; &amp;#x1F3FF;</code><td><code>&amp;#128077; &amp;#127999;</code><td>
<tr><td>&#x1F44E;<td><td>THUMBS DOWN SIGN &#x224A; thumbs down<td><code>&amp;#x1F44E;</code><td><code>&amp;#128078;</code><td>
<tr><td>&#x1F44E;&#x1F3FB;<td><td>thumbs down, type-1-2<td><code>&amp;#x1F44E; &amp;#x1F3FB;</code><td><code>&amp;#128078; &amp;#127995;</code><td>
<tr><td>&#x1F44E;&#x1F3FC;<td><td>thumbs down, type-3<td><code>&amp;#x1F44E; &amp;#x1F3FC;</code><td><code>&amp;#128078; &amp;#127996;</code><td>
<tr><td>&#x1F44E;&#x1F3FD;<td><td>thumbs down, type-4<td><code>&amp;#x1F44E; &amp;#x1F3FD;</code><td><code>&amp;#128078; &amp;#127997;</code><td>
<tr><td>&#x1F44E;&#x1F3FE;<td><td>thumbs down, type-5<td><code>&amp;#x1F44E; &amp;#x1F3FE;</code><td><code>&amp;#128078; &amp;#127998;</code><td>
<tr><td>&#x1F44E;&#x1F3FF;<td><td>thumbs down, type-6<td><code>&amp;#x1F44E; &amp;#x1F3FF;</code><td><code>&amp;#128078; &amp;#127999;</code><td>
<tr><td>&#x270A;<td><td>RAISED FIST<td><code>&amp;#x270A;</code><td><code>&amp;#9994;</code><td>
<tr><td>&#x270A;&#x1F3FB;<td><td>raised fist, type-1-2<td><code>&amp;#x270A; &amp;#x1F3FB;</code><td><code>&amp;#9994; &amp;#127995;</code><td>
<tr><td>&#x270A;&#x1F3FC;<td><td>raised fist, type-3<td><code>&amp;#x270A; &amp;#x1F3FC;</code><td><code>&amp;#9994; &amp;#127996;</code><td>
<tr><td>&#x270A;&#x1F3FD;<td><td>raised fist, type-4<td><code>&amp;#x270A; &amp;#x1F3FD;</code><td><code>&amp;#9994; &amp;#127997;</code><td>
<tr><td>&#x270A;&#x1F3FE;<td><td>raised fist, type-5<td><code>&amp;#x270A; &amp;#x1F3FE;</code><td><code>&amp;#9994; &amp;#127998;</code><td>
<tr><td>&#x270A;&#x1F3FF;<td><td>raised fist, type-6<td><code>&amp;#x270A; &amp;#x1F3FF;</code><td><code>&amp;#9994; &amp;#127999;</code><td>
<tr><td>&#x1F44A;<td><td>FISTED HAND SIGN &#x224A; oncoming fist<td><code>&amp;#x1F44A;</code><td><code>&amp;#128074;</code><td>
<tr><td>&#x1F44A;&#x1F3FB;<td><td>oncoming fist, type-1-2<td><code>&amp;#x1F44A; &amp;#x1F3FB;</code><td><code>&amp;#128074; &amp;#127995;</code><td>
<tr><td>&#x1F44A;&#x1F3FC;<td><td>oncoming fist, type-3<td><code>&amp;#x1F44A; &amp;#x1F3FC;</code><td><code>&amp;#128074; &amp;#127996;</code><td>
<tr><td>&#x1F44A;&#x1F3FD;<td><td>oncoming fist, type-4<td><code>&amp;#x1F44A; &amp;#x1F3FD;</code><td><code>&amp;#128074; &amp;#127997;</code><td>
<tr><td>&#x1F44A;&#x1F3FE;<td><td>oncoming fist, type-5<td><code>&amp;#x1F44A; &amp;#x1F3FE;</code><td><code>&amp;#128074; &amp;#127998;</code><td>
<tr><td>&#x1F44A;&#x1F3FF;<td><td>oncoming fist, type-6<td><code>&amp;#x1F44A; &amp;#x1F3FF;</code><td><code>&amp;#128074; &amp;#127999;</code><td>
<tr><td>&#x1F91B;<td><td>LEFT-FACING FIST<td><code>&amp;#x1F91B;</code><td><code>&amp;#129307;</code><td>
<tr><td>&#x1F91B;&#x1F3FB;<td><td>left-facing fist, type-1-2<td><code>&amp;#x1F91B; &amp;#x1F3FB;</code><td><code>&amp;#129307; &amp;#127995;</code><td>
<tr><td>&#x1F91B;&#x1F3FC;<td><td>left-facing fist, type-3<td><code>&amp;#x1F91B; &amp;#x1F3FC;</code><td><code>&amp;#129307; &amp;#127996;</code><td>
<tr><td>&#x1F91B;&#x1F3FD;<td><td>left-facing fist, type-4<td><code>&amp;#x1F91B; &amp;#x1F3FD;</code><td><code>&amp;#129307; &amp;#127997;</code><td>
<tr><td>&#x1F91B;&#x1F3FE;<td><td>left-facing fist, type-5<td><code>&amp;#x1F91B; &amp;#x1F3FE;</code><td><code>&amp;#129307; &amp;#127998;</code><td>
<tr><td>&#x1F91B;&#x1F3FF;<td><td>left-facing fist, type-6<td><code>&amp;#x1F91B; &amp;#x1F3FF;</code><td><code>&amp;#129307; &amp;#127999;</code><td>
<tr><td>&#x1F91C;<td><td>RIGHT-FACING FIST<td><code>&amp;#x1F91C;</code><td><code>&amp;#129308;</code><td>
<tr><td>&#x1F91C;&#x1F3FB;<td><td>right-facing fist, type-1-2<td><code>&amp;#x1F91C; &amp;#x1F3FB;</code><td><code>&amp;#129308; &amp;#127995;</code><td>
<tr><td>&#x1F91C;&#x1F3FC;<td><td>right-facing fist, type-3<td><code>&amp;#x1F91C; &amp;#x1F3FC;</code><td><code>&amp;#129308; &amp;#127996;</code><td>
<tr><td>&#x1F91C;&#x1F3FD;<td><td>right-facing fist, type-4<td><code>&amp;#x1F91C; &amp;#x1F3FD;</code><td><code>&amp;#129308; &amp;#127997;</code><td>
<tr><td>&#x1F91C;&#x1F3FE;<td><td>right-facing fist, type-5<td><code>&amp;#x1F91C; &amp;#x1F3FE;</code><td><code>&amp;#129308; &amp;#127998;</code><td>
<tr><td>&#x1F91C;&#x1F3FF;<td><td>right-facing fist, type-6<td><code>&amp;#x1F91C; &amp;#x1F3FF;</code><td><code>&amp;#129308; &amp;#127999;</code><td>
<tr><td>&#x1F91A;<td><td>RAISED BACK OF HAND<td><code>&amp;#x1F91A;</code><td><code>&amp;#129306;</code><td>
<tr><td>&#x1F91A;&#x1F3FB;<td><td>raised back of hand, type-1-2<td><code>&amp;#x1F91A; &amp;#x1F3FB;</code><td><code>&amp;#129306; &amp;#127995;</code><td>
<tr><td>&#x1F91A;&#x1F3FC;<td><td>raised back of hand, type-3<td><code>&amp;#x1F91A; &amp;#x1F3FC;</code><td><code>&amp;#129306; &amp;#127996;</code><td>
<tr><td>&#x1F91A;&#x1F3FD;<td><td>raised back of hand, type-4<td><code>&amp;#x1F91A; &amp;#x1F3FD;</code><td><code>&amp;#129306; &amp;#127997;</code><td>
<tr><td>&#x1F91A;&#x1F3FE;<td><td>raised back of hand, type-5<td><code>&amp;#x1F91A; &amp;#x1F3FE;</code><td><code>&amp;#129306; &amp;#127998;</code><td>
<tr><td>&#x1F91A;&#x1F3FF;<td><td>raised back of hand, type-6<td><code>&amp;#x1F91A; &amp;#x1F3FF;</code><td><code>&amp;#129306; &amp;#127999;</code><td>
<tr><td>&#x1F44B;<td><td>WAVING HAND SIGN &#x224A; waving hand<td><code>&amp;#x1F44B;</code><td><code>&amp;#128075;</code><td>
<tr><td>&#x1F44B;&#x1F3FB;<td><td>waving hand, type-1-2<td><code>&amp;#x1F44B; &amp;#x1F3FB;</code><td><code>&amp;#128075; &amp;#127995;</code><td>
<tr><td>&#x1F44B;&#x1F3FC;<td><td>waving hand, type-3<td><code>&amp;#x1F44B; &amp;#x1F3FC;</code><td><code>&amp;#128075; &amp;#127996;</code><td>
<tr><td>&#x1F44B;&#x1F3FD;<td><td>waving hand, type-4<td><code>&amp;#x1F44B; &amp;#x1F3FD;</code><td><code>&amp;#128075; &amp;#127997;</code><td>
<tr><td>&#x1F44B;&#x1F3FE;<td><td>waving hand, type-5<td><code>&amp;#x1F44B; &amp;#x1F3FE;</code><td><code>&amp;#128075; &amp;#127998;</code><td>
<tr><td>&#x1F44B;&#x1F3FF;<td><td>waving hand, type-6<td><code>&amp;#x1F44B; &amp;#x1F3FF;</code><td><code>&amp;#128075; &amp;#127999;</code><td>
<tr><td>&#x1F44F;<td><td>CLAPPING HANDS SIGN &#x224A; clapping hands<td><code>&amp;#x1F44F;</code><td><code>&amp;#128079;</code><td>
<tr><td>&#x1F44F;&#x1F3FB;<td><td>clapping hands, type-1-2<td><code>&amp;#x1F44F; &amp;#x1F3FB;</code><td><code>&amp;#128079; &amp;#127995;</code><td>
<tr><td>&#x1F44F;&#x1F3FC;<td><td>clapping hands, type-3<td><code>&amp;#x1F44F; &amp;#x1F3FC;</code><td><code>&amp;#128079; &amp;#127996;</code><td>
<tr><td>&#x1F44F;&#x1F3FD;<td><td>clapping hands, type-4<td><code>&amp;#x1F44F; &amp;#x1F3FD;</code><td><code>&amp;#128079; &amp;#127997;</code><td>
<tr><td>&#x1F44F;&#x1F3FE;<td><td>clapping hands, type-5<td><code>&amp;#x1F44F; &amp;#x1F3FE;</code><td><code>&amp;#128079; &amp;#127998;</code><td>
<tr><td>&#x1F44F;&#x1F3FF;<td><td>clapping hands, type-6<td><code>&amp;#x1F44F; &amp;#x1F3FF;</code><td><code>&amp;#128079; &amp;#127999;</code><td>
<tr><td>&#x270D;<td><td>WRITING HAND<td><code>&amp;#x270D;</code><td><code>&amp;#9997;</code><td>
<tr><td>&#x270D;&#x1F3FB;<td><td>writing hand, type-1-2<td><code>&amp;#x270D; &amp;#x1F3FB;</code><td><code>&amp;#9997; &amp;#127995;</code><td>
<tr><td>&#x270D;&#x1F3FC;<td><td>writing hand, type-3<td><code>&amp;#x270D; &amp;#x1F3FC;</code><td><code>&amp;#9997; &amp;#127996;</code><td>
<tr><td>&#x270D;&#x1F3FD;<td><td>writing hand, type-4<td><code>&amp;#x270D; &amp;#x1F3FD;</code><td><code>&amp;#9997; &amp;#127997;</code><td>
<tr><td>&#x270D;&#x1F3FE;<td><td>writing hand, type-5<td><code>&amp;#x270D; &amp;#x1F3FE;</code><td><code>&amp;#9997; &amp;#127998;</code><td>
<tr><td>&#x270D;&#x1F3FF;<td><td>writing hand, type-6<td><code>&amp;#x270D; &amp;#x1F3FF;</code><td><code>&amp;#9997; &amp;#127999;</code><td>
<tr><td>&#x1F450;<td><td>OPEN HANDS SIGN &#x224A; open hands<td><code>&amp;#x1F450;</code><td><code>&amp;#128080;</code><td>
<tr><td>&#x1F450;&#x1F3FB;<td><td>open hands, type-1-2<td><code>&amp;#x1F450; &amp;#x1F3FB;</code><td><code>&amp;#128080; &amp;#127995;</code><td>
<tr><td>&#x1F450;&#x1F3FC;<td><td>open hands, type-3<td><code>&amp;#x1F450; &amp;#x1F3FC;</code><td><code>&amp;#128080; &amp;#127996;</code><td>
<tr><td>&#x1F450;&#x1F3FD;<td><td>open hands, type-4<td><code>&amp;#x1F450; &amp;#x1F3FD;</code><td><code>&amp;#128080; &amp;#127997;</code><td>
<tr><td>&#x1F450;&#x1F3FE;<td><td>open hands, type-5<td><code>&amp;#x1F450; &amp;#x1F3FE;</code><td><code>&amp;#128080; &amp;#127998;</code><td>
<tr><td>&#x1F450;&#x1F3FF;<td><td>open hands, type-6<td><code>&amp;#x1F450; &amp;#x1F3FF;</code><td><code>&amp;#128080; &amp;#127999;</code><td>
<tr><td>&#x1F64C;<td><td>PERSON RAISING BOTH HANDS IN CELEBRATION &#x224A; person raising hands<td><code>&amp;#x1F64C;</code><td><code>&amp;#128588;</code><td>
<tr><td>&#x1F64C;&#x1F3FB;<td><td>person raising hands, type-1-2<td><code>&amp;#x1F64C; &amp;#x1F3FB;</code><td><code>&amp;#128588; &amp;#127995;</code><td>
<tr><td>&#x1F64C;&#x1F3FC;<td><td>person raising hands, type-3<td><code>&amp;#x1F64C; &amp;#x1F3FC;</code><td><code>&amp;#128588; &amp;#127996;</code><td>
<tr><td>&#x1F64C;&#x1F3FD;<td><td>person raising hands, type-4<td><code>&amp;#x1F64C; &amp;#x1F3FD;</code><td><code>&amp;#128588; &amp;#127997;</code><td>
<tr><td>&#x1F64C;&#x1F3FE;<td><td>person raising hands, type-5<td><code>&amp;#x1F64C; &amp;#x1F3FE;</code><td><code>&amp;#128588; &amp;#127998;</code><td>
<tr><td>&#x1F64C;&#x1F3FF;<td><td>person raising hands, type-6<td><code>&amp;#x1F64C; &amp;#x1F3FF;</code><td><code>&amp;#128588; &amp;#127999;</code><td>
<tr><td>&#x1F64F;<td><td>PERSON WITH FOLDED HANDS &#x224A; folded hands<td><code>&amp;#x1F64F;</code><td><code>&amp;#128591;</code><td>
<tr><td>&#x1F64F;&#x1F3FB;<td><td>folded hands, type-1-2<td><code>&amp;#x1F64F; &amp;#x1F3FB;</code><td><code>&amp;#128591; &amp;#127995;</code><td>
<tr><td>&#x1F64F;&#x1F3FC;<td><td>folded hands, type-3<td><code>&amp;#x1F64F; &amp;#x1F3FC;</code><td><code>&amp;#128591; &amp;#127996;</code><td>
<tr><td>&#x1F64F;&#x1F3FD;<td><td>folded hands, type-4<td><code>&amp;#x1F64F; &amp;#x1F3FD;</code><td><code>&amp;#128591; &amp;#127997;</code><td>
<tr><td>&#x1F64F;&#x1F3FE;<td><td>folded hands, type-5<td><code>&amp;#x1F64F; &amp;#x1F3FE;</code><td><code>&amp;#128591; &amp;#127998;</code><td>
<tr><td>&#x1F64F;&#x1F3FF;<td><td>folded hands, type-6<td><code>&amp;#x1F64F; &amp;#x1F3FF;</code><td><code>&amp;#128591; &amp;#127999;</code><td>
<tr><td>&#x1F91D;<td><td>HANDSHAKE<td><code>&amp;#x1F91D;</code><td><code>&amp;#129309;</code><td>
<tr><td>&#x1F91D;&#x1F3FB;<td><td>handshake, type-1-2<td><code>&amp;#x1F91D; &amp;#x1F3FB;</code><td><code>&amp;#129309; &amp;#127995;</code><td>
<tr><td>&#x1F91D;&#x1F3FC;<td><td>handshake, type-3<td><code>&amp;#x1F91D; &amp;#x1F3FC;</code><td><code>&amp;#129309; &amp;#127996;</code><td>
<tr><td>&#x1F91D;&#x1F3FD;<td><td>handshake, type-4<td><code>&amp;#x1F91D; &amp;#x1F3FD;</code><td><code>&amp;#129309; &amp;#127997;</code><td>
<tr><td>&#x1F91D;&#x1F3FE;<td><td>handshake, type-5<td><code>&amp;#x1F91D; &amp;#x1F3FE;</code><td><code>&amp;#129309; &amp;#127998;</code><td>
<tr><td>&#x1F91D;&#x1F3FF;<td><td>handshake, type-6<td><code>&amp;#x1F91D; &amp;#x1F3FF;</code><td><code>&amp;#129309; &amp;#127999;</code><td>
<tr><td>&#x1F485;<td><td>NAIL POLISH<td><code>&amp;#x1F485;</code><td><code>&amp;#128133;</code><td>
<tr><td>&#x1F485;&#x1F3FB;<td><td>nail polish, type-1-2<td><code>&amp;#x1F485; &amp;#x1F3FB;</code><td><code>&amp;#128133; &amp;#127995;</code><td>
<tr><td>&#x1F485;&#x1F3FC;<td><td>nail polish, type-3<td><code>&amp;#x1F485; &amp;#x1F3FC;</code><td><code>&amp;#128133; &amp;#127996;</code><td>
<tr><td>&#x1F485;&#x1F3FD;<td><td>nail polish, type-4<td><code>&amp;#x1F485; &amp;#x1F3FD;</code><td><code>&amp;#128133; &amp;#127997;</code><td>
<tr><td>&#x1F485;&#x1F3FE;<td><td>nail polish, type-5<td><code>&amp;#x1F485; &amp;#x1F3FE;</code><td><code>&amp;#128133; &amp;#127998;</code><td>
<tr><td>&#x1F485;&#x1F3FF;<td><td>nail polish, type-6<td><code>&amp;#x1F485; &amp;#x1F3FF;</code><td><code>&amp;#128133; &amp;#127999;</code><td>
<tr><td>&#x1F442;<td><td>EAR<td><code>&amp;#x1F442;</code><td><code>&amp;#128066;</code><td>
<tr><td>&#x1F442;&#x1F3FB;<td><td>ear, type-1-2<td><code>&amp;#x1F442; &amp;#x1F3FB;</code><td><code>&amp;#128066; &amp;#127995;</code><td>
<tr><td>&#x1F442;&#x1F3FC;<td><td>ear, type-3<td><code>&amp;#x1F442; &amp;#x1F3FC;</code><td><code>&amp;#128066; &amp;#127996;</code><td>
<tr><td>&#x1F442;&#x1F3FD;<td><td>ear, type-4<td><code>&amp;#x1F442; &amp;#x1F3FD;</code><td><code>&amp;#128066; &amp;#127997;</code><td>
<tr><td>&#x1F442;&#x1F3FE;<td><td>ear, type-5<td><code>&amp;#x1F442; &amp;#x1F3FE;</code><td><code>&amp;#128066; &amp;#127998;</code><td>
<tr><td>&#x1F442;&#x1F3FF;<td><td>ear, type-6<td><code>&amp;#x1F442; &amp;#x1F3FF;</code><td><code>&amp;#128066; &amp;#127999;</code><td>
<tr><td>&#x1F443;<td><td>NOSE<td><code>&amp;#x1F443;</code><td><code>&amp;#128067;</code><td>
<tr><td>&#x1F443;&#x1F3FB;<td><td>nose, type-1-2<td><code>&amp;#x1F443; &amp;#x1F3FB;</code><td><code>&amp;#128067; &amp;#127995;</code><td>
<tr><td>&#x1F443;&#x1F3FC;<td><td>nose, type-3<td><code>&amp;#x1F443; &amp;#x1F3FC;</code><td><code>&amp;#128067; &amp;#127996;</code><td>
<tr><td>&#x1F443;&#x1F3FD;<td><td>nose, type-4<td><code>&amp;#x1F443; &amp;#x1F3FD;</code><td><code>&amp;#128067; &amp;#127997;</code><td>
<tr><td>&#x1F443;&#x1F3FE;<td><td>nose, type-5<td><code>&amp;#x1F443; &amp;#x1F3FE;</code><td><code>&amp;#128067; &amp;#127998;</code><td>
<tr><td>&#x1F443;&#x1F3FF;<td><td>nose, type-6<td><code>&amp;#x1F443; &amp;#x1F3FF;</code><td><code>&amp;#128067; &amp;#127999;</code><td>
<tr><td>&#x1F463;<td><td>FOOTPRINTS<td><code>&amp;#x1F463;</code><td><code>&amp;#128099;</code><td>
<tr><td>&#x1F440;<td><td>EYES<td><code>&amp;#x1F440;</code><td><code>&amp;#128064;</code><td>
<tr><td>&#x1F441;<td><td>EYE<td><code>&amp;#x1F441;</code><td><code>&amp;#128065;</code><td>
<tr><td>&#x1F441;&#x200D;&#x1F5E8;     <td><td>eye, left speech bubble<td><code>&amp;#x1F441; &amp;#x200D; &amp;#x1F5E8;     </code><td><code>&amp;#128065; &amp;#8205; &amp;#128488;     </code><td>
<tr><td>&#x1F445;<td><td>TONGUE<td><code>&amp;#x1F445;</code><td><code>&amp;#128069;</code><td>
<tr><td>&#x1F444;<td><td>MOUTH<td><code>&amp;#x1F444;</code><td><code>&amp;#128068;</code><td>
<tr><td>&#x1F48B;<td><td>KISS MARK<td><code>&amp;#x1F48B;</code><td><code>&amp;#128139;</code><td>
<tr><td>&#x1F498;<td><td>HEART WITH ARROW<td><code>&amp;#x1F498;</code><td><code>&amp;#128152;</code><td>
<tr><td>&#x2764;<td><td>HEAVY BLACK HEART &#x224A; red heart<td><code>&amp;#x2764;</code><td><code>&amp;#10084;</code><td>
<tr><td>&#x1F493;<td><td>BEATING HEART<td><code>&amp;#x1F493;</code><td><code>&amp;#128147;</code><td>
<tr><td>&#x1F494;<td><td>BROKEN HEART<td><code>&amp;#x1F494;</code><td><code>&amp;#128148;</code><td>
<tr><td>&#x1F495;<td><td>TWO HEARTS<td><code>&amp;#x1F495;</code><td><code>&amp;#128149;</code><td>
<tr><td>&#x1F496;<td><td>SPARKLING HEART<td><code>&amp;#x1F496;</code><td><code>&amp;#128150;</code><td>
<tr><td>&#x1F497;<td><td>GROWING HEART<td><code>&amp;#x1F497;</code><td><code>&amp;#128151;</code><td>
<tr><td>&#x1F499;<td><td>BLUE HEART<td><code>&amp;#x1F499;</code><td><code>&amp;#128153;</code><td>
<tr><td>&#x1F49A;<td><td>GREEN HEART<td><code>&amp;#x1F49A;</code><td><code>&amp;#128154;</code><td>
<tr><td>&#x1F49B;<td><td>YELLOW HEART<td><code>&amp;#x1F49B;</code><td><code>&amp;#128155;</code><td>
<tr><td>&#x1F49C;<td><td>PURPLE HEART<td><code>&amp;#x1F49C;</code><td><code>&amp;#128156;</code><td>
<tr><td>&#x1F5A4;<td><td>BLACK HEART<td><code>&amp;#x1F5A4;</code><td><code>&amp;#128420;</code><td>
<tr><td>&#x1F49D;<td><td>HEART WITH RIBBON<td><code>&amp;#x1F49D;</code><td><code>&amp;#128157;</code><td>
<tr><td>&#x1F49E;<td><td>REVOLVING HEARTS<td><code>&amp;#x1F49E;</code><td><code>&amp;#128158;</code><td>
<tr><td>&#x1F49F;<td><td>HEART DECORATION<td><code>&amp;#x1F49F;</code><td><code>&amp;#128159;</code><td>
<tr><td>&#x2763;<td><td>HEAVY HEART EXCLAMATION MARK ORNAMENT &#x224A; heavy heart exclamation<td><code>&amp;#x2763;</code><td><code>&amp;#10083;</code><td>
<tr><td>&#x1F48C;<td><td>LOVE LETTER<td><code>&amp;#x1F48C;</code><td><code>&amp;#128140;</code><td>
<tr><td>&#x1F4A4;<td><td>SLEEPING SYMBOL &#x224A; zzz<td><code>&amp;#x1F4A4;</code><td><code>&amp;#128164;</code><td>
<tr><td>&#x1F4A2;<td><td>ANGER SYMBOL<td><code>&amp;#x1F4A2;</code><td><code>&amp;#128162;</code><td>
<tr><td>&#x1F4A3;<td><td>BOMB<td><code>&amp;#x1F4A3;</code><td><code>&amp;#128163;</code><td>
<tr><td>&#x1F4A5;<td><td>COLLISION SYMBOL &#x224A; collision<td><code>&amp;#x1F4A5;</code><td><code>&amp;#128165;</code><td>
<tr><td>&#x1F4A6;<td><td>SPLASHING SWEAT SYMBOL &#x224A; sweat droplets<td><code>&amp;#x1F4A6;</code><td><code>&amp;#128166;</code><td>
<tr><td>&#x1F4A8;<td><td>DASH SYMBOL &#x224A; dashing<td><code>&amp;#x1F4A8;</code><td><code>&amp;#128168;</code><td>
<tr><td>&#x1F4AB;<td><td>DIZZY SYMBOL &#x224A; dizzy<td><code>&amp;#x1F4AB;</code><td><code>&amp;#128171;</code><td>
<tr><td>&#x1F4AC;<td><td>SPEECH BALLOON<td><code>&amp;#x1F4AC;</code><td><code>&amp;#128172;</code><td>
<tr><td>&#x1F5E8;<td><td>LEFT SPEECH BUBBLE<td><code>&amp;#x1F5E8;</code><td><code>&amp;#128488;</code><td>
<tr><td>&#x1F5EF;<td><td>RIGHT ANGER BUBBLE<td><code>&amp;#x1F5EF;</code><td><code>&amp;#128495;</code><td>
<tr><td>&#x1F4AD;<td><td>THOUGHT BALLOON<td><code>&amp;#x1F4AD;</code><td><code>&amp;#128173;</code><td>
<tr><td>&#x1F573;<td><td>HOLE<td><code>&amp;#x1F573;</code><td><code>&amp;#128371;</code><td>
<tr><td>&#x1F453;<td><td>EYEGLASSES &#x224A; glasses<td><code>&amp;#x1F453;</code><td><code>&amp;#128083;</code><td>
<tr><td>&#x1F576;<td><td>DARK SUNGLASSES &#x224A; sunglasses<td><code>&amp;#x1F576;</code><td><code>&amp;#128374;</code><td>
<tr><td>&#x1F454;<td><td>NECKTIE<td><code>&amp;#x1F454;</code><td><code>&amp;#128084;</code><td>
<tr><td>&#x1F455;<td><td>T-SHIRT<td><code>&amp;#x1F455;</code><td><code>&amp;#128085;</code><td>
<tr><td>&#x1F456;<td><td>JEANS<td><code>&amp;#x1F456;</code><td><code>&amp;#128086;</code><td>
<tr><td>&#x1F457;<td><td>DRESS<td><code>&amp;#x1F457;</code><td><code>&amp;#128087;</code><td>
<tr><td>&#x1F458;<td><td>KIMONO<td><code>&amp;#x1F458;</code><td><code>&amp;#128088;</code><td>
<tr><td>&#x1F459;<td><td>BIKINI<td><code>&amp;#x1F459;</code><td><code>&amp;#128089;</code><td>
<tr><td>&#x1F45A;<td><td>WOMANS CLOTHES &#x224A; womanâ€™s clothes<td><code>&amp;#x1F45A;</code><td><code>&amp;#128090;</code><td>
<tr><td>&#x1F45B;<td><td>PURSE<td><code>&amp;#x1F45B;</code><td><code>&amp;#128091;</code><td>
<tr><td>&#x1F45C;<td><td>HANDBAG<td><code>&amp;#x1F45C;</code><td><code>&amp;#128092;</code><td>
<tr><td>&#x1F45D;<td><td>POUCH<td><code>&amp;#x1F45D;</code><td><code>&amp;#128093;</code><td>
<tr><td>&#x1F6CD;<td><td>SHOPPING BAGS<td><code>&amp;#x1F6CD;</code><td><code>&amp;#128717;</code><td>
<tr><td>&#x1F392;<td><td>SCHOOL SATCHEL &#x224A; school backpack<td><code>&amp;#x1F392;</code><td><code>&amp;#127890;</code><td>
<tr><td>&#x1F45E;<td><td>MANS SHOE &#x224A; manâ€™s shoe<td><code>&amp;#x1F45E;</code><td><code>&amp;#128094;</code><td>
<tr><td>&#x1F45F;<td><td>ATHLETIC SHOE &#x224A; running shoe<td><code>&amp;#x1F45F;</code><td><code>&amp;#128095;</code><td>
<tr><td>&#x1F460;<td><td>HIGH-HEELED SHOE<td><code>&amp;#x1F460;</code><td><code>&amp;#128096;</code><td>
<tr><td>&#x1F461;<td><td>WOMANS SANDAL &#x224A; womanâ€™s sandal<td><code>&amp;#x1F461;</code><td><code>&amp;#128097;</code><td>
<tr><td>&#x1F462;<td><td>WOMANS BOOTS &#x224A; womanâ€™s boot<td><code>&amp;#x1F462;</code><td><code>&amp;#128098;</code><td>
<tr><td>&#x1F451;<td><td>CROWN<td><code>&amp;#x1F451;</code><td><code>&amp;#128081;</code><td>
<tr><td>&#x1F452;<td><td>WOMANS HAT &#x224A; womanâ€™s hat<td><code>&amp;#x1F452;</code><td><code>&amp;#128082;</code><td>
<tr><td>&#x1F3A9;<td><td>TOP HAT<td><code>&amp;#x1F3A9;</code><td><code>&amp;#127913;</code><td>
<tr><td>&#x1F393;<td><td>GRADUATION CAP<td><code>&amp;#x1F393;</code><td><code>&amp;#127891;</code><td>
<tr><td>&#x26D1;<td><td>HELMET WITH WHITE CROSS<td><code>&amp;#x26D1;</code><td><code>&amp;#9937;</code><td>
<tr><td>&#x1F4FF;<td><td>PRAYER BEADS<td><code>&amp;#x1F4FF;</code><td><code>&amp;#128255;</code><td>
<tr><td>&#x1F484;<td><td>LIPSTICK<td><code>&amp;#x1F484;</code><td><code>&amp;#128132;</code><td>
<tr><td>&#x1F48D;<td><td>RING<td><code>&amp;#x1F48D;</code><td><code>&amp;#128141;</code><td>
<tr><td>&#x1F48E;<td><td>GEM STONE<td><code>&amp;#x1F48E;</code><td><code>&amp;#128142;</code><td>
<tr><td>&#x1F435;<td><td>MONKEY FACE<td><code>&amp;#x1F435;</code><td><code>&amp;#128053;</code><td>
<tr><td>&#x1F412;<td><td>MONKEY<td><code>&amp;#x1F412;</code><td><code>&amp;#128018;</code><td>
<tr><td>&#x1F98D;<td><td>GORILLA<td><code>&amp;#x1F98D;</code><td><code>&amp;#129421;</code><td>
<tr><td>&#x1F436;<td><td>DOG FACE<td><code>&amp;#x1F436;</code><td><code>&amp;#128054;</code><td>
<tr><td>&#x1F415;<td><td>DOG<td><code>&amp;#x1F415;</code><td><code>&amp;#128021;</code><td>
<tr><td>&#x1F429;<td><td>POODLE<td><code>&amp;#x1F429;</code><td><code>&amp;#128041;</code><td>
<tr><td>&#x1F43A;<td><td>WOLF FACE<td><code>&amp;#x1F43A;</code><td><code>&amp;#128058;</code><td>
<tr><td>&#x1F98A;<td><td>FOX FACE<td><code>&amp;#x1F98A;</code><td><code>&amp;#129418;</code><td>
<tr><td>&#x1F431;<td><td>CAT FACE<td><code>&amp;#x1F431;</code><td><code>&amp;#128049;</code><td>
<tr><td>&#x1F408;<td><td>CAT<td><code>&amp;#x1F408;</code><td><code>&amp;#128008;</code><td>
<tr><td>&#x1F981;<td><td>LION FACE<td><code>&amp;#x1F981;</code><td><code>&amp;#129409;</code><td>
<tr><td>&#x1F42F;<td><td>TIGER FACE<td><code>&amp;#x1F42F;</code><td><code>&amp;#128047;</code><td>
<tr><td>&#x1F405;<td><td>TIGER<td><code>&amp;#x1F405;</code><td><code>&amp;#128005;</code><td>
<tr><td>&#x1F406;<td><td>LEOPARD<td><code>&amp;#x1F406;</code><td><code>&amp;#128006;</code><td>
<tr><td>&#x1F434;<td><td>HORSE FACE<td><code>&amp;#x1F434;</code><td><code>&amp;#128052;</code><td>
<tr><td>&#x1F40E;<td><td>HORSE<td><code>&amp;#x1F40E;</code><td><code>&amp;#128014;</code><td>
<tr><td>&#x1F98C;<td><td>DEER<td><code>&amp;#x1F98C;</code><td><code>&amp;#129420;</code><td>
<tr><td>&#x1F984;<td><td>UNICORN FACE<td><code>&amp;#x1F984;</code><td><code>&amp;#129412;</code><td>
<tr><td>&#x1F42E;<td><td>COW FACE<td><code>&amp;#x1F42E;</code><td><code>&amp;#128046;</code><td>
<tr><td>&#x1F402;<td><td>OX<td><code>&amp;#x1F402;</code><td><code>&amp;#128002;</code><td>
<tr><td>&#x1F403;<td><td>WATER BUFFALO<td><code>&amp;#x1F403;</code><td><code>&amp;#128003;</code><td>
<tr><td>&#x1F404;<td><td>COW<td><code>&amp;#x1F404;</code><td><code>&amp;#128004;</code><td>
<tr><td>&#x1F437;<td><td>PIG FACE<td><code>&amp;#x1F437;</code><td><code>&amp;#128055;</code><td>
<tr><td>&#x1F416;<td><td>PIG<td><code>&amp;#x1F416;</code><td><code>&amp;#128022;</code><td>
<tr><td>&#x1F417;<td><td>BOAR<td><code>&amp;#x1F417;</code><td><code>&amp;#128023;</code><td>
<tr><td>&#x1F43D;<td><td>PIG NOSE<td><code>&amp;#x1F43D;</code><td><code>&amp;#128061;</code><td>
<tr><td>&#x1F40F;<td><td>RAM<td><code>&amp;#x1F40F;</code><td><code>&amp;#128015;</code><td>
<tr><td>&#x1F411;<td><td>SHEEP<td><code>&amp;#x1F411;</code><td><code>&amp;#128017;</code><td>
<tr><td>&#x1F410;<td><td>GOAT<td><code>&amp;#x1F410;</code><td><code>&amp;#128016;</code><td>
<tr><td>&#x1F42A;<td><td>DROMEDARY CAMEL &#x224A; camel<td><code>&amp;#x1F42A;</code><td><code>&amp;#128042;</code><td>
<tr><td>&#x1F42B;<td><td>BACTRIAN CAMEL &#x224A; two-hump camel<td><code>&amp;#x1F42B;</code><td><code>&amp;#128043;</code><td>
<tr><td>&#x1F418;<td><td>ELEPHANT<td><code>&amp;#x1F418;</code><td><code>&amp;#128024;</code><td>
<tr><td>&#x1F98F;<td><td>RHINOCEROS<td><code>&amp;#x1F98F;</code><td><code>&amp;#129423;</code><td>
<tr><td>&#x1F42D;<td><td>MOUSE FACE<td><code>&amp;#x1F42D;</code><td><code>&amp;#128045;</code><td>
<tr><td>&#x1F401;<td><td>MOUSE<td><code>&amp;#x1F401;</code><td><code>&amp;#128001;</code><td>
<tr><td>&#x1F400;<td><td>RAT<td><code>&amp;#x1F400;</code><td><code>&amp;#128000;</code><td>
<tr><td>&#x1F439;<td><td>HAMSTER FACE<td><code>&amp;#x1F439;</code><td><code>&amp;#128057;</code><td>
<tr><td>&#x1F430;<td><td>RABBIT FACE<td><code>&amp;#x1F430;</code><td><code>&amp;#128048;</code><td>
<tr><td>&#x1F407;<td><td>RABBIT<td><code>&amp;#x1F407;</code><td><code>&amp;#128007;</code><td>
<tr><td>&#x1F43F;<td><td>CHIPMUNK<td><code>&amp;#x1F43F;</code><td><code>&amp;#128063;</code><td>
<tr><td>&#x1F987;<td><td>BAT<td><code>&amp;#x1F987;</code><td><code>&amp;#129415;</code><td>
<tr><td>&#x1F43B;<td><td>BEAR FACE<td><code>&amp;#x1F43B;</code><td><code>&amp;#128059;</code><td>
<tr><td>&#x1F428;<td><td>KOALA<td><code>&amp;#x1F428;</code><td><code>&amp;#128040;</code><td>
<tr><td>&#x1F43C;<td><td>PANDA FACE<td><code>&amp;#x1F43C;</code><td><code>&amp;#128060;</code><td>
<tr><td>&#x1F43E;<td><td>PAW PRINTS<td><code>&amp;#x1F43E;</code><td><code>&amp;#128062;</code><td>
<tr><td>&#x1F983;<td><td>TURKEY<td><code>&amp;#x1F983;</code><td><code>&amp;#129411;</code><td>
<tr><td>&#x1F414;<td><td>CHICKEN<td><code>&amp;#x1F414;</code><td><code>&amp;#128020;</code><td>
<tr><td>&#x1F413;<td><td>ROOSTER<td><code>&amp;#x1F413;</code><td><code>&amp;#128019;</code><td>
<tr><td>&#x1F423;<td><td>HATCHING CHICK<td><code>&amp;#x1F423;</code><td><code>&amp;#128035;</code><td>
<tr><td>&#x1F424;<td><td>BABY CHICK<td><code>&amp;#x1F424;</code><td><code>&amp;#128036;</code><td>
<tr><td>&#x1F425;<td><td>FRONT-FACING BABY CHICK<td><code>&amp;#x1F425;</code><td><code>&amp;#128037;</code><td>
<tr><td>&#x1F426;<td><td>BIRD<td><code>&amp;#x1F426;</code><td><code>&amp;#128038;</code><td>
<tr><td>&#x1F427;<td><td>PENGUIN<td><code>&amp;#x1F427;</code><td><code>&amp;#128039;</code><td>
<tr><td>&#x1F54A;<td><td>DOVE OF PEACE &#x224A; dove<td><code>&amp;#x1F54A;</code><td><code>&amp;#128330;</code><td>
<tr><td>&#x1F985;<td><td>EAGLE<td><code>&amp;#x1F985;</code><td><code>&amp;#129413;</code><td>
<tr><td>&#x1F986;<td><td>DUCK<td><code>&amp;#x1F986;</code><td><code>&amp;#129414;</code><td>
<tr><td>&#x1F989;<td><td>OWL<td><code>&amp;#x1F989;</code><td><code>&amp;#129417;</code><td>
<tr><td>&#x1F438;<td><td>FROG FACE<td><code>&amp;#x1F438;</code><td><code>&amp;#128056;</code><td>
<tr><td>&#x1F40A;<td><td>CROCODILE<td><code>&amp;#x1F40A;</code><td><code>&amp;#128010;</code><td>
<tr><td>&#x1F422;<td><td>TURTLE<td><code>&amp;#x1F422;</code><td><code>&amp;#128034;</code><td>
<tr><td>&#x1F98E;<td><td>LIZARD<td><code>&amp;#x1F98E;</code><td><code>&amp;#129422;</code><td>
<tr><td>&#x1F40D;<td><td>SNAKE<td><code>&amp;#x1F40D;</code><td><code>&amp;#128013;</code><td>
<tr><td>&#x1F432;<td><td>DRAGON FACE<td><code>&amp;#x1F432;</code><td><code>&amp;#128050;</code><td>
<tr><td>&#x1F409;<td><td>DRAGON<td><code>&amp;#x1F409;</code><td><code>&amp;#128009;</code><td>
<tr><td>&#x1F433;<td><td>SPOUTING WHALE<td><code>&amp;#x1F433;</code><td><code>&amp;#128051;</code><td>
<tr><td>&#x1F40B;<td><td>WHALE<td><code>&amp;#x1F40B;</code><td><code>&amp;#128011;</code><td>
<tr><td>&#x1F42C;<td><td>DOLPHIN<td><code>&amp;#x1F42C;</code><td><code>&amp;#128044;</code><td>
<tr><td>&#x1F41F;<td><td>FISH<td><code>&amp;#x1F41F;</code><td><code>&amp;#128031;</code><td>
<tr><td>&#x1F420;<td><td>TROPICAL FISH<td><code>&amp;#x1F420;</code><td><code>&amp;#128032;</code><td>
<tr><td>&#x1F421;<td><td>BLOWFISH<td><code>&amp;#x1F421;</code><td><code>&amp;#128033;</code><td>
<tr><td>&#x1F988;<td><td>SHARK<td><code>&amp;#x1F988;</code><td><code>&amp;#129416;</code><td>
<tr><td>&#x1F419;<td><td>OCTOPUS<td><code>&amp;#x1F419;</code><td><code>&amp;#128025;</code><td>
<tr><td>&#x1F41A;<td><td>SPIRAL SHELL<td><code>&amp;#x1F41A;</code><td><code>&amp;#128026;</code><td>
<tr><td>&#x1F980;<td><td>CRAB<td><code>&amp;#x1F980;</code><td><code>&amp;#129408;</code><td>
<tr><td>&#x1F990;<td><td>SHRIMP<td><code>&amp;#x1F990;</code><td><code>&amp;#129424;</code><td>
<tr><td>&#x1F991;<td><td>SQUID<td><code>&amp;#x1F991;</code><td><code>&amp;#129425;</code><td>
<tr><td>&#x1F98B;<td><td>BUTTERFLY<td><code>&amp;#x1F98B;</code><td><code>&amp;#129419;</code><td>
<tr><td>&#x1F40C;<td><td>SNAIL<td><code>&amp;#x1F40C;</code><td><code>&amp;#128012;</code><td>
<tr><td>&#x1F41B;<td><td>BUG<td><code>&amp;#x1F41B;</code><td><code>&amp;#128027;</code><td>
<tr><td>&#x1F41C;<td><td>ANT<td><code>&amp;#x1F41C;</code><td><code>&amp;#128028;</code><td>
<tr><td>&#x1F41D;<td><td>HONEYBEE<td><code>&amp;#x1F41D;</code><td><code>&amp;#128029;</code><td>
<tr><td>&#x1F41E;<td><td>LADY BEETLE<td><code>&amp;#x1F41E;</code><td><code>&amp;#128030;</code><td>
<tr><td>&#x1F577;<td><td>SPIDER<td><code>&amp;#x1F577;</code><td><code>&amp;#128375;</code><td>
<tr><td>&#x1F578;<td><td>SPIDER WEB<td><code>&amp;#x1F578;</code><td><code>&amp;#128376;</code><td>
<tr><td>&#x1F982;<td><td>SCORPION<td><code>&amp;#x1F982;</code><td><code>&amp;#129410;</code><td>
<tr><td>&#x1F490;<td><td>BOUQUET<td><code>&amp;#x1F490;</code><td><code>&amp;#128144;</code><td>
<tr><td>&#x1F338;<td><td>CHERRY BLOSSOM<td><code>&amp;#x1F338;</code><td><code>&amp;#127800;</code><td>
<tr><td>&#x1F4AE;<td><td>WHITE FLOWER<td><code>&amp;#x1F4AE;</code><td><code>&amp;#128174;</code><td>
<tr><td>&#x1F3F5;<td><td>ROSETTE<td><code>&amp;#x1F3F5;</code><td><code>&amp;#127989;</code><td>
<tr><td>&#x1F339;<td><td>ROSE<td><code>&amp;#x1F339;</code><td><code>&amp;#127801;</code><td>
<tr><td>&#x1F940;<td><td>WILTED FLOWER<td><code>&amp;#x1F940;</code><td><code>&amp;#129344;</code><td>
<tr><td>&#x1F33A;<td><td>HIBISCUS<td><code>&amp;#x1F33A;</code><td><code>&amp;#127802;</code><td>
<tr><td>&#x1F33B;<td><td>SUNFLOWER<td><code>&amp;#x1F33B;</code><td><code>&amp;#127803;</code><td>
<tr><td>&#x1F33C;<td><td>BLOSSOM<td><code>&amp;#x1F33C;</code><td><code>&amp;#127804;</code><td>
<tr><td>&#x1F337;<td><td>TULIP<td><code>&amp;#x1F337;</code><td><code>&amp;#127799;</code><td>
<tr><td>&#x1F331;<td><td>SEEDLING<td><code>&amp;#x1F331;</code><td><code>&amp;#127793;</code><td>
<tr><td>&#x1F332;<td><td>EVERGREEN TREE &#x224A; evergreen<td><code>&amp;#x1F332;</code><td><code>&amp;#127794;</code><td>
<tr><td>&#x1F333;<td><td>DECIDUOUS TREE<td><code>&amp;#x1F333;</code><td><code>&amp;#127795;</code><td>
<tr><td>&#x1F334;<td><td>PALM TREE<td><code>&amp;#x1F334;</code><td><code>&amp;#127796;</code><td>
<tr><td>&#x1F335;<td><td>CACTUS<td><code>&amp;#x1F335;</code><td><code>&amp;#127797;</code><td>
<tr><td>&#x1F33E;<td><td>EAR OF RICE &#x224A; sheaf of rice<td><code>&amp;#x1F33E;</code><td><code>&amp;#127806;</code><td>
<tr><td>&#x1F33F;<td><td>HERB<td><code>&amp;#x1F33F;</code><td><code>&amp;#127807;</code><td>
<tr><td>&#x2618;<td><td>SHAMROCK<td><code>&amp;#x2618;</code><td><code>&amp;#9752;</code><td>
<tr><td>&#x1F340;<td><td>FOUR LEAF CLOVER<td><code>&amp;#x1F340;</code><td><code>&amp;#127808;</code><td>
<tr><td>&#x1F341;<td><td>MAPLE LEAF<td><code>&amp;#x1F341;</code><td><code>&amp;#127809;</code><td>
<tr><td>&#x1F342;<td><td>FALLEN LEAF<td><code>&amp;#x1F342;</code><td><code>&amp;#127810;</code><td>
<tr><td>&#x1F343;<td><td>LEAF FLUTTERING IN WIND<td><code>&amp;#x1F343;</code><td><code>&amp;#127811;</code><td>
<tr><td>&#x1F347;<td><td>GRAPES<td><code>&amp;#x1F347;</code><td><code>&amp;#127815;</code><td>
<tr><td>&#x1F348;<td><td>MELON<td><code>&amp;#x1F348;</code><td><code>&amp;#127816;</code><td>
<tr><td>&#x1F349;<td><td>WATERMELON<td><code>&amp;#x1F349;</code><td><code>&amp;#127817;</code><td>
<tr><td>&#x1F34A;<td><td>TANGERINE<td><code>&amp;#x1F34A;</code><td><code>&amp;#127818;</code><td>
<tr><td>&#x1F34B;<td><td>LEMON<td><code>&amp;#x1F34B;</code><td><code>&amp;#127819;</code><td>
<tr><td>&#x1F34C;<td><td>BANANA<td><code>&amp;#x1F34C;</code><td><code>&amp;#127820;</code><td>
<tr><td>&#x1F34D;<td><td>PINEAPPLE<td><code>&amp;#x1F34D;</code><td><code>&amp;#127821;</code><td>
<tr><td>&#x1F34E;<td><td>RED APPLE<td><code>&amp;#x1F34E;</code><td><code>&amp;#127822;</code><td>
<tr><td>&#x1F34F;<td><td>GREEN APPLE<td><code>&amp;#x1F34F;</code><td><code>&amp;#127823;</code><td>
<tr><td>&#x1F350;<td><td>PEAR<td><code>&amp;#x1F350;</code><td><code>&amp;#127824;</code><td>
<tr><td>&#x1F351;<td><td>PEACH<td><code>&amp;#x1F351;</code><td><code>&amp;#127825;</code><td>
<tr><td>&#x1F352;<td><td>CHERRIES<td><code>&amp;#x1F352;</code><td><code>&amp;#127826;</code><td>
<tr><td>&#x1F353;<td><td>STRAWBERRY<td><code>&amp;#x1F353;</code><td><code>&amp;#127827;</code><td>
<tr><td>&#x1F95D;<td><td>KIWIFRUIT<td><code>&amp;#x1F95D;</code><td><code>&amp;#129373;</code><td>
<tr><td>&#x1F345;<td><td>TOMATO<td><code>&amp;#x1F345;</code><td><code>&amp;#127813;</code><td>
<tr><td>&#x1F951;<td><td>AVOCADO<td><code>&amp;#x1F951;</code><td><code>&amp;#129361;</code><td>
<tr><td>&#x1F346;<td><td>AUBERGINE &#x224A; eggplant<td><code>&amp;#x1F346;</code><td><code>&amp;#127814;</code><td>
<tr><td>&#x1F954;<td><td>POTATO<td><code>&amp;#x1F954;</code><td><code>&amp;#129364;</code><td>
<tr><td>&#x1F955;<td><td>CARROT<td><code>&amp;#x1F955;</code><td><code>&amp;#129365;</code><td>
<tr><td>&#x1F33D;<td><td>EAR OF MAIZE &#x224A; ear of corn<td><code>&amp;#x1F33D;</code><td><code>&amp;#127805;</code><td>
<tr><td>&#x1F336;<td><td>HOT PEPPER<td><code>&amp;#x1F336;</code><td><code>&amp;#127798;</code><td>
<tr><td>&#x1F952;<td><td>CUCUMBER<td><code>&amp;#x1F952;</code><td><code>&amp;#129362;</code><td>
<tr><td>&#x1F344;<td><td>MUSHROOM<td><code>&amp;#x1F344;</code><td><code>&amp;#127812;</code><td>
<tr><td>&#x1F95C;<td><td>PEANUTS<td><code>&amp;#x1F95C;</code><td><code>&amp;#129372;</code><td>
<tr><td>&#x1F330;<td><td>CHESTNUT<td><code>&amp;#x1F330;</code><td><code>&amp;#127792;</code><td>
<tr><td>&#x1F35E;<td><td>BREAD<td><code>&amp;#x1F35E;</code><td><code>&amp;#127838;</code><td>
<tr><td>&#x1F950;<td><td>CROISSANT<td><code>&amp;#x1F950;</code><td><code>&amp;#129360;</code><td>
<tr><td>&#x1F956;<td><td>BAGUETTE BREAD<td><code>&amp;#x1F956;</code><td><code>&amp;#129366;</code><td>
<tr><td>&#x1F95E;<td><td>PANCAKES<td><code>&amp;#x1F95E;</code><td><code>&amp;#129374;</code><td>
<tr><td>&#x1F9C0;<td><td>CHEESE WEDGE<td><code>&amp;#x1F9C0;</code><td><code>&amp;#129472;</code><td>
<tr><td>&#x1F356;<td><td>MEAT ON BONE<td><code>&amp;#x1F356;</code><td><code>&amp;#127830;</code><td>
<tr><td>&#x1F357;<td><td>POULTRY LEG<td><code>&amp;#x1F357;</code><td><code>&amp;#127831;</code><td>
<tr><td>&#x1F953;<td><td>BACON<td><code>&amp;#x1F953;</code><td><code>&amp;#129363;</code><td>
<tr><td>&#x1F354;<td><td>HAMBURGER<td><code>&amp;#x1F354;</code><td><code>&amp;#127828;</code><td>
<tr><td>&#x1F35F;<td><td>FRENCH FRIES<td><code>&amp;#x1F35F;</code><td><code>&amp;#127839;</code><td>
<tr><td>&#x1F355;<td><td>SLICE OF PIZZA &#x224A; pizza<td><code>&amp;#x1F355;</code><td><code>&amp;#127829;</code><td>
<tr><td>&#x1F32D;<td><td>HOT DOG<td><code>&amp;#x1F32D;</code><td><code>&amp;#127789;</code><td>
<tr><td>&#x1F32E;<td><td>TACO<td><code>&amp;#x1F32E;</code><td><code>&amp;#127790;</code><td>
<tr><td>&#x1F32F;<td><td>BURRITO<td><code>&amp;#x1F32F;</code><td><code>&amp;#127791;</code><td>
<tr><td>&#x1F959;<td><td>STUFFED FLATBREAD<td><code>&amp;#x1F959;</code><td><code>&amp;#129369;</code><td>
<tr><td>&#x1F95A;<td><td>EGG<td><code>&amp;#x1F95A;</code><td><code>&amp;#129370;</code><td>
<tr><td>&#x1F373;<td><td>COOKING<td><code>&amp;#x1F373;</code><td><code>&amp;#127859;</code><td>
<tr><td>&#x1F958;<td><td>SHALLOW PAN OF FOOD<td><code>&amp;#x1F958;</code><td><code>&amp;#129368;</code><td>
<tr><td>&#x1F372;<td><td>POT OF FOOD<td><code>&amp;#x1F372;</code><td><code>&amp;#127858;</code><td>
<tr><td>&#x1F957;<td><td>GREEN SALAD<td><code>&amp;#x1F957;</code><td><code>&amp;#129367;</code><td>
<tr><td>&#x1F37F;<td><td>POPCORN<td><code>&amp;#x1F37F;</code><td><code>&amp;#127871;</code><td>
<tr><td>&#x1F371;<td><td>BENTO BOX<td><code>&amp;#x1F371;</code><td><code>&amp;#127857;</code><td>
<tr><td>&#x1F358;<td><td>RICE CRACKER<td><code>&amp;#x1F358;</code><td><code>&amp;#127832;</code><td>
<tr><td>&#x1F359;<td><td>RICE BALL<td><code>&amp;#x1F359;</code><td><code>&amp;#127833;</code><td>
<tr><td>&#x1F35A;<td><td>COOKED RICE<td><code>&amp;#x1F35A;</code><td><code>&amp;#127834;</code><td>
<tr><td>&#x1F35B;<td><td>CURRY AND RICE &#x224A; curry rice<td><code>&amp;#x1F35B;</code><td><code>&amp;#127835;</code><td>
<tr><td>&#x1F35C;<td><td>STEAMING BOWL<td><code>&amp;#x1F35C;</code><td><code>&amp;#127836;</code><td>
<tr><td>&#x1F35D;<td><td>SPAGHETTI<td><code>&amp;#x1F35D;</code><td><code>&amp;#127837;</code><td>
<tr><td>&#x1F360;<td><td>ROASTED SWEET POTATO<td><code>&amp;#x1F360;</code><td><code>&amp;#127840;</code><td>
<tr><td>&#x1F362;<td><td>ODEN<td><code>&amp;#x1F362;</code><td><code>&amp;#127842;</code><td>
<tr><td>&#x1F363;<td><td>SUSHI<td><code>&amp;#x1F363;</code><td><code>&amp;#127843;</code><td>
<tr><td>&#x1F364;<td><td>FRIED SHRIMP<td><code>&amp;#x1F364;</code><td><code>&amp;#127844;</code><td>
<tr><td>&#x1F365;<td><td>FISH CAKE WITH SWIRL DESIGN &#x224A; fish cake with swirl<td><code>&amp;#x1F365;</code><td><code>&amp;#127845;</code><td>
<tr><td>&#x1F361;<td><td>DANGO<td><code>&amp;#x1F361;</code><td><code>&amp;#127841;</code><td>
<tr><td>&#x1F366;<td><td>SOFT ICE CREAM<td><code>&amp;#x1F366;</code><td><code>&amp;#127846;</code><td>
<tr><td>&#x1F367;<td><td>SHAVED ICE<td><code>&amp;#x1F367;</code><td><code>&amp;#127847;</code><td>
<tr><td>&#x1F368;<td><td>ICE CREAM<td><code>&amp;#x1F368;</code><td><code>&amp;#127848;</code><td>
<tr><td>&#x1F369;<td><td>DOUGHNUT<td><code>&amp;#x1F369;</code><td><code>&amp;#127849;</code><td>
<tr><td>&#x1F36A;<td><td>COOKIE<td><code>&amp;#x1F36A;</code><td><code>&amp;#127850;</code><td>
<tr><td>&#x1F382;<td><td>BIRTHDAY CAKE<td><code>&amp;#x1F382;</code><td><code>&amp;#127874;</code><td>
<tr><td>&#x1F370;<td><td>SHORTCAKE<td><code>&amp;#x1F370;</code><td><code>&amp;#127856;</code><td>
<tr><td>&#x1F36B;<td><td>CHOCOLATE BAR<td><code>&amp;#x1F36B;</code><td><code>&amp;#127851;</code><td>
<tr><td>&#x1F36C;<td><td>CANDY<td><code>&amp;#x1F36C;</code><td><code>&amp;#127852;</code><td>
<tr><td>&#x1F36D;<td><td>LOLLIPOP<td><code>&amp;#x1F36D;</code><td><code>&amp;#127853;</code><td>
<tr><td>&#x1F36E;<td><td>CUSTARD<td><code>&amp;#x1F36E;</code><td><code>&amp;#127854;</code><td>
<tr><td>&#x1F36F;<td><td>HONEY POT<td><code>&amp;#x1F36F;</code><td><code>&amp;#127855;</code><td>
<tr><td>&#x1F37C;<td><td>BABY BOTTLE<td><code>&amp;#x1F37C;</code><td><code>&amp;#127868;</code><td>
<tr><td>&#x1F95B;<td><td>GLASS OF MILK<td><code>&amp;#x1F95B;</code><td><code>&amp;#129371;</code><td>
<tr><td>&#x2615;<td><td>HOT BEVERAGE<td><code>&amp;#x2615;</code><td><code>&amp;#9749;</code><td>
<tr><td>&#x1F375;<td><td>TEACUP WITHOUT HANDLE<td><code>&amp;#x1F375;</code><td><code>&amp;#127861;</code><td>
<tr><td>&#x1F376;<td><td>SAKE BOTTLE AND CUP &#x224A; sake<td><code>&amp;#x1F376;</code><td><code>&amp;#127862;</code><td>
<tr><td>&#x1F37E;<td><td>BOTTLE WITH POPPING CORK<td><code>&amp;#x1F37E;</code><td><code>&amp;#127870;</code><td>
<tr><td>&#x1F377;<td><td>WINE GLASS<td><code>&amp;#x1F377;</code><td><code>&amp;#127863;</code><td>
<tr><td>&#x1F378;<td><td>COCKTAIL GLASS<td><code>&amp;#x1F378;</code><td><code>&amp;#127864;</code><td>
<tr><td>&#x1F379;<td><td>TROPICAL DRINK<td><code>&amp;#x1F379;</code><td><code>&amp;#127865;</code><td>
<tr><td>&#x1F37A;<td><td>BEER MUG<td><code>&amp;#x1F37A;</code><td><code>&amp;#127866;</code><td>
<tr><td>&#x1F37B;<td><td>CLINKING BEER MUGS<td><code>&amp;#x1F37B;</code><td><code>&amp;#127867;</code><td>
<tr><td>&#x1F942;<td><td>CLINKING GLASSES<td><code>&amp;#x1F942;</code><td><code>&amp;#129346;</code><td>
<tr><td>&#x1F943;<td><td>TUMBLER GLASS<td><code>&amp;#x1F943;</code><td><code>&amp;#129347;</code><td>
<tr><td>&#x1F37D;<td><td>FORK AND KNIFE WITH PLATE<td><code>&amp;#x1F37D;</code><td><code>&amp;#127869;</code><td>
<tr><td>&#x1F374;<td><td>FORK AND KNIFE<td><code>&amp;#x1F374;</code><td><code>&amp;#127860;</code><td>
<tr><td>&#x1F944;<td><td>SPOON<td><code>&amp;#x1F944;</code><td><code>&amp;#129348;</code><td>
<tr><td>&#x1F52A;<td><td>HOCHO &#x224A; kitchen knife<td><code>&amp;#x1F52A;</code><td><code>&amp;#128298;</code><td>
<tr><td>&#x1F3FA;<td><td>AMPHORA<td><code>&amp;#x1F3FA;</code><td><code>&amp;#127994;</code><td>
<tr><td>&#x1F30D;<td><td>EARTH GLOBE EUROPE-AFRICA &#x224A; globe showing europe-africa<td><code>&amp;#x1F30D;</code><td><code>&amp;#127757;</code><td>
<tr><td>&#x1F30E;<td><td>EARTH GLOBE AMERICAS &#x224A; globe showing americas<td><code>&amp;#x1F30E;</code><td><code>&amp;#127758;</code><td>
<tr><td>&#x1F30F;<td><td>EARTH GLOBE ASIA-AUSTRALIA &#x224A; globe showing asia-australia<td><code>&amp;#x1F30F;</code><td><code>&amp;#127759;</code><td>
<tr><td>&#x1F310;<td><td>GLOBE WITH MERIDIANS<td><code>&amp;#x1F310;</code><td><code>&amp;#127760;</code><td>
<tr><td>&#x1F5FA;<td><td>WORLD MAP<td><code>&amp;#x1F5FA;</code><td><code>&amp;#128506;</code><td>
<tr><td>&#x1F5FE;<td><td>SILHOUETTE OF JAPAN &#x224A; map of japan<td><code>&amp;#x1F5FE;</code><td><code>&amp;#128510;</code><td>
<tr><td>&#x1F3D4;<td><td>SNOW CAPPED MOUNTAIN &#x224A; snow-capped mountain<td><code>&amp;#x1F3D4;</code><td><code>&amp;#127956;</code><td>
<tr><td>&#x26F0;<td><td>MOUNTAIN<td><code>&amp;#x26F0;</code><td><code>&amp;#9968;</code><td>
<tr><td>&#x1F30B;<td><td>VOLCANO<td><code>&amp;#x1F30B;</code><td><code>&amp;#127755;</code><td>
<tr><td>&#x1F5FB;<td><td>MOUNT FUJI<td><code>&amp;#x1F5FB;</code><td><code>&amp;#128507;</code><td>
<tr><td>&#x1F3D5;<td><td>CAMPING<td><code>&amp;#x1F3D5;</code><td><code>&amp;#127957;</code><td>
<tr><td>&#x1F3D6;<td><td>BEACH WITH UMBRELLA<td><code>&amp;#x1F3D6;</code><td><code>&amp;#127958;</code><td>
<tr><td>&#x1F3DC;<td><td>DESERT<td><code>&amp;#x1F3DC;</code><td><code>&amp;#127964;</code><td>
<tr><td>&#x1F3DD;<td><td>DESERT ISLAND<td><code>&amp;#x1F3DD;</code><td><code>&amp;#127965;</code><td>
<tr><td>&#x1F3DE;<td><td>NATIONAL PARK<td><code>&amp;#x1F3DE;</code><td><code>&amp;#127966;</code><td>
<tr><td>&#x1F3DF;<td><td>STADIUM<td><code>&amp;#x1F3DF;</code><td><code>&amp;#127967;</code><td>
<tr><td>&#x1F3DB;<td><td>CLASSICAL BUILDING<td><code>&amp;#x1F3DB;</code><td><code>&amp;#127963;</code><td>
<tr><td>&#x1F3D7;<td><td>BUILDING CONSTRUCTION<td><code>&amp;#x1F3D7;</code><td><code>&amp;#127959;</code><td>
<tr><td>&#x1F3D8;<td><td>HOUSE BUILDINGS<td><code>&amp;#x1F3D8;</code><td><code>&amp;#127960;</code><td>
<tr><td>&#x1F3D9;<td><td>CITYSCAPE<td><code>&amp;#x1F3D9;</code><td><code>&amp;#127961;</code><td>
<tr><td>&#x1F3DA;<td><td>DERELICT HOUSE BUILDING<td><code>&amp;#x1F3DA;</code><td><code>&amp;#127962;</code><td>
<tr><td>&#x1F3E0;<td><td>HOUSE BUILDING<td><code>&amp;#x1F3E0;</code><td><code>&amp;#127968;</code><td>
<tr><td>&#x1F3E1;<td><td>HOUSE WITH GARDEN<td><code>&amp;#x1F3E1;</code><td><code>&amp;#127969;</code><td>
<tr><td>&#x1F3E2;<td><td>OFFICE BUILDING<td><code>&amp;#x1F3E2;</code><td><code>&amp;#127970;</code><td>
<tr><td>&#x1F3E3;<td><td>JAPANESE POST OFFICE<td><code>&amp;#x1F3E3;</code><td><code>&amp;#127971;</code><td>
<tr><td>&#x1F3E4;<td><td>EUROPEAN POST OFFICE &#x224A; post office<td><code>&amp;#x1F3E4;</code><td><code>&amp;#127972;</code><td>
<tr><td>&#x1F3E5;<td><td>HOSPITAL<td><code>&amp;#x1F3E5;</code><td><code>&amp;#127973;</code><td>
<tr><td>&#x1F3E6;<td><td>BANK<td><code>&amp;#x1F3E6;</code><td><code>&amp;#127974;</code><td>
<tr><td>&#x1F3E8;<td><td>HOTEL<td><code>&amp;#x1F3E8;</code><td><code>&amp;#127976;</code><td>
<tr><td>&#x1F3E9;<td><td>LOVE HOTEL<td><code>&amp;#x1F3E9;</code><td><code>&amp;#127977;</code><td>
<tr><td>&#x1F3EA;<td><td>CONVENIENCE STORE<td><code>&amp;#x1F3EA;</code><td><code>&amp;#127978;</code><td>
<tr><td>&#x1F3EB;<td><td>SCHOOL<td><code>&amp;#x1F3EB;</code><td><code>&amp;#127979;</code><td>
<tr><td>&#x1F3EC;<td><td>DEPARTMENT STORE<td><code>&amp;#x1F3EC;</code><td><code>&amp;#127980;</code><td>
<tr><td>&#x1F3ED;<td><td>FACTORY<td><code>&amp;#x1F3ED;</code><td><code>&amp;#127981;</code><td>
<tr><td>&#x1F3EF;<td><td>JAPANESE CASTLE<td><code>&amp;#x1F3EF;</code><td><code>&amp;#127983;</code><td>
<tr><td>&#x1F3F0;<td><td>EUROPEAN CASTLE &#x224A; castle<td><code>&amp;#x1F3F0;</code><td><code>&amp;#127984;</code><td>
<tr><td>&#x1F492;<td><td>WEDDING<td><code>&amp;#x1F492;</code><td><code>&amp;#128146;</code><td>
<tr><td>&#x1F5FC;<td><td>TOKYO TOWER<td><code>&amp;#x1F5FC;</code><td><code>&amp;#128508;</code><td>
<tr><td>&#x1F5FD;<td><td>STATUE OF LIBERTY<td><code>&amp;#x1F5FD;</code><td><code>&amp;#128509;</code><td>
<tr><td>&#x26EA;<td><td>CHURCH<td><code>&amp;#x26EA;</code><td><code>&amp;#9962;</code><td>
<tr><td>&#x1F54C;<td><td>MOSQUE<td><code>&amp;#x1F54C;</code><td><code>&amp;#128332;</code><td>
<tr><td>&#x1F54D;<td><td>SYNAGOGUE<td><code>&amp;#x1F54D;</code><td><code>&amp;#128333;</code><td>
<tr><td>&#x26E9;<td><td>SHINTO SHRINE<td><code>&amp;#x26E9;</code><td><code>&amp;#9961;</code><td>
<tr><td>&#x1F54B;<td><td>KAABA<td><code>&amp;#x1F54B;</code><td><code>&amp;#128331;</code><td>
<tr><td>&#x26F2;<td><td>FOUNTAIN<td><code>&amp;#x26F2;</code><td><code>&amp;#9970;</code><td>
<tr><td>&#x26FA;<td><td>TENT<td><code>&amp;#x26FA;</code><td><code>&amp;#9978;</code><td>
<tr><td>&#x1F301;<td><td>FOGGY<td><code>&amp;#x1F301;</code><td><code>&amp;#127745;</code><td>
<tr><td>&#x1F303;<td><td>NIGHT WITH STARS<td><code>&amp;#x1F303;</code><td><code>&amp;#127747;</code><td>
<tr><td>&#x1F304;<td><td>SUNRISE OVER MOUNTAINS<td><code>&amp;#x1F304;</code><td><code>&amp;#127748;</code><td>
<tr><td>&#x1F305;<td><td>SUNRISE<td><code>&amp;#x1F305;</code><td><code>&amp;#127749;</code><td>
<tr><td>&#x1F306;<td><td>CITYSCAPE AT DUSK<td><code>&amp;#x1F306;</code><td><code>&amp;#127750;</code><td>
<tr><td>&#x1F307;<td><td>SUNSET OVER BUILDINGS &#x224A; sunset<td><code>&amp;#x1F307;</code><td><code>&amp;#127751;</code><td>
<tr><td>&#x1F309;<td><td>BRIDGE AT NIGHT<td><code>&amp;#x1F309;</code><td><code>&amp;#127753;</code><td>
<tr><td>&#x2668;<td><td>HOT SPRINGS<td><code>&amp;#x2668;</code><td><code>&amp;#9832;</code><td>
<tr><td>&#x1F30C;<td><td>MILKY WAY<td><code>&amp;#x1F30C;</code><td><code>&amp;#127756;</code><td>
<tr><td>&#x1F3A0;<td><td>CAROUSEL HORSE<td><code>&amp;#x1F3A0;</code><td><code>&amp;#127904;</code><td>
<tr><td>&#x1F3A1;<td><td>FERRIS WHEEL<td><code>&amp;#x1F3A1;</code><td><code>&amp;#127905;</code><td>
<tr><td>&#x1F3A2;<td><td>ROLLER COASTER<td><code>&amp;#x1F3A2;</code><td><code>&amp;#127906;</code><td>
<tr><td>&#x1F488;<td><td>BARBER POLE<td><code>&amp;#x1F488;</code><td><code>&amp;#128136;</code><td>
<tr><td>&#x1F3AA;<td><td>CIRCUS TENT<td><code>&amp;#x1F3AA;</code><td><code>&amp;#127914;</code><td>
<tr><td>&#x1F3AD;<td><td>PERFORMING ARTS<td><code>&amp;#x1F3AD;</code><td><code>&amp;#127917;</code><td>
<tr><td>&#x1F5BC;<td><td>FRAME WITH PICTURE<td><code>&amp;#x1F5BC;</code><td><code>&amp;#128444;</code><td>
<tr><td>&#x1F3A8;<td><td>ARTIST PALETTE<td><code>&amp;#x1F3A8;</code><td><code>&amp;#127912;</code><td>
<tr><td>&#x1F3B0;<td><td>SLOT MACHINE<td><code>&amp;#x1F3B0;</code><td><code>&amp;#127920;</code><td>
<tr><td>&#x1F682;<td><td>STEAM LOCOMOTIVE &#x224A; locomotive<td><code>&amp;#x1F682;</code><td><code>&amp;#128642;</code><td>
<tr><td>&#x1F683;<td><td>RAILWAY CAR<td><code>&amp;#x1F683;</code><td><code>&amp;#128643;</code><td>
<tr><td>&#x1F684;<td><td>HIGH-SPEED TRAIN<td><code>&amp;#x1F684;</code><td><code>&amp;#128644;</code><td>
<tr><td>&#x1F685;<td><td>HIGH-SPEED TRAIN WITH BULLET NOSE<td><code>&amp;#x1F685;</code><td><code>&amp;#128645;</code><td>
<tr><td>&#x1F686;<td><td>TRAIN<td><code>&amp;#x1F686;</code><td><code>&amp;#128646;</code><td>
<tr><td>&#x1F687;<td><td>METRO<td><code>&amp;#x1F687;</code><td><code>&amp;#128647;</code><td>
<tr><td>&#x1F688;<td><td>LIGHT RAIL<td><code>&amp;#x1F688;</code><td><code>&amp;#128648;</code><td>
<tr><td>&#x1F689;<td><td>STATION<td><code>&amp;#x1F689;</code><td><code>&amp;#128649;</code><td>
<tr><td>&#x1F68A;<td><td>TRAM<td><code>&amp;#x1F68A;</code><td><code>&amp;#128650;</code><td>
<tr><td>&#x1F69D;<td><td>MONORAIL<td><code>&amp;#x1F69D;</code><td><code>&amp;#128669;</code><td>
<tr><td>&#x1F69E;<td><td>MOUNTAIN RAILWAY<td><code>&amp;#x1F69E;</code><td><code>&amp;#128670;</code><td>
<tr><td>&#x1F68B;<td><td>TRAM CAR<td><code>&amp;#x1F68B;</code><td><code>&amp;#128651;</code><td>
<tr><td>&#x1F68C;<td><td>BUS<td><code>&amp;#x1F68C;</code><td><code>&amp;#128652;</code><td>
<tr><td>&#x1F68D;<td><td>ONCOMING BUS<td><code>&amp;#x1F68D;</code><td><code>&amp;#128653;</code><td>
<tr><td>&#x1F68E;<td><td>TROLLEYBUS<td><code>&amp;#x1F68E;</code><td><code>&amp;#128654;</code><td>
<tr><td>&#x1F690;<td><td>MINIBUS<td><code>&amp;#x1F690;</code><td><code>&amp;#128656;</code><td>
<tr><td>&#x1F691;<td><td>AMBULANCE<td><code>&amp;#x1F691;</code><td><code>&amp;#128657;</code><td>
<tr><td>&#x1F692;<td><td>FIRE ENGINE<td><code>&amp;#x1F692;</code><td><code>&amp;#128658;</code><td>
<tr><td>&#x1F693;<td><td>POLICE CAR<td><code>&amp;#x1F693;</code><td><code>&amp;#128659;</code><td>
<tr><td>&#x1F694;<td><td>ONCOMING POLICE CAR<td><code>&amp;#x1F694;</code><td><code>&amp;#128660;</code><td>
<tr><td>&#x1F695;<td><td>TAXI<td><code>&amp;#x1F695;</code><td><code>&amp;#128661;</code><td>
<tr><td>&#x1F696;<td><td>ONCOMING TAXI<td><code>&amp;#x1F696;</code><td><code>&amp;#128662;</code><td>
<tr><td>&#x1F697;<td><td>AUTOMOBILE<td><code>&amp;#x1F697;</code><td><code>&amp;#128663;</code><td>
<tr><td>&#x1F698;<td><td>ONCOMING AUTOMOBILE<td><code>&amp;#x1F698;</code><td><code>&amp;#128664;</code><td>
<tr><td>&#x1F699;<td><td>RECREATIONAL VEHICLE<td><code>&amp;#x1F699;</code><td><code>&amp;#128665;</code><td>
<tr><td>&#x1F69A;<td><td>DELIVERY TRUCK<td><code>&amp;#x1F69A;</code><td><code>&amp;#128666;</code><td>
<tr><td>&#x1F69B;<td><td>ARTICULATED LORRY<td><code>&amp;#x1F69B;</code><td><code>&amp;#128667;</code><td>
<tr><td>&#x1F69C;<td><td>TRACTOR<td><code>&amp;#x1F69C;</code><td><code>&amp;#128668;</code><td>
<tr><td>&#x1F6B2;<td><td>BICYCLE<td><code>&amp;#x1F6B2;</code><td><code>&amp;#128690;</code><td>
<tr><td>&#x1F6F4;<td><td>SCOOTER &#x224A; kick scooter<td><code>&amp;#x1F6F4;</code><td><code>&amp;#128756;</code><td>
<tr><td>&#x1F6F5;<td><td>MOTOR SCOOTER<td><code>&amp;#x1F6F5;</code><td><code>&amp;#128757;</code><td>
<tr><td>&#x1F3CE;<td><td>RACING CAR<td><code>&amp;#x1F3CE;</code><td><code>&amp;#127950;</code><td>
<tr><td>&#x1F3CD;<td><td>RACING MOTORCYCLE &#x224A; motorcycle<td><code>&amp;#x1F3CD;</code><td><code>&amp;#127949;</code><td>
<tr><td>&#x1F68F;<td><td>BUS STOP<td><code>&amp;#x1F68F;</code><td><code>&amp;#128655;</code><td>
<tr><td>&#x1F6E3;<td><td>MOTORWAY<td><code>&amp;#x1F6E3;</code><td><code>&amp;#128739;</code><td>
<tr><td>&#x1F6E4;<td><td>RAILWAY TRACK<td><code>&amp;#x1F6E4;</code><td><code>&amp;#128740;</code><td>
<tr><td>&#x26FD;<td><td>FUEL PUMP<td><code>&amp;#x26FD;</code><td><code>&amp;#9981;</code><td>
<tr><td>&#x1F6A8;<td><td>POLICE CARS REVOLVING LIGHT &#x224A; police carâ€™s light<td><code>&amp;#x1F6A8;</code><td><code>&amp;#128680;</code><td>
<tr><td>&#x1F6A5;<td><td>HORIZONTAL TRAFFIC LIGHT<td><code>&amp;#x1F6A5;</code><td><code>&amp;#128677;</code><td>
<tr><td>&#x1F6A6;<td><td>VERTICAL TRAFFIC LIGHT<td><code>&amp;#x1F6A6;</code><td><code>&amp;#128678;</code><td>
<tr><td>&#x1F6A7;<td><td>CONSTRUCTION SIGN &#x224A; construction<td><code>&amp;#x1F6A7;</code><td><code>&amp;#128679;</code><td>
<tr><td>&#x1F6D1;<td><td>OCTAGONAL SIGN &#x224A; stop sign<td><code>&amp;#x1F6D1;</code><td><code>&amp;#128721;</code><td>
<tr><td>&#x2693;<td><td>ANCHOR<td><code>&amp;#x2693;</code><td><code>&amp;#9875;</code><td>
<tr><td>&#x26F5;<td><td>SAILBOAT<td><code>&amp;#x26F5;</code><td><code>&amp;#9973;</code><td>
<tr><td>&#x1F6F6;<td><td>CANOE<td><code>&amp;#x1F6F6;</code><td><code>&amp;#128758;</code><td>
<tr><td>&#x1F6A4;<td><td>SPEEDBOAT<td><code>&amp;#x1F6A4;</code><td><code>&amp;#128676;</code><td>
<tr><td>&#x1F6F3;<td><td>PASSENGER SHIP<td><code>&amp;#x1F6F3;</code><td><code>&amp;#128755;</code><td>
<tr><td>&#x26F4;<td><td>FERRY<td><code>&amp;#x26F4;</code><td><code>&amp;#9972;</code><td>
<tr><td>&#x1F6E5;<td><td>MOTOR BOAT<td><code>&amp;#x1F6E5;</code><td><code>&amp;#128741;</code><td>
<tr><td>&#x1F6A2;<td><td>SHIP<td><code>&amp;#x1F6A2;</code><td><code>&amp;#128674;</code><td>
<tr><td>&#x2708;<td><td>AIRPLANE<td><code>&amp;#x2708;</code><td><code>&amp;#9992;</code><td>
<tr><td>&#x1F6E9;<td><td>SMALL AIRPLANE<td><code>&amp;#x1F6E9;</code><td><code>&amp;#128745;</code><td>
<tr><td>&#x1F6EB;<td><td>AIRPLANE DEPARTURE<td><code>&amp;#x1F6EB;</code><td><code>&amp;#128747;</code><td>
<tr><td>&#x1F6EC;<td><td>AIRPLANE ARRIVING &#x224A; airplane arrival<td><code>&amp;#x1F6EC;</code><td><code>&amp;#128748;</code><td>
<tr><td>&#x1F4BA;<td><td>SEAT<td><code>&amp;#x1F4BA;</code><td><code>&amp;#128186;</code><td>
<tr><td>&#x1F681;<td><td>HELICOPTER<td><code>&amp;#x1F681;</code><td><code>&amp;#128641;</code><td>
<tr><td>&#x1F69F;<td><td>SUSPENSION RAILWAY<td><code>&amp;#x1F69F;</code><td><code>&amp;#128671;</code><td>
<tr><td>&#x1F6A0;<td><td>MOUNTAIN CABLEWAY<td><code>&amp;#x1F6A0;</code><td><code>&amp;#128672;</code><td>
<tr><td>&#x1F6A1;<td><td>AERIAL TRAMWAY<td><code>&amp;#x1F6A1;</code><td><code>&amp;#128673;</code><td>
<tr><td>&#x1F680;<td><td>ROCKET<td><code>&amp;#x1F680;</code><td><code>&amp;#128640;</code><td>
<tr><td>&#x1F6F0;<td><td>SATELLITE<td><code>&amp;#x1F6F0;</code><td><code>&amp;#128752;</code><td>
<tr><td>&#x1F6CE;<td><td>BELLHOP BELL<td><code>&amp;#x1F6CE;</code><td><code>&amp;#128718;</code><td>
<tr><td>&#x1F6AA;<td><td>DOOR<td><code>&amp;#x1F6AA;</code><td><code>&amp;#128682;</code><td>
<tr><td>&#x1F6CC;<td><td>SLEEPING ACCOMMODATION &#x224A; person in bed<td><code>&amp;#x1F6CC;</code><td><code>&amp;#128716;</code><td>
<tr><td>&#x1F6CF;<td><td>BED<td><code>&amp;#x1F6CF;</code><td><code>&amp;#128719;</code><td>
<tr><td>&#x1F6CB;<td><td>COUCH AND LAMP<td><code>&amp;#x1F6CB;</code><td><code>&amp;#128715;</code><td>
<tr><td>&#x1F6BD;<td><td>TOILET<td><code>&amp;#x1F6BD;</code><td><code>&amp;#128701;</code><td>
<tr><td>&#x1F6BF;<td><td>SHOWER<td><code>&amp;#x1F6BF;</code><td><code>&amp;#128703;</code><td>
<tr><td>&#x1F6C0;<td><td>BATH &#x224A; person taking bath<td><code>&amp;#x1F6C0;</code><td><code>&amp;#128704;</code><td>
<tr><td>&#x1F6C0;&#x1F3FB;<td><td>person taking bath, type-1-2<td><code>&amp;#x1F6C0; &amp;#x1F3FB;</code><td><code>&amp;#128704; &amp;#127995;</code><td>
<tr><td>&#x1F6C0;&#x1F3FC;<td><td>person taking bath, type-3<td><code>&amp;#x1F6C0; &amp;#x1F3FC;</code><td><code>&amp;#128704; &amp;#127996;</code><td>
<tr><td>&#x1F6C0;&#x1F3FD;<td><td>person taking bath, type-4<td><code>&amp;#x1F6C0; &amp;#x1F3FD;</code><td><code>&amp;#128704; &amp;#127997;</code><td>
<tr><td>&#x1F6C0;&#x1F3FE;<td><td>person taking bath, type-5<td><code>&amp;#x1F6C0; &amp;#x1F3FE;</code><td><code>&amp;#128704; &amp;#127998;</code><td>
<tr><td>&#x1F6C0;&#x1F3FF;<td><td>person taking bath, type-6<td><code>&amp;#x1F6C0; &amp;#x1F3FF;</code><td><code>&amp;#128704; &amp;#127999;</code><td>
<tr><td>&#x1F6C1;<td><td>BATHTUB<td><code>&amp;#x1F6C1;</code><td><code>&amp;#128705;</code><td>
<tr><td>&#x231B;<td><td>HOURGLASS<td><code>&amp;#x231B;</code><td><code>&amp;#8987;</code><td>
<tr><td>&#x23F3;<td><td>HOURGLASS WITH FLOWING SAND<td><code>&amp;#x23F3;</code><td><code>&amp;#9203;</code><td>
<tr><td>&#x231A;<td><td>WATCH<td><code>&amp;#x231A;</code><td><code>&amp;#8986;</code><td>
<tr><td>&#x23F0;<td><td>ALARM CLOCK<td><code>&amp;#x23F0;</code><td><code>&amp;#9200;</code><td>
<tr><td>&#x23F1;<td><td>STOPWATCH<td><code>&amp;#x23F1;</code><td><code>&amp;#9201;</code><td>
<tr><td>&#x23F2;<td><td>TIMER CLOCK<td><code>&amp;#x23F2;</code><td><code>&amp;#9202;</code><td>
<tr><td>&#x1F570;<td><td>MANTELPIECE CLOCK<td><code>&amp;#x1F570;</code><td><code>&amp;#128368;</code><td>
<tr><td>&#x1F55B;<td><td>CLOCK FACE TWELVE OCLOCK &#x224A; twelve oâ€™clock<td><code>&amp;#x1F55B;</code><td><code>&amp;#128347;</code><td>
<tr><td>&#x1F567;<td><td>CLOCK FACE TWELVE-THIRTY &#x224A; twelve-thirty<td><code>&amp;#x1F567;</code><td><code>&amp;#128359;</code><td>
<tr><td>&#x1F550;<td><td>CLOCK FACE ONE OCLOCK &#x224A; one oâ€™clock<td><code>&amp;#x1F550;</code><td><code>&amp;#128336;</code><td>
<tr><td>&#x1F55C;<td><td>CLOCK FACE ONE-THIRTY &#x224A; one-thirty<td><code>&amp;#x1F55C;</code><td><code>&amp;#128348;</code><td>
<tr><td>&#x1F551;<td><td>CLOCK FACE TWO OCLOCK &#x224A; two oâ€™clock<td><code>&amp;#x1F551;</code><td><code>&amp;#128337;</code><td>
<tr><td>&#x1F55D;<td><td>CLOCK FACE TWO-THIRTY &#x224A; two-thirty<td><code>&amp;#x1F55D;</code><td><code>&amp;#128349;</code><td>
<tr><td>&#x1F552;<td><td>CLOCK FACE THREE OCLOCK &#x224A; three oâ€™clock<td><code>&amp;#x1F552;</code><td><code>&amp;#128338;</code><td>
<tr><td>&#x1F55E;<td><td>CLOCK FACE THREE-THIRTY &#x224A; three-thirty<td><code>&amp;#x1F55E;</code><td><code>&amp;#128350;</code><td>
<tr><td>&#x1F553;<td><td>CLOCK FACE FOUR OCLOCK &#x224A; four oâ€™clock<td><code>&amp;#x1F553;</code><td><code>&amp;#128339;</code><td>
<tr><td>&#x1F55F;<td><td>CLOCK FACE FOUR-THIRTY &#x224A; four-thirty<td><code>&amp;#x1F55F;</code><td><code>&amp;#128351;</code><td>
<tr><td>&#x1F554;<td><td>CLOCK FACE FIVE OCLOCK &#x224A; five oâ€™clock<td><code>&amp;#x1F554;</code><td><code>&amp;#128340;</code><td>
<tr><td>&#x1F560;<td><td>CLOCK FACE FIVE-THIRTY &#x224A; five-thirty<td><code>&amp;#x1F560;</code><td><code>&amp;#128352;</code><td>
<tr><td>&#x1F555;<td><td>CLOCK FACE SIX OCLOCK &#x224A; six oâ€™clock<td><code>&amp;#x1F555;</code><td><code>&amp;#128341;</code><td>
<tr><td>&#x1F561;<td><td>CLOCK FACE SIX-THIRTY &#x224A; six-thirty<td><code>&amp;#x1F561;</code><td><code>&amp;#128353;</code><td>
<tr><td>&#x1F556;<td><td>CLOCK FACE SEVEN OCLOCK &#x224A; seven oâ€™clock<td><code>&amp;#x1F556;</code><td><code>&amp;#128342;</code><td>
<tr><td>&#x1F562;<td><td>CLOCK FACE SEVEN-THIRTY &#x224A; seven-thirty<td><code>&amp;#x1F562;</code><td><code>&amp;#128354;</code><td>
<tr><td>&#x1F557;<td><td>CLOCK FACE EIGHT OCLOCK &#x224A; eight oâ€™clock<td><code>&amp;#x1F557;</code><td><code>&amp;#128343;</code><td>
<tr><td>&#x1F563;<td><td>CLOCK FACE EIGHT-THIRTY &#x224A; eight-thirty<td><code>&amp;#x1F563;</code><td><code>&amp;#128355;</code><td>
<tr><td>&#x1F558;<td><td>CLOCK FACE NINE OCLOCK &#x224A; nine oâ€™clock<td><code>&amp;#x1F558;</code><td><code>&amp;#128344;</code><td>
<tr><td>&#x1F564;<td><td>CLOCK FACE NINE-THIRTY &#x224A; nine-thirty<td><code>&amp;#x1F564;</code><td><code>&amp;#128356;</code><td>
<tr><td>&#x1F559;<td><td>CLOCK FACE TEN OCLOCK &#x224A; ten oâ€™clock<td><code>&amp;#x1F559;</code><td><code>&amp;#128345;</code><td>
<tr><td>&#x1F565;<td><td>CLOCK FACE TEN-THIRTY &#x224A; ten-thirty<td><code>&amp;#x1F565;</code><td><code>&amp;#128357;</code><td>
<tr><td>&#x1F55A;<td><td>CLOCK FACE ELEVEN OCLOCK &#x224A; eleven oâ€™clock<td><code>&amp;#x1F55A;</code><td><code>&amp;#128346;</code><td>
<tr><td>&#x1F566;<td><td>CLOCK FACE ELEVEN-THIRTY &#x224A; eleven-thirty<td><code>&amp;#x1F566;</code><td><code>&amp;#128358;</code><td>
<tr><td>&#x1F311;<td><td>NEW MOON SYMBOL &#x224A; new moon<td><code>&amp;#x1F311;</code><td><code>&amp;#127761;</code><td>
<tr><td>&#x1F312;<td><td>WAXING CRESCENT MOON SYMBOL &#x224A; waxing crescent moon<td><code>&amp;#x1F312;</code><td><code>&amp;#127762;</code><td>
<tr><td>&#x1F313;<td><td>FIRST QUARTER MOON SYMBOL &#x224A; first quarter moon<td><code>&amp;#x1F313;</code><td><code>&amp;#127763;</code><td>
<tr><td>&#x1F314;<td><td>WAXING GIBBOUS MOON SYMBOL &#x224A; waxing gibbous moon<td><code>&amp;#x1F314;</code><td><code>&amp;#127764;</code><td>
<tr><td>&#x1F315;<td><td>FULL MOON SYMBOL &#x224A; full moon<td><code>&amp;#x1F315;</code><td><code>&amp;#127765;</code><td>
<tr><td>&#x1F316;<td><td>WANING GIBBOUS MOON SYMBOL &#x224A; waning gibbous moon<td><code>&amp;#x1F316;</code><td><code>&amp;#127766;</code><td>
<tr><td>&#x1F317;<td><td>LAST QUARTER MOON SYMBOL &#x224A; last quarter moon<td><code>&amp;#x1F317;</code><td><code>&amp;#127767;</code><td>
<tr><td>&#x1F318;<td><td>WANING CRESCENT MOON SYMBOL &#x224A; waning crescent moon<td><code>&amp;#x1F318;</code><td><code>&amp;#127768;</code><td>
<tr><td>&#x1F319;<td><td>CRESCENT MOON<td><code>&amp;#x1F319;</code><td><code>&amp;#127769;</code><td>
<tr><td>&#x1F31A;<td><td>NEW MOON WITH FACE &#x224A; new moon face<td><code>&amp;#x1F31A;</code><td><code>&amp;#127770;</code><td>
<tr><td>&#x1F31B;<td><td>FIRST QUARTER MOON WITH FACE<td><code>&amp;#x1F31B;</code><td><code>&amp;#127771;</code><td>
<tr><td>&#x1F31C;<td><td>LAST QUARTER MOON WITH FACE<td><code>&amp;#x1F31C;</code><td><code>&amp;#127772;</code><td>
<tr><td>&#x1F321;<td><td>THERMOMETER<td><code>&amp;#x1F321;</code><td><code>&amp;#127777;</code><td>
<tr><td>&#x2600;<td><td>BLACK SUN WITH RAYS &#x224A; sun<td><code>&amp;#x2600;</code><td><code>&amp;#9728;</code><td>
<tr><td>&#x1F31D;<td><td>FULL MOON WITH FACE<td><code>&amp;#x1F31D;</code><td><code>&amp;#127773;</code><td>
<tr><td>&#x1F31E;<td><td>SUN WITH FACE<td><code>&amp;#x1F31E;</code><td><code>&amp;#127774;</code><td>
<tr><td>&#x2B50;<td><td>WHITE MEDIUM STAR<td><code>&amp;#x2B50;</code><td><code>&amp;#11088;</code><td>
<tr><td>&#x1F31F;<td><td>GLOWING STAR<td><code>&amp;#x1F31F;</code><td><code>&amp;#127775;</code><td>
<tr><td>&#x1F320;<td><td>SHOOTING STAR<td><code>&amp;#x1F320;</code><td><code>&amp;#127776;</code><td>
<tr><td>&#x2601;<td><td>CLOUD<td><code>&amp;#x2601;</code><td><code>&amp;#9729;</code><td>
<tr><td>&#x26C5;<td><td>SUN BEHIND CLOUD<td><code>&amp;#x26C5;</code><td><code>&amp;#9925;</code><td>
<tr><td>&#x26C8;<td><td>THUNDER CLOUD AND RAIN &#x224A; cloud with lightning and rain<td><code>&amp;#x26C8;</code><td><code>&amp;#9928;</code><td>
<tr><td>&#x1F324;<td><td>WHITE SUN WITH SMALL CLOUD &#x224A; sun behind small cloud<td><code>&amp;#x1F324;</code><td><code>&amp;#127780;</code><td>
<tr><td>&#x1F325;<td><td>WHITE SUN BEHIND CLOUD &#x224A; sun behind large cloud<td><code>&amp;#x1F325;</code><td><code>&amp;#127781;</code><td>
<tr><td>&#x1F326;<td><td>WHITE SUN BEHIND CLOUD WITH RAIN &#x224A; sun behind cloud with rain<td><code>&amp;#x1F326;</code><td><code>&amp;#127782;</code><td>
<tr><td>&#x1F327;<td><td>CLOUD WITH RAIN<td><code>&amp;#x1F327;</code><td><code>&amp;#127783;</code><td>
<tr><td>&#x1F328;<td><td>CLOUD WITH SNOW<td><code>&amp;#x1F328;</code><td><code>&amp;#127784;</code><td>
<tr><td>&#x1F329;<td><td>CLOUD WITH LIGHTNING<td><code>&amp;#x1F329;</code><td><code>&amp;#127785;</code><td>
<tr><td>&#x1F32A;<td><td>CLOUD WITH TORNADO &#x224A; tornado<td><code>&amp;#x1F32A;</code><td><code>&amp;#127786;</code><td>
<tr><td>&#x1F32B;<td><td>FOG<td><code>&amp;#x1F32B;</code><td><code>&amp;#127787;</code><td>
<tr><td>&#x1F32C;<td><td>WIND BLOWING FACE &#x224A; wind face<td><code>&amp;#x1F32C;</code><td><code>&amp;#127788;</code><td>
<tr><td>&#x1F300;<td><td>CYCLONE<td><code>&amp;#x1F300;</code><td><code>&amp;#127744;</code><td>
<tr><td>&#x1F308;<td><td>RAINBOW<td><code>&amp;#x1F308;</code><td><code>&amp;#127752;</code><td>
<tr><td>&#x1F302;<td><td>CLOSED UMBRELLA<td><code>&amp;#x1F302;</code><td><code>&amp;#127746;</code><td>
<tr><td>&#x2602;<td><td>UMBRELLA<td><code>&amp;#x2602;</code><td><code>&amp;#9730;</code><td>
<tr><td>&#x2614;<td><td>UMBRELLA WITH RAIN DROPS<td><code>&amp;#x2614;</code><td><code>&amp;#9748;</code><td>
<tr><td>&#x26F1;<td><td>UMBRELLA ON GROUND<td><code>&amp;#x26F1;</code><td><code>&amp;#9969;</code><td>
<tr><td>&#x26A1;<td><td>HIGH VOLTAGE SIGN &#x224A; high voltage<td><code>&amp;#x26A1;</code><td><code>&amp;#9889;</code><td>
<tr><td>&#x2744;<td><td>SNOWFLAKE<td><code>&amp;#x2744;</code><td><code>&amp;#10052;</code><td>
<tr><td>&#x2603;<td><td>SNOWMAN<td><code>&amp;#x2603;</code><td><code>&amp;#9731;</code><td>
<tr><td>&#x26C4;<td><td>SNOWMAN WITHOUT SNOW<td><code>&amp;#x26C4;</code><td><code>&amp;#9924;</code><td>
<tr><td>&#x2604;<td><td>COMET<td><code>&amp;#x2604;</code><td><code>&amp;#9732;</code><td>
<tr><td>&#x1F525;<td><td>FIRE<td><code>&amp;#x1F525;</code><td><code>&amp;#128293;</code><td>
<tr><td>&#x1F4A7;<td><td>DROPLET<td><code>&amp;#x1F4A7;</code><td><code>&amp;#128167;</code><td>
<tr><td>&#x1F30A;<td><td>WATER WAVE<td><code>&amp;#x1F30A;</code><td><code>&amp;#127754;</code><td>
<tr><td>&#x1F383;<td><td>JACK-O-LANTERN<td><code>&amp;#x1F383;</code><td><code>&amp;#127875;</code><td>
<tr><td>&#x1F384;<td><td>CHRISTMAS TREE<td><code>&amp;#x1F384;</code><td><code>&amp;#127876;</code><td>
<tr><td>&#x1F386;<td><td>FIREWORKS<td><code>&amp;#x1F386;</code><td><code>&amp;#127878;</code><td>
<tr><td>&#x1F387;<td><td>FIREWORK SPARKLER &#x224A; sparkler<td><code>&amp;#x1F387;</code><td><code>&amp;#127879;</code><td>
<tr><td>&#x2728;<td><td>SPARKLES<td><code>&amp;#x2728;</code><td><code>&amp;#10024;</code><td>
<tr><td>&#x1F388;<td><td>BALLOON<td><code>&amp;#x1F388;</code><td><code>&amp;#127880;</code><td>
<tr><td>&#x1F389;<td><td>PARTY POPPER<td><code>&amp;#x1F389;</code><td><code>&amp;#127881;</code><td>
<tr><td>&#x1F38A;<td><td>CONFETTI BALL<td><code>&amp;#x1F38A;</code><td><code>&amp;#127882;</code><td>
<tr><td>&#x1F38B;<td><td>TANABATA TREE<td><code>&amp;#x1F38B;</code><td><code>&amp;#127883;</code><td>
<tr><td>&#x1F38D;<td><td>PINE DECORATION<td><code>&amp;#x1F38D;</code><td><code>&amp;#127885;</code><td>
<tr><td>&#x1F38E;<td><td>JAPANESE DOLLS<td><code>&amp;#x1F38E;</code><td><code>&amp;#127886;</code><td>
<tr><td>&#x1F38F;<td><td>CARP STREAMER<td><code>&amp;#x1F38F;</code><td><code>&amp;#127887;</code><td>
<tr><td>&#x1F390;<td><td>WIND CHIME<td><code>&amp;#x1F390;</code><td><code>&amp;#127888;</code><td>
<tr><td>&#x1F391;<td><td>MOON VIEWING CEREMONY &#x224A; moon ceremony<td><code>&amp;#x1F391;</code><td><code>&amp;#127889;</code><td>
<tr><td>&#x1F380;<td><td>RIBBON<td><code>&amp;#x1F380;</code><td><code>&amp;#127872;</code><td>
<tr><td>&#x1F381;<td><td>WRAPPED PRESENT<td><code>&amp;#x1F381;</code><td><code>&amp;#127873;</code><td>
<tr><td>&#x1F397;<td><td>REMINDER RIBBON<td><code>&amp;#x1F397;</code><td><code>&amp;#127895;</code><td>
<tr><td>&#x1F39F;<td><td>ADMISSION TICKETS<td><code>&amp;#x1F39F;</code><td><code>&amp;#127903;</code><td>
<tr><td>&#x1F3AB;<td><td>TICKET<td><code>&amp;#x1F3AB;</code><td><code>&amp;#127915;</code><td>
<tr><td>&#x1F396;<td><td>MILITARY MEDAL<td><code>&amp;#x1F396;</code><td><code>&amp;#127894;</code><td>
<tr><td>&#x1F3C6;<td><td>TROPHY<td><code>&amp;#x1F3C6;</code><td><code>&amp;#127942;</code><td>
<tr><td>&#x1F3C5;<td><td>SPORTS MEDAL<td><code>&amp;#x1F3C5;</code><td><code>&amp;#127941;</code><td>
<tr><td>&#x1F947;<td><td>FIRST PLACE MEDAL &#x224A; 1st place medal<td><code>&amp;#x1F947;</code><td><code>&amp;#129351;</code><td>
<tr><td>&#x1F948;<td><td>SECOND PLACE MEDAL &#x224A; 2nd place medal<td><code>&amp;#x1F948;</code><td><code>&amp;#129352;</code><td>
<tr><td>&#x1F949;<td><td>THIRD PLACE MEDAL &#x224A; 3rd place medal<td><code>&amp;#x1F949;</code><td><code>&amp;#129353;</code><td>
<tr><td>&#x26BD;<td><td>SOCCER BALL<td><code>&amp;#x26BD;</code><td><code>&amp;#9917;</code><td>
<tr><td>&#x26BE;<td><td>BASEBALL<td><code>&amp;#x26BE;</code><td><code>&amp;#9918;</code><td>
<tr><td>&#x1F3C0;<td><td>BASKETBALL AND HOOP &#x224A; basketball<td><code>&amp;#x1F3C0;</code><td><code>&amp;#127936;</code><td>
<tr><td>&#x1F3D0;<td><td>VOLLEYBALL<td><code>&amp;#x1F3D0;</code><td><code>&amp;#127952;</code><td>
<tr><td>&#x1F3C8;<td><td>AMERICAN FOOTBALL<td><code>&amp;#x1F3C8;</code><td><code>&amp;#127944;</code><td>
<tr><td>&#x1F3C9;<td><td>RUGBY FOOTBALL<td><code>&amp;#x1F3C9;</code><td><code>&amp;#127945;</code><td>
<tr><td>&#x1F3BE;<td><td>TENNIS RACQUET AND BALL &#x224A; tennis<td><code>&amp;#x1F3BE;</code><td><code>&amp;#127934;</code><td>
<tr><td>&#x1F3B1;<td><td>BILLIARDS<td><code>&amp;#x1F3B1;</code><td><code>&amp;#127921;</code><td>
<tr><td>&#x1F3B3;<td><td>BOWLING<td><code>&amp;#x1F3B3;</code><td><code>&amp;#127923;</code><td>
<tr><td>&#x1F3CF;<td><td>CRICKET BAT AND BALL &#x224A; cricket<td><code>&amp;#x1F3CF;</code><td><code>&amp;#127951;</code><td>
<tr><td>&#x1F3D1;<td><td>FIELD HOCKEY STICK AND BALL &#x224A; field hockey<td><code>&amp;#x1F3D1;</code><td><code>&amp;#127953;</code><td>
<tr><td>&#x1F3D2;<td><td>ICE HOCKEY STICK AND PUCK<td><code>&amp;#x1F3D2;</code><td><code>&amp;#127954;</code><td>
<tr><td>&#x1F3D3;<td><td>TABLE TENNIS PADDLE AND BALL &#x224A; ping pong<td><code>&amp;#x1F3D3;</code><td><code>&amp;#127955;</code><td>
<tr><td>&#x1F3F8;<td><td>BADMINTON RACQUET AND SHUTTLECOCK &#x224A; badminton<td><code>&amp;#x1F3F8;</code><td><code>&amp;#127992;</code><td>
<tr><td>&#x1F94A;<td><td>BOXING GLOVE<td><code>&amp;#x1F94A;</code><td><code>&amp;#129354;</code><td>
<tr><td>&#x1F94B;<td><td>MARTIAL ARTS UNIFORM<td><code>&amp;#x1F94B;</code><td><code>&amp;#129355;</code><td>
<tr><td>&#x1F945;<td><td>GOAL NET<td><code>&amp;#x1F945;</code><td><code>&amp;#129349;</code><td>
<tr><td>&#x1F3AF;<td><td>DIRECT HIT<td><code>&amp;#x1F3AF;</code><td><code>&amp;#127919;</code><td>
<tr><td>&#x26F3;<td><td>FLAG IN HOLE<td><code>&amp;#x26F3;</code><td><code>&amp;#9971;</code><td>
<tr><td>&#x26F8;<td><td>ICE SKATE<td><code>&amp;#x26F8;</code><td><code>&amp;#9976;</code><td>
<tr><td>&#x1F3A3;<td><td>FISHING POLE AND FISH &#x224A; fishing pole<td><code>&amp;#x1F3A3;</code><td><code>&amp;#127907;</code><td>
<tr><td>&#x1F3BD;<td><td>RUNNING SHIRT WITH SASH &#x224A; running shirt<td><code>&amp;#x1F3BD;</code><td><code>&amp;#127933;</code><td>
<tr><td>&#x1F3BF;<td><td>SKI AND SKI BOOT &#x224A; skis<td><code>&amp;#x1F3BF;</code><td><code>&amp;#127935;</code><td>
<tr><td>&#x1F93A;<td><td>FENCER &#x224A; person fencing<td><code>&amp;#x1F93A;</code><td><code>&amp;#129338;</code><td>
<tr><td>&#x1F3C7;<td><td>HORSE RACING<td><code>&amp;#x1F3C7;</code><td><code>&amp;#127943;</code><td>
<tr><td>&#x26F7;<td><td>SKIER<td><code>&amp;#x26F7;</code><td><code>&amp;#9975;</code><td>
<tr><td>&#x1F3C2;<td><td>SNOWBOARDER<td><code>&amp;#x1F3C2;</code><td><code>&amp;#127938;</code><td>
<tr><td>&#x1F3CC;<td><td>GOLFER &#x224A; person golfing<td><code>&amp;#x1F3CC;</code><td><code>&amp;#127948;</code><td>
<tr><td>&#x1F3C4;<td><td>SURFER &#x224A; person surfing<td><code>&amp;#x1F3C4;</code><td><code>&amp;#127940;</code><td>
<tr><td>&#x1F3C4;&#x1F3FB;<td><td>person surfing, type-1-2<td><code>&amp;#x1F3C4; &amp;#x1F3FB;</code><td><code>&amp;#127940; &amp;#127995;</code><td>
<tr><td>&#x1F3C4;&#x1F3FC;<td><td>person surfing, type-3<td><code>&amp;#x1F3C4; &amp;#x1F3FC;</code><td><code>&amp;#127940; &amp;#127996;</code><td>
<tr><td>&#x1F3C4;&#x1F3FD;<td><td>person surfing, type-4<td><code>&amp;#x1F3C4; &amp;#x1F3FD;</code><td><code>&amp;#127940; &amp;#127997;</code><td>
<tr><td>&#x1F3C4;&#x1F3FE;<td><td>person surfing, type-5<td><code>&amp;#x1F3C4; &amp;#x1F3FE;</code><td><code>&amp;#127940; &amp;#127998;</code><td>
<tr><td>&#x1F3C4;&#x1F3FF;<td><td>person surfing, type-6<td><code>&amp;#x1F3C4; &amp;#x1F3FF;</code><td><code>&amp;#127940; &amp;#127999;</code><td>
<tr><td>&#x1F6A3;<td><td>ROWBOAT &#x224A; person rowing boat<td><code>&amp;#x1F6A3;</code><td><code>&amp;#128675;</code><td>
<tr><td>&#x1F6A3;&#x1F3FB;<td><td>person rowing boat, type-1-2<td><code>&amp;#x1F6A3; &amp;#x1F3FB;</code><td><code>&amp;#128675; &amp;#127995;</code><td>
<tr><td>&#x1F6A3;&#x1F3FC;<td><td>person rowing boat, type-3<td><code>&amp;#x1F6A3; &amp;#x1F3FC;</code><td><code>&amp;#128675; &amp;#127996;</code><td>
<tr><td>&#x1F6A3;&#x1F3FD;<td><td>person rowing boat, type-4<td><code>&amp;#x1F6A3; &amp;#x1F3FD;</code><td><code>&amp;#128675; &amp;#127997;</code><td>
<tr><td>&#x1F6A3;&#x1F3FE;<td><td>person rowing boat, type-5<td><code>&amp;#x1F6A3; &amp;#x1F3FE;</code><td><code>&amp;#128675; &amp;#127998;</code><td>
<tr><td>&#x1F6A3;&#x1F3FF;<td><td>person rowing boat, type-6<td><code>&amp;#x1F6A3; &amp;#x1F3FF;</code><td><code>&amp;#128675; &amp;#127999;</code><td>
<tr><td>&#x1F3CA;<td><td>SWIMMER &#x224A; person swimming<td><code>&amp;#x1F3CA;</code><td><code>&amp;#127946;</code><td>
<tr><td>&#x1F3CA;&#x1F3FB;<td><td>person swimming, type-1-2<td><code>&amp;#x1F3CA; &amp;#x1F3FB;</code><td><code>&amp;#127946; &amp;#127995;</code><td>
<tr><td>&#x1F3CA;&#x1F3FC;<td><td>person swimming, type-3<td><code>&amp;#x1F3CA; &amp;#x1F3FC;</code><td><code>&amp;#127946; &amp;#127996;</code><td>
<tr><td>&#x1F3CA;&#x1F3FD;<td><td>person swimming, type-4<td><code>&amp;#x1F3CA; &amp;#x1F3FD;</code><td><code>&amp;#127946; &amp;#127997;</code><td>
<tr><td>&#x1F3CA;&#x1F3FE;<td><td>person swimming, type-5<td><code>&amp;#x1F3CA; &amp;#x1F3FE;</code><td><code>&amp;#127946; &amp;#127998;</code><td>
<tr><td>&#x1F3CA;&#x1F3FF;<td><td>person swimming, type-6<td><code>&amp;#x1F3CA; &amp;#x1F3FF;</code><td><code>&amp;#127946; &amp;#127999;</code><td>
<tr><td>&#x26F9;<td><td>PERSON WITH BALL<td><code>&amp;#x26F9;</code><td><code>&amp;#9977;</code><td>
<tr><td>&#x26F9;&#x1F3FB;<td><td>person with ball, type-1-2<td><code>&amp;#x26F9; &amp;#x1F3FB;</code><td><code>&amp;#9977; &amp;#127995;</code><td>
<tr><td>&#x26F9;&#x1F3FC;<td><td>person with ball, type-3<td><code>&amp;#x26F9; &amp;#x1F3FC;</code><td><code>&amp;#9977; &amp;#127996;</code><td>
<tr><td>&#x26F9;&#x1F3FD;<td><td>person with ball, type-4<td><code>&amp;#x26F9; &amp;#x1F3FD;</code><td><code>&amp;#9977; &amp;#127997;</code><td>
<tr><td>&#x26F9;&#x1F3FE;<td><td>person with ball, type-5<td><code>&amp;#x26F9; &amp;#x1F3FE;</code><td><code>&amp;#9977; &amp;#127998;</code><td>
<tr><td>&#x26F9;&#x1F3FF;<td><td>person with ball, type-6<td><code>&amp;#x26F9; &amp;#x1F3FF;</code><td><code>&amp;#9977; &amp;#127999;</code><td>
<tr><td>&#x1F3CB;<td><td>WEIGHT LIFTER &#x224A; person weight lifting<td><code>&amp;#x1F3CB;</code><td><code>&amp;#127947;</code><td>
<tr><td>&#x1F3CB;&#x1F3FB;<td><td>person weight lifting, type-1-2<td><code>&amp;#x1F3CB; &amp;#x1F3FB;</code><td><code>&amp;#127947; &amp;#127995;</code><td>
<tr><td>&#x1F3CB;&#x1F3FC;<td><td>person weight lifting, type-3<td><code>&amp;#x1F3CB; &amp;#x1F3FC;</code><td><code>&amp;#127947; &amp;#127996;</code><td>
<tr><td>&#x1F3CB;&#x1F3FD;<td><td>person weight lifting, type-4<td><code>&amp;#x1F3CB; &amp;#x1F3FD;</code><td><code>&amp;#127947; &amp;#127997;</code><td>
<tr><td>&#x1F3CB;&#x1F3FE;<td><td>person weight lifting, type-5<td><code>&amp;#x1F3CB; &amp;#x1F3FE;</code><td><code>&amp;#127947; &amp;#127998;</code><td>
<tr><td>&#x1F3CB;&#x1F3FF;<td><td>person weight lifting, type-6<td><code>&amp;#x1F3CB; &amp;#x1F3FF;</code><td><code>&amp;#127947; &amp;#127999;</code><td>
<tr><td>&#x1F6B4;<td><td>BICYCLIST &#x224A; person biking<td><code>&amp;#x1F6B4;</code><td><code>&amp;#128692;</code><td>
<tr><td>&#x1F6B4;&#x1F3FB;<td><td>person biking, type-1-2<td><code>&amp;#x1F6B4; &amp;#x1F3FB;</code><td><code>&amp;#128692; &amp;#127995;</code><td>
<tr><td>&#x1F6B4;&#x1F3FC;<td><td>person biking, type-3<td><code>&amp;#x1F6B4; &amp;#x1F3FC;</code><td><code>&amp;#128692; &amp;#127996;</code><td>
<tr><td>&#x1F6B4;&#x1F3FD;<td><td>person biking, type-4<td><code>&amp;#x1F6B4; &amp;#x1F3FD;</code><td><code>&amp;#128692; &amp;#127997;</code><td>
<tr><td>&#x1F6B4;&#x1F3FE;<td><td>person biking, type-5<td><code>&amp;#x1F6B4; &amp;#x1F3FE;</code><td><code>&amp;#128692; &amp;#127998;</code><td>
<tr><td>&#x1F6B4;&#x1F3FF;<td><td>person biking, type-6<td><code>&amp;#x1F6B4; &amp;#x1F3FF;</code><td><code>&amp;#128692; &amp;#127999;</code><td>
<tr><td>&#x1F6B5;<td><td>MOUNTAIN BICYCLIST &#x224A; person mountain biking<td><code>&amp;#x1F6B5;</code><td><code>&amp;#128693;</code><td>
<tr><td>&#x1F6B5;&#x1F3FB;<td><td>person mountain biking, type-1-2<td><code>&amp;#x1F6B5; &amp;#x1F3FB;</code><td><code>&amp;#128693; &amp;#127995;</code><td>
<tr><td>&#x1F6B5;&#x1F3FC;<td><td>person mountain biking, type-3<td><code>&amp;#x1F6B5; &amp;#x1F3FC;</code><td><code>&amp;#128693; &amp;#127996;</code><td>
<tr><td>&#x1F6B5;&#x1F3FD;<td><td>person mountain biking, type-4<td><code>&amp;#x1F6B5; &amp;#x1F3FD;</code><td><code>&amp;#128693; &amp;#127997;</code><td>
<tr><td>&#x1F6B5;&#x1F3FE;<td><td>person mountain biking, type-5<td><code>&amp;#x1F6B5; &amp;#x1F3FE;</code><td><code>&amp;#128693; &amp;#127998;</code><td>
<tr><td>&#x1F6B5;&#x1F3FF;<td><td>person mountain biking, type-6<td><code>&amp;#x1F6B5; &amp;#x1F3FF;</code><td><code>&amp;#128693; &amp;#127999;</code><td>
<tr><td>&#x1F938;<td><td>PERSON DOING CARTWHEEL<td><code>&amp;#x1F938;</code><td><code>&amp;#129336;</code><td>
<tr><td>&#x1F938;&#x1F3FB;<td><td>person doing cartwheel, type-1-2<td><code>&amp;#x1F938; &amp;#x1F3FB;</code><td><code>&amp;#129336; &amp;#127995;</code><td>
<tr><td>&#x1F938;&#x1F3FC;<td><td>person doing cartwheel, type-3<td><code>&amp;#x1F938; &amp;#x1F3FC;</code><td><code>&amp;#129336; &amp;#127996;</code><td>
<tr><td>&#x1F938;&#x1F3FD;<td><td>person doing cartwheel, type-4<td><code>&amp;#x1F938; &amp;#x1F3FD;</code><td><code>&amp;#129336; &amp;#127997;</code><td>
<tr><td>&#x1F938;&#x1F3FE;<td><td>person doing cartwheel, type-5<td><code>&amp;#x1F938; &amp;#x1F3FE;</code><td><code>&amp;#129336; &amp;#127998;</code><td>
<tr><td>&#x1F938;&#x1F3FF;<td><td>person doing cartwheel, type-6<td><code>&amp;#x1F938; &amp;#x1F3FF;</code><td><code>&amp;#129336; &amp;#127999;</code><td>
<tr><td>&#x1F93C;<td><td>WRESTLERS &#x224A; people wrestling<td><code>&amp;#x1F93C;</code><td><code>&amp;#129340;</code><td>
<tr><td>&#x1F93C;&#x1F3FB;<td><td>people wrestling, type-1-2<td><code>&amp;#x1F93C; &amp;#x1F3FB;</code><td><code>&amp;#129340; &amp;#127995;</code><td>
<tr><td>&#x1F93C;&#x1F3FC;<td><td>people wrestling, type-3<td><code>&amp;#x1F93C; &amp;#x1F3FC;</code><td><code>&amp;#129340; &amp;#127996;</code><td>
<tr><td>&#x1F93C;&#x1F3FD;<td><td>people wrestling, type-4<td><code>&amp;#x1F93C; &amp;#x1F3FD;</code><td><code>&amp;#129340; &amp;#127997;</code><td>
<tr><td>&#x1F93C;&#x1F3FE;<td><td>people wrestling, type-5<td><code>&amp;#x1F93C; &amp;#x1F3FE;</code><td><code>&amp;#129340; &amp;#127998;</code><td>
<tr><td>&#x1F93C;&#x1F3FF;<td><td>people wrestling, type-6<td><code>&amp;#x1F93C; &amp;#x1F3FF;</code><td><code>&amp;#129340; &amp;#127999;</code><td>
<tr><td>&#x1F93D;<td><td>WATER POLO &#x224A; person playing water polo<td><code>&amp;#x1F93D;</code><td><code>&amp;#129341;</code><td>
<tr><td>&#x1F93D;&#x1F3FB;<td><td>person playing water polo, type-1-2<td><code>&amp;#x1F93D; &amp;#x1F3FB;</code><td><code>&amp;#129341; &amp;#127995;</code><td>
<tr><td>&#x1F93D;&#x1F3FC;<td><td>person playing water polo, type-3<td><code>&amp;#x1F93D; &amp;#x1F3FC;</code><td><code>&amp;#129341; &amp;#127996;</code><td>
<tr><td>&#x1F93D;&#x1F3FD;<td><td>person playing water polo, type-4<td><code>&amp;#x1F93D; &amp;#x1F3FD;</code><td><code>&amp;#129341; &amp;#127997;</code><td>
<tr><td>&#x1F93D;&#x1F3FE;<td><td>person playing water polo, type-5<td><code>&amp;#x1F93D; &amp;#x1F3FE;</code><td><code>&amp;#129341; &amp;#127998;</code><td>
<tr><td>&#x1F93D;&#x1F3FF;<td><td>person playing water polo, type-6<td><code>&amp;#x1F93D; &amp;#x1F3FF;</code><td><code>&amp;#129341; &amp;#127999;</code><td>
<tr><td>&#x1F93E;<td><td>HANDBALL &#x224A; person playing handball<td><code>&amp;#x1F93E;</code><td><code>&amp;#129342;</code><td>
<tr><td>&#x1F93E;&#x1F3FB;<td><td>person playing handball, type-1-2<td><code>&amp;#x1F93E; &amp;#x1F3FB;</code><td><code>&amp;#129342; &amp;#127995;</code><td>
<tr><td>&#x1F93E;&#x1F3FC;<td><td>person playing handball, type-3<td><code>&amp;#x1F93E; &amp;#x1F3FC;</code><td><code>&amp;#129342; &amp;#127996;</code><td>
<tr><td>&#x1F93E;&#x1F3FD;<td><td>person playing handball, type-4<td><code>&amp;#x1F93E; &amp;#x1F3FD;</code><td><code>&amp;#129342; &amp;#127997;</code><td>
<tr><td>&#x1F93E;&#x1F3FE;<td><td>person playing handball, type-5<td><code>&amp;#x1F93E; &amp;#x1F3FE;</code><td><code>&amp;#129342; &amp;#127998;</code><td>
<tr><td>&#x1F93E;&#x1F3FF;<td><td>person playing handball, type-6<td><code>&amp;#x1F93E; &amp;#x1F3FF;</code><td><code>&amp;#129342; &amp;#127999;</code><td>
<tr><td>&#x1F939;<td><td>JUGGLING &#x224A; person juggling<td><code>&amp;#x1F939;</code><td><code>&amp;#129337;</code><td>
<tr><td>&#x1F939;&#x1F3FB;<td><td>person juggling, type-1-2<td><code>&amp;#x1F939; &amp;#x1F3FB;</code><td><code>&amp;#129337; &amp;#127995;</code><td>
<tr><td>&#x1F939;&#x1F3FC;<td><td>person juggling, type-3<td><code>&amp;#x1F939; &amp;#x1F3FC;</code><td><code>&amp;#129337; &amp;#127996;</code><td>
<tr><td>&#x1F939;&#x1F3FD;<td><td>person juggling, type-4<td><code>&amp;#x1F939; &amp;#x1F3FD;</code><td><code>&amp;#129337; &amp;#127997;</code><td>
<tr><td>&#x1F939;&#x1F3FE;<td><td>person juggling, type-5<td><code>&amp;#x1F939; &amp;#x1F3FE;</code><td><code>&amp;#129337; &amp;#127998;</code><td>
<tr><td>&#x1F939;&#x1F3FF;<td><td>person juggling, type-6<td><code>&amp;#x1F939; &amp;#x1F3FF;</code><td><code>&amp;#129337; &amp;#127999;</code><td>
<tr><td>&#x1F3AE;<td><td>VIDEO GAME<td><code>&amp;#x1F3AE;</code><td><code>&amp;#127918;</code><td>
<tr><td>&#x1F579;<td><td>JOYSTICK<td><code>&amp;#x1F579;</code><td><code>&amp;#128377;</code><td>
<tr><td>&#x1F3B2;<td><td>GAME DIE<td><code>&amp;#x1F3B2;</code><td><code>&amp;#127922;</code><td>
<tr><td>&#x2660;<td><td>BLACK SPADE SUIT &#x224A; spade suit<td><code>&amp;#x2660;</code><td><code>&amp;#9824;</code><td>
<tr><td>&#x2665;<td><td>BLACK HEART SUIT &#x224A; heart suit<td><code>&amp;#x2665;</code><td><code>&amp;#9829;</code><td>
<tr><td>&#x2666;<td><td>BLACK DIAMOND SUIT &#x224A; diamond suit<td><code>&amp;#x2666;</code><td><code>&amp;#9830;</code><td>
<tr><td>&#x2663;<td><td>BLACK CLUB SUIT &#x224A; club suit<td><code>&amp;#x2663;</code><td><code>&amp;#9827;</code><td>
<tr><td>&#x1F0CF;<td><td>PLAYING CARD BLACK JOKER &#x224A; joker<td><code>&amp;#x1F0CF;</code><td><code>&amp;#127183;</code><td>
<tr><td>&#x1F004;<td><td>MAHJONG TILE RED DRAGON &#x224A; mahjong red dragon<td><code>&amp;#x1F004;</code><td><code>&amp;#126980;</code><td>
<tr><td>&#x1F3B4;<td><td>FLOWER PLAYING CARDS<td><code>&amp;#x1F3B4;</code><td><code>&amp;#127924;</code><td>
<tr><td>&#x1F507;<td><td>SPEAKER WITH CANCELLATION STROKE &#x224A; speaker off<td><code>&amp;#x1F507;</code><td><code>&amp;#128263;</code><td>
<tr><td>&#x1F508;<td><td>SPEAKER<td><code>&amp;#x1F508;</code><td><code>&amp;#128264;</code><td>
<tr><td>&#x1F509;<td><td>SPEAKER WITH ONE SOUND WAVE &#x224A; speaker on<td><code>&amp;#x1F509;</code><td><code>&amp;#128265;</code><td>
<tr><td>&#x1F50A;<td><td>SPEAKER WITH THREE SOUND WAVES &#x224A; speaker loud<td><code>&amp;#x1F50A;</code><td><code>&amp;#128266;</code><td>
<tr><td>&#x1F4E2;<td><td>PUBLIC ADDRESS LOUDSPEAKER &#x224A; loudspeaker<td><code>&amp;#x1F4E2;</code><td><code>&amp;#128226;</code><td>
<tr><td>&#x1F4E3;<td><td>CHEERING MEGAPHONE &#x224A; megaphone<td><code>&amp;#x1F4E3;</code><td><code>&amp;#128227;</code><td>
<tr><td>&#x1F4EF;<td><td>POSTAL HORN<td><code>&amp;#x1F4EF;</code><td><code>&amp;#128239;</code><td>
<tr><td>&#x1F514;<td><td>BELL<td><code>&amp;#x1F514;</code><td><code>&amp;#128276;</code><td>
<tr><td>&#x1F515;<td><td>BELL WITH CANCELLATION STROKE &#x224A; bell with slash<td><code>&amp;#x1F515;</code><td><code>&amp;#128277;</code><td>
<tr><td>&#x1F3BC;<td><td>MUSICAL SCORE<td><code>&amp;#x1F3BC;</code><td><code>&amp;#127932;</code><td>
<tr><td>&#x1F3B5;<td><td>MUSICAL NOTE<td><code>&amp;#x1F3B5;</code><td><code>&amp;#127925;</code><td>
<tr><td>&#x1F3B6;<td><td>MULTIPLE MUSICAL NOTES &#x224A; musical notes<td><code>&amp;#x1F3B6;</code><td><code>&amp;#127926;</code><td>
<tr><td>&#x1F399;<td><td>STUDIO MICROPHONE<td><code>&amp;#x1F399;</code><td><code>&amp;#127897;</code><td>
<tr><td>&#x1F39A;<td><td>LEVEL SLIDER<td><code>&amp;#x1F39A;</code><td><code>&amp;#127898;</code><td>
<tr><td>&#x1F39B;<td><td>CONTROL KNOBS<td><code>&amp;#x1F39B;</code><td><code>&amp;#127899;</code><td>
<tr><td>&#x1F3A4;<td><td>MICROPHONE<td><code>&amp;#x1F3A4;</code><td><code>&amp;#127908;</code><td>
<tr><td>&#x1F3A7;<td><td>HEADPHONE<td><code>&amp;#x1F3A7;</code><td><code>&amp;#127911;</code><td>
<tr><td>&#x1F4FB;<td><td>RADIO<td><code>&amp;#x1F4FB;</code><td><code>&amp;#128251;</code><td>
<tr><td>&#x1F3B7;<td><td>SAXOPHONE<td><code>&amp;#x1F3B7;</code><td><code>&amp;#127927;</code><td>
<tr><td>&#x1F3B8;<td><td>GUITAR<td><code>&amp;#x1F3B8;</code><td><code>&amp;#127928;</code><td>
<tr><td>&#x1F3B9;<td><td>MUSICAL KEYBOARD<td><code>&amp;#x1F3B9;</code><td><code>&amp;#127929;</code><td>
<tr><td>&#x1F3BA;<td><td>TRUMPET<td><code>&amp;#x1F3BA;</code><td><code>&amp;#127930;</code><td>
<tr><td>&#x1F3BB;<td><td>VIOLIN<td><code>&amp;#x1F3BB;</code><td><code>&amp;#127931;</code><td>
<tr><td>&#x1F941;<td><td>DRUM WITH DRUMSTICKS &#x224A; drum<td><code>&amp;#x1F941;</code><td><code>&amp;#129345;</code><td>
<tr><td>&#x1F4F1;<td><td>MOBILE PHONE<td><code>&amp;#x1F4F1;</code><td><code>&amp;#128241;</code><td>
<tr><td>&#x1F4F2;<td><td>MOBILE PHONE WITH RIGHTWARDS ARROW AT LEFT &#x224A; mobile phone with arrow<td><code>&amp;#x1F4F2;</code><td><code>&amp;#128242;</code><td>
<tr><td>&#x260E;<td><td>BLACK TELEPHONE &#x224A; telephone<td><code>&amp;#x260E;</code><td><code>&amp;#9742;</code><td>
<tr><td>&#x1F4DE;<td><td>TELEPHONE RECEIVER<td><code>&amp;#x1F4DE;</code><td><code>&amp;#128222;</code><td>
<tr><td>&#x1F4DF;<td><td>PAGER<td><code>&amp;#x1F4DF;</code><td><code>&amp;#128223;</code><td>
<tr><td>&#x1F4E0;<td><td>FAX MACHINE<td><code>&amp;#x1F4E0;</code><td><code>&amp;#128224;</code><td>
<tr><td>&#x1F50B;<td><td>BATTERY<td><code>&amp;#x1F50B;</code><td><code>&amp;#128267;</code><td>
<tr><td>&#x1F50C;<td><td>ELECTRIC PLUG<td><code>&amp;#x1F50C;</code><td><code>&amp;#128268;</code><td>
<tr><td>&#x1F4BB;<td><td>PERSONAL COMPUTER &#x224A; laptop computer<td><code>&amp;#x1F4BB;</code><td><code>&amp;#128187;</code><td>
<tr><td>&#x1F5A5;<td><td>DESKTOP COMPUTER<td><code>&amp;#x1F5A5;</code><td><code>&amp;#128421;</code><td>
<tr><td>&#x1F5A8;<td><td>PRINTER<td><code>&amp;#x1F5A8;</code><td><code>&amp;#128424;</code><td>
<tr><td>&#x2328;<td><td>KEYBOARD<td><code>&amp;#x2328;</code><td><code>&amp;#9000;</code><td>
<tr><td>&#x1F5B1;<td><td>THREE BUTTON MOUSE &#x224A; computer mouse<td><code>&amp;#x1F5B1;</code><td><code>&amp;#128433;</code><td>
<tr><td>&#x1F5B2;<td><td>TRACKBALL<td><code>&amp;#x1F5B2;</code><td><code>&amp;#128434;</code><td>
<tr><td>&#x1F4BD;<td><td>MINIDISC<td><code>&amp;#x1F4BD;</code><td><code>&amp;#128189;</code><td>
<tr><td>&#x1F4BE;<td><td>FLOPPY DISK<td><code>&amp;#x1F4BE;</code><td><code>&amp;#128190;</code><td>
<tr><td>&#x1F4BF;<td><td>OPTICAL DISC<td><code>&amp;#x1F4BF;</code><td><code>&amp;#128191;</code><td>
<tr><td>&#x1F4C0;<td><td>DVD<td><code>&amp;#x1F4C0;</code><td><code>&amp;#128192;</code><td>
<tr><td>&#x1F3A5;<td><td>MOVIE CAMERA<td><code>&amp;#x1F3A5;</code><td><code>&amp;#127909;</code><td>
<tr><td>&#x1F39E;<td><td>FILM FRAMES<td><code>&amp;#x1F39E;</code><td><code>&amp;#127902;</code><td>
<tr><td>&#x1F4FD;<td><td>FILM PROJECTOR<td><code>&amp;#x1F4FD;</code><td><code>&amp;#128253;</code><td>
<tr><td>&#x1F3AC;<td><td>CLAPPER BOARD<td><code>&amp;#x1F3AC;</code><td><code>&amp;#127916;</code><td>
<tr><td>&#x1F4FA;<td><td>TELEVISION<td><code>&amp;#x1F4FA;</code><td><code>&amp;#128250;</code><td>
<tr><td>&#x1F4F7;<td><td>CAMERA<td><code>&amp;#x1F4F7;</code><td><code>&amp;#128247;</code><td>
<tr><td>&#x1F4F8;<td><td>CAMERA WITH FLASH<td><code>&amp;#x1F4F8;</code><td><code>&amp;#128248;</code><td>
<tr><td>&#x1F4F9;<td><td>VIDEO CAMERA<td><code>&amp;#x1F4F9;</code><td><code>&amp;#128249;</code><td>
<tr><td>&#x1F4FC;<td><td>VIDEOCASSETTE<td><code>&amp;#x1F4FC;</code><td><code>&amp;#128252;</code><td>
<tr><td>&#x1F50D;<td><td>LEFT-POINTING MAGNIFYING GLASS<td><code>&amp;#x1F50D;</code><td><code>&amp;#128269;</code><td>
<tr><td>&#x1F50E;<td><td>RIGHT-POINTING MAGNIFYING GLASS<td><code>&amp;#x1F50E;</code><td><code>&amp;#128270;</code><td>
<tr><td>&#x1F52C;<td><td>MICROSCOPE<td><code>&amp;#x1F52C;</code><td><code>&amp;#128300;</code><td>
<tr><td>&#x1F52D;<td><td>TELESCOPE<td><code>&amp;#x1F52D;</code><td><code>&amp;#128301;</code><td>
<tr><td>&#x1F4E1;<td><td>SATELLITE ANTENNA<td><code>&amp;#x1F4E1;</code><td><code>&amp;#128225;</code><td>
<tr><td>&#x1F56F;<td><td>CANDLE<td><code>&amp;#x1F56F;</code><td><code>&amp;#128367;</code><td>
<tr><td>&#x1F4A1;<td><td>ELECTRIC LIGHT BULB &#x224A; light bulb<td><code>&amp;#x1F4A1;</code><td><code>&amp;#128161;</code><td>
<tr><td>&#x1F526;<td><td>ELECTRIC TORCH &#x224A; flashlight<td><code>&amp;#x1F526;</code><td><code>&amp;#128294;</code><td>
<tr><td>&#x1F3EE;<td><td>IZAKAYA LANTERN &#x224A; red paper lantern<td><code>&amp;#x1F3EE;</code><td><code>&amp;#127982;</code><td>
<tr><td>&#x1F4D4;<td><td>NOTEBOOK WITH DECORATIVE COVER<td><code>&amp;#x1F4D4;</code><td><code>&amp;#128212;</code><td>
<tr><td>&#x1F4D5;<td><td>CLOSED BOOK<td><code>&amp;#x1F4D5;</code><td><code>&amp;#128213;</code><td>
<tr><td>&#x1F4D6;<td><td>OPEN BOOK<td><code>&amp;#x1F4D6;</code><td><code>&amp;#128214;</code><td>
<tr><td>&#x1F4D7;<td><td>GREEN BOOK<td><code>&amp;#x1F4D7;</code><td><code>&amp;#128215;</code><td>
<tr><td>&#x1F4D8;<td><td>BLUE BOOK<td><code>&amp;#x1F4D8;</code><td><code>&amp;#128216;</code><td>
<tr><td>&#x1F4D9;<td><td>ORANGE BOOK<td><code>&amp;#x1F4D9;</code><td><code>&amp;#128217;</code><td>
<tr><td>&#x1F4DA;<td><td>BOOKS<td><code>&amp;#x1F4DA;</code><td><code>&amp;#128218;</code><td>
<tr><td>&#x1F4D3;<td><td>NOTEBOOK<td><code>&amp;#x1F4D3;</code><td><code>&amp;#128211;</code><td>
<tr><td>&#x1F4D2;<td><td>LEDGER<td><code>&amp;#x1F4D2;</code><td><code>&amp;#128210;</code><td>
<tr><td>&#x1F4C3;<td><td>PAGE WITH CURL<td><code>&amp;#x1F4C3;</code><td><code>&amp;#128195;</code><td>
<tr><td>&#x1F4DC;<td><td>SCROLL<td><code>&amp;#x1F4DC;</code><td><code>&amp;#128220;</code><td>
<tr><td>&#x1F4C4;<td><td>PAGE FACING UP<td><code>&amp;#x1F4C4;</code><td><code>&amp;#128196;</code><td>
<tr><td>&#x1F4F0;<td><td>NEWSPAPER<td><code>&amp;#x1F4F0;</code><td><code>&amp;#128240;</code><td>
<tr><td>&#x1F5DE;<td><td>ROLLED-UP NEWSPAPER<td><code>&amp;#x1F5DE;</code><td><code>&amp;#128478;</code><td>
<tr><td>&#x1F4D1;<td><td>BOOKMARK TABS<td><code>&amp;#x1F4D1;</code><td><code>&amp;#128209;</code><td>
<tr><td>&#x1F516;<td><td>BOOKMARK<td><code>&amp;#x1F516;</code><td><code>&amp;#128278;</code><td>
<tr><td>&#x1F3F7;<td><td>LABEL<td><code>&amp;#x1F3F7;</code><td><code>&amp;#127991;</code><td>
<tr><td>&#x1F4B0;<td><td>MONEY BAG<td><code>&amp;#x1F4B0;</code><td><code>&amp;#128176;</code><td>
<tr><td>&#x1F4B4;<td><td>BANKNOTE WITH YEN SIGN &#x224A; yen banknote<td><code>&amp;#x1F4B4;</code><td><code>&amp;#128180;</code><td>
<tr><td>&#x1F4B5;<td><td>BANKNOTE WITH DOLLAR SIGN &#x224A; dollar banknote<td><code>&amp;#x1F4B5;</code><td><code>&amp;#128181;</code><td>
<tr><td>&#x1F4B6;<td><td>BANKNOTE WITH EURO SIGN &#x224A; euro banknote<td><code>&amp;#x1F4B6;</code><td><code>&amp;#128182;</code><td>
<tr><td>&#x1F4B7;<td><td>BANKNOTE WITH POUND SIGN &#x224A; pound banknote<td><code>&amp;#x1F4B7;</code><td><code>&amp;#128183;</code><td>
<tr><td>&#x1F4B8;<td><td>MONEY WITH WINGS<td><code>&amp;#x1F4B8;</code><td><code>&amp;#128184;</code><td>
<tr><td>&#x1F4B3;<td><td>CREDIT CARD<td><code>&amp;#x1F4B3;</code><td><code>&amp;#128179;</code><td>
<tr><td>&#x1F4B9;<td><td>CHART WITH UPWARDS TREND AND YEN SIGN &#x224A; chart increasing with yen<td><code>&amp;#x1F4B9;</code><td><code>&amp;#128185;</code><td>
<tr><td>&#x1F4B1;<td><td>CURRENCY EXCHANGE<td><code>&amp;#x1F4B1;</code><td><code>&amp;#128177;</code><td>
<tr><td>&#x1F4B2;<td><td>HEAVY DOLLAR SIGN<td><code>&amp;#x1F4B2;</code><td><code>&amp;#128178;</code><td>
<tr><td>&#x2709;<td><td>ENVELOPE<td><code>&amp;#x2709;</code><td><code>&amp;#9993;</code><td>
<tr><td>&#x1F4E7;<td><td>E-MAIL SYMBOL &#x224A; e-mail<td><code>&amp;#x1F4E7;</code><td><code>&amp;#128231;</code><td>
<tr><td>&#x1F4E8;<td><td>INCOMING ENVELOPE<td><code>&amp;#x1F4E8;</code><td><code>&amp;#128232;</code><td>
<tr><td>&#x1F4E9;<td><td>ENVELOPE WITH DOWNWARDS ARROW ABOVE &#x224A; envelope with arrow<td><code>&amp;#x1F4E9;</code><td><code>&amp;#128233;</code><td>
<tr><td>&#x1F4E4;<td><td>OUTBOX TRAY<td><code>&amp;#x1F4E4;</code><td><code>&amp;#128228;</code><td>
<tr><td>&#x1F4E5;<td><td>INBOX TRAY<td><code>&amp;#x1F4E5;</code><td><code>&amp;#128229;</code><td>
<tr><td>&#x1F4E6;<td><td>PACKAGE<td><code>&amp;#x1F4E6;</code><td><code>&amp;#128230;</code><td>
<tr><td>&#x1F4EB;<td><td>CLOSED MAILBOX WITH RAISED FLAG<td><code>&amp;#x1F4EB;</code><td><code>&amp;#128235;</code><td>
<tr><td>&#x1F4EA;<td><td>CLOSED MAILBOX WITH LOWERED FLAG<td><code>&amp;#x1F4EA;</code><td><code>&amp;#128234;</code><td>
<tr><td>&#x1F4EC;<td><td>OPEN MAILBOX WITH RAISED FLAG<td><code>&amp;#x1F4EC;</code><td><code>&amp;#128236;</code><td>
<tr><td>&#x1F4ED;<td><td>OPEN MAILBOX WITH LOWERED FLAG<td><code>&amp;#x1F4ED;</code><td><code>&amp;#128237;</code><td>
<tr><td>&#x1F4EE;<td><td>POSTBOX<td><code>&amp;#x1F4EE;</code><td><code>&amp;#128238;</code><td>
<tr><td>&#x1F5F3;<td><td>BALLOT BOX WITH BALLOT<td><code>&amp;#x1F5F3;</code><td><code>&amp;#128499;</code><td>
<tr><td>&#x270F;<td><td>PENCIL<td><code>&amp;#x270F;</code><td><code>&amp;#9999;</code><td>
<tr><td>&#x2712;<td><td>BLACK NIB<td><code>&amp;#x2712;</code><td><code>&amp;#10002;</code><td>
<tr><td>&#x1F58B;<td><td>LOWER LEFT FOUNTAIN PEN &#x224A; fountain pen<td><code>&amp;#x1F58B;</code><td><code>&amp;#128395;</code><td>
<tr><td>&#x1F58A;<td><td>LOWER LEFT BALLPOINT PEN &#x224A; pen<td><code>&amp;#x1F58A;</code><td><code>&amp;#128394;</code><td>
<tr><td>&#x1F58C;<td><td>LOWER LEFT PAINTBRUSH &#x224A; paintbrush<td><code>&amp;#x1F58C;</code><td><code>&amp;#128396;</code><td>
<tr><td>&#x1F58D;<td><td>LOWER LEFT CRAYON &#x224A; crayon<td><code>&amp;#x1F58D;</code><td><code>&amp;#128397;</code><td>
<tr><td>&#x1F4DD;<td><td>MEMO<td><code>&amp;#x1F4DD;</code><td><code>&amp;#128221;</code><td>
<tr><td>&#x1F4BC;<td><td>BRIEFCASE<td><code>&amp;#x1F4BC;</code><td><code>&amp;#128188;</code><td>
<tr><td>&#x1F4C1;<td><td>FILE FOLDER<td><code>&amp;#x1F4C1;</code><td><code>&amp;#128193;</code><td>
<tr><td>&#x1F4C2;<td><td>OPEN FILE FOLDER<td><code>&amp;#x1F4C2;</code><td><code>&amp;#128194;</code><td>
<tr><td>&#x1F5C2;<td><td>CARD INDEX DIVIDERS<td><code>&amp;#x1F5C2;</code><td><code>&amp;#128450;</code><td>
<tr><td>&#x1F4C5;<td><td>CALENDAR<td><code>&amp;#x1F4C5;</code><td><code>&amp;#128197;</code><td>
<tr><td>&#x1F4C6;<td><td>TEAR-OFF CALENDAR<td><code>&amp;#x1F4C6;</code><td><code>&amp;#128198;</code><td>
<tr><td>&#x1F5D2;<td><td>SPIRAL NOTE PAD &#x224A; spiral notepad<td><code>&amp;#x1F5D2;</code><td><code>&amp;#128466;</code><td>
<tr><td>&#x1F5D3;<td><td>SPIRAL CALENDAR PAD &#x224A; spiral calendar<td><code>&amp;#x1F5D3;</code><td><code>&amp;#128467;</code><td>
<tr><td>&#x1F4C7;<td><td>CARD INDEX<td><code>&amp;#x1F4C7;</code><td><code>&amp;#128199;</code><td>
<tr><td>&#x1F4C8;<td><td>CHART WITH UPWARDS TREND &#x224A; chart increasing<td><code>&amp;#x1F4C8;</code><td><code>&amp;#128200;</code><td>
<tr><td>&#x1F4C9;<td><td>CHART WITH DOWNWARDS TREND &#x224A; chart decreasing<td><code>&amp;#x1F4C9;</code><td><code>&amp;#128201;</code><td>
<tr><td>&#x1F4CA;<td><td>BAR CHART<td><code>&amp;#x1F4CA;</code><td><code>&amp;#128202;</code><td>
<tr><td>&#x1F4CB;<td><td>CLIPBOARD<td><code>&amp;#x1F4CB;</code><td><code>&amp;#128203;</code><td>
<tr><td>&#x1F4CC;<td><td>PUSHPIN<td><code>&amp;#x1F4CC;</code><td><code>&amp;#128204;</code><td>
<tr><td>&#x1F4CD;<td><td>ROUND PUSHPIN<td><code>&amp;#x1F4CD;</code><td><code>&amp;#128205;</code><td>
<tr><td>&#x1F4CE;<td><td>PAPERCLIP<td><code>&amp;#x1F4CE;</code><td><code>&amp;#128206;</code><td>
<tr><td>&#x1F587;<td><td>LINKED PAPERCLIPS<td><code>&amp;#x1F587;</code><td><code>&amp;#128391;</code><td>
<tr><td>&#x1F4CF;<td><td>STRAIGHT RULER<td><code>&amp;#x1F4CF;</code><td><code>&amp;#128207;</code><td>
<tr><td>&#x1F4D0;<td><td>TRIANGULAR RULER<td><code>&amp;#x1F4D0;</code><td><code>&amp;#128208;</code><td>
<tr><td>&#x2702;<td><td>BLACK SCISSORS &#x224A; scissors<td><code>&amp;#x2702;</code><td><code>&amp;#9986;</code><td>
<tr><td>&#x1F5C3;<td><td>CARD FILE BOX<td><code>&amp;#x1F5C3;</code><td><code>&amp;#128451;</code><td>
<tr><td>&#x1F5C4;<td><td>FILE CABINET<td><code>&amp;#x1F5C4;</code><td><code>&amp;#128452;</code><td>
<tr><td>&#x1F5D1;<td><td>WASTEBASKET<td><code>&amp;#x1F5D1;</code><td><code>&amp;#128465;</code><td>
<tr><td>&#x1F512;<td><td>LOCK<td><code>&amp;#x1F512;</code><td><code>&amp;#128274;</code><td>
<tr><td>&#x1F513;<td><td>OPEN LOCK<td><code>&amp;#x1F513;</code><td><code>&amp;#128275;</code><td>
<tr><td>&#x1F50F;<td><td>LOCK WITH INK PEN &#x224A; lock with pen<td><code>&amp;#x1F50F;</code><td><code>&amp;#128271;</code><td>
<tr><td>&#x1F510;<td><td>CLOSED LOCK WITH KEY<td><code>&amp;#x1F510;</code><td><code>&amp;#128272;</code><td>
<tr><td>&#x1F511;<td><td>KEY<td><code>&amp;#x1F511;</code><td><code>&amp;#128273;</code><td>
<tr><td>&#x1F5DD;<td><td>OLD KEY<td><code>&amp;#x1F5DD;</code><td><code>&amp;#128477;</code><td>
<tr><td>&#x1F528;<td><td>HAMMER<td><code>&amp;#x1F528;</code><td><code>&amp;#128296;</code><td>
<tr><td>&#x26CF;<td><td>PICK<td><code>&amp;#x26CF;</code><td><code>&amp;#9935;</code><td>
<tr><td>&#x2692;<td><td>HAMMER AND PICK<td><code>&amp;#x2692;</code><td><code>&amp;#9874;</code><td>
<tr><td>&#x1F6E0;<td><td>HAMMER AND WRENCH<td><code>&amp;#x1F6E0;</code><td><code>&amp;#128736;</code><td>
<tr><td>&#x1F5E1;<td><td>DAGGER KNIFE &#x224A; dagger<td><code>&amp;#x1F5E1;</code><td><code>&amp;#128481;</code><td>
<tr><td>&#x2694;<td><td>CROSSED SWORDS<td><code>&amp;#x2694;</code><td><code>&amp;#9876;</code><td>
<tr><td>&#x1F52B;<td><td>PISTOL<td><code>&amp;#x1F52B;</code><td><code>&amp;#128299;</code><td>
<tr><td>&#x1F3F9;<td><td>BOW AND ARROW<td><code>&amp;#x1F3F9;</code><td><code>&amp;#127993;</code><td>
<tr><td>&#x1F6E1;<td><td>SHIELD<td><code>&amp;#x1F6E1;</code><td><code>&amp;#128737;</code><td>
<tr><td>&#x1F527;<td><td>WRENCH<td><code>&amp;#x1F527;</code><td><code>&amp;#128295;</code><td>
<tr><td>&#x1F529;<td><td>NUT AND BOLT<td><code>&amp;#x1F529;</code><td><code>&amp;#128297;</code><td>
<tr><td>&#x2699;<td><td>GEAR<td><code>&amp;#x2699;</code><td><code>&amp;#9881;</code><td>
<tr><td>&#x1F5DC;<td><td>COMPRESSION<td><code>&amp;#x1F5DC;</code><td><code>&amp;#128476;</code><td>
<tr><td>&#x2697;<td><td>ALEMBIC<td><code>&amp;#x2697;</code><td><code>&amp;#9879;</code><td>
<tr><td>&#x2696;<td><td>SCALES &#x224A; balance scale<td><code>&amp;#x2696;</code><td><code>&amp;#9878;</code><td>
<tr><td>&#x1F517;<td><td>LINK SYMBOL &#x224A; link<td><code>&amp;#x1F517;</code><td><code>&amp;#128279;</code><td>
<tr><td>&#x26D3;<td><td>CHAINS<td><code>&amp;#x26D3;</code><td><code>&amp;#9939;</code><td>
<tr><td>&#x1F489;<td><td>SYRINGE<td><code>&amp;#x1F489;</code><td><code>&amp;#128137;</code><td>
<tr><td>&#x1F48A;<td><td>PILL<td><code>&amp;#x1F48A;</code><td><code>&amp;#128138;</code><td>
<tr><td>&#x1F6AC;<td><td>SMOKING SYMBOL &#x224A; smoking<td><code>&amp;#x1F6AC;</code><td><code>&amp;#128684;</code><td>
<tr><td>&#x26B0;<td><td>COFFIN<td><code>&amp;#x26B0;</code><td><code>&amp;#9904;</code><td>
<tr><td>&#x26B1;<td><td>FUNERAL URN<td><code>&amp;#x26B1;</code><td><code>&amp;#9905;</code><td>
<tr><td>&#x1F5FF;<td><td>MOYAI &#x224A; moai<td><code>&amp;#x1F5FF;</code><td><code>&amp;#128511;</code><td>
<tr><td>&#x1F6E2;<td><td>OIL DRUM<td><code>&amp;#x1F6E2;</code><td><code>&amp;#128738;</code><td>
<tr><td>&#x1F52E;<td><td>CRYSTAL BALL<td><code>&amp;#x1F52E;</code><td><code>&amp;#128302;</code><td>
<tr><td>&#x1F6D2;<td><td>SHOPPING TROLLEY &#x224A; shopping cart<td><code>&amp;#x1F6D2;</code><td><code>&amp;#128722;</code><td>
<tr><td>&#x1F3E7;<td><td>AUTOMATED TELLER MACHINE &#x224A; ATM sign<td><code>&amp;#x1F3E7;</code><td><code>&amp;#127975;</code><td>
<tr><td>&#x1F6AE;<td><td>PUT LITTER IN ITS PLACE SYMBOL &#x224A; litter in bin sign<td><code>&amp;#x1F6AE;</code><td><code>&amp;#128686;</code><td>
<tr><td>&#x1F6B0;<td><td>POTABLE WATER SYMBOL &#x224A; potable water<td><code>&amp;#x1F6B0;</code><td><code>&amp;#128688;</code><td>
<tr><td>&#x267F;<td><td>WHEELCHAIR SYMBOL &#x224A; wheelchair<td><code>&amp;#x267F;</code><td><code>&amp;#9855;</code><td>
<tr><td>&#x1F6B9;<td><td>MENS SYMBOL &#x224A; menâ€™s room<td><code>&amp;#x1F6B9;</code><td><code>&amp;#128697;</code><td>
<tr><td>&#x1F6BA;<td><td>WOMENS SYMBOL &#x224A; womenâ€™s room<td><code>&amp;#x1F6BA;</code><td><code>&amp;#128698;</code><td>
<tr><td>&#x1F6BB;<td><td>RESTROOM<td><code>&amp;#x1F6BB;</code><td><code>&amp;#128699;</code><td>
<tr><td>&#x1F6BC;<td><td>BABY SYMBOL<td><code>&amp;#x1F6BC;</code><td><code>&amp;#128700;</code><td>
<tr><td>&#x1F6BE;<td><td>WATER CLOSET<td><code>&amp;#x1F6BE;</code><td><code>&amp;#128702;</code><td>
<tr><td>&#x1F6C2;<td><td>PASSPORT CONTROL<td><code>&amp;#x1F6C2;</code><td><code>&amp;#128706;</code><td>
<tr><td>&#x1F6C3;<td><td>CUSTOMS<td><code>&amp;#x1F6C3;</code><td><code>&amp;#128707;</code><td>
<tr><td>&#x1F6C4;<td><td>BAGGAGE CLAIM<td><code>&amp;#x1F6C4;</code><td><code>&amp;#128708;</code><td>
<tr><td>&#x1F6C5;<td><td>LEFT LUGGAGE<td><code>&amp;#x1F6C5;</code><td><code>&amp;#128709;</code><td>
<tr><td>&#x26A0;<td><td>WARNING SIGN &#x224A; warning<td><code>&amp;#x26A0;</code><td><code>&amp;#9888;</code><td>
<tr><td>&#x1F6B8;<td><td>CHILDREN CROSSING<td><code>&amp;#x1F6B8;</code><td><code>&amp;#128696;</code><td>
<tr><td>&#x26D4;<td><td>NO ENTRY<td><code>&amp;#x26D4;</code><td><code>&amp;#9940;</code><td>
<tr><td>&#x1F6AB;<td><td>NO ENTRY SIGN &#x224A; prohibited<td><code>&amp;#x1F6AB;</code><td><code>&amp;#128683;</code><td>
<tr><td>&#x1F6B3;<td><td>NO BICYCLES<td><code>&amp;#x1F6B3;</code><td><code>&amp;#128691;</code><td>
<tr><td>&#x1F6AD;<td><td>NO SMOKING SYMBOL &#x224A; no smoking<td><code>&amp;#x1F6AD;</code><td><code>&amp;#128685;</code><td>
<tr><td>&#x1F6AF;<td><td>DO NOT LITTER SYMBOL &#x224A; no littering<td><code>&amp;#x1F6AF;</code><td><code>&amp;#128687;</code><td>
<tr><td>&#x1F6B1;<td><td>NON-POTABLE WATER SYMBOL &#x224A; non-potable water<td><code>&amp;#x1F6B1;</code><td><code>&amp;#128689;</code><td>
<tr><td>&#x1F6B7;<td><td>NO PEDESTRIANS<td><code>&amp;#x1F6B7;</code><td><code>&amp;#128695;</code><td>
<tr><td>&#x1F4F5;<td><td>NO MOBILE PHONES<td><code>&amp;#x1F4F5;</code><td><code>&amp;#128245;</code><td>
<tr><td>&#x1F51E;<td><td>NO ONE UNDER EIGHTEEN SYMBOL &#x224A; no one under eighteen<td><code>&amp;#x1F51E;</code><td><code>&amp;#128286;</code><td>
<tr><td>&#x2622;<td><td>RADIOACTIVE SIGN &#x224A; radioactive<td><code>&amp;#x2622;</code><td><code>&amp;#9762;</code><td>
<tr><td>&#x2623;<td><td>BIOHAZARD SIGN &#x224A; biohazard<td><code>&amp;#x2623;</code><td><code>&amp;#9763;</code><td>
<tr><td>&#x2B06;<td><td>UPWARDS BLACK ARROW &#x224A; up arrow<td><code>&amp;#x2B06;</code><td><code>&amp;#11014;</code><td>
<tr><td>&#x2197;<td><td>NORTH EAST ARROW &#x224A; up-right arrow<td><code>&amp;#x2197;</code><td><code>&amp;#8599;</code><td>
<tr><td>&#x27A1;<td><td>BLACK RIGHTWARDS ARROW &#x224A; right arrow<td><code>&amp;#x27A1;</code><td><code>&amp;#10145;</code><td>
<tr><td>&#x2198;<td><td>SOUTH EAST ARROW &#x224A; down-right arrow<td><code>&amp;#x2198;</code><td><code>&amp;#8600;</code><td>
<tr><td>&#x2B07;<td><td>DOWNWARDS BLACK ARROW &#x224A; down arrow<td><code>&amp;#x2B07;</code><td><code>&amp;#11015;</code><td>
<tr><td>&#x2199;<td><td>SOUTH WEST ARROW &#x224A; down-left arrow<td><code>&amp;#x2199;</code><td><code>&amp;#8601;</code><td>
<tr><td>&#x2B05;<td><td>LEFTWARDS BLACK ARROW &#x224A; left arrow<td><code>&amp;#x2B05;</code><td><code>&amp;#11013;</code><td>
<tr><td>&#x2196;<td><td>NORTH WEST ARROW &#x224A; up-left arrow<td><code>&amp;#x2196;</code><td><code>&amp;#8598;</code><td>
<tr><td>&#x2195;<td><td>UP DOWN ARROW &#x224A; up-down arrow<td><code>&amp;#x2195;</code><td><code>&amp;#8597;</code><td>
<tr><td>&#x2194;<td><td>LEFT RIGHT ARROW &#x224A; left-right arrow<td><code>&amp;#x2194;</code><td><code>&amp;#8596;</code><td>
<tr><td>&#x21A9;<td><td>LEFTWARDS ARROW WITH HOOK &#x224A; right arrow curving left<td><code>&amp;#x21A9;</code><td><code>&amp;#8617;</code><td>
<tr><td>&#x21AA;<td><td>RIGHTWARDS ARROW WITH HOOK &#x224A; left arrow curving right<td><code>&amp;#x21AA;</code><td><code>&amp;#8618;</code><td>
<tr><td>&#x2934;<td><td>ARROW POINTING RIGHTWARDS THEN CURVING UPWARDS &#x224A; right arrow curving up<td><code>&amp;#x2934;</code><td><code>&amp;#10548;</code><td>
<tr><td>&#x2935;<td><td>ARROW POINTING RIGHTWARDS THEN CURVING DOWNWARDS &#x224A; right arrow curving down<td><code>&amp;#x2935;</code><td><code>&amp;#10549;</code><td>
<tr><td>&#x1F503;<td><td>CLOCKWISE DOWNWARDS AND UPWARDS OPEN CIRCLE ARROWS &#x224A; clockwise vertical arrows<td><code>&amp;#x1F503;</code><td><code>&amp;#128259;</code><td>
<tr><td>&#x1F504;<td><td>ANTICLOCKWISE DOWNWARDS AND UPWARDS OPEN CIRCLE ARROWS &#x224A; anticlockwise arrows button<td><code>&amp;#x1F504;</code><td><code>&amp;#128260;</code><td>
<tr><td>&#x1F519;<td><td>BACK WITH LEFTWARDS ARROW ABOVE &#x224A; back arrow<td><code>&amp;#x1F519;</code><td><code>&amp;#128281;</code><td>
<tr><td>&#x1F51A;<td><td>END WITH LEFTWARDS ARROW ABOVE &#x224A; end arrow<td><code>&amp;#x1F51A;</code><td><code>&amp;#128282;</code><td>
<tr><td>&#x1F51B;<td><td>ON WITH EXCLAMATION MARK WITH LEFT RIGHT ARROW ABOVE &#x224A; on! arrow<td><code>&amp;#x1F51B;</code><td><code>&amp;#128283;</code><td>
<tr><td>&#x1F51C;<td><td>SOON WITH RIGHTWARDS ARROW ABOVE &#x224A; soon arrow<td><code>&amp;#x1F51C;</code><td><code>&amp;#128284;</code><td>
<tr><td>&#x1F51D;<td><td>TOP WITH UPWARDS ARROW ABOVE &#x224A; top arrow<td><code>&amp;#x1F51D;</code><td><code>&amp;#128285;</code><td>
<tr><td>&#x1F6D0;<td><td>PLACE OF WORSHIP<td><code>&amp;#x1F6D0;</code><td><code>&amp;#128720;</code><td>
<tr><td>&#x269B;<td><td>ATOM SYMBOL<td><code>&amp;#x269B;</code><td><code>&amp;#9883;</code><td>
<tr><td>&#x1F549;<td><td>OM SYMBOL &#x224A; om<td><code>&amp;#x1F549;</code><td><code>&amp;#128329;</code><td>
<tr><td>&#x2721;<td><td>STAR OF DAVID<td><code>&amp;#x2721;</code><td><code>&amp;#10017;</code><td>
<tr><td>&#x2638;<td><td>WHEEL OF DHARMA<td><code>&amp;#x2638;</code><td><code>&amp;#9784;</code><td>
<tr><td>&#x262F;<td><td>YIN YANG<td><code>&amp;#x262F;</code><td><code>&amp;#9775;</code><td>
<tr><td>&#x271D;<td><td>LATIN CROSS<td><code>&amp;#x271D;</code><td><code>&amp;#10013;</code><td>
<tr><td>&#x2626;<td><td>ORTHODOX CROSS<td><code>&amp;#x2626;</code><td><code>&amp;#9766;</code><td>
<tr><td>&#x262A;<td><td>STAR AND CRESCENT<td><code>&amp;#x262A;</code><td><code>&amp;#9770;</code><td>
<tr><td>&#x262E;<td><td>PEACE SYMBOL<td><code>&amp;#x262E;</code><td><code>&amp;#9774;</code><td>
<tr><td>&#x1F54E;<td><td>MENORAH WITH NINE BRANCHES &#x224A; menorah<td><code>&amp;#x1F54E;</code><td><code>&amp;#128334;</code><td>
<tr><td>&#x1F52F;<td><td>SIX POINTED STAR WITH MIDDLE DOT &#x224A; dotted six-pointed star<td><code>&amp;#x1F52F;</code><td><code>&amp;#128303;</code><td>
<tr><td>&#x267B;<td><td>BLACK UNIVERSAL RECYCLING SYMBOL &#x224A; recycling symbol<td><code>&amp;#x267B;</code><td><code>&amp;#9851;</code><td>
<tr><td>&#x1F4DB;<td><td>NAME BADGE<td><code>&amp;#x1F4DB;</code><td><code>&amp;#128219;</code><td>
<tr><td>&#x269C;<td><td>FLEUR-DE-LIS<td><code>&amp;#x269C;</code><td><code>&amp;#9884;</code><td>
<tr><td>&#x1F530;<td><td>JAPANESE SYMBOL FOR BEGINNER<td><code>&amp;#x1F530;</code><td><code>&amp;#128304;</code><td>
<tr><td>&#x1F531;<td><td>TRIDENT EMBLEM<td><code>&amp;#x1F531;</code><td><code>&amp;#128305;</code><td>
<tr><td>&#x2B55;<td><td>HEAVY LARGE CIRCLE<td><code>&amp;#x2B55;</code><td><code>&amp;#11093;</code><td>
<tr><td>&#x2705;<td><td>WHITE HEAVY CHECK MARK<td><code>&amp;#x2705;</code><td><code>&amp;#9989;</code><td>
<tr><td>&#x2611;<td><td>BALLOT BOX WITH CHECK<td><code>&amp;#x2611;</code><td><code>&amp;#9745;</code><td>
<tr><td>&#x2714;<td><td>HEAVY CHECK MARK<td><code>&amp;#x2714;</code><td><code>&amp;#10004;</code><td>
<tr><td>&#x2716;<td><td>HEAVY MULTIPLICATION X<td><code>&amp;#x2716;</code><td><code>&amp;#10006;</code><td>
<tr><td>&#x274C;<td><td>CROSS MARK<td><code>&amp;#x274C;</code><td><code>&amp;#10060;</code><td>
<tr><td>&#x274E;<td><td>NEGATIVE SQUARED CROSS MARK &#x224A; cross mark button<td><code>&amp;#x274E;</code><td><code>&amp;#10062;</code><td>
<tr><td>&#x2795;<td><td>HEAVY PLUS SIGN<td><code>&amp;#x2795;</code><td><code>&amp;#10133;</code><td>
<tr><td>&#x2796;<td><td>HEAVY MINUS SIGN<td><code>&amp;#x2796;</code><td><code>&amp;#10134;</code><td>
<tr><td>&#x2797;<td><td>HEAVY DIVISION SIGN<td><code>&amp;#x2797;</code><td><code>&amp;#10135;</code><td>
<tr><td>&#x27B0;<td><td>CURLY LOOP<td><code>&amp;#x27B0;</code><td><code>&amp;#10160;</code><td>
<tr><td>&#x27BF;<td><td>DOUBLE CURLY LOOP<td><code>&amp;#x27BF;</code><td><code>&amp;#10175;</code><td>
<tr><td>&#x303D;<td><td>PART ALTERNATION MARK<td><code>&amp;#x303D;</code><td><code>&amp;#12349;</code><td>
<tr><td>&#x2733;<td><td>EIGHT SPOKED ASTERISK &#x224A; eight-spoked asterisk<td><code>&amp;#x2733;</code><td><code>&amp;#10035;</code><td>
<tr><td>&#x2734;<td><td>EIGHT POINTED BLACK STAR &#x224A; eight-pointed star<td><code>&amp;#x2734;</code><td><code>&amp;#10036;</code><td>
<tr><td>&#x2747;<td><td>SPARKLE<td><code>&amp;#x2747;</code><td><code>&amp;#10055;</code><td>
<tr><td>&#x203C;<td><td>DOUBLE EXCLAMATION MARK<td><code>&amp;#x203C;</code><td><code>&amp;#8252;</code><td>
<tr><td>&#x2049;<td><td>EXCLAMATION QUESTION MARK<td><code>&amp;#x2049;</code><td><code>&amp;#8265;</code><td>
<tr><td>&#x2753;<td><td>BLACK QUESTION MARK ORNAMENT &#x224A; question mark<td><code>&amp;#x2753;</code><td><code>&amp;#10067;</code><td>
<tr><td>&#x2754;<td><td>WHITE QUESTION MARK ORNAMENT &#x224A; white question mark<td><code>&amp;#x2754;</code><td><code>&amp;#10068;</code><td>
<tr><td>&#x2755;<td><td>WHITE EXCLAMATION MARK ORNAMENT &#x224A; white exclamation mark<td><code>&amp;#x2755;</code><td><code>&amp;#10069;</code><td>
<tr><td>&#x2757;<td><td>HEAVY EXCLAMATION MARK SYMBOL &#x224A; exclamation mark<td><code>&amp;#x2757;</code><td><code>&amp;#10071;</code><td>
<tr><td>&#x3030;<td><td>WAVY DASH<td><code>&amp;#x3030;</code><td><code>&amp;#12336;</code><td>
<tr><td>&#x00A9;<td><td>COPYRIGHT SIGN &#x224A; copyright<td><code>&amp;#x00A9;</code><td><code>&amp;#169;</code><td>
<tr><td>&#x00AE;<td><td>REGISTERED SIGN &#x224A; registered<td><code>&amp;#x00AE;</code><td><code>&amp;#174;</code><td>
<tr><td>&#x2122;<td><td>TRADE MARK SIGN &#x224A; trade mark<td><code>&amp;#x2122;</code><td><code>&amp;#8482;</code><td>
<tr><td>&#x2648;<td><td>ARIES<td><code>&amp;#x2648;</code><td><code>&amp;#9800;</code><td>
<tr><td>&#x2649;<td><td>TAURUS<td><code>&amp;#x2649;</code><td><code>&amp;#9801;</code><td>
<tr><td>&#x264A;<td><td>GEMINI<td><code>&amp;#x264A;</code><td><code>&amp;#9802;</code><td>
<tr><td>&#x264B;<td><td>CANCER<td><code>&amp;#x264B;</code><td><code>&amp;#9803;</code><td>
<tr><td>&#x264C;<td><td>LEO<td><code>&amp;#x264C;</code><td><code>&amp;#9804;</code><td>
<tr><td>&#x264D;<td><td>VIRGO<td><code>&amp;#x264D;</code><td><code>&amp;#9805;</code><td>
<tr><td>&#x264E;<td><td>LIBRA<td><code>&amp;#x264E;</code><td><code>&amp;#9806;</code><td>
<tr><td>&#x264F;<td><td>SCORPIUS<td><code>&amp;#x264F;</code><td><code>&amp;#9807;</code><td>
<tr><td>&#x2650;<td><td>SAGITTARIUS<td><code>&amp;#x2650;</code><td><code>&amp;#9808;</code><td>
<tr><td>&#x2651;<td><td>CAPRICORN<td><code>&amp;#x2651;</code><td><code>&amp;#9809;</code><td>
<tr><td>&#x2652;<td><td>AQUARIUS<td><code>&amp;#x2652;</code><td><code>&amp;#9810;</code><td>
<tr><td>&#x2653;<td><td>PISCES<td><code>&amp;#x2653;</code><td><code>&amp;#9811;</code><td>
<tr><td>&#x26CE;<td><td>OPHIUCHUS<td><code>&amp;#x26CE;</code><td><code>&amp;#9934;</code><td>
<tr><td>&#x1F500;<td><td>TWISTED RIGHTWARDS ARROWS &#x224A; shuffle tracks button<td><code>&amp;#x1F500;</code><td><code>&amp;#128256;</code><td>
<tr><td>&#x1F501;<td><td>CLOCKWISE RIGHTWARDS AND LEFTWARDS OPEN CIRCLE ARROWS &#x224A; repeat button<td><code>&amp;#x1F501;</code><td><code>&amp;#128257;</code><td>
<tr><td>&#x1F502;<td><td>CLOCKWISE RIGHTWARDS AND LEFTWARDS OPEN CIRCLE ARROWS WITH CIRCLED ONE OVERLAY &#x224A; repeat single button<td><code>&amp;#x1F502;</code><td><code>&amp;#128258;</code><td>
<tr><td>&#x25B6;<td><td>BLACK RIGHT-POINTING TRIANGLE &#x224A; play button<td><code>&amp;#x25B6;</code><td><code>&amp;#9654;</code><td>
<tr><td>&#x23E9;<td><td>BLACK RIGHT-POINTING DOUBLE TRIANGLE &#x224A; fast-forword button<td><code>&amp;#x23E9;</code><td><code>&amp;#9193;</code><td>
<tr><td>&#x23ED;<td><td>BLACK RIGHT-POINTING DOUBLE TRIANGLE WITH VERTICAL BAR &#x224A; next track button<td><code>&amp;#x23ED;</code><td><code>&amp;#9197;</code><td>
<tr><td>&#x23EF;<td><td>BLACK RIGHT-POINTING TRIANGLE WITH DOUBLE VERTICAL BAR &#x224A; play or pause button<td><code>&amp;#x23EF;</code><td><code>&amp;#9199;</code><td>
<tr><td>&#x25C0;<td><td>BLACK LEFT-POINTING TRIANGLE &#x224A; reverse button<td><code>&amp;#x25C0;</code><td><code>&amp;#9664;</code><td>
<tr><td>&#x23EA;<td><td>BLACK LEFT-POINTING DOUBLE TRIANGLE &#x224A; fast reverse button<td><code>&amp;#x23EA;</code><td><code>&amp;#9194;</code><td>
<tr><td>&#x23EE;<td><td>BLACK LEFT-POINTING DOUBLE TRIANGLE WITH VERTICAL BAR &#x224A; last track button<td><code>&amp;#x23EE;</code><td><code>&amp;#9198;</code><td>
<tr><td>&#x1F53C;<td><td>UP-POINTING SMALL RED TRIANGLE &#x224A; up button<td><code>&amp;#x1F53C;</code><td><code>&amp;#128316;</code><td>
<tr><td>&#x23EB;<td><td>BLACK UP-POINTING DOUBLE TRIANGLE &#x224A; fast up button<td><code>&amp;#x23EB;</code><td><code>&amp;#9195;</code><td>
<tr><td>&#x1F53D;<td><td>DOWN-POINTING SMALL RED TRIANGLE &#x224A; down button<td><code>&amp;#x1F53D;</code><td><code>&amp;#128317;</code><td>
<tr><td>&#x23EC;<td><td>BLACK DOWN-POINTING DOUBLE TRIANGLE &#x224A; fast down button<td><code>&amp;#x23EC;</code><td><code>&amp;#9196;</code><td>
<tr><td>&#x23F8;<td><td>DOUBLE VERTICAL BAR &#x224A; pause button<td><code>&amp;#x23F8;</code><td><code>&amp;#9208;</code><td>
<tr><td>&#x23F9;<td><td>BLACK SQUARE FOR STOP &#x224A; stop button<td><code>&amp;#x23F9;</code><td><code>&amp;#9209;</code><td>
<tr><td>&#x23FA;<td><td>BLACK CIRCLE FOR RECORD &#x224A; record button<td><code>&amp;#x23FA;</code><td><code>&amp;#9210;</code><td>
<tr><td>&#x23CF;<td><td>EJECT SYMBOL &#x224A; eject button<td><code>&amp;#x23CF;</code><td><code>&amp;#9167;</code><td>
<tr><td>&#x1F3A6;<td><td>CINEMA<td><code>&amp;#x1F3A6;</code><td><code>&amp;#127910;</code><td>
<tr><td>&#x1F505;<td><td>LOW BRIGHTNESS SYMBOL &#x224A; dim button<td><code>&amp;#x1F505;</code><td><code>&amp;#128261;</code><td>
<tr><td>&#x1F506;<td><td>HIGH BRIGHTNESS SYMBOL &#x224A; bright button<td><code>&amp;#x1F506;</code><td><code>&amp;#128262;</code><td>
<tr><td>&#x1F4F6;<td><td>ANTENNA WITH BARS &#x224A; antenna bars<td><code>&amp;#x1F4F6;</code><td><code>&amp;#128246;</code><td>
<tr><td>&#x1F4F3;<td><td>VIBRATION MODE<td><code>&amp;#x1F4F3;</code><td><code>&amp;#128243;</code><td>
<tr><td>&#x1F4F4;<td><td>MOBILE PHONE OFF<td><code>&amp;#x1F4F4;</code><td><code>&amp;#128244;</code><td>
<tr><td>&#x0023;&#xFE0F;&#x20E3;     <td><td>keycap #<td><code>&amp;#x0023; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#35; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x002A;&#xFE0F;&#x20E3;     <td><td>keycap *<td><code>&amp;#x002A; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#42; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0030;&#xFE0F;&#x20E3;     <td><td>keycap 0<td><code>&amp;#x0030; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#48; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0031;&#xFE0F;&#x20E3;     <td><td>keycap 1<td><code>&amp;#x0031; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#49; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0032;&#xFE0F;&#x20E3;     <td><td>keycap 2<td><code>&amp;#x0032; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#50; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0033;&#xFE0F;&#x20E3;     <td><td>keycap 3<td><code>&amp;#x0033; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#51; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0034;&#xFE0F;&#x20E3;     <td><td>keycap 4<td><code>&amp;#x0034; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#52; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0035;&#xFE0F;&#x20E3;     <td><td>keycap 5<td><code>&amp;#x0035; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#53; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0036;&#xFE0F;&#x20E3;     <td><td>keycap 6<td><code>&amp;#x0036; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#54; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0037;&#xFE0F;&#x20E3;     <td><td>keycap 7<td><code>&amp;#x0037; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#55; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0038;&#xFE0F;&#x20E3;     <td><td>keycap 8<td><code>&amp;#x0038; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#56; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x0039;&#xFE0F;&#x20E3;     <td><td>keycap 9<td><code>&amp;#x0039; &amp;#xFE0F; &amp;#x20E3;     </code><td><code>&amp;#57; &amp;#65039; &amp;#8419;     </code><td>
<tr><td>&#x1F51F;<td><td>KEYCAP TEN &#x224A; keycap #<td><code>&amp;#x1F51F;</code><td><code>&amp;#128287;</code><td>
<tr><td>&#x1F4AF;<td><td>HUNDRED POINTS SYMBOL &#x224A; hundred points<td><code>&amp;#x1F4AF;</code><td><code>&amp;#128175;</code><td>
<tr><td>&#x1F520;<td><td>INPUT SYMBOL FOR LATIN CAPITAL LETTERS &#x224A; input latin uppercase<td><code>&amp;#x1F520;</code><td><code>&amp;#128288;</code><td>
<tr><td>&#x1F521;<td><td>INPUT SYMBOL FOR LATIN SMALL LETTERS &#x224A; input latin lowercase<td><code>&amp;#x1F521;</code><td><code>&amp;#128289;</code><td>
<tr><td>&#x1F522;<td><td>INPUT SYMBOL FOR NUMBERS &#x224A; input numbers<td><code>&amp;#x1F522;</code><td><code>&amp;#128290;</code><td>
<tr><td>&#x1F523;<td><td>INPUT SYMBOL FOR SYMBOLS &#x224A; input symbols<td><code>&amp;#x1F523;</code><td><code>&amp;#128291;</code><td>
<tr><td>&#x1F524;<td><td>INPUT SYMBOL FOR LATIN LETTERS &#x224A; input latin letters<td><code>&amp;#x1F524;</code><td><code>&amp;#128292;</code><td>
<tr><td>&#x1F170;<td><td>NEGATIVE SQUARED LATIN CAPITAL LETTER A &#x224A; a button<td><code>&amp;#x1F170;</code><td><code>&amp;#127344;</code><td>
<tr><td>&#x1F18E;<td><td>NEGATIVE SQUARED AB &#x224A; ab button<td><code>&amp;#x1F18E;</code><td><code>&amp;#127374;</code><td>
<tr><td>&#x1F171;<td><td>NEGATIVE SQUARED LATIN CAPITAL LETTER B &#x224A; b button<td><code>&amp;#x1F171;</code><td><code>&amp;#127345;</code><td>
<tr><td>&#x1F191;<td><td>SQUARED CL<td><code>&amp;#x1F191;</code><td><code>&amp;#127377;</code><td>
<tr><td>&#x1F192;<td><td>SQUARED COOL<td><code>&amp;#x1F192;</code><td><code>&amp;#127378;</code><td>
<tr><td>&#x1F193;<td><td>SQUARED FREE<td><code>&amp;#x1F193;</code><td><code>&amp;#127379;</code><td>
<tr><td>&#x2139;<td><td>INFORMATION SOURCE<td><code>&amp;#x2139;</code><td><code>&amp;#8505;</code><td>
<tr><td>&#x1F194;<td><td>SQUARED ID<td><code>&amp;#x1F194;</code><td><code>&amp;#127380;</code><td>
<tr><td>&#x24C2;<td><td>CIRCLED LATIN CAPITAL LETTER M &#x224A; circled letter m<td><code>&amp;#x24C2;</code><td><code>&amp;#9410;</code><td>
<tr><td>&#x1F195;<td><td>SQUARED NEW<td><code>&amp;#x1F195;</code><td><code>&amp;#127381;</code><td>
<tr><td>&#x1F196;<td><td>SQUARED NG<td><code>&amp;#x1F196;</code><td><code>&amp;#127382;</code><td>
<tr><td>&#x1F17E;<td><td>NEGATIVE SQUARED LATIN CAPITAL LETTER O &#x224A; o button<td><code>&amp;#x1F17E;</code><td><code>&amp;#127358;</code><td>
<tr><td>&#x1F197;<td><td>SQUARED OK<td><code>&amp;#x1F197;</code><td><code>&amp;#127383;</code><td>
<tr><td>&#x1F17F;<td><td>NEGATIVE SQUARED LATIN CAPITAL LETTER P &#x224A; p button<td><code>&amp;#x1F17F;</code><td><code>&amp;#127359;</code><td>
<tr><td>&#x1F198;<td><td>SQUARED SOS<td><code>&amp;#x1F198;</code><td><code>&amp;#127384;</code><td>
<tr><td>&#x1F199;<td><td>SQUARED UP WITH EXCLAMATION MARK &#x224A; up! button<td><code>&amp;#x1F199;</code><td><code>&amp;#127385;</code><td>
<tr><td>&#x1F19A;<td><td>SQUARED VS<td><code>&amp;#x1F19A;</code><td><code>&amp;#127386;</code><td>
<tr><td>&#x1F201;<td><td>SQUARED KATAKANA KOKO<td><code>&amp;#x1F201;</code><td><code>&amp;#127489;</code><td>
<tr><td>&#x1F202;<td><td>SQUARED KATAKANA SA<td><code>&amp;#x1F202;</code><td><code>&amp;#127490;</code><td>
<tr><td>&#x1F237;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-6708 &#x224A; squared moon ideograph<td><code>&amp;#x1F237;</code><td><code>&amp;#127543;</code><td>
<tr><td>&#x1F236;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-6709 &#x224A; squared exist ideograph<td><code>&amp;#x1F236;</code><td><code>&amp;#127542;</code><td>
<tr><td>&#x1F22F;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-6307 &#x224A; squared finger ideograph<td><code>&amp;#x1F22F;</code><td><code>&amp;#127535;</code><td>
<tr><td>&#x1F250;<td><td>CIRCLED IDEOGRAPH ADVANTAGE &#x224A; circled advantage ideograph<td><code>&amp;#x1F250;</code><td><code>&amp;#127568;</code><td>
<tr><td>&#x1F239;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-5272 &#x224A; squared divide ideograph<td><code>&amp;#x1F239;</code><td><code>&amp;#127545;</code><td>
<tr><td>&#x1F21A;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-7121 &#x224A; squared negation ideograph<td><code>&amp;#x1F21A;</code><td><code>&amp;#127514;</code><td>
<tr><td>&#x1F232;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-7981 &#x224A; squared prohibit ideograph<td><code>&amp;#x1F232;</code><td><code>&amp;#127538;</code><td>
<tr><td>&#x1F251;<td><td>CIRCLED IDEOGRAPH ACCEPT &#x224A; circled accept ideograph<td><code>&amp;#x1F251;</code><td><code>&amp;#127569;</code><td>
<tr><td>&#x1F238;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-7533 &#x224A; squared apply ideograph<td><code>&amp;#x1F238;</code><td><code>&amp;#127544;</code><td>
<tr><td>&#x1F234;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-5408 &#x224A; squared together ideograph<td><code>&amp;#x1F234;</code><td><code>&amp;#127540;</code><td>
<tr><td>&#x1F233;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-7A7A &#x224A; squared empty ideograph<td><code>&amp;#x1F233;</code><td><code>&amp;#127539;</code><td>
<tr><td>&#x3297;<td><td>CIRCLED IDEOGRAPH CONGRATULATION &#x224A; circled congratulate ideograph<td><code>&amp;#x3297;</code><td><code>&amp;#12951;</code><td>
<tr><td>&#x3299;<td><td>CIRCLED IDEOGRAPH SECRET &#x224A; circled secret ideograph<td><code>&amp;#x3299;</code><td><code>&amp;#12953;</code><td>
<tr><td>&#x1F23A;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-55B6 &#x224A; squared operating ideograph<td><code>&amp;#x1F23A;</code><td><code>&amp;#127546;</code><td>
<tr><td>&#x1F235;<td><td>SQUARED CJK UNIFIED IDEOGRAPH-6E80 &#x224A; squared fullness ideograph<td><code>&amp;#x1F235;</code><td><code>&amp;#127541;</code><td>
<tr><td>&#x25AA;<td><td>BLACK SMALL SQUARE<td><code>&amp;#x25AA;</code><td><code>&amp;#9642;</code><td>
<tr><td>&#x25AB;<td><td>WHITE SMALL SQUARE<td><code>&amp;#x25AB;</code><td><code>&amp;#9643;</code><td>
<tr><td>&#x25FB;<td><td>WHITE MEDIUM SQUARE<td><code>&amp;#x25FB;</code><td><code>&amp;#9723;</code><td>
<tr><td>&#x25FC;<td><td>BLACK MEDIUM SQUARE<td><code>&amp;#x25FC;</code><td><code>&amp;#9724;</code><td>
<tr><td>&#x25FD;<td><td>WHITE MEDIUM SMALL SQUARE &#x224A; white medium-small square<td><code>&amp;#x25FD;</code><td><code>&amp;#9725;</code><td>
<tr><td>&#x25FE;<td><td>BLACK MEDIUM SMALL SQUARE &#x224A; black medium-small square<td><code>&amp;#x25FE;</code><td><code>&amp;#9726;</code><td>
<tr><td>&#x2B1B;<td><td>BLACK LARGE SQUARE<td><code>&amp;#x2B1B;</code><td><code>&amp;#11035;</code><td>
<tr><td>&#x2B1C;<td><td>WHITE LARGE SQUARE<td><code>&amp;#x2B1C;</code><td><code>&amp;#11036;</code><td>
<tr><td>&#x1F536;<td><td>LARGE ORANGE DIAMOND<td><code>&amp;#x1F536;</code><td><code>&amp;#128310;</code><td>
<tr><td>&#x1F537;<td><td>LARGE BLUE DIAMOND<td><code>&amp;#x1F537;</code><td><code>&amp;#128311;</code><td>
<tr><td>&#x1F538;<td><td>SMALL ORANGE DIAMOND<td><code>&amp;#x1F538;</code><td><code>&amp;#128312;</code><td>
<tr><td>&#x1F539;<td><td>SMALL BLUE DIAMOND<td><code>&amp;#x1F539;</code><td><code>&amp;#128313;</code><td>
<tr><td>&#x1F53A;<td><td>UP-POINTING RED TRIANGLE &#x224A; red triangle pointed up<td><code>&amp;#x1F53A;</code><td><code>&amp;#128314;</code><td>
<tr><td>&#x1F53B;<td><td>DOWN-POINTING RED TRIANGLE &#x224A; red triangle pointed down<td><code>&amp;#x1F53B;</code><td><code>&amp;#128315;</code><td>
<tr><td>&#x1F4A0;<td><td>DIAMOND SHAPE WITH A DOT INSIDE &#x224A; diamond with a dot<td><code>&amp;#x1F4A0;</code><td><code>&amp;#128160;</code><td>
<tr><td>&#x1F518;<td><td>RADIO BUTTON<td><code>&amp;#x1F518;</code><td><code>&amp;#128280;</code><td>
<tr><td>&#x1F532;<td><td>BLACK SQUARE BUTTON<td><code>&amp;#x1F532;</code><td><code>&amp;#128306;</code><td>
<tr><td>&#x1F533;<td><td>WHITE SQUARE BUTTON<td><code>&amp;#x1F533;</code><td><code>&amp;#128307;</code><td>
<tr><td>&#x26AA;<td><td>MEDIUM WHITE CIRCLE &#x224A; white circle<td><code>&amp;#x26AA;</code><td><code>&amp;#9898;</code><td>
<tr><td>&#x26AB;<td><td>MEDIUM BLACK CIRCLE &#x224A; black circle<td><code>&amp;#x26AB;</code><td><code>&amp;#9899;</code><td>
<tr><td>&#x1F534;<td><td>LARGE RED CIRCLE &#x224A; red circle<td><code>&amp;#x1F534;</code><td><code>&amp;#128308;</code><td>
<tr><td>&#x1F535;<td><td>LARGE BLUE CIRCLE &#x224A; blue circle<td><code>&amp;#x1F535;</code><td><code>&amp;#128309;</code><td>
<tr><td>&#x1F3C1;<td><td>CHEQUERED FLAG<td><code>&amp;#x1F3C1;</code><td><code>&amp;#127937;</code><td>
<tr><td>&#x1F6A9;<td><td>TRIANGULAR FLAG ON POST &#x224A; triangular flag<td><code>&amp;#x1F6A9;</code><td><code>&amp;#128681;</code><td>
<tr><td>&#x1F38C;<td><td>CROSSED FLAGS<td><code>&amp;#x1F38C;</code><td><code>&amp;#127884;</code><td>
<tr><td>&#x1F3F4;<td><td>WAVING BLACK FLAG<td><code>&amp;#x1F3F4;</code><td><code>&amp;#127988;</code><td>
<tr><td>&#x1F3F3;<td><td>WAVING WHITE FLAG<td><code>&amp;#x1F3F3;</code><td><code>&amp;#127987;</code><td>
<tr><td>&#x1F1E6;&#x1F1E8;<td><td>Ascension Island<td><code>&amp;#x1F1E6; &amp;#x1F1E8;</code><td><code>&amp;#127462; &amp;#127464;</code><td>
<tr><td>&#x1F1E6;&#x1F1E9;<td><td>Andorra<td><code>&amp;#x1F1E6; &amp;#x1F1E9;</code><td><code>&amp;#127462; &amp;#127465;</code><td>
<tr><td>&#x1F1E6;&#x1F1EA;<td><td>United Arab Emirates<td><code>&amp;#x1F1E6; &amp;#x1F1EA;</code><td><code>&amp;#127462; &amp;#127466;</code><td>
<tr><td>&#x1F1E6;&#x1F1EB;<td><td>Afghanistan<td><code>&amp;#x1F1E6; &amp;#x1F1EB;</code><td><code>&amp;#127462; &amp;#127467;</code><td>
<tr><td>&#x1F1E6;&#x1F1EC;<td><td>Antigua &amp; Barbuda<td><code>&amp;#x1F1E6; &amp;#x1F1EC;</code><td><code>&amp;#127462; &amp;#127468;</code><td>
<tr><td>&#x1F1E6;&#x1F1EE;<td><td>Anguilla<td><code>&amp;#x1F1E6; &amp;#x1F1EE;</code><td><code>&amp;#127462; &amp;#127470;</code><td>
<tr><td>&#x1F1E6;&#x1F1F1;<td><td>Albania<td><code>&amp;#x1F1E6; &amp;#x1F1F1;</code><td><code>&amp;#127462; &amp;#127473;</code><td>
<tr><td>&#x1F1E6;&#x1F1F2;<td><td>Armenia<td><code>&amp;#x1F1E6; &amp;#x1F1F2;</code><td><code>&amp;#127462; &amp;#127474;</code><td>
<tr><td>&#x1F1E6;&#x1F1F4;<td><td>Angola<td><code>&amp;#x1F1E6; &amp;#x1F1F4;</code><td><code>&amp;#127462; &amp;#127476;</code><td>
<tr><td>&#x1F1E6;&#x1F1F6;<td><td>Antarctica<td><code>&amp;#x1F1E6; &amp;#x1F1F6;</code><td><code>&amp;#127462; &amp;#127478;</code><td>
<tr><td>&#x1F1E6;&#x1F1F7;<td><td>Argentina<td><code>&amp;#x1F1E6; &amp;#x1F1F7;</code><td><code>&amp;#127462; &amp;#127479;</code><td>
<tr><td>&#x1F1E6;&#x1F1F8;<td><td>American Samoa<td><code>&amp;#x1F1E6; &amp;#x1F1F8;</code><td><code>&amp;#127462; &amp;#127480;</code><td>
<tr><td>&#x1F1E6;&#x1F1F9;<td><td>Austria<td><code>&amp;#x1F1E6; &amp;#x1F1F9;</code><td><code>&amp;#127462; &amp;#127481;</code><td>
<tr><td>&#x1F1E6;&#x1F1FA;<td><td>Australia<td><code>&amp;#x1F1E6; &amp;#x1F1FA;</code><td><code>&amp;#127462; &amp;#127482;</code><td>
<tr><td>&#x1F1E6;&#x1F1FC;<td><td>Aruba<td><code>&amp;#x1F1E6; &amp;#x1F1FC;</code><td><code>&amp;#127462; &amp;#127484;</code><td>
<tr><td>&#x1F1E6;&#x1F1FD;<td><td>Ã…land Islands<td><code>&amp;#x1F1E6; &amp;#x1F1FD;</code><td><code>&amp;#127462; &amp;#127485;</code><td>
<tr><td>&#x1F1E6;&#x1F1FF;<td><td>Azerbaijan<td><code>&amp;#x1F1E6; &amp;#x1F1FF;</code><td><code>&amp;#127462; &amp;#127487;</code><td>
<tr><td>&#x1F1E7;&#x1F1E6;<td><td>Bosnia &amp; Herzegovina<td><code>&amp;#x1F1E7; &amp;#x1F1E6;</code><td><code>&amp;#127463; &amp;#127462;</code><td>
<tr><td>&#x1F1E7;&#x1F1E7;<td><td>Barbados<td><code>&amp;#x1F1E7; &amp;#x1F1E7;</code><td><code>&amp;#127463; &amp;#127463;</code><td>
<tr><td>&#x1F1E7;&#x1F1E9;<td><td>Bangladesh<td><code>&amp;#x1F1E7; &amp;#x1F1E9;</code><td><code>&amp;#127463; &amp;#127465;</code><td>
<tr><td>&#x1F1E7;&#x1F1EA;<td><td>Belgium<td><code>&amp;#x1F1E7; &amp;#x1F1EA;</code><td><code>&amp;#127463; &amp;#127466;</code><td>
<tr><td>&#x1F1E7;&#x1F1EB;<td><td>Burkina Faso<td><code>&amp;#x1F1E7; &amp;#x1F1EB;</code><td><code>&amp;#127463; &amp;#127467;</code><td>
<tr><td>&#x1F1E7;&#x1F1EC;<td><td>Bulgaria<td><code>&amp;#x1F1E7; &amp;#x1F1EC;</code><td><code>&amp;#127463; &amp;#127468;</code><td>
<tr><td>&#x1F1E7;&#x1F1ED;<td><td>Bahrain<td><code>&amp;#x1F1E7; &amp;#x1F1ED;</code><td><code>&amp;#127463; &amp;#127469;</code><td>
<tr><td>&#x1F1E7;&#x1F1EE;<td><td>Burundi<td><code>&amp;#x1F1E7; &amp;#x1F1EE;</code><td><code>&amp;#127463; &amp;#127470;</code><td>
<tr><td>&#x1F1E7;&#x1F1EF;<td><td>Benin<td><code>&amp;#x1F1E7; &amp;#x1F1EF;</code><td><code>&amp;#127463; &amp;#127471;</code><td>
<tr><td>&#x1F1E7;&#x1F1F1;<td><td>St. BarthÃ©lemy<td><code>&amp;#x1F1E7; &amp;#x1F1F1;</code><td><code>&amp;#127463; &amp;#127473;</code><td>
<tr><td>&#x1F1E7;&#x1F1F2;<td><td>Bermuda<td><code>&amp;#x1F1E7; &amp;#x1F1F2;</code><td><code>&amp;#127463; &amp;#127474;</code><td>
<tr><td>&#x1F1E7;&#x1F1F3;<td><td>Brunei<td><code>&amp;#x1F1E7; &amp;#x1F1F3;</code><td><code>&amp;#127463; &amp;#127475;</code><td>
<tr><td>&#x1F1E7;&#x1F1F4;<td><td>Bolivia<td><code>&amp;#x1F1E7; &amp;#x1F1F4;</code><td><code>&amp;#127463; &amp;#127476;</code><td>
<tr><td>&#x1F1E7;&#x1F1F6;<td><td>Caribbean Netherlands<td><code>&amp;#x1F1E7; &amp;#x1F1F6;</code><td><code>&amp;#127463; &amp;#127478;</code><td>
<tr><td>&#x1F1E7;&#x1F1F7;<td><td>Brazil<td><code>&amp;#x1F1E7; &amp;#x1F1F7;</code><td><code>&amp;#127463; &amp;#127479;</code><td>
<tr><td>&#x1F1E7;&#x1F1F8;<td><td>Bahamas<td><code>&amp;#x1F1E7; &amp;#x1F1F8;</code><td><code>&amp;#127463; &amp;#127480;</code><td>
<tr><td>&#x1F1E7;&#x1F1F9;<td><td>Bhutan<td><code>&amp;#x1F1E7; &amp;#x1F1F9;</code><td><code>&amp;#127463; &amp;#127481;</code><td>
<tr><td>&#x1F1E7;&#x1F1FB;<td><td>Bouvet Island<td><code>&amp;#x1F1E7; &amp;#x1F1FB;</code><td><code>&amp;#127463; &amp;#127483;</code><td>
<tr><td>&#x1F1E7;&#x1F1FC;<td><td>Botswana<td><code>&amp;#x1F1E7; &amp;#x1F1FC;</code><td><code>&amp;#127463; &amp;#127484;</code><td>
<tr><td>&#x1F1E7;&#x1F1FE;<td><td>Belarus<td><code>&amp;#x1F1E7; &amp;#x1F1FE;</code><td><code>&amp;#127463; &amp;#127486;</code><td>
<tr><td>&#x1F1E7;&#x1F1FF;<td><td>Belize<td><code>&amp;#x1F1E7; &amp;#x1F1FF;</code><td><code>&amp;#127463; &amp;#127487;</code><td>
<tr><td>&#x1F1E8;&#x1F1E6;<td><td>Canada<td><code>&amp;#x1F1E8; &amp;#x1F1E6;</code><td><code>&amp;#127464; &amp;#127462;</code><td>
<tr><td>&#x1F1E8;&#x1F1E8;<td><td>Cocos (Keeling) Islands<td><code>&amp;#x1F1E8; &amp;#x1F1E8;</code><td><code>&amp;#127464; &amp;#127464;</code><td>
<tr><td>&#x1F1E8;&#x1F1E9;<td><td>Congo - Kinshasa<td><code>&amp;#x1F1E8; &amp;#x1F1E9;</code><td><code>&amp;#127464; &amp;#127465;</code><td>
<tr><td>&#x1F1E8;&#x1F1EB;<td><td>Central African Republic<td><code>&amp;#x1F1E8; &amp;#x1F1EB;</code><td><code>&amp;#127464; &amp;#127467;</code><td>
<tr><td>&#x1F1E8;&#x1F1EC;<td><td>Congo - Brazzaville<td><code>&amp;#x1F1E8; &amp;#x1F1EC;</code><td><code>&amp;#127464; &amp;#127468;</code><td>
<tr><td>&#x1F1E8;&#x1F1ED;<td><td>Switzerland<td><code>&amp;#x1F1E8; &amp;#x1F1ED;</code><td><code>&amp;#127464; &amp;#127469;</code><td>
<tr><td>&#x1F1E8;&#x1F1EE;<td><td>CÃ´te dâ€™Ivoire<td><code>&amp;#x1F1E8; &amp;#x1F1EE;</code><td><code>&amp;#127464; &amp;#127470;</code><td>
<tr><td>&#x1F1E8;&#x1F1F0;<td><td>Cook Islands<td><code>&amp;#x1F1E8; &amp;#x1F1F0;</code><td><code>&amp;#127464; &amp;#127472;</code><td>
<tr><td>&#x1F1E8;&#x1F1F1;<td><td>Chile<td><code>&amp;#x1F1E8; &amp;#x1F1F1;</code><td><code>&amp;#127464; &amp;#127473;</code><td>
<tr><td>&#x1F1E8;&#x1F1F2;<td><td>Cameroon<td><code>&amp;#x1F1E8; &amp;#x1F1F2;</code><td><code>&amp;#127464; &amp;#127474;</code><td>
<tr><td>&#x1F1E8;&#x1F1F3;<td><td>China<td><code>&amp;#x1F1E8; &amp;#x1F1F3;</code><td><code>&amp;#127464; &amp;#127475;</code><td>
<tr><td>&#x1F1E8;&#x1F1F4;<td><td>Colombia<td><code>&amp;#x1F1E8; &amp;#x1F1F4;</code><td><code>&amp;#127464; &amp;#127476;</code><td>
<tr><td>&#x1F1E8;&#x1F1F5;<td><td>Clipperton Island<td><code>&amp;#x1F1E8; &amp;#x1F1F5;</code><td><code>&amp;#127464; &amp;#127477;</code><td>
<tr><td>&#x1F1E8;&#x1F1F7;<td><td>Costa Rica<td><code>&amp;#x1F1E8; &amp;#x1F1F7;</code><td><code>&amp;#127464; &amp;#127479;</code><td>
<tr><td>&#x1F1E8;&#x1F1FA;<td><td>Cuba<td><code>&amp;#x1F1E8; &amp;#x1F1FA;</code><td><code>&amp;#127464; &amp;#127482;</code><td>
<tr><td>&#x1F1E8;&#x1F1FB;<td><td>Cape Verde<td><code>&amp;#x1F1E8; &amp;#x1F1FB;</code><td><code>&amp;#127464; &amp;#127483;</code><td>
<tr><td>&#x1F1E8;&#x1F1FC;<td><td>CuraÃ§ao<td><code>&amp;#x1F1E8; &amp;#x1F1FC;</code><td><code>&amp;#127464; &amp;#127484;</code><td>
<tr><td>&#x1F1E8;&#x1F1FD;<td><td>Christmas Island<td><code>&amp;#x1F1E8; &amp;#x1F1FD;</code><td><code>&amp;#127464; &amp;#127485;</code><td>
<tr><td>&#x1F1E8;&#x1F1FE;<td><td>Cyprus<td><code>&amp;#x1F1E8; &amp;#x1F1FE;</code><td><code>&amp;#127464; &amp;#127486;</code><td>
<tr><td>&#x1F1E8;&#x1F1FF;<td><td>Czech Republic<td><code>&amp;#x1F1E8; &amp;#x1F1FF;</code><td><code>&amp;#127464; &amp;#127487;</code><td>
<tr><td>&#x1F1E9;&#x1F1EA;<td><td>Germany<td><code>&amp;#x1F1E9; &amp;#x1F1EA;</code><td><code>&amp;#127465; &amp;#127466;</code><td>
<tr><td>&#x1F1E9;&#x1F1EC;<td><td>Diego Garcia<td><code>&amp;#x1F1E9; &amp;#x1F1EC;</code><td><code>&amp;#127465; &amp;#127468;</code><td>
<tr><td>&#x1F1E9;&#x1F1EF;<td><td>Djibouti<td><code>&amp;#x1F1E9; &amp;#x1F1EF;</code><td><code>&amp;#127465; &amp;#127471;</code><td>
<tr><td>&#x1F1E9;&#x1F1F0;<td><td>Denmark<td><code>&amp;#x1F1E9; &amp;#x1F1F0;</code><td><code>&amp;#127465; &amp;#127472;</code><td>
<tr><td>&#x1F1E9;&#x1F1F2;<td><td>Dominica<td><code>&amp;#x1F1E9; &amp;#x1F1F2;</code><td><code>&amp;#127465; &amp;#127474;</code><td>
<tr><td>&#x1F1E9;&#x1F1F4;<td><td>Dominican Republic<td><code>&amp;#x1F1E9; &amp;#x1F1F4;</code><td><code>&amp;#127465; &amp;#127476;</code><td>
<tr><td>&#x1F1E9;&#x1F1FF;<td><td>Algeria<td><code>&amp;#x1F1E9; &amp;#x1F1FF;</code><td><code>&amp;#127465; &amp;#127487;</code><td>
<tr><td>&#x1F1EA;&#x1F1E6;<td><td>Ceuta &amp; Melilla<td><code>&amp;#x1F1EA; &amp;#x1F1E6;</code><td><code>&amp;#127466; &amp;#127462;</code><td>
<tr><td>&#x1F1EA;&#x1F1E8;<td><td>Ecuador<td><code>&amp;#x1F1EA; &amp;#x1F1E8;</code><td><code>&amp;#127466; &amp;#127464;</code><td>
<tr><td>&#x1F1EA;&#x1F1EA;<td><td>Estonia<td><code>&amp;#x1F1EA; &amp;#x1F1EA;</code><td><code>&amp;#127466; &amp;#127466;</code><td>
<tr><td>&#x1F1EA;&#x1F1EC;<td><td>Egypt<td><code>&amp;#x1F1EA; &amp;#x1F1EC;</code><td><code>&amp;#127466; &amp;#127468;</code><td>
<tr><td>&#x1F1EA;&#x1F1ED;<td><td>Western Sahara<td><code>&amp;#x1F1EA; &amp;#x1F1ED;</code><td><code>&amp;#127466; &amp;#127469;</code><td>
<tr><td>&#x1F1EA;&#x1F1F7;<td><td>Eritrea<td><code>&amp;#x1F1EA; &amp;#x1F1F7;</code><td><code>&amp;#127466; &amp;#127479;</code><td>
<tr><td>&#x1F1EA;&#x1F1F8;<td><td>Spain<td><code>&amp;#x1F1EA; &amp;#x1F1F8;</code><td><code>&amp;#127466; &amp;#127480;</code><td>
<tr><td>&#x1F1EA;&#x1F1F9;<td><td>Ethiopia<td><code>&amp;#x1F1EA; &amp;#x1F1F9;</code><td><code>&amp;#127466; &amp;#127481;</code><td>
<tr><td>&#x1F1EA;&#x1F1FA;<td><td>European Union<td><code>&amp;#x1F1EA; &amp;#x1F1FA;</code><td><code>&amp;#127466; &amp;#127482;</code><td>
<tr><td>&#x1F1EB;&#x1F1EE;<td><td>Finland<td><code>&amp;#x1F1EB; &amp;#x1F1EE;</code><td><code>&amp;#127467; &amp;#127470;</code><td>
<tr><td>&#x1F1EB;&#x1F1EF;<td><td>Fiji<td><code>&amp;#x1F1EB; &amp;#x1F1EF;</code><td><code>&amp;#127467; &amp;#127471;</code><td>
<tr><td>&#x1F1EB;&#x1F1F0;<td><td>Falkland Islands<td><code>&amp;#x1F1EB; &amp;#x1F1F0;</code><td><code>&amp;#127467; &amp;#127472;</code><td>
<tr><td>&#x1F1EB;&#x1F1F2;<td><td>Micronesia<td><code>&amp;#x1F1EB; &amp;#x1F1F2;</code><td><code>&amp;#127467; &amp;#127474;</code><td>
<tr><td>&#x1F1EB;&#x1F1F4;<td><td>Faroe Islands<td><code>&amp;#x1F1EB; &amp;#x1F1F4;</code><td><code>&amp;#127467; &amp;#127476;</code><td>
<tr><td>&#x1F1EB;&#x1F1F7;<td><td>France<td><code>&amp;#x1F1EB; &amp;#x1F1F7;</code><td><code>&amp;#127467; &amp;#127479;</code><td>
<tr><td>&#x1F1EC;&#x1F1E6;<td><td>Gabon<td><code>&amp;#x1F1EC; &amp;#x1F1E6;</code><td><code>&amp;#127468; &amp;#127462;</code><td>
<tr><td>&#x1F1EC;&#x1F1E7;<td><td>United Kingdom<td><code>&amp;#x1F1EC; &amp;#x1F1E7;</code><td><code>&amp;#127468; &amp;#127463;</code><td>
<tr><td>&#x1F1EC;&#x1F1E9;<td><td>Grenada<td><code>&amp;#x1F1EC; &amp;#x1F1E9;</code><td><code>&amp;#127468; &amp;#127465;</code><td>
<tr><td>&#x1F1EC;&#x1F1EA;<td><td>Georgia<td><code>&amp;#x1F1EC; &amp;#x1F1EA;</code><td><code>&amp;#127468; &amp;#127466;</code><td>
<tr><td>&#x1F1EC;&#x1F1EB;<td><td>French Guiana<td><code>&amp;#x1F1EC; &amp;#x1F1EB;</code><td><code>&amp;#127468; &amp;#127467;</code><td>
<tr><td>&#x1F1EC;&#x1F1EC;<td><td>Guernsey<td><code>&amp;#x1F1EC; &amp;#x1F1EC;</code><td><code>&amp;#127468; &amp;#127468;</code><td>
<tr><td>&#x1F1EC;&#x1F1ED;<td><td>Ghana<td><code>&amp;#x1F1EC; &amp;#x1F1ED;</code><td><code>&amp;#127468; &amp;#127469;</code><td>
<tr><td>&#x1F1EC;&#x1F1EE;<td><td>Gibraltar<td><code>&amp;#x1F1EC; &amp;#x1F1EE;</code><td><code>&amp;#127468; &amp;#127470;</code><td>
<tr><td>&#x1F1EC;&#x1F1F1;<td><td>Greenland<td><code>&amp;#x1F1EC; &amp;#x1F1F1;</code><td><code>&amp;#127468; &amp;#127473;</code><td>
<tr><td>&#x1F1EC;&#x1F1F2;<td><td>Gambia<td><code>&amp;#x1F1EC; &amp;#x1F1F2;</code><td><code>&amp;#127468; &amp;#127474;</code><td>
<tr><td>&#x1F1EC;&#x1F1F3;<td><td>Guinea<td><code>&amp;#x1F1EC; &amp;#x1F1F3;</code><td><code>&amp;#127468; &amp;#127475;</code><td>
<tr><td>&#x1F1EC;&#x1F1F5;<td><td>Guadeloupe<td><code>&amp;#x1F1EC; &amp;#x1F1F5;</code><td><code>&amp;#127468; &amp;#127477;</code><td>
<tr><td>&#x1F1EC;&#x1F1F6;<td><td>Equatorial Guinea<td><code>&amp;#x1F1EC; &amp;#x1F1F6;</code><td><code>&amp;#127468; &amp;#127478;</code><td>
<tr><td>&#x1F1EC;&#x1F1F7;<td><td>Greece<td><code>&amp;#x1F1EC; &amp;#x1F1F7;</code><td><code>&amp;#127468; &amp;#127479;</code><td>
<tr><td>&#x1F1EC;&#x1F1F8;<td><td>South Georgia &amp; South Sandwich Islands<td><code>&amp;#x1F1EC; &amp;#x1F1F8;</code><td><code>&amp;#127468; &amp;#127480;</code><td>
<tr><td>&#x1F1EC;&#x1F1F9;<td><td>Guatemala<td><code>&amp;#x1F1EC; &amp;#x1F1F9;</code><td><code>&amp;#127468; &amp;#127481;</code><td>
<tr><td>&#x1F1EC;&#x1F1FA;<td><td>Guam<td><code>&amp;#x1F1EC; &amp;#x1F1FA;</code><td><code>&amp;#127468; &amp;#127482;</code><td>
<tr><td>&#x1F1EC;&#x1F1FC;<td><td>Guinea-Bissau<td><code>&amp;#x1F1EC; &amp;#x1F1FC;</code><td><code>&amp;#127468; &amp;#127484;</code><td>
<tr><td>&#x1F1EC;&#x1F1FE;<td><td>Guyana<td><code>&amp;#x1F1EC; &amp;#x1F1FE;</code><td><code>&amp;#127468; &amp;#127486;</code><td>
<tr><td>&#x1F1ED;&#x1F1F0;<td><td>Hong Kong SAR China<td><code>&amp;#x1F1ED; &amp;#x1F1F0;</code><td><code>&amp;#127469; &amp;#127472;</code><td>
<tr><td>&#x1F1ED;&#x1F1F2;<td><td>Heard &amp; McDonald Islands<td><code>&amp;#x1F1ED; &amp;#x1F1F2;</code><td><code>&amp;#127469; &amp;#127474;</code><td>
<tr><td>&#x1F1ED;&#x1F1F3;<td><td>Honduras<td><code>&amp;#x1F1ED; &amp;#x1F1F3;</code><td><code>&amp;#127469; &amp;#127475;</code><td>
<tr><td>&#x1F1ED;&#x1F1F7;<td><td>Croatia<td><code>&amp;#x1F1ED; &amp;#x1F1F7;</code><td><code>&amp;#127469; &amp;#127479;</code><td>
<tr><td>&#x1F1ED;&#x1F1F9;<td><td>Haiti<td><code>&amp;#x1F1ED; &amp;#x1F1F9;</code><td><code>&amp;#127469; &amp;#127481;</code><td>
<tr><td>&#x1F1ED;&#x1F1FA;<td><td>Hungary<td><code>&amp;#x1F1ED; &amp;#x1F1FA;</code><td><code>&amp;#127469; &amp;#127482;</code><td>
<tr><td>&#x1F1EE;&#x1F1E8;<td><td>Canary Islands<td><code>&amp;#x1F1EE; &amp;#x1F1E8;</code><td><code>&amp;#127470; &amp;#127464;</code><td>
<tr><td>&#x1F1EE;&#x1F1E9;<td><td>Indonesia<td><code>&amp;#x1F1EE; &amp;#x1F1E9;</code><td><code>&amp;#127470; &amp;#127465;</code><td>
<tr><td>&#x1F1EE;&#x1F1EA;<td><td>Ireland<td><code>&amp;#x1F1EE; &amp;#x1F1EA;</code><td><code>&amp;#127470; &amp;#127466;</code><td>
<tr><td>&#x1F1EE;&#x1F1F1;<td><td>Israel<td><code>&amp;#x1F1EE; &amp;#x1F1F1;</code><td><code>&amp;#127470; &amp;#127473;</code><td>
<tr><td>&#x1F1EE;&#x1F1F2;<td><td>Isle of Man<td><code>&amp;#x1F1EE; &amp;#x1F1F2;</code><td><code>&amp;#127470; &amp;#127474;</code><td>
<tr><td>&#x1F1EE;&#x1F1F3;<td><td>India<td><code>&amp;#x1F1EE; &amp;#x1F1F3;</code><td><code>&amp;#127470; &amp;#127475;</code><td>
<tr><td>&#x1F1EE;&#x1F1F4;<td><td>British Indian Ocean Territory<td><code>&amp;#x1F1EE; &amp;#x1F1F4;</code><td><code>&amp;#127470; &amp;#127476;</code><td>
<tr><td>&#x1F1EE;&#x1F1F6;<td><td>Iraq<td><code>&amp;#x1F1EE; &amp;#x1F1F6;</code><td><code>&amp;#127470; &amp;#127478;</code><td>
<tr><td>&#x1F1EE;&#x1F1F7;<td><td>Iran<td><code>&amp;#x1F1EE; &amp;#x1F1F7;</code><td><code>&amp;#127470; &amp;#127479;</code><td>
<tr><td>&#x1F1EE;&#x1F1F8;<td><td>Iceland<td><code>&amp;#x1F1EE; &amp;#x1F1F8;</code><td><code>&amp;#127470; &amp;#127480;</code><td>
<tr><td>&#x1F1EE;&#x1F1F9;<td><td>Italy<td><code>&amp;#x1F1EE; &amp;#x1F1F9;</code><td><code>&amp;#127470; &amp;#127481;</code><td>
<tr><td>&#x1F1EF;&#x1F1EA;<td><td>Jersey<td><code>&amp;#x1F1EF; &amp;#x1F1EA;</code><td><code>&amp;#127471; &amp;#127466;</code><td>
<tr><td>&#x1F1EF;&#x1F1F2;<td><td>Jamaica<td><code>&amp;#x1F1EF; &amp;#x1F1F2;</code><td><code>&amp;#127471; &amp;#127474;</code><td>
<tr><td>&#x1F1EF;&#x1F1F4;<td><td>Jordan<td><code>&amp;#x1F1EF; &amp;#x1F1F4;</code><td><code>&amp;#127471; &amp;#127476;</code><td>
<tr><td>&#x1F1EF;&#x1F1F5;<td><td>Japan<td><code>&amp;#x1F1EF; &amp;#x1F1F5;</code><td><code>&amp;#127471; &amp;#127477;</code><td>
<tr><td>&#x1F1F0;&#x1F1EA;<td><td>Kenya<td><code>&amp;#x1F1F0; &amp;#x1F1EA;</code><td><code>&amp;#127472; &amp;#127466;</code><td>
<tr><td>&#x1F1F0;&#x1F1EC;<td><td>Kyrgyzstan<td><code>&amp;#x1F1F0; &amp;#x1F1EC;</code><td><code>&amp;#127472; &amp;#127468;</code><td>
<tr><td>&#x1F1F0;&#x1F1ED;<td><td>Cambodia<td><code>&amp;#x1F1F0; &amp;#x1F1ED;</code><td><code>&amp;#127472; &amp;#127469;</code><td>
<tr><td>&#x1F1F0;&#x1F1EE;<td><td>Kiribati<td><code>&amp;#x1F1F0; &amp;#x1F1EE;</code><td><code>&amp;#127472; &amp;#127470;</code><td>
<tr><td>&#x1F1F0;&#x1F1F2;<td><td>Comoros<td><code>&amp;#x1F1F0; &amp;#x1F1F2;</code><td><code>&amp;#127472; &amp;#127474;</code><td>
<tr><td>&#x1F1F0;&#x1F1F3;<td><td>St. Kitts &amp; Nevis<td><code>&amp;#x1F1F0; &amp;#x1F1F3;</code><td><code>&amp;#127472; &amp;#127475;</code><td>
<tr><td>&#x1F1F0;&#x1F1F5;<td><td>North Korea<td><code>&amp;#x1F1F0; &amp;#x1F1F5;</code><td><code>&amp;#127472; &amp;#127477;</code><td>
<tr><td>&#x1F1F0;&#x1F1F7;<td><td>South Korea<td><code>&amp;#x1F1F0; &amp;#x1F1F7;</code><td><code>&amp;#127472; &amp;#127479;</code><td>
<tr><td>&#x1F1F0;&#x1F1FC;<td><td>Kuwait<td><code>&amp;#x1F1F0; &amp;#x1F1FC;</code><td><code>&amp;#127472; &amp;#127484;</code><td>
<tr><td>&#x1F1F0;&#x1F1FE;<td><td>Cayman Islands<td><code>&amp;#x1F1F0; &amp;#x1F1FE;</code><td><code>&amp;#127472; &amp;#127486;</code><td>
<tr><td>&#x1F1F0;&#x1F1FF;<td><td>Kazakhstan<td><code>&amp;#x1F1F0; &amp;#x1F1FF;</code><td><code>&amp;#127472; &amp;#127487;</code><td>
<tr><td>&#x1F1F1;&#x1F1E6;<td><td>Laos<td><code>&amp;#x1F1F1; &amp;#x1F1E6;</code><td><code>&amp;#127473; &amp;#127462;</code><td>
<tr><td>&#x1F1F1;&#x1F1E7;<td><td>Lebanon<td><code>&amp;#x1F1F1; &amp;#x1F1E7;</code><td><code>&amp;#127473; &amp;#127463;</code><td>
<tr><td>&#x1F1F1;&#x1F1E8;<td><td>St. Lucia<td><code>&amp;#x1F1F1; &amp;#x1F1E8;</code><td><code>&amp;#127473; &amp;#127464;</code><td>
<tr><td>&#x1F1F1;&#x1F1EE;<td><td>Liechtenstein<td><code>&amp;#x1F1F1; &amp;#x1F1EE;</code><td><code>&amp;#127473; &amp;#127470;</code><td>
<tr><td>&#x1F1F1;&#x1F1F0;<td><td>Sri Lanka<td><code>&amp;#x1F1F1; &amp;#x1F1F0;</code><td><code>&amp;#127473; &amp;#127472;</code><td>
<tr><td>&#x1F1F1;&#x1F1F7;<td><td>Liberia<td><code>&amp;#x1F1F1; &amp;#x1F1F7;</code><td><code>&amp;#127473; &amp;#127479;</code><td>
<tr><td>&#x1F1F1;&#x1F1F8;<td><td>Lesotho<td><code>&amp;#x1F1F1; &amp;#x1F1F8;</code><td><code>&amp;#127473; &amp;#127480;</code><td>
<tr><td>&#x1F1F1;&#x1F1F9;<td><td>Lithuania<td><code>&amp;#x1F1F1; &amp;#x1F1F9;</code><td><code>&amp;#127473; &amp;#127481;</code><td>
<tr><td>&#x1F1F1;&#x1F1FA;<td><td>Luxembourg<td><code>&amp;#x1F1F1; &amp;#x1F1FA;</code><td><code>&amp;#127473; &amp;#127482;</code><td>
<tr><td>&#x1F1F1;&#x1F1FB;<td><td>Latvia<td><code>&amp;#x1F1F1; &amp;#x1F1FB;</code><td><code>&amp;#127473; &amp;#127483;</code><td>
<tr><td>&#x1F1F1;&#x1F1FE;<td><td>Libya<td><code>&amp;#x1F1F1; &amp;#x1F1FE;</code><td><code>&amp;#127473; &amp;#127486;</code><td>
<tr><td>&#x1F1F2;&#x1F1E6;<td><td>Morocco<td><code>&amp;#x1F1F2; &amp;#x1F1E6;</code><td><code>&amp;#127474; &amp;#127462;</code><td>
<tr><td>&#x1F1F2;&#x1F1E8;<td><td>Monaco<td><code>&amp;#x1F1F2; &amp;#x1F1E8;</code><td><code>&amp;#127474; &amp;#127464;</code><td>
<tr><td>&#x1F1F2;&#x1F1E9;<td><td>Moldova<td><code>&amp;#x1F1F2; &amp;#x1F1E9;</code><td><code>&amp;#127474; &amp;#127465;</code><td>
<tr><td>&#x1F1F2;&#x1F1EA;<td><td>Montenegro<td><code>&amp;#x1F1F2; &amp;#x1F1EA;</code><td><code>&amp;#127474; &amp;#127466;</code><td>
<tr><td>&#x1F1F2;&#x1F1EB;<td><td>St. Martin<td><code>&amp;#x1F1F2; &amp;#x1F1EB;</code><td><code>&amp;#127474; &amp;#127467;</code><td>
<tr><td>&#x1F1F2;&#x1F1EC;<td><td>Madagascar<td><code>&amp;#x1F1F2; &amp;#x1F1EC;</code><td><code>&amp;#127474; &amp;#127468;</code><td>
<tr><td>&#x1F1F2;&#x1F1ED;<td><td>Marshall Islands<td><code>&amp;#x1F1F2; &amp;#x1F1ED;</code><td><code>&amp;#127474; &amp;#127469;</code><td>
<tr><td>&#x1F1F2;&#x1F1F0;<td><td>Macedonia<td><code>&amp;#x1F1F2; &amp;#x1F1F0;</code><td><code>&amp;#127474; &amp;#127472;</code><td>
<tr><td>&#x1F1F2;&#x1F1F1;<td><td>Mali<td><code>&amp;#x1F1F2; &amp;#x1F1F1;</code><td><code>&amp;#127474; &amp;#127473;</code><td>
<tr><td>&#x1F1F2;&#x1F1F2;<td><td>Myanmar (Burma)<td><code>&amp;#x1F1F2; &amp;#x1F1F2;</code><td><code>&amp;#127474; &amp;#127474;</code><td>
<tr><td>&#x1F1F2;&#x1F1F3;<td><td>Mongolia<td><code>&amp;#x1F1F2; &amp;#x1F1F3;</code><td><code>&amp;#127474; &amp;#127475;</code><td>
<tr><td>&#x1F1F2;&#x1F1F4;<td><td>Macau SAR China<td><code>&amp;#x1F1F2; &amp;#x1F1F4;</code><td><code>&amp;#127474; &amp;#127476;</code><td>
<tr><td>&#x1F1F2;&#x1F1F5;<td><td>Northern Mariana Islands<td><code>&amp;#x1F1F2; &amp;#x1F1F5;</code><td><code>&amp;#127474; &amp;#127477;</code><td>
<tr><td>&#x1F1F2;&#x1F1F6;<td><td>Martinique<td><code>&amp;#x1F1F2; &amp;#x1F1F6;</code><td><code>&amp;#127474; &amp;#127478;</code><td>
<tr><td>&#x1F1F2;&#x1F1F7;<td><td>Mauritania<td><code>&amp;#x1F1F2; &amp;#x1F1F7;</code><td><code>&amp;#127474; &amp;#127479;</code><td>
<tr><td>&#x1F1F2;&#x1F1F8;<td><td>Montserrat<td><code>&amp;#x1F1F2; &amp;#x1F1F8;</code><td><code>&amp;#127474; &amp;#127480;</code><td>
<tr><td>&#x1F1F2;&#x1F1F9;<td><td>Malta<td><code>&amp;#x1F1F2; &amp;#x1F1F9;</code><td><code>&amp;#127474; &amp;#127481;</code><td>
<tr><td>&#x1F1F2;&#x1F1FA;<td><td>Mauritius<td><code>&amp;#x1F1F2; &amp;#x1F1FA;</code><td><code>&amp;#127474; &amp;#127482;</code><td>
<tr><td>&#x1F1F2;&#x1F1FB;<td><td>Maldives<td><code>&amp;#x1F1F2; &amp;#x1F1FB;</code><td><code>&amp;#127474; &amp;#127483;</code><td>
<tr><td>&#x1F1F2;&#x1F1FC;<td><td>Malawi<td><code>&amp;#x1F1F2; &amp;#x1F1FC;</code><td><code>&amp;#127474; &amp;#127484;</code><td>
<tr><td>&#x1F1F2;&#x1F1FD;<td><td>Mexico<td><code>&amp;#x1F1F2; &amp;#x1F1FD;</code><td><code>&amp;#127474; &amp;#127485;</code><td>
<tr><td>&#x1F1F2;&#x1F1FE;<td><td>Malaysia<td><code>&amp;#x1F1F2; &amp;#x1F1FE;</code><td><code>&amp;#127474; &amp;#127486;</code><td>
<tr><td>&#x1F1F2;&#x1F1FF;<td><td>Mozambique<td><code>&amp;#x1F1F2; &amp;#x1F1FF;</code><td><code>&amp;#127474; &amp;#127487;</code><td>
<tr><td>&#x1F1F3;&#x1F1E6;<td><td>Namibia<td><code>&amp;#x1F1F3; &amp;#x1F1E6;</code><td><code>&amp;#127475; &amp;#127462;</code><td>
<tr><td>&#x1F1F3;&#x1F1E8;<td><td>New Caledonia<td><code>&amp;#x1F1F3; &amp;#x1F1E8;</code><td><code>&amp;#127475; &amp;#127464;</code><td>
<tr><td>&#x1F1F3;&#x1F1EA;<td><td>Niger<td><code>&amp;#x1F1F3; &amp;#x1F1EA;</code><td><code>&amp;#127475; &amp;#127466;</code><td>
<tr><td>&#x1F1F3;&#x1F1EB;<td><td>Norfolk Island<td><code>&amp;#x1F1F3; &amp;#x1F1EB;</code><td><code>&amp;#127475; &amp;#127467;</code><td>
<tr><td>&#x1F1F3;&#x1F1EC;<td><td>Nigeria<td><code>&amp;#x1F1F3; &amp;#x1F1EC;</code><td><code>&amp;#127475; &amp;#127468;</code><td>
<tr><td>&#x1F1F3;&#x1F1EE;<td><td>Nicaragua<td><code>&amp;#x1F1F3; &amp;#x1F1EE;</code><td><code>&amp;#127475; &amp;#127470;</code><td>
<tr><td>&#x1F1F3;&#x1F1F1;<td><td>Netherlands<td><code>&amp;#x1F1F3; &amp;#x1F1F1;</code><td><code>&amp;#127475; &amp;#127473;</code><td>
<tr><td>&#x1F1F3;&#x1F1F4;<td><td>Norway<td><code>&amp;#x1F1F3; &amp;#x1F1F4;</code><td><code>&amp;#127475; &amp;#127476;</code><td>
<tr><td>&#x1F1F3;&#x1F1F5;<td><td>Nepal<td><code>&amp;#x1F1F3; &amp;#x1F1F5;</code><td><code>&amp;#127475; &amp;#127477;</code><td>
<tr><td>&#x1F1F3;&#x1F1F7;<td><td>Nauru<td><code>&amp;#x1F1F3; &amp;#x1F1F7;</code><td><code>&amp;#127475; &amp;#127479;</code><td>
<tr><td>&#x1F1F3;&#x1F1FA;<td><td>Niue<td><code>&amp;#x1F1F3; &amp;#x1F1FA;</code><td><code>&amp;#127475; &amp;#127482;</code><td>
<tr><td>&#x1F1F3;&#x1F1FF;<td><td>New Zealand<td><code>&amp;#x1F1F3; &amp;#x1F1FF;</code><td><code>&amp;#127475; &amp;#127487;</code><td>
<tr><td>&#x1F1F4;&#x1F1F2;<td><td>Oman<td><code>&amp;#x1F1F4; &amp;#x1F1F2;</code><td><code>&amp;#127476; &amp;#127474;</code><td>
<tr><td>&#x1F1F5;&#x1F1E6;<td><td>Panama<td><code>&amp;#x1F1F5; &amp;#x1F1E6;</code><td><code>&amp;#127477; &amp;#127462;</code><td>
<tr><td>&#x1F1F5;&#x1F1EA;<td><td>Peru<td><code>&amp;#x1F1F5; &amp;#x1F1EA;</code><td><code>&amp;#127477; &amp;#127466;</code><td>
<tr><td>&#x1F1F5;&#x1F1EB;<td><td>French Polynesia<td><code>&amp;#x1F1F5; &amp;#x1F1EB;</code><td><code>&amp;#127477; &amp;#127467;</code><td>
<tr><td>&#x1F1F5;&#x1F1EC;<td><td>Papua New Guinea<td><code>&amp;#x1F1F5; &amp;#x1F1EC;</code><td><code>&amp;#127477; &amp;#127468;</code><td>
<tr><td>&#x1F1F5;&#x1F1ED;<td><td>Philippines<td><code>&amp;#x1F1F5; &amp;#x1F1ED;</code><td><code>&amp;#127477; &amp;#127469;</code><td>
<tr><td>&#x1F1F5;&#x1F1F0;<td><td>Pakistan<td><code>&amp;#x1F1F5; &amp;#x1F1F0;</code><td><code>&amp;#127477; &amp;#127472;</code><td>
<tr><td>&#x1F1F5;&#x1F1F1;<td><td>Poland<td><code>&amp;#x1F1F5; &amp;#x1F1F1;</code><td><code>&amp;#127477; &amp;#127473;</code><td>
<tr><td>&#x1F1F5;&#x1F1F2;<td><td>St. Pierre &amp; Miquelon<td><code>&amp;#x1F1F5; &amp;#x1F1F2;</code><td><code>&amp;#127477; &amp;#127474;</code><td>
<tr><td>&#x1F1F5;&#x1F1F3;<td><td>Pitcairn Islands<td><code>&amp;#x1F1F5; &amp;#x1F1F3;</code><td><code>&amp;#127477; &amp;#127475;</code><td>
<tr><td>&#x1F1F5;&#x1F1F7;<td><td>Puerto Rico<td><code>&amp;#x1F1F5; &amp;#x1F1F7;</code><td><code>&amp;#127477; &amp;#127479;</code><td>
<tr><td>&#x1F1F5;&#x1F1F8;<td><td>Palestinian Territories<td><code>&amp;#x1F1F5; &amp;#x1F1F8;</code><td><code>&amp;#127477; &amp;#127480;</code><td>
<tr><td>&#x1F1F5;&#x1F1F9;<td><td>Portugal<td><code>&amp;#x1F1F5; &amp;#x1F1F9;</code><td><code>&amp;#127477; &amp;#127481;</code><td>
<tr><td>&#x1F1F5;&#x1F1FC;<td><td>Palau<td><code>&amp;#x1F1F5; &amp;#x1F1FC;</code><td><code>&amp;#127477; &amp;#127484;</code><td>
<tr><td>&#x1F1F5;&#x1F1FE;<td><td>Paraguay<td><code>&amp;#x1F1F5; &amp;#x1F1FE;</code><td><code>&amp;#127477; &amp;#127486;</code><td>
<tr><td>&#x1F1F6;&#x1F1E6;<td><td>Qatar<td><code>&amp;#x1F1F6; &amp;#x1F1E6;</code><td><code>&amp;#127478; &amp;#127462;</code><td>
<tr><td>&#x1F1F7;&#x1F1EA;<td><td>RÃ©union<td><code>&amp;#x1F1F7; &amp;#x1F1EA;</code><td><code>&amp;#127479; &amp;#127466;</code><td>
<tr><td>&#x1F1F7;&#x1F1F4;<td><td>Romania<td><code>&amp;#x1F1F7; &amp;#x1F1F4;</code><td><code>&amp;#127479; &amp;#127476;</code><td>
<tr><td>&#x1F1F7;&#x1F1F8;<td><td>Serbia<td><code>&amp;#x1F1F7; &amp;#x1F1F8;</code><td><code>&amp;#127479; &amp;#127480;</code><td>
<tr><td>&#x1F1F7;&#x1F1FA;<td><td>Russia<td><code>&amp;#x1F1F7; &amp;#x1F1FA;</code><td><code>&amp;#127479; &amp;#127482;</code><td>
<tr><td>&#x1F1F7;&#x1F1FC;<td><td>Rwanda<td><code>&amp;#x1F1F7; &amp;#x1F1FC;</code><td><code>&amp;#127479; &amp;#127484;</code><td>
<tr><td>&#x1F1F8;&#x1F1E6;<td><td>Saudi Arabia<td><code>&amp;#x1F1F8; &amp;#x1F1E6;</code><td><code>&amp;#127480; &amp;#127462;</code><td>
<tr><td>&#x1F1F8;&#x1F1E7;<td><td>Solomon Islands<td><code>&amp;#x1F1F8; &amp;#x1F1E7;</code><td><code>&amp;#127480; &amp;#127463;</code><td>
<tr><td>&#x1F1F8;&#x1F1E8;<td><td>Seychelles<td><code>&amp;#x1F1F8; &amp;#x1F1E8;</code><td><code>&amp;#127480; &amp;#127464;</code><td>
<tr><td>&#x1F1F8;&#x1F1E9;<td><td>Sudan<td><code>&amp;#x1F1F8; &amp;#x1F1E9;</code><td><code>&amp;#127480; &amp;#127465;</code><td>
<tr><td>&#x1F1F8;&#x1F1EA;<td><td>Sweden<td><code>&amp;#x1F1F8; &amp;#x1F1EA;</code><td><code>&amp;#127480; &amp;#127466;</code><td>
<tr><td>&#x1F1F8;&#x1F1EC;<td><td>Singapore<td><code>&amp;#x1F1F8; &amp;#x1F1EC;</code><td><code>&amp;#127480; &amp;#127468;</code><td>
<tr><td>&#x1F1F8;&#x1F1ED;<td><td>St. Helena<td><code>&amp;#x1F1F8; &amp;#x1F1ED;</code><td><code>&amp;#127480; &amp;#127469;</code><td>
<tr><td>&#x1F1F8;&#x1F1EE;<td><td>Slovenia<td><code>&amp;#x1F1F8; &amp;#x1F1EE;</code><td><code>&amp;#127480; &amp;#127470;</code><td>
<tr><td>&#x1F1F8;&#x1F1EF;<td><td>Svalbard &amp; Jan Mayen<td><code>&amp;#x1F1F8; &amp;#x1F1EF;</code><td><code>&amp;#127480; &amp;#127471;</code><td>
<tr><td>&#x1F1F8;&#x1F1F0;<td><td>Slovakia<td><code>&amp;#x1F1F8; &amp;#x1F1F0;</code><td><code>&amp;#127480; &amp;#127472;</code><td>
<tr><td>&#x1F1F8;&#x1F1F1;<td><td>Sierra Leone<td><code>&amp;#x1F1F8; &amp;#x1F1F1;</code><td><code>&amp;#127480; &amp;#127473;</code><td>
<tr><td>&#x1F1F8;&#x1F1F2;<td><td>San Marino<td><code>&amp;#x1F1F8; &amp;#x1F1F2;</code><td><code>&amp;#127480; &amp;#127474;</code><td>
<tr><td>&#x1F1F8;&#x1F1F3;<td><td>Senegal<td><code>&amp;#x1F1F8; &amp;#x1F1F3;</code><td><code>&amp;#127480; &amp;#127475;</code><td>
<tr><td>&#x1F1F8;&#x1F1F4;<td><td>Somalia<td><code>&amp;#x1F1F8; &amp;#x1F1F4;</code><td><code>&amp;#127480; &amp;#127476;</code><td>
<tr><td>&#x1F1F8;&#x1F1F7;<td><td>Suriname<td><code>&amp;#x1F1F8; &amp;#x1F1F7;</code><td><code>&amp;#127480; &amp;#127479;</code><td>
<tr><td>&#x1F1F8;&#x1F1F8;<td><td>South Sudan<td><code>&amp;#x1F1F8; &amp;#x1F1F8;</code><td><code>&amp;#127480; &amp;#127480;</code><td>
<tr><td>&#x1F1F8;&#x1F1F9;<td><td>SÃ£o TomÃ© &amp; PrÃ­ncipe<td><code>&amp;#x1F1F8; &amp;#x1F1F9;</code><td><code>&amp;#127480; &amp;#127481;</code><td>
<tr><td>&#x1F1F8;&#x1F1FB;<td><td>El Salvador<td><code>&amp;#x1F1F8; &amp;#x1F1FB;</code><td><code>&amp;#127480; &amp;#127483;</code><td>
<tr><td>&#x1F1F8;&#x1F1FD;<td><td>Sint Maarten<td><code>&amp;#x1F1F8; &amp;#x1F1FD;</code><td><code>&amp;#127480; &amp;#127485;</code><td>
<tr><td>&#x1F1F8;&#x1F1FE;<td><td>Syria<td><code>&amp;#x1F1F8; &amp;#x1F1FE;</code><td><code>&amp;#127480; &amp;#127486;</code><td>
<tr><td>&#x1F1F8;&#x1F1FF;<td><td>Swaziland<td><code>&amp;#x1F1F8; &amp;#x1F1FF;</code><td><code>&amp;#127480; &amp;#127487;</code><td>
<tr><td>&#x1F1F9;&#x1F1E6;<td><td>Tristan da Cunha<td><code>&amp;#x1F1F9; &amp;#x1F1E6;</code><td><code>&amp;#127481; &amp;#127462;</code><td>
<tr><td>&#x1F1F9;&#x1F1E8;<td><td>Turks &amp; Caicos Islands<td><code>&amp;#x1F1F9; &amp;#x1F1E8;</code><td><code>&amp;#127481; &amp;#127464;</code><td>
<tr><td>&#x1F1F9;&#x1F1E9;<td><td>Chad<td><code>&amp;#x1F1F9; &amp;#x1F1E9;</code><td><code>&amp;#127481; &amp;#127465;</code><td>
<tr><td>&#x1F1F9;&#x1F1EB;<td><td>French Southern Territories<td><code>&amp;#x1F1F9; &amp;#x1F1EB;</code><td><code>&amp;#127481; &amp;#127467;</code><td>
<tr><td>&#x1F1F9;&#x1F1EC;<td><td>Togo<td><code>&amp;#x1F1F9; &amp;#x1F1EC;</code><td><code>&amp;#127481; &amp;#127468;</code><td>
<tr><td>&#x1F1F9;&#x1F1ED;<td><td>Thailand<td><code>&amp;#x1F1F9; &amp;#x1F1ED;</code><td><code>&amp;#127481; &amp;#127469;</code><td>
<tr><td>&#x1F1F9;&#x1F1EF;<td><td>Tajikistan<td><code>&amp;#x1F1F9; &amp;#x1F1EF;</code><td><code>&amp;#127481; &amp;#127471;</code><td>
<tr><td>&#x1F1F9;&#x1F1F0;<td><td>Tokelau<td><code>&amp;#x1F1F9; &amp;#x1F1F0;</code><td><code>&amp;#127481; &amp;#127472;</code><td>
<tr><td>&#x1F1F9;&#x1F1F1;<td><td>Timor-Leste<td><code>&amp;#x1F1F9; &amp;#x1F1F1;</code><td><code>&amp;#127481; &amp;#127473;</code><td>
<tr><td>&#x1F1F9;&#x1F1F2;<td><td>Turkmenistan<td><code>&amp;#x1F1F9; &amp;#x1F1F2;</code><td><code>&amp;#127481; &amp;#127474;</code><td>
<tr><td>&#x1F1F9;&#x1F1F3;<td><td>Tunisia<td><code>&amp;#x1F1F9; &amp;#x1F1F3;</code><td><code>&amp;#127481; &amp;#127475;</code><td>
<tr><td>&#x1F1F9;&#x1F1F4;<td><td>Tonga<td><code>&amp;#x1F1F9; &amp;#x1F1F4;</code><td><code>&amp;#127481; &amp;#127476;</code><td>
<tr><td>&#x1F1F9;&#x1F1F7;<td><td>Turkey<td><code>&amp;#x1F1F9; &amp;#x1F1F7;</code><td><code>&amp;#127481; &amp;#127479;</code><td>
<tr><td>&#x1F1F9;&#x1F1F9;<td><td>Trinidad &amp; Tobago<td><code>&amp;#x1F1F9; &amp;#x1F1F9;</code><td><code>&amp;#127481; &amp;#127481;</code><td>
<tr><td>&#x1F1F9;&#x1F1FB;<td><td>Tuvalu<td><code>&amp;#x1F1F9; &amp;#x1F1FB;</code><td><code>&amp;#127481; &amp;#127483;</code><td>
<tr><td>&#x1F1F9;&#x1F1FC;<td><td>Taiwan<td><code>&amp;#x1F1F9; &amp;#x1F1FC;</code><td><code>&amp;#127481; &amp;#127484;</code><td>
<tr><td>&#x1F1F9;&#x1F1FF;<td><td>Tanzania<td><code>&amp;#x1F1F9; &amp;#x1F1FF;</code><td><code>&amp;#127481; &amp;#127487;</code><td>
<tr><td>&#x1F1FA;&#x1F1E6;<td><td>Ukraine<td><code>&amp;#x1F1FA; &amp;#x1F1E6;</code><td><code>&amp;#127482; &amp;#127462;</code><td>
<tr><td>&#x1F1FA;&#x1F1EC;<td><td>Uganda<td><code>&amp;#x1F1FA; &amp;#x1F1EC;</code><td><code>&amp;#127482; &amp;#127468;</code><td>
<tr><td>&#x1F1FA;&#x1F1F2;<td><td>U.S. Outlying Islands<td><code>&amp;#x1F1FA; &amp;#x1F1F2;</code><td><code>&amp;#127482; &amp;#127474;</code><td>
<tr><td>&#x1F1FA;&#x1F1F8;<td><td>United States<td><code>&amp;#x1F1FA; &amp;#x1F1F8;</code><td><code>&amp;#127482; &amp;#127480;</code><td>
<tr><td>&#x1F1FA;&#x1F1FE;<td><td>Uruguay<td><code>&amp;#x1F1FA; &amp;#x1F1FE;</code><td><code>&amp;#127482; &amp;#127486;</code><td>
<tr><td>&#x1F1FA;&#x1F1FF;<td><td>Uzbekistan<td><code>&amp;#x1F1FA; &amp;#x1F1FF;</code><td><code>&amp;#127482; &amp;#127487;</code><td>
<tr><td>&#x1F1FB;&#x1F1E6;<td><td>Vatican City<td><code>&amp;#x1F1FB; &amp;#x1F1E6;</code><td><code>&amp;#127483; &amp;#127462;</code><td>
<tr><td>&#x1F1FB;&#x1F1E8;<td><td>St. Vincent &amp; Grenadines<td><code>&amp;#x1F1FB; &amp;#x1F1E8;</code><td><code>&amp;#127483; &amp;#127464;</code><td>
<tr><td>&#x1F1FB;&#x1F1EA;<td><td>Venezuela<td><code>&amp;#x1F1FB; &amp;#x1F1EA;</code><td><code>&amp;#127483; &amp;#127466;</code><td>
<tr><td>&#x1F1FB;&#x1F1EC;<td><td>British Virgin Islands<td><code>&amp;#x1F1FB; &amp;#x1F1EC;</code><td><code>&amp;#127483; &amp;#127468;</code><td>
<tr><td>&#x1F1FB;&#x1F1EE;<td><td>U.S. Virgin Islands<td><code>&amp;#x1F1FB; &amp;#x1F1EE;</code><td><code>&amp;#127483; &amp;#127470;</code><td>
<tr><td>&#x1F1FB;&#x1F1F3;<td><td>Vietnam<td><code>&amp;#x1F1FB; &amp;#x1F1F3;</code><td><code>&amp;#127483; &amp;#127475;</code><td>
<tr><td>&#x1F1FB;&#x1F1FA;<td><td>Vanuatu<td><code>&amp;#x1F1FB; &amp;#x1F1FA;</code><td><code>&amp;#127483; &amp;#127482;</code><td>
<tr><td>&#x1F1FC;&#x1F1EB;<td><td>Wallis &amp; Futuna<td><code>&amp;#x1F1FC; &amp;#x1F1EB;</code><td><code>&amp;#127484; &amp;#127467;</code><td>
<tr><td>&#x1F1FC;&#x1F1F8;<td><td>Samoa<td><code>&amp;#x1F1FC; &amp;#x1F1F8;</code><td><code>&amp;#127484; &amp;#127480;</code><td>
<tr><td>&#x1F1FD;&#x1F1F0;<td><td>Kosovo<td><code>&amp;#x1F1FD; &amp;#x1F1F0;</code><td><code>&amp;#127485; &amp;#127472;</code><td>
<tr><td>&#x1F1FE;&#x1F1EA;<td><td>Yemen<td><code>&amp;#x1F1FE; &amp;#x1F1EA;</code><td><code>&amp;#127486; &amp;#127466;</code><td>
<tr><td>&#x1F1FE;&#x1F1F9;<td><td>Mayotte<td><code>&amp;#x1F1FE; &amp;#x1F1F9;</code><td><code>&amp;#127486; &amp;#127481;</code><td>
<tr><td>&#x1F1FF;&#x1F1E6;<td><td>South Africa<td><code>&amp;#x1F1FF; &amp;#x1F1E6;</code><td><code>&amp;#127487; &amp;#127462;</code><td>
<tr><td>&#x1F1FF;&#x1F1F2;<td><td>Zambia<td><code>&amp;#x1F1FF; &amp;#x1F1F2;</code><td><code>&amp;#127487; &amp;#127474;</code><td>
<tr><td>&#x1F1FF;&#x1F1FC;<td><td>Zimbabwe<td><code>&amp;#x1F1FF; &amp;#x1F1FC;</code><td><code>&amp;#127487; &amp;#127484;</code><td>`
