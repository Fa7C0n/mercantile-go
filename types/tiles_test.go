package types

import (
	"testing"

	"github.com/Fa7C0n/mercantile-go/maths"
	"github.com/stretchr/testify/assert"
)

func TestTile_TopLeft(t *testing.T) {
	type fields struct {
		x uint
		y uint
		z uint
	}
	tests := []struct {
		name    string
		fields  fields
		want    *LongLat
		wantErr bool
	}{
		{
			name: "WithinLimits",
			fields: fields{
				x: 486,
				y: 332,
				z: 10,
			},
			want: &LongLat{
				Long: -9.140625,
				Lat:  53.330872983017045,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tile{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			got, err := tr.Ul()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tile.TopLeft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, maths.Round(got.Lat, 10e-7), maths.Round(tt.want.Lat, 10e-7))
			assert.Equal(t, maths.Round(got.Long, 10e-7), maths.Round(tt.want.Long, 10e-7))
		})
	}
}

func TestBoundsLatLong(t *testing.T) {
	tile := Tile{
		x: 486,
		y: 332,
		z: 10,
	}
	want := LongLatBbox{
		BBox{
			Left:   -9.140625,
			Bottom: 53.12040528310657,
			Right:  -8.7890625,
			Top:    53.33087298301705,
		},
	}

	got, err := tile.BoundsLatLong()

	assert.Nil(t, err)
	assert.Equal(t, 0*10e-7, maths.Round(got.Bottom-want.Bottom, 10e-7))
	assert.Equal(t, 0*10e-7, maths.Round(got.Left-want.Left, 10e-7))
	assert.Equal(t, 0*10e-7, maths.Round(got.Right-want.Right, 10e-7))
	assert.Equal(t, 0*10e-7, maths.Round(got.Top-want.Top, 10e-7))
}
