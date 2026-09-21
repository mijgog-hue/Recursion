package main
 
import "testing"
 
func treeEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && treeEqual(a.Left, b.Left) && treeEqual(a.Right, b.Right)
}
 
func TestInvertTree(t *testing.T) {
	s := Solution{}
 
	tests := []struct {
		name  string
		input *TreeNode
		want  *TreeNode
	}{
		{
			name:  "пустое дерево",
			input: nil,
			want:  nil,
		},
		{
			name:  "один узел",
			input: &TreeNode{Val: 1},
			want:  &TreeNode{Val: 1},
		},
		{
			name: "два уровня",
			input: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 7},
			},
			want: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		{
			name: "три уровня, классический кейс",
			input: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7,
					Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: 7,
					Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
				Right: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
			},
		},
	}
 
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.InvertTree(tt.input)
			if !treeEqual(got, tt.want) {
				t.Errorf("InvertTree() failed")
			}
		})
	}
}
