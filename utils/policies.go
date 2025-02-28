package utils

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sparrow/structures"
)

func LoadPolicies() []structures.SPIF {
	dirPath := "config/spifs/"
	var loadedSPIFs []structures.SPIF
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
			var spif structures.SPIF
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

func FindPolicy(spifs []structures.SPIF, policy string) (structures.SPIF, error) {
	for _, spif := range spifs {
		// Look the the right spif
		if spif.SecurityPolicyId.Name == policy {
			return spif, nil
		}
	}
	return structures.SPIF{}, errors.New("policy not found")
}

func GetPolicies(spifs []structures.SPIF) []string {
	var policies []string
	for _, spif := range spifs {
		policies = append(policies, spif.SecurityPolicyId.Name)
	}
	return policies
}
