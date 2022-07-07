/*
Copyright (c) 2022 İlteriş Yağıztegin Eroğlu (linuxgemini)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package utils

var (
	illerAlfabetik map[string]int = map[string]int{
		"ADANA":          1,
		"ADIYAMAN":       2,
		"AFYONKARAHİSAR": 3,
		"AĞRI":           4,
		"AKSARAY":        68,
		"AMASYA":         5,
		"ANKARA":         6,
		"ANTALYA":        7,
		"ARDAHAN":        75,
		"ARTVİN":         8,
		"AYDIN":          9,
		"BALIKESİR":      10,
		"BARTIN":         74,
		"BATMAN":         72,
		"BAYBURT":        69,
		"BİLECİK":        11,
		"BİNGÖL":         12,
		"BİTLİS":         13,
		"BOLU":           14,
		"BURDUR":         15,
		"BURSA":          16,
		"ÇANAKKALE":      17,
		"ÇANKIRI":        18,
		"ÇORUM":          19,
		"DENİZLİ":        20,
		"DİYARBAKIR":     21,
		"DÜZCE":          81,
		"EDİRNE":         22,
		"ELAZIĞ":         23,
		"ERZİNCAN":       24,
		"ERZURUM":        25,
		"ESKİŞEHİR":      26,
		"GAZİANTEP":      27,
		"GİRESUN":        28,
		"GÜMÜŞHANE":      29,
		"HAKKARİ":        30,
		"HATAY":          31,
		"IĞDIR":          76,
		"ISPARTA":        32,
		"İSTANBUL":       34,
		"İZMİR":          35,
		"KAHRAMANMARAŞ":  46,
		"KARABüK":        78,
		"KARAMAN":        70,
		"KARS":           36,
		"KASTAMONU":      37,
		"KAYSERİ":        38,
		"KİLİS":          79,
		"KIRIKKALE":      71,
		"KIRKLARELİ":     39,
		"KIRŞEHİR":       40,
		"KOCAELİ":        41,
		"KONYA":          42,
		"KÜTAHYA":        43,
		"MALATYA":        44,
		"MANİSA":         45,
		"MARDİN":         47,
		"MERSİN":         33,
		"MUĞLA":          48,
		"MUŞ":            49,
		"NEVŞEHİR":       50,
		"NİĞDE":          51,
		"ORDU":           52,
		"OSMANİYE":       80,
		"RİZE":           53,
		"SAKARYA":        54,
		"SAMSUN":         55,
		"ŞANLIURFA":      63,
		"SİİRT":          56,
		"SİNOP":          57,
		"ŞIRNAK":         73,
		"SİVAS":          58,
		"TEKİRDAĞ":       59,
		"TOKAT":          60,
		"TRABZON":        61,
		"TUNCELİ":        62,
		"UŞAK":           64,
		"VAN":            65,
		"YALOVA":         77,
		"YOZGAT":         66,
		"ZONGULDAK":      67,
	}
)

func GetIlListesi() map[string]int {
	return illerAlfabetik
}
