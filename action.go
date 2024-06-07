package gotfy

// ActionButtonType specifies all the currently supported action buttons
// that you can use for a notification.
// See: https://docs.ntfy.sh/publish/#action-buttons
type ActionButtonType int8

const (
	ActionButtonTypeUnspecified ActionButtonType = iota
	ActionButtonTypeView
	ActionButtonTypeHTTP
	ActionButtonTypeBroadcast
)

// ActionButton is an interface for Ntfy action buttons.
// All it requires is having a ButtonType and the ability to marshal itself to a JSON structure
// that Ntfy understands; see: https://docs.ntfy.sh/publish/#using-a-json-array
type ActionButton interface {
	ButtonType() ActionButtonType
	MarshalJSON() ([]byte, error)
}
