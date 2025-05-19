class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        jewels_set = set(jewels)

        return sum(1 for stone in stones if stone in jewels_set )

jewels: str = input("Jewels: ")
stones: str = input("Stones: ")
solution: Solution = Solution()
result: int = solution.numJewelsInStones(jewels, stones)
print(result)