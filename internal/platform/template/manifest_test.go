package template

import (
	"path/filepath"
	"testing"
)

func TestRepositoryManifestIsValid(t *testing.T) {
	manifest, err := Load(filepath.Join("..", "..", "..", ".template", "manifest.json"))
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestManifestRejectsInvalidCommit(t *testing.T) {
	manifest := validManifest()
	manifest.TemplateCommit = "not-a-commit"

	if err := manifest.Validate(); err == nil {
		t.Fatal("expected invalid template commit to fail validation")
	}
}

func TestManifestAcceptsEmptySourceCommit(t *testing.T) {
	manifest := validManifest()
	manifest.TemplateCommit = ""

	if err := manifest.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func validManifest() Manifest {
	return Manifest{
		TemplateID:       "template-go-api",
		TemplateSource:   ExpectedTemplateSource,
		TemplateVersion:  "0.1.0",
		Repository:       ExpectedTemplateSource,
		License:          ExpectedLicense,
		MinimumGoVersion: ExpectedGoVersion,
		Database:         ExpectedDatabase,
		MigrationTool:    ExpectedMigrationTool,
		Architecture:     ExpectedArchitecture,
		SupportedEnvironments: []string{
			"development", "test", "production",
		},
		DependencyVersions: map[string]string{
			"github.com/golang-migrate/migrate/v4": "v4.19.1",
			"github.com/jackc/pgx/v5":              "v5.8.0",
		},
		Compatibility: Compatibility{
			Go:         ExpectedGoVersion,
			PostgreSQL: ExpectedPostgreSQL,
		},
	}
}
