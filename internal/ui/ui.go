package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/qro/password/internal/generator"
	"github.com/qro/password/internal/strength"
)

func Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("github.com/qro/password")
	myWindow.Resize(fyne.NewSize(520, 480))

	generatorTab := container.NewTabItem("Generator", buildGeneratorTab(myWindow))
	strengthTab := container.NewTabItem("Strength Checker", buildStrengthTab())

	tabs := container.NewAppTabs(generatorTab, strengthTab)
	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func buildGeneratorTab(myWindow fyne.Window) fyne.CanvasObject {
	// Character set checkboxes
	uppercaseCheck := widget.NewCheck("Uppercase (A-Z)", nil)
	uppercaseCheck.SetChecked(true)

	lowercaseCheck := widget.NewCheck("Lowercase (a-z)", nil)
	lowercaseCheck.SetChecked(true)

	numbersCheck := widget.NewCheck("Numbers (0-9)", nil)
	numbersCheck.SetChecked(true)

	symbolsCheck := widget.NewCheck("Symbols (!@#$...)", nil)
	symbolsCheck.SetChecked(false)

	checkboxGroup := container.NewGridWithColumns(2,
		uppercaseCheck, lowercaseCheck,
		numbersCheck, symbolsCheck,
	)

	// Password length slider
	lengthLabel := widget.NewLabel("Length: 16")
	lengthSlider := widget.NewSlider(8, 128)
	lengthSlider.Value = 16
	lengthSlider.Step = 1
	lengthSlider.OnChanged = func(val float64) {
		lengthLabel.SetText("Length: " + intToStr(int(val)))
	}

	lengthRow := container.NewBorder(nil, nil, lengthLabel, nil, lengthSlider)

	// Output field
	outputEntry := widget.NewPasswordEntry()
	outputEntry.SetPlaceHolder("Generated password will appear here...")

	// Buttons
	generateBtn := widget.NewButton("Generate", func() {
		opts := generator.Options{
			Length:  int(lengthSlider.Value),
			Upper:   uppercaseCheck.Checked,
			Lower:   lowercaseCheck.Checked,
			Digits:  numbersCheck.Checked,
			Symbols: symbolsCheck.Checked,
		}
		pw, err := generator.Generate(opts)
		if err != nil {
			outputEntry.SetText("Error: " + err.Error())
			return
		}
		outputEntry.SetText(pw)
	})
	generateBtn.Importance = widget.HighImportance

	// Copy
	copyBtn := widget.NewButton("Copy", func() {
		if outputEntry.Text != "" && myWindow != nil {
			myWindow.Clipboard().SetContent(outputEntry.Text)
		}
	})

	buttonRow := container.NewHBox(generateBtn, copyBtn, layout.NewSpacer())

	// Order
	return container.NewVBox(
		widget.NewLabel("Character Sets:"),
		checkboxGroup,
		widget.NewSeparator(),
		lengthRow,
		widget.NewSeparator(),
		buttonRow,
		outputEntry,
	)
}

func buildStrengthTab() fyne.CanvasObject {
	// Password input
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password...")

	// Results
	entropyLabel := widget.NewLabel("Entropy: —")
	ratingLabel := widget.NewLabel("Rating: —")
	breachLabel := widget.NewLabel("Breach Status: —")

	strengthBar := widget.NewProgressBar()
	strengthBar.SetValue(0)

	// Check button
	checkBtn := widget.NewButton("Check Strength", func() {
		result := strength.Check(passwordEntry.Text)
		entropyLabel.SetText(fmt.Sprintf("Entropy: %.1f bits", result.Entropy))
		ratingLabel.SetText("Rating: " + result.Rating)
		if result.Breached {
			breachLabel.SetText("Breach Status: Found in breaches!")
		} else {
			breachLabel.SetText("Breach Status: Not found")
		}
		// Normalise entropy to 0-1 for the progress bar (cap at 128 bits)
		bar := result.Entropy / 128
		if bar > 1 {
			bar = 1
		}
		strengthBar.SetValue(bar)
	})
	checkBtn.Importance = widget.HighImportance

	resultsCard := widget.NewCard("Results", "", container.NewVBox(
		entropyLabel,
		ratingLabel,
		breachLabel,
		strengthBar,
	))

	// Order
	return container.NewVBox(
		widget.NewLabel("Enter a password to evaluate:"),
		passwordEntry,
		checkBtn,
		widget.NewSeparator(),
		resultsCard,
	)
}

// intToStr is a tiny helper to avoid importing strconv just for this.
func intToStr(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}
