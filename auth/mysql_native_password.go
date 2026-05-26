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

package auth

// CheckScrambledPassword check scrambled password received from client.
// The new authentication is performed in following manner:
//
//	SERVER:  public_seed=create_random_string()
//	         send(public_seed)
//	CLIENT:  recv(public_seed)
//	         hash_stage1=sha1("password")
//	         hash_stage2=sha1(hash_stage1)
//	         reply=xor(hash_stage1, sha1(public_seed,hash_stage2)
//	         // this three steps are done in scramble()
//	         send(reply)
//	SERVER:  recv(reply)
//	         hash_stage1=xor(reply, sha1(public_seed,hash_stage2))
//	         candidate_hash2=sha1(hash_stage1)
//	         check(candidate_hash2==hash_stage2)
//	         // this three steps are done in check_scramble()
func CheckScrambledPassword(salt, hpwd, auth []byte) bool { _ = "STUB: not implemented"; return false }

// token = scrambleHash XOR stage1Hash

// Sha1Hash is an util function to calculate sha1 hash.
func Sha1Hash(bs []byte) []byte { _ = "STUB: not implemented"; return nil }

// EncodePassword converts plaintext password to hashed hex string.
func EncodePassword(pwd string) string { _ = "STUB: not implemented"; return "" }

// DecodePassword converts hex string password without prefix '*' to byte array.
func DecodePassword(pwd string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
