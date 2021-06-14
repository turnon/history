package epoch

import "time"

const (
	epoch = 11644473600
	zoom  = 1000000
)

func From(sec int64, format string) string {
	timing := time.Unix((-epoch + sec/zoom), 0)
	return timing.Format(format)
}

func To(date string, format string) int64 {
	timing, err := time.Parse(format, date)
	if err != nil {
		panic(err)
	}
	return (timing.Unix() + epoch) * zoom
}
