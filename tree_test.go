package treenode

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_bfs(t *testing.T) {
	type args struct {
		q NodeQueue
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1-5 tree",
			args{q: NodeQueue{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
				},
			}},
			[]string{"1", "2", "3", "null", "null", "4", "5"},
		},
		{
			"empty tree",
			args{q: NodeQueue{nil}},
			nil,
		},
		{
			"unbalance tree",
			args{q: NodeQueue{
				&TreeNode{
					1,
					nil,
					&TreeNode{
						2,
						nil,
						&TreeNode{
							3,
							nil,
							nil,
						},
					},
				},
			}},
			[]string{"1", "null", "2", "null", "3"},
		},
		{
			"single node",
			args{q: NodeQueue{&TreeNode{
				1,
				nil,
				nil,
			}}},
			[]string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfs(&tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_String(t1 *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Test 1",
			fields{
				1,
				nil,
				&TreeNode{
					2,
					nil,
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
			},
			"[1,null,2,null,3]",
		},
		{
			"Test 2",
			fields{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
			},
			"[1,2,3,null,null,4,5]",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bfsBuild(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	want := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	err := bfsBuild(&NodeQueue{root}, []string{"2", "3", "null", "null", "4", "5"})
	require.NoError(t, err)
	assert.Equal(t, want.String(), root.String())

	root = &TreeNode{1, nil, nil}
	err = bfsBuild(&NodeQueue{root}, []string{"2", "3", "null", "null", "spam", "5"})
	require.Error(t, err)
	assert.EqualError(t, err, "strconv.Atoi: parsing \"spam\": invalid syntax")

	err = bfsBuild(nil, []string{"2", "3", "null", "null", "spam", "5"})
	assert.NoError(t, err)
}

func TestNewTreeNode(t *testing.T) {
	data := "[1,2,3,null,null,4,5]"
	root, err := NewTreeNode(data)
	require.NoError(t, err)
	assert.Equal(t, data, root.String())
}

func BenchmarkEmptyTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewTreeNode("[]")
	}
}

func Benchmark1_5Tree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewTreeNode("[1,2,3,null,null,4,5]")
	}
}
