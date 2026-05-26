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
	"github.com/pingcap/parser/types"
)

// node is the struct implements Node interface except for Accept method.
// Node implementations should embed it in.
type node struct {
	text   string
	offset int
}

// SetOriginTextPosition implements Node interface.
func (n *node) SetOriginTextPosition(offset int) {
	_ = "STUB: not implemented"

	// OriginTextPosition implements Node interface.
	return
}

func (n *node) OriginTextPosition() int {
	_ = "STUB: not implemented"

	// SetText implements Node interface.
	return 0
}

func (n *node) SetText(text string) {
	_ = "STUB: not implemented"

	// Text implements Node interface.
	return
}

func (n *node) Text() string {
	_ = "STUB: not implemented"

	// stmtNode implements StmtNode interface.
	// Statement implementations should embed it in.
	return ""
}

type stmtNode struct {
	node
}

// statement implements StmtNode interface.
func (sn *stmtNode) statement() {
	_ = "STUB: not implemented"

	// ddlNode implements DDLNode interface.
	// DDL implementations should embed it in.
	return
}

type ddlNode struct {
	stmtNode
}

// ddlStatement implements DDLNode interface.
func (dn *ddlNode) ddlStatement() {
	_ = "STUB: not implemented"

	// dmlNode is the struct implements DMLNode interface.
	// DML implementations should embed it in.
	return
}

type dmlNode struct {
	stmtNode
}

// dmlStatement implements DMLNode interface.
func (dn *dmlNode) dmlStatement() {
	_ = "STUB: not implemented"

	// exprNode is the struct implements Expression interface.
	// Expression implementations should embed it in.
	return
}

type exprNode struct {
	node
	Type types.FieldType
	flag uint64
}

// TexprNode is exported for parser driver.
type TexprNode = exprNode

// SetType implements ExprNode interface.
func (en *exprNode) SetType(tp *types.FieldType) {
	_ = "STUB: not implemented"

	// GetType implements ExprNode interface.
	return
}

func (en *exprNode) GetType() *types.FieldType {
	_ = "STUB: not implemented"

	// SetFlag implements ExprNode interface.
	return nil
}

func (en *exprNode) SetFlag(flag uint64) {
	_ = "STUB: not implemented"

	// GetFlag implements ExprNode interface.
	return
}

func (en *exprNode) GetFlag() uint64 { _ = "STUB: not implemented"; return 0 }

type funcNode struct {
	exprNode
}

// functionExpression implements FunctionNode interface.
func (fn *funcNode) functionExpression() { _ = "STUB: not implemented"; return }
