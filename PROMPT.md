# Prompt

Implementiere das Composite Design Pattern in Go anhand folgender UML-Struktur:

- Item ist ein Interface mit den Methoden Price() und Display().
- Glass und Frame sind konkrete Klassen, die Item implementieren und jeweils einen eigenen Preis besitzen.
- Window implementiert ebenfalls Item und ist das Composite. Ein Window enthält 1..* Item-Objekte, also Glass, Frame oder weitere Window-Objekte.
- Price() gibt den jeweiligen Preis zurück. Bei Window wird der Gesamtpreis aller enthaltenen Elemente rekursiv berechnet.
- Display() gibt die Fensterstruktur übersichtlich aus.

Erstelle ein vollständiges, ausführbares Go-Programm mit einer main()-Funktion, die mehrere Fenster mit unterschiedlichen Glas- und Rahmenelementen erstellt und deren Struktur und Gesamtpreis ausgibt.

Verwende ausschließlich die Go-Standardbibliothek. Halte den Code einfach, verständlich und möglichst kompakt. Erstelle außerdem eine README.md mit einer kurzen Erklärung, eine Datei PROMPT.md mit diesem Prompt und einfache Tests.

Prüfe abschließend, ob das Programm mit go run . ausgeführt werden kann und alle Tests mit go test ./... erfolgreich sind.
