"""
Write a function in Python that accepts a credit card number.
It should return a string where all the characters are hidden with an asterisk except the last four.
For example, if the function gets sent “4444444444444444”, then it should return “4444”.
"""

def hide_card_number(num: int) -> str:
    if num > 9999999999999999:
        raise ValueError("Invalid credit card number! please input 16 digit valid number.")
    return "*" * 12 + str(num%10000)


if __name__=="__main__":
    print(hide_card_number(1234567812345678))
