# Composite Design Pattern – Fenster

Einfaches Go-Beispiel für das **Composite Pattern** (nur Standardbibliothek).

## Struktur

| Rolle       | Typ      | Beschreibung                                          |
|-------------|----------|-------------------------------------------------------|
| Component   | `Item`   | Interface mit `Price()` und `Display()`               |
| Leaf        | `Glass`  | Glaselement mit eigenem Preis                         |
| Leaf        | `Frame`  | Rahmenelement mit eigenem Preis                       |
| Composite   | `Window` | Enthält 1..* `Item`s (Glass, Frame oder weitere Windows) |

- `Price()` liefert bei Blättern den eigenen Preis, bei `Window` die rekursive Summe aller enthaltenen Elemente.
- `Display()` liefert die Struktur als eingerückten Baum. Die Einrückung bei verschachtelten Fenstern übernimmt die private Hilfsmethode `display(indent)` von `Window`.
- `NewWindow(name, first, rest...)` verlangt mindestens ein Element und erzwingt so die Multiplizität 1..*.

Der Client behandelt einzelne Teile und ganze Fenster einheitlich über das `Item`-Interface.

## Dateien

- `item.go` – Interface und Typen
- `main.go` – Beispiel mit mehreren (auch verschachtelten) Fenstern
- `item_test.go` – Tests

## Ausführen

```sh
go run .
go test ./...
```
