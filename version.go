package grammalecte

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	// ErrVersionNotFound signals no version assignment was detected.
	ErrVersionNotFound = errors.New("version not found")

	dunderVersionRe = regexp.MustCompile(`^\s*__version__\s*=\s*["']([^"']+)["']`)
	versionRe       = regexp.MustCompile(`^\s*version\s*=\s*["']([^"']+)["']`)
)

// VersionFromPythonFile extracts the version string from a Python source file.
func VersionFromPythonFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return VersionFromPythonReader(file)
}

// VersionFromPythonSource extracts the version string from Python source text.
func VersionFromPythonSource(source string) (string, error) {
	return VersionFromPythonReader(strings.NewReader(source))
}

// VersionFromPythonReader extracts the version string from a Python reader.
func VersionFromPythonReader(reader io.Reader) (string, error) {
	lines, err := readLines(reader)
	if err != nil {
		return "", err
	}

	if version := findVersion(lines, dunderVersionRe); version != "" {
		return version, nil
	}
	if version := findVersion(lines, versionRe); version != "" {
		return version, nil
	}

	return "", ErrVersionNotFound
}

func readLines(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func findVersion(lines []string, regex *regexp.Regexp) string {
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if match := regex.FindStringSubmatch(line); match != nil {
			return match[2]
		}
	}
	return ""
}
