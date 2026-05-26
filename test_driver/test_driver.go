// Copyright 2019 PingCAP, Inc.
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

//go:build !codes
// +build !codes

package test_driver

import (
	"io"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
)

func init() {
	ast.NewValueExpr = newValueExpr
	ast.NewParamMarkerExpr = newParamMarkerExpr
	ast.NewDecimal = func(str string) (interface{}, error) {
		dec := new(MyDecimal)
		err := dec.FromString([]byte(str))
		return dec, err
	}
	ast.NewHexLiteral = func(str string) (interface{}, error) {
		h, err := NewHexLiteral(str)
		return h, err
	}
	ast.NewBitLiteral = func(str string) (interface{}, error) {
		b, err := NewBitLiteral(str)
		return b, err
	}
}

var (
	_ ast.ParamMarkerExpr = &ParamMarkerExpr{}
	_ ast.ValueExpr       = &ValueExpr{}
)

// ValueExpr is the simple value expression.
type ValueExpr struct {
	ast.TexprNode
	Datum
	projectionOffset int
}

// Restore implements Node interface.
func (n *ValueExpr) Restore(ctx *format.RestoreCtx) error { _ = "STUB: not implemented"; return nil }

// TODO implement Restore function

// GetDatumString implements the ValueExpr interface.
func (n *ValueExpr) GetDatumString() string { _ = "STUB: not implemented"; return "" }

// Format the ExprNode into a Writer.
func (n *ValueExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// newValueExpr creates a ValueExpr with value, and sets default field type.
func newValueExpr(value interface{}, charset string, collate string) ast.ValueExpr {
	_ = "STUB: not implemented"
	return *new(ast.ValueExpr)
}

// SetProjectionOffset sets ValueExpr.projectionOffset for logical plan builder.
func (n *ValueExpr) SetProjectionOffset(offset int) { _ = "STUB: not implemented"; return }

// GetProjectionOffset returns ValueExpr.projectionOffset.
func (n *ValueExpr) GetProjectionOffset() int { _ = "STUB: not implemented"; return 0 }

// Accept implements Node interface.
func (n *ValueExpr) Accept(v ast.Visitor) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// ParamMarkerExpr expression holds a place for another expression.
// Used in parsing prepare statement.
type ParamMarkerExpr struct {
	ValueExpr
	Offset    int
	Order     int
	InExecute bool
}

// Restore implements Node interface.
func (n *ParamMarkerExpr) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func newParamMarkerExpr(offset int) ast.ParamMarkerExpr {
	_ = "STUB: not implemented"
	return *new(ast.ParamMarkerExpr)
}

// Format the ExprNode into a Writer.
func (n *ParamMarkerExpr) Format(w io.Writer) { _ = "STUB: not implemented"; return }

// Accept implements Node Accept interface.
func (n *ParamMarkerExpr) Accept(v ast.Visitor) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// SetOrder implements the ParamMarkerExpr interface.
func (n *ParamMarkerExpr) SetOrder(order int) { _ = "STUB: not implemented"; return }
