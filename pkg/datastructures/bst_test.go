package datastructures

import (
	"reflect"
	"testing"
)

func TestNewBSTNode(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		args args
		want *BSTNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBSTNode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBSTNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBST(t *testing.T) {
	tests := []struct {
		name string
		want *BST
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBST(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Insert(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			b.Insert(tt.args.data)
		})
	}
}

func TestBST_insert(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	type args struct {
		newNode *BSTNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			b.insert(tt.args.newNode)
		})
	}
}

func TestBST_Search(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BSTNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := b.Search(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BST.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Delete(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	type args struct {
		node *BSTNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			b.Delete(tt.args.node)
		})
	}
}

func TestBST_BreadthFirst(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := b.BreadthFirst(); got != tt.want {
				t.Errorf("BST.BreadthFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_DepthFirst(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := b.DepthFirst(); got != tt.want {
				t.Errorf("BST.DepthFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_InOrder(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	tests := []struct {
		name     string
		fields   fields
		initData []int
		want     string
	}{
		{
			name:     "evenly built",
			initData: []int{32, 16, 48, 8, 24, 40, 56, 4, 12, 20, 28, 36, 44, 52, 60},
			want:     "4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48, 52, 56, 60",
		},
		{
			name:     "all lesser",
			initData: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want:     "0, 1, 2, 3, 4, 5, 6, 7, 8, 9",
		},
		{
			name:     "all greater",
			initData: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:     "0, 1, 2, 3, 4, 5, 6, 7, 8, 9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			for _, num := range tt.initData {
				b.Insert(num)
			}
			if got := b.InOrder(); got != tt.want {
				t.Errorf("BST.InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_PreOrder(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := b.PreOrder(); got != tt.want {
				t.Errorf("BST.PreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_PostOrder(t *testing.T) {
	type fields struct {
		root *BSTNode
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
				size: tt.fields.size,
			}
			if got := b.PostOrder(); got != tt.want {
				t.Errorf("BST.PostOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
