package utils

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func LoadPolicies() []SPIF {
	dirPath := "config/spifs/s4774/"
	var loadedSPIFs []SPIF
	// Open the directory
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file has an .xml extension
		if !info.IsDir() && filepath.Ext(info.Name()) == ".xml" {
			fmt.Println("Found XML file:", path, "- Loading it")
			file, err := os.Open(path)
			if err != nil {
				fmt.Println("Error opening file:", err)

			}
			defer file.Close()
			// Parse the XML
			var spif SPIF
			err = xml.NewDecoder(file).Decode(&spif)
			if err != nil {
				fmt.Println("Error decoding XML:", err)
			}
			loadedSPIFs = append(loadedSPIFs, spif)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
	}

	return loadedSPIFs
}

func FindPolicy(spifs []SPIF, policy string) (SPIF, error) {
	for _, spif := range spifs {
		// Look the the right spif
		if spif.SecurityPolicyID.Name == policy {
			return spif, nil
		}
	}
	return SPIF{}, errors.New("policy not found")
}

func GetPolicies(spifs []SPIF) []string {
	var policies []string
	for _, spif := range spifs {
		policies = append(policies, spif.SecurityPolicyID.Name)
	}
	return policies
}
