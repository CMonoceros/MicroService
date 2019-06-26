package log

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

type ORMLogger struct {
	gorm.LogWriter
	DataSource string
}

func (logger ORMLogger) Print(values ...interface{}) {
	logger.Println(GetORMFormatter(append(values, logger.DataSource)...)...)
}

func GetORMDefaultWriter(dataSource string) *ORMLogger {
	return &ORMLogger{
		LogWriter: log.New(
			io.MultiWriter(
				os.Stdout,
				fileHandler.fws[_ormInfoIdx],
			),
			"", 0,
		),
		DataSource: dataSource,
	}
}

func ormIsPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func GetORMFormatter(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			title           = "[GORM]"
			sql             string
			formattedValues []string
			level           = values[0]
			currentTime     = "[" + time.Now().Format("2006/01/02 15:04:05.999") + "]"
			source          = fmt.Sprintf("%v", values[1])
		)

		messages = []interface{}{currentTime, title}

		if level == "sql" {
			messages = append(messages, fmt.Sprintf("[DSN:%v]", values[6]))
			messages = append(messages, fmt.Sprintf("[%.2fms]", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			messages = append(messages, source)
			messages = append(messages, fmt.Sprintf("[ROWS:%v]\n", strconv.FormatInt(values[5].(int64), 10)))

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); ormIsPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						switch value.(type) {
						case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
							formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
						default:
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						}
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if regexp.MustCompile(`\$\d+`).MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range regexp.MustCompile(`\?`).Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, fmt.Sprintf("[SQL] %v", sql))
		} else {
			messages = append(messages, values[2:]...)
		}
	}

	return
}
