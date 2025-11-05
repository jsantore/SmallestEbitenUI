package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

type uiDemo struct {
	demoUI  ebitenui.UI
	counter int
	name    string
	state   gameState
}

type gameState int

const (
	gameStateMainMenu gameState = iota
	gameStateGameshow
)

func (u *uiDemo) Update() error {
	if u.state == gameStateMainMenu {
		u.demoUI.Update()
	} else {
		//put post UI here
	}
	return nil
}

func (u uiDemo) Draw(screen *ebiten.Image) {
	if u.state == gameStateMainMenu {
		u.demoUI.Draw(screen)
	} else {
		//draw post UI here
	}
}

func (u uiDemo) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(800, 720)
	ebiten.SetWindowTitle("ebitenui demo")
	gui := NewUI()
	game := uiDemo{
		demoUI: gui,
	}
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Fatal(err)
	}
}

func NewUI() ebitenui.UI {
	outerContainer := widget.NewContainer(
		//create an achor layout with 50 pixals of padding on the left
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(50)),
		)),
		//widget.ContainerOpts.WidgetOpts()
	)
	innerContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(
				widget.DirectionVertical,
			),

			widget.RowLayoutOpts.Spacing(5),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    true,
			})),
	)
	textfont := DefaultFont(16)
	buttonFont := DefaultFont(12)
	textInputPict := LoadImage("TextInput1.png")
	buttonPressedPict := LoadImage("Button_02A_Pressed.png")
	buttonNormalPict := LoadImage("Button_02A_Normal.png")
	buttonSelectedPict := LoadImage("Button_02A_Selected.png")
	directions := widget.NewLabel(
		widget.LabelOpts.LabelColor(&widget.LabelColor{
			Idle:     colornames.Aquamarine,
			Disabled: colornames.Gray,
		}),
		widget.LabelOpts.LabelFace(&buttonFont),
	)
	directions.Label = "Enter your Character Name:"
	innerContainer.AddChild(directions)
	nameInput := widget.NewTextInput(
		widget.TextInputOpts.Face(&textfont),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          colornames.Bisque,
			Disabled:      colornames.Gray,
			Caret:         colornames.Black,
			DisabledCaret: colornames.Gray,
		}),
		widget.TextInputOpts.Image(
			&widget.TextInputImage{
				Idle:      image.NewNineSliceBorder(textInputPict, 14),
				Disabled:  image.NewNineSliceBorder(textInputPict, 14),
				Highlight: image.NewNineSliceBorder(textInputPict, 14),
			}),
		widget.TextInputOpts.WidgetOpts(widget.WidgetOpts.MinSize(200, 60)),
		widget.TextInputOpts.Padding(&widget.Insets{
			Top:    0,
			Left:   20,
			Right:  10,
			Bottom: 0,
		}),
	)
	innerContainer.AddChild(nameInput)
	startButton := widget.NewButton(
		widget.ButtonOpts.TextLabel("Start Game"),
		widget.ButtonOpts.TextFace(&buttonFont),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:     colornames.Azure,
			Disabled: colornames.Gray,
			Hover:    colornames.Aquamarine,
			Pressed:  colornames.Aquamarine,
		}),
		widget.ButtonOpts.Image(
			&widget.ButtonImage{
				Idle:            image.NewNineSliceBorder(buttonNormalPict, 10),
				Hover:           image.NewNineSliceBorder(buttonSelectedPict, 10),
				Pressed:         image.NewNineSliceBorder(buttonPressedPict, 10),
				PressedHover:    image.NewNineSliceBorder(buttonPressedPict, 10),
				Disabled:        image.NewNineSliceBorder(buttonNormalPict, 10),
				PressedDisabled: nil,
			}),
		widget.ButtonOpts.TextPadding(&widget.Insets{
			Bottom: 60,
		}),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.MinSize(200, 100)),
	)
	outerContainer.AddChild(innerContainer)
	bottomContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout()),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),

			widget.WidgetOpts.MinSize(200, 100),
		),
	)
	bottomContainer.AddChild(startButton)
	outerContainer.AddChild(bottomContainer)

	GUItoDisplay := ebitenui.UI{
		Container: outerContainer,
	}
	return GUItoDisplay
}

// pulled from the EBiten UI button demo
// uses build in font provided by GO
func DefaultFont(size float64) text.Face {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		panic(err)
	}
	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}
}

func LoadImage(name string) *ebiten.Image {
	Pict, _, err := ebitenutil.NewImageFromFile(name)
	if err != nil {
		fmt.Println("Unable to load background image:", err)
	}
	return Pict
}
