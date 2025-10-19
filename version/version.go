package version

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//go:embed version.json
var versionData []byte

// Version defines the version of the Aerium software.
// It follows the Semantic Versioning 2.0.0 spec: http://semver.org/.
//
// Update this struct with each new release by adjusting the Major, Minor, or Patch numbers.
// For major releases, clear the Meta field (set to an empty string).
// Use the Meta field to indicate pre-release stages, such as "rc1", "rc2", or "beta" during development.
type Version struct {
	Major uint   `json:"major"` // Major version number
	Minor uint   `json:"minor"` // Minor version number
	Patch uint   `json:"patch"` // Patch version number
	Meta  string `json:"meta"`  // Metadata for version (e.g., "beta", "rc1")
	Alias string `json:"alias"` // Alias for version (e.g., "London")
}

var nodeVersion *Version

// NodeVersion represents the current version of the node software.
func NodeVersion() Version {
	if nodeVersion == nil {
		// Initialize the version from the embedded version.json.
		nodeVersion = new(Version)
		if err := json.Unmarshal(versionData, nodeVersion); err != nil {
			panic(err)
		}
	}

	return *nodeVersion
}

// ParseVersion parses a version string into a Version struct.
// The format should be "Major.Minor.Patch-Meta", where Meta is optional.
// Returns the parsed Version struct and an error if parsing fails.
func ParseVersion(versionStr string) (Version, error) {
	var ver Version

	if versionStr == "" {
		return ver, errors.New("empty version string")
	}

	if versionStr[0] == 'v' {
		versionStr = versionStr[1:]
	}

	parts := strings.Split(versionStr, ".")
	if len(parts) != 3 {
		return ver, errors.New("invalid version format")
	}

	parseUintPart := func(part string, name string) (uint, error) {
		val, err := strconv.ParseUint(part, 10, 0)
		if err != nil {
			return 0, fmt.Errorf("failed to parse %s version: %w", name, err)
		}

		return uint(val), nil
	}
	var err error
	if ver.Major, err = parseUintPart(parts[0], "Major"); err != nil {
		return ver, err
	}
	if ver.Minor, err = parseUintPart(parts[1], "Minor"); err != nil {
		return ver, err
	}

	patchMeta := strings.Split(parts[2], "-")
	if len(patchMeta) > 2 {
		return ver, errors.New("invalid Patch and Meta format")
	}

	if ver.Patch, err = parseUintPart(patchMeta[0], "Patch"); err != nil {
		return ver, err
	}

	if len(patchMeta) == 2 {
		ver.Meta = patchMeta[1]
	}

	return ver, nil
}

// StringWithAlias returns a string representation of the Version object with the alias.
func (v Version) StringWithAlias() string {
	if v.Alias == "" {
		return v.String()
	}

	return fmt.Sprintf("%s (%s)", v.String(), v.Alias)
}

// String returns a string representation of the Version object.
func (v Version) String() string {
	ver := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Meta != "" {
		ver = fmt.Sprintf("%s-%s", ver, v.Meta)
	}

	return ver
}

// Compare compares the current version (v) with another version (other)
// and returns:
//
//	-1 if v < other
//	 0 if v == other
//	 1 if v > other
func (v Version) Compare(other Version) int {
	if v.Major != other.Major {
		return compareInt(v.Major, other.Major)
	}
	if v.Minor != other.Minor {
		return compareInt(v.Minor, other.Minor)
	}

	return compareInt(v.Patch, other.Patch)
}

func compareInt(a, b uint) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}

	return 0
}

var _ = NodeVersion()
