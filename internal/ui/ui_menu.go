package ui

import (
	"errors"
	"giteasy/internal/lib/git"
	"giteasy/internal/logger"
	"giteasy/internal/model"
	"giteasy/internal/utils"
	"image/color/palette"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func MenuItemFileClone(parent fyne.Canvas, responseBox *widget.Box) *fyne.MenuItem {
	return fyne.NewMenuItem("Clone", func() {
		remoteRepo := widget.NewEntry()
		localRepo := widget.NewEntry()
		userName := widget.NewEntry()
		password := widget.NewPasswordEntry()
		var popup *widget.PopUp
		form := widget.NewForm(
			widget.NewFormItem("Remote:", remoteRepo),
			widget.NewFormItem("Local: ", localRepo),
			widget.NewFormItem("User Name: ", userName),
			widget.NewFormItem("Password: ", password),
		)
		form.OnSubmit = func() {
			popup.Hide()
			logger.Debug(remoteRepo.Text)
			logger.Debug(localRepo.Text)
			model.CurrentProfile.RemoteRepo = remoteRepo.Text
			model.CurrentProfile.LocalRepo = utils.DerriveLocalRepo(remoteRepo.Text, localRepo.Text)
			model.CurrentProfile.UserName = userName.Text
			model.CurrentProfile.Password = password.Text
			err := git.Clone()
			if err != nil {
				logger.Error(err)
				responseBox.Append(canvas.NewText(err.Error(), palette.WebSafe[181]))
				return
			}
			responseBox.Append(canvas.NewText("Clone successful", palette.WebSafe[67]))
		}
		form.OnCancel = func() { popup.Hide() }
		popup = widget.NewModalPopUp(form, parent)
	})
}

func MenuItemFilePull(parent fyne.Window) *fyne.MenuItem {
	return fyne.NewMenuItem("Pull", func() {
		logger.Debug("Pulling")
	})
}

func MenuItemFileOpen(parent fyne.Window) *fyne.MenuItem {
	return fyne.NewMenuItem("Open", func() {
		input := widget.NewEntry()
		input.SetPlaceHolder("Your local project path")
		dialog.ShowCustomConfirm("Git project path", "Ok", "Cancel", input, func(ok bool) {
			if ok {
				model.CurrentProfile.LocalRepo = input.Text
				git.Status()
			}
		}, parent)
	})
}

func MenuItemFileCheck(parent fyne.Window) *fyne.MenuItem {
	return fyne.NewMenuItem("Check", func() {
		if model.CurrentProfile.LocalRepo == "" {
			dialog.NewError(errors.New("No local repo found! Select using File -> Open"), parent).Show()
			return
		}
		git.Status()
	})
}

func MenuItemFileExit(parent fyne.Window) *fyne.MenuItem {
	return fyne.NewMenuItem("Exit", func() {
		parent.Close()
	})
}
