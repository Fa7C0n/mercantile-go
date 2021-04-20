package types

import (
	"reflect"
	"testing"
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
			got, err := tr.TopLeft()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tile.TopLeft() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tile.TopLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
