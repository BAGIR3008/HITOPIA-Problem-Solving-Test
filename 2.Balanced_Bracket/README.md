# Balanced Bracket

This Go implementation verifies whether a provided string of brackets is balanced. The `balanced_bracket` function checks if brackets are balanced by confirming every opening bracket has a matching closing bracket and are correctly nested.

## Analysis of Complexity

The time complexity of the function is 𝑂(𝑛), with 𝑛 being the length of the input string. Every character in the string is handled only once.

In the worst-case scenario, where all characters are opening brackets and pushed onto the stack, the space complexity is 𝑂(𝑛).