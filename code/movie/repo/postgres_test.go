package repository

import (
	"stockbit4/code/movie"
	"testing"
)

func Test_repositoryStruct_LogSearchData(t *testing.T) {
	type args struct {
		searchData movie.SearchData
	}
	tests := []struct {
		name    string
		p       *repositoryStruct
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.LogSearchData(tt.args.searchData); (err != nil) != tt.wantErr {
				t.Errorf("repositoryStruct.LogSearchData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
