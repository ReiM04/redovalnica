package main

import (
    "context"
	"log"
    "os"
    "github.com/urfave/cli/v3"
    "github.com/ReiM04/redovalnica/pkg"
)

func main() {
	cmd := &cli.Command{
		Name: "redovalnica",
		Usage: "redovalnica za dodajanje in izpis ocen",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "stOcen",
				Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name: "minOcena",
				Usage: "najmanjša možna ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name: "maxOcena",
				Usage: "največja možna ocena",
				Value: 10,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			return run(stOcen, minOcena, maxOcena)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(stOcen int, minOcena int, maxOcena int) error {

	studenti := map[string]redovalnica.Student{
		"63210001" : {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8}},
		"63210002" : {Ime: "Boris", Priimek: "Kralj", Ocene:[]int{6,7,5,8}},
		"63210003" : {Ime: "Janez", Priimek: "Novaj", Ocene:[]int{4,5,3,5}},
	}

	redovalnica.DodajOceno(studenti, "63210002", 10, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210002", 9, minOcena, maxOcena)

	redovalnica.DodajOceno(studenti, "63210002", 42, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "12345678", 10, minOcena, maxOcena)

	redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

	redovalnica.DodajOceno(studenti, "63210001", 9, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210001", 9, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210001", 9, minOcena, maxOcena)

	redovalnica.IzpisVsehOcen(studenti)

	redovalnica.DodajOceno(studenti, "63210003", 4, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210003", 4, minOcena, maxOcena)
	redovalnica.DodajOceno(studenti, "63210003", 4, minOcena, maxOcena)

	redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

	return nil

}


