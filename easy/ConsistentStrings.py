class Solution:
    def countConsistentStrings(self, allowed: str, words: list[str]) -> int:
        allowed_set: set[str] = set(allowed)
        consistent_count: int = 0

        for word in words:
            if all(char in allowed_set for char in word):
                consistent_count += 1

        return consistent_count

solution: Solution = Solution()
allowed: str = "ab"
words: list[str] = ["ad", "bd", "aaab", "baa", "badab"]
result: int = solution.countConsistentStrings(allowed, words)
print(result) # 2
