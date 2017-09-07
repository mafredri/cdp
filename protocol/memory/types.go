// Code generated by cdpgen. DO NOT EDIT.

package memory

// PressureLevel Memory pressure level.
type PressureLevel string

// PressureLevel as enums.
const (
	PressureLevelNotSet   PressureLevel = ""
	PressureLevelModerate PressureLevel = "moderate"
	PressureLevelCritical PressureLevel = "critical"
)

func (e PressureLevel) Valid() bool {
	switch e {
	case "moderate", "critical":
		return true
	default:
		return false
	}
}

func (e PressureLevel) String() string {
	return string(e)
}
