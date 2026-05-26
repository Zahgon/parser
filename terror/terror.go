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

package terror

import (
	"sync"

	"github.com/pingcap/errors"
	"github.com/pingcap/parser/mysql"
)

// ErrCode represents a specific error type in a error class.
// Same error code can be used in different error classes.
type ErrCode int

const (
	// Executor error codes.

	// CodeUnknown is for errors of unknown reason.
	CodeUnknown ErrCode = -1
	// CodeExecResultIsEmpty indicates execution result is empty.
	CodeExecResultIsEmpty ErrCode = 3

	// Expression error codes.

	// CodeMissConnectionID indicates connection id is missing.
	CodeMissConnectionID ErrCode = 1

	// Special error codes.

	// CodeResultUndetermined indicates the sql execution result is undetermined.
	CodeResultUndetermined ErrCode = 2
)

// ErrClass represents a class of errors.
type ErrClass int

type Error = errors.Error

// Error classes.
var (
	ClassAutoid     = RegisterErrorClass(1, "autoid")
	ClassDDL        = RegisterErrorClass(2, "ddl")
	ClassDomain     = RegisterErrorClass(3, "domain")
	ClassEvaluator  = RegisterErrorClass(4, "evaluator")
	ClassExecutor   = RegisterErrorClass(5, "executor")
	ClassExpression = RegisterErrorClass(6, "expression")
	ClassAdmin      = RegisterErrorClass(7, "admin")
	ClassKV         = RegisterErrorClass(8, "kv")
	ClassMeta       = RegisterErrorClass(9, "meta")
	ClassOptimizer  = RegisterErrorClass(10, "planner")
	ClassParser     = RegisterErrorClass(11, "parser")
	ClassPerfSchema = RegisterErrorClass(12, "perfschema")
	ClassPrivilege  = RegisterErrorClass(13, "privilege")
	ClassSchema     = RegisterErrorClass(14, "schema")
	ClassServer     = RegisterErrorClass(15, "server")
	ClassStructure  = RegisterErrorClass(16, "structure")
	ClassVariable   = RegisterErrorClass(17, "variable")
	ClassXEval      = RegisterErrorClass(18, "xeval")
	ClassTable      = RegisterErrorClass(19, "table")
	ClassTypes      = RegisterErrorClass(20, "types")
	ClassGlobal     = RegisterErrorClass(21, "global")
	ClassMockTikv   = RegisterErrorClass(22, "mocktikv")
	ClassJSON       = RegisterErrorClass(23, "json")
	ClassTiKV       = RegisterErrorClass(24, "tikv")
	ClassSession    = RegisterErrorClass(25, "session")
	ClassPlugin     = RegisterErrorClass(26, "plugin")
	ClassUtil       = RegisterErrorClass(27, "util")
	// Add more as needed.
)

var errClass2Desc = make(map[ErrClass]string)
var rfcCode2errClass = newCode2ErrClassMap()

type code2ErrClassMap struct {
	data sync.Map
}

func newCode2ErrClassMap() *code2ErrClassMap { _ = "STUB: not implemented"; return nil }

func (m *code2ErrClassMap) Get(key string) (ErrClass, bool) {
	_ = "STUB: not implemented"
	return *new(ErrClass), false
}

func (m *code2ErrClassMap) Put(key string, err ErrClass) { _ = "STUB: not implemented"; return }

var registerFinish uint32

// RegisterFinish makes the register of new error panic.
// The use pattern should be register all the errors during initialization, and then call RegisterFinish.
func RegisterFinish() { _ = "STUB: not implemented"; return }

func frozen() bool { _ = "STUB: not implemented"; return false }

// RegisterErrorClass registers new error class for terror.
func RegisterErrorClass(classCode int, desc string) ErrClass {
	_ = "STUB: not implemented"
	return *new(ErrClass)
}

// String implements fmt.Stringer interface.
func (ec ErrClass) String() string { _ = "STUB: not implemented"; return "" }

// EqualClass returns true if err is *Error with the same class.
func (ec ErrClass) EqualClass(err error) bool { _ = "STUB: not implemented"; return false }

// NotEqualClass returns true if err is not *Error with the same class.
func (ec ErrClass) NotEqualClass(err error) bool { _ = "STUB: not implemented"; return false }

func (ec ErrClass) initError(code ErrCode) string { _ = "STUB: not implemented"; return "" }

// New defines an *Error with an error code and an error message.
// Usually used to create base *Error.
// Attention:
// this method is not goroutine-safe and
// usually be used in global variable initializer
//
// Deprecated: use NewStd or NewStdErr instead.
func (ec ErrClass) New(code ErrCode, message string) *Error { _ = "STUB: not implemented"; return nil }

// NewStdErr defines an *Error with an error code, an error
// message and workaround to create standard error.
func (ec ErrClass) NewStdErr(code ErrCode, message *mysql.ErrMessage) *Error {
	_ = "STUB: not implemented"
	return nil
}

// NewStd calls New using the standard message for the error code
// Attention:
// this method is not goroutine-safe and
// usually be used in global variable initializer
func (ec ErrClass) NewStd(code ErrCode) *Error { _ = "STUB: not implemented"; return nil }

// Synthesize synthesizes an *Error in the air
// it didn't register error into ErrClassToMySQLCodes
// so it's goroutine-safe
// and often be used to create Error came from other systems like TiKV.
func (ec ErrClass) Synthesize(code ErrCode, message string) *Error {
	_ = "STUB: not implemented"
	return nil
}

// ToSQLError convert Error to mysql.SQLError.
func ToSQLError(e *Error) *mysql.SQLError { _ = "STUB: not implemented"; return nil }

var defaultMySQLErrorCode uint16

func getMySQLErrorCode(e *Error) uint16 { _ = "STUB: not implemented"; return 0 }

var (
	// ErrClassToMySQLCodes is the map of ErrClass to code-set.
	ErrClassToMySQLCodes  = make(map[ErrClass]map[ErrCode]struct{})
	ErrCritical           = ClassGlobal.NewStdErr(CodeExecResultIsEmpty, mysql.Message("critical error %v", nil))
	ErrResultUndetermined = ClassGlobal.NewStdErr(CodeResultUndetermined, mysql.Message("execution result undetermined", nil))
)

func init() {
	defaultMySQLErrorCode = mysql.ErrUnknown
}

// ErrorEqual returns a boolean indicating whether err1 is equal to err2.
func ErrorEqual(err1, err2 error) bool { _ = "STUB: not implemented"; return false }

// ErrorNotEqual returns a boolean indicating whether err1 isn't equal to err2.
func ErrorNotEqual(err1, err2 error) bool { _ = "STUB: not implemented"; return false }

// MustNil cleans up and fatals if err is not nil.
func MustNil(err error, closeFuns ...func()) { _ = "STUB: not implemented"; return }

// Call executes a function and checks the returned err.
func Call(fn func() error) { _ = "STUB: not implemented"; return }

// Log logs the error if it is not nil.
func Log(err error) { _ = "STUB: not implemented"; return }

func GetErrClass(e *Error) ErrClass { _ = "STUB: not implemented"; return *new(ErrClass) }
