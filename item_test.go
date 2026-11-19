package main

import (
	"strings"
	"testing"
)

func TestLeafPrices(t *testing.T) {
	if p := (Glass{"G", 80}).Price(); p != 80 {
		t.Errorf("Glass.Price() = %.2f, erwartet 80", p)
	}
	if p := (Frame{"F", 120}).Price(); p != 120 {
		t.Errorf("Frame.Price() = %.2f, erwartet 120", p)
	}
}

func TestWindowPrice(t *testing.T) {
	w := NewWindow("W", Frame{"F", 100}, Glass{"G", 50})
	if p := w.Price(); p != 150 {
		t.Errorf("Window.Price() = %.2f, erwartet 150", p)
	}
}

func TestNestedWindowPrice(t *testing.T) {
	inner := NewWindow("Innen", Frame{"F", 10}, Glass{"G", 20})
	outer := NewWindow("Außen", Frame{"F", 100}, inner)
	outer.Add(Glass{"G", 5})
	if p := outer.Price(); p != 135 {
		t.Errorf("verschachtelter Preis = %.2f, erwartet 135", p)
	}
}

func TestDisplay(t *testing.T) {
	inner := NewWindow("Innen", Glass{"Klarglas", 20})
	out := NewWindow("Außen", Frame{"Holz", 100}, inner).Display()

	for _, want := range []string{"Fenster: Außen", "Rahmen: Holz", "  Fenster: Innen", "    Glas: Klarglas"} {
		if !strings.Contains(out, want) {
			t.Errorf("Display enthält %q nicht:\n%s", want, out)
		}
	}
}
