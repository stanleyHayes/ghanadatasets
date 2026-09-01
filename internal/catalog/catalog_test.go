package catalog

import (
	"os"
	"testing"
)

func fixture(t *testing.T) *Dataset {
	t.Helper()
	raw, e := os.ReadFile("../../data/datasets.json")
	if e != nil {
		t.Fatal(e)
	}
	d, e := Parse(raw)
	if e != nil {
		t.Fatal(e)
	}
	return d
}
func TestInvariants(t *testing.T) {
	d := fixture(t)
	if len(d.Datasets) != 10 {
		t.Fatalf("count=%d", len(d.Datasets))
	}
	for _, r := range d.Datasets {
		if len(r.Distributions) == 0 || r.Licence == "" || r.CheckedAt == "" {
			t.Fatalf("incomplete %+v", r)
		}
	}
}
func TestSearchFilters(t *testing.T) {
	d := fixture(t)
	if got := d.Search(Query{Search: "PHC 2021"}); len(got) != 2 {
		t.Fatalf("phc=%d", len(got))
	}
	if got := d.Search(Query{Topic: "Agriculture"}); len(got) != 2 {
		t.Fatalf("agriculture=%d", len(got))
	}
	if got := d.Search(Query{Format: "CATALOG_RECORD"}); len(got) != 1 || got[0].ID != "gss-phc-2021-microdata" {
		t.Fatalf("format=%+v", got)
	}
}
func TestUnknownLicenceIsNotOpen(t *testing.T) {
	d := fixture(t)
	for _, r := range d.Datasets {
		if r.Licence == "unknown" && r.AccessCondition == "open" {
			t.Fatalf("%s mislabels licence", r.ID)
		}
	}
}
