package runtime

import (
	"bytes"
	"fmt"
)

// Error implements error for ExceptionDetails.
func (r ExceptionDetails) Error() string {
	var desc string
	if r.Exception.Description != nil {
		desc = ": " + *r.Exception.Description
	}
	return fmt.Sprintf("runtime.ExceptionDetails: %s exception at %d:%d%s", r.Text, r.LineNumber, r.ColumnNumber, desc)
}

var (
	_ error = (*ExceptionDetails)(nil)
)

// String returns a human readable string of a runtime object.
func (r RemoteObject) String() string {
	switch r.Type {
	case "undefined":
		return "undefined"
	case "object":
		switch {
		case r.Preview != nil:
			return r.Preview.String()
		}
	default:
		if r.UnserializableValue.Valid() {
			return r.UnserializableValue.String()
		}
	}

	if len(r.Value) == 0 && r.Description != nil {
		return *r.Description
	}

	return string(r.Value)
}

// String returns a human readable string of the object preview.
func (r ObjectPreview) String() string {
	var desc string
	if r.Description != nil {
		desc = *r.Description
	}

	var b bytes.Buffer

	switch r.Type {
	case "object":
		var stype string
		if r.Subtype != nil {
			stype = *r.Subtype
		}
		switch stype {
		case "null":
			return "null"
		case "array", "typedarray":
			if desc != "" {
				b.WriteString(desc)
				b.WriteByte(' ')
			}
			b.WriteByte('[')
			for _, prop := range r.Properties {
				b.WriteString(prop.string(false))
				b.WriteString(", ")
			}
			if b.Len() >= 2 && len(r.Properties) > 0 {
				b.Truncate(b.Len() - 2)
			}
			b.WriteByte(']')
			return b.String()
		case "date", "regexp":
			return desc
		default:
			if val, ok := primitiveValue(r.Properties); ok {
				fmt.Fprintf(&b, "%s(%s)", desc, val)
				return b.String()
			}
		}

		b.WriteString(desc)
		b.WriteString(" {")
		for _, prop := range r.Properties {
			b.WriteString(prop.String())
			b.WriteString(", ")
		}
		for _, entry := range r.Entries {
			b.WriteString(entry.String())
			b.WriteString(", ")
		}

		if r.Overflow {
			b.WriteString("...")
		} else if b.Len() >= 2 && (len(r.Properties) > 0 || len(r.Entries) > 0) {
			b.Truncate(b.Len() - 2)
		}

		b.WriteByte('}')
		return b.String()
	case "string":
		fmt.Fprintf(&b, "%q", desc)
		return b.String()
	default:
		return desc
	}
}

// String returns a human readable string of the property.
func (r PropertyPreview) String() string {
	return r.string(true)
}

func (r PropertyPreview) string(showName bool) string {
	var b bytes.Buffer
	if showName {
		b.WriteString(r.Name)
		b.WriteString(": ")
	}
	if r.Value != nil {
		if r.Type == "string" {
			fmt.Fprintf(&b, "%q", *r.Value)
		} else {
			b.WriteString(*r.Value)
		}
	}
	if r.ValuePreview != nil {
		b.WriteString(r.ValuePreview.String())
	}
	return b.String()
}

// String returns a human readable string of the entry preview.
func (r EntryPreview) String() string {
	var b bytes.Buffer
	if r.Key != nil {
		b.WriteString(r.Key.String())
		b.WriteString(" => ")
	}
	b.WriteString(r.Value.String())
	return b.String()
}

const primitiveValueKey = "[[PrimitiveValue]]"

func primitiveValue(props []PropertyPreview) (string, bool) {
	for _, prop := range props {
		if prop.Name == primitiveValueKey && prop.Value != nil {
			val := *prop.Value
			if prop.Type == "string" {
				val = fmt.Sprintf("%q", val)
			}
			return val, true
		}
	}
	return "", false
}
