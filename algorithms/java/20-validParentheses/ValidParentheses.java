import java.util.ArrayList;
import java.util.Stack;
import java.util.List;

public class ValidParentheses {
    public static void main(String[] args) {
        // Solution s = new Solution();
        // System.out.println(s.isValid("()"));
        // System.out.println(s.isValid("()[]{}"));
        // System.out.println(s.isValid("(]"));
        // System.out.println(s.isValid("([)]"));
        // System.out.println(s.isValid("{[]}"));

        Solution s1 = new Solution();
        System.out.println(s1.isValid("()"));
        System.out.println(s1.isValid("()[]{}"));
        System.out.println(s1.isValid("(]"));
        System.out.println(s1.isValid("([)]"));
        System.out.println(s1.isValid("{[]}"));

        System.out.println(Character.codePointAt("(", 0));
        System.out.println(Character.codePointAt(")", 0));
        System.out.println(Character.codePointAt("{", 0));
        System.out.println(Character.codePointAt("}", 0));
        System.out.println(Character.codePointAt("[", 0));
        System.out.println(Character.codePointAt("]", 0));
        System.out.println(Character.toChars(123));
        System.out.println(Character.toChars(124));
        System.out.println(Character.toChars(125));
    }
}

class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<Character>();
        char c1, c2;
        for (int i = 0; i < s.length(); i++) {
            c1 = s.charAt(i);
            switch (c1) {
            case '(':
            case '{':
            case '[':
                stack.push(c1);
                break;
            case ')':
            case '}':
            case ']':
                if (stack.empty()) {
                    return false;
                }
                c2 = stack.pop();
                if ((c1 == ')' && c2 != '(') || (c1 == '}' && c2 != '{') || (c1 == ']' && c2 != '[')) {
                    return false;
                }
                break;
            }
        }
        return stack.empty();
    }
}

class SolutionOpt1 {
    public boolean isValid(String s) {
        List<Character> list = new ArrayList<Character>();
        char c1, c2;
        for (int i = 0; i < s.length(); i++) {
            c1 = s.charAt(i);
            switch (c1) {
            case '(':
            case '{':
            case '[':
                list.add(c1);
                break;
            case ')':
            case '}':
            case ']':
                if (list.size() == 0) {
                    return false;
                }
                c2 = list.remove(list.size() - 1);
                if ((c1 == ')' && c2 != '(') || (c1 == '}' && c2 != '{') || (c1 == ']' && c2 != '[')) {
                    return false;
                }
                break;
            }
        }
        return list.size() == 0;
    }
}