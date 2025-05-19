class Solution:
    def convertDateToBinary(self, date: str) -> str:
        year_str, month_str, day_str = date.split('-')

        year_int: int = int(year_str)
        month_int: int = int(month_str)
        day_int: int = int(day_str)

        year_bin: str = bin(year_int)[2:]
        month_bin: str = bin(month_int)[2:]
        day_bin: str = bin(day_int)[2:]

        result: str = f"{year_bin}-{month_bin}-{day_bin}"
        return result

solution: Solution = Solution()
date: str = "2080-02-29"
result: str = solution.convertDateToBinary(date)
print(result)