package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/bytes"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var newMap AnnotationStickers

	tmp := map[string]string{"1": "1", "2": "2", "3": "3"}

	newMap = tmp

	v, err := newMap.Value()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)
}

type AnnotationStickers map[string]string

// Value implements the database/sql Valuer interface for adding AnnotationStickers to the database
// Stickers are stored in the dtabase as a slice of strings like "[key=val]"
// They are encoded into a JSON string for storing intothe database, and the JSON sqlite extension is
// able to manipvalte them like an object
func (a AnnotationStickers) Value() (string, error) {
	stickeSlice := make([]string, 0, len(a))

	for k, v := range a {
		stickeSlice = append(stickeSlice, fmt.Sprintf("%s=%s", k, v))
	}

	sticks, err := json.Marshal(stickeSlice)
	if err != nil {
		return "", err
	}

	return string(sticks), nil
}

// Scan implements the databasse/sql Scanner interface for retrieving AnnotationStickers from the database
// The string is decoded into a slice of strings, which are then converted back into a map
func (a *AnnotationStickers) Scan(value interface{}) error {
	vString, ok := value.(string)
	if !ok {
		return &Error{
			Code: EInvalid,
			Msg:  "could not load stickers from sqlite",
		}
	}

	var stickSlice []string
	if err := json.NewDecoder(strings.NewReader(vString)).Decode(&stickSlice); err != nil {
		return err
	}

	stickMap, err := stickSliceToMap(stickSlice)
	if err != nil {
		return nil
	}

	*a = stickMap

	return nil
}

func stickSliceToMap(stickers []string) (map[string]string, error) {
	stickMap := map[string]string{}

	for i := range stickers {
		if stick0, stick1, found := strings.Cut(stickers[i], "="); found {
			stickMap[stick0] = stick1
		} else {
			return nil, invalidStickerError(stickers[i])
		}
	}

	return stickMap, nil
}

func invalidStickerError(s string) error {
	return &Error{
		Code: EInternal,
		Msg:  fmt.Sprintf("invalid sticker: %q", s),
	}
}

type Error struct {
	Code string
	Msg  string
	Op   string
	Err  error
}

// Error implements the error interface by writing out the recursive messages.
func (e *Error) Error() string {
	if e.Msg != "" && e.Err != nil {
		var b strings.Builder
		b.WriteString(e.Msg)
		b.WriteString(": ")
		b.WriteString(e.Err.Error())
		return b.String()
	} else if e.Msg != "" {
		return e.Msg
	} else if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("<%s>", e.Code)
}

const (
	EInternal       = "internal error"
	ENotImplemented = "not implemented"
	ENotFound       = "not found"
	EConflict       = "conflict"
	EInvalid        = "invalid"
)

type Authorization struct {
	ID          string `json:"id"`
	Token       string `json:"token"`
	Status      string `json:"status"`
	Description string `json:"description"`
	OrgID       string `json:"orgID"`
	UserID      string `json:"userID"`
	Permissions string `json:"permissions"`
}

func IsActive(a *Authorization) bool {
	return a.IsActive()
}

func (a *Authorization) IsActive() bool {
	return a.Status == Active
}

func (a *Authorization) GetUserID() string {
	return a.UserID
}

func (a *Authorization) Kind() string {
	return AuthorizationKind
}

func (a *Authorization) Identifier() string {
	return a.ID
}

const Ext = ".tmpl"

type pathSpec struct {
	in, out string
}

func (p *pathSpec) String() string {
	return p.in + " ->" + p.out
}

func (p *pathSpec) IsGoFile() bool {
	return filepath.Ext(p.out) == ".go"
}

func parserPath(path string) (string, string) {
	p := strings.IndexByte(path, '=')
	if p == -1 {
		if filepath.Ext(path) != Ext {
			errExit("template file '%s' must have .tmpl extension", path)
		}

		return path, path[:len(path)-len(Ext)]
	}

	return path[:p], path[p+1]
}

type data struct {
	In interface{}
	D  listValue
}

func errExit(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

type listValue map[string]string

func (l listValue) String() string {
	res := make([]string, 0, len(l))
	for k, v := range l {
		res = append(res, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(res, ", ")
}

func (l listValue) Set(v string) error {
	mv := strings.Split(v, "=")
	if len(mv) != 2 {
		return fmt.Errorf("expected NAME=VALUE, got %s", v)
	}

	l[mv[0]] = mv[1]
	return nil
}

func mustReadAll(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		errExit(err.Error())
	}

	return data
}

func readData(path string) interface{} {
	data := mustReadAll(path)
	var v interface{}

	if err := json.Unmarshal(StripComments(data), &v); err != nil {
		errExit("invalid JSON data: %s", err.Error())
	}

	return v
}

func fileMode(path string) os.FileMode {
	stat, err := os.Stat(path)
	if err != nil {
		errExit(err.Error())
	}

	return stat.Mode()
}

var funcs = template.FuncMap{
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
}

func process(data interface{}, specs []pathSpec) {
	for _, spec := range specs {
		var (
			t   *template.Template
			err error
		)
		t, err = template.New("gen").Funcs(funcs).Parse(string(mustReadAll(spec.in)))
		if err != nil {
			errExit("error procession template '%s' : %s", spec.in, err.Error())
		}

		var buf bytes.Buffer
		if spec.IsGoFile() {
			// preamble
			fmt.Fprintf(&buf, "// Code generated by %s, DO　NOT　EDIT．\n", spec.in)
			fmt.Fprintln(&buf)
		}

		err = t.Execute(&buf, data)
		if err != nil {
			errExit("error executing template '%s': %s", spec.in, err.Error())
		}

		generated := buf.Bytes()
		if spec.IsGoFile() {
			generated, err = formatter(generated)
			if err != nil {
				errExit("error formatting '%s': %s", spec.in, err.Error())
			}

			os.WriteFile(spec.out, generated, fileMode(spec.in))
		}
	}
}

var (
	formatter func([]byte) ([]byte, error)
)

func formatSource(in []byte) ([]byte, error) {
	r := bytes.NewReader(in)
	cmd := exec.Command("goimports")
	cmd.Stdin = r
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("error running gomports: %s", string(ee.Stderr))
		}

		return nil, fmt.Errorf("error running goimports: %s", string(out))
	}

	return out, nil
}
