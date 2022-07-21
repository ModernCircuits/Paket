package innosetup

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func parseSetupSection(lines []string) (*SetupSection, error) {
	setupSectionLines, err := getSetupSectionLines(lines)
	if err != nil {
		return nil, err
	}

	m, err := parseLinesToMap(setupSectionLines[1:])
	if err != nil {
		return nil, err
	}

	setup := &SetupSection{}
	for k, v := range m {
		if err := setSetupStructField(setup, k, v); err != nil {
			return nil, err
		}
	}

	return setup, nil
}

func setSetupStructField(setup *SetupSection, tag string, value string) error {
	e := reflect.ValueOf(setup).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		if varName != tag {
			continue
		}

		varType := e.Type().Field(i).Type
		switch varType.Kind() {
		case reflect.String:
			e.Field(i).SetString(value)
		case reflect.Bool:
			e.Field(i).SetBool(value == "yes")
		default:
			return fmt.Errorf("unimplemented type %v", varType)
		}
	}
	return nil
}

func parseLinesToMap(lines []string) (map[string]string, error) {
	m := map[string]string{}

	for _, line := range lines {
		nameAndValue := strings.Split(line, "=")
		if len(nameAndValue) != 2 {
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		name := nameAndValue[0]
		if name == "" {
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		value := nameAndValue[1]
		if value == "" {
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		m[name] = value
	}

	return m, nil
}

func readAllLines(r io.Reader) []string {
	lines := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, strings.Trim(scanner.Text(), " \t\r\n"))
	}
	return lines
}

func removeAllCommentLines(lines []string) []string {
	cleaned := []string{}
	for _, line := range lines {
		if isCommentLine(line) {
			continue
		}
		if isBlankLine(line) {
			continue
		}
		cleaned = append(cleaned, line)
	}
	return cleaned
}

func isCommentLine(line string) bool {
	return strings.HasPrefix(line, ";")
}

func isBlankLine(line string) bool {
	return line == ""
}

func findSetupStartIndex(lines []string) int {
	for i, line := range lines {
		if strings.Contains(line, "[Setup]") {
			return i
		}
	}
	return -1
}

func findSetupEndIndex(lines []string, startIndex int) int {
	for i, line := range lines {
		if strings.Contains(line, "[Files]") {
			return i + startIndex
		}
	}
	return -1
}

func getSetupSectionLines(lines []string) ([]string, error) {
	start := findSetupStartIndex(lines)
	if start == -1 {
		return nil, fmt.Errorf("setup section start not found")
	}

	end := findSetupEndIndex(lines, start)
	if end == -1 {
		return nil, fmt.Errorf("setup section end not found")
	}

	return lines[start:end], nil
}
