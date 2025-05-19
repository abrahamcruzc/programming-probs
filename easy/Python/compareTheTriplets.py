from typing import List

def compare_triplets(a: List[int], b: List[int]) -> List[int]:
    alice: int = 0
    bob: int = 0

    if len(a) != len(b):
        return []

    for i in range(len(a)):
        if a[i] > b[i]:
            alice += 1
        elif b[i] > a[i]:
            bob += 1

    return [alice, bob]

a: List[int] = [5, 6, 7]
b: List[int] = [3, 6, 10]
result: List[int] = compare_triplets(a, b)
print(result)

