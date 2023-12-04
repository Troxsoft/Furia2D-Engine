package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tawesoft/golib/v2/dialog"
)

type UiButton struct {
	pos    Position
	size   Size
	Zone   *UiColorZone
	Text   *UiText
	hide   bool
	colorA Color
	wr     bool
}

func NewUiButton(scene *Scene, txt string, color Color, pos Position, siz Size) *UiButton {
	p := &UiButton{
		pos:    pos,
		size:   siz,
		hide:   false,
		Zone:   NewUiColorZone(scene, color, pos, siz),
		Text:   NewUiText(scene, txt, NewPosition(int32(siz.W)/2+pos.X-int32(len(txt))*3, int32(siz.H)/2+pos.Y), 15),
		colorA: color,
		wr:     false,
	}
	p = p.add(scene)

	return p
}
func (b *UiButton) SetColor(color Color) {
	b.colorA = color
}
func (b *UiButton) Color() Color {
	return b.colorA
}
func (b *UiButton) Hide() {
	b.hide = true
	b.Zone.hide = true
	b.Text.hide = true
}
func (b *UiButton) Show() {
	b.hide = false
	b.Zone.hide = false
	b.Text.hide = false
}
func (b *UiButton) add(scene *Scene) *UiButton {
	scene.ui_button = append(scene.ui_button, b)
	return scene.ui_button[len(scene.ui_button)-1]
}
func (b *UiButton) Draw() {
	b.Text.pos.X = int32(b.size.W)/2 + b.pos.X - int32(len(b.Text.text))*3
	b.Text.pos.Y = int32(b.size.H)/2 + b.pos.Y
	x := GetMousePosition().X
	y := GetMousePosition().Y

	if b.Zone.color.R-20 > 0 && b.Zone.color.G-20 > 0 && b.Zone.color.B-20 > 0 {
		//b.colorA = *NewColor2(b.Zone.color.R-20, b.Zone.color.G-20, b.Zone.color.B-20)
		if rl.CheckCollisionRecs(rl.NewRectangle(float32(x), float32(y), 2, 2), rl.NewRectangle(float32(b.pos.X), float32(b.pos.Y), float32(b.size.W), float32(b.size.H))) {
			if b.wr == false {
				b.Zone.color = NewColor2(b.Zone.color.R-20, b.Zone.color.G-20, b.Zone.color.B-20)

			}
			if b.wr == false {
				b.wr = true
			}
		} else {
			b.Zone.color = b.colorA
			b.wr = false
		}
	} else {
		b.wr = false
		b.Zone.SetColor(b.colorA)
	}
	rl.DrawRectangleRounded(rl.NewRectangle(float32(b.pos.X)-2, float32(b.pos.Y)-2, float32(b.size.W)+2, float32(b.size.H)+2), b.Zone.roundness, int32(b.Zone.segments+1000), ConvertColor(NewColor(50, 50, 70, 255)))
	b.Zone.Draw()
	b.Text.Draw()

}
func (b *UiButton) IsPressed() bool {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) {
		x := GetMousePosition().X
		y := GetMousePosition().Y
		if rl.CheckCollisionRecs(rl.NewRectangle(float32(x), float32(y), 2, 2), rl.NewRectangle(float32(b.pos.X), float32(b.pos.Y), float32(b.size.W), float32(b.size.H))) {
			return true
		}
	}
	return false
}
func (b *UiButton) Size() Size {
	return b.size
}
func (b *UiButton) SetSize(nsize Size) {
	b.size = nsize
	b.Zone.size = nsize
	b.Text.pos.X = int32(b.size.W)/2 + b.pos.X - int32(len(b.Text.text))*3
	b.Text.pos.Y = int32(b.size.H)/2 + b.pos.Y
	//NewPosition((pos.X/2)+(pos.X-int32(siz.W)/2), (pos.Y/2)+(pos.Y-int32(siz.H)/2))
}

func (b *UiButton) Position() Position {
	return b.pos
}
func (b *UiButton) SetPosition(nPos Position) {
	b.pos = nPos
	b.Zone.pos = b.pos
	b.Text.pos.X = int32(b.size.W)/2 + b.pos.X - int32(len(b.Text.text))*3
	b.Text.pos.Y = int32(b.size.H)/2 + b.pos.Y
}

type UiColorZone struct {
	roundness float32
	color     Color
	pos       Position
	size      Size
	segments  int
	hide      bool
}

func NewUiColorZone(scene *Scene, color Color, pos Position, size Size) *UiColorZone {
	p := &UiColorZone{
		roundness: 0,
		color:     color,
		pos:       pos,
		size:      size,
		segments:  50,
		hide:      false,
	}
	p = p.add(scene)
	return p
}
func (c *UiColorZone) SetSegments(s int) {
	c.segments = s
}
func (c *UiColorZone) Segments() int {
	return c.segments
}

func (c *UiColorZone) SetRoundness(r float32) {
	c.roundness = r
}
func (c *UiColorZone) Roundness() float32 {
	return c.roundness
}
func (c *UiColorZone) Hide() {
	//rl.DrawRectangle
	c.hide = true
}
func (c *UiColorZone) Show() {
	c.hide = false
}
func (c *UiColorZone) Draw() {
	if c.hide == false {
		rl.DrawRectangleRounded(rl.NewRectangle(float32(c.pos.X), float32(c.pos.Y), float32(c.size.W), float32(c.size.H)), c.roundness, int32(c.segments), ConvertColor(c.color))
	}
}
func (c *UiColorZone) add(scene *Scene) *UiColorZone {
	scene.ui_colorZone = append(scene.ui_colorZone, c)
	return scene.ui_colorZone[len(scene.ui_colorZone)-1]
}
func (c *UiColorZone) SetColor(nColor Color) {
	c.color = nColor
}
func (c *UiColorZone) SetColor2(r, g, b, a uint8) {
	c.color = NewColor(r, g, b, a)
}
func (c *UiColorZone) SetColor3(r, g, b uint8) {
	c.color = NewColor2(r, g, b)
}
func (c *UiColorZone) Color() Color {
	return c.color
}
func (c *UiColorZone) SetSize(nSize Size) {
	c.size = nSize
}
func (c *UiColorZone) Size() Size {
	return c.size
}
func (c *UiColorZone) SetPosition(nPos Position) {
	c.pos = nPos
}
func (c *UiColorZone) Position() Position {
	return c.pos
}

type UiText struct {
	text  string
	pos   Position
	size  uint16
	font  *string
	color Color
	hide  bool
}

func NewUiText(scene *Scene, text string, pos Position, size uint16) *UiText {
	p := &UiText{
		text:  text,
		pos:   pos,
		size:  size,
		font:  nil,
		color: NewColor2(255, 0, 0),
		hide:  false,
	}
	p = p.add(scene)
	return p
}

func NewUiTextWithFont(scene *Scene, text string, font string, pos Position, size uint16) *UiText {
	p := &UiText{
		text:  text,
		pos:   pos,
		size:  size,
		font:  &font,
		color: NewColor2(255, 0, 0),
		hide:  false,
	}
	p = p.add(scene)
	return p
}
func (text *UiText) Hide() {
	//rl.DrawRectangle
	text.hide = true
}
func (text *UiText) Show() {
	text.hide = false
}
func (text *UiText) Draw() {
	if text.hide == false {

		if text.font == nil {

			rl.DrawText(text.text, text.pos.X, text.pos.Y, int32(text.size), ConvertColor(text.color))
		} else {
			rl.DrawTextEx(rl.LoadFont(text.FontName()), text.text, rl.NewVector2(float32(text.pos.X), float32(text.pos.Y)), float32(text.size), 0.5, ConvertColor(text.color))
		}
	}
}
func (text *UiText) add(scene *Scene) *UiText {
	scene.ui_text = append(scene.ui_text, text)
	return scene.ui_text[len(scene.ui_text)-1]
}
func (text *UiText) SetColor(nColor Color) {
	text.color = nColor
}
func (text *UiText) SetColor2(r, g, b, a uint8) {
	text.color = NewColor(r, g, b, a)
}
func (text *UiText) SetColor3(r, g, b uint8) {
	text.color = NewColor2(r, g, b)
}
func (text *UiText) Color() Color {
	return text.color
}
func (text *UiText) SetFontName(font string) {
	if font == "" || font == "default" {
		text.font = nil
	} else {
		text.font = &font
	}
}

func (text *UiText) FontName() string {
	if text.font == nil {
		return "default"
	} else {
		return *text.font
	}
}
func (text *UiText) SetSize(nSize uint16) {
	text.size = nSize
}
func (text *UiText) Size() uint16 {
	return text.size
}
func (text *UiText) SetPosition(nPos Position) {
	text.pos = nPos
}
func (text *UiText) Position() Position {
	return text.pos
}
func (text *UiText) SetText(nText string) {
	text.text = nText
}
func (text *UiText) SetTextAsFormat(format string, pa ...any) {
	text.text = fmt.Sprintf(format, pa...)
}
func (text *UiText) SetTextAsString(pa any) {
	text.text = fmt.Sprint(pa)
}
func (text *UiText) Text() string {
	return text.text
}

type UiMessageBoxAlert struct {
	title   string
	message string
}

func ShowMessageBoxAlert(title, message string) {
	NewMessageBoxAlert(title, message).Show()
}
func NewMessageBoxAlert(title, message string) *UiMessageBoxAlert {
	return &UiMessageBoxAlert{
		message: message,
		title:   title,
	}
}
func (mb *UiMessageBoxAlert) Show() {
	dialog.Message{
		Title:  mb.title,
		Format: mb.message,
		Icon:   dialog.IconWarning,
	}.Raise()
}

//info

type UiMessageBoxInfo struct {
	title   string
	message string
}

func ShowMessageBoxInfo(title, message string) {
	NewMessageBoxInfo(title, message).Show()
}
func NewMessageBoxInfo(title, message string) *UiMessageBoxInfo {
	return &UiMessageBoxInfo{
		message: message,
		title:   title,
	}
}
func (mb *UiMessageBoxInfo) Show() {
	dialog.Message{
		Title:  mb.title,
		Format: mb.message,
		Icon:   dialog.IconInfo,
	}.Raise()
}

type UiMessageBoxError struct {
	title   string
	message string
}

func ShowMessageBoxError(title, message string) {
	NewMessageBoxError(title, message).Show()
}
func NewMessageBoxError(title, message string) *UiMessageBoxError {
	return &UiMessageBoxError{
		message: message,
		title:   title,
	}
}
func (mb *UiMessageBoxError) Show() {
	dialog.Message{
		Title:  mb.title,
		Format: mb.message,
		Icon:   dialog.IconError,
	}.Raise()
}
