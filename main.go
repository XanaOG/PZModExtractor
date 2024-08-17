package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "/mnt/f/SteamLibrary/steamapps/workshop/content/108600"
	modIDList := []string{}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, "mod.info") {
			id, err := extractModID(path)
			if err == nil && id != "" {
				fmt.Printf("Processing mod ID: %s from file: %s\n", id, path)
				modIDList = append(modIDList, id)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", rootDir, err)
		return
	}

	outputFilePath := "./mod_list.txt"
	err = writeModListToFile(modIDList, outputFilePath)
	if err != nil {
		fmt.Printf("Error writing mod list to file: %v\n", err)
	} else {
		fmt.Println("Mod list successfully written to", outputFilePath)
	}
}

func extractModID(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "id=") {
			return strings.TrimPrefix(line, "id="), nil
		}
	}

	return "", scanner.Err()
}

func writeModListToFile(modIDList []string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString("# Enter the mod loading ID here. It can be found in \\Steam\\steamapps\\workshop\\modID\\mods\\modName\\info.txt\n")
	writer.WriteString("Mods=")
	writer.WriteString(strings.Join(modIDList, ";"))
	writer.WriteString("\n")

	return nil
}
