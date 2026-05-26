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

// HasAggFlag checks if the expr contains FlagHasAggregateFunc.
func HasAggFlag(expr ExprNode) bool { _ = "STUB: not implemented"; return false }

func HasWindowFlag(expr ExprNode) bool { _ = "STUB: not implemented"; return false }

// SetFlag sets flag for expression.
func SetFlag(n Node) { _ = "STUB: not implemented"; return }

type flagSetter struct {
}

func (f *flagSetter) Enter(in Node) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (f *flagSetter) Leave(in Node) (Node, bool) {
	_ = "STUB: not implemented"
	return *new(Node), false
}

func (f *flagSetter) caseExpr(x *CaseExpr) { _ = "STUB: not implemented"; return }

func (f *flagSetter) patternIn(x *PatternInExpr) { _ = "STUB: not implemented"; return }

func (f *flagSetter) patternLike(x *PatternLikeExpr) { _ = "STUB: not implemented"; return }

func (f *flagSetter) patternRegexp(x *PatternRegexpExpr) { _ = "STUB: not implemented"; return }

func (f *flagSetter) row(x *RowExpr) { _ = "STUB: not implemented"; return }

func (f *flagSetter) funcCall(x *FuncCallExpr) { _ = "STUB: not implemented"; return }

func (f *flagSetter) aggregateFunc(x *AggregateFuncExpr) { _ = "STUB: not implemented"; return }

func (f *flagSetter) windowFunc(x *WindowFuncExpr) { _ = "STUB: not implemented"; return }
