package octapi

import "testing"

//Test must has file name like *_test.go (File for tested function must be *.go)
//Test must has name of main function like TestXxxx()
//Run test in terminal by command "go test" or "go test -v ."

func TestSum(t *testing.T) {

	t.Run("5 + 5 should be 10", func(t *testing.T) {
		total := Sum(5, 5)
		if total != 10 {
			t.Errorf("result was incorrect, got %d, want %d", total, 10)
		}
	})

}
