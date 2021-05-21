package conf

import (
	"os"
	"path/filepath"
	"testing"

	jww "github.com/spf13/jwalterweatherman"
)

func init() {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	df := filepath.Join(UserHomeDir(), ".shoppe", "demo")
	if !Exists(df) {
		os.MkdirAll(df, os.ModePerm)
		var yamlExample = []byte(`
Hacker: false
name: steve2
`)

		os.WriteFile(filepath.Join(df, "app.yaml"), yamlExample, os.ModePerm)
	}
}

func TestInit(t *testing.T) {

	Init("demo")

	name := GetString("name")
	if name != "steve2" {
		t.Fatal("getstring error")
	}

	hacker := GetBool("Hacker")
	if !hacker {
		t.Fatal("getBool error")
	}

	clothing := GetMap("clothing")

	if clothing["jacket"] != "leather" {
		t.Fatal("getmap error")
	}

	hobbies := GetStrings("hobbies")

	if hobbies[0] != "skateboarding" {
		t.Fatal("GetStrings error")
	}

}
