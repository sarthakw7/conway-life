# Conway's Game of Life — Terminal Edition 🧬

A simple, animated simulation of Conway’s Game of Life rendered right in your terminal using Go.  
This project is part of my journey to build **interactive, visually rich terminal-based applications** using clean architecture and modern Go practices.

---

## 🚀 Vision

I want to build a terminal simulation that feels alive — cleanly structured, interactive, and easy to extend.  
The goal is to start small and gradually evolve it into a powerful, keyboard-controlled, pattern-rich system with beautiful terminal visuals.

This project also serves as a foundation for exploring:
- Animation in terminal
- Grid-based simulation logic
- Input handling and ANSI rendering
- Modular Go architecture & clean Git workflows

---

## ✅ Features (So Far)

- Core Game of Life logic (generation rules)
- Modular project structure (`game`, `render`, `input`)
- Randomized initial state
- Terminal animation loop
- Interactive keyboard controls:
  - `p` = pause/resume
  - `+ / -` = control speed
  - `q` = quit

---

## 🧩 Next Features (Phase 3+)

- [ ] Predefined patterns (glider, blinker, etc.)
- [ ] Save/load initial patterns
- [ ] Toroidal grid (wrap-around edge logic)
- [ ] Unicode block rendering or color-enhanced output
- [ ] Adjustable grid size from CLI
- [ ] Optional FPS counter

---

## 🛠️ How to Run

```bash
git clone https://github.com/sarthakw7/conway-life.git
cd conway-life
go run main.go
