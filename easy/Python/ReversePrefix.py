from math import ceil
class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        try:
            index = word.index(ch)
            return word[:index + 1][:: - 1] + word[index + 1:]
        except ValueError:
            return word

s: Solution = Solution()
print(s.reversePrefix('abcdefd', 'd'))

