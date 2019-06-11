import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashSet;

class Solution {
    public int numIslands(char[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) {
            return 0;
        }

        int count = 0;
        HashSet<Long> visited = new HashSet<>();
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                long key = ((long) i) << 32 + j;
                if (!visited.contains(key) && grid[i][j] == 1) {
                    count++;
                    markSurroundings(grid, i, j, visited);
                }
            }
        }
        return count;
    }

    private void markSurroundings(char[][] grid, int i, int j, HashSet<Long> visited) {
        ArrayDeque<Long> surroundings = new ArrayDeque<>();
        long key = ((long) i) << 32 + j;
        surroundings.push(key);

        int[][] directions = {
            {-1,-1},
            {-1,0},
            {-1, 1},
            {0, -1},
            {0, 1},
            {1, -1},
            {1, 0},
            {1, 1},
        };

        while (surroundings.isEmpty()) {
            key = surroundings.pop();
            i = (int) (key >> 32);
            j = (int) key;

            for (int k = 0; k < 8; k++) {
                int newI = i + directions[k][0];
                int newJ = j + directions[k][1];

                if (newI < 0 || newJ < 0 || newI >= grid.length || newJ >= grid[0].length) {
                    continue;
                }

                key = ((long) newI) << 32 + newJ;
                if (!visited.contains(key)) {
                    surroundings.push(key);
                    visited.add(key);
                }
            }
        }
    }
}

public class NumberOfIslands {
    public static void main(String[] args) {
        {
            char[][] grid = { { 1, 1, 1, 1, 0 }, { 1, 1, 0, 1, 0 }, { 1, 1, 0, 0, 0 }, { 0, 0, 0, 0, 0 }, };
            System.out.println(new Solution().numIslands(grid));
        }
        
        {
            char[][] grid = {
                {1,1,0,0,0},
                {1,1,0,0,0},
                {0,0,1,0,0},
                {0,0,0,1,1},
            };
            System.out.println(new Solution().numIslands(grid));
        }
    }
}