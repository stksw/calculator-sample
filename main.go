package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type cdata struct {
	mem int
	cal string
	flg bool // 計算直後かどうかを判定
}

func createNumButtons(f func(v int)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(3),
		widget.NewButton(strconv.Itoa(7), func() { f(7) }),
		widget.NewButton(strconv.Itoa(8), func() { f(8) }),
		widget.NewButton(strconv.Itoa(9), func() { f(9) }),
		widget.NewButton(strconv.Itoa(4), func() { f(4) }),
		widget.NewButton(strconv.Itoa(5), func() { f(5) }),
		widget.NewButton(strconv.Itoa(6), func() { f(6) }),
		widget.NewButton(strconv.Itoa(1), func() { f(1) }),
		widget.NewButton(strconv.Itoa(2), func() { f(2) }),
		widget.NewButton(strconv.Itoa(3), func() { f(3) }),
		widget.NewButton(strconv.Itoa(0), func() { f(0) }),
	)
	return c
}

func createCalcButtons(f func(c string)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewButton("CL", func() { f("CL") }),
		widget.NewButton("/", func() { f("/") }),
		widget.NewButton("*", func() { f("*") }),
		widget.NewButton("+", func() { f("+") }),
		widget.NewButton("-", func() { f("-") }),
	)
	return c
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("calculator")
	window.SetFixedSize(true)
	label := widget.NewLabel("0")
	label.Alignment = fyne.TextAlignTrailing

	data := cdata{
		mem: 0,
		cal: "",
		flg: false,
	}

	calc := func(num int) {
		switch data.cal {
		case "":
			data.mem = num
		case "+":
			data.mem += num
		case "-":
			data.mem -= num
		case "*":
			data.mem *= num
		case "/":	
			data.mem /= num
		}
		label.SetText(strconv.Itoa(data.mem))
		data.flg = true
	}

	// number button action
	pushNum := func(v int) {
		str := label.Text
		if data.flg {
			str = "0"
			data.flg = false
		}
		str += strconv.Itoa(v)
		num, err := strconv.Atoi(str)
		if err == nil {
			label.SetText(strconv.Itoa(num))
		}
	}

	pushCalc := func(c string) {
		if c == "CL" {
			label.SetText("0")
			data.mem = 0
			data.flg = false
			data.cal = ""
			return
		}

		num, err := strconv.Atoi(label.Text)
		if err != nil {
			return 
		}
		calc(num)
		data.cal = c
	}

	pushEnter := func() {
		num, err := strconv.Atoi(label.Text)
		if err != nil {
			return 
		}
		calc(num)
		data.cal = ""
	}

	btnNum := createNumButtons(pushNum)
	btnCal := createCalcButtons(pushCalc)
	btnEnt := widget.NewButton("Enter", pushEnter)
	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(label, btnEnt, nil, btnCal),
			label, btnEnt, btnNum, btnCal,
		),
	)

	window.Resize(fyne.NewSize(300, 200))
	window.ShowAndRun()
}