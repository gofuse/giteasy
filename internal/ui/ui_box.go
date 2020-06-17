package ui

import (
	"errors"
	"fmt"

	"giteasy/internal/constants"
	"giteasy/internal/lib/git"
	"giteasy/internal/logger"
	"giteasy/internal/model"
	"giteasy/internal/observer"

	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}
func unstagedBox(parent fyne.Window, responseBox *widget.Box) *fyne.Container {
	button := widget.NewButton("Stage", func() {
		if model.CurrentProfile.LocalRepo == "" {
			dialog.NewError(errors.New("No local repo found! Select using File -> Open"), parent).Show()
			return
		}
		git.Stage()
	})
	label := widget.NewLabel("Unstaged files")
	var box = widget.NewVBox(canvas.NewRectangle(color.Transparent))
	toolBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), button, label)

	var container = fyne.Container{
		Objects: []fyne.CanvasObject{fyne.NewContainerWithLayout(layout.NewVBoxLayout(), toolBar, box)},
	}
	me := observer.UnstageObserver{
		OnNotify: func() {
			logger.Debug("UNSTAGED BOX ", model.Get(constants.UNSTAGED))
			box = widget.NewVBox(canvas.NewRectangle(color.Transparent))
			container.Objects = []fyne.CanvasObject{fyne.NewContainerWithLayout(layout.NewVBoxLayout(), toolBar, box)}
			for path, statusCode := range model.Get(constants.UNSTAGED) {
				logger.Debug(path)
				logger.Debug(statusCode)
				box.Append(canvas.NewText(constants.StatusCodeColorMap[statusCode].Status+": "+path, constants.StatusCodeColorMap[statusCode].Color))
			}
			container.Refresh()
			// parent.Refresh()
		},
	}
	observer.Register(constants.UNSTAGED, me)
	return &container
}

func stagedBox(parent fyne.Window, responseBox *widget.Box) *fyne.Container {
	button := widget.NewButton("Commit", func() {
		if model.CurrentProfile.LocalRepo == "" {
			dialog.NewError(errors.New("No local repo found! Select using File -> Open"), parent).Show()
			return
		}
		input := widget.NewEntry()
		dialog.NewCustomConfirm("Commit message", "Ok", "Cancel", input, func(ok bool) {
			if ok {
				git.Commit(input.Text)
			}
		}, parent).Show()
	})
	label := widget.NewLabel("Staged files")
	var box = widget.NewVBox(canvas.NewRectangle(color.Transparent))
	toolBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), button, label)
	var container = fyne.Container{
		Objects: []fyne.CanvasObject{fyne.NewContainerWithLayout(layout.NewVBoxLayout(), toolBar, box)},
	}
	me := observer.UnstageObserver{
		OnNotify: func() {
			logger.Debug("STAGED BOX ", model.Get(constants.STAGED))
			box = widget.NewVBox(canvas.NewRectangle(color.Transparent))
			container.Objects = []fyne.CanvasObject{fyne.NewContainerWithLayout(layout.NewVBoxLayout(), toolBar, box)}
			for path, statusCode := range model.Get(constants.STAGED) {
				logger.Debug(path)
				logger.Debug(statusCode)
				box.Append(canvas.NewText(constants.StatusCodeColorMap[statusCode].Status+": "+path, constants.StatusCodeColorMap[statusCode].Color))
			}
			container.Refresh()
			// parent.Refresh()
		},
	}
	observer.Register(constants.STAGED, me)
	return &container
}

func commitedBox(parent fyne.Window, responseBox *widget.Box) *fyne.Container {
	button := widget.NewButton("Push", func() {
		if model.CurrentProfile.LocalRepo == "" {
			dialog.NewError(errors.New("No local repo found! Select using File -> Open"), parent).Show()
			return
		}
		if model.CurrentProfile.UserName == "" || model.CurrentProfile.Password == "" {
			userName := widget.NewEntry()
			password := widget.NewPasswordEntry()
			var popup *widget.PopUp
			form := widget.NewForm(
				widget.NewFormItem("User Name: ", userName),
				widget.NewFormItem("Password: ", password),
			)
			form.OnSubmit = func() {
				popup.Hide()
				model.CurrentProfile.UserName = userName.Text
				model.CurrentProfile.Password = password.Text
				// defer func() {
				// 	if r := recover(); r != nil {
				// 		logger.Error(r)
				// 		responseBox.Append(canvas.NewText("Failed to push ", color.RGBA{128, 0, 0, 1}))
				// 		return
				// 	}
				// }()
				err := git.Push()
				if err != nil {
					responseBox.Append(canvas.NewText(err.Error(), color.RGBA{128, 0, 0, 1}))
					return
				}
				responseBox.Append(canvas.NewText("Done", color.RGBA{0, 128, 0, 1}))
			}
			form.OnCancel = func() { popup.Hide() }
			popup = widget.NewModalPopUp(form, parent.Canvas())
		}
	})
	label := widget.NewLabel("Commited files")
	var box = widget.NewVBox(canvas.NewRectangle(color.Transparent))
	toolBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), button, label)
	var container = fyne.Container{
		Objects: []fyne.CanvasObject{fyne.NewContainerWithLayout(layout.NewVBoxLayout(), toolBar, box)},
	}
	me := observer.UnstageObserver{
		OnNotify: func() {
			logger.Debug("PUSHED BOX ", model.Get(constants.COMMITED))
			box = widget.NewVBox(canvas.NewRectangle(color.Transparent))
			container.Objects = []fyne.CanvasObject{fyne.NewContainerWithLayout(layout.NewVBoxLayout(), toolBar, box)}
			for path, statusCode := range model.Get(constants.COMMITED) {
				logger.Debug(path)
				logger.Debug(statusCode)
				box.Append(canvas.NewText(constants.StatusCodeColorMap[statusCode].Status+": "+path, constants.StatusCodeColorMap[statusCode].Color))
			}
			container.Refresh()
			// parent.Refresh()
		},
	}
	observer.Register(constants.COMMITED, me)
	return &container
}
