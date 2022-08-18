package GoConsole

import (
	"errors"
	"flag"
	"reflect"
	"strings"

	Util "github.com/artziel/go-utilities"
)

type fieldTags struct {
	Name  string
	Usage string
}

type Command struct {

	// Public
	RootRequired bool
	RootExcluded bool
	Name         string
	Help         string
	Example      string

	Flags interface{}

	Init func() error
	Run  func(f interface{}) error

	// Privated
	flagSet *flag.FlagSet
}

func (c *Command) Execute(args []string) error {

	isRoot, _ := Util.IsRoot()
	if c.RootExcluded && isRoot {
		return ErrSubCommandRootExcluded
	}

	if c.RootRequired && !isRoot {
		return ErrSubCommandRootRequired
	}

	if err := c.parseStruct(c.Flags, args); err != nil {
		return err
	}
	return c.Run(c.Flags)
}

func (c *Command) parseTags(tags string) (fieldTags, error) {
	ft := fieldTags{}
	tags = strings.TrimSpace(tags)

	if len(tags) == 0 {
		return ft, ErrTagNoFieldTag
	}

	name := "GoConsole:"
	res := strings.Index(tags, name)
	if res < 0 {
		return ft, ErrTagNoFieldTag
	}
	tags = tags[res+len(name):]
	if len(tags) == 0 {
		return ft, ErrTagEmptyFieldTag
	}

	res = strings.Index(tags, "\"")
	if res < 0 {
		return ft, ErrTagEmptyFieldTag
	}
	tags = tags[res+1:]

	res = strings.Index(tags, "\"")
	if res < 0 {
		return ft, ErrTagEmptyFieldTag
	}
	tags = strings.TrimSpace(tags[:res])
	if len(tags) < 1 {
		return ft, ErrTagEmptyFieldTag
	}
	options := strings.Split(tags, ",")

	if len(options) == 0 {
		return ft, ErrTagEmptyFieldTag
	}

	for _, o := range options {
		pair := strings.SplitN(o, ":", 2)
		key := strings.ToLower(strings.TrimSpace(pair[0]))
		val := ""
		if len(pair) > 1 {
			val = strings.TrimSpace(pair[1])
		}

		switch key {
		case "name":
			ft.Name = val
		case "usage":
			ft.Usage = val
		}
	}

	return ft, nil
}

func (c *Command) parseStruct(r interface{}, args []string) error {
	if r == nil {
		return nil
	}

	if reflect.ValueOf(r).Type().Kind() != reflect.Ptr {
		return ErrWrongFalgsType
	}

	s := reflect.Indirect(reflect.ValueOf(r))
	kind := s.Kind().String()

	if kind != "struct" {
		return errors.New("flag are not an struct")
	}

	mStr := map[string]*string{}
	mBool := map[string]*bool{}
	mInt := map[string]*int{}
	c.flagSet = flag.NewFlagSet(c.Name, flag.ContinueOnError)

	for i := 0; i < s.NumField(); i++ {
		tags, err := c.parseTags(string(s.Type().Field(i).Tag))
		if err == nil {
			f := s.Field(i)
			switch f.Kind() {
			case reflect.String:
				val := f.String()
				mStr[s.Type().Field(i).Name] = &val
				c.flagSet.StringVar(mStr[s.Type().Field(i).Name], tags.Name, val, tags.Usage)
			case reflect.Bool:
				val := f.Bool()
				mBool[s.Type().Field(i).Name] = &val
				c.flagSet.BoolVar(mBool[s.Type().Field(i).Name], tags.Name, val, tags.Usage)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val := int(f.Int())
				mInt[s.Type().Field(i).Name] = &val
				c.flagSet.IntVar(mInt[s.Type().Field(i).Name], tags.Name, val, tags.Usage)
			}
		}
	}

	err := c.flagSet.Parse(args)
	if err != nil {
		return err
	}

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Kind() {
		case reflect.String:
			if _, found := mStr[s.Type().Field(i).Name]; found {
				f.SetString(*mStr[s.Type().Field(i).Name])
			}
		case reflect.Bool:
			if _, found := mBool[s.Type().Field(i).Name]; found {
				f.SetBool(*mBool[s.Type().Field(i).Name])
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if _, found := mInt[s.Type().Field(i).Name]; found {
				v := *mInt[s.Type().Field(i).Name]
				f.SetInt(int64(v))
			}
		}
	}

	return nil
}

func (c *Command) GetHelp() string {
	help := c.Help + "\n"
	if c.Example != "" {
		help = help + "\n    Example: " + c.Example + "\n\n"

	}

	if c.Flags != nil {
		s := reflect.Indirect(reflect.ValueOf(c.Flags))
		kind := s.Kind().String()

		if kind != "struct" {
			return ""
		}

		for i := 0; i < s.NumField(); i++ {
			tags, err := c.parseTags(string(s.Type().Field(i).Tag))
			if err == nil {
				help = help + Colored("    <blue>-"+tags.Name+"</blue>: "+tags.Usage+"\n")
			}
		}
	}

	return help
}
