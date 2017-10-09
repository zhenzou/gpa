package common

type Args []interface{}

type Create struct {
	Func  *Func
	Table string
}

type Update struct {
	Func       *Func
	Table      string
	Predicates []*Predicate
}

type Find struct {
	Func       *Func
	Table      string
	Predicates []*Predicate

	GroupBy string
	SortBy  Sort
}

type Delete struct {
	Func       *Func
	Table      string
	Predicates []*Predicate
}

type Sort struct {
	Field string
	Desc  bool
}

const (
	OpBetween          = iota + 1
	OpNotNull
	OpNull
	OpLessThan
	OpLessThanEqual
	OpGreaterThan
	OpGreaterThanEqual
	OpBefore
	OpAfter
	OpLike
	OpNotLike
	OpStartWith
	OpEndWith
	OpNotEmpty
	OpEmpty
	OpContain
	OpNotIn
	OpIn
	OpRegex
	OpExists
	OpTrue
	OpFalse
	OpNot
	OpEqual
)

var (
	opBetween          = []string{"IsBetween", "Between"}
	opNotNull          = []string{"IsNotNull", "NotNull"}
	opNull             = []string{"IsNull", "Null"}
	opLessThan         = []string{"IsLessThan", "LessThan", "LT"}
	opLessThanEqual    = []string{"IsLessThanEqual", "LessThanEqual", "LE"}
	opGreaterThan      = []string{"IsGreaterThan", "GreaterThan", "GT"}
	opGreaterThanEqual = []string{"IsGreaterThanEqual", "GreaterThanEqual", "GE"}
	opBefore           = []string{"IsBefore", "Before"}
	opAfter            = []string{"IsAfter", "After"}
	opLike             = []string{"IsLike", "Like"}
	opNotLike          = []string{"IsNotLike", "NotLike"}
	opStartWith        = []string{"IsStartWith", "StartWith", "HasPrefix"}
	opEndWith          = []string{"IsEndWith", "EndWith", "HasSuffix"}
	opNotEmpty         = []string{"IsNotEmpty", "NotEmpty"}
	opEmpty            = []string{"IsEmpty", "Empty"}
	opContain          = []string{"IsContain", "Contain"}
	opNotIn            = []string{"IsNotIn", "NotIN"}
	opIn               = []string{"IsIn", "IN"}
	opRegex            = []string{"MatchesRegex", "Matches", "Regex"}
	opExists           = []string{"Exists"}
	opTrue             = []string{"IsTrue", "True"}
	opFalse            = []string{"IsFalse", "False"}
	opNot              = []string{"IsNot", "Not"}
	opEqual            = []string{"Is", "Equals"} //默认

	LogicAnd = "And"
	LogicOr  = "Or"

	SortBy = "SortBy"

	AllLogic = []string{"And", "Or"}

	AllOp = Concat(
		opBetween,
		opNotNull,
		opNull,
		opLessThan,
		opLessThanEqual,
		opGreaterThan,
		opGreaterThanEqual,
		opBefore,
		opAfter,
		opLike,
		opNotLike,
		opStartWith,
		opEndWith,
		opNotEmpty,
		opEmpty,
		opContain,
		opNotIn,
		opIn,
		opRegex,
		opExists,
		opTrue,
		opFalse,
		opNot,
		opEqual,
	)
)

type Predicate struct {
	Field      string
	OpCode     int
	OpText     string
	Logic      string
	ParamCount int
}

func NewPredicate(field, op, logic string) *Predicate {
	AssertIn([]string{"", LogicAnd, LogicOr}, logic)

	p := &Predicate{Field: field, OpText: op, Logic: logic, ParamCount: 1}
	if IsInStringSlice(opBetween, op) {
		p.OpCode = OpBetween
		p.ParamCount = 2
	} else if IsInStringSlice(opNotNull, op) {
		p.OpCode = OpNotNull
		p.ParamCount = 0
	} else if IsInStringSlice(opNull, op) {
		p.OpCode = OpNull
		p.ParamCount = 0
	} else if IsInStringSlice(opLessThan, op) {
		p.OpCode = OpLessThan
	} else if IsInStringSlice(opLessThanEqual, op) {
		p.OpCode = OpLessThanEqual
	} else if IsInStringSlice(opGreaterThan, op) {
		p.OpCode = OpGreaterThan
	} else if IsInStringSlice(opGreaterThanEqual, op) {
		p.OpCode = OpGreaterThanEqual
	} else if IsInStringSlice(opLike, op) {
		p.OpCode = OpLike
	} else if IsInStringSlice(opNotLike, op) {
		p.OpCode = OpNotLike
	} else if IsInStringSlice(opStartWith, op) {
		p.OpCode = OpStartWith
	} else if IsInStringSlice(opEndWith, op) {
		p.OpCode = OpEndWith
	} else if IsInStringSlice(opNotEmpty, op) {
		p.OpCode = OpNotEmpty
		p.ParamCount = 0
	} else if IsInStringSlice(opEmpty, op) {
		p.OpCode = OpEmpty
		p.ParamCount = 0
	} else if IsInStringSlice(opContain, op) {
		p.OpCode = OpContain
	} else if IsInStringSlice(opIn, op) {
		p.OpCode = OpIn
	} else if IsInStringSlice(opNotIn, op) {
		p.OpCode = OpNotIn
	} else if IsInStringSlice(opRegex, op) {
		p.OpCode = OpRegex
	} else if IsInStringSlice(opExists, op) {
		p.OpCode = OpExists
		p.ParamCount = 0
	} else if IsInStringSlice(opTrue, op) {
		p.OpCode = OpTrue
		p.ParamCount = 0
	} else if IsInStringSlice(opFalse, op) {
		p.OpCode = OpFalse
		p.ParamCount = 0
	} else if IsInStringSlice(opNot, op) {
		p.OpCode = OpNot
	} else if IsInStringSlice(opEqual, op) || op == "" {
		p.OpCode = OpEqual
	}

	return p
}

func IsLogic(str string) bool {
	return IsInStringSlice(AllLogic, str)
}

func IsOp(str string) bool {
	return IsInStringSlice(AllOp, str)
}
