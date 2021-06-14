package epoch

import "testing"

func TestEpoch(t *testing.T) {
	format := "2006-01-02"
	expectedDate := "2021-06-14"
	sec := To(expectedDate, format)
	actualDate := From(sec, format)
	if expectedDate != actualDate {
		t.Error(actualDate)
	}
}
