// Package utils 工具集
package utils

type PreLevel string

const (
	AREA     PreLevel = "区"
	CITY     PreLevel = "市"
	PROVINCE PreLevel = "省"
)

type Email string

const (
	EMAIL_SYMBOL Email = "@"
)

type GeneratorType int

const (
	ALL_TYPE GeneratorType = 1
	NUMBER   GeneratorType = 2
)
