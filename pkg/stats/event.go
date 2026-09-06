package stats

import (
	"errors"
	"sync"
)

type EventIndex int

type Level int

const (
	LevelDisabled Level = iota
	LevelBase
	LevelDetailed
)

type Event interface {
	Index() EventIndex
	Level() Level
}

type event struct {
	idx   EventIndex
	level Level
}

func (e event) Index() EventIndex { _ = "STUB: not implemented"; return *new(EventIndex) }

func (e event) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

const (
	_ EventIndex = iota
	serverHandleStart
	serverHandleFinish
	clientConnStart
	clientConnFinish
	rpcStart
	rpcFinish
	readStart
	readFinish
	waitReadStart
	waitReadFinish
	writeStart
	writeFinish
	streamRecv
	streamSend
	checksumGenerateStart
	checksumGenerateFinish
	checksumValidateStart
	checksumValidateFinish
	streamStart
	streamRecvHeader
	streamFinish

	predefinedEventNum
)

var (
	RPCStart  = newEvent(rpcStart, LevelBase)
	RPCFinish = newEvent(rpcFinish, LevelBase)

	ServerHandleStart      = newEvent(serverHandleStart, LevelDetailed)
	ServerHandleFinish     = newEvent(serverHandleFinish, LevelDetailed)
	ClientConnStart        = newEvent(clientConnStart, LevelDetailed)
	ClientConnFinish       = newEvent(clientConnFinish, LevelDetailed)
	ReadStart              = newEvent(readStart, LevelDetailed)
	ReadFinish             = newEvent(readFinish, LevelDetailed)
	WaitReadStart          = newEvent(waitReadStart, LevelDetailed)
	WaitReadFinish         = newEvent(waitReadFinish, LevelDetailed)
	WriteStart             = newEvent(writeStart, LevelDetailed)
	WriteFinish            = newEvent(writeFinish, LevelDetailed)
	ChecksumValidateStart  = newEvent(checksumValidateStart, LevelDetailed)
	ChecksumValidateFinish = newEvent(checksumValidateFinish, LevelDetailed)
	ChecksumGenerateStart  = newEvent(checksumGenerateStart, LevelDetailed)
	ChecksumGenerateFinish = newEvent(checksumGenerateFinish, LevelDetailed)

	StreamRecv       = newEvent(streamRecv, LevelDetailed)
	StreamSend       = newEvent(streamSend, LevelDetailed)
	StreamStart      = newEvent(streamStart, LevelDetailed)
	StreamRecvHeader = newEvent(streamRecvHeader, LevelDetailed)
	StreamFinish     = newEvent(streamFinish, LevelDetailed)
)

var (
	ErrNotAllowed = errors.New("event definition is not allowed after initialization")
	ErrDuplicated = errors.New("event name is already defined")
)

var (
	lock        sync.RWMutex
	inited      int32
	userDefined = make(map[string]Event)
	maxEventNum = int(predefinedEventNum)
)

func FinishInitialization() { _ = "STUB: not implemented"; return }

func DefineNewEvent(name string, level Level) (Event, error) {
	_ = "STUB: not implemented"
	return *new(Event), nil
}

func MaxEventNum() int { _ = "STUB: not implemented"; return 0 }

func PredefinedEventNum() int { _ = "STUB: not implemented"; return 0 }

func newEvent(idx EventIndex, level Level) Event { _ = "STUB: not implemented"; return *new(Event) }
