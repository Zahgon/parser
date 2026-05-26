// Copyright 2021 PingCAP, Inc.
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

package auth

// Resources:
// - https://dev.mysql.com/doc/refman/8.0/en/caching-sha2-pluggable-authentication.html
// - https://dev.mysql.com/doc/dev/mysql-server/latest/page_caching_sha2_authentication_exchanges.html
// - https://dev.mysql.com/doc/dev/mysql-server/latest/namespacesha2__password.html
// - https://www.akkadia.org/drepper/SHA-crypt.txt
// - https://dev.mysql.com/worklog/task/?id=9591
//
// CREATE USER 'foo'@'%' IDENTIFIED BY 'foobar';
// SELECT HEX(authentication_string) FROM mysql.user WHERE user='foo';
// 24412430303524031A69251C34295C4B35167C7F1E5A7B63091349503974624D34504B5A424679354856336868686F52485A736E4A733368786E427575516C73446469496537
//
// Format:
// Split on '$':
// - digest type ("A")
// - iterations (divided by ITERATION_MULTIPLIER)
// - salt+hash
//

const (
	MIXCHARS             = 32
	SALT_LENGTH          = 20
	ITERATION_MULTIPLIER = 1000
)

func b64From24bit(b []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

func sha256crypt(plaintext string, salt []byte, iterations int) string {
	_ = "STUB: not implemented"
	// Numbers in the comments refer to the description of the algorithm on https://www.akkadia.org/drepper/SHA-crypt.txt
	return ""
}

// 1, 2, 3

// 4, 5, 6, 7, 8

// 9, 10

// 11

// 12

// 13, 14, 15

// 16

// 17, 18, 19

// 20

// 21

// 22

// FIXME

// Checks if a MySQL style caching_sha2 authentication string matches a password
func CheckShaPassword(pwhash []byte, password string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func NewSha2Password(pwd string) string { _ = "STUB: not implemented"; return "" }

// Restrict to 7-bit to avoid multi-byte UTF-8

// '$' or NUL
