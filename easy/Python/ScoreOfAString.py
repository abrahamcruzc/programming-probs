class Solution:
    def scoreOfString(self, s: str) -> int:
        score: int = 0

        for i in range(len(s)-1):
            score += abs(ord(s[i]) - ord(s[i+1]))

        return score

s: str = input()
solution: Solution = Solution()
result: int = solution.scoreOfString(s)
print(result)