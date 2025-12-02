package redovalnica

import "fmt"


type Student struct {
	Ime 	string
	Priimek string
	Ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) {
	if ocena < minOcena || ocena > maxOcena {
		fmt.Println("Ocena ni v med", minOcena , " in " maxOcena)
		return
	}

	e, ok := studenti[vpisnaStevilka]
	if ok {
		e.Ocene = append(e.Ocene, ocena)
		studenti[vpisnaStevilka] = e
	} else {
		fmt.Println("Studenta ni v seznamu")
	}
	return
}

func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
	e, ok := studenti[vpisnaStevilka]
	if !ok {return -1}
	if len(e.Ocene) < stOcen {return 0.0}
	var povprecje int
	for _, v := range e.Ocene {
		povprecje = povprecje + v
	}

	return (float64(povprecje) / float64(len(e.Ocene)))
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for i, v := range studenti {
		fmt.Printf("%s - %s %s: %v\n", i, v.Ime,v.Priimek, v.Ocene)
	}
	return
}

func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int) {
	for i, v := range studenti {
		var ocena = povprecje(studenti, i, stOcen)

		if ocena >= 9 {
			fmt.Printf("%s %s: povprecna ocena %.1f -> Odličen študent!\n", v.Ime, v.Priimek, ocena)
		} else if ocena > 6 {
			fmt.Printf("%s %s: povprecna ocena %.1f -> Povprečen študent!\n", v.Ime, v.Priimek, ocena)
		} else {
			fmt.Printf("%s %s: povprecna ocena %.1f -> Neuspešen študent!\n", v.Ime, v.Priimek, ocena)
		}
	}
	return
}

