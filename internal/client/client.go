package client

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bufbuild/connect-go"
	shitlistv1 "github.com/mikeder/shitlist/pkg/go/shitlist/v1"
	"github.com/mikeder/shitlist/pkg/go/shitlist/v1/shitlistv1connect"
)

const refreshInterval time.Duration = time.Second * 3

type Application struct {
	cc     shitlistv1connect.ShitlistServiceClient
	app    fyne.App
	window fyne.Window

	leaders []*shitlistv1.Clicker

	state
}

type state struct {
	tabIdx int // current tab index
}

func New() *Application {
	return &Application{
		cc: shitlistv1connect.NewShitlistServiceClient(
			http.DefaultClient,
			"https://click.sqweeb.net",
		),
		app: app.NewWithID("net.sqweeb.click.app"),
	}
}

func (a *Application) setContent() {
	menu := fyne.NewMainMenu(fyne.NewMenu("Options",
		fyne.NewMenuItem("Close", func() {
			a.window.Close()
		}),
	))
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), newHomeTab()),
		container.NewTabItemWithIcon("Leaders", theme.AccountIcon(), newLeaderTab(a.leaders)),
	)
	tabs.SelectIndex(a.state.tabIdx)
	tabs.OnSelected = func(ti *container.TabItem) {
		a.state.tabIdx = tabs.CurrentTabIndex()
	}
	a.window.SetMainMenu(menu)
	a.window.SetContent(tabs)
}

func (a *Application) periodicRefresh() {
	for range time.Tick(refreshInterval) {
		a.getLeaders()
		a.setContent()
		log.Println("refreshed")
	}
}

func (a *Application) Run() {
	a.window = a.app.NewWindow("Clicker Client")
	a.window.Resize(fyne.NewSize(100, 250))

	go a.periodicRefresh()

	a.setContent()
	a.window.ShowAndRun()
}

func newHomeTab() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Home Tab"),
		widget.NewButton("Click", func() {
			log.Println("click button clicked")
		}),
	)
}

func newLeaderTab(leaders []*shitlistv1.Clicker) fyne.CanvasObject {
	header := []fyne.CanvasObject{
		widget.NewLabel("Index"), widget.NewLabel("User ID"), widget.NewLabel("Click Count"),
	}

	var content []fyne.CanvasObject
	content = append(content, header...)
	for i := range leaders {
		row := []fyne.CanvasObject{
			widget.NewLabel(strconv.Itoa(i)),
			widget.NewLabel(leaders[i].UserId),
			widget.NewLabel(strconv.Itoa(int(leaders[i].Clicks))),
		}
		content = append(content, row...)
	}

	return container.NewVScroll(container.NewAdaptiveGrid(3, content...))
}

func (a *Application) getLeaders() {
	log.Println("get leaders")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := a.cc.Leaders(ctx, connect.NewRequest(&shitlistv1.LeadersRequest{}))
	if err != nil {
		log.Printf("get leaders: %v\n", err)
		return
	}
	a.leaders = res.Msg.GetTopClickers()
}
