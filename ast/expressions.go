// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package ast

import (
	"io"
	"regexp"

	"github.com/pingcap/parser/format"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/opcode"
)

var (
	_ ExprNode = &BetweenExpr{}
	_ ExprNode = &BinaryOperationExpr{}
	_ ExprNode = &CaseExpr{}
	_ ExprNode = &ColumnNameExpr{}
	_ ExprNode = &TableNameExpr{}
	_ ExprNode = &CompareSubqueryExpr{}
	_ ExprNode = &DefaultExpr{}
	_ ExprNode = &ExistsSubqueryExpr{}
	_ ExprNode = &IsNullExpr{}
	_ ExprNode = &IsTruthExpr{}
	_ ExprNode = &ParenthesesExpr{}
	_ ExprNode = &PatternInExpr{}
	_ ExprNode = &PatternLikeExpr{}
	_ ExprNode = &PatternRegexpExpr{}
	_ ExprNode = &PositionExpr{}
	_ ExprNode = &RowExpr{}
	_ ExprNode = &SubqueryExpr{}
	_ ExprNode = &UnaryOperationExpr{}
	_ ExprNode = &ValuesExpr{}
	_ ExprNode = &VariableExpr{}
	_ ExprNode = &MatchAgainst{}
	_ ExprNode = &SetCollationExpr{}

	_ Node = &ColumnName{}
	_ Node = &WhenClause{}
)

// ValueExpr define a interface for ValueExpr.
type ValueExpr interface {
	ExprNode
	SetValue(val interface{})
	GetValue() interface{}
	GetDatumString() string
	GetString() string
	GetProjectionOffset() int
	SetProjectionOffset(offset int)
}

// NewValueExpr creates a ValueExpr with value, and sets default field type.
var NewValueExpr func(value interface{}, charset string, collate string) ValueExpr

// NewParamMarkerExpr creates a ParamMarkerExpr.
var NewParamMarkerExpr func(offset int) ParamMarkerExpr

// BetweenExpr is for "between and" or "not between and" expression.
type BetweenExpr struct {
	exprNode
	// Expr is the expression to be checked.
	Expr ExprNode
	// Left is the expression for minimal value in the range.
	Left ExprNode
	// Right is the expression for maximum value in the range.
	Right ExprNode
	// Not is true, the expression is "not between and".
	Not bool
}

// Restore implements Node interface.
func (n *BetweenExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *BetweenExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node interface.
func (n *BetweenExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// BinaryOperationExpr is for binary operation like `1 + 1`, `1 - 1`, etc.
type BinaryOperationExpr struct {
	exprNode
	// Op is the operator code for BinaryOperation.
	Op opcode.Op
	// L is the left expression in BinaryOperation.
	L ExprNode
	// R is the right expression in BinaryOperation.
	R ExprNode
}

func restoreBinaryOpWithSpacesAround(ctx *format.RestoreCtx, op opcode.Op) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to annotate, the caller will annotate.

// Restore implements Node interface.
func (n *BinaryOperationExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *BinaryOperationExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node interface.
func (n *BinaryOperationExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// WhenClause is the when clause in Case expression for "when condition then result".
type WhenClause struct {
	node
	// Expr is the condition expression in WhenClause.
	Expr ExprNode
	// Result is the result expression in WhenClause.
	Result ExprNode
}

// Restore implements Node interface.
func (n *WhenClause) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *WhenClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// CaseExpr is the case expression.
type CaseExpr struct {
	exprNode
	// Value is the compare value expression.
	Value ExprNode
	// WhenClauses is the condition check expression.
	WhenClauses []*WhenClause
	// ElseClause is the else result expression.
	ElseClause ExprNode
}

// Restore implements Node interface.
func (n *CaseExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *CaseExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Because the presence of `case when` syntax, `Value` could be nil and we need check this.

// Accept implements Node Accept interface.
func (n *CaseExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SubqueryExpr represents a subquery.
type SubqueryExpr struct {
	exprNode
	// Query is the query SelectNode.
	Query      ResultSetNode
	Evaluated  bool
	Correlated bool
	MultiRows  bool
	Exists     bool
}

func (*SubqueryExpr) resultSet() {
	_ = "STUB: not implemented"

	// Restore implements Node interface.
	return
}

func (n *SubqueryExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *SubqueryExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *SubqueryExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// CompareSubqueryExpr is the expression for "expr cmp (select ...)".
// See https://dev.mysql.com/doc/refman/5.7/en/comparisons-using-subqueries.html
// See https://dev.mysql.com/doc/refman/5.7/en/any-in-some-subqueries.html
// See https://dev.mysql.com/doc/refman/5.7/en/all-subqueries.html
type CompareSubqueryExpr struct {
	exprNode
	// L is the left expression
	L ExprNode
	// Op is the comparison opcode.
	Op opcode.Op
	// R is the subquery for right expression, may be rewritten to other type of expression.
	R ExprNode
	// All is true, we should compare all records in subquery.
	All bool
}

// Restore implements Node interface.
func (n *CompareSubqueryExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *CompareSubqueryExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *CompareSubqueryExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// TableNameExpr represents a table-level object name expression, such as sequence/table/view etc.
type TableNameExpr struct {
	exprNode

	// Name is the referenced object name expression.
	Name *TableName
}

// Restore implements Node interface.
func (n *TableNameExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *TableNameExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *TableNameExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ColumnName represents column name.
type ColumnName struct {
	node
	Schema model.CIStr
	Table  model.CIStr
	Name   model.CIStr
}

// Restore implements Node interface.
func (n *ColumnName) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ColumnName) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// String implements Stringer interface.
func (n *ColumnName) String() string { _ = "STUB: not implemented"; return "" }

// OrigColName returns the full original column name.
func (n *ColumnName) OrigColName() (ret string) { _ = "STUB: not implemented"; return "" }

// ColumnNameExpr represents a column name expression.
type ColumnNameExpr struct {
	exprNode

	// Name is the referenced column name.
	Name *ColumnName

	// Refer is the result field the column name refers to.
	// The value of Refer.Expr is used as the value of the expression.
	Refer *ResultField
}

// Restore implements Node interface.
func (n *ColumnNameExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *ColumnNameExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *ColumnNameExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DefaultExpr is the default expression using default value for a column.
type DefaultExpr struct {
	exprNode
	// Name is the column name.
	Name *ColumnName
}

// Restore implements Node interface.
func (n *DefaultExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *DefaultExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *DefaultExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ExistsSubqueryExpr is the expression for "exists (select ...)".
// See https://dev.mysql.com/doc/refman/5.7/en/exists-and-not-exists-subqueries.html
type ExistsSubqueryExpr struct {
	exprNode
	// Sel is the subquery, may be rewritten to other type of expression.
	Sel ExprNode
	// Not is true, the expression is "not exists".
	Not bool
}

// Restore implements Node interface.
func (n *ExistsSubqueryExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *ExistsSubqueryExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *ExistsSubqueryExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// PatternInExpr is the expression for in operator, like "expr in (1, 2, 3)" or "expr in (select c from t)".
type PatternInExpr struct {
	exprNode
	// Expr is the value expression to be compared.
	Expr ExprNode
	// List is the list expression in compare list.
	List []ExprNode
	// Not is true, the expression is "not in".
	Not bool
	// Sel is the subquery, may be rewritten to other type of expression.
	Sel ExprNode
}

// Restore implements Node interface.
func (n *PatternInExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *PatternInExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *PatternInExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// IsNullExpr is the expression for null check.
type IsNullExpr struct {
	exprNode
	// Expr is the expression to be checked.
	Expr ExprNode
	// Not is true, the expression is "is not null".
	Not bool
}

// Restore implements Node interface.
func (n *IsNullExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *IsNullExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *IsNullExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// IsTruthExpr is the expression for true/false check.
type IsTruthExpr struct {
	exprNode
	// Expr is the expression to be checked.
	Expr ExprNode
	// Not is true, the expression is "is not true/false".
	Not bool
	// True indicates checking true or false.
	True int64
}

// Restore implements Node interface.
func (n *IsTruthExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *IsTruthExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *IsTruthExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// PatternLikeExpr is the expression for like operator, e.g, expr like "%123%"
type PatternLikeExpr struct {
	exprNode
	// Expr is the expression to be checked.
	Expr ExprNode
	// Pattern is the like expression.
	Pattern ExprNode
	// Not is true, the expression is "not like".
	Not bool

	Escape byte

	PatChars []byte
	PatTypes []byte
}

// Restore implements Node interface.
func (n *PatternLikeExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *PatternLikeExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *PatternLikeExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ParamMarkerExpr expression holds a place for another expression.
// Used in parsing prepare statement.
type ParamMarkerExpr interface {
	ValueExpr
	SetOrder(int)
}

// ParenthesesExpr is the parentheses expression.
type ParenthesesExpr struct {
	exprNode
	// Expr is the expression in parentheses.
	Expr ExprNode
}

// Restore implements Node interface.
func (n *ParenthesesExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *ParenthesesExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *ParenthesesExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// PositionExpr is the expression for order by and group by position.
// MySQL use position expression started from 1, it looks a little confused inner.
// maybe later we will use 0 at first.
type PositionExpr struct {
	exprNode
	// N is the position, started from 1 now.
	N int
	// P is the parameterized position.
	P ExprNode
	// Refer is the result field the position refers to.
	Refer *ResultField
}

// Restore implements Node interface.
func (n *PositionExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *PositionExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *PositionExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// PatternRegexpExpr is the pattern expression for pattern match.
type PatternRegexpExpr struct {
	exprNode
	// Expr is the expression to be checked.
	Expr ExprNode
	// Pattern is the expression for pattern.
	Pattern ExprNode
	// Not is true, the expression is "not rlike",
	Not bool

	// Re is the compiled regexp.
	Re *regexp.Regexp
	// Sexpr is the string for Expr expression.
	Sexpr *string
}

// Restore implements Node interface.
func (n *PatternRegexpExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *PatternRegexpExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *PatternRegexpExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RowExpr is the expression for row constructor.
// See https://dev.mysql.com/doc/refman/5.7/en/row-subqueries.html
type RowExpr struct {
	exprNode

	Values []ExprNode
}

// Restore implements Node interface.
func (n *RowExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *RowExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *RowExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// UnaryOperationExpr is the expression for unary operator.
type UnaryOperationExpr struct {
	exprNode
	// Op is the operator opcode.
	Op opcode.Op
	// V is the unary expression.
	V ExprNode
}

// Restore implements Node interface.
func (n *UnaryOperationExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *UnaryOperationExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *UnaryOperationExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ValuesExpr is the expression used in INSERT VALUES.
type ValuesExpr struct {
	exprNode
	// Column is column name.
	Column *ColumnNameExpr
}

// Restore implements Node interface.
func (n *ValuesExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *ValuesExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *ValuesExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// `node` may be *ast.ValueExpr, to avoid panic, we write `_` and do not use
// it.

// VariableExpr is the expression for variable.
type VariableExpr struct {
	exprNode
	// Name is the variable name.
	Name string
	// IsGlobal indicates whether this variable is global.
	IsGlobal bool
	// IsSystem indicates whether this variable is a system variable in current session.
	IsSystem bool
	// ExplicitScope indicates whether this variable scope is set explicitly.
	ExplicitScope bool
	// Value is the variable value.
	Value ExprNode
}

// Restore implements Node interface.
func (n *VariableExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *VariableExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *VariableExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// MaxValueExpr is the expression for "maxvalue" used in partition.
type MaxValueExpr struct {
	exprNode
}

// Restore implements Node interface.
func (n *MaxValueExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Format the ExprNode into a Writer.
func (n *MaxValueExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *MaxValueExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// MatchAgainst is the expression for matching against fulltext index.
type MatchAgainst struct {
	exprNode
	// ColumnNames are the columns to match.
	ColumnNames []*ColumnName
	// Against
	Against ExprNode
	// Modifier
	Modifier FulltextSearchModifier
}

func (n *MatchAgainst) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

func (n *MatchAgainst) Format(w io.Writer) { _ = "STUB: not implemented"; return }

func (n *MatchAgainst) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SetCollationExpr is the expression for the `COLLATE collation_name` clause.
type SetCollationExpr struct {
	exprNode
	// Expr is the expression to be set.
	Expr ExprNode
	// Collate is the name of collation to set.
	Collate string
}

// Restore implements Node interface.
func (n *SetCollationExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the ExprNode into a Writer.
func (n *SetCollationExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *SetCollationExpr) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type exprTextPositionCleaner struct {
	oldTextPos []int
	restore    bool
}

func (e *exprTextPositionCleaner) BeginRestore() { _ = "STUB: not implemented"; return }

func (e *exprTextPositionCleaner) Enter(n Node) (node Node, skipChildren bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (e *exprTextPositionCleaner) Leave(n Node) (node Node, ok bool) {
	_ = "STUB: not implemented"

	// ExpressionDeepEqual compares the equivalence of two expressions.
	return *new(Node), false
}

func ExpressionDeepEqual(a ExprNode, b ExprNode) bool { _ = "STUB: not implemented"; return false }
