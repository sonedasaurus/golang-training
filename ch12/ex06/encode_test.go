package sexpr

import (
	"fmt"
	"testing"
)

func TestZero(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	var tests = []Movie{
		Movie{Title: "", Subtitle: "", Year: 0, Actor: nil, Oscars: nil},
		Movie{Title: "Title", Subtitle: "", Year: 0, Actor: nil, Oscars: nil},
		Movie{Title: "", Subtitle: "subtitle", Year: 0, Actor: nil, Oscars: nil},
		Movie{Title: "", Subtitle: "", Year: 1989, Actor: nil, Oscars: nil},
		Movie{Title: "", Subtitle: "", Year: 0, Actor: map[string]string{
			"Brig. Gen. Jack D. Ripper": "Sterling Hayden",
			`Maj. T.J. "King" Kong`:     "Slim Pickens",
		}, Oscars: nil},
		Movie{Title: "", Subtitle: "", Year: 0, Actor: nil,
			Oscars: []string{
				"Best Actor (Nomin.)",
				"Best Adapted Screenplay (Nomin.)",
			},
		},
	}

	for _, test := range tests {
		data, err := Marshal(test)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		t.Logf("Marshal() = %s\n", data)
		fmt.Println(string(data))
	}
}
