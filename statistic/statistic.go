package statistic

import (
	"io"
	"sort"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/turnon/history/db"
)

type visitsPerDate struct {
	vpd   map[string]map[string]int
	where func(*db.Visit) string
	sum   func(*db.Visit) int
	dates []string
}

const maxDateCount = 366

func NewVisitsPerDate(w func(*db.Visit) string, s func(*db.Visit) int, start string, end string) *visitsPerDate {
	dates := []string{}
	date := start
	t, _ := time.Parse(db.EpochFormat, start)
	for date <= end && len(dates) <= maxDateCount {
		dates = append(dates, date)
		t = t.AddDate(0, 0, 1)
		date = t.Format(db.EpochFormat)
	}

	_vpd := &visitsPerDate{
		vpd:   make(map[string]map[string]int),
		where: w,
		sum:   s,
		dates: dates,
	}
	return _vpd
}

func (vpd *visitsPerDate) AddVisits(vs []*db.Visit) {
	for _, v := range vs {
		vpd.addVisit(v)
	}
}

func (vpd *visitsPerDate) addVisit(v *db.Visit) {
	where := vpd.where(v)
	sum := vpd.sum(v)

	visistWhere, exists := vpd.vpd[where]
	if !exists {
		visistWhere = make(map[string]int)
		vpd.vpd[where] = visistWhere
	}
	visistWhere[v.VisitTimeString] = visistWhere[v.VisitTimeString] + sum
}

func (vpd *visitsPerDate) Render(w io.Writer) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithLegendOpts(opts.Legend{Show: true, Top: "5%"}),
		charts.WithInitializationOpts(opts.Initialization{Width: "1800px", Height: "800px"}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "slider"}),
	)

	line.SetXAxis(vpd.dates)

	for _, category := range vpd.categories() {
		visits := vpd.vpd[category]
		items := make([]opts.LineData, 0)
		for _, date := range vpd.dates {
			sum := visits[date]
			items = append(items, opts.LineData{Value: sum})
		}

		line.AddSeries(category, items)
	}

	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	line.Render(w)
}

func (vpd *visitsPerDate) categories() []string {
	cates := make([]string, 0, len(vpd.vpd))
	for category, _ := range vpd.vpd {
		cates = append(cates, category)

	}
	sort.Strings(cates)
	return cates
}
