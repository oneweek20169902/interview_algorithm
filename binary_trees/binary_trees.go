package binary_trees

type Node struct {
	Value     int
	LeftNode  *Node
	RightNode *Node
}

var rootNode = &Node{
	Value: 1,
	LeftNode: &Node{
		Value: 2,
		LeftNode: &Node{
			Value: 4,
			RightNode: &Node{
				Value:     6,
				LeftNode:  &Node{Value: 7},
				RightNode: &Node{Value: 8},
			},
		},
	},
	RightNode: &Node{
		Value:     3,
		RightNode: &Node{Value: 5},
	},
}

/*
二叉树打印顺序：
1、先序：头、左、右
2、中序：左、头、右
3、后序：左、右、头

*/

/*
非递归遍历：
先序：x
	1、从栈中弹出一个节点cur
	2、打印处理cur
	3、先右再左（如果有）
Stack<Node>stack 二 new Stack<Node>();
	stack.add(head);
while (Istack. isEmpty()） {
	head = stack.pop();
	System. out.print (head .value +""）;
	if (head. right != nu11) R
	stack.push (head. right)
}
stack• push(head.right);
if (head. left != null) { stack.push (head. left);


后序：
准备两个栈：stack1和stack2
以头左右的顺序写入stack1，逻辑为：
1、弹cur
2、cur放入stack2
3、先左再右
public static void posOrderUnRecur1(Node head)
System.out.print("pos-order: ");
if (head != nul1){
Stack<Node> s1 = new Stack<Node>();
Stack<Node> new Stack<Node>();
s1.push(head);
while (!s1. isEmpty()){
head = s1.pop();
s2.push (head);
if (head. left l= nul1) { s1.push(head. left);
if (head. right != nul1) {
s1.push(head.right);
}


System. out.print(s2.pop().value +"");
while (!s2. isEmpty()){
System.out. printin();

中序：
每棵子树左边界(顺着左节点查找)进栈，依次弹出的过程中对弹出的右节点重复
*/

/*
搜索二叉树：左树都比右树小
*/

var preValue int

func IsBST(head *Node) bool {
	if head == nil {
		return false
	}
	isLeftBst := IsBST(head.LeftNode)
	if !isLeftBst {
		return false
	}
	if head.Value <= preValue {
		return false
	} else {
		preValue = head.Value
	}
	return IsBST(head)
}

/*
如何判断完全二叉树：
1、如果右节点没有左节点
2、如果左右节点不双全的情况，后面的遇到的节点都必须是叶节点
*/

/*
如何判断满二叉树：
1、求最大深度为L
2、求节点数N
3、如果是满二叉树必然满足N=2^L-1
*/

/*
平衡树是左节点和右节点的的高度差为1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor 二叉树最近的公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}
