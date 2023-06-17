package gprogress

/*
   @File: progress.go
   @Author: khaosles
   @Time: 2023/6/4 01:27
   @Desc:
*/

type EventType int

const (
	// TransferStartedEvent transfer started, set TotalBytes
	TransferStartedEvent EventType = 1 + iota
	// TransferDataEvent transfer data, set ConsumedBytes anmd TotalBytes
	TransferDataEvent
	// TransferCompletedEvent transfer completed
	TransferCompletedEvent
	// TransferFailedEvent transfer encounters an error
	TransferFailedEvent
)

// Event defines progress event
type Event struct {
	ConsumedBytes int64
	TotalBytes    int64
	RwBytes       int64
	EventType     EventType
}

// Listener listens progress change
type Listener interface {
	ProgressChanged(event *Event)
}
