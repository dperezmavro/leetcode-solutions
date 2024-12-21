from typing import List

class Solution:
    def __init__(self):
        self.visited = set()
        self.grid = List[List[int]]

    def numIslands(self, grid: List[List[str]]) -> int:
        m = len(grid)
        n = len(grid[0])
        answer = 0

        self.grid = grid
        for i, row in enumerate(self.grid):
            for j, pos in enumerate(row):
                if pos == "0" or (i, j) in self.visited:
                    continue
                self.exploreNeighbours(m, n, i, j)
                answer += 1

        return answer

    def exploreNeighbours(
        self,
        m: int,
        n: int,
        i: int,
        j: int,
    ):
        to_explore = []
        self.visited.add((i, j))
        if i - 1 >= 0 and self.grid[i - 1][j] == "1" :
            to_explore.append((i - 1, j))

        if i + 1 < m and self.grid[i + 1][j] == "1":
            to_explore.append((i + 1, j))

        if j + 1 < n and self.grid[i][j + 1] == "1":
            to_explore.append((i, j + 1))

        if j - 1 >= 0 and self.grid[i][j - 1] == "1":
            to_explore.append((i, j - 1))

        for p in to_explore:
            if (p[0], p[1]) not in self.visited:
                self.exploreNeighbours(m, n, p[0], p[1])


if __name__ == "__main__":
    s = Solution()

    g = [
        ["1", "1", "1", "1", "0"],
        ["1", "1", "0", "1", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "0", "0", "0"],
    ]
    print(s.numIslands(g))

    s = Solution()
    g = [
        ["1", "1", "0", "0", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "1", "0", "0"],
        ["0", "0", "0", "1", "1"],
    ]
    print(s.numIslands(g))
