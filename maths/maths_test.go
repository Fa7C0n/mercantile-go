package maths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleRound() {
	fmt.Printf("%.4f\n", Round(0.363636, 0.001)) // 0.364
	fmt.Printf("%.4f\n", Round(0.363636, 0.01))  // 0.36
	fmt.Printf("%.4f\n", Round(0.363636, 0.1))   // 0.4
	fmt.Printf("%.4f\n", Round(0.363636, 0.05))  // 0.35
	fmt.Printf("%.4f\n", Round(3.2, 1))          // 3
	fmt.Printf("%.4f\n", Round(32, 5))           // 30
	fmt.Printf("%.4f\n", Round(33, 5))           // 35
	fmt.Printf("%.4f\n", Round(32, 10))          // 30

	fmt.Printf("%.4f\n", Round(-0.363636, 0.001)) // -0.364
	fmt.Printf("%.4f\n", Round(-0.363636, 0.01))  // -0.36
	fmt.Printf("%.4f\n", Round(-0.363636, 0.1))   // -0.4
	fmt.Printf("%.4f\n", Round(-0.363636, 0.05))  // -0.35
	fmt.Printf("%.4f\n", Round(-3.2, 1))          // -3
	fmt.Printf("%.4f\n", Round(-32, 5))           // -30
	fmt.Printf("%.4f\n", Round(-33, 5))           // -35
	fmt.Printf("%.4f\n", Round(-32, 10))          // -30

	// Output:
	// 0.3640
	// 0.3600
	// 0.4000
	// 0.3500
	// 3.0000
	// 30.0000
	// 35.0000
	// 30.0000
	// -0.3640
	// -0.3600
	// -0.4000
	// -0.3500
	// -3.0000
	// -30.0000
	// -35.0000
	// -30.0000
}
func TestRound(t *testing.T) {
	type args struct {
		x    float64
		unit float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0.001",
			args: args{
				x:    0.363636,
				unit: 0.001,
			},
			want: 0.364,
		},
		{
			name: "0.001",
			args: args{
				x:    0.363636,
				unit: 0.01,
			},
			want: 0.36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Round(tt.args.x, tt.args.unit)
			assert.Equal(t, got, tt.want)
		})
	}
}
