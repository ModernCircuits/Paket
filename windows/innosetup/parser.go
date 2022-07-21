package innosetup

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func parseSetupSection(lines []string) (*SetupSection, error) {
	m, err := parseLinesToMap(lines)
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
	set := false
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		if varName != tag {
			continue
		}

		varType := e.Type().Field(i).Type
		set = true
		switch varType.Kind() {
		case reflect.String:
			e.Field(i).SetString(value)
		case reflect.Bool:
			e.Field(i).SetBool(value == "yes")
		default:
			return fmt.Errorf("unimplemented type %v", varType)
		}
	}

	if !set {
		return fmt.Errorf("unimplemented tag %q", tag)
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

func splitSectionLines(lines []string) map[string][]string {
	sections := map[string][]string{}
	key := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			key = strings.Replace(line, "[", "", 1)
			key = strings.Replace(key, "]", "", 1)
			sections[key] = []string{}
			continue
		}

		sections[key] = append(sections[key], line)
	}

	return sections
}

func getSetupSectionLines(lines []string) ([]string, error) {
	sectionLines, ok := splitSectionLines(lines)["Setup"]
	if !ok {
		return nil, fmt.Errorf("no [Setup] section found")
	}
	return sectionLines, nil
}
