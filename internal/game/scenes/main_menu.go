package scenes

import (
	"bytes"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/joelschutz/stagehand"
	darkfalltypes "github.com/mavrw/ravenfall-proto/pkg/darkfall_types"
	"golang.org/x/image/font/gofont/goregular"
)

type MainMenuScene struct {
	BaseScene

	face        text.Face
	headerLbl   *widget.Text
	startButton *widget.Button
	exitButton  *widget.Button
}

func (s *MainMenuScene) Update() error {
	// your update code
	s.ui.Update()

	return nil
}

func (s *MainMenuScene) Draw(screen *ebiten.Image) {
	// your draw code

	s.ui.Draw(screen)
}

func (s *MainMenuScene) Load(state darkfalltypes.DarkfallGameState, manager stagehand.SceneController[darkfalltypes.DarkfallGameState]) {
	// your load code
	s.ui = s.getUI()

	s.sm = manager.(*stagehand.SceneManager[darkfalltypes.DarkfallGameState]) // This type assertion is important
}

func (s *MainMenuScene) Unload() darkfalltypes.DarkfallGameState {
	// your unload code

	return s.state
}

func (s *MainMenuScene) getUI() *ebitenui.UI {
	// load label text font
	s.face, _ = loadFont(18)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(5)))),
	)

	// Because this container has a backgroundImage set we track that the ui is hovered over.
	headerContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eimage.NewNineSliceColor(color.NRGBA{R: 200, G: 200, B: 200, A: 100})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    false,
			}),
			// Uncomment this to not track that you are hovering over this header
			// widget.WidgetOpts.TrackHover(false),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	s.headerLbl = widget.NewText(
		widget.TextOpts.Text("Game Demo!", s.face, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  false,
				StretchVertical:    false,
			}),
			// Uncomment to force tracking hover of this element
			// widget.WidgetOpts.TrackHover(true),
		),
	)
	headerContainer.AddChild(s.headerLbl)
	rootContainer.AddChild(headerContainer)

	// Button container
	buttonContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eimage.NewNineSliceColor(color.NRGBA{R: 200, G: 200, B: 200, A: 100})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  false,
				StretchVertical:    false,
			}),
			// Uncomment this to not track that you are hovering over this header
			// widget.WidgetOpts.TrackHover(false),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	rootContainer.AddChild(buttonContainer)

	// start button
	s.startButton = s.getStartButton()
	buttonContainer.AddChild(s.startButton)

	// TODO: Figure out what I am doing wrong that makes these buttons stack on each other

	// exit button
	s.exitButton = s.getExitButton()
	buttonContainer.AddChild(s.exitButton)

	hProgressbar := widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    false,
			}),
			// Set the minimum size for the progress bar.
			// This is necessary if you wish to have the progress bar be larger than
			// the provided track image. In this exampe since we are using NineSliceColor
			// which is 1px x 1px we must set a minimum size.
			widget.WidgetOpts.MinSize(200, 20),
			// Set this parameter to indicate we want do not want to track that this ui element is being hovered over.
			// widget.WidgetOpts.TrackHover(false),
		),
		widget.ProgressBarOpts.Images(
			// Set the track images (Idle, Disabled).
			&widget.ProgressBarImage{
				Idle: eimage.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			// Set the progress images (Idle, Disabled).
			&widget.ProgressBarImage{
				Idle: eimage.NewNineSliceColor(color.NRGBA{255, 255, 100, 255}),
			},
		),
		// Set the min, max, and current values.
		widget.ProgressBarOpts.Values(0, 10, 7),
		// Set how much of the track is displayed when the bar is overlayed.
		widget.ProgressBarOpts.TrackPadding(widget.Insets{
			Top:    2,
			Bottom: 2,
		}),
	)

	rootContainer.AddChild(hProgressbar)

	// Create a label to show the percentage on top of the progress bar
	label2 := widget.NewText(
		widget.TextOpts.Text("70%", s.face, color.Black),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  false,
				StretchVertical:    false,
			}),
		),
	)
	rootContainer.AddChild(label2)

	return &ebitenui.UI{
		Container: rootContainer,
		//Call a render method after the rootContainer is drawn but before any ebitenui.Windows are drawn
		//PostRenderHook: s.Render,
	}
}

func (s *MainMenuScene) getExampleButton() *widget.Button {
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	var button *widget.Button
	// construct a button
	button = widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		//widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
		widget.ButtonOpts.Text("Hello, [color=FF00FF]World![/color]", s.face, &widget.ButtonTextColor{
			Idle:    color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
			Hover:   color.NRGBA{0, 255, 128, 255},
			Pressed: color.NRGBA{255, 0, 0, 255},
		}),
		widget.ButtonOpts.TextProcessBBCode(true),
		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		//Move the text down and right on press
		widget.ButtonOpts.PressedHandler(func(args *widget.ButtonPressedEventArgs) {
			button.Text().Inset.Top = 4
			button.Text().Inset.Left = 4
			button.GetWidget().CustomData = true
		}),
		//Move the text back to start on press released
		widget.ButtonOpts.ReleasedHandler(func(args *widget.ButtonReleasedEventArgs) {
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
			button.GetWidget().CustomData = false
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("button clicked")
		}),

		// add a handler that reacts to entering the button with the cursor
		widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor entered button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
			//If we moved the Text because we clicked on this button previously, move the text down and right
			if button.GetWidget().CustomData == true {
				button.Text().Inset.Top = 4
				button.Text().Inset.Left = 4
			}
		}),

		// add a handler that reacts to moving the cursor on the button
		widget.ButtonOpts.CursorMovedHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor moved on button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY, "diffX =", args.DiffX, "diffY =", args.DiffY)
		}),

		// add a handler that reacts to exiting the button with the cursor
		widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor exited button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
			//Reset the Text inset if the cursor is no longer over the button
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
		}),

		// Indicate that this button should not be submitted when enter or space are pressed
		// widget.ButtonOpts.DisableDefaultKeys(),
	)

	return button
}

func (s *MainMenuScene) getStartButton() *widget.Button {
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	var button *widget.Button
	// construct a button
	button = widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		//widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
		widget.ButtonOpts.Text("Start", s.face, &widget.ButtonTextColor{
			Idle:    color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
			Hover:   color.NRGBA{0, 255, 128, 255},
			Pressed: color.NRGBA{255, 0, 0, 255},
		}),
		widget.ButtonOpts.TextProcessBBCode(true),
		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		//Move the text down and right on press
		widget.ButtonOpts.PressedHandler(func(args *widget.ButtonPressedEventArgs) {
			button.Text().Inset.Top = 4
			button.Text().Inset.Left = 4
			button.GetWidget().CustomData = true
		}),
		//Move the text back to start on press released
		widget.ButtonOpts.ReleasedHandler(func(args *widget.ButtonReleasedEventArgs) {
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
			button.GetWidget().CustomData = false
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("start clicked")
		}),

		// add a handler that reacts to entering the button with the cursor
		widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor entered start button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
			//If we moved the Text because we clicked on this button previously, move the text down and right
			if button.GetWidget().CustomData == true {
				button.Text().Inset.Top = 4
				button.Text().Inset.Left = 4
			}
		}),

		// add a handler that reacts to moving the cursor on the button
		widget.ButtonOpts.CursorMovedHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor moved on start button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY, "diffX =", args.DiffX, "diffY =", args.DiffY)
		}),

		// add a handler that reacts to exiting the button with the cursor
		widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor exited start button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
			//Reset the Text inset if the cursor is no longer over the button
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
		}),

		// Indicate that this button should not be submitted when enter or space are pressed
		// widget.ButtonOpts.DisableDefaultKeys(),
	)

	return button
}

func (s *MainMenuScene) getExitButton() *widget.Button {
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	var button *widget.Button
	// construct a button
	button = widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		//widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
		widget.ButtonOpts.Text("Exit", s.face, &widget.ButtonTextColor{
			Idle:    color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
			Hover:   color.NRGBA{0, 255, 128, 255},
			Pressed: color.NRGBA{255, 0, 0, 255},
		}),
		widget.ButtonOpts.TextProcessBBCode(true),
		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		//Move the text down and right on press
		widget.ButtonOpts.PressedHandler(func(args *widget.ButtonPressedEventArgs) {
			button.Text().Inset.Top = 4
			button.Text().Inset.Left = 4
			button.GetWidget().CustomData = true
		}),
		//Move the text back to start on press released
		widget.ButtonOpts.ReleasedHandler(func(args *widget.ButtonReleasedEventArgs) {
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
			button.GetWidget().CustomData = false
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("exit clicked")
		}),

		// add a handler that reacts to entering the button with the cursor
		widget.ButtonOpts.CursorEnteredHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor entered exit button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
			//If we moved the Text because we clicked on this button previously, move the text down and right
			if button.GetWidget().CustomData == true {
				button.Text().Inset.Top = 4
				button.Text().Inset.Left = 4
			}
		}),

		// add a handler that reacts to moving the cursor on the button
		widget.ButtonOpts.CursorMovedHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor moved on exit button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY, "diffX =", args.DiffX, "diffY =", args.DiffY)
		}),

		// add a handler that reacts to exiting the button with the cursor
		widget.ButtonOpts.CursorExitedHandler(func(args *widget.ButtonHoverEventArgs) {
			println("cursor exited exit button: entered =", args.Entered, "offsetX =", args.OffsetX, "offsetY =", args.OffsetY)
			//Reset the Text inset if the cursor is no longer over the button
			button.Text().Inset.Top = 0
			button.Text().Inset.Left = 0
		}),

		// Indicate that this button should not be submitted when enter or space are pressed
		// widget.ButtonOpts.DisableDefaultKeys(),
	)

	return button
}

func loadFont(size float64) (text.Face, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size,
	}, nil
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := eimage.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := eimage.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := eimage.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
