package day202404_test

import (
	"testing"

	day202404 "github.com/LukasJenicek/aoc/years/2024/day04"
)

func TestFourthDay_RunFirstPart(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "first",
			args: args{
				input: `X..S
MXSA
AMAM
SAMX
.SX.`,
			},
			want:    "5",
			wantErr: false,
		},
		{
			name: "second",
			args: args{
				input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			},
			want: "18",
		},
		{
			name: "third",
			args: args{
				input: `S..S..S
.A.A.A.
..MMM..
SAMXMAS
..MMM..
.A.A.A.
S..S..S`,
			},
			want: "8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fo := &day202404.FourthDay{}
			got, err := fo.RunFirstPart(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunFirstPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RunFirstPart() got = %v, want %v", got, tt.want)
			}
		})
	}
}
