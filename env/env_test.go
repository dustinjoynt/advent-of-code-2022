package env

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {

	t.Run("successfully load env", func(t *testing.T) {

		err := Load(".test_env")
		if err != nil {
			t.Fatalf("failed to load env: %v", err)
		}

		env1 := os.Getenv("TEST")
		want := "testing"
		if env1 != want {
			t.Fatalf("failed to load env var TEST. want: %v, got: %v", want, env1)
		}

		env2 := os.Getenv("TEST2")
		want2 := "testing2"
		if env2 != want2 {
			t.Fatalf("failed to load env var TEST2. want: %v, got: %v", want2, env2)
		}
	})
}
