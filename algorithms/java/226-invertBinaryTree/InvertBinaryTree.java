import java.util.ArrayDeque;

public class InvertBinaryTree {
    public static void main(String[] args) {

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
    public TreeNode invertTree(TreeNode root) {
        if (root == null) {
            return root;
        }

        ArrayDeque<TreeNode> l = new ArrayDeque<>();
        l.push(root);
        while (!l.isEmpty()) {
            TreeNode n = l.pop();

            TreeNode tmp = n.left;
            n.left = n.right;
            n.right = tmp;

            if (n.left != null) {
                l.push(n.left);
            }

            if (n.right != null) {
                l.push(n.right);
            }
        }

        return root;
    }
}