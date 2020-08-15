package action

type Event string

const (
	EventBlur        Event = "blur"
	EventFocus       Event = "focus"
	EventFocusin     Event = "focusin"
	EventFocusout    Event = "focusout"
	EventLoad        Event = "load"
	EventResize      Event = "resize"
	EventScroll      Event = "scroll"
	EventUnload      Event = "unload"
	EventClick       Event = "click"
	EventDblclick    Event = "dblclick"
	EventMousedown   Event = "mousedown"
	EventMouseup     Event = "mouseup"
	EventMousemove   Event = "mousemove"
	EventMouseover   Event = "mouseover"
	EventMouseout    Event = "mouseout"
	EventMouseenter  Event = "mouseenter"
	EventMouseleave  Event = "mouseleave"
	EventChange      Event = "change"
	EventSelect      Event = "select"
	EventSubmit      Event = "submit"
	EventKeydown     Event = "keydown"
	EventKeypress    Event = "keypress"
	EventKeyup       Event = "keyup"
	EventError       Event = "error"
	EventContextmenu Event = "contextmenu"
)
