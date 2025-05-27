package migrations

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/primeiro/internal/infra/database"
	"gorm.io/gorm"
)

type Migration struct {
	Version int
	Name    string
	UpSQL   string
	DownSQL string
}

func RunMigrations() error {
	// Obter o diretório de trabalho atual
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("erro ao obter diretório de trabalho: %v", err)
	}

	// Construir o caminho para o diretório de migrations
	migrationsDir := filepath.Join(workDir, "internal", "infra", "database", "migrations", "sql")

	// Ler os arquivos do diretório
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("erro ao ler diretório de migrations: %v", err)
	}

	var migrations []Migration
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".up.sql") {
			version := extractVersion(file.Name())

			// Ler arquivo up
			upSQL, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				return fmt.Errorf("erro ao ler arquivo de migration up: %v", err)
			}

			// Ler arquivo down
			downFileName := strings.Replace(file.Name(), ".up.sql", ".down.sql", 1)
			downSQL, err := os.ReadFile(filepath.Join(migrationsDir, downFileName))
			if err != nil {
				return fmt.Errorf("erro ao ler arquivo de migration down: %v", err)
			}

			migrations = append(migrations, Migration{
				Version: version,
				Name:    strings.TrimSuffix(file.Name(), ".up.sql"),
				UpSQL:   string(upSQL),
				DownSQL: string(downSQL),
			})
		}
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	// Criar tabela de migrations se não existir
	err = database.DB.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			version INT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`).Error
	if err != nil {
		return fmt.Errorf("erro ao criar tabela de migrations: %v", err)
	}

	// Executar migrations pendentes
	for _, migration := range migrations {
		var count int64
		database.DB.Table("migrations").Where("version = ?", migration.Version).Count(&count)
		if count == 0 {
			log.Printf("Executando migration: %s", migration.Name)
			err = database.DB.Transaction(func(tx *gorm.DB) error {
				if err := tx.Exec(migration.UpSQL).Error; err != nil {
					return err
				}
				return tx.Exec("INSERT INTO migrations (version, name) VALUES (?, ?)", migration.Version, migration.Name).Error
			})
			if err != nil {
				return fmt.Errorf("erro ao executar migration %s: %v", migration.Name, err)
			}
		}
	}

	return nil
}

func extractVersion(filename string) int {
	var version int
	fmt.Sscanf(filename, "%d_", &version)
	return version
}
