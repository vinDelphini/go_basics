"""
Write a function in Python that accepts a string.
The function should return a string and add “.” in between each letter.
For example, if you send the function “skills” as a parameter, it should return “s.k.i.l.l.s”.
"""

def addDots(word: str) -> str:
    length = len(word)
    return ".".join(word)


if __name__ == "__main__":
    print(addDots("skills"))
