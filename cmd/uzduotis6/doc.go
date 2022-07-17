/* 6. Kreksint nuo vieno tvenkinio iki kito
Ši patirtis ančiukams buvo išties nervinga, tačiau jie jau išėjo iš interviu ir iš pastato. Laikas informuoti šeimą apie tai, kaip sekėsi. Ne, antys nenaudoja SMS. Jie turi savo sistemą – QQS. Jos šifras grindžiamas supaprastinta labai populiarios kodavimo mašinos versija.

Šiame kode naudojami 3 lietuviškos abėcėlės masyvai [A,Ą,B,C,Č,D,E,Ę,Ė,F,G,H,I,Į,Y,J,K,L,M,N,O,P,R,S,Š,T,U,Ų,Ū,V,Z,Ž], kur X apibrėžia, kuri raidė yra pirma masyve (t.y. jei X = 1,  tada pirmoji raidė yra A, jei X = 5, tada pirmoji raidė yra Č ir t.t.).

Raidė užkoduojama paimant raidės vietą Y lietuviškoje abėcėlėje (A=1, B=3, D=6 ir t.t.) ir pažiūrint kokia raidė yra vietoje Y pastumtame masyve (t.y., jeigu masyvo X yra 2, o Y yra 1, tai A pavirstų Ą, B pavirstų C ir t.t.).

Raidė taip pat užkoduota per visus 3 masyvus, po kurių pirmasis masyvas perkeliamas viena pozicija, o antrasis masyvas perkeliamas tik tada, kai pirmasis masyvas visiškai apsisuka, o trečiasis – tik tada, kai antrasis visiškai apsisuka.

Pavyzdžiui, jei X1 = 5, X2 = 6, X3 = 9, masyvai atrodys taip:

[Č,D,E,Ę,Ė,F,G,H,I,Į,Y,J,K,L,M,N,O,P,R,S,Š,T,U,Ų,Ū,V,Z,Ž,A,Ą,B,C]

[D,E,Ę,Ė,F,G,H,I,Į,Y,J,K,L,M,N,O,P,R,S,Š,T,U,Ų,Ū, V,Z,Ž,A,Ą,B,C,Č]

[Ė,F,G,H,I,Į,Y,J,K,L,M,N,O,P,R,S,Š,T,U,Ų,Ū,V,Z,Ž,A,Ą,B,C,Č,D,E,Ę]

Antys išsiuntė žinią savo šeimai, tačiau pamiršo, kokios pradinės masyvų pozicijos buvo naudojamos koduojant tekstą. Vienintelis dalykas, kurį jie žino, yra tai, kad pranešimas prasideda TELIA. Žemiau – ančių žinutė.

Kokios buvo pradinės pozicijos, naudojamos tekstui koduoti? Atsakymų langelyje įveskite tris skaičius, atskirdami tarpu.

ĄJŲSIVBOČĄCĮRPĮYPŽEHĖKZIVĘYNZŪAEĮČŽDUUDĘĄMIAĖRKIGEŲRIŠČĖUYKTŲMĮŠR

*/

package main
