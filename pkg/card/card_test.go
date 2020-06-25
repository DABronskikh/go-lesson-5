package card

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Валидный номер",
			args: args{
				number: "5106 2100 0000 0007",
			},
			want: true,
		},
		{
			name: "НЕ Валидный номер",
			args: args{
				number: "5106 2100 0000 0000",
			},
			want: false,
		},
		{
			name: "НЕ Валидный номер (есть буквы)",
			args: args{
				number: "5106 2100 0000 000X",
			},
			want: false,
		},
	}
	for _, tt := range tests {
			if got := IsValid(tt.args.number); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
	}
}
