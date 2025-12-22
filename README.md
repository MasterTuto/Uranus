# Uranus

Uranus is a lightweight experimental UI rendering engine written in Go, leveraging SDL2 for window management, text rendering, and event handling. The project’s primary goal is to **implement a CSS-like styling and layout system** directly in Go, allowing developers to describe UI elements declaratively using familiar concepts such as colors, fonts, dimensions, and positioning.

This engine serves as a foundation for building custom UI frameworks, in-game interfaces, or visualization tools using a styling model inspired by CSS.

---

## Features

* CSS-inspired styling model (colors, fonts, sizes, element types).
* Basic layout engine for placing elements on screen.
* Simplified window creation and lifecycle management using SDL2.
* Declarative element definitions (`LayoutedElement`).
* Text rendering with custom fonts.
* Event loop management.

---

## Getting Started

### Requirements

* **Go 1.21+**
* **SDL2** installed on your system
  (Ubuntu/Debian example: `sudo apt install libsdl2-dev libsdl2-ttf-dev`)

### Installation

```bash
git clone https://github.com/MasterTuto/Uranus.git
cd Uranus
go mod tidy
```

### Running the Example

```bash
go run ./cmd/uranus
```

You should see a window displaying the text **“Hello, World”**, styled via the engine’s CSS-like API (font, size, color, etc.).

---

## Project Structure

```
Uranus/
├── cmd/
│   └── uranus/
│       └── main.go          # Example application
├── pkg/
│   ├── renderer/            # Rendering engine (layout + CSS-style rules)
│   └── window/              # Window and event lifecycle
├── go.mod
└── go.sum
```

### `cmd/uranus/main.go`

The example demonstrates:

1. Initializing the windowing system.
2. Creating an SDL2-based rendering context.
3. Declaring an element with CSS-inspired attributes.
4. Rendering UI elements.
5. Running the event loop.

```go
elements := []renderer.LayoutedElement{
    {
        Type: renderer.Text,
        PosX: 0, PosY: 0,
        Height: 100, Width: 100,
        Element: renderer.Element{
            TextContent: "Hello, World",
            Style: renderer.ElementStyle{
                Color: "#fff",
                FontFamily: "font.ttf",
                FontSize: 32,
            },
        },
    },
}
```

---

## CSS-like Styling

The `ElementStyle` structure is inspired by CSS properties:

| CSS Concept    | Uranus Field            |
| -------------- | ----------------------- |
| `color`        | `Color` (`#RRGGBB` hex) |
| `font-family`  | `FontFamily`            |
| `font-size`    | `FontSize`              |
| `width/height` | Element width/height    |
| `position`     | `PosX`, `PosY`          |

The long-term intention is to expand this model to include:

* margin/padding
* auto layout
* flexbox-like rules
* borders & background
* element nesting / DOM-like hierarchy

---

## Dependencies

This project uses:

* **github.com/veandco/go-sdl2 v0.4.35** — Go bindings for SDL2

```go
require github.com/veandco/go-sdl2 v0.4.35
```

---

## Development

Build everything:

```bash
go build ./...
```

Run all tests (future):

```bash
go test ./...
```

Clear build cache:

```bash
go clean -cache
```

---

## Roadmap

* Expand CSS-like style model
* Add layout engines (flexbox, grid)
* Add image components
* Add event-aware components (buttons, inputs)
* Component tree / DOM abstraction
* Animation and transitions

---

## License

MIT License. See the `LICENSE` file for details.
