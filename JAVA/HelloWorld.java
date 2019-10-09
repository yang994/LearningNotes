import java.util.*;

public class HelloWorld{
    public static void main(String[] args) {
        ArrayList<Integer> arr = new ArrayList<>();
        arr.add(2);
        arr.add(3);
    }
}
class Solution1 {
    public ArrayList<ArrayList<Integer>> FindContinuousSequence(int sum) {
        int s=0;
        ArrayList<ArrayList<Integer>> outList = new ArrayList<ArrayList<Integer>>();
        ArrayList<Integer> workList = new ArrayList<>();
        for (int i=1;i<sum;i++){
            if(s==sum){
                outList.add(copyList(workList));
                s+=i;
                workList.add(i);
                for(;s>=sum;){
                    s-=workList.remove(0);
                }

            }
            else if(s<sum){
                s+=i;
                workList.add(i);
            }
            else if(s>sum){
                s+=i;
                workList.add(i);
                for(;s>=sum;){
                    s-=workList.remove(0);
                }
            }
            
        }
        return outList;
    }
    public ArrayList<Integer> copyList(ArrayList<Integer> a){
        ArrayList<Integer> out=new ArrayList<>();
        for(int i=0;i<a.size();i++){
            out.add(a.get(i));
        }
        return out;
    }
}

class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;

    public TreeNode(int val) {
        this.val = val;
    }
}
class Solution2 {
    String Serialize(TreeNode root) {
        StringBuffer sb=new StringBuffer();
        //层序遍历
        if (root==null) return sb.toString();
        ArrayList<TreeNode> arr = new ArrayList<>();
        arr.add(root);
        TreeNode p;
        TreeNode voidNode = new TreeNode(-1);
        int levelCount=1,thisCount=0;
        for(;arr.size()!=0;){
            int voidCount = 0;
            p=arr.remove(0);
            thisCount++;
            if (p==voidNode){
                sb.append("#!");
                voidCount++;
            }else{
                sb.append(Integer.toString(p.val)+"!");
                if(p.left!=null) arr.add(p.left);
                else arr.add(voidNode);
                if(p.right!=null) arr.add(p.right);
                else arr.add(voidNode);
            }
            if(thisCount==levelCount){
                if (voidCount == thisCount) break;
                levelCount*=2;
                thisCount = 0;
            }
        }
        return sb.toString();
    }
    TreeNode Deserialize(String str) {
        String[] strList = str.split("!");
        return construct(strList,0);
    }
    TreeNode construct(String[] strs,int index){
        if (index<strs.length){
            if (strs[index].equals("#")) return null;
            int thisVal = Integer.parseInt(strs[index]);
            System.out.println(index);
            TreeNode root = new TreeNode(thisVal);
            root.left = construct(strs,(index*2));
            root.right = construct(strs,(index*2+1));
            return root;
        }
        return null;
        
    }
}