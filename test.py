"""
Write a function in Python that accepts a list of any length that contains a mix of non-negative integers and strings.
The function should return a list with only the integers in the original list in the same order.
"""
from typing import List, Union

def onlyInts(nums: List[Union[int, str]]) -> List[int]:
    return [ele for ele in nums if isinstance(ele, int) and not isinstance(ele, bool)]


print(onlyInts([1, 2, 4, "five", True, 6, "seven"]))


"""
Create a Python function that accepts a string. The function should return a string,
with each character in the original string doubled. If you send the function “now” as a parameter,
it should return “nnooww,” and if you send “123a!”, it should return “112233aa!!”.
"""

def doubleLetters(word: str) -> str:
    return "".join([letter * 2 for letter in word])

print(doubleLetters("123a!"))


"""
For this challenge, create a Python function that accepts a string.
The function should return a string, with each lowercase character in the original string returned as uppercase characters.
If you send the function “goodbye” as a parameter, it should return “GOODBYE”.
"""

print("abc".upper())


"""
Write a function in Python that accepts a string.
The function should return a string and add “.” in between each letter.
For example, if you send the function “skills” as a parameter, it should return “s.k.i.l.l.s”.
"""

# brute force
word = "skills"
new_word = ""
for i in range(len(word)):
    new_word += word[i]
    print(i)
    if i < len(word) - 1:
        new_word += "."
print(new_word)

# optimized
print(".".join(word))


"""
Given an integer N, return the number of digits in N.
"""
num = 12345
count = 0

while num > 0:
    count += 1
    num = num//10

print(count)

"""
Given an integer N return the reverse of the given number
"""

n = 4123
rev_n = 0
while n > 0:
    digit = n % 10
    rev_n = rev_n * 10 + digit
    n //= 10
print(rev_n)


"""
Given an integer N, return true it is an Armstrong number otherwise return false.
An Amrstrong number is a number that is equal to the sum of its own digits each raised to the power of the number of digits.
"""

n = 153
string_n = str(n)
len_n = len(string_n)
sum_ = 0
for num in string_n:
    sum_ += int(num) ** len_n
print(sum_)
