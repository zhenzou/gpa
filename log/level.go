package log

type level int

func (l level) String() string {
	switch l {
	case LevelDebug:
		return "[D] "
	case LevelInfo:
		return "[I] "
	case LevelWarn:
		return "[W] "
	case LevelError:
		return "[E] "
	}
	return ""
}

const (
	LevelDebug level = iota + 1
	LevelInfo
	LevelWarn
	LevelError
)
