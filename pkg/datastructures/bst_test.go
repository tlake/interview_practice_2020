package datastructures

import (
	"reflect"
	"testing"
)

func buildBST(initData []int) *BST {
	b := &BST{}
	for _, num := range initData {
		b.Insert(num)
	}
	return b
}

func TestNewBSTNode(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		args args
		want *BSTNode
	}{
		{
			name: "new node is initialized",
			args: args{data: 3},
			want: &BSTNode{
				data:         3,
				parent:       nil,
				lesserChild:  nil,
				greaterChild: nil,
			},
		},
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
		{
			name: "new bst is initialized",
			want: &BST{
				root: nil,
				size: 0,
			},
		},
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
	type args struct {
		data int
	}
	tests := []struct {
		name            string
		initData        []int
		args            args
		wantBFTraversal string
		wantSize        int
	}{
		{
			name:            "successfully creates node, then inserts it",
			initData:        []int{3, 1, 5, 0, 4, 6},
			args:            args{data: 2},
			wantBFTraversal: "3, 1, 5, 0, 2, 4, 6",
			wantSize:        7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
			b.Insert(tt.args.data)
			if got := b.BreadthFirst(); got != tt.wantBFTraversal {
				t.Errorf("BST.Insert() traversal = %v, want = %v", got, tt.wantBFTraversal)
			}
			if b.size != tt.wantSize {
				t.Errorf("BST.Insert() size = %v, want = %v", b.size, tt.wantSize)
			}
		})
	}
}

func TestBST_insert(t *testing.T) {
	type args struct {
		newNode *BSTNode
	}
	tests := []struct {
		name            string
		initData        []int
		args            args
		wantBFTraversal string
		wantSize        int
	}{
		{
			name:            "insert into empty tree",
			initData:        []int{},
			args:            args{newNode: &BSTNode{data: 3}},
			wantBFTraversal: "3",
			wantSize:        1,
		},
		{
			name:            "insert balanced",
			initData:        []int{3, 1, 5, 0, 4, 6},
			args:            args{newNode: &BSTNode{data: 2}},
			wantBFTraversal: "3, 1, 5, 0, 2, 4, 6",
			wantSize:        7,
		},
		{
			name:            "insert all lesser",
			initData:        []int{5, 4, 3, 2, 1},
			args:            args{newNode: &BSTNode{data: 0}},
			wantBFTraversal: "5, 4, 3, 2, 1, 0",
			wantSize:        6,
		},
		{
			name:            "insert all greater",
			initData:        []int{0, 1, 2, 3, 4},
			args:            args{newNode: &BSTNode{data: 5}},
			wantBFTraversal: "0, 1, 2, 3, 4, 5",
			wantSize:        6,
		},
		{
			name:            "insert existing",
			initData:        []int{3, 1, 5, 0, 2, 4, 6},
			args:            args{newNode: &BSTNode{data: 2}},
			wantBFTraversal: "3, 1, 5, 0, 2, 4, 6",
			wantSize:        7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
			b.insert(tt.args.newNode)
			if got := b.BreadthFirst(); got != tt.wantBFTraversal {
				t.Errorf("BST.insert() traversal = %v, want = %v", got, tt.wantBFTraversal)
			}
			if b.size != tt.wantSize {
				t.Errorf("BST.insert() size = %v, want = %v", b.size, tt.wantSize)
			}
		})
	}
}

func TestBST_Search(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name     string
		initData []int
		args     args
		want     *BSTNode
	}{
		{
			name:     "not found",
			initData: []int{5, 3, 4, 1, 2, 0, 7, 6, 9, 8},
			args:     args{data: 11},
			want:     nil,
		},
		{
			name:     "find at root",
			initData: []int{5, 3, 4, 1, 2, 0, 7, 6, 9, 8},
			args:     args{data: 5},
			want:     &BSTNode{data: 5},
		},
		{
			name:     "find at left",
			initData: []int{5, 3, 4, 1, 2, 0, 7, 6, 9, 8},
			args:     args{data: 1},
			want:     &BSTNode{data: 1},
		},
		{
			name:     "find at right",
			initData: []int{5, 3, 4, 1, 2, 0, 7, 6, 9, 8},
			args:     args{data: 9},
			want:     &BSTNode{data: 9},
		},
		{
			name:     "find in middle",
			initData: []int{5, 3, 4, 1, 2, 0, 7, 6, 9, 8},
			args:     args{data: 3},
			want:     &BSTNode{data: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
			got := b.Search(tt.args.data)
			if tt.want == nil {
				if !(reflect.DeepEqual(got, tt.want)) {
					t.Errorf("BST.Search() = %v, want %v", got, tt.want)
				}
			} else if got.data != tt.args.data {
				t.Errorf("BST.Search() = %v, want %v", got.data, tt.args.data)
			}
		})
	}
}

func TestBST_Delete(t *testing.T) {
	type args struct {
		node *BSTNode
	}
	tests := []struct {
		name                string
		initData            []int
		targetValue         int
		expectedBFTraversal string
		wantSize            int
	}{
		/*
			        7
			   3         11
			 1   5     9    13
			0 2 4 6   8       14

			insert:
			7, 3, 11, 1, 5, 9, 13, 0, 2, 4, 6, 8, 14
		*/
		{
			name:                "delete node with no children",
			initData:            []int{7, 3, 1, 5, 0, 2, 4, 6, 11, 9, 13, 8, 14},
			targetValue:         2,
			expectedBFTraversal: "7, 3, 11, 1, 5, 9, 13, 0, 4, 6, 8, 14",
			wantSize:            12,
		},
		{
			name:                "delete lesser node with only lesser child",
			initData:            []int{7, 3, 1, 5, 0, 2, 4, 6, 11, 9, 13, 8, 14},
			targetValue:         9,
			expectedBFTraversal: "7, 3, 11, 1, 5, 8, 13, 0, 2, 4, 6, 14",
			wantSize:            12,
		},
		{
			name:                "delete greater node with only greater child",
			initData:            []int{7, 3, 1, 5, 0, 2, 4, 6, 11, 9, 13, 8, 14},
			targetValue:         13,
			expectedBFTraversal: "7, 3, 11, 1, 5, 9, 14, 0, 2, 4, 6, 8",
			wantSize:            12,
		},
		{
			name:                "delete node with two children",
			initData:            []int{7, 3, 1, 5, 0, 2, 4, 6, 11, 9, 13, 8, 14},
			targetValue:         3,
			expectedBFTraversal: "7, 4, 11, 1, 5, 9, 13, 0, 2, 6, 8, 14",
			wantSize:            12,
		},
		{
			name:                "delete root node",
			initData:            []int{7, 3, 1, 5, 0, 2, 4, 6, 11, 9, 13, 8, 14},
			targetValue:         7,
			expectedBFTraversal: "8, 3, 11, 1, 5, 9, 13, 0, 2, 4, 6, 14",
			wantSize:            12,
		},
		/*
			        7
			   3         11
			 1   5     9    13
			0 2 4 6     10 12

			insert:
			7, 3, 11, 1, 5, 9, 13, 0, 2, 4, 6, 10, 12
		*/
		{
			name:                "delete greater node with only lesser child",
			initData:            []int{7, 3, 11, 1, 5, 9, 13, 0, 2, 4, 6, 10, 12},
			targetValue:         13,
			expectedBFTraversal: "7, 3, 11, 1, 5, 9, 12, 0, 2, 4, 6, 10",
			wantSize:            12,
		},
		{
			name:                "delete lesser node with only greater child",
			initData:            []int{7, 3, 11, 1, 5, 9, 13, 0, 2, 4, 6, 10, 12},
			targetValue:         9,
			expectedBFTraversal: "7, 3, 11, 1, 5, 10, 13, 0, 2, 4, 6, 12",
			wantSize:            12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
			n := b.Search(tt.targetValue)
			b.Delete(n)

			if got := b.BreadthFirst(); got != tt.expectedBFTraversal {
				t.Errorf("BST.Delete()got %v want = %v", got, tt.expectedBFTraversal)
			}
			if b.size != tt.wantSize {
				t.Errorf("BST.Delete() size = %v, want = %v", b.size, tt.wantSize)
			}
		})
	}
}

func TestBST_BreadthFirst(t *testing.T) {
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
			name:     "empty",
			initData: []int{},
			want:     "",
		},
		{
			name:     "balanced",
			initData: []int{32, 16, 48, 8, 24, 40, 56, 4, 12, 20, 28, 36, 44, 52, 60},
			want:     "32, 16, 48, 8, 24, 40, 56, 4, 12, 20, 28, 36, 44, 52, 60",
		},
		{
			name:     "all lesser",
			initData: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want:     "9, 8, 7, 6, 5, 4, 3, 2, 1, 0",
		},
		{
			name:     "all greater",
			initData: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:     "0, 1, 2, 3, 4, 5, 6, 7, 8, 9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
			if got := b.BreadthFirst(); got != tt.want {
				t.Errorf("BST.BreadthFirst() = %v, want %v", got, tt.want)
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
			name:     "empty",
			initData: []int{},
			want:     "",
		},
		{
			name:     "balanced",
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
			b := buildBST(tt.initData)
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
		name     string
		fields   fields
		initData []int
		want     string
	}{
		{
			name:     "empty",
			initData: []int{},
			want:     "",
		},
		{
			name:     "balanced",
			initData: []int{32, 16, 48, 8, 24, 40, 56, 4, 12, 20, 28, 36, 44, 52, 60},
			want:     "32, 16, 8, 4, 12, 24, 20, 28, 48, 40, 36, 44, 56, 52, 60",
		},
		{
			name:     "all lesser",
			initData: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want:     "9, 8, 7, 6, 5, 4, 3, 2, 1, 0",
		},
		{
			name:     "all greater",
			initData: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:     "0, 1, 2, 3, 4, 5, 6, 7, 8, 9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
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
		name     string
		fields   fields
		initData []int
		want     string
	}{
		{
			name:     "empty",
			initData: []int{},
			want:     "",
		},
		{
			name:     "balanced",
			initData: []int{32, 16, 48, 8, 24, 40, 56, 4, 12, 20, 28, 36, 44, 52, 60},
			want:     "4, 12, 8, 20, 28, 24, 16, 36, 44, 40, 52, 60, 56, 48, 32",
		},
		{
			name:     "all lesser",
			initData: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want:     "0, 1, 2, 3, 4, 5, 6, 7, 8, 9",
		},
		{
			name:     "all greater",
			initData: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:     "9, 8, 7, 6, 5, 4, 3, 2, 1, 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
			if got := b.PostOrder(); got != tt.want {
				t.Errorf("BST.PostOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_getLeast(t *testing.T) {
	tests := []struct {
		name             string
		initData         []int
		targetStartValue int
		want             *BSTNode
	}{
		{
			name:             "given root",
			initData:         []int{3, 1, 5, 0, 2, 4, 6},
			targetStartValue: 3,
			want:             &BSTNode{data: 0},
		},
		{
			name:             "given empty",
			initData:         []int{3, 1, 5, 0, 2, 4, 6},
			targetStartValue: -1,
			want:             nil,
		},
		{
			name:             "given mid",
			initData:         []int{3, 1, 5, 0, 2, 4, 6},
			targetStartValue: 5,
			want:             &BSTNode{data: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buildBST(tt.initData)
			targetNode := b.Search(tt.targetStartValue)
			got := b.getLeast(targetNode)
			if tt.want == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("BST.getLeast() = %v, want = %v", got, tt.want)
				}
			} else if got.data != tt.want.data {
				t.Errorf("BST.getLeast() = %v, want %v", got.data, tt.want.data)
			}
		})
	}
}
