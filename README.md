MIT LICENCE


# Furia2D-Engine
Un motor de videojuego escrito usando Golang 100% usando la poderosa libreria Raylib
                  Y MUY FACIL

![Raylib](https://github.com/raysan5/raylib)
![Raylib-go](https://github.com/gen2brain/raylib-go)
### Ejemplo
- Hola cuadrado
```go
package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you object", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		obj.Instance(nil)

		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}

```
- mas ejemplos en 'examples'
#### Furia2D esta escrito en ingles ,pero el github,documentacion NO.
