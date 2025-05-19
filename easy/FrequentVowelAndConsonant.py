from collections import Counter

class Solution:
    def maxFreqSum(self, s: str) -> int:
        vowels: set[str] = set('aeiou')
        consonants: set[str] = set('bcdfghjklmnpqrstvwxyz') 

        freq = Counter(s)

        max_vowel: int = max((freq[c] for c in freq if c in vowels), default=0)
        max_consonant: int = max((freq[c] for c in freq if c in consonants), default=0)

        return max_vowel + max_consonant
    
solution: Solution = Solution()
s: str = "aeiaeia"
output: int = solution.maxFreqSum(s)
print(output)