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
	"fmt"
	"log"

	"github.com/ibm-hyper-protect/contract-cli/common"
	"github.com/ibm-hyper-protect/contract-go/contract"
	"github.com/spf13/cobra"
)

const (
	inputMissingMessageBase64 = "input data is missing"
	invalidInputMessageBase64 = "invalid input format"
	successMessageBase64      = "successfully generated Base64"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   common.Base64ParamName,
	Short: common.Base64ParamShortDescription,
	Long:  common.Base64ParamLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		inputData, formatType, outputPath, err := validateInputBase64(cmd)
		if err != nil {
			log.Fatal(err)
		}

		base64String, err := processBase64(inputData, formatType)
		if err != nil {
			log.Fatal(err)
		}

		err = printBase64(base64String, outputPath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// init - cobra init function
func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.PersistentFlags().String(common.FileInFlagName, "", common.Base64InputFlagDescription)
	base64Cmd.PersistentFlags().String(common.DataFormatFlagName, common.DataFormatText, common.Base64InputFormatFlagDescription)
	base64Cmd.PersistentFlags().String(common.FileOutFlagName, "", common.Base64OutputPathFlagDescription)
}

// validateInputBase64 - function to validate base64 command input
func validateInputBase64(cmd *cobra.Command) (string, string, string, error) {
	inputData, err := cmd.Flags().GetString(common.FileInFlagName)
	if err != nil {
		return "", "", "", err
	}

	formatType, err := cmd.Flags().GetString(common.DataFormatFlagName)
	if err != nil {
		return "", "", "", err
	}

	outputPath, err := cmd.Flags().GetString(common.FileOutFlagName)
	if err != nil {
		return "", "", "", err
	}

	return inputData, formatType, outputPath, nil
}

// processBase64 - function to process base64 command inputs
func processBase64(inputData, formatType string) (string, error) {
	var base64String string
	var err error

	if inputData == "" {
		return "", fmt.Errorf(inputMissingMessageBase64)
	}

	if formatType == common.DataFormatText {
		base64String, _, _, err = contract.HpcrText(inputData)
		if err != nil {
			return "", err
		}
	} else if formatType == common.DataFormatJson {
		base64String, _, _, err = contract.HpcrJson(inputData)
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf(invalidInputMessageBase64)
	}

	return base64String, nil
}

// printBase64 - function to print base64 string or redirect to a file
func printBase64(base64String, outputPath string) error {
	if outputPath != "" {
		err := common.WriteDataToFile(outputPath, base64String)
		if err != nil {
			return err
		}
		fmt.Println(successMessageBase64)
	} else {
		fmt.Println(base64String)
	}

	return nil
}
