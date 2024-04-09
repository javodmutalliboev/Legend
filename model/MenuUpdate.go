package model

import (
	"Legend/database"
	"fmt"
	"strings"
)

func (m *Menu) MenuUpdate() error {
	database := database.DB()
	defer database.Close()

	var fields []string
	var args []interface{}
	i := 1

	if m.TitleUz != "" {
		fields = append(fields, fmt.Sprintf("title_uz = $%d", i))
		args = append(args, m.TitleUz)
		i++
	}

	if m.TitleRu != "" {
		fields = append(fields, fmt.Sprintf("title_ru = $%d", i))
		args = append(args, m.TitleRu)
		i++
	}

	if m.TitleEn != "" {
		fields = append(fields, fmt.Sprintf("title_en = $%d", i))
		args = append(args, m.TitleEn)
		i++
	}

	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// build the SQL query
	sql := fmt.Sprintf("UPDATE menu SET %s, updated_at = NOW() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, m.ID)

	_, err := database.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
