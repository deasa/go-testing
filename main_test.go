package main

import "testing"

func BenchmarkGetPathToDirectory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetPathToDirectory("./integrationtests", "data-seeding-api")
	}
}

func BenchmarkGetPathToDirectory_WalkDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetPathToDirectory_WalkDir("./integrationtests", "data-seeding-api")
	}
}
