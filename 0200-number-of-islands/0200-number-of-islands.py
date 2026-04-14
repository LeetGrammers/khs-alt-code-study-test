class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        answer = 0
        m = len(grid)
        n = len(grid[0])
        visited = [[False] * n for _ in range(m)]
        dr, dc = [0, 1, 0, -1], [1, 0, -1, 0]
        def dfs(r, c, m, n):
            visited[r][c] = True
            for k in range(4):
                nr, nc = r + dr[k], c + dc[k]
                if nr < 0 or nc < 0 or nr >= m or nc >= n or visited[nr][nc] or grid[nr][nc] == "0":
                    continue
                dfs(nr, nc, m, n)
        for r in range(m):
            for c in range(n):
                if grid[r][c] == "1" and not visited[r][c]:
                    answer += 1
                    dfs(r, c, m, n)
        return answer            
