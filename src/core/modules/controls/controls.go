package controls

type ControlType string

type ResponseTypeControlValue string

const (
	ResponseTypeControlValueNormal ResponseTypeControlValue = "Normal"
	ResponseTypeControlValueEmpty  ResponseTypeControlValue = "Empty"
	ResponseTypeControlValueError  ResponseTypeControlValue = "Error"
)

type Control struct {
	ResponseType ResponseTypeControlValue
}
