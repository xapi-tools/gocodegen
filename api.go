package gocodegen

import "strings"

type GoCodeWriter interface {
	ToStringBuilder() (*strings.Builder, error)
	ToFile(string) error
}
