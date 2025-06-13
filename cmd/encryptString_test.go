// Copyright (c) 2025 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"testing"

	"github.com/ibm-hyper-protect/contract-cli/common"
	"github.com/stretchr/testify/assert"
)

const (
	sampleEncryptStringInput = "testing"
	sampleEncryptStringJson  = `{"type":"workload"}`

	sampleEncryptStringOutputPlain = "../build/encrypt_string.txt"
	sampleEncryptStringOutputJson  = "../build/encrypt_string.json"
)

var (
	sampleEncryptStringValidCommandText = []string{common.EncryptStrParamName, "--in", sampleEncryptStringInput, "--format", common.DataFormatText, "--out", sampleEncryptStringOutputPlain}
	sampleEncryptStringValidCommandJson = []string{common.EncryptStrParamName, "--in", sampleEncryptStringJson, "--format", common.DataFormatJson, "--out", sampleEncryptStringOutputJson}
)

// Testcase to check if encrypt-string is able to generate encrypted string of plain text
func TestEncryptStringCmdText(t *testing.T) {
	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs(sampleEncryptStringValidCommandText)
	err := encryptStringCmd.Execute()

	assert.NoError(t, err)
}

// Testcase to check if encrypt-string is able to generate encrypted string of JSON
func TestEncryptStringCmdJson(t *testing.T) {
	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs(sampleEncryptStringValidCommandJson)
	err := encryptStringCmd.Execute()

	assert.NoError(t, err)
}
