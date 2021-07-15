package models

// TemplateData holds the data send from handlers to template
type TemplateData struct {
	StringMap  map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CSRFTocken string
	Flash      string
	Warning    string
	Error      string
}
