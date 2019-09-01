/*
>>>>>>>>>>>>>>>>树子结构判断>>>>>>>>>>>>>>>>>>>>>>>>>
输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）
>>>>>>>>>>>>>>>>树子结构判断>>>>>>>>>>>>>>>>>>>>>>>>>
*/
public class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;

    public TreeNode(int val) {
        this.val = val;

    }

}
//思路：遍历root1所有节点，与root2根节点比较，相等，则继续比较左右子节点，不等，返回false
public class Solution {
    public boolean HasSubtree(TreeNode root1,TreeNode root2) {
        //前序遍历root1所有节点
        if(root2==null) return false;
        if (root1!=null){
            if (Equal(root1,root2)){
                return true;
            }
            return (Equal(root1.left,root2)||Equal(root1.right,root2));
        }else return false;
        
    }
    public boolean Equal(TreeNode root1,TreeNode root2){
        if (root2==null) return true;
        if (root1==null) return false;
        if (root1.val==root2.val){
            return (Equal(root1.left,root2.left)&&Equal(root1.right,root2.right));
        }else return false;
    }
}



