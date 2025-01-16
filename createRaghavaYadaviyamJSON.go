package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func createRaghavaYadaviyamJSON(text string) map[string]map[string]string {
	textParts := strings.Split(text, "||")
	result := make(map[string]map[string]string)

	for i := 1; i <= 30; i++ {
		part := strings.TrimSpace(textParts[i-1])
		result[fmt.Sprintf("%d", i)] = map[string]string{
			"anulom":   part,
			"pratilom": reverseString(part),
		}
	}

	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	text := `vande ahaM devaM taM shrItaM rantAraM kAlaM bhAsA ya:|
rAma: rAmAdhI: ApyAga:leelAm Ara Ayodhye vAse||
sevAdhyeyo rAmAlAlee gopyArAdhI mArAmorA:|
yassAbhAla~NkAraM tAraM taM shrItaM vande ahaM devam||
sAketAkhyA jyAyAmAseet yA viprAdIptA AryAdhArA|
pU: AjIta adevAdyAvishvAsA agryA sAvAshArAvA||
ayodyA mathurA mayA kAshIM kA~nciravantikA
pUrI dvAravaTI caiva saptaidAH mokShadAyikAH||
vArAshAvAsAgryA sAshvAvidyAvAdetAjeerA pU:|
rAdhAryAptA dIprA vidyAsImA yA jyAkhyAtA ke sA||
kAmabhArassthalasArashreesaudhA asau ghanvApikA|
sArasAravapeenA sarAgAkArasubhUribhU:||
bhUribhUsurakAgArAsanA peevarasArasA|
kA api va anaghasaudha asau shreerasAlasthabhAmakA||
rAmadhAma samAnenam Agorodhanam Asa tAm||
nAmahAm akShararasaM tArAbhA: tu na veda yA||
yAdavenaH tu bhArAtA saMrarakSha mahAmanAH|
tAM saH mAnadharaH gomAn anemAsamadhAmarAH||
yan gAdheyaH yogee rAgee vaitAne saumye saukhye asau|
taM khyAtaM sheetaM spheetaM bheemAn Ama ashreehAtA trAtam||
taM trAtA hA shreemAn Ama abheetaM spheetaM sheetaM khyAtaM|
saukhye saumye asau netA vai geerAgee yaH yodhe gAyan||
mAramaM sukumArAbhaM rasAja Apa nRRitAshritaM|
kAvirAmadalApa gosama avAmatarA nate||
tena rAtam avAma asa gopalAt amarAvika|
taM shrita nRRipajA sArabhaM rAmA kusumaM ramA||
rAmanAmA sadA khedabhAve dayAvAn atApeenatejAH ripau Anate|
kAdimodAsahAtA svabhAsA rasAme sugaH reNukAgAtraje bhUrume||
merubhUjetragA kANure gosume sA arasA bhAsvatA hA sadA modikA|
tena vA pArijAtena peetA navA yAdave abhAt akhedA samAnAmara||
sArasAsamadhAta akShibhUmnA dhAmasu sItayA|
sAdhu asau iha reme kSheme aram AsurasArahA ||
hArasArasumA ramyakShemera iha visAdhvasA|
ya atasIsumadhAmnA bhUkShitA dhAma sasAra sA||
sAgasA bharatAya ibhamAbhAtA manyumattayA |
sa atra madhyamaya tApe potAya adhigatA rasA ||
sAratAgadhiyA tApopetA yA madhyamatrasA|
yAttamanyumataa bhAmA bhayetA rabhasAgasA||
tAnavAt apakA umAbhA rAme kAnanada Asa sA|
yA latA avRRiddhasevAkA kaikeyI mahada ahaha||
haha dAhamayee kekaikAvAseddhavRRitAlayA |
sA sadAnanakA AmerA bhAmA kopadavAnatA ||
varamAnadasatyAsahreetapitrAdarAt aho|
bhAsvara: sthiradheera: apahArorA: vanagAmee asau||
somyagAnavarArohApara: dheera: sthirasvabhAvA:|
ho darAt atra Apitahree satyAsadanam Ara vA ||
yA nayAnaghadheetAdA rasAyA: tanayA dave|
sA gatA hi viyAtA hreesatApA na kila UnAbhA||
bhAn aloki na pAtA sa: hreetA yA vihitAgasA |
vedayAna: tayA sAradAta dheeghanayA anayA ||
rAgirAdhutigarvAdAradAha: mahasA haha |
yAn agAta bharadvAjam AyAsee damagAhina: ||
no hi gAm adaseeyAmAjat va Arabhata gA: na yA |
haha sA Aha mahodAradArvAgatidhurA girA ||
yAturAjidabhAbhAraM dyAM va mArutagandhagam |
sa: agam Ara padaM yakShatu~NgAbha: anaghayAtrayA ||
yAtrayA ghanabha: gAtuM kShayadaM paramAgasa: |
gandhagam tarum Ava dyAM rambhAbhAdajirA tu yA ||
daNDakAM pradama: rAjAlya: hatAmayakArihA |
sa: samAnavatAnenobhogyAbha: na tadA Asa na ||
na sadAtanabhogyAbha: no netA vanam Asa sa: |
hArikAyamatAhalyAjArAmodaprakADajam ||
sa: aram Arat anaj~nAna: vederAkaNThakumbhajam |
taM drusArapaTa: anAgA: nAnAdoSavirAdhahA ||
hA dharAviSada: nAnAgAnAToparasAt drutam |
jambhakuNThakarA: devena: aj~nAnadaram Ara saH ||
sAgamAkarapAtA hAka~NkenAvanata: hi sa: |
na samAnarda mA arAmA la~NkArAjasvasA ratam ||
taM rasAsu ajarAkAlaM ma ArAmardanam Asa na |
sa: hita: anavanAkekaM hAtA apArakaM AgasA ||
tAM sa: goramadoshreeda: vigrAm asadara: atata |
vairam Asa palAhArA vinAsA ravivaMshake ||
keshavaM virasAnAviH Aha AlApasamAravai: |
tatarodasam agrAvida: ashreeda: amaraga: asatAm ||
godyugoma: svamAya: abhUt ashreegakharasenayA |
saha sAhavadhAra: avikala: arAjat arAtihA ||
hA atirAdajarAloka virodhAvahasAhasa |
yAnaserakhaga shreeda bhUya: ma svam aga: dyuga: ||
hatapApacaye heya: la~Nkesha: ayam asAradhee: |
rajirAviraterApa: hA hA aham graham Ara gha: ||
ghoram Aha grahaM hAhApa: arAteH ravirAjirA: |
dheerasAmayashoke alaM ya: heye ca papAta ha: ||
tATakeyalavAt enohAree hArigira Asa sa: |
hA asahAyajanA siitA anAptenA adamanA: bhuvi ||
vibhunA madanAptena Ata AseenAjayahAsahA |
sa: sarA: girihAree ha no devAlayake aTatA ||
bhAramA kudashAkena AsharAdheekuhakena hA |
cArudheevanapAlokyA vaidehee mahitA hRRitA ||
tA: hRRitA: hi maheedevaikyAlopAnavadheerucA |
hAnakehakudheerAshanAkeshA adakumArabhA: ||
hAritoyadabha: rAmAviyoge anaghavAyuja: |
taM rumAmahita: apetAmodA: asAraj~na: Ama ya: ||
ya: amarAj~na: asAdoma: atApeta: himamArutam |
ja: yuvA ghanageya: vim Ara Abhodayata: arihA ||
bhAnubhAnutabhA: vAmA sadAmodapara: hataM |
taM ha tAmarasAbhAkSha: atirAtA akRRita vAsavim ||
viM sa: vAtakRRitArAtikShobhAsAramatAhataM |
taM haropadama: dAsam Ava AbhatanubhanubhA: ||
haMsajAruddhabalajA parodArasubhA ajani |
rAji rAvaNa rakShoravighAtAya ramA Ara yam ||
yaM ramaa Ara yatAgha virakShoraNavarAjira |
nijabhA surada ropajAlabaddha rujAsaham ||
sAgaratigam AbhAtinAkesha: asuramAsaha: |
taM sa: mArutajam goptA abhAt AsAdya gata: agajam ||
aM gata: gadee asAdAbhAptA gojaM tarum Asa tam |
ha: samArasushokena atibhAmAgati: Agasa ||
veeravAnarasenasya trAta abhAt avatA hi sa: |
toyadhou arigoyAdasi ayata: navasetunA ||
nA tu sevanata: yasya dayAga: arivadhAyata: |
sa hi tAvat abhata trAsee anase: anavAravii ||
hArisAhasala~NkenAsubhedee mahita: hi sa: |
cArubhUtanuja: rAma: aram ArAdhayadArtihA ||
ha ArtidAya dharAm Ara morA: ja: nutabhU: rucA |
sa: hita: hi madeebhe sunAke alaM sahasA arihA ||
nAlikera subhAkArAgArA asau surasApikA |
rAvaNArikShamerA pU: Abheje hi na na amunA ||
na amunA nahi jebhera pU: Ame akShariNA varA |
kA api sArasusaurAgA rAkAbhAsurakelinA ||
sA agryatAmarasAgArAm akShAmA ghanabhA Ara gau: |
nijade aparajiti Asa shree: rAme sugarAjabhA ||
bhA ajaraga sumerA shreesatyAjirapade ajani |
gaurabhA anaghamA kShAmarAgA sa aramata agryasA ||
`

	result := createRaghavaYadaviyamJSON(text)
	jsonOutput, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonOutput))
	saveJSONToFile(jsonOutput, "raghavaYadaviyam.json")
}

// Function to save this jsonOutput to a file in the same directory
func saveJSONToFile(jsonOutput []byte, filename string) error {
	return os.WriteFile(filename, jsonOutput, 0644)
}
func saveToFile(jsonOutput []byte, filename string) error {
	return os.WriteFile(filename, jsonOutput, 0644)
}
