package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	resultLines = "*3\r\n$3\r\nSET\r\n$5\r\nprice\r\n$5\r\n99.99\r\n*3\r\n$3\r\nSET\r\n$5\r\ncolor\r\n$3\r\nred\r\n*3\r\n$3\r\nSET\r\n$4\r\nunit\r\n$7\r\nCelsius\r\n"
)

var (
	fileSourceLines = []string{
		"SET price 99.99",
		"SET color red",
		"SET unit Celsius",
	}
)

func TestReadLines(t *testing.T) {
	lines, err := readLines("sample_source.txt")
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, fileSourceLines, lines, "They should be equal")
}

func TestExists(t *testing.T) {
	assert.True(t, exists("./README.md"), "README.md should exist")
	assert.False(t, exists("./ABCDEFG"), "This file should not exist")
}

func TestGenerateRedisScript(t *testing.T) {
	var script = generateRedisScript(fileSourceLines)

	assert.Equal(t, resultLines, script, "They should be equal")
}
