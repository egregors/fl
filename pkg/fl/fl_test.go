// Package fl – Find Links. Tools for searching downloadable links on the web page
//
// @egregors 2019
package fl

import (
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func Test_unique(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{[]string{"a", "b", "a", "a", "b", "c", "c", "a"}}, []string{"a", "b", "c"}},
		{"", args{[]string{"a", "a", "a", "a", "a", "a", "a", "a"}}, []string{"a"}},
		{"", args{[]string{"abc", "abc", "aaa", "aaa", "abc", "abc"}}, []string{"abc", "aaa"}},
		{"", args{[]string{"a"}}, []string{"a"}},
		{"", args{[]string{}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinder_getDoc(t *testing.T) {
	type fields struct {
		url     string
		exts    []string
		links   []string
		linksCh chan string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *html.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Finder{
				url:     tt.fields.url,
				exts:    tt.fields.exts,
				links:   tt.fields.links,
				linksCh: tt.fields.linksCh,
			}
			got, err := f.getDoc()
			if (err != nil) != tt.wantErr {
				t.Errorf("Finder.getDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finder.getDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinder_findAll(t *testing.T) {
	type fields struct {
		url     string
		exts    []string
		links   []string
		linksCh chan string
	}
	type args struct {
		n *html.Node
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Finder{
				url:     tt.fields.url,
				exts:    tt.fields.exts,
				links:   tt.fields.links,
				linksCh: tt.fields.linksCh,
			}
			got, err := f.findAll(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finder.findAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finder.findAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinder_find(t *testing.T) {
	type fields struct {
		url     string
		exts    []string
		links   []string
		linksCh chan string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Finder{
				url:     tt.fields.url,
				exts:    tt.fields.exts,
				links:   tt.fields.links,
				linksCh: tt.fields.linksCh,
			}
			if err := f.find(); (err != nil) != tt.wantErr {
				t.Errorf("Finder.find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinder_GetLinks(t *testing.T) {
	type fields struct {
		url     string
		exts    []string
		links   []string
		linksCh chan string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Finder{
				url:     tt.fields.url,
				exts:    tt.fields.exts,
				links:   tt.fields.links,
				linksCh: tt.fields.linksCh,
			}
			got, err := f.GetLinks()
			if (err != nil) != tt.wantErr {
				t.Errorf("Finder.GetLinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finder.GetLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinder_GetChan(t *testing.T) {
	type fields struct {
		url     string
		exts    []string
		links   []string
		linksCh chan string
	}
	tests := []struct {
		name    string
		fields  fields
		want    chan string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Finder{
				url:     tt.fields.url,
				exts:    tt.fields.exts,
				links:   tt.fields.links,
				linksCh: tt.fields.linksCh,
			}
			got, err := f.GetChan()
			if (err != nil) != tt.wantErr {
				t.Errorf("Finder.GetChan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finder.GetChan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFinder(t *testing.T) {
	type args struct {
		pageURL string
		exts    []string
	}
	tests := []struct {
		name string
		args args
		want Finder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFinder(tt.args.pageURL, tt.args.exts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}
