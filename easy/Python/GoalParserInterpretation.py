class Solution:
    def interpret(self, command: str) -> str:
        return command.replace("()", "o").replace("(al)", "al")

solution: Solution = Solution()
command: str = "(al)G(al)()()G"
result: str = solution.interpret(command)
print(result)