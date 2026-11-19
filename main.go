package main

import "fmt"

func main() {
	// Einfaches Fenster
	kipp := NewWindow("Kippfenster",
		Frame{"Holzrahmen", 120},
		Glass{"Doppelverglasung", 80},
	)

	// Fenster mit verschachtelten Fensterflügeln
	balkon := NewWindow("Balkontür",
		Frame{"Kunststoffrahmen", 150},
		NewWindow("Linker Flügel", Frame{"Flügelrahmen", 60}, Glass{"Dreifachverglasung", 110}),
		NewWindow("Rechter Flügel", Frame{"Flügelrahmen", 60}, Glass{"Sicherheitsglas", 140}),
	)

	// Dachfenster, nachträglich ergänzt
	dach := NewWindow("Dachfenster", Frame{"Aluminiumrahmen", 200})
	dach.Add(Glass{"Wärmeschutzglas", 170})

	haus := NewWindow("Haus (alle Fenster)", kipp, balkon, dach)

	for _, w := range []Item{kipp, balkon, dach, haus} {
		fmt.Print(w.Display())
		fmt.Printf("=> Gesamtpreis: %.2f €\n\n", w.Price())
	}
}
