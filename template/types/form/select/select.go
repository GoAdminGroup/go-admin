package selection

import "fmt"

type Data struct {
	Results    Options    `json:"results"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	More bool `json:"more"`
}

type Options []Option

type Option struct {
	ID       interface{} `json:"id"`
	Text     string      `json:"text"`
	Selected bool        `json:"selected,omitempty"`
	Disabled bool        `json:"disabled,omitempty"`
}

// TODO: make clear

type Configuration struct {
	AdaptContainerCssClass string                 `json:"adaptContainerCssClass,omitempty"`
	AdaptDropdownCssClass  string                 `json:"adaptDropdownCssClass,omitempty"`
	Ajax                   map[string]interface{} `json:"ajax,omitempty"`
	AllowClear             bool                   `json:"allowClear,omitempty"`
	AmdBase                string                 `json:"amdBase,omitempty"`
	AmdLanguageBase        string                 `json:"amdLanguageBase,omitempty"`
	CloseOnSelect          bool                   `json:"closeOnSelect,omitempty"`
	ContainerCss           map[string]interface{} `json:"containerCss,omitempty"`
	ContainerCssClass      string                 `json:"containerCssClass,omitempty"`
	Data                   Options                `json:"data,omitempty"`
	Debug                  bool                   `json:"debug,omitempty"`
	Disabled               bool                   `json:"disabled,omitempty"`

	DropdownAutoWidth bool                   `json:"dropdownAutoWidth,omitempty"`
	DropdownCss       map[string]interface{} `json:"dropdownCss,omitempty"`
	DropdownCssClass  string                 `json:"dropdownCssClass,omitempty"`
	DropdownParent    string                 `json:"dropdownParent,omitempty"`

	EscapeMarkup  func()      `json:"escapeMarkup,omitempty"`
	InitSelection func()      `json:"initSelection,omitempty"`
	Language      interface{} `json:"language,omitempty"`
	Matcher       func()      `json:"matcher,omitempty"`

	MaximumInputLength      int `json:"maximumInputLength,omitempty"`
	MaximumSelectionLength  int `json:"maximumSelectionLength,omitempty"`
	MinimumInputLength      int `json:"minimumInputLength,omitempty"`
	MinimumResultsForSearch int `json:"minimumResultsForSearch,omitempty"`

	Multiple      bool        `json:"multiple,omitempty"`
	Placeholder   interface{} `json:"placeholder,omitempty"`
	Query         func()      `json:"query,omitempty"`
	SelectOnClose bool        `json:"selectOnClose,omitempty"`
	Sorter        func()      `json:"sorter,omitempty"`
	Tags          bool        `json:"tags,omitempty"`

	TemplateResultFns    []Function
	TemplateResult       string `json:"templateResult,omitempty"`
	TemplateSelectionFns []Function
	TemplateSelection    string `json:"templateSelection,omitempty"`

	Theme             string   `json:"theme,omitempty"`
	Tokenizer         func()   `json:"tokenizer,omitempty"`
	TokenSeparators   []string `json:"tokenSeparators,omitempty"`
	Width             string   `json:"width,omitempty"`
	ScrollAfterSelect bool     `json:"scrollAfterSelect,omitempty"`
}

type Function struct {
	Format string
	Args   []Arg
	Next   *Function
	P      func(f string, args []Arg, next *Function) string
}

type ArgType int

const (
	ArgInt ArgType = iota
	ArgString
	ArgOperation
)

type Arg interface {
	Type() ArgType
	String() string
	Wrap(string) string
}

type BaseArg string

func (b BaseArg) String() string {
	return string(b)
}

func (b BaseArg) Wrap(s string) string {
	return s
}

type StringArg BaseArg

func (s StringArg) Type() ArgType {
	return ArgString
}

func (s StringArg) Wrap(ss string) string {
	return `"` + ss + `"`
}

type IntArg BaseArg

func (s IntArg) Type() ArgType {
	return ArgInt
}

type OperationArg BaseArg

func (s OperationArg) Type() ArgType {
	return ArgOperation
}

func If(operation, arg Arg, next *Function) Function {
	return Function{
		Format: `if (%s ` + operation.Wrap("%s") + " " + arg.Wrap("%s") + `) {
	%s
}
`,
		Next: next,
		Args: []Arg{operation, arg},
		P: func(f string, args []Arg, next *Function) string {
			return fmt.Sprintf(f, args[0], args[1], args[2],
				next.P(next.Format, append([]Arg{args[0]}, next.Args...), next.Next))
		},
	}
}

func Return() Function {
	return Function{
		Format: `return %s`,
		P: func(f string, args []Arg, next *Function) string {
			return fmt.Sprintf(f, args[0])
		},
	}
}

func Add(arg Arg) Function {
	return Function{
		Format: `%s += ` + arg.Wrap("%s"),
		Args:   []Arg{arg},
		P: func(f string, args []Arg, next *Function) string {
			return fmt.Sprintf(f, args[0], args[1])
		},
	}
}

func AddFront(arg Arg) Function {
	return Function{
		Format: `%s = ` + arg.Wrap("%s") + ` + %s`,
		Args:   []Arg{arg},
		P: func(f string, args []Arg, next *Function) string {
			return fmt.Sprintf(f, args[0], args[1], args[0])
		},
	}
}
