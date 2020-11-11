package heimdallr

import "testing"

func TestValidateHKID(t *testing.T) {

	tests := []struct {
		name       string
		id         string
		shouldBeOK bool
	}{
		{name: "Unnormalized-M Card", id: "E364912(5)", shouldBeOK: true},
		{name: "Unnormalized-M Card", id: "M812318(2)", shouldBeOK: true},
		{name: "OK-M Card", id: "M8123182", shouldBeOK: true},
		{name: "OK-S Card", id: "J4479871", shouldBeOK: true},
		{name: "OK-I Card", id: "I336251A", shouldBeOK: true},
		{name: "OK-T Card", id: "T4293376", shouldBeOK: true},
		{name: "OK-T Card", id: "R914749(6)", shouldBeOK: true},
		// {name: "OK-Double Character", id: "WE4316017", shouldBeOK: true},
		{name: "OK-End with Char", id: "W392026A", shouldBeOK: true},
		{name: "Invalid-O not allowed", id: "O8123182", shouldBeOK: false},
		{name: "Invalid-ShortLen", id: "M812318", shouldBeOK: false},
		{name: "Invalid-LongLen", id: "M81231888", shouldBeOK: false},
		{name: "Invalid-Letter in wrong place", id: "M812D188", shouldBeOK: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := ValidateHKID(tt.id)
			if err != nil {
				t.Error(err)
			}
			if ok != tt.shouldBeOK {
				t.Fail()
			}
		})
	}
}

func BenchmarkValidateHKID(b *testing.B) {
	// run the Fib function b.N times
	const ID = "M8123182"
	for n := 0; n < b.N; n++ {
		ValidateHKID(ID)
	}
}
