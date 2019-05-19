import java.util.ArrayDeque;

public class PathSum {
public static void main(String[] args) {
        {
            TreeNode n1 = new TreeNode(5);
            TreeNode n2 = new TreeNode(4);
            TreeNode n3 = new TreeNode(8);
            TreeNode n4 = new TreeNode(11);
            TreeNode n5 = new TreeNode(7);
            TreeNode n6 = new TreeNode(2);
            TreeNode n7 = new TreeNode(13);
            TreeNode n8 = new TreeNode(4);
            TreeNode n9 = new TreeNode(1);

            n1.left = n2;
            n1.right = n3;
            n2.left = n4;
            n4.left = n5;
            n4.right = n6;
            n3.left = n7;
            n3.right = n8;
            n8.right = n9;
            System.out.println(new Solution().hasPathSum(n1, 22));
        }
    
        
}
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

class Solution {
    public boolean hasPathSum(TreeNode root, int sum) {
        if (root == null) {
            return false;
        }

        ArrayDeque<TreeNode> nodeDeque = new ArrayDeque<>();
        ArrayDeque<Integer> sumDeque = new ArrayDeque<>();

        nodeDeque.push(root);
        sumDeque.push(root.val);

        while (!nodeDeque.isEmpty()) {
            TreeNode node = nodeDeque.pop();
            int currentSum = sumDeque.pop();

            if (node.left == null && node.right == null) {
                if (currentSum == sum) {
                    return true;
                }
            }

            if (node.left != null) {
                nodeDeque.push(node.left);
                sumDeque.push(currentSum + node.left.val);
            }

            if (node.right != null) {
                nodeDeque.push(node.right);
                sumDeque.push(currentSum + node.right.val);
            }
        }

        return false;
    }
}