package main

import (
	"fmt"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//создание нового приложения
	calculator := app.New()

	//графическое окно приложения
	window := calculator.NewWindow("calculator")
	window.Resize(fyne.NewSize(550, 550))
	calculator.Settings().SetTheme(theme.DarkTheme())

	//ссылки
	url_1, _ := url.Parse("https://www.desmos.com/scientific")
	url_2, _ := url.Parse("https://t.me/dvrkv")
	link_1 := widget.NewHyperlinkWithStyle("more advanced calculator", url_1, fyne.TextAlign(20), fyne.TextStyle{Italic: true})
	link_2 := widget.NewHyperlinkWithStyle("designed by @dvrkv", url_2, fyne.TextAlign(20), fyne.TextStyle{Italic: true})

	//картинки
	img_addition := canvas.NewImageFromFile("img/addition.svg")
	img_subtraction := canvas.NewImageFromFile("img/subtraction.svg")
	img_multiplication := canvas.NewImageFromFile("img/multiplication.svg")
	img_division := canvas.NewImageFromFile("img/division.svg")

	//поля ввода
	entry_1 := widget.NewEntry()
	entry_1.Resize(fyne.NewSize(500, 40))
	entry_1.Move(fyne.NewPos(20, 30))

	entry_2 := widget.NewEntry()
	entry_2.Resize(fyne.NewSize(500, 40))
	entry_2.Move(fyne.NewPos(20, 80))

	//подсказки ввода
	entry_1.SetPlaceHolder("Введите первое число...")
	entry_2.SetPlaceHolder("Введите второе число...")

	//надписи
	label_name := widget.NewLabel("КАЛЬКУЛЯТОР")
	label_name.TextStyle = fyne.TextStyle{Bold: true}
	label_name.Move(fyne.NewPos(12, 0))

	label_result := widget.NewLabel("Результат")
	label_result.Move(fyne.NewPos(12, 180))

	label_theme := widget.NewLabel("ТЕМА")
	label_theme.TextStyle = fyne.TextStyle{Bold: true}
	label_theme.Move(fyne.NewPos(12, 320))

	label_links := widget.NewLabel("ССЫЛКИ")
	label_links.TextStyle = fyne.TextStyle{Bold: true}
	label_links.Move(fyne.NewPos(12, 340))

	result := widget.NewLabel("")
	result.TextStyle = fyne.TextStyle{Bold: true}
	result.Move(fyne.NewPos(12, 210))

	//кнопки
	button_addition := container.New(
		layout.NewMaxLayout(),
		img_addition,
		widget.NewButton("", func() {
			//конвертация в float64 и получение значений из полей ввода
			num_1, err_1 := strconv.ParseFloat(entry_1.Text, 64)
			num_2, err_2 := strconv.ParseFloat(entry_2.Text, 64)
			if err_1 != nil || err_2 != nil {
				//всплывающее диалоговое окно с ошибкой
				dialog.ShowCustom(
					"OOOPS :(",
					"Ok",
					widget.NewLabel("Ошибка ввода !"),
					window,
				)
			} else {
				addition := num_1 + num_2

				//вывод ответа и конвертация в string
				result.SetText(fmt.Sprintf("%.2f", addition))
			}
		}))
	button_addition.Move(fyne.NewPos(20, 130))
	button_addition.Resize(fyne.NewSize(40, 40))

	button_subtraction := container.New(
		layout.NewMaxLayout(),
		img_subtraction,
		widget.NewButton("", func() {
			//конвертация в float64 и получение значений из полей ввода
			num_1, err_1 := strconv.ParseFloat(entry_1.Text, 64)
			num_2, err_2 := strconv.ParseFloat(entry_2.Text, 64)
			if err_1 != nil || err_2 != nil {
				//всплывающее диалоговое окно с ошибкой
				dialog.ShowCustom(
					"OOOPS :(",
					"Ok",
					widget.NewLabel("Ошибка ввода !"),
					window,
				)
			} else {
				subtraction := num_1 - num_2

				//вывод ответа и конвертация в string
				result.SetText(fmt.Sprintf("%.2f", subtraction))
			}
		}))
	button_subtraction.Resize(fyne.NewSize(40, 40))
	button_subtraction.Move(fyne.NewPos(75, 130))

	button_multiplication := container.New(
		layout.NewMaxLayout(),
		img_multiplication,
		widget.NewButton("", func() {
			//конвертация в float64 и получение значений из полей ввода
			num_1, err_1 := strconv.ParseFloat(entry_1.Text, 64)
			num_2, err_2 := strconv.ParseFloat(entry_2.Text, 64)
			if err_1 != nil || err_2 != nil {
				//всплывающее диалоговое окно с ошибкой
				dialog.ShowCustom(
					"OOOPS :(",
					"Ok",
					widget.NewLabel("Ошибка ввода !"),
					window,
				)
			} else {
				multiplication := num_1 * num_2

				//вывод ответа и конвертация в string
				result.SetText(fmt.Sprintf("%.2f", multiplication))
			}
		}))
	button_multiplication.Resize(fyne.NewSize(40, 40))
	button_multiplication.Move(fyne.NewPos(130, 130))

	button_division := container.New(
		layout.NewMaxLayout(),
		img_division,
		widget.NewButton("", func() {
			//конвертация в float64 и получение значений из полей ввода
			num_1, err_1 := strconv.ParseFloat(entry_1.Text, 64)
			num_2, err_2 := strconv.ParseFloat(entry_2.Text, 64)
			if err_1 != nil || err_2 != nil {
				//всплывающее диалоговое окно с ошибкой
				dialog.ShowCustom(
					"OOOPS :(",
					"Ok",
					widget.NewLabel("Ошибка ввода !"),
					window,
				)
			} else if num_2 == 0 {
				dialog.ShowCustom(
					"OOOPS :(",
					"Ok",
					widget.NewLabel("На 0 делить нельзя !"),
					window,
				)
			} else {
				division := num_1 / num_2

				//вывод ответа и конвертация в string
				result.SetText(fmt.Sprintf("%.2f", division))
			}
		}))
	button_division.Resize(fyne.NewSize(40, 40))
	button_division.Move(fyne.NewPos(185, 130))

	button_light := widget.NewButton("Светлая", func() {
		calculator.Settings().SetTheme(theme.LightTheme())
	})
	button_light.Resize(fyne.NewSize(95, 40))
	button_light.Move(fyne.NewPos(20, 360))

	button_dark := widget.NewButton("Темная", func() {
		calculator.Settings().SetTheme(theme.DarkTheme())
	})
	button_dark.Resize(fyne.NewSize(95, 40))
	button_dark.Move(fyne.NewPos(130, 360))

	//панель с цифрами
	n_1 := widget.NewButton("1", func() {
		entry_1.SetText("1")
	})
	n_1.Resize(fyne.NewSize(70, 40))
	n_1.Move(fyne.NewPos(300, 130))

	n_2 := widget.NewButton("2", func() {
		entry_1.SetText("2")
	})
	n_2.Resize(fyne.NewSize(70, 40))
	n_2.Move(fyne.NewPos(375, 130))

	n_3 := widget.NewButton("3", func() {
		entry_1.SetText("3")
	})
	n_3.Resize(fyne.NewSize(70, 40))
	n_3.Move(fyne.NewPos(450, 130))

	n_4 := widget.NewButton("4", func() {
		entry_1.SetText("4")
	})
	n_4.Resize(fyne.NewSize(70, 40))
	n_4.Move(fyne.NewPos(300, 175))

	n_5 := widget.NewButton("5", func() {
		entry_1.SetText("5")
	})
	n_5.Resize(fyne.NewSize(70, 40))
	n_5.Move(fyne.NewPos(375, 175))

	n_6 := widget.NewButton("6", func() {
		entry_1.SetText("6")
	})
	n_6.Resize(fyne.NewSize(70, 40))
	n_6.Move(fyne.NewPos(450, 175))

	n_7 := widget.NewButton("7", func() {
		entry_1.SetText("7")
	})
	n_7.Resize(fyne.NewSize(70, 40))
	n_7.Move(fyne.NewPos(300, 220))

	n_8 := widget.NewButton("8", func() {
		entry_1.SetText("8")
	})
	n_8.Resize(fyne.NewSize(70, 40))
	n_8.Move(fyne.NewPos(375, 220))

	n_9 := widget.NewButton("9", func() {
		entry_1.SetText("9")
	})
	n_9.Resize(fyne.NewSize(70, 40))
	n_9.Move(fyne.NewPos(450, 220))

	n_0 := widget.NewButton("0", func() {
		entry_1.SetText("0")
	})
	n_0.Resize(fyne.NewSize(220, 40))
	n_0.Move(fyne.NewPos(300, 265))

	//коробки
	box_calculator := container.NewWithoutLayout(
		label_name,
		entry_1,
		entry_2,
		button_addition,
		button_subtraction,
		button_multiplication,
		button_division,
	)
	box_result := container.NewWithoutLayout(
		label_result,
		result,
	)
	box_theme := container.NewWithoutLayout(
		label_theme,
		button_dark,
		button_light,
	)
	box_links := container.NewGridWithColumns(
		1,
		label_links,
		link_1,
		link_2,
	)
	box_links.Move(fyne.NewPos(12, 420))
	box_numbers := container.NewWithoutLayout(
		n_1,
		n_2,
		n_3,
		n_4,
		n_5,
		n_6,
		n_7,
		n_8,
		n_9,
		n_0,
	)

	//содержимое окна приложения
	window.SetContent(container.NewWithoutLayout(
		box_calculator,
		box_result,
		box_theme,
		box_links,
		box_numbers,
	))

	//запуск приложения
	window.ShowAndRun()
}
