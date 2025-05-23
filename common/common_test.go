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

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sampleFile = "../samples/attestation/se-checksums.txt.enc"

	simpleSampleTextPath = "../samples/simple_file.txt"
	simpleSampleText     = "Testing"

	simpleSampleWritePath = "../build/simple_file.txt"

	samplePrivateKeyPath = "../samples/sign/private.pem"

	sampleCertPath = "../samples/contract-expiry/personal_ca.crt"
)

// Testcase to check if CheckFileFolderExists() is able to check if a folder exists
func TestCheckFileFolderExists(t *testing.T) {
	result := CheckFileFolderExists(sampleFile)

	assert.True(t, result)
}

// Testcase to check if ReadDataFromFile() is able to read data from file
func TestReadDataFromFile(t *testing.T) {
	content, err := ReadDataFromFile(simpleSampleTextPath)
	if err != nil {
		t.Errorf("failed to read text from file - %v", err)
	}

	assert.Equal(t, content, simpleSampleText)
}

// Testcase to check if WriteDataToFile() is able to write data to file
func TestWriteDataToFile(t *testing.T) {
	err := WriteDataToFile(simpleSampleWritePath, simpleSampleText)
	if err != nil {
		t.Errorf("failed to write data to file - %v", err)
	}
}

// Testcase to check if ExecCommand() is able to execute command
func TestExecCommand(t *testing.T) {
	_, err := ExecCommand("openssl", "", "version")
	if err != nil {
		t.Errorf("failed to execute command - %v", err)
	}
}

// Testcase to check if OpensslCheck() is able to check if openssl is present
func TestOpensslCheck(t *testing.T) {
	err := OpensslCheck()
	if err != nil {
		t.Errorf("openssl check failed - %v", err)
	}
}

// Testcase to check if GetPrivateKey() is able to get key from file
func TestGetPrivateKeyNoKey(t *testing.T) {
	result, err := GetPrivateKey(samplePrivateKeyPath)
	if err != nil {
		t.Errorf("failed to get private key - %v", err)
	}

	assert.NotEmpty(t, result)
}

// Testcase to check if GetPrivateKey() is able generate private key
func TestGetPrivateKey(t *testing.T) {
	result, err := GetPrivateKey("")
	if err != nil {
		t.Errorf("failed to get private key - %v", err)
	}

	assert.NotEmpty(t, result)
}

// Testcase to check if generatePrivateKey() is able to generate private key
func TestGeneratePrivateKey(t *testing.T) {
	result, err := generatePrivateKey()
	if err != nil {
		t.Errorf("failed to generate private key - %v", err)
	}

	assert.NotEmpty(t, result)
}

// Testcase to check if GetDataFromFile() is able to get data form file
func TestGetDataFromFileWithData(t *testing.T) {
	result, err := GetDataFromFile(sampleCertPath)
	if err != nil {
		t.Errorf("failed to get data from file - %v", err)
	}

	assert.NotEmpty(t, result)
}

// Testcase to check if GetDataFromFile() is throwing error for invalid file path
func TestGetDataFromFileWithoutData(t *testing.T) {
	result, err := GetDataFromFile("")
	if err != nil {
		t.Errorf("failed to get data from file - %v", err)
	}

	assert.Empty(t, result)
}
