# is105-uke04

Navnliste: 
Abdikani Gureye, 
Brede Knutsen Meli, 
Eirik Aanestad Fintland, 
Jan Kevin Henriksen, 
Mats Skjærvik, 
Mikael Kimerud, 
Morten Johan Mygland, 
Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor kan det vise at det er bare 5/8 som har lastet opp noe. 

Oppgave 1.2.1
Titall til totall:
1 = 0001
2 = 0010
5 = 0101
8 = 1000
16 = 0001 0000
256 = 0001 0000 0000

Verdier på de ulike plassene:                 Eks. omregning fra ti- til totallsystemet.
2^0=1                                                          1 = (2^0 x 1) = 1
2^1=2                                                          2 = (2^1 x 1) + (2^0 x 0) = 10
2^2=4                                                          5 = (2^2 x 1) + (2^1 x 0) + (2^0 x 1) = 101
2^3=8                                                          osv.
2^4=16      
2^5=32
2^6=64
2^7=128
2^8=256
2^9=512

Totall til titall: 
100 = 1*2^2 + 0*2^1 + 0*2^0 = 4 + 0 + 0 = 4
1001 = 1*2^3 + 0*2^2 + 0*2^1 + 1*2^0 = 8 + 0 + 0 +1 = 9
1100110011 = 512 + 256 + 0x128 + 0x64 + 32 + 16 + 0x8 + 0x4 + 2 + 1 = 819 
^10 sifre, derfor vet man det må hvertfall være 2^9

Når man skal konvertere fra totallsystemet til titallsystemet, så bruker man det faktum at det kun brukes to sifre i totallsystemet, 0 og 1. 0 og 1 representerer da ulike verdier, basert på hvilken posisjon de har. Ut ifra dette bruker man grunntallet 2, og opphøyer dette i sifferets posisjon nummer og multipliserer dette med sifferverdien.  10 i totallsystemet blir da f.eks: (1 * 2^1) + (0 * 2^0) = 2 + 0 = 2 i titallsystemet. 

Fra titallsystemet til totallsystemet bruker man det samme prinsippet med at alle sifrene har posisjoner med forskjellige verdier, bare motsatt vei. Og da må man regne ut slik at de ulike verdiene til sammen blir 10. Over har vi laget en liten tabell som skriver hvilke verdier de ulike posisjonene i totallsystemet har. Ut ifra den kan man se at man trenger 1 på 2^3 plassen (for det blir 8, ”åtterplass”) og 1 i 2^1 plassen (det blir 2, ”toerplass”). De andre plassene vil da få 0. Altså blir 10 i titallsystemet da f.eks: 1010 i totallsystemet. (1 * 2^3) + (0 * 2^2) + (1 * 2^1) + (0 * 2^0) = 8 + 0 + 2 + 0 = 10. 

Oppgave 1.2.2
3-bit = 8 muligheter (tall, 0-7). (000, 001, 010, 011, 100, 101, 110, 111).

(1) 4 Oddetall (1, 3, 5 og 7)
Vi har en formel som sier noe om hvordan man skal regne informasjonsmengden basert på antall sannsynlige valgmuligheter, og også hvis man får mer informasjon om et valg som begrenser disse valgmulighetene igjen. Formelen er log2(1 / (M/N) ), der N er antall valgmuligheter (8), og M antall informasjon som begrenser valgmuligheter(4) N. Ved bruk av denne formelen vil svaret bli:  log2(1/(4/8)) = log2(2) = 1 bit.

(2) Ikke multiplum av 3 (0, 3, 6). Da gjenstår 1, 2, 4, 5, 7.
Log2(1/(5/8)) = 0.678072 bit. 

(3) 011, 101, 110
Log2(1/(3/8)) =  ~1.4 bit.

(4) Ut i fra opplysningene i oppgave 1,2 og 3, så vet vi at det eneste tallet som passer alle beskrivelsene er  5. Da blir formelen:
Log2(1/(1/8)) = 3 bit. Hun har altså all den informasjonen hun trenger for å vite tallet.

Jo mer informasjon (bit) en person har, jo større sannsynlighet har de for å gjette riktig tall.


1.2.3
Begynte med å opprette et repository fra GitHub og klonet dette ved bruk av
git clone <URL>. Så laget vi filen "Hello.go" og brukte git add Hello.go, git status for å se om filen ble lagt til i staging, git commit -m "melding" og git push origin master for å opplaste den til GitHub.



1.2.4

(1) Hvilken fordeler og ulemper har en git-flow-modell med en hovedrepository?

Link til hovedrepository: https://github.com/Bredema/is105-uke04 
Fordeler: Parallell jobbing fungerer godt, enkelte kan gjøre endringer, så kan andre 
i gruppen sjekke om jobber er gjort bra. Det er også lettere å gå tilbake å gjør endringer
hvis det er noe som ikke fungerer. Samtlige har også mulighet til å kun gjøre endringer i sin egen branch, slik at man unngår å overskrive andres arbeid.
Ulemper: Kanskje hvis man ikke er flink til å update/merge branchen sin ofte, så kan det bli
mye og uoversiktlig.


(2) Finn ut hva heter objektfiler for de mest brukte platformer (Unix/Linux, MS Windows, Mac OS X)! Hvorfor, etter deres mening, har disse platformene så forskjellige objektfil-formater?

Gjorde mye av det samme som i 1.2.3 når vi først hadde fått tilgang til hovedrepositoren,
bare at vi gjorde endringer i egen branch (som ble laget ved git checkout -b <branchnavn>), og sendte enn pull request og merget med master. (git merge <branchnavn>).

Oversikt over objektfiler for de forskjellige plattformene:
Windows: DLL, EXE
Linux: ELF (buildavhengig)
Mac:EXEC.

De forskjellige plattformene har ulike objekt formater fordi de har ulike måter de leser og kaller på binære tallkoder.


(3) Hvilke forskjeller ser dere i forhold til programmeringsspråket Java?

En av de største forskjellene mellom Java og Golang er at Golang leser koden en linje av gangen, og kan dermed kjøre store deler av programmet før den crasher, mens java krever at hele koden er feilfri.
Golang er heller ikke objektorientert, noe som java er.


(4) 
(Master-> log)
Hvilke viktige poeng illustrerer denne øvelsen når det gjelder bruk av et programmeringsmiljø på en platform?

Programmering på en platform har mange fordeler, der i blant versjonskontroll, effektivtet og åpenhet. Vi fant også ut at mappestruktur er veldig viktig for å kunne importere en selvlaget pakke, i dette tilfellet “package log”, som ble importert i main.go via “import ./log”. 


(5)
(Master -> log -> logbcli.exe) 
Er det hensiktsmessig å legge inn denne filen i git repository? Begrunn svaret!

Nei! Den eneste hensikten med å ha et utførbart program i Repositoryen er for å dele den med resten av gruppen. MEN, siden alle har tilgang til .go filene vil alle kunne builde sine egne programmer, så mye av hensikten faller egentlig bort.
i tilleg så vil man måtte “builde” forskjellige filer for forskjellige operativsystemer.

(6)
(Master -> log -> logbcli.exe)
(Kildekode ligger som log.go og logbcli.go) 
Hvordan skiller pakken log​, som dere har implementert, seg fra andre pakker i go, som, for eksempel, fmt​? 


Log er en lokal pakke som vi har skrevet selv, og vi har full kontroll over funksjonaliteten til denne pakken. fmt er en del av standardbiblioteket for golang, og vi kan kun bruke funksjonene som definert i fmt dersom vi importerer fmt. 
