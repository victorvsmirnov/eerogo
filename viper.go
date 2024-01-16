package eerogo

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func ViperSetDefault(vpr *viper.Viper, prefix string) error {
	params := os.Environ()
	for _, kv := range params {
		m := strings.Split(kv, `=`)
		if len(m) < 2 {
			continue
		}
		p := m[0]
		if prefix != `` {
			vpr.SetDefault(fmt.Sprintf("%s.%s", prefix, p), nil)
		} else {
			vpr.SetDefault(p, nil)

		}
	}
	return nil
}

func LoadViperConfiguration(file string, configuration any) error {
	vpr := viper.New()
	vpr.SetDefault("ENV", "development")
	vpr.AddConfigPath(".")
	vpr.AddConfigPath("/run/secrets/")
	vpr.SetConfigType("env")
	vpr.SetConfigName(file)
	vpr.AutomaticEnv()
	ViperSetDefault(vpr, "")

	if err := vpr.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("config file %s not found. Using ENV\n", file)
		} else {
			fmt.Printf("config file %s has errors: %s\n", file, err)
			return err
		}
	}
	err := vpr.Unmarshal(configuration)
	if err != nil {
		return err
	}
	return ValidateConfiguration(configuration)
}

func tagNameWithOptions(tag reflect.StructTag) (string, map[string]struct{}) {
	chunks := strings.Split(tag.Get("mapstructure"), ",")
	r := map[string]struct{}{}
	if len(chunks) == 1 {
		return chunks[0], r
	}
	for i := range chunks[1:] {
		r[chunks[i+1]] = struct{}{}
	}
	return chunks[0], r
}

func tagHasOption(tag reflect.StructTag, option string) bool {
	_, opts := tagNameWithOptions(tag)
	_, found := opts[option]
	return found
}

func ValidateConfiguration(configuration any) error {
	val := reflect.ValueOf(configuration)
	typ := reflect.TypeOf(configuration)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	errs := make([]error, 0)
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.FieldByIndex([]int{i})
		fieldType := typ.FieldByIndex([]int{i})
		tagname, _ := tagNameWithOptions(fieldType.Tag)
		if tagname == "" {
			tagname = strings.ToUpper(fieldType.Name)
		}
		if tagname == `-` {
			continue
		}
		if !tagHasOption(fieldType.Tag, "ignorenil") && (fieldValue.IsZero()) {
			errs = append(errs, fmt.Errorf("missing %s configuration variable. Provide or use ignorenil tag option", tagname))
		}
		if fieldType.Type.Kind() == reflect.Struct {
			err := ValidateConfiguration(fieldValue.Interface())
			if err != nil {
				errs = append(errs, err)
			}
		}
		if fieldType.Type.Kind() == reflect.Map {
			for _, key := range fieldValue.MapKeys() {
				err := ValidateConfiguration(fieldValue.MapIndex(key).Interface())
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	if len(errs) != 0 {
		errstr := make([]string, 0)
		for _, e := range errs {
			errstr = append(errstr, e.Error())
		}
		return fmt.Errorf(strings.Join(errstr, "\n"))
	}
	return nil
}
