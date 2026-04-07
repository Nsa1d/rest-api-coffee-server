package database

type Record struct {
	ID               int `json:"id"`
	Name             string `json:"name"`
	Price            int `json:"price"`
	InStock          bool   `json:"inStock"`
	ContainsCaffeine bool   `json:"containsCaffeine"`
	Volume           int `json:"volume"`
	Description      string `json:"description"`
}

var records []Record

func SaveRecords(newRecords []Record) error {
	records = make([]Record, len(newRecords))
	copy(records, newRecords)
	return nil
}

func LoadRecords() ([]Record, error) {
	result := make([]Record, len(records))
	copy(result, records)
	return result, nil
}
