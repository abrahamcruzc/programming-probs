def truncateSentence(s: str, k: int) -> str:
    words: list[str] = s.split(" ")
    selected = words[:k]
    return ' '.join(selected)