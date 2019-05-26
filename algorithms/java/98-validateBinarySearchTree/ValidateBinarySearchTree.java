import java.util.ArrayDeque;
import java.util.HashSet;

public class ValidateBinarySearchTree {
    public static void main(String[] args) {
        {
            TreeNode n1 = new TreeNode(2);
            TreeNode n2 = new TreeNode(1);
            TreeNode n3 = new TreeNode(3);
            n1.left = n2;
            n1.right = n3;
            System.out.println(new Solution().isValidBST(n1));
        }

        {
            TreeNode n1 = new TreeNode(5);
            TreeNode n2 = new TreeNode(1);
            TreeNode n3 = new TreeNode(4);
            TreeNode n4 = new TreeNode(3);
            TreeNode n5 = new TreeNode(6);
            n1.left = n2;
            n1.right = n3;
            n3.left = n4;
            n3.right = n5;
            System.out.println(new Solution().isValidBST(n1));
        }
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
    public boolean isValidBST(TreeNode root) {
        if (root == null) {
            return true;
        }

        ArrayDeque<TreeNode> ad = new ArrayDeque<>();
        HashSet<TreeNode> visitedNodes = new HashSet<>();
        int prevVal = 0;
        boolean first = true;
        boolean nodeFromStack = false;

        TreeNode currentNode = root;
        // traverse in mid-order
        while (currentNode != null) {
            // find the left-est node
            if (!nodeFromStack && currentNode.left != null) {
                ad.push(currentNode);
                currentNode = currentNode.left;
                continue;
            }

            // visit current node
            if (first) {
                first = false;
            } else {
                if (prevVal >= currentNode.val) {
                    return false;
                }
            }
            prevVal = currentNode.val;

            if (currentNode.right != null) {
                currentNode = currentNode.right;
                nodeFromStack = false;
            } else {
                if (ad.isEmpty()) {
                    // no node at all
                    currentNode = null;
                } else {
                    currentNode = ad.pop();
                    nodeFromStack = true;
                }
            }
        }

        return true;        
    }
}