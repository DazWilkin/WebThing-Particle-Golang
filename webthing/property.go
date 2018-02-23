package webthing

type Property struct {
	Type        string `json:"type,omitempty"`
	Unit        string `json:"unit,omitempty"`
	Description string `json:"description,omitempty"`
	HREF        string `json:"href,omitempty"`
}
