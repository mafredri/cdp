// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// DataReceivedClient is a client for DataReceived events. Fired when data
// chunk was received over the network.
type DataReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*DataReceivedReply, error)
	rpcc.Stream
}

// DataReceivedReply is the reply for DataReceived events.
type DataReceivedReply struct {
	RequestID         RequestID     `json:"requestId"`         // Request identifier.
	Timestamp         MonotonicTime `json:"timestamp"`         // Timestamp.
	DataLength        int           `json:"dataLength"`        // Data chunk length.
	EncodedDataLength int           `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
}

// EventSourceMessageReceivedClient is a client for EventSourceMessageReceived events.
// Fired when EventSource message is received.
type EventSourceMessageReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*EventSourceMessageReceivedReply, error)
	rpcc.Stream
}

// EventSourceMessageReceivedReply is the reply for EventSourceMessageReceived events.
type EventSourceMessageReceivedReply struct {
	RequestID RequestID     `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime `json:"timestamp"` // Timestamp.
	EventName string        `json:"eventName"` // Message type.
	EventID   string        `json:"eventId"`   // Message identifier.
	Data      string        `json:"data"`      // Message content.
}

// LoadingFailedClient is a client for LoadingFailed events. Fired when HTTP
// request has failed to load.
type LoadingFailedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadingFailedReply, error)
	rpcc.Stream
}

// LoadingFailedReply is the reply for LoadingFailed events.
type LoadingFailedReply struct {
	RequestID       RequestID        `json:"requestId"`                 // Request identifier.
	Timestamp       MonotonicTime    `json:"timestamp"`                 // Timestamp.
	Type            ResourceType     `json:"type"`                      // Resource type.
	ErrorText       string           `json:"errorText"`                 // User friendly error message.
	Canceled        *bool            `json:"canceled,omitempty"`        // True if loading was canceled.
	BlockedReason   BlockedReason    `json:"blockedReason,omitempty"`   // The reason why loading was blocked, if any.
	CORSErrorStatus *CORSErrorStatus `json:"corsErrorStatus,omitempty"` // The reason why loading was blocked by CORS, if any.
}

// LoadingFinishedClient is a client for LoadingFinished events. Fired when
// HTTP request has finished loading.
type LoadingFinishedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadingFinishedReply, error)
	rpcc.Stream
}

// LoadingFinishedReply is the reply for LoadingFinished events.
type LoadingFinishedReply struct {
	RequestID                RequestID     `json:"requestId"`                          // Request identifier.
	Timestamp                MonotonicTime `json:"timestamp"`                          // Timestamp.
	EncodedDataLength        float64       `json:"encodedDataLength"`                  // Total number of bytes received for this request.
	ShouldReportCorbBlocking *bool         `json:"shouldReportCorbBlocking,omitempty"` // Set when 1) response was blocked by Cross-Origin Read Blocking and also 2) this needs to be reported to the DevTools console.
}

// RequestInterceptedClient is a client for RequestIntercepted events. Details
// of an intercepted HTTP request, which must be either allowed, blocked,
// modified or mocked. Deprecated, use Fetch.requestPaused instead.
type RequestInterceptedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestInterceptedReply, error)
	rpcc.Stream
}

// RequestInterceptedReply is the reply for RequestIntercepted events.
type RequestInterceptedReply struct {
	InterceptionID      InterceptionID       `json:"interceptionId"`                // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
	Request             Request              `json:"request"`                       // No description.
	FrameID             internal.PageFrameID `json:"frameId"`                       // The id of the frame that initiated the request.
	ResourceType        ResourceType         `json:"resourceType"`                  // How the requested resource will be used.
	IsNavigationRequest bool                 `json:"isNavigationRequest"`           // Whether this is a navigation request, which can abort the navigation completely.
	IsDownload          *bool                `json:"isDownload,omitempty"`          // Set if the request is a navigation that will result in a download. Only present after response is received from the server (i.e. HeadersReceived stage).
	RedirectURL         *string              `json:"redirectUrl,omitempty"`         // Redirect location, only sent if a redirect was intercepted.
	AuthChallenge       *AuthChallenge       `json:"authChallenge,omitempty"`       // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
	ResponseErrorReason ErrorReason          `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage or if redirect occurred while intercepting request.
	ResponseStatusCode  *int                 `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage or if redirect occurred while intercepting request or auth retry occurred.
	ResponseHeaders     Headers              `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage or if redirect occurred while intercepting request or auth retry occurred.
	RequestID           *RequestID           `json:"requestId,omitempty"`           // If the intercepted request had a corresponding requestWillBeSent event fired for it, then this requestId will be the same as the requestId present in the requestWillBeSent event.
}

// RequestServedFromCacheClient is a client for RequestServedFromCache events.
// Fired if request ended up loading from cache.
type RequestServedFromCacheClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestServedFromCacheReply, error)
	rpcc.Stream
}

// RequestServedFromCacheReply is the reply for RequestServedFromCache events.
type RequestServedFromCacheReply struct {
	RequestID RequestID `json:"requestId"` // Request identifier.
}

// RequestWillBeSentClient is a client for RequestWillBeSent events. Fired
// when page is about to send HTTP request.
type RequestWillBeSentClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestWillBeSentReply, error)
	rpcc.Stream
}

// RequestWillBeSentReply is the reply for RequestWillBeSent events.
type RequestWillBeSentReply struct {
	RequestID        RequestID             `json:"requestId"`                  // Request identifier.
	LoaderID         LoaderID              `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
	DocumentURL      string                `json:"documentURL"`                // URL of the document this request is loaded for.
	Request          Request               `json:"request"`                    // Request data.
	Timestamp        MonotonicTime         `json:"timestamp"`                  // Timestamp.
	WallTime         TimeSinceEpoch        `json:"wallTime"`                   // Timestamp.
	Initiator        Initiator             `json:"initiator"`                  // Request initiator.
	RedirectResponse *Response             `json:"redirectResponse,omitempty"` // Redirect response data.
	Type             ResourceType          `json:"type,omitempty"`             // Type of this resource.
	FrameID          *internal.PageFrameID `json:"frameId,omitempty"`          // Frame identifier.
	HasUserGesture   *bool                 `json:"hasUserGesture,omitempty"`   // Whether the request is initiated by a user gesture. Defaults to false.
}

// ResourceChangedPriorityClient is a client for ResourceChangedPriority events.
// Fired when resource loading priority is changed
type ResourceChangedPriorityClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResourceChangedPriorityReply, error)
	rpcc.Stream
}

// ResourceChangedPriorityReply is the reply for ResourceChangedPriority events.
type ResourceChangedPriorityReply struct {
	RequestID   RequestID        `json:"requestId"`   // Request identifier.
	NewPriority ResourcePriority `json:"newPriority"` // New priority
	Timestamp   MonotonicTime    `json:"timestamp"`   // Timestamp.
}

// SignedExchangeReceivedClient is a client for SignedExchangeReceived events.
// Fired when a signed exchange was received over the network
type SignedExchangeReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*SignedExchangeReceivedReply, error)
	rpcc.Stream
}

// SignedExchangeReceivedReply is the reply for SignedExchangeReceived events.
type SignedExchangeReceivedReply struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Info      SignedExchangeInfo `json:"info"`      // Information about the signed exchange response.
}

// ResponseReceivedClient is a client for ResponseReceived events. Fired when
// HTTP response is available.
type ResponseReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResponseReceivedReply, error)
	rpcc.Stream
}

// ResponseReceivedReply is the reply for ResponseReceived events.
type ResponseReceivedReply struct {
	RequestID RequestID             `json:"requestId"`         // Request identifier.
	LoaderID  LoaderID              `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
	Timestamp MonotonicTime         `json:"timestamp"`         // Timestamp.
	Type      ResourceType          `json:"type"`              // Resource type.
	Response  Response              `json:"response"`          // Response data.
	FrameID   *internal.PageFrameID `json:"frameId,omitempty"` // Frame identifier.
}

// WebSocketClosedClient is a client for WebSocketClosed events. Fired when
// WebSocket is closed.
type WebSocketClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketClosedReply, error)
	rpcc.Stream
}

// WebSocketClosedReply is the reply for WebSocketClosed events.
type WebSocketClosedReply struct {
	RequestID RequestID     `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime `json:"timestamp"` // Timestamp.
}

// WebSocketCreatedClient is a client for WebSocketCreated events. Fired upon
// WebSocket creation.
type WebSocketCreatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketCreatedReply, error)
	rpcc.Stream
}

// WebSocketCreatedReply is the reply for WebSocketCreated events.
type WebSocketCreatedReply struct {
	RequestID RequestID  `json:"requestId"`           // Request identifier.
	URL       string     `json:"url"`                 // WebSocket request URL.
	Initiator *Initiator `json:"initiator,omitempty"` // Request initiator.
}

// WebSocketFrameErrorClient is a client for WebSocketFrameError events. Fired
// when WebSocket message error occurs.
type WebSocketFrameErrorClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameErrorReply, error)
	rpcc.Stream
}

// WebSocketFrameErrorReply is the reply for WebSocketFrameError events.
type WebSocketFrameErrorReply struct {
	RequestID    RequestID     `json:"requestId"`    // Request identifier.
	Timestamp    MonotonicTime `json:"timestamp"`    // Timestamp.
	ErrorMessage string        `json:"errorMessage"` // WebSocket error message.
}

// WebSocketFrameReceivedClient is a client for WebSocketFrameReceived events.
// Fired when WebSocket message is received.
type WebSocketFrameReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameReceivedReply, error)
	rpcc.Stream
}

// WebSocketFrameReceivedReply is the reply for WebSocketFrameReceived events.
type WebSocketFrameReceivedReply struct {
	RequestID RequestID      `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime  `json:"timestamp"` // Timestamp.
	Response  WebSocketFrame `json:"response"`  // WebSocket response data.
}

// WebSocketFrameSentClient is a client for WebSocketFrameSent events. Fired
// when WebSocket message is sent.
type WebSocketFrameSentClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameSentReply, error)
	rpcc.Stream
}

// WebSocketFrameSentReply is the reply for WebSocketFrameSent events.
type WebSocketFrameSentReply struct {
	RequestID RequestID      `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime  `json:"timestamp"` // Timestamp.
	Response  WebSocketFrame `json:"response"`  // WebSocket response data.
}

// WebSocketHandshakeResponseReceivedClient is a client for WebSocketHandshakeResponseReceived events.
// Fired when WebSocket handshake response becomes available.
type WebSocketHandshakeResponseReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketHandshakeResponseReceivedReply, error)
	rpcc.Stream
}

// WebSocketHandshakeResponseReceivedReply is the reply for WebSocketHandshakeResponseReceived events.
type WebSocketHandshakeResponseReceivedReply struct {
	RequestID RequestID         `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime     `json:"timestamp"` // Timestamp.
	Response  WebSocketResponse `json:"response"`  // WebSocket response data.
}

// WebSocketWillSendHandshakeRequestClient is a client for WebSocketWillSendHandshakeRequest events.
// Fired when WebSocket is about to initiate handshake.
type WebSocketWillSendHandshakeRequestClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketWillSendHandshakeRequestReply, error)
	rpcc.Stream
}

// WebSocketWillSendHandshakeRequestReply is the reply for WebSocketWillSendHandshakeRequest events.
type WebSocketWillSendHandshakeRequestReply struct {
	RequestID RequestID        `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime    `json:"timestamp"` // Timestamp.
	WallTime  TimeSinceEpoch   `json:"wallTime"`  // UTC Timestamp.
	Request   WebSocketRequest `json:"request"`   // WebSocket request data.
}

// WebTransportCreatedClient is a client for WebTransportCreated events. Fired
// upon WebTransport creation.
type WebTransportCreatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebTransportCreatedReply, error)
	rpcc.Stream
}

// WebTransportCreatedReply is the reply for WebTransportCreated events.
type WebTransportCreatedReply struct {
	TransportID RequestID     `json:"transportId"`         // WebTransport identifier.
	URL         string        `json:"url"`                 // WebTransport request URL.
	Timestamp   MonotonicTime `json:"timestamp"`           // Timestamp.
	Initiator   *Initiator    `json:"initiator,omitempty"` // Request initiator.
}

// WebTransportConnectionEstablishedClient is a client for WebTransportConnectionEstablished events.
// Fired when WebTransport handshake is finished.
type WebTransportConnectionEstablishedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebTransportConnectionEstablishedReply, error)
	rpcc.Stream
}

// WebTransportConnectionEstablishedReply is the reply for WebTransportConnectionEstablished events.
type WebTransportConnectionEstablishedReply struct {
	TransportID RequestID     `json:"transportId"` // WebTransport identifier.
	Timestamp   MonotonicTime `json:"timestamp"`   // Timestamp.
}

// WebTransportClosedClient is a client for WebTransportClosed events. Fired
// when WebTransport is disposed.
type WebTransportClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebTransportClosedReply, error)
	rpcc.Stream
}

// WebTransportClosedReply is the reply for WebTransportClosed events.
type WebTransportClosedReply struct {
	TransportID RequestID     `json:"transportId"` // WebTransport identifier.
	Timestamp   MonotonicTime `json:"timestamp"`   // Timestamp.
}

// RequestWillBeSentExtraInfoClient is a client for RequestWillBeSentExtraInfo events.
// Fired when additional information about a requestWillBeSent event is
// available from the network stack. Not every requestWillBeSent event will
// have an additional requestWillBeSentExtraInfo fired for it, and there is no
// guarantee whether requestWillBeSent or requestWillBeSentExtraInfo will be
// fired first for the same request.
type RequestWillBeSentExtraInfoClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestWillBeSentExtraInfoReply, error)
	rpcc.Stream
}

// RequestWillBeSentExtraInfoReply is the reply for RequestWillBeSentExtraInfo events.
type RequestWillBeSentExtraInfoReply struct {
	RequestID           RequestID                 `json:"requestId"`                     // Request identifier. Used to match this information to an existing requestWillBeSent event.
	AssociatedCookies   []BlockedCookieWithReason `json:"associatedCookies"`             // A list of cookies potentially associated to the requested URL. This includes both cookies sent with the request and the ones not sent; the latter are distinguished by having blockedReason field set.
	Headers             Headers                   `json:"headers"`                       // Raw request headers as they will be sent over the wire.
	ClientSecurityState *ClientSecurityState      `json:"clientSecurityState,omitempty"` // The client security state set for the request.
}

// ResponseReceivedExtraInfoClient is a client for ResponseReceivedExtraInfo events.
// Fired when additional information about a responseReceived event is
// available from the network stack. Not every responseReceived event will have
// an additional responseReceivedExtraInfo for it, and
// responseReceivedExtraInfo may be fired before or after responseReceived.
type ResponseReceivedExtraInfoClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResponseReceivedExtraInfoReply, error)
	rpcc.Stream
}

// ResponseReceivedExtraInfoReply is the reply for ResponseReceivedExtraInfo events.
type ResponseReceivedExtraInfoReply struct {
	RequestID              RequestID                    `json:"requestId"`              // Request identifier. Used to match this information to another responseReceived event.
	BlockedCookies         []BlockedSetCookieWithReason `json:"blockedCookies"`         // A list of cookies which were not stored from the response along with the corresponding reasons for blocking. The cookies here may not be valid due to syntax errors, which are represented by the invalid cookie line string instead of a proper cookie.
	Headers                Headers                      `json:"headers"`                // Raw response headers as they were received over the wire.
	ResourceIPAddressSpace IPAddressSpace               `json:"resourceIPAddressSpace"` // The IP address space of the resource. The address space can only be determined once the transport established the connection, so we can't send it in `requestWillBeSentExtraInfo`.
	HeadersText            *string                      `json:"headersText,omitempty"`  // Raw response header text as it was received over the wire. The raw text may not always be available, such as in the case of HTTP/2 or QUIC.
}

// TrustTokenOperationDoneClient is a client for TrustTokenOperationDone events.
// Fired exactly once for each Trust Token operation. Depending on the type of
// the operation and whether the operation succeeded or failed, the event is
// fired before the corresponding request was sent or after the response was
// received.
type TrustTokenOperationDoneClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*TrustTokenOperationDoneReply, error)
	rpcc.Stream
}

// TrustTokenOperationDoneReply is the reply for TrustTokenOperationDone events.
type TrustTokenOperationDoneReply struct {
	// Status Detailed success or error status of the operation.
	// 'AlreadyExists' also signifies a successful operation, as the result
	// of the operation already exists und thus, the operation was abort
	// preemptively (e.g. a cache hit).
	//
	// Values: "Ok", "InvalidArgument", "FailedPrecondition", "ResourceExhausted", "AlreadyExists", "Unavailable", "BadResponse", "InternalError", "UnknownError", "FulfilledLocally".
	Status           string                  `json:"status"`
	Type             TrustTokenOperationType `json:"type"`                       // No description.
	RequestID        RequestID               `json:"requestId"`                  // No description.
	TopLevelOrigin   *string                 `json:"topLevelOrigin,omitempty"`   // Top level origin. The context in which the operation was attempted.
	IssuerOrigin     *string                 `json:"issuerOrigin,omitempty"`     // Origin of the issuer in case of a "Issuance" or "Redemption" operation.
	IssuedTokenCount *int                    `json:"issuedTokenCount,omitempty"` // The number of obtained Trust Tokens on a successful "Issuance" operation.
}

// SubresourceWebBundleMetadataReceivedClient is a client for SubresourceWebBundleMetadataReceived events.
// Fired once when parsing the .wbn file has succeeded. The event contains the
// information about the web bundle contents.
type SubresourceWebBundleMetadataReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*SubresourceWebBundleMetadataReceivedReply, error)
	rpcc.Stream
}

// SubresourceWebBundleMetadataReceivedReply is the reply for SubresourceWebBundleMetadataReceived events.
type SubresourceWebBundleMetadataReceivedReply struct {
	RequestID RequestID `json:"requestId"` // Request identifier. Used to match this information to another event.
	URLs      []string  `json:"urls"`      // A list of URLs of resources in the subresource Web Bundle.
}

// SubresourceWebBundleMetadataErrorClient is a client for SubresourceWebBundleMetadataError events.
// Fired once when parsing the .wbn file has failed.
type SubresourceWebBundleMetadataErrorClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*SubresourceWebBundleMetadataErrorReply, error)
	rpcc.Stream
}

// SubresourceWebBundleMetadataErrorReply is the reply for SubresourceWebBundleMetadataError events.
type SubresourceWebBundleMetadataErrorReply struct {
	RequestID    RequestID `json:"requestId"`    // Request identifier. Used to match this information to another event.
	ErrorMessage string    `json:"errorMessage"` // Error message
}

// SubresourceWebBundleInnerResponseParsedClient is a client for SubresourceWebBundleInnerResponseParsed events.
// Fired when handling requests for resources within a .wbn file. Note: this
// will only be fired for resources that are requested by the webpage.
type SubresourceWebBundleInnerResponseParsedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*SubresourceWebBundleInnerResponseParsedReply, error)
	rpcc.Stream
}

// SubresourceWebBundleInnerResponseParsedReply is the reply for SubresourceWebBundleInnerResponseParsed events.
type SubresourceWebBundleInnerResponseParsedReply struct {
	InnerRequestID  RequestID  `json:"innerRequestId"`            // Request identifier of the subresource request
	InnerRequestURL string     `json:"innerRequestURL"`           // URL of the subresource resource.
	BundleRequestID *RequestID `json:"bundleRequestId,omitempty"` // Bundle request identifier. Used to match this information to another event. This made be absent in case when the instrumentation was enabled only after webbundle was parsed.
}

// SubresourceWebBundleInnerResponseErrorClient is a client for SubresourceWebBundleInnerResponseError events.
// Fired when request for resources within a .wbn file failed.
type SubresourceWebBundleInnerResponseErrorClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*SubresourceWebBundleInnerResponseErrorReply, error)
	rpcc.Stream
}

// SubresourceWebBundleInnerResponseErrorReply is the reply for SubresourceWebBundleInnerResponseError events.
type SubresourceWebBundleInnerResponseErrorReply struct {
	InnerRequestID  RequestID  `json:"innerRequestId"`            // Request identifier of the subresource request
	InnerRequestURL string     `json:"innerRequestURL"`           // URL of the subresource resource.
	ErrorMessage    string     `json:"errorMessage"`              // Error message
	BundleRequestID *RequestID `json:"bundleRequestId,omitempty"` // Bundle request identifier. Used to match this information to another event. This made be absent in case when the instrumentation was enabled only after webbundle was parsed.
}
