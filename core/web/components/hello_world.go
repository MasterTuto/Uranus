package components

type helloWorld struct {}

func (c helloWorld) Get() {
    return c.
}

func (c helloWorld) Style() renderer.ElementStyle {
    return renderer.ElementStyle{
        Color:           "#ff0000",
        BackgroundColor: "#f0f",
        FontFamily:      "font.ttf",
        FontSize:        32,
    };
}

func HelloWorld() helloWorld {
    return helloWorld{}
}


