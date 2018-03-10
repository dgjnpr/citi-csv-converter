package main

import (
	"io"
	"reflect"
	"strings"
	"testing"

	cc "github.com/dgjnpr/citi-csv-converter/citiconverter"
)

func TestCitiIngest(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *[][]string
		wantErr bool
	}{
		{
			name: "inflow and outflow in one statement",
			args: args{strings.NewReader(`"Account Number","Account Name","Transaction Date","Post Date","Reference Number","Transaction Detail","Billing Amount","Source Currency","Source Amount","Customer Ref","Employee Number"
"XXXXXXXXXX","Foo","01/01/2018","02/01/2018","12345","my company"," -1,000","GBP"," -1,000",,"98765"
"XXXXXXXXXX","Foo","02/01/2018","03/01/2018","23456","a shop","10.00","GBP","10.00",,"98765"`)},
			want: &[][]string{
				{"Account Number", "Account Name", "Transaction Date", "Post Date", "Reference Number", "Transaction Detail", "Billing Amount", "Source Currency", "Source Amount", "Customer Ref", "Employee Number"},
				{"XXXXXXXXXX", "Foo", "01/01/2018", "02/01/2018", "12345", "my company", " -1,000", "GBP", " -1,000", "", "98765"},
				{"XXXXXXXXXX", "Foo", "02/01/2018", "03/01/2018", "23456", "a shop", "10.00", "GBP", "10.00", "", "98765"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cc.CitiIngest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CitiIngest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CitiIngest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToYnab(t *testing.T) {
	type args struct {
		citi *[][]string
	}
	tests := []struct {
		name string
		args args
		want *[][]string
	}{
		{
			name: "inflow and outflow in one statement",
			args: args{&[][]string{
				{"Account Number", "Account Name", "Transaction Date", "Post Date", "Reference Number", "Transaction Detail", "Billing Amount", "Source Currency", "Source Amount", "Customer Ref", "Employee Number"},
				{"XXXXXXXXXX", "Foo", "01/01/2018", "02/01/2018", "12345", "my company", " -1,000", "GBP", " -1,000", "", "98765"},
				{"XXXXXXXXXX", "Foo", "02/01/2018", "03/01/2018", "23456", "a shop", "10.00", "GBP", "10.00", "", "98765"},
			}},
			want: &[][]string{
				{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"},
				{"01/01/2018", "my company", "Job Expense", "", "", "1,000"},
				{"02/01/2018", "a shop", "Job Expense", "", "10.00", ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cc.ToYnab(tt.args.citi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToYnab(): got %v, want %v", got, tt.want)
			}
		})
	}
}
