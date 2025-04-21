"""
Create a Python function that accepts a string.
The function should return a string, with each character in the original string doubled.
If you send the function “now” as a parameter, it should return “nnooww,” and if you send “123a!”, it should return “112233aa!!”.
"""

def repeatChar(word: str) -> str:
    return "".join([letter + letter for letter in word])


if __name__ == "__main__":
    print(repeatChar("123a!"))
