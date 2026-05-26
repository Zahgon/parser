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
	"github.com/pingcap/parser/auth"
	"github.com/pingcap/parser/format"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/mysql"
)

var (
	_ DMLNode = &DeleteStmt{}
	_ DMLNode = &InsertStmt{}
	_ DMLNode = &SetOprStmt{}
	_ DMLNode = &UpdateStmt{}
	_ DMLNode = &SelectStmt{}
	_ DMLNode = &CallStmt{}
	_ DMLNode = &ShowStmt{}
	_ DMLNode = &LoadDataStmt{}
	_ DMLNode = &SplitRegionStmt{}

	_ Node = &Assignment{}
	_ Node = &ByItem{}
	_ Node = &FieldList{}
	_ Node = &GroupByClause{}
	_ Node = &HavingClause{}
	_ Node = &AsOfClause{}
	_ Node = &Join{}
	_ Node = &Limit{}
	_ Node = &OnCondition{}
	_ Node = &OrderByClause{}
	_ Node = &SelectField{}
	_ Node = &TableName{}
	_ Node = &TableRefsClause{}
	_ Node = &TableSource{}
	_ Node = &SetOprSelectList{}
	_ Node = &WildCardField{}
	_ Node = &WindowSpec{}
	_ Node = &PartitionByClause{}
	_ Node = &FrameClause{}
	_ Node = &FrameBound{}
)

// JoinType is join type, including cross/left/right/full.
type JoinType int

const (
	// CrossJoin is cross join type.
	CrossJoin JoinType = iota + 1
	// LeftJoin is left Join type.
	LeftJoin
	// RightJoin is right Join type.
	RightJoin
)

// Join represents table join.
type Join struct {
	node

	// Left table can be TableSource or JoinNode.
	Left ResultSetNode
	// Right table can be TableSource or JoinNode or nil.
	Right ResultSetNode
	// Tp represents join type.
	Tp JoinType
	// On represents join on condition.
	On *OnCondition
	// Using represents join using clause.
	Using []*ColumnName
	// NaturalJoin represents join is natural join.
	NaturalJoin bool
	// StraightJoin represents a straight join.
	StraightJoin   bool
	ExplicitParens bool
}

func (*Join) resultSet() {
	_ = "STUB: not implemented"

	// NewCrossJoin builds a cross join without `on` or `using` clause.
	// If the right child is a join tree, we need to handle it differently to make the precedence get right.
	// Here is the example: t1 join t2 join t3
	//
	//	               JOIN ON t2.a = t3.a
	//	t1    join    /    \
	//	            t2      t3
	//
	// (left)         (right)
	//
	// We can not build it directly to:
	//
	//	  JOIN
	//	 /    \
	//	t1	   JOIN ON t2.a = t3.a
	//	      /   \
	//	     t2    t3
	//
	// The precedence would be t1 join (t2 join t3 on t2.a=t3.a), not (t1 join t2) join t3 on t2.a=t3.a
	// We need to find the left-most child of the right child, and build a cross join of the left-hand side
	// of the left child(t1), and the right hand side with the original left-most child of the right child(t2).
	//
	//	    JOIN t2.a = t3.a
	//	   /    \
	//	 JOIN    t3
	//	 /  \
	//	t1  t2
	//
	// Besides, if the right handle side join tree's join type is right join and has explicit parentheses, we need to rewrite it to left join.
	// So t1 join t2 right join t3 would be rewrite to t1 join t3 left join t2.
	// If not, t1 join (t2 right join t3) would be (t1 join t2) right join t3. After rewrite the right join to left join.
	// We get (t1 join t3) left join t2, the semantics is correct.
	return
}

func NewCrossJoin(left, right ResultSetNode) (n *Join) { _ = "STUB: not implemented"; return nil }

// Walk down the right hand side.

// Rewrite right join to left join.

// Restore implements Node interface.
func (n *Join) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *Join) Accept(v Visitor) (Node, bool) { _ = "STUB: not implemented"; return *new(Node), false }

// TableName represents a table name.
type TableName struct {
	node

	Schema model.CIStr
	Name   model.CIStr

	DBInfo    *model.DBInfo
	TableInfo *model.TableInfo

	IndexHints     []*IndexHint
	PartitionNames []model.CIStr
	TableSample    *TableSample
	// AS OF is used to see the data as it was at a specific point in time.
	AsOf *AsOfClause
}

func (*TableName) resultSet() {
	_ = "STUB: not implemented"

	// Restore implements Node interface.
	return
}

func (n *TableName) restoreName(ctx *format.RestoreCtx) { _ = "STUB: not implemented"; return }

// Try CTE, for a CTE table name, we shouldn't write the database name.

func (n *TableName) restorePartitions(ctx *format.RestoreCtx) { _ = "STUB: not implemented"; return }

func (n *TableName) restoreIndexHints(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *TableName) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// IndexHintType is the type for index hint use, ignore or force.
type IndexHintType int

// IndexHintUseType values.
const (
	HintUse IndexHintType = iota + 1
	HintIgnore
	HintForce
)

// IndexHintScope is the type for index hint for join, order by or group by.
type IndexHintScope int

// Index hint scopes.
const (
	HintForScan IndexHintScope = iota + 1
	HintForJoin
	HintForOrderBy
	HintForGroupBy
)

// IndexHint represents a hint for optimizer to use/ignore/force for join/order by/group by.
type IndexHint struct {
	IndexNames []model.CIStr
	HintType   IndexHintType
	HintScope  IndexHintScope
}

// IndexHint Restore (The const field uses switch to facilitate understanding)
func (n *IndexHint) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Prevent accidents

// Prevent accidents

// Accept implements Node Accept interface.
func (n *TableName) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DeleteTableList is the tablelist used in delete statement multi-table mode.
type DeleteTableList struct {
	node
	Tables []*TableName
}

// Restore implements Node interface.
func (n *DeleteTableList) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *DeleteTableList) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// OnCondition represents JOIN on condition.
type OnCondition struct {
	node

	Expr ExprNode
}

// Restore implements Node interface.
func (n *OnCondition) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *OnCondition) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// TableSource represents table source with a name.
type TableSource struct {
	node

	// Source is the source of the data, can be a TableName,
	// a SelectStmt, a SetOprStmt, or a JoinNode.
	Source ResultSetNode

	// AsName is the alias name of the table source.
	AsName model.CIStr
}

func (*TableSource) resultSet() {
	_ = "STUB: not implemented"

	// Restore implements Node interface.
	return
}

func (n *TableSource) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *TableSource) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SelectLockType is the lock type for SelectStmt.
type SelectLockType int

// Select lock types.
const (
	SelectLockNone SelectLockType = iota
	SelectLockForUpdate
	SelectLockForShare
	SelectLockForUpdateNoWait
	SelectLockForUpdateWaitN
	SelectLockForShareNoWait
	SelectLockForUpdateSkipLocked
	SelectLockForShareSkipLocked
)

type SelectLockInfo struct {
	LockType SelectLockType
	WaitSec  uint64
	Tables   []*TableName
}

// String implements fmt.Stringer.
func (n SelectLockType) String() string { _ = "STUB: not implemented"; return "" }

// WildCardField is a special type of select field content.
type WildCardField struct {
	node

	Table  model.CIStr
	Schema model.CIStr
}

// Restore implements Node interface.
func (n *WildCardField) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *WildCardField) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SelectField represents fields in select statement.
// There are two type of select field: wildcard
// and expression with optional alias name.
type SelectField struct {
	node

	// Offset is used to get original text.
	Offset int
	// WildCard is not nil, Expr will be nil.
	WildCard *WildCardField
	// Expr is not nil, WildCard will be nil.
	Expr ExprNode
	// AsName is alias name for Expr.
	AsName model.CIStr
	// Auxiliary stands for if this field is auxiliary.
	// When we add a Field into SelectField list which is used for having/orderby clause but the field is not in select clause,
	// we should set its Auxiliary to true. Then the TrimExec will trim the field.
	Auxiliary bool
}

// Restore implements Node interface.
func (n *SelectField) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *SelectField) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// FieldList represents field list in select statement.
type FieldList struct {
	node

	Fields []*SelectField
}

// Restore implements Node interface.
func (n *FieldList) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *FieldList) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// TableRefsClause represents table references clause in dml statement.
type TableRefsClause struct {
	node

	TableRefs *Join
}

// Restore implements Node interface.
func (n *TableRefsClause) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *TableRefsClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ByItem represents an item in order by or group by.
type ByItem struct {
	node

	Expr      ExprNode
	Desc      bool
	NullOrder bool
}

// Restore implements Node interface.
func (n *ByItem) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ByItem) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// GroupByClause represents group by clause.
type GroupByClause struct {
	node
	Items []*ByItem
}

// Restore implements Node interface.
func (n *GroupByClause) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *GroupByClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// HavingClause represents having clause.
type HavingClause struct {
	node
	Expr ExprNode
}

// Restore implements Node interface.
func (n *HavingClause) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *HavingClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// OrderByClause represents order by clause.
type OrderByClause struct {
	node
	Items    []*ByItem
	ForUnion bool
}

// Restore implements Node interface.
func (n *OrderByClause) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *OrderByClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type SampleMethodType int8

const (
	SampleMethodTypeNone SampleMethodType = iota
	SampleMethodTypeSystem
	SampleMethodTypeBernoulli
	SampleMethodTypeTiDBRegion
)

type SampleClauseUnitType int8

const (
	SampleClauseUnitTypeDefault SampleClauseUnitType = iota
	SampleClauseUnitTypeRow
	SampleClauseUnitTypePercent
)

type TableSample struct {
	node
	SampleMethod     SampleMethodType
	Expr             ExprNode
	SampleClauseUnit SampleClauseUnitType
	RepeatableSeed   ExprNode
}

func (s *TableSample) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

func (s *TableSample) Accept(v Visitor) (node Node, ok bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type SelectStmtKind uint8

const (
	SelectStmtKindSelect SelectStmtKind = iota
	SelectStmtKindTable
	SelectStmtKindValues
)

func (s *SelectStmtKind) String() string { _ = "STUB: not implemented"; return "" }

type CommonTableExpression struct {
	node

	Name        model.CIStr
	Query       *SubqueryExpr
	ColNameList []model.CIStr
}

type WithClause struct {
	node

	IsRecursive bool
	CTEs        []*CommonTableExpression
}

// SelectStmt represents the select query node.
// See https://dev.mysql.com/doc/refman/5.7/en/select.html
type SelectStmt struct {
	dmlNode

	// SelectStmtOpts wraps around select hints and switches.
	*SelectStmtOpts
	// Distinct represents whether the select has distinct option.
	Distinct bool
	// From is the from clause of the query.
	From *TableRefsClause
	// Where is the where clause in select statement.
	Where ExprNode
	// Fields is the select expression list.
	Fields *FieldList
	// GroupBy is the group by expression list.
	GroupBy *GroupByClause
	// Having is the having condition.
	Having *HavingClause
	// WindowSpecs is the window specification list.
	WindowSpecs []WindowSpec
	// OrderBy is the ordering expression list.
	OrderBy *OrderByClause
	// Limit is the limit clause.
	Limit *Limit
	// LockInfo is the lock type
	LockInfo *SelectLockInfo
	// TableHints represents the table level Optimizer Hint for join type
	TableHints []*TableOptimizerHint
	// IsInBraces indicates whether it's a stmt in brace.
	IsInBraces bool
	// WithBeforeBraces indicates whether stmt's with clause is before the brace.
	// It's used to distinguish (with xxx select xxx) and with xxx (select xxx)
	WithBeforeBraces bool
	// QueryBlockOffset indicates the order of this SelectStmt if counted from left to right in the sql text.
	QueryBlockOffset int
	// SelectIntoOpt is the select-into option.
	SelectIntoOpt *SelectIntoOption
	// AfterSetOperator indicates the SelectStmt after which type of set operator
	AfterSetOperator *SetOprType
	// Kind refer to three kind of statement: SelectStmt, TableStmt and ValuesStmt
	Kind SelectStmtKind
	// Lists is filled only when Kind == SelectStmtKindValues
	Lists []*RowExpr
	With  *WithClause
}

func (*SelectStmt) resultSet() { _ = "STUB: not implemented"; return }

func (n *WithClause) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// If the CTE is recursive, we should make it visible for the CTE's query.
// Otherwise, we should put it to stack after building the CTE's query.

func (n *WithClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Restore implements Node interface.
func (n *SelectStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

func restoreTables(ctx *format.RestoreCtx, ts []*TableName) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *SelectStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SetOprSelectList represents the SelectStmt/TableStmt/ValuesStmt list in a union statement.
type SetOprSelectList struct {
	node

	With             *WithClause
	AfterSetOperator *SetOprType
	Selects          []Node
}

// Restore implements Node interface.
func (n *SetOprSelectList) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *SetOprSelectList) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type SetOprType uint8

const (
	Union SetOprType = iota
	UnionAll
	Except
	ExceptAll
	Intersect
	IntersectAll
)

func (s *SetOprType) String() string { _ = "STUB: not implemented"; return "" }

// SetOprStmt represents "union/except/intersect statement"
// See https://dev.mysql.com/doc/refman/5.7/en/union.html
// See https://mariadb.com/kb/en/intersect/
// See https://mariadb.com/kb/en/except/
type SetOprStmt struct {
	dmlNode

	IsInBraces bool
	SelectList *SetOprSelectList
	OrderBy    *OrderByClause
	Limit      *Limit
	With       *WithClause
}

func (*SetOprStmt) resultSet() {
	_ = "STUB: not implemented"

	// Restore implements Node interface.
	return
}

func (n *SetOprStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *SetOprStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Assignment is the expression for assignment, like a = 1.
type Assignment struct {
	node
	// Column is the column name to be assigned.
	Column *ColumnName
	// Expr is the expression assigning to ColName.
	Expr ExprNode
}

// Restore implements Node interface.
func (n *Assignment) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *Assignment) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type ColumnNameOrUserVar struct {
	node
	ColumnName *ColumnName
	UserVar    *VariableExpr
}

func (n *ColumnNameOrUserVar) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *ColumnNameOrUserVar) Accept(v Visitor) (node Node, ok bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// LoadDataStmt is a statement to load data from a specified file, then insert this rows into an existing table.
// See https://dev.mysql.com/doc/refman/5.7/en/load-data.html
type LoadDataStmt struct {
	dmlNode

	IsLocal           bool
	Path              string
	OnDuplicate       OnDuplicateKeyHandlingType
	Table             *TableName
	Columns           []*ColumnName
	FieldsInfo        *FieldsClause
	LinesInfo         *LinesClause
	IgnoreLines       uint64
	ColumnAssignments []*Assignment

	ColumnsAndUserVars []*ColumnNameOrUserVar
}

// Restore implements Node interface.
func (n *LoadDataStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *LoadDataStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

const (
	Terminated = iota
	Enclosed
	Escaped
)

type FieldItem struct {
	Type        int
	Value       string
	OptEnclosed bool
}

// FieldsClause represents fields references clause in load data statement.
type FieldsClause struct {
	Terminated  string
	Enclosed    byte
	Escaped     byte
	OptEnclosed bool
}

// Restore for FieldsClause
func (n *FieldsClause) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// LinesClause represents lines references clause in load data statement.
type LinesClause struct {
	Starting   string
	Terminated string
}

// Restore for LinesClause
func (n *LinesClause) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// CallStmt represents a call procedure query node.
// See https://dev.mysql.com/doc/refman/5.7/en/call.html
type CallStmt struct {
	dmlNode

	Procedure *FuncCallExpr
}

// Restore implements Node interface.
func (n *CallStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *CallStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// InsertStmt is a statement to insert new rows into an existing table.
// See https://dev.mysql.com/doc/refman/5.7/en/insert.html
type InsertStmt struct {
	dmlNode

	IsReplace   bool
	IgnoreErr   bool
	Table       *TableRefsClause
	Columns     []*ColumnName
	Lists       [][]ExprNode
	Setlist     []*Assignment
	Priority    mysql.PriorityEnum
	OnDuplicate []*Assignment
	Select      ResultSetNode
	// TableHints represents the table level Optimizer Hint for join type.
	TableHints     []*TableOptimizerHint
	PartitionNames []model.CIStr
}

// Restore implements Node interface.
func (n *InsertStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *InsertStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DeleteStmt is a statement to delete rows from table.
// See https://dev.mysql.com/doc/refman/5.7/en/delete.html
type DeleteStmt struct {
	dmlNode

	// TableRefs is used in both single table and multiple table delete statement.
	TableRefs *TableRefsClause
	// Tables is only used in multiple table delete statement.
	Tables       *DeleteTableList
	Where        ExprNode
	Order        *OrderByClause
	Limit        *Limit
	Priority     mysql.PriorityEnum
	IgnoreErr    bool
	Quick        bool
	IsMultiTable bool
	BeforeFrom   bool
	// TableHints represents the table level Optimizer Hint for join type.
	TableHints []*TableOptimizerHint
	With       *WithClause
}

// Restore implements Node interface.
func (n *DeleteStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Multiple-Table Syntax

// Single-Table Syntax

// Accept implements Node Accept interface.
func (n *DeleteStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// UpdateStmt is a statement to update columns of existing rows in tables with new values.
// See https://dev.mysql.com/doc/refman/5.7/en/update.html
type UpdateStmt struct {
	dmlNode

	TableRefs     *TableRefsClause
	List          []*Assignment
	Where         ExprNode
	Order         *OrderByClause
	Limit         *Limit
	Priority      mysql.PriorityEnum
	IgnoreErr     bool
	MultipleTable bool
	TableHints    []*TableOptimizerHint
	With          *WithClause
}

// Restore implements Node interface.
func (n *UpdateStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *UpdateStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Limit is the limit clause.
type Limit struct {
	node

	Count  ExprNode
	Offset ExprNode
}

// Restore implements Node interface.
func (n *Limit) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *Limit) Accept(v Visitor) (Node, bool) { _ = "STUB: not implemented"; return *new(Node), false }

// ShowStmtType is the type for SHOW statement.
type ShowStmtType int

// Show statement types.
const (
	ShowNone = iota
	ShowEngines
	ShowDatabases
	ShowTables
	ShowTableStatus
	ShowColumns
	ShowWarnings
	ShowCharset
	ShowVariables
	ShowStatus
	ShowCollation
	ShowCreateTable
	ShowCreateView
	ShowCreateUser
	ShowCreateSequence
	ShowCreatePlacementPolicy
	ShowGrants
	ShowTriggers
	ShowProcedureStatus
	ShowIndex
	ShowProcessList
	ShowCreateDatabase
	ShowConfig
	ShowEvents
	ShowStatsExtended
	ShowStatsMeta
	ShowStatsHistograms
	ShowStatsTopN
	ShowStatsBuckets
	ShowStatsHealthy
	ShowPlugins
	ShowProfile
	ShowProfiles
	ShowMasterStatus
	ShowPrivileges
	ShowErrors
	ShowBindings
	ShowPumpStatus
	ShowDrainerStatus
	ShowOpenTables
	ShowAnalyzeStatus
	ShowRegions
	ShowBuiltins
	ShowTableNextRowId
	ShowBackups
	ShowRestores
	ShowImports
	ShowCreateImport
	ShowPlacement
	ShowPlacementForDatabase
	ShowPlacementForTable
	ShowPlacementForPartition
	ShowPlacementLabels
)

const (
	ProfileTypeInvalid = iota
	ProfileTypeCPU
	ProfileTypeMemory
	ProfileTypeBlockIo
	ProfileTypeContextSwitch
	ProfileTypePageFaults
	ProfileTypeIpc
	ProfileTypeSwaps
	ProfileTypeSource
	ProfileTypeAll
)

// ShowStmt is a statement to provide information about databases, tables, columns and so on.
// See https://dev.mysql.com/doc/refman/5.7/en/show.html
type ShowStmt struct {
	dmlNode

	Tp          ShowStmtType // Databases/Tables/Columns/....
	DBName      string
	Table       *TableName  // Used for showing columns.
	Partition   model.CIStr // Used for showing partition.
	Column      *ColumnName // Used for `desc table column`.
	IndexName   model.CIStr
	Flag        int // Some flag parsed from sql, such as FULL.
	Full        bool
	User        *auth.UserIdentity   // Used for show grants/create user.
	Roles       []*auth.RoleIdentity // Used for show grants .. using
	IfNotExists bool                 // Used for `show create database if not exists`
	Extended    bool                 // Used for `show extended columns from ...`

	// GlobalScope is used by `show variables` and `show bindings`
	GlobalScope bool
	Pattern     *PatternLikeExpr
	Where       ExprNode

	ShowProfileTypes []int  // Used for `SHOW PROFILE` syntax
	ShowProfileArgs  *int64 // Used for `SHOW PROFILE` syntax
	ShowProfileLimit *Limit // Used for `SHOW PROFILE` syntax
}

// Restore implements Node interface.
func (n *ShowStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// FROM OR IN

// ShowTargetFilterable

// here can be INDEX INDEXES KEYS
// FROM or IN

// TODO: remember to check this case
// equivalent to SHOW FIELDS

// FROM or IN

// Accept implements Node Accept interface.
func (n *ShowStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// WindowSpec is the specification of a window.
type WindowSpec struct {
	node

	Name model.CIStr
	// Ref is the reference window of this specification. For example, in `w2 as (w1 order by a)`,
	// the definition of `w2` references `w1`.
	Ref model.CIStr

	PartitionBy *PartitionByClause
	OrderBy     *OrderByClause
	Frame       *FrameClause

	// OnlyAlias will set to true of the first following case.
	// To make compatible with MySQL, we need to distinguish `select func over w` from `select func over (w)`.
	OnlyAlias bool
}

// Restore implements Node interface.
func (n *WindowSpec) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *WindowSpec) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type SelectIntoType int

const (
	SelectIntoOutfile SelectIntoType = iota + 1
	SelectIntoDumpfile
	SelectIntoVars
)

type SelectIntoOption struct {
	node

	Tp         SelectIntoType
	FileName   string
	FieldsInfo *FieldsClause
	LinesInfo  *LinesClause
}

// Restore implements Node interface.
func (n *SelectIntoOption) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// only support SELECT/TABLE/VALUES ... INTO OUTFILE statement now

// Accept implements Node Accept interface.
func (n *SelectIntoOption) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// PartitionByClause represents partition by clause.
type PartitionByClause struct {
	node

	Items []*ByItem
}

// Restore implements Node interface.
func (n *PartitionByClause) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *PartitionByClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// FrameType is the type of window function frame.
type FrameType int

// Window function frame types.
// MySQL only supports `ROWS` and `RANGES`.
const (
	Rows = iota
	Ranges
	Groups
)

// FrameClause represents frame clause.
type FrameClause struct {
	node

	Type   FrameType
	Extent FrameExtent
}

// Restore implements Node interface.
func (n *FrameClause) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *FrameClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// FrameExtent represents frame extent.
type FrameExtent struct {
	Start FrameBound
	End   FrameBound
}

// FrameType is the type of window function frame bound.
type BoundType int

// Frame bound types.
const (
	Following = iota
	Preceding
	CurrentRow
)

// FrameBound represents frame bound.
type FrameBound struct {
	node

	Type      BoundType
	UnBounded bool
	Expr      ExprNode
	// `Unit` is used to indicate the units in which the `Expr` should be interpreted.
	// For example: '2:30' MINUTE_SECOND.
	Unit TimeUnitType
}

// Restore implements Node interface.
func (n *FrameBound) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *FrameBound) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type SplitRegionStmt struct {
	dmlNode

	Table          *TableName
	IndexName      model.CIStr
	PartitionNames []model.CIStr

	SplitSyntaxOpt *SplitSyntaxOption

	SplitOpt *SplitOption
}

type SplitOption struct {
	Lower      []ExprNode
	Upper      []ExprNode
	Num        int64
	ValueLists [][]ExprNode
}

type SplitSyntaxOption struct {
	HasRegionFor bool
	HasPartition bool
}

func (n *SplitRegionStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *SplitRegionStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *SplitOption) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

type FulltextSearchModifier int

const (
	FulltextSearchModifierNaturalLanguageMode = 0
	FulltextSearchModifierBooleanMode         = 1
	FulltextSearchModifierModeMask            = 0xF
	FulltextSearchModifierWithQueryExpansion  = 1 << 4
)

func (m FulltextSearchModifier) IsBooleanMode() bool { _ = "STUB: not implemented"; return false }

func (m FulltextSearchModifier) IsNaturalLanguageMode() bool {
	_ = "STUB: not implemented"
	return false
}

func (m FulltextSearchModifier) WithQueryExpansion() bool { _ = "STUB: not implemented"; return false }

type AsOfClause struct {
	node
	TsExpr ExprNode
}

// Restore implements Node interface.
func (n *AsOfClause) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *AsOfClause) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}
