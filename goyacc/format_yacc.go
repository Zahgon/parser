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

package main

import (
	"bufio"
	"os"
	"regexp"

	parser "github.com/cznic/parser/yacc"
	"github.com/cznic/strutil"
	"github.com/pingcap/parser/format"
)

func Format(inputFilename string, goldenFilename string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseFileToSpec(inputFilename string) (*parser.Specification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Definition represents data reduced by productions:
//
//	Definition:
//	        START IDENTIFIER
//	|       UNION                      // Case 1
//	|       LCURL RCURL                // Case 2
//	|       ReservedWord Tag NameList  // Case 3
//	|       ReservedWord Tag           // Case 4
//	|       ERROR_VERBOSE              // Case 5
const (
	StartIdentifierCase = iota
	UnionDefinitionCase
	LCURLRCURLCase
	ReservedWordTagNameListCase
	ReservedWordTagCase
)

func printDefinitions(formatter format.Formatter, definitions []*parser.Definition) error {
	_ = "STUB: not implemented"
	return nil
}

func handleStart(f format.Formatter, definition *parser.Definition) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUnion(f format.Formatter, definition *parser.Definition) error {
	_ = "STUB: not implemented"
	return nil
}

func handleProlog(f format.Formatter, definition *parser.Definition) error {
	_ = "STUB: not implemented"
	return nil
}

func handleReservedWordTagNameList(f format.Formatter, def *parser.Definition) error {
	_ = "STUB: not implemented"
	return nil
}

func joinTag(tag *parser.Tag) string { _ = "STUB: not implemented"; return "" }

type stringLayout int8

const (
	spanStringLayout stringLayout = iota
	divStringLayout
	divNewLineStringLayout
)

func getTokenComment(token *parser.Token, layout stringLayout) string {
	_ = "STUB: not implemented"
	return ""
}

func printNameListVertical(f format.Formatter, names NameArr) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func joinNames(names NameArr) string { _ = "STUB: not implemented"; return "" }

func printSingleName(f format.Formatter, name *parser.Name, maxCharLength int) error {
	_ = "STUB: not implemented"
	return nil
}

type NameArr []*parser.Name

func (ns NameArr) span(pred func(*parser.Name) bool) (NameArr, NameArr) {
	_ = "STUB: not implemented"
	return *new(NameArr), *new(NameArr)
}

func (ns NameArr) takeWhile(pred func(*parser.Name) bool) NameArr {
	_ = "STUB: not implemented"
	return *new(NameArr)
}

func (ns NameArr) findMaxLength() int { _ = "STUB: not implemented"; return 0 }

func hasComments(n *parser.Name) bool { _ = "STUB: not implemented"; return false }

func noComment(n *parser.Name) bool { _ = "STUB: not implemented"; return false }

func containsActionInRule(rule *parser.Rule) bool { _ = "STUB: not implemented"; return false }

type RuleArr []*parser.Rule

func printRules(f format.Formatter, rules RuleArr) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type ruleItemType int8

const (
	identRuleItemType      ruleItemType = 1
	actionRuleItemType     ruleItemType = 2
	strLiteralRuleItemType ruleItemType = 3
)

func printRuleBody(f format.Formatter, rule *parser.Rule) error {
	_ = "STUB: not implemented"
	return nil
}

func handleAction(f format.Formatter, rule *parser.Rule, action *parser.Action, isFirstItem bool) error {
	_ = "STUB: not implemented"
	return nil
}

func handlePrecedence(f format.Formatter, p *parser.Precedence, isFirstItem bool) error {
	_ = "STUB: not implemented"
	return nil
}

func formatGoSnippet(actVal []*parser.ActionValue) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func collectGoSnippet(tran *SpecialActionValTransformer, actionValArr []*parser.ActionValue) string {
	_ = "STUB: not implemented"
	return ""
}

var lineBeginBlankRegex = regexp.MustCompile("(?m)^[\t ]+")

func removeLineBeginBlanks(src string) string { _ = "STUB: not implemented"; return "" }

type SpecialActionValTransformer struct {
	store map[string]string
}

const yaccFmtVar = "_yaccfmt_var_"

var yaccFmtVarRegex = regexp.MustCompile("_yaccfmt_var_[0-9]{1,5}")

func (s *SpecialActionValTransformer) transform(val string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SpecialActionValTransformer) restore(src string) string {
	_ = "STUB: not implemented"
	return ""
}

type OutputFormatter struct {
	file      *os.File
	out       *bufio.Writer
	formatter strutil.Formatter
}

func (y *OutputFormatter) Setup(filename string) (err error) { _ = "STUB: not implemented"; return nil }

func (y *OutputFormatter) Teardown() error { _ = "STUB: not implemented"; return nil }

func (y *OutputFormatter) Format(format string, args ...interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (y *OutputFormatter) Write(bytes []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type NotNilAssert struct {
	idx int
	err error
}

func (n *NotNilAssert) and(target interface{}) *NotNilAssert { _ = "STUB: not implemented"; return nil }

func (n *NotNilAssert) NotNil() error { _ = "STUB: not implemented"; return nil }

func Ensure(target interface{}) *NotNilAssert { _ = "STUB: not implemented"; return nil }

func escapePercent(src string) string { _ = "STUB: not implemented"; return "" }

func checkInconsistencyInYaccParser(f format.Formatter, rule *parser.Rule, counter int) error {
	_ = "STUB: not implemented"
	return nil
}

// pickup rule item in ruleBody
