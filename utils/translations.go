package utils

import (
	"fmt"
)

// Translation represents translation format written in JSON file
type Translation map[string]map[string]string

// GetExistingPool returns Translation which cointains all of available translations
func GetExistingPool() (Translation, error) {
	translationFilePaths, err := GetTranslationFilePaths()
	if err != nil {
		return nil, err
	}
	translationPoolSl := []Translation{}
	for _, trFilePath := range translationFilePaths {
		translation, err := GetJSONTranslationData(trFilePath)
		if err != nil {
			return nil, err
		}
		translationPoolSl = append(translationPoolSl, translation)
	}
	return mergeTranslations(translationPoolSl), nil
}

func mergeTranslations(translations []Translation) Translation {
	pool := Translation{}
	for _, translation := range translations {
		fmt.Println(translation)
		for lang, keys := range translation {
			existing, ok := pool[lang]
			if ok {
				for key, value := range keys {
					if len(value) > 0 {
						existing[key] = value
					}
				}
			} else {
				pool[lang] = keys
			}
		}
	}
	fmt.Println(pool)
	return pool
}
