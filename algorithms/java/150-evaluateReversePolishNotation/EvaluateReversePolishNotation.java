import java.util.ArrayDeque;

public class Trouble {
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.evalRPN(new String[] { "2", "1", "+", "3", "*" }));
        System.out.println(s.evalRPN(new String[] { "4", "13", "5", "/", "+" }));
        System.out.println(s.evalRPN(new String[]{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}));
        System.out.println(6 / -11);
        System.out.println(s.evalRPN(new String[]{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}));
    }
}

class Solution {
    public int evalRPN(String[] tokens) {
        // store operand
        ArrayDeque<Integer> stack = new ArrayDeque<>();
        Integer num1 = 0;
        Integer num2 =0;
        for (int i = 0; i < tokens.length; i++) {
            switch (tokens[i]) {
            case "+":
                num1 = stack.pop();
                num2 = stack.pop();
                stack.push(num2 + num1);
                break;
            case "-":
                num1 = stack.pop();
                num2 = stack.pop();
                stack.push(num2 - num1);
                break;
            case "*":
                num1 = stack.pop();
                num2 = stack.pop();
                stack.push(num2 * num1);
                break;
            case "/":
                num1 = stack.pop();
                num2 = stack.pop();
                stack.push(num2 / num1);
                break;
            default:
                // number
                stack.push(Integer.valueOf(tokens[i]));
                break;
            }
        }
        return stack.pop();
    }
}
