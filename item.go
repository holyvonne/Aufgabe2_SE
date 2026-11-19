package main

import (
	"fmt"
	"strings"
)

// Item ist die gemeinsame Schnittstelle (Component) für Blätter und Composites.
type Item interface {
	Price() float64
	Display() string
}

// Glass ist ein Blatt mit eigenem Preis.
type Glass struct {
	Name string
	Cost float64
}

func (g Glass) Price() float64 { return g.Cost }

func (g Glass) Display() string {
	return fmt.Sprintf("Glas: %s (%.2f €)\n", g.Name, g.Cost)
}

// Frame ist ein Blatt mit eigenem Preis.
type Frame struct {
	Name string
	Cost float64
}

func (f Frame) Price() float64 { return f.Cost }

func (f Frame) Display() string {
	return fmt.Sprintf("Rahmen: %s (%.2f €)\n", f.Name, f.Cost)
}

// Window ist das Composite und enthält 1..* Items (Glass, Frame oder Window).
type Window struct {
	Name  string
	Items []Item
}

// NewWindow erzwingt die Multiplizität 1..*.
func NewWindow(name string, first Item, rest ...Item) *Window {
	return &Window{Name: name, Items: append([]Item{first}, rest...)}
}

// Add fügt ein weiteres Element hinzu.
func (w *Window) Add(item Item) { w.Items = append(w.Items, item) }

// Price berechnet den Gesamtpreis rekursiv.
func (w *Window) Price() float64 {
	total := 0.0
	for _, it := range w.Items {
		total += it.Price()
	}
	return total
}

// Display gibt die Struktur rekursiv mit Einrückung aus.
func (w *Window) Display() string { return w.display(0) }

// display ist die private Hilfsmethode, die die Einrückungstiefe mitführt.
func (w *Window) display(indent int) string {
	pad := strings.Repeat("  ", indent)
	var sb strings.Builder
	fmt.Fprintf(&sb, "%sFenster: %s (gesamt %.2f €)\n", pad, w.Name, w.Price())
	for _, it := range w.Items {
		if child, ok := it.(*Window); ok {
			sb.WriteString(child.display(indent + 1))
		} else {
			sb.WriteString(pad + "  " + it.Display())
		}
	}
	return sb.String()
}
