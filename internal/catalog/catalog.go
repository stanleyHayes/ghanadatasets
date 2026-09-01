package catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Distribution struct {
	Format    string `json:"format"`
	AccessURL string `json:"accessUrl"`
}
type DatasetRecord struct {
	ID              string         `json:"id"`
	Title           string         `json:"title"`
	Publisher       string         `json:"publisher"`
	Topic           string         `json:"topic"`
	Description     string         `json:"description"`
	Licence         string         `json:"licence"`
	AccessCondition string         `json:"accessCondition"`
	SourceID        string         `json:"sourceId"`
	CheckedAt       string         `json:"checkedAt"`
	Status          string         `json:"status"`
	StatusNote      string         `json:"statusNote,omitempty"`
	Distributions   []Distribution `json:"distributions"`
}
type Dataset struct {
	DatasetVersion string          `json:"datasetVersion"`
	Coverage       string          `json:"coverage"`
	ReviewedAt     string          `json:"reviewedAt"`
	Datasets       []DatasetRecord `json:"datasets"`
}
type Query struct{ Search, Publisher, Topic, Format string }

func Parse(raw []byte) (*Dataset, error) {
	var d Dataset
	if err := json.Unmarshal(raw, &d); err != nil {
		return nil, fmt.Errorf("decode dataset: %w", err)
	}
	if d.DatasetVersion == "" || d.Coverage == "" || d.ReviewedAt == "" {
		return nil, errors.New("metadata incomplete")
	}
	if len(d.Datasets) == 0 {
		return nil, errors.New("no records")
	}
	seen := map[string]bool{}
	for _, r := range d.Datasets {
		if r.ID == "" || r.Title == "" || r.Publisher == "" || r.Topic == "" || r.Licence == "" || r.AccessCondition == "" || r.SourceID == "" || r.CheckedAt == "" || r.Status == "" || len(r.Distributions) == 0 {
			return nil, fmt.Errorf("record %q incomplete", r.ID)
		}
		if seen[r.ID] {
			return nil, fmt.Errorf("duplicate id %q", r.ID)
		}
		seen[r.ID] = true
		for _, dist := range r.Distributions {
			if dist.Format == "" || !strings.HasPrefix(dist.AccessURL, "https://") {
				return nil, fmt.Errorf("record %q invalid distribution", r.ID)
			}
		}
	}
	return &d, nil
}
func norm(v string) string {
	return strings.Join(strings.FieldsFunc(strings.ToLower(v), func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }), " ")
}
func (d *Dataset) Search(q Query) []DatasetRecord {
	needle, publisher, topic, format := norm(q.Search), norm(q.Publisher), norm(q.Topic), norm(q.Format)
	out := []DatasetRecord{}
	for _, r := range d.Datasets {
		if publisher != "" && norm(r.Publisher) != publisher {
			continue
		}
		if topic != "" && norm(r.Topic) != topic {
			continue
		}
		if format != "" {
			ok := false
			for _, dist := range r.Distributions {
				ok = ok || norm(dist.Format) == format
			}
			if !ok {
				continue
			}
		}
		if needle != "" && !strings.Contains(norm(r.ID+" "+r.Title+" "+r.Description+" "+r.Topic+" "+r.Publisher), needle) {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Title < out[j].Title })
	return out
}
func (d *Dataset) Get(id string) (DatasetRecord, bool) {
	for _, r := range d.Datasets {
		if r.ID == id {
			return r, true
		}
	}
	return DatasetRecord{}, false
}
