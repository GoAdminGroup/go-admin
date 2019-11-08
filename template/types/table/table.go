package table

type Type uint8

const (
	Text Type = iota
	Textarea
	Select
	Date
	Datetime
	Year
	Month
	Day
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
	case Year:
		return "year"
	case Month:
		return "month"
	case Day:
		return "day"
	case Datetime:
		return "datetime"
	default:
		panic("wrong form type")
	}
}
