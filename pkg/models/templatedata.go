package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // when you don't know the data type you just put the interface with {}
	CSRFToken string
	Flash     string //success message to user
	Warning   string
	Error     string
}
