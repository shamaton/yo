// Package {{ .Package }} contains the types for schema '{{ .Schema }}'.
package {{ .Package }}

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

