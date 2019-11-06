package table

type Type uint8

const (
	Text Type = iota
	Textarea
	Select
	Date
	Datetime
	Dateui
	Combodate
	Html5types
	Checklist
	Wysihtml5
	Typeahead
	Typeaheadjs
	Select2
)

func (t Type) String() string {
	switch t {
	case Text:
		return "text"
	case Select:
		return "select"
	case Textarea:
		return "textarea"
	case Date:
		return "date"
	case Datetime:
		return "datetime"
	case Dateui:
		return "dateui"
	case Combodate:
		return "combodate"
	case Html5types:
		return "html5types"
	case Checklist:
		return "checklist"
	case Wysihtml5:
		return "wysihtml5"
	case Typeahead:
		return "typeahead"
	case Typeaheadjs:
		return "typeaheadjs"
	case Select2:
		return "select2"
	default:
		panic("wrong form type")
	}
}
