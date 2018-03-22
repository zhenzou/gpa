package gen

import (
	"github.com/zhenzou/gpa/util"
)

type Args []interface{}

type CreateFunc struct {
	Func  *Func
	Table string
}

type UpdateFunc struct {
	Func       *Func
	Table      string
	ParamCount int
	Predicates []*Predicate
}

type FindFunc struct {
	Func       *Func
	Table      string
	Predicates []*Predicate
	ParamCount int
	GroupBy    string
	SortBy     SortExp
}

type DeleteFunc struct {
	Func       *Func
	Table      string
	ParamCount int
	Predicates []*Predicate
}

type SortExp struct {
	Field string
	Desc  bool
}

type OperationInfo struct {
	OperationCode int
	ParamCount    int
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
	opEqual            = []string{"Is", "Equals"} // 默认

	opBetweenInfo          = &OperationInfo{OpBetween, 2}
	opNotNullInfo          = &OperationInfo{OpNotNull, 0}
	opNullInfo             = &OperationInfo{OpNull, 0}
	opLessThanInfo         = &OperationInfo{OpLessThan, 1}
	opLessThanEqualInfo    = &OperationInfo{OpLessThanEqual, 1}
	opGreaterThanInfo      = &OperationInfo{OpGreaterThan, 1}
	opGreaterThanEqualInfo = &OperationInfo{OpGreaterThanEqual, 1}
	opBeforeInfo           = &OperationInfo{OpBefore, 1}
	opAfterInfo            = &OperationInfo{OpAfter, 1}
	opLikeInfo             = &OperationInfo{OpLike, 1}
	opNotLikeInfo          = &OperationInfo{OpNotLike, 1}
	opStartWithInfo        = &OperationInfo{OpStartWith, 1}
	opEndWithInfo          = &OperationInfo{OpEndWith, 1}
	opNotEmptyInfo         = &OperationInfo{OpNotEmpty, 0}
	opEmptyInfo            = &OperationInfo{OpEmpty, 0}
	opContainInfo          = &OperationInfo{OpContain, 1}
	opNotInInfo            = &OperationInfo{OpNotIn, 1}
	opInInfo               = &OperationInfo{OpIn, 1}
	opRegexInfo            = &OperationInfo{OpRegex, 1}
	opExistsInfo           = &OperationInfo{OpExists, 0}
	opTrueInfo             = &OperationInfo{OpTrue, 0}
	opFalseInfo            = &OperationInfo{OpFalse, 0}
	opNotInfo              = &OperationInfo{OpNot, 1}
	opEqualInfo            = &OperationInfo{OpEqual, 1}

	OpCodeMap = map[string]*OperationInfo{
		"IsBetween":          opBetweenInfo, "Between": opBetweenInfo,
		"IsNotNull":          opNotNullInfo, "NotNull": opNotNullInfo,
		"IsNull":             opNullInfo, "Null": opNullInfo,
		"IsLessThan":         opLessThanInfo, "LessThan": opLessThanInfo, "LT": opLessThanInfo,
		"IsLessThanEqual":    opLessThanEqualInfo, "LessThanEqual": opLessThanEqualInfo, "LE": opLessThanEqualInfo,
		"IsGreaterThan":      opGreaterThanInfo, "GreaterThan": opGreaterThanInfo, "GT": opGreaterThanInfo,
		"IsGreaterThanEqual": opGreaterThanEqualInfo, "GreaterThanEqual": opGreaterThanEqualInfo, "GE": opGreaterThanEqualInfo,
		"IsBefore":           opBeforeInfo, "Before": opBeforeInfo,
		"IsAfter":            opAfterInfo, "After": opAfterInfo,
		"IsLike":             opLikeInfo, "Like": opLikeInfo,
		"IsNotLike":          opNotLikeInfo, "NotLike": opNotLikeInfo,
		"IsStartWith":        opStartWithInfo, "StartWith": opStartWithInfo, "HasPrefix": opStartWithInfo,
		"IsEndWith":          opEndWithInfo, "EndWith": opEndWithInfo, "HasSuffix": opEndWithInfo,
		"IsNotEmpty":         opNotEmptyInfo, "NotEmpty": opNotEmptyInfo,
		"IsEmpty":            opEmptyInfo, "Empty": opEmptyInfo,
		"IsContain":          opContainInfo, "Contain": opContainInfo,
		"IsNotIn":            opNotInInfo, "NotIN": opNotInInfo,
		"IsIn":               opInInfo, "IN": opInInfo,
		"MatchesRegex":       opRegexInfo, "Matches": opRegexInfo, "Regex": opRegexInfo,
		"Exists":             opExistsInfo,
		"IsTrue":             opTrueInfo, "True": opTrueInfo,
		"IsFalse":            opFalseInfo, "False": opFalseInfo,
		"IsNot":              opNotInfo, "Not": opNotInfo,
		"Is":                 opEqualInfo, "Equals": opEqualInfo, "": opEqualInfo,
	}

	LogicAnd = "And"
	LogicOr  = "Or"

	SortBy  = "SortBy"
	GroupBy = "GroupBy"

	AllLogic = []string{LogicAnd, LogicOr}

	AllOp = util.Concat(
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
	util.AssertIn([]string{"", LogicAnd, LogicOr}, logic)

	p := &Predicate{Field: field, OpText: op, Logic: logic, ParamCount: 1}
	if info, ok := OpCodeMap[op]; ok {
		p.OpCode = info.OperationCode
		p.ParamCount = info.ParamCount
	}
	return p
}

func IsLogic(str string) bool {
	return util.IsInStringSlice(AllLogic, str)
}

func IsOp(str string) bool {
	return util.IsInStringSlice(AllOp, str)
}
