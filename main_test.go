package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

func TestYnabParser(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "inflow and outflow in one statement",
			args: args{strings.NewReader(`"Account Number","Account Name","Transaction Date","Post Date","Reference Number","Transaction Detail","Billing Amount","Source Currency","Source Amount","Customer Ref","Employee Number"
"XXXXXXXXXX","Foo","01/01/2018","02/01/2018","12345","my company"," -1,000","GBP"," -1,000",,"98765"
"XXXXXXXXXX","Foo","02/01/2018","03/01/2018","23456","a shop","10.00","GBP","10.00",,"98765"`)},
			want: [][]string{
				{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"},
				{"01/01/2018", "my company", "Job Expense", "", "", "1,000"},
				{"02/01/2018", "a shop", "Job Expense", "", "10.00", ""},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YnabParser(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("YnabParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YnabParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
