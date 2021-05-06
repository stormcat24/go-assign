package example

import "embed"

var (
	//go:embed testdata/*
	TestDataDir embed.FS
)