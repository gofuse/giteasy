package constants

type StatusType int64

const (
	UNSTAGED StatusType = 0
	STAGED   StatusType = 1
	COMMITED StatusType = 2
)
