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
	"github.com/pingcap/parser/terror"
	"github.com/pingcap/parser/types"
)

var (
	_ DDLNode = &AlterTableStmt{}
	_ DDLNode = &AlterSequenceStmt{}
	_ DDLNode = &AlterPlacementPolicyStmt{}
	_ DDLNode = &CreateDatabaseStmt{}
	_ DDLNode = &CreateIndexStmt{}
	_ DDLNode = &CreateTableStmt{}
	_ DDLNode = &CreateViewStmt{}
	_ DDLNode = &CreateSequenceStmt{}
	_ DDLNode = &CreatePlacementPolicyStmt{}
	_ DDLNode = &DropDatabaseStmt{}
	_ DDLNode = &DropIndexStmt{}
	_ DDLNode = &DropTableStmt{}
	_ DDLNode = &DropSequenceStmt{}
	_ DDLNode = &DropPlacementPolicyStmt{}
	_ DDLNode = &RenameTableStmt{}
	_ DDLNode = &TruncateTableStmt{}
	_ DDLNode = &RepairTableStmt{}

	_ Node = &AlterTableSpec{}
	_ Node = &ColumnDef{}
	_ Node = &ColumnOption{}
	_ Node = &ColumnPosition{}
	_ Node = &Constraint{}
	_ Node = &IndexPartSpecification{}
	_ Node = &ReferenceDef{}
)

// CharsetOpt is used for parsing charset option from SQL.
type CharsetOpt struct {
	Chs string
	Col string
}

// NullString represents a string that may be nil.
type NullString struct {
	String string
	Empty  bool // Empty is true if String is empty backtick.
}

// DatabaseOptionType is the type for database options.
type DatabaseOptionType int

// Database option types.
const (
	DatabaseOptionNone DatabaseOptionType = iota
	DatabaseOptionCharset
	DatabaseOptionCollate
	DatabaseOptionEncryption
	DatabaseOptionPlacementPrimaryRegion       = DatabaseOptionType(PlacementOptionPrimaryRegion)
	DatabaseOptionPlacementRegions             = DatabaseOptionType(PlacementOptionRegions)
	DatabaseOptionPlacementFollowerCount       = DatabaseOptionType(PlacementOptionFollowerCount)
	DatabaseOptionPlacementVoterCount          = DatabaseOptionType(PlacementOptionVoterCount)
	DatabaseOptionPlacementLearnerCount        = DatabaseOptionType(PlacementOptionLearnerCount)
	DatabaseOptionPlacementSchedule            = DatabaseOptionType(PlacementOptionSchedule)
	DatabaseOptionPlacementConstraints         = DatabaseOptionType(PlacementOptionConstraints)
	DatabaseOptionPlacementLeaderConstraints   = DatabaseOptionType(PlacementOptionLeaderConstraints)
	DatabaseOptionPlacementLearnerConstraints  = DatabaseOptionType(PlacementOptionLearnerConstraints)
	DatabaseOptionPlacementFollowerConstraints = DatabaseOptionType(PlacementOptionFollowerConstraints)
	DatabaseOptionPlacementVoterConstraints    = DatabaseOptionType(PlacementOptionVoterConstraints)
	DatabaseOptionPlacementPolicy              = DatabaseOptionType(PlacementOptionPolicy)
)

// DatabaseOption represents database option.
type DatabaseOption struct {
	Tp        DatabaseOptionType
	Value     string
	UintValue uint64
}

// Restore implements Node interface.
func (n *DatabaseOption) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateDatabaseStmt is a statement to create a database.
// See https://dev.mysql.com/doc/refman/5.7/en/create-database.html
type CreateDatabaseStmt struct {
	ddlNode

	IfNotExists bool
	Name        string
	Options     []*DatabaseOption
}

// Restore implements Node interface.
func (n *CreateDatabaseStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *CreateDatabaseStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// AlterDatabaseStmt is a statement to change the structure of a database.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-database.html
type AlterDatabaseStmt struct {
	ddlNode

	Name                 string
	AlterDefaultDatabase bool
	Options              []*DatabaseOption
}

// Restore implements Node interface.
func (n *AlterDatabaseStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *AlterDatabaseStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropDatabaseStmt is a statement to drop a database and all tables in the database.
// See https://dev.mysql.com/doc/refman/5.7/en/drop-database.html
type DropDatabaseStmt struct {
	ddlNode

	IfExists bool
	Name     string
}

// Restore implements Node interface.
func (n *DropDatabaseStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *DropDatabaseStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// IndexPartSpecifications is used for parsing index column name or index expression from SQL.
type IndexPartSpecification struct {
	node

	Column *ColumnName
	Length int
	Expr   ExprNode
}

// Restore implements Node interface.
func (n *IndexPartSpecification) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *IndexPartSpecification) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// MatchType is the type for reference match type.
type MatchType int

// match type
const (
	MatchNone MatchType = iota
	MatchFull
	MatchPartial
	MatchSimple
)

// ReferenceDef is used for parsing foreign key reference option from SQL.
// See http://dev.mysql.com/doc/refman/5.7/en/create-table-foreign-keys.html
type ReferenceDef struct {
	node

	Table                   *TableName
	IndexPartSpecifications []*IndexPartSpecification
	OnDelete                *OnDeleteOpt
	OnUpdate                *OnUpdateOpt
	Match                   MatchType
}

// Restore implements Node interface.
func (n *ReferenceDef) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ReferenceDef) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ReferOptionType is the type for refer options.
type ReferOptionType int

// Refer option types.
const (
	ReferOptionNoOption ReferOptionType = iota
	ReferOptionRestrict
	ReferOptionCascade
	ReferOptionSetNull
	ReferOptionNoAction
	ReferOptionSetDefault
)

// String implements fmt.Stringer interface.
func (r ReferOptionType) String() string { _ = "STUB: not implemented"; return "" }

// OnDeleteOpt is used for optional on delete clause.
type OnDeleteOpt struct {
	node
	ReferOpt ReferOptionType
}

// Restore implements Node interface.
func (n *OnDeleteOpt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *OnDeleteOpt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// OnUpdateOpt is used for optional on update clause.
type OnUpdateOpt struct {
	node
	ReferOpt ReferOptionType
}

// Restore implements Node interface.
func (n *OnUpdateOpt) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *OnUpdateOpt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ColumnOptionType is the type for ColumnOption.
type ColumnOptionType int

// ColumnOption types.
const (
	ColumnOptionNoOption ColumnOptionType = iota
	ColumnOptionPrimaryKey
	ColumnOptionNotNull
	ColumnOptionAutoIncrement
	ColumnOptionDefaultValue
	ColumnOptionUniqKey
	ColumnOptionNull
	ColumnOptionOnUpdate // For Timestamp and Datetime only.
	ColumnOptionFulltext
	ColumnOptionComment
	ColumnOptionGenerated
	ColumnOptionReference
	ColumnOptionCollate
	ColumnOptionCheck
	ColumnOptionColumnFormat
	ColumnOptionStorage
	ColumnOptionAutoRandom
)

var (
	invalidOptionForGeneratedColumn = map[ColumnOptionType]struct{}{
		ColumnOptionAutoIncrement: {},
		ColumnOptionOnUpdate:      {},
		ColumnOptionDefaultValue:  {},
	}
)

// ColumnOption is used for parsing column constraint info from SQL.
type ColumnOption struct {
	node

	Tp ColumnOptionType
	// Expr is used for ColumnOptionDefaultValue/ColumnOptionOnUpdateColumnOptionGenerated.
	// For ColumnOptionDefaultValue or ColumnOptionOnUpdate, it's the target value.
	// For ColumnOptionGenerated, it's the target expression.
	Expr ExprNode
	// Stored is only for ColumnOptionGenerated, default is false.
	Stored bool
	// Refer is used for foreign key.
	Refer               *ReferenceDef
	StrValue            string
	AutoRandomBitLength int
	// Enforced is only for Check, default is true.
	Enforced bool
	// Name is only used for Check Constraint name.
	ConstraintName string
	PrimaryKeyTp   model.PrimaryKeyType
}

// Restore implements Node interface.
func (n *ColumnOption) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ColumnOption) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// IndexVisibility is the option for index visibility.
type IndexVisibility int

// IndexVisibility options.
const (
	IndexVisibilityDefault IndexVisibility = iota
	IndexVisibilityVisible
	IndexVisibilityInvisible
)

// IndexOption is the index options.
//
//	  KEY_BLOCK_SIZE [=] value
//	| index_type
//	| WITH PARSER parser_name
//	| COMMENT 'string'
//
// See http://dev.mysql.com/doc/refman/5.7/en/create-table.html
type IndexOption struct {
	node

	KeyBlockSize uint64
	Tp           model.IndexType
	Comment      string
	ParserName   model.CIStr
	Visibility   IndexVisibility
	PrimaryKeyTp model.PrimaryKeyType
}

// Restore implements Node interface.
func (n *IndexOption) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *IndexOption) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ConstraintType is the type for Constraint.
type ConstraintType int

// ConstraintTypes
const (
	ConstraintNoConstraint ConstraintType = iota
	ConstraintPrimaryKey
	ConstraintKey
	ConstraintIndex
	ConstraintUniq
	ConstraintUniqKey
	ConstraintUniqIndex
	ConstraintForeignKey
	ConstraintFulltext
	ConstraintCheck
)

// Constraint is constraint for table definition.
type Constraint struct {
	node

	// only supported by MariaDB 10.0.2+ (ADD {INDEX|KEY}, ADD FOREIGN KEY),
	// see https://mariadb.com/kb/en/library/alter-table/
	IfNotExists bool

	Tp   ConstraintType
	Name string

	Keys []*IndexPartSpecification // Used for PRIMARY KEY, UNIQUE, ......

	Refer *ReferenceDef // Used for foreign key.

	Option *IndexOption // Index Options

	Expr ExprNode // Used for Check

	Enforced bool // Used for Check

	InColumn bool // Used for Check

	InColumnName string // Used for Check
	IsEmptyIndex bool   // Used for Check
}

// Restore implements Node interface.
func (n *Constraint) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *Constraint) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// ColumnDef is used for parsing column definition from SQL.
type ColumnDef struct {
	node

	Name    *ColumnName
	Tp      *types.FieldType
	Options []*ColumnOption
}

// Restore implements Node interface.
func (n *ColumnDef) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *ColumnDef) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Validate checks if a column definition is legal.
// For example, generated column definitions that contain such
// column options as `ON UPDATE`, `AUTO_INCREMENT`, `DEFAULT`
// are illegal.
func (n *ColumnDef) Validate() bool { _ = "STUB: not implemented"; return false }

type TemporaryKeyword int

const (
	TemporaryNone TemporaryKeyword = iota
	TemporaryGlobal
	TemporaryLocal
)

// CreateTableStmt is a statement to create a table.
// See https://dev.mysql.com/doc/refman/5.7/en/create-table.html
type CreateTableStmt struct {
	ddlNode

	IfNotExists bool
	TemporaryKeyword
	// Meanless when TemporaryKeyword is not TemporaryGlobal.
	// ON COMMIT DELETE ROWS => true
	// ON COMMIT PRESERVE ROW => false
	OnCommitDelete bool
	Table          *TableName
	ReferTable     *TableName
	Cols           []*ColumnDef
	Constraints    []*Constraint
	Options        []*TableOption
	Partition      *PartitionOptions
	OnDuplicate    OnDuplicateKeyHandlingType
	Select         ResultSetNode
}

// Restore implements Node interface.
func (n *CreateTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *CreateTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropTableStmt is a statement to drop one or more tables.
// See https://dev.mysql.com/doc/refman/5.7/en/drop-table.html
type DropTableStmt struct {
	ddlNode

	IfExists         bool
	Tables           []*TableName
	IsView           bool
	TemporaryKeyword // make sense ONLY if/when IsView == false
}

// Restore implements Node interface.
func (n *DropTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *DropTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropPlacementPolicyStmt is a statement to drop a Policy.
type DropPlacementPolicyStmt struct {
	ddlNode

	IfExists   bool
	PolicyName model.CIStr
}

// Restore implements Restore interface.
func (n *DropPlacementPolicyStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *DropPlacementPolicyStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropSequenceStmt is a statement to drop a Sequence.
type DropSequenceStmt struct {
	ddlNode

	IfExists  bool
	Sequences []*TableName
}

// Restore implements Node interface.
func (n *DropSequenceStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *DropSequenceStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RenameTableStmt is a statement to rename a table.
// See http://dev.mysql.com/doc/refman/5.7/en/rename-table.html
type RenameTableStmt struct {
	ddlNode

	TableToTables []*TableToTable
}

// Restore implements Node interface.
func (n *RenameTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *RenameTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// TableToTable represents renaming old table to new table used in RenameTableStmt.
type TableToTable struct {
	node
	OldTable *TableName
	NewTable *TableName
}

// Restore implements Node interface.
func (n *TableToTable) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// Accept implements Node Accept interface.
func (n *TableToTable) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// CreateViewStmt is a statement to create a View.
// See https://dev.mysql.com/doc/refman/5.7/en/create-view.html
type CreateViewStmt struct {
	ddlNode

	OrReplace   bool
	ViewName    *TableName
	Cols        []model.CIStr
	Select      StmtNode
	SchemaCols  []model.CIStr
	Algorithm   model.ViewAlgorithm
	Definer     *auth.UserIdentity
	Security    model.ViewSecurity
	CheckOption model.ViewCheckOption
}

// Restore implements Node interface.
func (n *CreateViewStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// todo Use n.Definer.Restore(ctx) to replace this part

// Accept implements Node Accept interface.
func (n *CreateViewStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// CreatePlacementPolicyStmt is a statement to create a policy.
type CreatePlacementPolicyStmt struct {
	ddlNode

	IfNotExists      bool
	PolicyName       model.CIStr
	PlacementOptions []*PlacementOption
}

// Restore implements Node interface.
func (n *CreatePlacementPolicyStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *CreatePlacementPolicyStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// CreateSequenceStmt is a statement to create a Sequence.
type CreateSequenceStmt struct {
	ddlNode

	// TODO : support or replace if need : care for it will conflict on temporaryOpt.
	IfNotExists bool
	Name        *TableName
	SeqOptions  []*SequenceOption
	TblOptions  []*TableOption
}

// Restore implements Node interface.
func (n *CreateSequenceStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *CreateSequenceStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// IndexLockAndAlgorithm stores the algorithm option and the lock option.
type IndexLockAndAlgorithm struct {
	node

	LockTp      LockType
	AlgorithmTp AlgorithmType
}

// Restore implements Node interface.
func (n *IndexLockAndAlgorithm) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *IndexLockAndAlgorithm) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// IndexKeyType is the type for index key.
type IndexKeyType int

// Index key types.
const (
	IndexKeyTypeNone IndexKeyType = iota
	IndexKeyTypeUnique
	IndexKeyTypeSpatial
	IndexKeyTypeFullText
)

// CreateIndexStmt is a statement to create an index.
// See https://dev.mysql.com/doc/refman/5.7/en/create-index.html
type CreateIndexStmt struct {
	ddlNode

	// only supported by MariaDB 10.0.2+,
	// see https://mariadb.com/kb/en/library/create-index/
	IfNotExists bool

	IndexName               string
	Table                   *TableName
	IndexPartSpecifications []*IndexPartSpecification
	IndexOption             *IndexOption
	KeyType                 IndexKeyType
	LockAlg                 *IndexLockAndAlgorithm
}

// Restore implements Node interface.
func (n *CreateIndexStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *CreateIndexStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// DropIndexStmt is a statement to drop the index.
// See https://dev.mysql.com/doc/refman/5.7/en/drop-index.html
type DropIndexStmt struct {
	ddlNode

	IfExists  bool
	IndexName string
	Table     *TableName
	LockAlg   *IndexLockAndAlgorithm
}

// Restore implements Node interface.
func (n *DropIndexStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *DropIndexStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// LockTablesStmt is a statement to lock tables.
type LockTablesStmt struct {
	ddlNode

	TableLocks []TableLock
}

// TableLock contains the table name and lock type.
type TableLock struct {
	Table *TableName
	Type  model.TableLockType
}

// Accept implements Node Accept interface.
func (n *LockTablesStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Restore implements Node interface.
func (n *LockTablesStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// UnlockTablesStmt is a statement to unlock tables.
type UnlockTablesStmt struct {
	ddlNode
}

// Accept implements Node Accept interface.
func (n *UnlockTablesStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Restore implements Node interface.
func (n *UnlockTablesStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanupTableLockStmt is a statement to cleanup table lock.
type CleanupTableLockStmt struct {
	ddlNode

	Tables []*TableName
}

// Accept implements Node Accept interface.
func (n *CleanupTableLockStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Restore implements Node interface.
func (n *CleanupTableLockStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// RepairTableStmt is a statement to repair tableInfo.
type RepairTableStmt struct {
	ddlNode
	Table      *TableName
	CreateStmt *CreateTableStmt
}

// Accept implements Node Accept interface.
func (n *RepairTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// Restore implements Node interface.
func (n *RepairTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// PlacementOptionType is the type for PlacementOption
type PlacementOptionType int

// PlacementOption types.
const (
	PlacementOptionPrimaryRegion PlacementOptionType = 0x3000 + iota
	PlacementOptionRegions
	PlacementOptionFollowerCount
	PlacementOptionVoterCount
	PlacementOptionLearnerCount
	PlacementOptionSchedule
	PlacementOptionConstraints
	PlacementOptionLeaderConstraints
	PlacementOptionLearnerConstraints
	PlacementOptionFollowerConstraints
	PlacementOptionVoterConstraints
	PlacementOptionPolicy
)

// PlacementOption is used for parsing placement option.
type PlacementOption struct {
	Tp        PlacementOptionType
	StrValue  string
	UintValue uint64
}

func (n *PlacementOption) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteSpecialComment

// TableOptionType is the type for TableOption
type TableOptionType int

// TableOption types.
const (
	TableOptionNone TableOptionType = iota
	TableOptionEngine
	TableOptionCharset
	TableOptionCollate
	TableOptionAutoIdCache
	TableOptionAutoIncrement
	TableOptionAutoRandomBase
	TableOptionComment
	TableOptionAvgRowLength
	TableOptionCheckSum
	TableOptionCompression
	TableOptionConnection
	TableOptionPassword
	TableOptionKeyBlockSize
	TableOptionMaxRows
	TableOptionMinRows
	TableOptionDelayKeyWrite
	TableOptionRowFormat
	TableOptionStatsPersistent
	TableOptionStatsAutoRecalc
	TableOptionShardRowID
	TableOptionPreSplitRegion
	TableOptionPackKeys
	TableOptionTablespace
	TableOptionNodegroup
	TableOptionDataDirectory
	TableOptionIndexDirectory
	TableOptionStorageMedia
	TableOptionStatsSamplePages
	TableOptionSecondaryEngine
	TableOptionSecondaryEngineNull
	TableOptionInsertMethod
	TableOptionTableCheckSum
	TableOptionUnion
	TableOptionEncryption
	TableOptionPlacementPrimaryRegion       = TableOptionType(PlacementOptionPrimaryRegion)
	TableOptionPlacementRegions             = TableOptionType(PlacementOptionRegions)
	TableOptionPlacementFollowerCount       = TableOptionType(PlacementOptionFollowerCount)
	TableOptionPlacementVoterCount          = TableOptionType(PlacementOptionVoterCount)
	TableOptionPlacementLearnerCount        = TableOptionType(PlacementOptionLearnerCount)
	TableOptionPlacementSchedule            = TableOptionType(PlacementOptionSchedule)
	TableOptionPlacementConstraints         = TableOptionType(PlacementOptionConstraints)
	TableOptionPlacementLeaderConstraints   = TableOptionType(PlacementOptionLeaderConstraints)
	TableOptionPlacementLearnerConstraints  = TableOptionType(PlacementOptionLearnerConstraints)
	TableOptionPlacementFollowerConstraints = TableOptionType(PlacementOptionFollowerConstraints)
	TableOptionPlacementVoterConstraints    = TableOptionType(PlacementOptionVoterConstraints)
	TableOptionPlacementPolicy              = TableOptionType(PlacementOptionPolicy)
)

// RowFormat types
const (
	RowFormatDefault uint64 = iota + 1
	RowFormatDynamic
	RowFormatFixed
	RowFormatCompressed
	RowFormatRedundant
	RowFormatCompact
	TokuDBRowFormatDefault
	TokuDBRowFormatFast
	TokuDBRowFormatSmall
	TokuDBRowFormatZlib
	TokuDBRowFormatQuickLZ
	TokuDBRowFormatLzma
	TokuDBRowFormatSnappy
	TokuDBRowFormatUncompressed
)

// OnDuplicateKeyHandlingType is the option that handle unique key values in 'CREATE TABLE ... SELECT' or `LOAD DATA`.
// See https://dev.mysql.com/doc/refman/5.7/en/create-table-select.html
// See https://dev.mysql.com/doc/refman/5.7/en/load-data.html
type OnDuplicateKeyHandlingType int

// OnDuplicateKeyHandling types
const (
	OnDuplicateKeyHandlingError OnDuplicateKeyHandlingType = iota
	OnDuplicateKeyHandlingIgnore
	OnDuplicateKeyHandlingReplace
)

const (
	TableOptionCharsetWithoutConvertTo uint64 = 0
	TableOptionCharsetWithConvertTo    uint64 = 1
)

// TableOption is used for parsing table option from SQL.
type TableOption struct {
	Tp         TableOptionType
	Default    bool
	StrValue   string
	UintValue  uint64
	BoolValue  bool
	TableNames []*TableName
}

func (n *TableOption) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// TODO: not support

// TODO: not support

// SequenceOptionType is the type for SequenceOption
type SequenceOptionType int

// SequenceOption types.
const (
	SequenceOptionNone SequenceOptionType = iota
	SequenceOptionIncrementBy
	SequenceStartWith
	SequenceNoMinValue
	SequenceMinValue
	SequenceNoMaxValue
	SequenceMaxValue
	SequenceNoCache
	SequenceCache
	SequenceNoCycle
	SequenceCycle
	// SequenceRestart is only used in alter sequence statement.
	SequenceRestart
	SequenceRestartWith
)

// SequenceOption is used for parsing sequence option from SQL.
type SequenceOption struct {
	Tp       SequenceOptionType
	IntValue int64
}

func (n *SequenceOption) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// ColumnPositionType is the type for ColumnPosition.
type ColumnPositionType int

// ColumnPosition Types
const (
	ColumnPositionNone ColumnPositionType = iota
	ColumnPositionFirst
	ColumnPositionAfter
)

// ColumnPosition represent the position of the newly added column
type ColumnPosition struct {
	node
	// Tp is either ColumnPositionNone, ColumnPositionFirst or ColumnPositionAfter.
	Tp ColumnPositionType
	// RelativeColumn is the column the newly added column after if type is ColumnPositionAfter
	RelativeColumn *ColumnName
}

// Restore implements Node interface.
func (n *ColumnPosition) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

// Accept implements Node Accept interface.
func (n *ColumnPosition) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// AlterTableType is the type for AlterTableSpec.
type AlterTableType int

// AlterTable types.
const (
	AlterTableOption AlterTableType = iota + 1
	AlterTableAddColumns
	AlterTableAddConstraint
	AlterTableDropColumn
	AlterTableDropPrimaryKey
	AlterTableDropIndex
	AlterTableDropForeignKey
	AlterTableModifyColumn
	AlterTableChangeColumn
	AlterTableRenameColumn
	AlterTableRenameTable
	AlterTableAlterColumn
	AlterTableLock
	AlterTableWriteable
	AlterTableAlgorithm
	AlterTableRenameIndex
	AlterTableForce
	AlterTableAddPartitions
	AlterTableAlterPartition
	AlterTablePartitionAttributes
	AlterTablePartitionOptions
	AlterTableCoalescePartitions
	AlterTableDropPartition
	AlterTableTruncatePartition
	AlterTablePartition
	AlterTableEnableKeys
	AlterTableDisableKeys
	AlterTableRemovePartitioning
	AlterTableWithValidation
	AlterTableWithoutValidation
	AlterTableSecondaryLoad
	AlterTableSecondaryUnload
	AlterTableRebuildPartition
	AlterTableReorganizePartition
	AlterTableCheckPartitions
	AlterTableExchangePartition
	AlterTableOptimizePartition
	AlterTableRepairPartition
	AlterTableImportPartitionTablespace
	AlterTableDiscardPartitionTablespace
	AlterTableAlterCheck
	AlterTableDropCheck
	AlterTableImportTablespace
	AlterTableDiscardTablespace
	AlterTableIndexInvisible
	// TODO: Add more actions
	AlterTableOrderByColumns
	// AlterTableSetTiFlashReplica uses to set the table TiFlash replica.
	AlterTableSetTiFlashReplica
	AlterTablePlacement
	AlterTableAddStatistics
	AlterTableDropStatistics
	AlterTableAttributes
)

// LockType is the type for AlterTableSpec.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-table.html#alter-table-concurrency
type LockType byte

func (n LockType) String() string { _ = "STUB: not implemented"; return "" }

// Lock Types.
const (
	LockTypeNone LockType = iota + 1
	LockTypeDefault
	LockTypeShared
	LockTypeExclusive
)

// AlgorithmType is the algorithm of the DDL operations.
// See https://dev.mysql.com/doc/refman/8.0/en/alter-table.html#alter-table-performance.
type AlgorithmType byte

// DDL algorithms.
// For now, TiDB only supported inplace and instance algorithms. If the user specify `copy`,
// will get an error.
const (
	AlgorithmTypeDefault AlgorithmType = iota
	AlgorithmTypeCopy
	AlgorithmTypeInplace
	AlgorithmTypeInstant
)

func (a AlgorithmType) String() string { _ = "STUB: not implemented"; return "" }

// AlterTableSpec represents alter table specification.
type AlterTableSpec struct {
	node

	// only supported by MariaDB 10.0.2+ (DROP COLUMN, CHANGE COLUMN, MODIFY COLUMN, DROP INDEX, DROP FOREIGN KEY, DROP PARTITION)
	// see https://mariadb.com/kb/en/library/alter-table/
	IfExists bool

	// only supported by MariaDB 10.0.2+ (ADD COLUMN, ADD PARTITION)
	// see https://mariadb.com/kb/en/library/alter-table/
	IfNotExists bool

	NoWriteToBinlog bool
	OnAllPartitions bool

	Tp              AlterTableType
	Name            string
	IndexName       model.CIStr
	Constraint      *Constraint
	Options         []*TableOption
	OrderByList     []*AlterOrderItem
	NewTable        *TableName
	NewColumns      []*ColumnDef
	NewConstraints  []*Constraint
	OldColumnName   *ColumnName
	NewColumnName   *ColumnName
	Position        *ColumnPosition
	LockType        LockType
	Algorithm       AlgorithmType
	Comment         string
	FromKey         model.CIStr
	ToKey           model.CIStr
	Partition       *PartitionOptions
	PartitionNames  []model.CIStr
	PartDefinitions []*PartitionDefinition
	WithValidation  bool
	Num             uint64
	Visibility      IndexVisibility
	TiFlashReplica  *TiFlashReplicaSpec
	PlacementSpecs  []*PlacementSpec
	Writeable       bool
	Statistics      *StatisticsSpec
	AttributesSpec  *AttributesSpec
}

type TiFlashReplicaSpec struct {
	Count  uint64
	Labels []string
}

// AlterOrderItem represents an item in order by at alter table stmt.
type AlterOrderItem struct {
	node
	Column *ColumnName
	Desc   bool
}

// Restore implements Node interface.
func (n *AlterOrderItem) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Restore implements Node interface.
func (n *AlterTableSpec) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: RestrictOrCascadeOpt not support

// TODO: not support

// TODO: not support

// Accept implements Node Accept interface.
func (n *AlterTableSpec) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// AlterTableStmt is a statement to change the structure of a table.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-table.html
type AlterTableStmt struct {
	ddlNode

	Table *TableName
	Specs []*AlterTableSpec
}

// Restore implements Node interface.
func (n *AlterTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *AlterTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// TruncateTableStmt is a statement to empty a table completely.
// See https://dev.mysql.com/doc/refman/5.7/en/truncate-table.html
type TruncateTableStmt struct {
	ddlNode

	Table *TableName
}

// Restore implements Node interface.
func (n *TruncateTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *TruncateTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

var (
	ErrNoParts                              = terror.ClassDDL.NewStd(mysql.ErrNoParts)
	ErrPartitionColumnList                  = terror.ClassDDL.NewStd(mysql.ErrPartitionColumnList)
	ErrPartitionRequiresValues              = terror.ClassDDL.NewStd(mysql.ErrPartitionRequiresValues)
	ErrPartitionsMustBeDefined              = terror.ClassDDL.NewStd(mysql.ErrPartitionsMustBeDefined)
	ErrPartitionWrongNoPart                 = terror.ClassDDL.NewStd(mysql.ErrPartitionWrongNoPart)
	ErrPartitionWrongNoSubpart              = terror.ClassDDL.NewStd(mysql.ErrPartitionWrongNoSubpart)
	ErrPartitionWrongValues                 = terror.ClassDDL.NewStd(mysql.ErrPartitionWrongValues)
	ErrRowSinglePartitionField              = terror.ClassDDL.NewStd(mysql.ErrRowSinglePartitionField)
	ErrSubpartition                         = terror.ClassDDL.NewStd(mysql.ErrSubpartition)
	ErrSystemVersioningWrongPartitions      = terror.ClassDDL.NewStd(mysql.ErrSystemVersioningWrongPartitions)
	ErrTooManyValues                        = terror.ClassDDL.NewStd(mysql.ErrTooManyValues)
	ErrWrongPartitionTypeExpectedSystemTime = terror.ClassDDL.NewStd(mysql.ErrWrongPartitionTypeExpectedSystemTime)
	ErrUnknownCharacterSet                  = terror.ClassDDL.NewStd(mysql.ErrUnknownCharacterSet)
)

type SubPartitionDefinition struct {
	Name    model.CIStr
	Options []*TableOption
}

func (spd *SubPartitionDefinition) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

type PartitionDefinitionClause interface {
	restore(ctx *format.RestoreCtx) error
	acceptInPlace(v Visitor) bool
	// Validate checks if the clause is consistent with the given options.
	// `pt` can be 0 and `columns` can be -1 to skip checking the clause against
	// the partition type or number of columns in the expression list.
	Validate(pt model.PartitionType, columns int) error
}

type PartitionDefinitionClauseNone struct{}

func (n *PartitionDefinitionClauseNone) restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *PartitionDefinitionClauseNone) acceptInPlace(v Visitor) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *PartitionDefinitionClauseNone) Validate(pt model.PartitionType, columns int) error {
	_ = "STUB: not implemented"
	return nil
}

type PartitionDefinitionClauseLessThan struct {
	Exprs []ExprNode
}

func (n *PartitionDefinitionClauseLessThan) restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *PartitionDefinitionClauseLessThan) acceptInPlace(v Visitor) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *PartitionDefinitionClauseLessThan) Validate(pt model.PartitionType, columns int) error {
	_ = "STUB: not implemented"
	return nil
}

type PartitionDefinitionClauseIn struct {
	Values [][]ExprNode
}

func (n *PartitionDefinitionClauseIn) restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	// we special-case an empty list of values to mean MariaDB's "DEFAULT" clause.
	return nil
}

func (n *PartitionDefinitionClauseIn) acceptInPlace(v Visitor) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *PartitionDefinitionClauseIn) Validate(pt model.PartitionType, columns int) error {
	_ = "STUB: not implemented"
	return nil
}

type PartitionDefinitionClauseHistory struct {
	Current bool
}

func (n *PartitionDefinitionClauseHistory) restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *PartitionDefinitionClauseHistory) acceptInPlace(v Visitor) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *PartitionDefinitionClauseHistory) Validate(pt model.PartitionType, columns int) error {
	_ = "STUB: not implemented"
	return nil
}

// PartitionDefinition defines a single partition.
type PartitionDefinition struct {
	Name    model.CIStr
	Clause  PartitionDefinitionClause
	Options []*TableOption
	Sub     []*SubPartitionDefinition
}

// Comment returns the comment option given to this definition.
// The second return value indicates if the comment option exists.
func (n *PartitionDefinition) Comment() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (n *PartitionDefinition) acceptInPlace(v Visitor) bool {
	_ = "STUB: not implemented"
	return false
}

// Restore implements Node interface.
func (n *PartitionDefinition) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// PartitionMethod describes how partitions or subpartitions are constructed.
type PartitionMethod struct {
	// Tp is the type of the partition function
	Tp model.PartitionType
	// Linear is a modifier to the HASH and KEY type for choosing a different
	// algorithm
	Linear bool
	// Expr is an expression used as argument of HASH, RANGE, LIST and
	// SYSTEM_TIME types
	Expr ExprNode
	// ColumnNames is a list of column names used as argument of KEY,
	// RANGE COLUMNS and LIST COLUMNS types
	ColumnNames []*ColumnName
	// Unit is a time unit used as argument of SYSTEM_TIME type
	Unit TimeUnitType
	// Limit is a row count used as argument of the SYSTEM_TIME type
	Limit uint64

	// Num is the number of (sub)partitions required by the method.
	Num uint64

	// KeyAlgorithm is the optional hash algorithm type for `PARTITION BY [LINEAR] KEY` syntax.
	KeyAlgorithm *PartitionKeyAlgorithm
}

type PartitionKeyAlgorithm struct {
	Type uint64
}

// Restore implements the Node interface
func (n *PartitionMethod) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// acceptInPlace is like Node.Accept but does not allow replacing the node itself.
func (n *PartitionMethod) acceptInPlace(v Visitor) bool { _ = "STUB: not implemented"; return false }

// PartitionOptions specifies the partition options.
type PartitionOptions struct {
	node
	PartitionMethod
	Sub         *PartitionMethod
	Definitions []*PartitionDefinition
}

// Validate checks if the partition is well-formed.
func (n *PartitionOptions) Validate() error {
	_ = "STUB: not implemented"
	// if both a partition list and the partition numbers are specified, their values must match
	return nil
}

// now check the subpartition count

// ensure the subpartition count for every partitions are the same
// then normalize n.Num and n.Sub.Num so equality comparison works.

// ensure the partition definition types match the methods,
// e.g. RANGE partitions only allows VALUES LESS THAN

func (n *PartitionOptions) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *PartitionOptions) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// RecoverTableStmt is a statement to recover dropped table.
type RecoverTableStmt struct {
	ddlNode

	JobID  int64
	Table  *TableName
	JobNum int64
}

// Restore implements Node interface.
func (n *RecoverTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *RecoverTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// FlashBackTableStmt is a statement to restore a dropped/truncate table.
type FlashBackTableStmt struct {
	ddlNode

	Table   *TableName
	NewName string
}

// Restore implements Node interface.
func (n *FlashBackTableStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// Accept implements Node Accept interface.
func (n *FlashBackTableStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type PlacementActionType int

const (
	PlacementAdd PlacementActionType = iota + 1
	PlacementAlter
	PlacementDrop
)

type PlacementRole int

const (
	PlacementRoleNone PlacementRole = iota
	PlacementRoleLeader
	PlacementRoleFollower
	PlacementRoleLearner
	PlacementRoleVoter
)

type PlacementSpec struct {
	node

	Tp          PlacementActionType
	Constraints string
	Role        PlacementRole
	Replicas    uint64
}

func (n *PlacementSpec) restoreRole(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *PlacementSpec) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *PlacementSpec) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

type AttributesSpec struct {
	node

	Attributes string
	Default    bool
}

func (n *AttributesSpec) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *AttributesSpec) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// AlterPlacementPolicyStmt is a statement to alter placement policy option.
type AlterPlacementPolicyStmt struct {
	ddlNode

	PolicyName       model.CIStr
	IfExists         bool
	PlacementOptions []*PlacementOption
}

func (n *AlterPlacementPolicyStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *AlterPlacementPolicyStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

// AlterSequenceStmt is a statement to alter sequence option.
type AlterSequenceStmt struct {
	ddlNode

	// sequence name
	Name *TableName

	IfExists   bool
	SeqOptions []*SequenceOption
}

func (n *AlterSequenceStmt) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *AlterSequenceStmt) Accept(v Visitor) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}
