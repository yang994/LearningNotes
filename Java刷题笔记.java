import java.util.*;

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
        if (root1==null) return false;  //这里的判断方法需要注意
        if (root1.val==root2.val){
            return (Equal(root1.left,root2.left)&&Equal(root1.right,root2.right));
        }else return false;
    }
}




/*
>>>>>>>>>>>>>>>>>>>>>栈结构的应用>>>>>>>>>>>>>>>>>>>>>>>>>>>>
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））。
>>>>>>>>>>>>>>>>>>>>>栈结构的应用>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//思路：主要看栈相关的操作，和双栈的思路
public class Solution {
    Stack<Integer> s1,s2;
    public Solution(){
        this.s1 = new Stack<Integer>();
        this.s2 = new Stack<Integer>();
    }
    public void push(int node) {
        this.s1.push(node);
        if (this.s2.empty()){       //判断栈空
            this.s2.push(node);
        }else{
            if (this.s2.peek()<=node){
                this.s2.push(s2.peek());
            }else{
                this.s2.push(node);
            }
        }
    }
    public void pop() {
        this.s1.pop();
        this.s2.pop();
    }
    public int top() {
        return this.s1.peek();
    }
    public int min() {
        return this.s2.peek();
    }
}





/*
>>>>>>>>>>>>>>>>>>>>>栈结构的应用>>>>>>>>>>>>>>>>>>>>>>>>>>>>
从上往下打印出二叉树的每个节点，同层节点从左至右打印。
>>>>>>>>>>>>>>>>>>>>>栈结构的应用>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//思路：树结构的顺序遍历，用ArrayList模拟队列操作，注意ArrayLIst的用法
public class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;
    public TreeNode(int val) {
        this.val = val;
    }
}
public class Solution {
    public ArrayList<Integer> PrintFromTopToBottom(TreeNode root) {
        ArrayList<Integer> out = new ArrayList<Integer>();
        if (root==null)return out;
        ArrayList<TreeNode> queue = new ArrayList<TreeNode>();
        queue.add(root);
        for (;queue.size()!=0;){
            TreeNode t=queue.remove(0);     //remove函数应用
            if(t.left!=null) queue.add(t.left); //add函数应用
            if(t.right!=null)queue.add(t.right);
            out.add(t.val);
        }
        return out;
    }
}





/*
>>>>>>>>>>>>>>>>>>>>>判断是否为树结构>>>>>>>>>>>>>>>>>>>>>>>>>>>>
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。
>>>>>>>>>>>>>>>>>>>>>判断是否为树结构>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//思路 递归思想应用，序列中最后一个节点为根节点，利用根节点将序列分为左右子树，递归判断
public class Solution {
    public boolean VerifySquenceOfBST(int [] sequence) {
        //先转化为ArrayList便于操作
        if (sequence.length==0) return false;
        ArrayList<Integer> a=new ArrayList<Integer>();
        for (int i=0;i<sequence.length;i++){
            a.add(sequence[i]);
        }
        return Verify(a);
    }
    public boolean Verify(ArrayList<Integer> sequence){
        if (sequence.size()<=1) return true;
        int root=sequence.get(sequence.size()-1);   //get(i) 取元素方法   size()获取大小方法
        ArrayList<Integer> left = new ArrayList<Integer>();
        ArrayList<Integer> right = new ArrayList<Integer>();
        int i=0;
        for (;i<sequence.size()-1;i++){ //整理出左子树
            if (sequence.get(i)<root){
                left.add(sequence.get(i));
            }else break;
        }
        for (;i<sequence.size()-1;i++){//整理出整理出右子树 当遇到比根小的元素，可以直接返回false了
            if (sequence.get(i)>root){
                right.add(sequence.get(i));
            }else return false;
        }
        return (Verify(left)||Verify(right)); //递归判断
    }
}



/*
>>>>>>>>>>>>>>>>>>>>>回溯法 深度优先遍历>>>>>>>>>>>>>>>>>>>>>>>>>>>>
输入一颗二叉树的根节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。(注意: 在返回值的list中，数组长度大的数组靠前)
>>>>>>>>>>>>>>>>>>>>>回溯法 深度优先遍历>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//思路 深度优先遍历 寻路方法，利用栈记录当前寻路结果
public class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;

    public TreeNode(int val) {
        this.val = val;

    }

}
public class Solution {
    private ArrayList<ArrayList<Integer>> paths = new ArrayList<>();
    private Stack<Integer> path = new Stack();

    public ArrayList<ArrayList<Integer>> FindPath(TreeNode root, int target) {
        if (root == null) {
            return new ArrayList<>();
        }
        find(root, target, 0);
        return paths;
    }

    private void find(TreeNode root, int target, int cur) {
        //返回条件。当满足target并且走到最后，加入路径并返回，没满足target但是走到了最后直接返回
        if (cur == target && root == null) {
            paths.add(new ArrayList(path));
            return;
        }
        if (root == null||cur>target) {
            return;
        }
        
        path.push(root.val);
        find(root.left, target, cur + root.val);//深度优先遍历左子树
        path.pop();//回溯

        if (root.left == null && root.right == null) {    //这里在遍历完左子树后 再检查是否为叶子节点，如果不是，说明右子树还没遍历
            return;
        }

        path.push(root.val);
        find(root.right, target, cur + root.val);//深度优先遍历右子树
        path.pop();//回溯    这里 遍历完右子树就直接返回了，不会重复遍历
    }
}


/*
>>>>>>>>>>>>>>>>>>>>>动态规划，全排列>>>>>>>>>>>>>>>>>>>>>>>>>>>>
输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串abc,则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
>>>>>>>>>>>>>>>>>>>>>动态规划，全排列>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//递归应用，注意字符串 和其他数据结构用法
public class Solution {
    TreeSet<String> paths = new TreeSet<>();
    StringBuilder path = new StringBuilder();
    boolean[] visited;
    public ArrayList<String> Permutation(String str) {
        char[] strs = str.toCharArray();    //字符串和字符数组的转换
        visited = new boolean[strs.length]; //记录每个元素是否访问
        Arrays.sort(strs);      //对数组排序的方法
        ArrayList<String> out = new ArrayList<String>();
        if (strs.length==0) return out;
        order(strs,0);
        out.addAll(paths);  //TreeSet放进ArrayList
        return out;
    }
    public void order(char[] strs,int index){
        if (path.length()==strs.length){
            paths.add(path.toString()); //TreeSet添加元素
            return;
        }
        for(int i=0;i<strs.length;i++){
            if (visited[i]==false){
                visited[i]=true;
                path.append(strs[i]);   //StringBuilder 添加元素
                order(strs,i+1);
                visited[i]=false;
                path.deleteCharAt(path.length()-1); //StringBuilder 删除元素
            }
        }

    }
}



/*
>>>>>>>>>>>>>>>>>>>>>只出现一次的数字>>>>>>>>>>>>>>>>>>>>>>>>>>>>
一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。
>>>>>>>>>>>>>>>>>>>>>只出现一次的数字>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
public class Solution {
    public void FindNumsAppearOnce(int [] array,int num1[] , int num2[]) {
        ArrayList<Integer> arr = new ArrayList<>();
        for (int i=0;i<array.length;i++){
            if (arr.contains(array[i])){
                arr.remove(Integer.valueOf(array[i]));  //ArrayList 移除整形元素的方法
            }else{
                arr.add(array[i]);
            }
        }
        num1[0] = arr.remove(0);
        num2[0] = arr.remove(0);
    }
}

/*
>>>>>>>>>>>>>>>>>>>>>循环左移>>>>>>>>>>>>>>>>>>>>>>>>>>>>
汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，就是用字符串模拟这个指令的运算结果。对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc”。是不是很简单？OK，搞定它！
>>>>>>>>>>>>>>>>>>>>>循环左移>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//注意字符串操作
public class Solution {
    public String LeftRotateString(String str,int n) {
        StringBuffer sb = new StringBuffer();
        if (n>str.length()) return str;
        for (int i=n;i<str.length();i++){   //长度操作length()
            sb.append(str.charAt(i));   //String取元素操作 charAt
        }
        for (int i=0;i<n;i++){
            sb.append(str.charAt(i));//StringBuffer添加元素操作append
        }
        return sb.toString();
    }
}



/*
>>>>>>>>>>>>>>>>>>>>>调用Math包>>>>>>>>>>>>>>>>>>>>>>>>>>>>
求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
>>>>>>>>>>>>>>>>>>>>>调用Math包>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//乘方操作、移位操作

public class Solution {
    public int Sum_Solution(int n) {
        int sum=(int)Math.pow(n, 2)+n;  //pow返回double类型的数
        return sum>>1;
    }
}


/*
>>>>>>>>>>>>>>>>>>>>>位运算>>>>>>>>>>>>>>>>>>>>>>>>>>>>
写一个函数，求两个整数之和，要求在函数体内不得使用+、-、*、/四则运算符号。
>>>>>>>>>>>>>>>>>>>>>位运算>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/
//位运算的思想理解

public class Solution {
    public int Add(int num1,int num2) {
        int a=num1^num2;    //异或，此时没有进位
        int b = num1&num2;  //与 获取进位为1，没进位为0
        int c = b<<1;       //进位符右移一位，准备相加
        if(b!=0){
            a=Add(a,c);     //若有进位需要，则递归加 精辟，直到最后没有进位  精辟啊！
        }
        return a;
    }
}


