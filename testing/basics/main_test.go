package main

import "testing"

func setup()    {}
func teardown() {}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	m.Run()
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "first larger",
			a:    10,
			b:    3,
			want: 10,
		},
		{
			name: "second is larger",
			a:    4,
			b:    99,
			want: 99,
		},
		{
			name: "number are equal",
			a:    99,
			b:    99,
			want: 99,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := max(test.a, test.b)
			if got != test.want {
				t.Errorf(
					"Max(%d, %d) = %d; want %d",
					test.a,
					test.b,
					got,	
					test.want,
				)
			}
		})

	}

}

// func TestAddWithSubtests(t *testing.T) {
// 	// setup
//
// 	t.Run("test 1", func(t *testing.T) {
// 		got := Add(2, 5)
// 		if got != 7 {
// 			t.Errorf("Expected 5, got %d", got)
// 		}
// 	})
// 	t.Run("test 2", func(t *testing.T) {
// 		got := Add(55, 45)
// 		if got != 100 {
// 			t.Errorf("Expected 5, got %d", got)
// 		}
// 	})
//
// }
