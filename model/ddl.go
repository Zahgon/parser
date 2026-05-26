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

package model

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/pingcap/errors"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/parser/terror"
)

// ActionType is the type for DDL action.
type ActionType byte

// List DDL actions.
const (
	ActionNone                          ActionType = 0
	ActionCreateSchema                  ActionType = 1
	ActionDropSchema                    ActionType = 2
	ActionCreateTable                   ActionType = 3
	ActionDropTable                     ActionType = 4
	ActionAddColumn                     ActionType = 5
	ActionDropColumn                    ActionType = 6
	ActionAddIndex                      ActionType = 7
	ActionDropIndex                     ActionType = 8
	ActionAddForeignKey                 ActionType = 9
	ActionDropForeignKey                ActionType = 10
	ActionTruncateTable                 ActionType = 11
	ActionModifyColumn                  ActionType = 12
	ActionRebaseAutoID                  ActionType = 13
	ActionRenameTable                   ActionType = 14
	ActionSetDefaultValue               ActionType = 15
	ActionShardRowID                    ActionType = 16
	ActionModifyTableComment            ActionType = 17
	ActionRenameIndex                   ActionType = 18
	ActionAddTablePartition             ActionType = 19
	ActionDropTablePartition            ActionType = 20
	ActionCreateView                    ActionType = 21
	ActionModifyTableCharsetAndCollate  ActionType = 22
	ActionTruncateTablePartition        ActionType = 23
	ActionDropView                      ActionType = 24
	ActionRecoverTable                  ActionType = 25
	ActionModifySchemaCharsetAndCollate ActionType = 26
	ActionLockTable                     ActionType = 27
	ActionUnlockTable                   ActionType = 28
	ActionRepairTable                   ActionType = 29
	ActionSetTiFlashReplica             ActionType = 30
	ActionUpdateTiFlashReplicaStatus    ActionType = 31
	ActionAddPrimaryKey                 ActionType = 32
	ActionDropPrimaryKey                ActionType = 33
	ActionCreateSequence                ActionType = 34
	ActionAlterSequence                 ActionType = 35
	ActionDropSequence                  ActionType = 36
	ActionAddColumns                    ActionType = 37
	ActionDropColumns                   ActionType = 38
	ActionModifyTableAutoIdCache        ActionType = 39
	ActionRebaseAutoRandomBase          ActionType = 40
	ActionAlterIndexVisibility          ActionType = 41
	ActionExchangeTablePartition        ActionType = 42
	ActionAddCheckConstraint            ActionType = 43
	ActionDropCheckConstraint           ActionType = 44
	ActionAlterCheckConstraint          ActionType = 45
	ActionAlterTableAlterPartition      ActionType = 46
	ActionRenameTables                  ActionType = 47
	ActionDropIndexes                   ActionType = 48
	ActionAlterTableAttributes          ActionType = 49
	ActionAlterTablePartitionAttributes ActionType = 50
	ActionCreatePlacementPolicy         ActionType = 51
	ActionAlterPlacementPolicy          ActionType = 52
	ActionDropPlacementPolicy           ActionType = 53
	ActionAlterTablePartitionPolicy     ActionType = 54
	ActionModifySchemaDefaultPlacement  ActionType = 55
)

var actionMap = map[ActionType]string{
	ActionCreateSchema:                  "create schema",
	ActionDropSchema:                    "drop schema",
	ActionCreateTable:                   "create table",
	ActionDropTable:                     "drop table",
	ActionAddColumn:                     "add column",
	ActionDropColumn:                    "drop column",
	ActionAddIndex:                      "add index",
	ActionDropIndex:                     "drop index",
	ActionAddForeignKey:                 "add foreign key",
	ActionDropForeignKey:                "drop foreign key",
	ActionTruncateTable:                 "truncate table",
	ActionModifyColumn:                  "modify column",
	ActionRebaseAutoID:                  "rebase auto_increment ID",
	ActionRenameTable:                   "rename table",
	ActionSetDefaultValue:               "set default value",
	ActionShardRowID:                    "shard row ID",
	ActionModifyTableComment:            "modify table comment",
	ActionRenameIndex:                   "rename index",
	ActionAddTablePartition:             "add partition",
	ActionDropTablePartition:            "drop partition",
	ActionCreateView:                    "create view",
	ActionModifyTableCharsetAndCollate:  "modify table charset and collate",
	ActionTruncateTablePartition:        "truncate partition",
	ActionDropView:                      "drop view",
	ActionRecoverTable:                  "recover table",
	ActionModifySchemaCharsetAndCollate: "modify schema charset and collate",
	ActionLockTable:                     "lock table",
	ActionUnlockTable:                   "unlock table",
	ActionRepairTable:                   "repair table",
	ActionSetTiFlashReplica:             "set tiflash replica",
	ActionUpdateTiFlashReplicaStatus:    "update tiflash replica status",
	ActionAddPrimaryKey:                 "add primary key",
	ActionDropPrimaryKey:                "drop primary key",
	ActionCreateSequence:                "create sequence",
	ActionAlterSequence:                 "alter sequence",
	ActionDropSequence:                  "drop sequence",
	ActionAddColumns:                    "add multi-columns",
	ActionDropColumns:                   "drop multi-columns",
	ActionModifyTableAutoIdCache:        "modify auto id cache",
	ActionRebaseAutoRandomBase:          "rebase auto_random ID",
	ActionAlterIndexVisibility:          "alter index visibility",
	ActionExchangeTablePartition:        "exchange partition",
	ActionAddCheckConstraint:            "add check constraint",
	ActionDropCheckConstraint:           "drop check constraint",
	ActionAlterCheckConstraint:          "alter check constraint",
	ActionAlterTableAlterPartition:      "alter partition",
	ActionDropIndexes:                   "drop multi-indexes",
	ActionAlterTableAttributes:          "alter table attributes",
	ActionAlterTablePartitionAttributes: "alter table partition attributes",
	ActionCreatePlacementPolicy:         "create placement policy",
	ActionAlterPlacementPolicy:          "alter placement policy",
	ActionDropPlacementPolicy:           "drop placement policy",
	ActionModifySchemaDefaultPlacement:  "modify schema default placement",
}

// String return current ddl action in string
func (action ActionType) String() string { _ = "STUB: not implemented"; return "" }

// HistoryInfo is used for binlog.
type HistoryInfo struct {
	SchemaVersion int64
	DBInfo        *DBInfo
	TableInfo     *TableInfo
	FinishedTS    uint64
}

// AddDBInfo adds schema version and schema information that are used for binlog.
// dbInfo is added in the following operations: create database, drop database.
func (h *HistoryInfo) AddDBInfo(schemaVer int64, dbInfo *DBInfo) { _ = "STUB: not implemented"; return }

// AddTableInfo adds schema version and table information that are used for binlog.
// tblInfo is added except for the following operations: create database, drop database.
func (h *HistoryInfo) AddTableInfo(schemaVer int64, tblInfo *TableInfo) {
	_ = "STUB: not implemented"
	return
}

// Clean cleans history information.
func (h *HistoryInfo) Clean() { _ = "STUB: not implemented"; return }

// DDLReorgMeta is meta info of DDL reorganization.
type DDLReorgMeta struct {
	// EndHandle is the last handle of the adding indices table.
	// We should only backfill indices in the range [startHandle, EndHandle].
	EndHandle int64 `json:"end_handle"`

	SQLMode       mysql.SQLMode                    `json:"sql_mode"`
	Warnings      map[errors.ErrorID]*terror.Error `json:"warnings"`
	WarningsCount map[errors.ErrorID]int64         `json:"warnings_count"`
}

// NewDDLReorgMeta new a DDLReorgMeta.
func NewDDLReorgMeta() *DDLReorgMeta { _ = "STUB: not implemented"; return nil }

// MultiSchemaInfo keeps some information for multi schema change.
type MultiSchemaInfo struct {
	Warnings []*errors.Error
}

// Job is for a DDL operation.
type Job struct {
	ID         int64         `json:"id"`
	Type       ActionType    `json:"type"`
	SchemaID   int64         `json:"schema_id"`
	TableID    int64         `json:"table_id"`
	SchemaName string        `json:"schema_name"`
	State      JobState      `json:"state"`
	Error      *terror.Error `json:"err"`
	// ErrorCount will be increased, every time we meet an error when running job.
	ErrorCount int64 `json:"err_count"`
	// RowCount means the number of rows that are processed.
	RowCount int64      `json:"row_count"`
	Mu       sync.Mutex `json:"-"`
	// CtxVars are variables attached to the job. It is for internal usage.
	// E.g. passing arguments between functions by one single *Job pointer.
	CtxVars []interface{} `json:"-"`
	Args    []interface{} `json:"-"`
	// RawArgs : We must use json raw message to delay parsing special args.
	RawArgs     json.RawMessage `json:"raw_args"`
	SchemaState SchemaState     `json:"schema_state"`
	// SnapshotVer means snapshot version for this job.
	SnapshotVer uint64 `json:"snapshot_ver"`
	// RealStartTS uses timestamp allocated by TSO.
	// Now it's the TS when we actually start the job.
	RealStartTS uint64 `json:"real_start_ts"`
	// StartTS uses timestamp allocated by TSO.
	// Now it's the TS when we put the job to TiKV queue.
	StartTS uint64 `json:"start_ts"`
	// DependencyID is the job's ID that the current job depends on.
	DependencyID int64 `json:"dependency_id"`
	// Query string of the ddl job.
	Query      string       `json:"query"`
	BinlogInfo *HistoryInfo `json:"binlog"`

	// Version indicates the DDL job version. For old jobs, it will be 0.
	Version int64 `json:"version"`

	// ReorgMeta is meta info of ddl reorganization.
	// This field is depreciated.
	ReorgMeta *DDLReorgMeta `json:"reorg_meta"`

	// MultiSchemaInfo keeps some warning now for multi schema change.
	MultiSchemaInfo *MultiSchemaInfo `json:"multi_schema_info"`

	// Priority is only used to set the operation priority of adding indices.
	Priority int `json:"priority"`
}

// FinishTableJob is called when a job is finished.
// It updates the job's state information and adds tblInfo to the binlog.
func (job *Job) FinishTableJob(jobState JobState, schemaState SchemaState, ver int64, tblInfo *TableInfo) {
	_ = "STUB: not implemented"
	return
}

// FinishDBJob is called when a job is finished.
// It updates the job's state information and adds dbInfo the binlog.
func (job *Job) FinishDBJob(jobState JobState, schemaState SchemaState, ver int64, dbInfo *DBInfo) {
	_ = "STUB: not implemented"
	return
}

// TSConvert2Time converts timestamp to time.
func TSConvert2Time(ts uint64) time.Time {
	_ = "STUB: not implemented"
	// 18 is for the logical time.
	return *new(time.Time)
}

// SetRowCount sets the number of rows. Make sure it can pass `make race`.
func (job *Job) SetRowCount(count int64) { _ = "STUB: not implemented"; return }

// GetRowCount gets the number of rows. Make sure it can pass `make race`.
func (job *Job) GetRowCount() int64 { _ = "STUB: not implemented"; return 0 }

// SetWarnings sets the warnings of rows handled.
func (job *Job) SetWarnings(warnings map[errors.ErrorID]*terror.Error, warningsCount map[errors.ErrorID]int64) {
	_ = "STUB: not implemented"
	return
}

// GetWarnings gets the warnings of the rows handled.
func (job *Job) GetWarnings() (map[errors.ErrorID]*terror.Error, map[errors.ErrorID]int64) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encode encodes job with json format.
// updateRawArgs is used to determine whether to update the raw args.
func (job *Job) Encode(updateRawArgs bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode decodes job from the json buffer, we must use DecodeArgs later to
// decode special args for this job.
func (job *Job) Decode(b []byte) error { _ = "STUB: not implemented"; return nil }

// DecodeArgs decodes job args.
func (job *Job) DecodeArgs(args ...interface{}) error { _ = "STUB: not implemented"; return nil }

// String implements fmt.Stringer interface.
func (job *Job) String() string { _ = "STUB: not implemented"; return "" }

func (job *Job) hasDependentSchema(other *Job) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsDependentOn returns whether the job depends on "other".
// How to check the job depends on "other"?
// 1. The two jobs handle the same database when one of the two jobs is an ActionDropSchema or ActionCreateSchema type.
// 2. Or the two jobs handle the same table.
func (job *Job) IsDependentOn(other *Job) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// TODO: If a job is ActionRenameTable, we need to check table name.

// IsFinished returns whether job is finished or not.
// If the job state is Done or Cancelled, it is finished.
func (job *Job) IsFinished() bool { _ = "STUB: not implemented"; return false }

// IsCancelled returns whether the job is cancelled or not.
func (job *Job) IsCancelled() bool { _ = "STUB: not implemented"; return false }

// IsRollbackDone returns whether the job is rolled back or not.
func (job *Job) IsRollbackDone() bool { _ = "STUB: not implemented"; return false }

// IsRollingback returns whether the job is rolling back or not.
func (job *Job) IsRollingback() bool { _ = "STUB: not implemented"; return false }

// IsCancelling returns whether the job is cancelling or not.
func (job *Job) IsCancelling() bool { _ = "STUB: not implemented"; return false }

// IsSynced returns whether the DDL modification is synced among all TiDB servers.
func (job *Job) IsSynced() bool { _ = "STUB: not implemented"; return false }

// IsDone returns whether job is done.
func (job *Job) IsDone() bool { _ = "STUB: not implemented"; return false }

// IsRunning returns whether job is still running or not.
func (job *Job) IsRunning() bool { _ = "STUB: not implemented"; return false }

// JobState is for job state.
type JobState byte

// List job states.
const (
	JobStateNone    JobState = 0
	JobStateRunning JobState = 1
	// When DDL encountered an unrecoverable error at reorganization state,
	// some keys has been added already, we need to remove them.
	// JobStateRollingback is the state to do the rolling back job.
	JobStateRollingback  JobState = 2
	JobStateRollbackDone JobState = 3
	JobStateDone         JobState = 4
	JobStateCancelled    JobState = 5
	// JobStateSynced is used to mark the information about the completion of this job
	// has been synchronized to all servers.
	JobStateSynced JobState = 6
	// JobStateCancelling is used to mark the DDL job is cancelled by the client, but the DDL work hasn't handle it.
	JobStateCancelling JobState = 7
)

// String implements fmt.Stringer interface.
func (s JobState) String() string { _ = "STUB: not implemented"; return "" }

// SchemaDiff contains the schema modification at a particular schema version.
// It is used to reduce schema reload cost.
type SchemaDiff struct {
	Version  int64      `json:"version"`
	Type     ActionType `json:"type"`
	SchemaID int64      `json:"schema_id"`
	TableID  int64      `json:"table_id"`

	// OldTableID is the table ID before truncate, only used by truncate table DDL.
	OldTableID int64 `json:"old_table_id"`
	// OldSchemaID is the schema ID before rename table, only used by rename table DDL.
	OldSchemaID int64 `json:"old_schema_id"`

	AffectedOpts []*AffectedOption `json:"affected_options"`
}

// AffectedOption is used when a ddl affects multi tables.
type AffectedOption struct {
	SchemaID    int64 `json:"schema_id"`
	TableID     int64 `json:"table_id"`
	OldTableID  int64 `json:"old_table_id"`
	OldSchemaID int64 `json:"old_schema_id"`
}
