package pkg

import "errors"

var (
	// ErrSliceLength input slice len invalid
	ErrSliceLength = errors.New("invalid slice len")

	// ErrSliceOutOfRange input bytes len out of allow
	ErrSliceOutOfRange = errors.New("the input slice is out of it's range")

	// ErrLoopByRange the value has out of range, recount from range of start or end
	ErrLoopByRange = errors.New("the value has out of range, result maybe parse of loop range")
)

var (
	// ErrIntReturn can got the except value,used in function return with nil
	ErrIntReturn = -1

	// ErrIndexReturn can got the except index,used in function return with nil
	ErrIndexReturn = -1
)
