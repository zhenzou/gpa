package common

type Args []interface{}

type Create struct {
	Func  *Func
	Table string
}

type Update struct {
	Func       *Func
	Table      string
	ParamCount int
	Predicates []*Predicate
}

type Find struct {
	Func       *Func
	Table      string
	Predicates []*Predicate
	ParamCount int
	GroupBy    string
	SortBy     Sort
}

type Delete struct {
	Func       *Func
	Table      string
	ParamCount int
	Predicates []*Predicate
}

type Sort struct {
	Field string
	Desc  bool
}

type OpInfo struct {
	OpCode     int
	ParamCount int
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

	opBetweenInfo          = &OpInfo{OpBetween, 2}
	opNotNullInfo          = &OpInfo{OpNotNull, 0}
	opNullInfo             = &OpInfo{OpNull, 0}
	opLessThanInfo         = &OpInfo{OpLessThan, 1}
	opLessThanEqualInfo    = &OpInfo{OpLessThanEqual, 1}
	opGreaterThanInfo      = &OpInfo{OpGreaterThan, 1}
	opGreaterThanEqualInfo = &OpInfo{OpGreaterThanEqual, 1}
	opBeforeInfo           = &OpInfo{OpBefore, 1}
	opAfterInfo            = &OpInfo{OpAfter, 1}
	opLikeInfo             = &OpInfo{OpLike, 1}
	opNotLikeInfo          = &OpInfo{OpNotLike, 1}
	opStartWithInfo        = &OpInfo{OpStartWith, 1}
	opEndWithInfo          = &OpInfo{OpEndWith, 1}
	opNotEmptyInfo         = &OpInfo{OpNotEmpty, 0}
	opEmptyInfo            = &OpInfo{OpEmpty, 0}
	opContainInfo          = &OpInfo{OpContain, 1}
	opNotInInfo            = &OpInfo{OpNotIn, 1}
	opInInfo               = &OpInfo{OpIn, 1}
	opRegexInfo            = &OpInfo{OpRegex, 1}
	opExistsInfo           = &OpInfo{OpExists, 0}
	opTrueInfo             = &OpInfo{OpTrue, 0}
	opFalseInfo            = &OpInfo{OpFalse, 0}
	opNotInfo              = &OpInfo{OpNot, 1}
	opEqualInfo            = &OpInfo{OpEqual, 1}

	OpCodeMap = map[string]*OpInfo{
		"IsBetween":          opBetweenInfo,            "Between": opBetweenInfo,
		"IsNotNull":          opNotNullInfo,            "NotNull": opNotNullInfo,
		"IsNull":             opNullInfo,               "Null": opNullInfo,
		"IsLessThan":         opLessThanInfo,           "LessThan": opLessThanInfo, 				"LT": opLessThanInfo,
		"IsLessThanEqual":    opLessThanEqualInfo,      "LessThanEqual": opLessThanEqualInfo, 		"LE": opLessThanEqualInfo,
		"IsGreaterThan":      opGreaterThanInfo,        "GreaterThan": opGreaterThanInfo, 			"GT": opGreaterThanInfo,
		"IsGreaterThanEqual": opGreaterThanEqualInfo,   "GreaterThanEqual": opGreaterThanEqualInfo, "GE": opGreaterThanEqualInfo,
		"IsBefore":           opBeforeInfo,             "Before": opBeforeInfo,
		"IsAfter":            opAfterInfo,              "After": opAfterInfo,
		"IsLike":             opLikeInfo, 				"Like": opLikeInfo,
		"IsNotLike":          opNotLikeInfo, 			"NotLike": opNotLikeInfo,
		"IsStartWith":        opStartWithInfo, 			"StartWith": opStartWithInfo, 				"HasPrefix": opStartWithInfo,
		"IsEndWith":          opEndWithInfo, 			"EndWith": opEndWithInfo, 					"HasSuffix": opEndWithInfo,
		"IsNotEmpty":         opNotEmptyInfo, 			"NotEmpty": opNotEmptyInfo,
		"IsEmpty":            opEmptyInfo, 				"Empty": opEmptyInfo,
		"IsContain":          opContainInfo, 			"Contain": opContainInfo,
		"IsNotIn":            opNotInInfo, 				"NotIN": opNotInInfo,
		"IsIn":               opInInfo, 				"IN": opInInfo,
		"MatchesRegex":       opRegexInfo, 				"Matches": opRegexInfo, 					"Regex": opRegexInfo,
		"Exists":             opExistsInfo,
		"IsTrue":             opTrueInfo, 				"True": opTrueInfo,
		"IsFalse":            opFalseInfo, 				"False": opFalseInfo,
		"IsNot":              opNotInfo, 				"Not": opNotInfo,
		"Is":                 opEqualInfo, 				"Equals": opEqualInfo, 						"": opEqualInfo,
	}

	LogicAnd = "And"
	LogicOr  = "Or"

	SortBy  = "SortBy"
	GroupBy = "GroupBy"

	AllLogic = []string{LogicAnd, LogicOr}

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
	if info, ok := OpCodeMap[op]; ok {
		p.OpCode = info.OpCode
		p.ParamCount = info.ParamCount
	}
	return p
}

func IsLogic(str string) bool {
	return IsInStringSlice(AllLogic, str)
}

func IsOp(str string) bool {
	return IsInStringSlice(AllOp, str)
}
