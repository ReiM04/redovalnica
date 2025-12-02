Peket redovalnica omogoča izdelavo redovalnice študentov. Paket omogoča dodajanje ocen, izpis redovalnice ter izpis končnih ocen. Primer uporabe: 

studenti := map[string]redovalnica.Student{
		"63210001" : {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8}},
		"63210002" : {Ime: "Boris", Priimek: "Kralj", Ocene:[]int{6,7,5,8}},
		"63210003" : {Ime: "Janez", Priimek: "Novaj", Ocene:[]int{4,5,3,5}},
}

redovalnica.DodajOceno(studenti, "63210002", 10, minOcena, maxOcena)
redovalnica.IzpisiKoncniUspeh(studenti, stOcen)
redovalnica.IzpisVsehOcen(studenti)
