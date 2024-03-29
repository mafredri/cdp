// Code generated by cdpgen. DO NOT EDIT.

package storage

import (
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/rpcc"
)

// CacheStorageContentUpdatedClient is a client for CacheStorageContentUpdated events.
// A cache's contents have been modified.
type CacheStorageContentUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CacheStorageContentUpdatedReply, error)
	rpcc.Stream
}

// CacheStorageContentUpdatedReply is the reply for CacheStorageContentUpdated events.
type CacheStorageContentUpdatedReply struct {
	Origin    string `json:"origin"`    // Origin to update.
	CacheName string `json:"cacheName"` // Name of cache in origin.
}

// CacheStorageListUpdatedClient is a client for CacheStorageListUpdated events.
// A cache has been added/deleted.
type CacheStorageListUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CacheStorageListUpdatedReply, error)
	rpcc.Stream
}

// CacheStorageListUpdatedReply is the reply for CacheStorageListUpdated events.
type CacheStorageListUpdatedReply struct {
	Origin string `json:"origin"` // Origin to update.
}

// IndexedDBContentUpdatedClient is a client for IndexedDBContentUpdated events.
// The origin's IndexedDB object store has been modified.
type IndexedDBContentUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*IndexedDBContentUpdatedReply, error)
	rpcc.Stream
}

// IndexedDBContentUpdatedReply is the reply for IndexedDBContentUpdated events.
type IndexedDBContentUpdatedReply struct {
	Origin          string `json:"origin"`          // Origin to update.
	DatabaseName    string `json:"databaseName"`    // Database to update.
	ObjectStoreName string `json:"objectStoreName"` // ObjectStore to update.
}

// IndexedDBListUpdatedClient is a client for IndexedDBListUpdated events. The
// origin's IndexedDB database list has been modified.
type IndexedDBListUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*IndexedDBListUpdatedReply, error)
	rpcc.Stream
}

// IndexedDBListUpdatedReply is the reply for IndexedDBListUpdated events.
type IndexedDBListUpdatedReply struct {
	Origin string `json:"origin"` // Origin to update.
}

// InterestGroupAccessedClient is a client for InterestGroupAccessed events.
// One of the interest groups was accessed by the associated page.
type InterestGroupAccessedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*InterestGroupAccessedReply, error)
	rpcc.Stream
}

// InterestGroupAccessedReply is the reply for InterestGroupAccessed events.
type InterestGroupAccessedReply struct {
	AccessTime  network.TimeSinceEpoch  `json:"accessTime"`  // No description.
	Type        InterestGroupAccessType `json:"type"`        // No description.
	OwnerOrigin string                  `json:"ownerOrigin"` // No description.
	Name        string                  `json:"name"`        // No description.
}
