class Solution:
    def finalValueAfterOperations(self, operations: list[str]) -> int:
        x: int = 0
        
        for operation in operations:
            if "++" in operation:
                x += 1
            elif "--" in operation:
                x -= 1

        return x

operations: list[str] = ["--X", "X++", "X++"]
solution: Solution = Solution()
result: int = solution.finalValueAfterOperations(operations)
print(result)