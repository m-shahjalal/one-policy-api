package config

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"gorm.io/gorm"
)

type ModelVersion struct {
	ID        uint   `gorm:"primaryKey"`
	ModelName string `gorm:"unique;not null"`
	Hash      string `gorm:"not null"`
	UpdatedAt time.Time
}

type SmartMigrator struct {
	db            *gorm.DB
	modelVersions map[string]*ModelVersion
	modelHashes   map[string]string
}

func newSmartMigrator(db *gorm.DB) *SmartMigrator {
	db.AutoMigrate(&ModelVersion{})

	// Load all model versions into memory
	var versions []ModelVersion
	db.Find(&versions)

	versionMap := make(map[string]*ModelVersion)
	for i := range versions {
		versionMap[versions[i].ModelName] = &versions[i]
	}

	return &SmartMigrator{
		db:            db,
		modelVersions: versionMap,
		modelHashes:   make(map[string]string),
	}
}

func (sm *SmartMigrator) generateModelHash(model any) string {
	modelName := reflect.TypeOf(model).Elem().Name()

	// Check if hash is already cached
	if hash, exists := sm.modelHashes[modelName]; exists {
		return hash
	}

	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var fields []string
	for i := range t.NumField() {
		field := t.Field(i)
		gormTag := field.Tag.Get("gorm")
		jsonTag := field.Tag.Get("json")

		fieldInfo := fmt.Sprintf("%s:%s:%s:%s",
			field.Name, field.Type.String(), gormTag, jsonTag)
		fields = append(fields, fieldInfo)
	}

	data, _ := json.Marshal(fields)
	hash := fmt.Sprintf("%x", md5.Sum(data))

	// Cache the hash
	sm.modelHashes[modelName] = hash
	return hash
}

func (sm *SmartMigrator) migrateIfChanged(models ...any) error {
	// Batch process all models
	var modelsToMigrate []interface{}
	var versionUpdates []*ModelVersion

	for _, model := range models {
		modelName := reflect.TypeOf(model).Elem().Name()
		currentHash := sm.generateModelHash(model)

		version, exists := sm.modelVersions[modelName]

		if !exists {
			fmt.Printf("Will migrate %s (new model)\n", modelName)
			modelsToMigrate = append(modelsToMigrate, model)
			newVersion := &ModelVersion{
				ModelName: modelName,
				Hash:      currentHash,
				UpdatedAt: time.Now(),
			}
			versionUpdates = append(versionUpdates, newVersion)
			sm.modelVersions[modelName] = newVersion
		} else if version.Hash != currentHash {
			fmt.Printf("Will migrate %s (hash changed)\n", modelName)
			modelsToMigrate = append(modelsToMigrate, model)
			version.Hash = currentHash
			version.UpdatedAt = time.Now()
			versionUpdates = append(versionUpdates, version)
		} else {
			fmt.Printf("Skipping %s (no changes)\n", modelName)
		}
	}

	// If there are models to migrate, do it in a single transaction
	if len(modelsToMigrate) > 0 {
		return sm.db.Transaction(func(tx *gorm.DB) error {
			// Migrate all changed models
			if err := tx.AutoMigrate(modelsToMigrate...); err != nil {
				return err
			}

			// Update all version records
			for _, version := range versionUpdates {
				if version.ID == 0 {
					if err := tx.Create(version).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Save(version).Error; err != nil {
						return err
					}
				}
			}
			return nil
		})
	}

	return nil
}

func InitMigration(DB *gorm.DB) error {
	migrator := newSmartMigrator(DB)

	if err := migrator.migrateIfChanged(&model.User{}); err != nil {
		return err
	}

	if err := migrator.migrateIfChanged(&model.Policy{}); err != nil {
		return err
	}

	return nil
}
