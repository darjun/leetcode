import java.util.ArrayDeque;

public class LongestValidParentheses {
    public static void main(String[] args) 
    {
        Solution s = new Solution();
        System.out.println(s.longestValidParentheses("(())"));
        System.out.println(s.longestValidParentheses(")()())"));
        System.out.println(s.longestValidParentheses("()(())"));
        System.out.println(s.longestValidParentheses(")()())()()("));
        System.out.println(s.longestValidParentheses("(()))())("));
        System.out.println(s.longestValidParentheses(")(()())())(((()))(()()()(()(()(())))(())()((()()(((()())()))(()()())())(())(()(()()()()))(((()())))(((()))))()()())))(()))))())(((()"));
    }
}

class Solution {
    public int longestValidParentheses(String s) {
        ArrayDeque<Integer> stack = new ArrayDeque<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            switch (c) {
            case '(':
                stack.push(-1);
                break;
            case ')':
                if (stack.isEmpty()) {
                    // stack is empty
                    // does nothing
                    break;
                }

                Integer top = stack.peek();
                if (top == -1) {
                    // -1 represents (, match current ), pop it
                    stack.pop();
                    // ()'s length is 2
                    int num = 2;
                    if (!stack.isEmpty()) {
                        // check if last is num?
                        // pattern: `expr`()
                        Integer last = stack.peek();
                        if (last > 0) {
                            num += last;
                            stack.pop();
                        }
                    }
                    stack.push(num);
                } else if (top == -2) {
                    // -2 represents )
                    // does nothing
                } else {
                    if (stack.size() == 1) {
                        stack.push(-2);
                        break;
                    }

                    Integer num = top;                    
                    stack.pop();

                    Integer last = stack.peek();
                    // check pattern: (`expr`)
                    if (last == -1) {
                        num += 2;
                        // pop last
                        stack.pop();

                        if (stack.size() > 0) {
                            // check pattern: `expr`(`expr`)
                            last = stack.peek();
                            if (last > 0) {
                                num += last;
                                stack.pop();
                            }
                        }
                        stack.push(num);
                    } else {
                        // push back top
                        stack.push(top);
                        stack.push(-2);
                    }
                }
                break;
            }
        }
        
        Integer max = 0;
        for (Integer e : stack) {
            if (max < e) {
                max = e;
            }
        }

        return max;
    }
}