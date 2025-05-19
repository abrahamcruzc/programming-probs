class Solution:
    def findWordsContaining(self, words: list[str], x: str) -> list[int]:

        return [i for i in range(len(words)) if x in words[i]]

words: list[str] = ["abc","bcd","aaaa","cbc"]
x: str = "a"
solution: Solution = Solution()
result: list[int] = solution.findWordsContaining(words, x)
print(result)
