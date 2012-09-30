// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type MainWindow struct {
	Widget   **walk.MainWindow
	Name     string
	Font     Font
	Title    string
	Layout   Layout
	Children []Widget
	Menu     Menu
	ToolBar  ToolBar
}

func (mw MainWindow) Create(parent walk.Container) error {
	w, err := walk.NewMainWindow()
	if err != nil {
		return err
	}

	return InitWidget(mw, w, func() error {
		if err := w.SetTitle(mw.Title); err != nil {
			return err
		}

		if err := mw.Menu.init(w.Menu()); err != nil {
			return err
		}

		imageList, err := walk.NewImageList(walk.Size{16, 16}, 0)
		if err != nil {
			return err
		}
		w.ToolBar().SetImageList(imageList)

		if err := mw.ToolBar.initActions(w.ToolBar()); err != nil {
			return err
		}

		if mw.Widget != nil {
			*mw.Widget = w
		}

		return nil
	})
}

func (mw MainWindow) CommonInfo() (name string, stretchFactor, row, rowSpan, column, columnSpan int) {
	return mw.Name, 0, 0, 0, 0, 0
}

func (mw MainWindow) Font_() *Font {
	return &mw.Font
}

func (mw MainWindow) ContainerInfo() (Layout, []Widget) {
	return mw.Layout, mw.Children
}
