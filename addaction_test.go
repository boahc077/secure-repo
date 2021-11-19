package main

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestAddAction(t *testing.T) {
	type args struct {
		inputYaml string
		action    string
	}
	const inputDirectory = "./testfiles/addaction/input"
	const outputDirectory = "./testfiles/addaction/output"
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "one job", args: args{inputYaml: "action-issues.yml", action: "step-security/harden-runner@main"}, want: "action-issues.yml", wantErr: false},
		{name: "two jobs", args: args{inputYaml: "2jobs.yml", action: "step-security/harden-runner@main"}, want: "2jobs.yml", wantErr: false},
		{name: "already present", args: args{inputYaml: "alreadypresent.yml", action: "step-security/harden-runner@main"}, want: "alreadypresent.yml", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := ioutil.ReadFile(path.Join(inputDirectory, tt.args.inputYaml))
			if err != nil {
				t.Fatalf("error reading test file")
			}
			got, err := AddAction(string(input), tt.args.action)

			if (err != nil) != tt.wantErr {
				t.Errorf("AddAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			output, err := ioutil.ReadFile(path.Join(outputDirectory, tt.args.inputYaml))
			if err != nil {
				t.Fatalf("error reading test file")
			}
			if got != string(output) {
				t.Errorf("AddAction() = %v, want %v", got, string(output))
			}
		})
	}
}