package webthing

type WebThing struct {
	Name        string              `json:"name,omitempty"`
	Type        string              `json:"type,omitempty"`
	Description string              `json:"description,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
}
