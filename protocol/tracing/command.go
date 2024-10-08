// Code generated by cdpgen. DO NOT EDIT.

package tracing

// GetCategoriesReply represents the return values for GetCategories in the Tracing domain.
type GetCategoriesReply struct {
	Categories []string `json:"categories"` // A list of supported tracing categories.
}

// RecordClockSyncMarkerArgs represents the arguments for RecordClockSyncMarker in the Tracing domain.
type RecordClockSyncMarkerArgs struct {
	SyncID string `json:"syncId"` // The ID of this clock sync marker
}

// NewRecordClockSyncMarkerArgs initializes RecordClockSyncMarkerArgs with the required arguments.
func NewRecordClockSyncMarkerArgs(syncID string) *RecordClockSyncMarkerArgs {
	args := new(RecordClockSyncMarkerArgs)
	args.SyncID = syncID
	return args
}

// RequestMemoryDumpArgs represents the arguments for RequestMemoryDump in the Tracing domain.
type RequestMemoryDumpArgs struct {
	Deterministic *bool                   `json:"deterministic,omitempty"` // Enables more deterministic results by forcing garbage collection
	LevelOfDetail MemoryDumpLevelOfDetail `json:"levelOfDetail,omitempty"` // Specifies level of details in memory dump. Defaults to "detailed".
}

// NewRequestMemoryDumpArgs initializes RequestMemoryDumpArgs with the required arguments.
func NewRequestMemoryDumpArgs() *RequestMemoryDumpArgs {
	args := new(RequestMemoryDumpArgs)

	return args
}

// SetDeterministic sets the Deterministic optional argument. Enables
// more deterministic results by forcing garbage collection
func (a *RequestMemoryDumpArgs) SetDeterministic(deterministic bool) *RequestMemoryDumpArgs {
	a.Deterministic = &deterministic
	return a
}

// SetLevelOfDetail sets the LevelOfDetail optional argument.
// Specifies level of details in memory dump. Defaults to "detailed".
func (a *RequestMemoryDumpArgs) SetLevelOfDetail(levelOfDetail MemoryDumpLevelOfDetail) *RequestMemoryDumpArgs {
	a.LevelOfDetail = levelOfDetail
	return a
}

// RequestMemoryDumpReply represents the return values for RequestMemoryDump in the Tracing domain.
type RequestMemoryDumpReply struct {
	DumpGUID string `json:"dumpGuid"` // GUID of the resulting global memory dump.
	Success  bool   `json:"success"`  // True iff the global memory dump succeeded.
}

// StartArgs represents the arguments for Start in the Tracing domain.
type StartArgs struct {
	// Categories is deprecated.
	//
	// Deprecated: Category/tag filter
	//
	// Note: This property is experimental.
	Categories *string `json:"categories,omitempty"`
	// Options is deprecated.
	//
	// Deprecated: Tracing options
	//
	// Note: This property is experimental.
	Options *string `json:"options,omitempty"`
	// BufferUsageReportingInterval If set, the agent will issue
	// bufferUsage events at this interval, specified in milliseconds
	//
	// Note: This property is experimental.
	BufferUsageReportingInterval *float64 `json:"bufferUsageReportingInterval,omitempty"`
	// TransferMode Whether to report trace events as series of
	// dataCollected events or to save trace to a stream (defaults to
	// `ReportEvents`).
	//
	// Values: "ReportEvents", "ReturnAsStream".
	TransferMode *string      `json:"transferMode,omitempty"`
	StreamFormat StreamFormat `json:"streamFormat,omitempty"` // Trace data format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `json`).
	// StreamCompression Compression format to use. This only applies when
	// using `ReturnAsStream` transfer mode (defaults to `none`)
	//
	// Note: This property is experimental.
	StreamCompression StreamCompression `json:"streamCompression,omitempty"`
	TraceConfig       *TraceConfig      `json:"traceConfig,omitempty"` // No description.
	// PerfettoConfig Base64-encoded serialized
	// perfetto.protos.TraceConfig protobuf message When specified, the
	// parameters `categories`, `options`, `traceConfig` are ignored.
	// (Encoded as a base64 string when passed over JSON)
	//
	// Note: This property is experimental.
	PerfettoConfig []byte `json:"perfettoConfig,omitempty"`
	// TracingBackend Backend type (defaults to `auto`)
	//
	// Note: This property is experimental.
	TracingBackend Backend `json:"tracingBackend,omitempty"`
}

// NewStartArgs initializes StartArgs with the required arguments.
func NewStartArgs() *StartArgs {
	args := new(StartArgs)

	return args
}

// SetCategories sets the Categories optional argument.
//
// Deprecated: Category/tag
// filter
//
// Note: This property is experimental.
func (a *StartArgs) SetCategories(categories string) *StartArgs {
	a.Categories = &categories
	return a
}

// SetOptions sets the Options optional argument.
//
// Deprecated: Tracing options
//
// Note: This property is experimental.
func (a *StartArgs) SetOptions(options string) *StartArgs {
	a.Options = &options
	return a
}

// SetBufferUsageReportingInterval sets the BufferUsageReportingInterval optional argument.
// If set, the agent will issue bufferUsage events at this interval,
// specified in milliseconds
//
// Note: This property is experimental.
func (a *StartArgs) SetBufferUsageReportingInterval(bufferUsageReportingInterval float64) *StartArgs {
	a.BufferUsageReportingInterval = &bufferUsageReportingInterval
	return a
}

// SetTransferMode sets the TransferMode optional argument. Whether to
// report trace events as series of dataCollected events or to save
// trace to a stream (defaults to `ReportEvents`).
//
// Values: "ReportEvents", "ReturnAsStream".
func (a *StartArgs) SetTransferMode(transferMode string) *StartArgs {
	a.TransferMode = &transferMode
	return a
}

// SetStreamFormat sets the StreamFormat optional argument. Trace data
// format to use. This only applies when using `ReturnAsStream`
// transfer mode (defaults to `json`).
func (a *StartArgs) SetStreamFormat(streamFormat StreamFormat) *StartArgs {
	a.StreamFormat = streamFormat
	return a
}

// SetStreamCompression sets the StreamCompression optional argument.
// Compression format to use. This only applies when using
// `ReturnAsStream` transfer mode (defaults to `none`)
//
// Note: This property is experimental.
func (a *StartArgs) SetStreamCompression(streamCompression StreamCompression) *StartArgs {
	a.StreamCompression = streamCompression
	return a
}

// SetTraceConfig sets the TraceConfig optional argument.
func (a *StartArgs) SetTraceConfig(traceConfig TraceConfig) *StartArgs {
	a.TraceConfig = &traceConfig
	return a
}

// SetPerfettoConfig sets the PerfettoConfig optional argument.
// Base64-encoded serialized perfetto.protos.TraceConfig protobuf
// message When specified, the parameters `categories`, `options`,
// `traceConfig` are ignored. (Encoded as a base64 string when passed
// over JSON)
//
// Note: This property is experimental.
func (a *StartArgs) SetPerfettoConfig(perfettoConfig []byte) *StartArgs {
	a.PerfettoConfig = perfettoConfig
	return a
}

// SetTracingBackend sets the TracingBackend optional argument.
// Backend type (defaults to `auto`)
//
// Note: This property is experimental.
func (a *StartArgs) SetTracingBackend(tracingBackend Backend) *StartArgs {
	a.TracingBackend = tracingBackend
	return a
}
