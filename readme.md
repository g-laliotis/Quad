# Quad Series (A–E)

This project implements five different functions (`QuadA`, `QuadB`, `QuadC`, `QuadD`, `QuadE`) in **Go** that draw rectangles of a given width `x` and height `y`.  
Each function prints the rectangle to the standard output using specific ASCII characters, following the defined patterns.

---

## 📖 Project Overview

The goal is to practice loops, conditional statements, and formatted output in Go by implementing different rectangle-drawing algorithms.  

Each function has the following signature:

```go
func QuadX(x, y int)
x → width of the rectangle
y → height of the rectangle
If x or y are non-positive, the function prints nothing.
The functions differ in the way they render corners, edges, and borders.

🖼️ Examples
QuadA
piscine.QuadA(5, 3)
Output:

o---o
|   |
o---o

QuadB
piscine.QuadB(5, 3)
Output:

/***\
*   *
\***/

QuadC
piscine.QuadC(5, 3)
Output:

ABBBA
B   B
CBBBC

QuadD
piscine.QuadD(5, 3)
Output:

ABBBC
B   B
ABBBC

QuadE
piscine.QuadE(5, 3)
Output:

ABBBC
B   B
CBBBA

📂 Project Structure
.
├── piscine/
│   ├── quada.go
│   ├── quadb.go
│   ├── quadc.go
│   ├── quadd.go
│   ├── quade.go
│   ├── go.mod
│   ├── go.sum
│   └── README.md
└── main.go        // test runner (optional)
Each quadX.go file contains the corresponding function implementation.

▶️ Usage
Clone this repository:
https://github.com/g-laliotis/Quad.git
cd quad-series
Run any of the provided test programs. Example:
Program:

package main

import "piscine"

func main() {
    piscine.QuadA(5, 3)
}
Execution:
go run .
Output:
o---o
|   |
o---o

✅ Requirements
Go 1.18+
Standard Go tools (go run, go build, etc.)

✨ Key Learning Points
Working with loops in Go
Handling edge cases for single rows/columns
Printing ASCII-based graphics
Clean modular project structure

📝 License
Giorgos Laliotis
