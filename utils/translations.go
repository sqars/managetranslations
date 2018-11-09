package utils

import (
	"fmt"
)

// Translation represents translation format written in JSON file
type Translation map[string]map[string]string

// GetExistingPool returns Translation which cointains all of available translations
func GetExistingPool() (Translation, error) {
	translationFilePaths, err := GetFilePaths("/*i18n*")
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
	return pool
}

func csvToTranslationFormat(data [][]string) Translation {
	languageIdxMap := make(map[int]string)
	translations := Translation{}
	for lineIdx, line := range data {
		for valIdx, val := range line {
			if lineIdx == 0 && valIdx > 0 {
				languageIdxMap[valIdx] = val
			} else if lineIdx > 0 && valIdx > 0 {
				lang, ok := languageIdxMap[valIdx]
				if !ok {
					panic("Cant map lineIdx to language something may be wrong with CSV format")
				}
				_, ok = translations[lang]
				if !ok {
					translations[lang] = make(map[string]string)
				}
				translations[lang][line[0]] = val
			}
		}
	}
	return translations
}

// UpdateTranslations updates translation from pool(existing translations)
func UpdateTranslations(data, pool Translation) Translation {
	count := 0
	for lang, keys := range data {
		for key, translation := range keys {
			existingTranslation, ok := pool[lang][key]
			if len(translation) == 0 && ok {
				if len(existingTranslation) > 0 {
					data[lang][key] = existingTranslation
					count++
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("%d translation keys updated", count))
	return data
}
