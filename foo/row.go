package foo

type Row struct {
	Name    string            `json:"name,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
	Columns []string          `json:"columns"`
	Values  [][]interface{}   `json:"values,omitempty"`
	Err     error             `json:"err,omitempty"`
}

type Rows []*Row
