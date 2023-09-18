package Database

import (
	"context"
	"encoding/json"
	"os"
	"strings"
)

func Migrate() error {
	scripts, list, err := getSqlScripts()
	var usedMigrationsList []string
	if err == nil {
		for key, value := range scripts {
			_, used := list[key]
			if !used {
				_, err := Connection.Exec(context.Background(), value)
				if err != nil {
					return err
				}
				list[key] = true
			}
		}
		for key, _ := range list {
			usedMigrationsList = append(usedMigrationsList, key)
		}
		data, err := json.Marshal(usedMigrationsList)
		if err == nil {
			return os.WriteFile("Migrations/migrations_list.json", data, 0777)
		}
	}
	return err
}
func getSqlScripts() (map[string]string, map[string]bool, error) {
	dir, err := os.ReadDir("Migrations")
	if err != nil {
		return nil, nil, err
	}
	list := make(map[string]bool)
	scripts := make(map[string]string)
	var straightList []string
	for index, _ := range dir {
		if !dir[index].IsDir() {
			filename := dir[index].Name()
			if strings.HasSuffix(filename, ".sql") {
				data, err := os.ReadFile("Migrations/" + filename)
				if err == nil {
					scripts[filename] = string(data)
				} else {
					return nil, nil, err
				}
			} else {
				if filename == "migrations_list.json" {
					data, err := os.ReadFile("Migrations/" + filename)
					if err == nil {
						err = json.Unmarshal(data, &straightList)
						if err == nil {
							for _, item := range straightList {
								list[item] = true
							}
						} else {
							err := os.WriteFile("Migrations/"+filename, []byte("[]"), 0777)
							if err != nil {
								return nil, nil, err
							}
						}
					} else {
						return nil, nil, err
					}
				}
			}
		}
	}
	return scripts, list, nil
}
