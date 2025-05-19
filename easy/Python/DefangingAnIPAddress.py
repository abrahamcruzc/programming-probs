
class Solution:
    def defangIPaddr(self, address: str) -> str:
        return address.replace('.', '[.]')

address: str = input()
solution: Solution = Solution()
result: str = solution.defangIPaddr(address)
print(result)