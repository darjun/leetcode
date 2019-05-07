import java.util.ArrayDeque;

public class MaximumDepthOfBinaryTree {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(3);
        TreeNode left = new TreeNode(9);
        TreeNode right = new TreeNode(20);
        root.left = left;
        root.right = right;

        right.left = new TreeNode(15);
        right.right = new TreeNode(7);

        System.out.println(new Solution().maxDepth(root));
    }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

class Solution {
    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int depth = 0;
        ArrayDeque<TreeNode> l = new ArrayDeque<>();
        l.push(root);
        while (true) {
            depth += 1;
            
            ArrayDeque<TreeNode> l2 = new ArrayDeque<>();
            for (TreeNode n : l) {
                if (n.left != null) {
                    l2.push(n.left);
                }

                if (n.right != null) {
                    l2.push(n.right);
                }
            }

            if (l2.isEmpty()) {
                break;
            }

            l = l2;
        }

        return depth;
    }
}