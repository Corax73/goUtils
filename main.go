package goutils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"slices"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// LogInit using the passed path string, creates, if missing, a log file and assigns errors to be output to it.
func LogInit(path string) {
	defaultPath := "./logs/app.log"
	if path == "" {
		path = defaultPath
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}

// Logging writes an indefinite number of transmitted non-empty errors to the log.
func Logging(errors ...error) {
	for _, err := range errors {
		if err != nil {
			log.Println(err)
		}
	}
}

// GetConfFromEnvFile receives data for the database from the environment file. If successful, returns a non-empty map.
func GetConfFromEnvFile(filename string) map[string]string {
	resp := make(map[string]string)
	if filename == "" {
		filename = ".env"
	}
	envFile, err := godotenv.Read(filename)
	if err == nil {
		resp = envFile
	} else {
		Logging(err)
	}
	return resp
}

// GCRunAndPrintMemory runs a garbage collection and if setting the APP_ENV environment variable as "dev" prints currently allocated number of bytes on the heap.
func GCRunAndPrintMemory() {
	debugSet := false
	settings := GetConfFromEnvFile(".env")
	if val, ok := settings["APP_ENV"]; ok && val == "dev" {
		debugSet = true
	}
	if debugSet {
		var stat runtime.MemStats
		runtime.ReadMemStats(&stat)
		fmt.Println(stat.Alloc / 1024)
	}
	if val, ok := settings["GC_MANUAL_RUN"]; ok && val == "true" {
		runtime.GC()
	}
}

// ConcatSlice returns a string from the elements of the passed slice with strings. Separator - space.
func ConcatSlice(strSlice []string) string {
	resp := ""
	if len(strSlice) > 0 {
		var strBuilder strings.Builder
		for _, val := range strSlice {
			strBuilder.WriteString(val)
		}
		resp = strBuilder.String()
		strBuilder.Reset()
	}
	return resp
}

// CompareMapsByStringKeys for map-arguments, checks the keys of the first argument that contain non-empty values
// ​​to see if they are present in the second argument.
func CompareMapsByStringKeys(map1, map2 map[string]string) bool {
	var resp bool
	len1 := len(map1)
	len2 := len(map2)
	if len1 == len2 {
		keysSlice1 := GetMapKeysWithValue(map1)
		keysSlice2 := GetMapKeysWithValue(map2)
		check := true
		for _, val := range keysSlice1 {
			if !slices.Contains(keysSlice2, val) {
				check = false
				break
			}
		}
		resp = check
	}
	return resp
}

// GetMapKeysWithValue returns from the argument map, a map with keys with non-empty values.
func GetMapKeysWithValue(mapArg map[string]string) []string {
	var resp []string
	if len(mapArg) > 0 {
		for key, val := range mapArg {
			if val != "" {
				resp = append(resp, key)
			}
		}
	}
	return resp
}

// GetMapValues from the passed map returns a slice with its non-empty values.
func GetMapValues(mapArg map[string]string) []string {
	var resp []string
	if len(mapArg) > 0 {
		for _, value := range mapArg {
			if value != "" {
				resp = append(resp, value)
			}
		}
	}
	return resp
}

// GetIndexByStrValue returns the integer index of the passed value in the passed slice; if the value is missing, then -1.
func GetIndexByStrValue(data []string, value string) int {
	resp := -1
	for i, val := range data {
		if val == value {
			resp = i
			break
		}
	}
	return resp
}

// SqlToMap the values ​​of the passed response structure are returned by the database as a map.
func SqlToMap(rows *sql.Rows) []map[string]any {
	resp := make([]map[string]any, 0)
	columns, err := rows.Columns()
	if err != nil {
		Logging(err)
	} else {
		scanArgs := make([]any, len(columns))
		values := make([]any, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				Logging(err)
			}
			record := make(map[string]any)
			for i, col := range values {
				if col != nil {
					switch col.(type) {
					case bool:
						record[columns[i]] = col.(bool)
					case int:
						record[columns[i]] = col.(int)
					case int64:
						record[columns[i]] = col.(int64)
					case float64:
						record[columns[i]] = col.(float64)
					case string:
						record[columns[i]] = col.(string)
					case time.Time:
						record[columns[i]] = col.(time.Time)
					case []byte:
						record[columns[i]] = string(col.([]byte))
					default:
						record[columns[i]] = col
					}
				}
			}
			resp = append(resp, record)
		}
	}
	return resp
}

// GetMapKeys returns a slice of the keys of the passed map.
func GetMapKeys(argMap map[string]string) []string {
	resp := make([]string, len(argMap))
	var i int
	for k := range argMap {
		resp[i] = k
		i++
	}
	return resp
}

// PresenceMapKeysInOtherMap returns a Boolean answer whether the keys of the first passed card are contained in the second.
func PresenceMapKeysInOtherMap(map1, map2 map[string]string) bool {
	var resp bool
	keys1 := GetMapKeys(map1)
	keys2 := GetMapKeys(map2)
	check := true
	for _, val := range keys1 {
		if !slices.Contains(keys2, val) {
			check = false
			break
		}
	}
	resp = check
	return resp
}

// GetMapWithoutKeys returns the transferred map without the transferred key.
func GetMapWithoutKeys(map1 map[string]string, exceptKeys []string) map[string]string {
	resp := make(map[string]string, len(map1)-len(exceptKeys))
	for k, v := range map1 {
		if !slices.Contains(exceptKeys, k) {
			resp[k] = v
		}
	}
	return resp
}

// IsEmail returns a boolean value checking the passed string against the pattern of matching the email address.
func IsEmail(email string) bool {
	emailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegexp.MatchString(email)
}

// HashPassword returns the hash of the passed string and a possible error.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash returns a boolean value comparing the passed strings (password and hash)
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ClearStringOfCharacters in the passed string, replaces the characters from the passed slice with an empty string.
func ClearStringOfCharacters(str string, characters []string) string {
	if str != "" {
		for _, char := range characters {
			str = strings.ReplaceAll(str, char, "")
		}
	}
	return str
}
