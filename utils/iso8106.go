package utils

import (
	"fmt"
	"time"
)

// Parsování ISO8601 stringu.
//
//	@param timestamp
//	@return *time.Time
//	@return error
func ParseISO8601String(timestamp string) (*time.Time, error) {
	// Definovat část řetězců rozložení reprezentujících různé formáty ISO 8601
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05.999999999Z07:00", // Nanosekundová přesnost
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	var parsedTime time.Time
	var err error

	// Projděte každé rozložení a zkuste analyzovat časové razítko.
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, timestamp)
		if err == nil {
			// Pokud se parsování podaří, vrátí se ukazatel na parsovaný čas.
			return &parsedTime, nil
		}
	}

	// Pokud všechny pokusy o rozbor selžou, vrátí se hodnota nil a chybová zpráva.
	return nil, fmt.Errorf("neplatný formát ISO 8601: %s", timestamp)
}
