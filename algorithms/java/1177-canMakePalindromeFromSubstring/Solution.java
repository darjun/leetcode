import java.util.List;
import java.util.ArrayList;

class Solution {
  public List<Boolean> canMakePaliQueries(String s, int[][] queries) {
    int[][] sum = new int[s.length()+1][];
    sum[0] = new int[26];
    for (int i = 0; i < s.length(); i++) {
      sum[i + 1] = sum[i].clone();
      ++sum[i + 1][s.charAt(i) - 'a'];
    }

    List<Boolean> can = new ArrayList<>(queries.length);
    for (int[] q: queries) {
      int diff = 0;
      for (int j = 0; j < 26; j++) {
        diff += (sum[q[1] + 1][j] - sum[q[0]][j]) & 1;
      }
      System.out.println(diff);
      can.add(diff/2<=q[2]);
    }
    return can;
  }
}