import java.util.*;

public class HelloWorld{
    public static void main(String[] args) {
        Scanner s=new Scanner(System.in);
        int a=s.nextInt();
        System.out.println(a);
    }
}
class Solution {
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