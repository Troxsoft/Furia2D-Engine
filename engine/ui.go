package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tawesoft/golib/v2/dialog"
)

var (
	/*				[ON]	[ON]	[ON]			[ON]	[ON]->DRAW NO
	ui objects: UiButton - UiText - UiTextBox - UiSelect - UiMessageBox ->ALLS TYPES OF MESSAGEBOX
	*/

	//text
	ui_text []*UiText
)

type UiText struct {
	text  string
	pos   Position
	size  uint16
	font  *string
	color Color
	hide  bool
}

func NewUiText(text string, pos Position, size uint16) *UiText {
	p := &UiText{
		text:  text,
		pos:   pos,
		size:  size,
		font:  nil,
		color: NewColor2(255, 0, 0),
		hide:  false,
	}
	p = p.add()
	return p
}

func NewUiTextWithFont(text string, font string, pos Position, size uint16) *UiText {
	p := &UiText{
		text:  text,
		pos:   pos,
		size:  size,
		font:  &font,
		color: NewColor2(255, 0, 0),
		hide:  false,
	}
	p = p.add()
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
func (text *UiText) add() *UiText {
	ui_text = append(ui_text, text)
	return ui_text[len(ui_text)-1]
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
