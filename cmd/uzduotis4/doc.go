/*  4. Kai antys išskrenda žiemoti į pietus

Laikas interviu. Kodis eina pirmas. Atrankų specialistas atrodo pakankamai draugiškas, o jį lydintis programuotojas neatrodo panašus į tuos hakerius, matomus interneto memuose. Neįsitempkime, atsipalaiduokime ir tiesiog būkime savimi.

Interviu prasideda pradeda nuo pamatinių klausimų, o vėliau kalba pereina prie Kodžio asmenybės ir patirties. Paklaustas apie savo požiūrį į darbą iš užsienio, jis tiesiog negali tylėti. Kodis turi atskleisti, kad atėjus žiemai jis dirbs iš pietų. Fiziškai tiesiog nėra jokio būdo, kad jis galėtų likti čia, šaltyje.

Be didelių dvejonių, programuotojas pasinaudoja proga užduoti svarbiausią klausimą – kaip antys suderina savo V formos išsidėstymą skrisdamos į pietus? Kodis netgi yra tam klausimui parašę matricą, bet kaip tik tada jo galvoje spengianti tyle. Ar galite jam padėti?

Jis žino, kad yra keli žingsnių, kurių reikia norint pasiekti tobulą formą. Raskite žingsnius, būtinuss norint surinkti raides TGSL trejose raidžių matricose (5×5). Pradedant vieta (0;0), raskite trumpiausią kelią surinkti reikiamas raides (T, G, S, L) bet kokia tvarka (matricose negalima judėti įstrižai). Atsakymų lauke įveskite visų trumpiausių maršrutų sandaugą iš visų trijų matricų.

Pavyzdys:
Z	K	T	M	G
A	B	C	D	E
Y	S	R	Q	X
L	J	H	P	O
A	B	Y	J	P

Trumpiausias kelias susidėtų iš 10 žingnsių (L -> S -> T -> G)

Koordinatės: (0;0) 3 žingsniai -> L(0;3) 2 žingsniai -> S(1;2) 3 žingsniai -> T(2;0) 2 žingsniai -> G(4;0) = 3*2*3*2 = 36 žingsniai iš viso.
A 	B 	C 	D 	S
F 	Z 	H 	I 	J
K 	L 	M 	N 	O
P 	Q 	T 	E 	R
U 	V 	W 	X 	G

A 	Z 	C 	D 	O
F 	G 	H 	J 	I
K 	V 	M 	N 	T
P 	Q 	R 	E 	S
U 	L 	W 	X 	B

A 	B 	C 	D 	S
F 	Z 	H 	I 	J
K 	P 	M 	N 	O
L 	G 	R 	E 	T
U 	V 	W 	X 	Q

*/

package main
