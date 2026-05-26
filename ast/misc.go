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
	_ StmtNode = &AdminStmt{}
	_ StmtNode = &AlterUserStmt{}
	_ StmtNode = &BeginStmt{}
	_ StmtNode = &BinlogStmt{}
	_ StmtNode = &CommitStmt{}
	_ StmtNode = &CreateUserStmt{}
	_ StmtNode = &DeallocateStmt{}
	_ StmtNode = &DoStmt{}
	_ StmtNode = &ExecuteStmt{}
	_ StmtNode = &ExplainStmt{}
	_ StmtNode = &GrantStmt{}
	_ StmtNode = &PrepareStmt{}
	_ StmtNode = &RollbackStmt{}
	_ StmtNode = &SetPwdStmt{}
	_ StmtNode = &SetRoleStmt{}
	_ StmtNode = &SetDefaultRoleStmt{}
	_ StmtNode = &SetStmt{}
	_ StmtNode = &UseStmt{}
	_ StmtNode = &FlushStmt{}
	_ StmtNode = &KillStmt{}
	_ StmtNode = &CreateBindingStmt{}
	_ StmtNode = &DropBindingStmt{}
	_ StmtNode = &ShutdownStmt{}
	_ StmtNode = &RestartStmt{}
	_ StmtNode = &RenameUserStmt{}
	_ StmtNode = &HelpStmt{}
	_ StmtNode = &PlanRecreatorStmt{}

	_ Node = &PrivElem{}
	_ Node = &VariableAssignment{}
)

// Isolation level constants.
const (
	ReadCommitted   = "READ-COMMITTED"
	ReadUncommitted = "READ-UNCOMMITTED"
	Serializable    = "SERIALIZABLE"
	RepeatableRead  = "REPEATABLE-READ"

	PumpType    = "PUMP"
	DrainerType = "DRAINER"
)

// Transaction mode constants.
const (
	Optimistic  = "OPTIMISTIC"
	Pessimistic = "PESSIMISTIC"
)

// TypeOpt is used for parsing data type option from SQL.
type TypeOpt struct {
	IsUnsigned bool
	IsZerofill bool
}

// FloatOpt is used for parsing floating-point type option from SQL.
// See http://dev.mysql.com/doc/refman/5.7/en/floating-point-types.html
type FloatOpt struct {
	Flen    int
	Decimal int
}

// AuthOption is used for parsing create use statement.
type AuthOption struct {
	// ByAuthString set as true, if AuthString is used for authorization. Otherwise, authorization is done by HashString.
	ByAuthString bool
	AuthString   string
	HashString   string
	AuthPlugin   string
}

// Restore implements Node interface.
func (n *AuthOption) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// TraceStmt is a statement to trace what sql actually does at background.
type TraceStmt struct {
	stmtNode

	Stmt   StmtNode
	Format string
}

// Restore implements Node interface.
func (n *TraceStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *TraceStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ExplainForStmt is a statement to provite information about how is SQL statement executeing
// in connection #ConnectionID
// See https://dev.mysql.com/doc/refman/5.7/en/explain.html
type ExplainForStmt struct {
	stmtNode

	Format       string
	ConnectionID uint64
}

// Restore implements Node interface.
func (n *ExplainForStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *ExplainForStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ExplainStmt is a statement to provide information about how is SQL statement executed
// or get columns information in a table.
// See https://dev.mysql.com/doc/refman/5.7/en/explain.html
type ExplainStmt struct {
	stmtNode

	Stmt    StmtNode
	Format  string
	Analyze bool
}

// Restore implements Node interface.
func (n *ExplainStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ExplainStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// PlanRecreatorStmt is a statement to dump or load information for recreating plans
type PlanRecreatorStmt struct {
	stmtNode

	Stmt    StmtNode
	Analyze bool
	Load    bool
	File    string
	// Where is the where clause in select statement.
	Where ExprNode
	// OrderBy is the ordering expression list.
	OrderBy *OrderByClause
	// Limit is the limit clause.
	Limit *Limit
}

// Restore implements Node interface.
func (n *PlanRecreatorStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *PlanRecreatorStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// PrepareStmt is a statement to prepares a SQL statement which contains placeholders,
// and it is executed with ExecuteStmt and released with DeallocateStmt.
// See https://dev.mysql.com/doc/refman/5.7/en/prepare.html
type PrepareStmt struct {
	stmtNode

	Name    string
	SQLText string
	SQLVar  *VariableExpr
}

// Restore implements Node interface.
func (n *PrepareStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *PrepareStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DeallocateStmt is a statement to release PreparedStmt.
// See https://dev.mysql.com/doc/refman/5.7/en/deallocate-prepare.html
type DeallocateStmt struct {
	stmtNode

	Name string
}

// Restore implements Node interface.
func (n *DeallocateStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *DeallocateStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Prepared represents a prepared statement.
type Prepared struct {
	Stmt          StmtNode
	StmtType      string
	Params        []ParamMarkerExpr
	SchemaVersion int64
	UseCache      bool
	CachedPlan    interface{}
	CachedNames   interface{}
}

// ExecuteStmt is a statement to execute PreparedStmt.
// See https://dev.mysql.com/doc/refman/5.7/en/execute.html
type ExecuteStmt struct {
	stmtNode

	Name       string
	UsingVars  []ExprNode
	BinaryArgs interface{}
	ExecID     uint32
	IdxInMulti int
}

// Restore implements Node interface.
func (n *ExecuteStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ExecuteStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// BeginStmt is a statement to start a new transaction.
// See https://dev.mysql.com/doc/refman/5.7/en/commit.html
type BeginStmt struct {
	stmtNode
	Mode                  string
	CausalConsistencyOnly bool
	ReadOnly              bool
	// AS OF is used to read the data at a specific point of time.
	// Should only be used when ReadOnly is true.
	AsOf *AsOfClause
}

// Restore implements Node interface.
func (n *BeginStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *BeginStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// BinlogStmt is an internal-use statement.
// We just parse and ignore it.
// See http://dev.mysql.com/doc/refman/5.7/en/binlog.html
type BinlogStmt struct {
	stmtNode
	Str string
}

// Restore implements Node interface.
func (n *BinlogStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *BinlogStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// CompletionType defines completion_type used in COMMIT and ROLLBACK statements
type CompletionType int8

const (
	// CompletionTypeDefault refers to NO_CHAIN
	CompletionTypeDefault CompletionType = iota
	CompletionTypeChain
	CompletionTypeRelease
)

func (n CompletionType) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// CommitStmt is a statement to commit the current transaction.
// See https://dev.mysql.com/doc/refman/5.7/en/commit.html
type CommitStmt struct {
	stmtNode
	// CompletionType overwrites system variable `completion_type` within transaction
	CompletionType CompletionType
}

// Restore implements Node interface.
func (n *CommitStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *CommitStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RollbackStmt is a statement to roll back the current transaction.
// See https://dev.mysql.com/doc/refman/5.7/en/commit.html
type RollbackStmt struct {
	stmtNode
	// CompletionType overwrites system variable `completion_type` within transaction
	CompletionType CompletionType
}

// Restore implements Node interface.
func (n *RollbackStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *RollbackStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// UseStmt is a statement to use the DBName database as the current database.
// See https://dev.mysql.com/doc/refman/5.7/en/use.html
type UseStmt struct {
	stmtNode

	DBName string
}

// Restore implements Node interface.
func (n *UseStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *UseStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

const (
	// SetNames is the const for set names stmt.
	// If VariableAssignment.Name == Names, it should be set names stmt.
	SetNames = "SetNAMES"
	// SetCharset is the const for set charset stmt.
	SetCharset = "SetCharset"
)

// VariableAssignment is a variable assignment struct.
type VariableAssignment struct {
	node
	Name     string
	Value    ExprNode
	IsGlobal bool
	IsSystem bool

	// ExtendValue is a way to store extended info.
	// VariableAssignment should be able to store information for SetCharset/SetPWD Stmt.
	// For SetCharsetStmt, Value is charset, ExtendValue is collation.
	// TODO: Use SetStmt to implement set password statement.
	ExtendValue ValueExpr
}

// Restore implements Node interface.
func (n *VariableAssignment) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node interface.
func (n *VariableAssignment) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// FlushStmtType is the type for FLUSH statement.
type FlushStmtType int

// Flush statement types.
const (
	FlushNone FlushStmtType = iota
	FlushTables
	FlushPrivileges
	FlushStatus
	FlushTiDBPlugin
	FlushHosts
	FlushLogs
	FlushClientErrorsSummary
)

// LogType is the log type used in FLUSH statement.
type LogType int8

const (
	LogTypeDefault LogType = iota
	LogTypeBinary
	LogTypeEngine
	LogTypeError
	LogTypeGeneral
	LogTypeSlow
)

// FlushStmt is a statement to flush tables/privileges/optimizer costs and so on.
type FlushStmt struct {
	stmtNode

	Tp              FlushStmtType // Privileges/Tables/...
	NoWriteToBinLog bool
	LogType         LogType
	Tables          []*TableName // For FlushTableStmt, if Tables is empty, it means flush all tables.
	ReadLock        bool
	Plugins         []string
}

// Restore implements Node interface.
func (n *FlushStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *FlushStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// KillStmt is a statement to kill a query or connection.
type KillStmt struct {
	stmtNode

	// Query indicates whether terminate a single query on this connection or the whole connection.
	// If Query is true, terminates the statement the connection is currently executing, but leaves the connection itself intact.
	// If Query is false, terminates the connection associated with the given ConnectionID, after terminating any statement the connection is executing.
	Query        bool
	ConnectionID uint64
	// TiDBExtension is used to indicate whether the user knows he is sending kill statement to the right tidb-server.
	// When the SQL grammar is "KILL TIDB [CONNECTION | QUERY] connectionID", TiDBExtension will be set.
	// It's a special grammar extension in TiDB. This extension exists because, when the connection is:
	// client -> LVS proxy -> TiDB, and type Ctrl+C in client, the following action will be executed:
	// new a connection; kill xxx;
	// kill command may send to the wrong TiDB, because the exists of LVS proxy, and kill the wrong session.
	// So, "KILL TIDB" grammar is introduced, and it REQUIRES DIRECT client -> TiDB TOPOLOGY.
	// TODO: The standard KILL grammar will be supported once we have global connectionID.
	TiDBExtension bool
}

// Restore implements Node interface.
func (n *KillStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *KillStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SetStmt is the statement to set variables.
type SetStmt struct {
	stmtNode
	// Variables is the list of variable assignment.
	Variables []*VariableAssignment
}

// Restore implements Node interface.
func (n *SetStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *SetStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SetConfigStmt is the statement to set cluster configs.
type SetConfigStmt struct {
	stmtNode

	Type     string // TiDB, TiKV, PD
	Instance string // '127.0.0.1:3306'
	Name     string // the variable name
	Value    ExprNode
}

func (n *SetConfigStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *SetConfigStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

/*
// SetCharsetStmt is a statement to assign values to character and collation variables.
// See https://dev.mysql.com/doc/refman/5.7/en/set-statement.html
type SetCharsetStmt struct {
	stmtNode

	Charset string
	Collate string
}

// Accept implements Node Accept interface.
func (n *SetCharsetStmt) Accept(v Visitor) (Node, bool) {
	newNode, skipChildren := v.Enter(n)
	if skipChildren {
		return v.Leave(newNode)
	}
	n = newNode.(*SetCharsetStmt)
	return v.Leave(n)
}
*/

// SetPwdStmt is a statement to assign a password to user account.
// See https://dev.mysql.com/doc/refman/5.7/en/set-password.html
type SetPwdStmt struct {
	stmtNode

	User     *auth.UserIdentity
	Password string
}

// Restore implements Node interface.
func (n *SetPwdStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// SecureText implements SensitiveStatement interface.
func (n *SetPwdStmt) SecureText() string { _ = "STUB: not implemented"; return "" }

// Accept implements Node Accept interface.
func (n *SetPwdStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type ChangeStmt struct {
	stmtNode

	NodeType string
	State    string
	NodeID   string
}

// Restore implements Node interface.
func (n *ChangeStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// SecureText implements SensitiveStatement interface.
func (n *ChangeStmt) SecureText() string { _ = "STUB: not implemented"; return "" }

// Accept implements Node Accept interface.
func (n *ChangeStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SetRoleStmtType is the type for FLUSH statement.
type SetRoleStmtType int

// SetRole statement types.
const (
	SetRoleDefault SetRoleStmtType = iota
	SetRoleNone
	SetRoleAll
	SetRoleAllExcept
	SetRoleRegular
)

type SetRoleStmt struct {
	stmtNode

	SetRoleOpt SetRoleStmtType
	RoleList   []*auth.RoleIdentity
}

func (n *SetRoleStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *SetRoleStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type SetDefaultRoleStmt struct {
	stmtNode

	SetRoleOpt SetRoleStmtType
	RoleList   []*auth.RoleIdentity
	UserList   []*auth.UserIdentity
}

func (n *SetDefaultRoleStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *SetDefaultRoleStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// UserSpec is used for parsing create user statement.
type UserSpec struct {
	User    *auth.UserIdentity
	AuthOpt *AuthOption
	IsRole  bool
}

// Restore implements Node interface.
func (n *UserSpec) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// SecurityString formats the UserSpec without password information.
func (n *UserSpec) SecurityString() string { _ = "STUB: not implemented"; return "" }

// EncodedPassword returns the encoded password (which is the real data mysql.user).
// The boolean value indicates input's password format is legal or not.
func (n *UserSpec) EncodedPassword() (string, bool) { _ = "STUB: not implemented"; return "", false }

// In case we have 'IDENTIFIED WITH <plugin>' but no 'BY <password>' to set an empty password.

// Not a legal password string.

const (
	TlsNone = iota
	Ssl
	X509
	Cipher
	Issuer
	Subject
	SAN
)

type TLSOption struct {
	Type  int
	Value string
}

func (t *TLSOption) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

const (
	MaxQueriesPerHour = iota + 1
	MaxUpdatesPerHour
	MaxConnectionsPerHour
	MaxUserConnections
)

type ResourceOption struct {
	Type  int
	Count int64
}

func (r *ResourceOption) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	PasswordExpire = iota + 1
	PasswordExpireDefault
	PasswordExpireNever
	PasswordExpireInterval
	Lock
	Unlock
)

type PasswordOrLockOption struct {
	Type  int
	Count int64
}

func (p *PasswordOrLockOption) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateUserStmt creates user account.
// See https://dev.mysql.com/doc/refman/5.7/en/create-user.html
type CreateUserStmt struct {
	stmtNode

	IsCreateRole          bool
	IfNotExists           bool
	Specs                 []*UserSpec
	TLSOptions            []*TLSOption
	ResourceOptions       []*ResourceOption
	PasswordOrLockOptions []*PasswordOrLockOption
}

// Restore implements Node interface.
func (n *CreateUserStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *CreateUserStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// SecureText implements SensitiveStatement interface.
func (n *CreateUserStmt) SecureText() string { _ = "STUB: not implemented"; return "" }

// AlterUserStmt modifies user account.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-user.html
type AlterUserStmt struct {
	stmtNode

	IfExists              bool
	CurrentAuth           *AuthOption
	Specs                 []*UserSpec
	TLSOptions            []*TLSOption
	ResourceOptions       []*ResourceOption
	PasswordOrLockOptions []*PasswordOrLockOption
}

// Restore implements Node interface.
func (n *AlterUserStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// SecureText implements SensitiveStatement interface.
func (n *AlterUserStmt) SecureText() string { _ = "STUB: not implemented"; return "" }

// Accept implements Node Accept interface.
func (n *AlterUserStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// AlterInstanceStmt modifies instance.
// See https://dev.mysql.com/doc/refman/8.0/en/alter-instance.html
type AlterInstanceStmt struct {
	stmtNode

	ReloadTLS         bool
	NoRollbackOnError bool
}

// Restore implements Node interface.
func (n *AlterInstanceStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *AlterInstanceStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropUserStmt creates user account.
// See http://dev.mysql.com/doc/refman/5.7/en/drop-user.html
type DropUserStmt struct {
	stmtNode

	IfExists   bool
	IsDropRole bool
	UserList   []*auth.UserIdentity
}

// Restore implements Node interface.
func (n *DropUserStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *DropUserStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// CreateBindingStmt creates sql binding hint.
type CreateBindingStmt struct {
	stmtNode

	GlobalScope bool
	OriginNode  StmtNode
	HintedNode  StmtNode
}

func (n *CreateBindingStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *CreateBindingStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropBindingStmt deletes sql binding hint.
type DropBindingStmt struct {
	stmtNode

	GlobalScope bool
	OriginNode  StmtNode
	HintedNode  StmtNode
}

func (n *DropBindingStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *DropBindingStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Extended statistics types.
const (
	StatsTypeCardinality uint8 = iota
	StatsTypeDependency
	StatsTypeCorrelation
)

// StatisticsSpec is the specification for ADD /DROP STATISTICS.
type StatisticsSpec struct {
	StatsName string
	StatsType uint8
	Columns   []*ColumnName
}

// CreateStatisticsStmt is a statement to create extended statistics.
// Examples:
//
//	CREATE STATISTICS stats1 (cardinality) ON t(a, b, c);
//	CREATE STATISTICS stats2 (dependency) ON t(a, b);
//	CREATE STATISTICS stats3 (correlation) ON t(a, b);
type CreateStatisticsStmt struct {
	stmtNode

	IfNotExists bool
	StatsName   string
	StatsType   uint8
	Table       *TableName
	Columns     []*ColumnName
}

// Restore implements Node interface.
func (n *CreateStatisticsStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *CreateStatisticsStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropStatisticsStmt is a statement to drop extended statistics.
// Examples:
//
//	DROP STATISTICS stats1;
type DropStatisticsStmt struct {
	stmtNode

	StatsName string
}

// Restore implements Node interface.
func (n *DropStatisticsStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *DropStatisticsStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DoStmt is the struct for DO statement.
type DoStmt struct {
	stmtNode

	Exprs []ExprNode
}

// Restore implements Node interface.
func (n *DoStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *DoStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// AdminStmtType is the type for admin statement.
type AdminStmtType int

// Admin statement types.
const (
	AdminShowDDL = iota + 1
	AdminCheckTable
	AdminShowDDLJobs
	AdminCancelDDLJobs
	AdminCheckIndex
	AdminRecoverIndex
	AdminCleanupIndex
	AdminCheckIndexRange
	AdminShowDDLJobQueries
	AdminChecksumTable
	AdminShowSlow
	AdminShowNextRowID
	AdminReloadExprPushdownBlacklist
	AdminReloadOptRuleBlacklist
	AdminPluginDisable
	AdminPluginEnable
	AdminFlushBindings
	AdminCaptureBindings
	AdminEvolveBindings
	AdminReloadBindings
	AdminShowTelemetry
	AdminResetTelemetryID
	AdminReloadStatistics
)

// HandleRange represents a range where handle value >= Begin and < End.
type HandleRange struct {
	Begin int64
	End   int64
}

// ShowSlowType defines the type for SlowSlow statement.
type ShowSlowType int

const (
	// ShowSlowTop is a ShowSlowType constant.
	ShowSlowTop ShowSlowType = iota
	// ShowSlowRecent is a ShowSlowType constant.
	ShowSlowRecent
)

// ShowSlowKind defines the kind for SlowSlow statement when the type is ShowSlowTop.
type ShowSlowKind int

const (
	// ShowSlowKindDefault is a ShowSlowKind constant.
	ShowSlowKindDefault ShowSlowKind = iota
	// ShowSlowKindInternal is a ShowSlowKind constant.
	ShowSlowKindInternal
	// ShowSlowKindAll is a ShowSlowKind constant.
	ShowSlowKindAll
)

// ShowSlow is used for the following command:
//
//	admin show slow top [ internal | all] N
//	admin show slow recent N
type ShowSlow struct {
	Tp    ShowSlowType
	Count uint64
	Kind  ShowSlowKind
}

// Restore implements Node interface.
func (n *ShowSlow) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// do nothing

// AdminStmt is the struct for Admin statement.
type AdminStmt struct {
	stmtNode

	Tp        AdminStmtType
	Index     string
	Tables    []*TableName
	JobIDs    []int64
	JobNumber int64

	HandleRanges []HandleRange
	ShowSlow     *ShowSlow
	Plugins      []string
	Where        ExprNode
}

// Restore implements Node interface.
func (n *AdminStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *AdminStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RoleOrPriv is a temporary structure to be further processed into auth.RoleIdentity or PrivElem
type RoleOrPriv struct {
	Symbols string      // hold undecided symbols
	Node    interface{} // hold auth.RoleIdentity or PrivElem that can be sure when parsing
}

func (n *RoleOrPriv) ToRole() (*auth.RoleIdentity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *RoleOrPriv) ToPriv() (*PrivElem, error) { _ = "STUB: not implemented"; return nil, nil }

// PrivElem is the privilege type and optional column list.
type PrivElem struct {
	node

	Priv mysql.PrivilegeType
	Cols []*ColumnName
	Name string
}

// Restore implements Node interface.
func (n *PrivElem) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *PrivElem) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ObjectTypeType is the type for object type.
type ObjectTypeType int

const (
	// ObjectTypeNone is for empty object type.
	ObjectTypeNone ObjectTypeType = iota + 1
	// ObjectTypeTable means the following object is a table.
	ObjectTypeTable
	// ObjectTypeFunction means the following object is a stored function.
	ObjectTypeFunction
	// ObjectTypeProcedure means the following object is a stored procedure.
	ObjectTypeProcedure
)

// Restore implements Node interface.
func (n ObjectTypeType) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

// GrantLevelType is the type for grant level.
type GrantLevelType int

const (
	// GrantLevelNone is the dummy const for default value.
	GrantLevelNone GrantLevelType = iota + 1
	// GrantLevelGlobal means the privileges are administrative or apply to all databases on a given server.
	GrantLevelGlobal
	// GrantLevelDB means the privileges apply to all objects in a given database.
	GrantLevelDB
	// GrantLevelTable means the privileges apply to all columns in a given table.
	GrantLevelTable
)

// GrantLevel is used for store the privilege scope.
type GrantLevel struct {
	Level     GrantLevelType
	DBName    string
	TableName string
}

// Restore implements Node interface.
func (n *GrantLevel) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// RevokeStmt is the struct for REVOKE statement.
type RevokeStmt struct {
	stmtNode

	Privs      []*PrivElem
	ObjectType ObjectTypeType
	Level      *GrantLevel
	Users      []*UserSpec
}

// Restore implements Node interface.
func (n *RevokeStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *RevokeStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RevokeStmt is the struct for REVOKE statement.
type RevokeRoleStmt struct {
	stmtNode

	Roles []*auth.RoleIdentity
	Users []*auth.UserIdentity
}

// Restore implements Node interface.
func (n *RevokeRoleStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *RevokeRoleStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// GrantStmt is the struct for GRANT statement.
type GrantStmt struct {
	stmtNode

	Privs      []*PrivElem
	ObjectType ObjectTypeType
	Level      *GrantLevel
	Users      []*UserSpec
	TLSOptions []*TLSOption
	WithGrant  bool
}

// Restore implements Node interface.
func (n *GrantStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// SecureText implements SensitiveStatement interface.
func (n *GrantStmt) SecureText() string {
	_ = "STUB: not implemented"

	// Filter "identified by xxx" because it would expose password information.
	return ""
}

// Accept implements Node Accept interface.
func (n *GrantStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// GrantProxyStmt is the struct for GRANT PROXY statement.
type GrantProxyStmt struct {
	stmtNode

	LocalUser     *auth.UserIdentity
	ExternalUsers []*auth.UserIdentity
	WithGrant     bool
}

// Accept implements Node Accept interface.
func (n *GrantProxyStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Restore implements Node interface.
func (n *GrantProxyStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// GrantRoleStmt is the struct for GRANT TO statement.
type GrantRoleStmt struct {
	stmtNode

	Roles []*auth.RoleIdentity
	Users []*auth.UserIdentity
}

// Accept implements Node Accept interface.
func (n *GrantRoleStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Restore implements Node interface.
func (n *GrantRoleStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// SecureText implements SensitiveStatement interface.
func (n *GrantRoleStmt) SecureText() string {
	_ = "STUB: not implemented"

	// Filter "identified by xxx" because it would expose password information.
	return ""
}

// ShutdownStmt is a statement to stop the TiDB server.
// See https://dev.mysql.com/doc/refman/5.7/en/shutdown.html
type ShutdownStmt struct {
	stmtNode
}

// Restore implements Node interface.
func (n *ShutdownStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ShutdownStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RestartStmt is a statement to restart the TiDB server.
// See https://dev.mysql.com/doc/refman/8.0/en/restart.html
type RestartStmt struct {
	stmtNode
}

// Restore implements Node interface.
func (n *RestartStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *RestartStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// HelpStmt is a statement for server side help
// See https://dev.mysql.com/doc/refman/8.0/en/help.html
type HelpStmt struct {
	stmtNode

	Topic string
}

// Restore implements Node interface.
func (n *HelpStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *HelpStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RenameUserStmt is a statement to rename a user.
// See http://dev.mysql.com/doc/refman/5.7/en/rename-user.html
type RenameUserStmt struct {
	stmtNode

	UserToUsers []*UserToUser
}

// Restore implements Node interface.
func (n *RenameUserStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *RenameUserStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// UserToUser represents renaming old user to new user used in RenameUserStmt.
type UserToUser struct {
	node
	OldUser *auth.UserIdentity
	NewUser *auth.UserIdentity
}

// Restore implements Node interface.
func (n *UserToUser) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *UserToUser) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type BRIEKind uint8
type BRIEOptionType uint16

const (
	BRIEKindBackup BRIEKind = iota
	BRIEKindRestore

	// common BRIE options
	BRIEOptionRateLimit BRIEOptionType = iota + 1
	BRIEOptionConcurrency
	BRIEOptionChecksum
	BRIEOptionSendCreds
	BRIEOptionCheckpoint
	// backup options
	BRIEOptionBackupTimeAgo
	BRIEOptionBackupTS
	BRIEOptionBackupTSO
	BRIEOptionLastBackupTS
	BRIEOptionLastBackupTSO
	// restore options
	BRIEOptionOnline
	// import options
	BRIEOptionAnalyze
	BRIEOptionBackend
	BRIEOptionOnDuplicate
	BRIEOptionSkipSchemaFiles
	BRIEOptionStrictFormat
	BRIEOptionTiKVImporter
	BRIEOptionResume
	// CSV options
	BRIEOptionCSVBackslashEscape
	BRIEOptionCSVDelimiter
	BRIEOptionCSVHeader
	BRIEOptionCSVNotNull
	BRIEOptionCSVNull
	BRIEOptionCSVSeparator
	BRIEOptionCSVTrimLastSeparators

	BRIECSVHeaderIsColumns = ^uint64(0)
)

type BRIEOptionLevel uint64

const (
	BRIEOptionLevelOff      BRIEOptionLevel = iota // equals FALSE
	BRIEOptionLevelRequired                        // equals TRUE
	BRIEOptionLevelOptional
)

func (kind BRIEKind) String() string { _ = "STUB: not implemented"; return "" }

func (kind BRIEOptionType) String() string { _ = "STUB: not implemented"; return "" }

func (level BRIEOptionLevel) String() string { _ = "STUB: not implemented"; return "" }

type BRIEOption struct {
	Tp        BRIEOptionType
	StrValue  string
	UintValue uint64
}

func (opt *BRIEOption) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// BACKUP/RESTORE doesn't support OPTIONAL value for now, should warn at executor

// BRIEStmt is a statement for backup, restore, import and export.
type BRIEStmt struct {
	stmtNode

	Kind    BRIEKind
	Schemas []string
	Tables  []*TableName
	Storage string
	Options []*BRIEOption
}

func (n *BRIEStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *BRIEStmt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// SecureText implements SensitiveStmtNode
func (n *BRIEStmt) SecureText() string {
	_ = "STUB: not implemented"
	// FIXME: this solution is not scalable, and duplicates some logic from BR.
	return ""
}

type PurgeImportStmt struct {
	stmtNode

	TaskID uint64
}

func (n *PurgeImportStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *PurgeImportStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrorHandlingOption is used in async IMPORT related stmt
type ErrorHandlingOption uint64

const (
	ErrorHandleError ErrorHandlingOption = iota
	ErrorHandleReplace
	ErrorHandleSkipAll
	ErrorHandleSkipConstraint
	ErrorHandleSkipDuplicate
	ErrorHandleSkipStrict
)

func (o ErrorHandlingOption) String() string { _ = "STUB: not implemented"; return "" }

type CreateImportStmt struct {
	stmtNode

	IfNotExists   bool
	Name          string
	Storage       string
	ErrorHandling ErrorHandlingOption
	Options       []*BRIEOption
}

func (n *CreateImportStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *CreateImportStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// SecureText implements SensitiveStmtNode
func (n *CreateImportStmt) SecureText() string {
	_ = "STUB: not implemented"
	// FIXME: this solution is not scalable, and duplicates some logic from BR.
	return ""
}

type StopImportStmt struct {
	stmtNode

	IfRunning bool
	Name      string
}

func (n *StopImportStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *StopImportStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

type ResumeImportStmt struct {
	stmtNode

	IfNotRunning bool
	Name         string
}

func (n *ResumeImportStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *ResumeImportStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

type ImportTruncate struct {
	IsErrorsOnly bool
	TableNames   []*TableName
}

type AlterImportStmt struct {
	stmtNode

	Name          string
	ErrorHandling ErrorHandlingOption
	Options       []*BRIEOption
	Truncate      *ImportTruncate
}

func (n *AlterImportStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *AlterImportStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

type DropImportStmt struct {
	stmtNode

	IfExists bool
	Name     string
}

func (n *DropImportStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *DropImportStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

type ShowImportStmt struct {
	stmtNode

	Name       string
	ErrorsOnly bool
	TableNames []*TableName
}

func (n *ShowImportStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (n *ShowImportStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Ident is the table identifier composed of schema name and table name.
type Ident struct {
	Schema model.CIStr
	Name   model.CIStr
}

// String implements fmt.Stringer interface.
func (i Ident) String() string { _ = "STUB: not implemented"; return "" }

// SelectStmtOpts wrap around select hints and switches
type SelectStmtOpts struct {
	Distinct        bool
	SQLBigResult    bool
	SQLBufferResult bool
	SQLCache        bool
	SQLSmallResult  bool
	CalcFoundRows   bool
	StraightJoin    bool
	Priority        mysql.PriorityEnum
	TableHints      []*TableOptimizerHint
	ExplicitAll     bool
}

// TableOptimizerHint is Table level optimizer hint
type TableOptimizerHint struct {
	node
	// HintName is the name or alias of the table(s) which the hint will affect.
	// Table hints has no schema info
	// It allows only table name or alias (if table has an alias)
	HintName model.CIStr
	// HintData is the payload of the hint. The actual type of this field
	// is defined differently as according `HintName`. Define as following:
	//
	// Statement Execution Time Optimizer Hints
	// See https://dev.mysql.com/doc/refman/5.7/en/optimizer-hints.html#optimizer-hints-execution-time
	// - MAX_EXECUTION_TIME  => uint64
	// - MEMORY_QUOTA        => int64
	// - QUERY_TYPE          => model.CIStr
	//
	// Time Range is used to hint the time range of inspection tables
	// e.g: select /*+ time_range('','') */ * from information_schema.inspection_result.
	// - TIME_RANGE          => ast.HintTimeRange
	// - READ_FROM_STORAGE   => model.CIStr
	// - USE_TOJA            => bool
	// - NTH_PLAN            => int64
	HintData interface{}
	// QBName is the default effective query block of this hint.
	QBName  model.CIStr
	Tables  []HintTable
	Indexes []model.CIStr
}

// HintTimeRange is the payload of `TIME_RANGE` hint
type HintTimeRange struct {
	From string
	To   string
}

// HintSetVar is the payload of `SET_VAR` hint
type HintSetVar struct {
	VarName string
	Value   string
}

// HintTable is table in the hint. It may have query block info.
type HintTable struct {
	DBName        model.CIStr
	TableName     model.CIStr
	QBName        model.CIStr
	PartitionList []model.CIStr
}

func (ht *HintTable) Restore(ctx *format.RestoreCtx) { _ = "STUB: not implemented"; return }

// Restore implements Node interface.
func (n *TableOptimizerHint) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Hints without args except query block.

// Hints with args except query block.

// Accept implements Node Accept interface.
func (n *TableOptimizerHint) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type BinaryLiteral interface {
	ToString() string
}

// NewDecimal creates a types.Decimal value, it's provided by parser driver.
var NewDecimal func(string) (interface{}, error)

// NewHexLiteral creates a types.HexLiteral value, it's provided by parser driver.
var NewHexLiteral func(string) (interface{}, error)

// NewBitLiteral creates a types.BitLiteral value, it's provided by parser driver.
var NewBitLiteral func(string) (interface{}, error)
