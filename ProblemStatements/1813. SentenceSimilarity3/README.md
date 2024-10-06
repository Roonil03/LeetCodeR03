# 1813. Sentence Similarity III
## Question Level: Medium
### Description:
You are given two strings sentence1 and sentence2, each representing a sentence composed of words. A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of only uppercase and lowercase English characters.

Two sentences s1 and s2 are considered similar if it is possible to insert an arbitrary sentence (possibly empty) inside one of these sentences such that the two sentences become equal. Note that the inserted sentence must be separated from existing words by spaces.

For example,<br>
s1 = "Hello Jane" and s2 = "Hello my name is Jane" can be made equal by inserting "my name is" between "Hello" and "Jane" in s1.
s1 = "Frog cool" and s2 = "Frogs are cool" are not similar, since although there is a sentence "s are" inserted into s1, it is not separated from "Frog" by a space.<br>
Given two sentences sentence1 and sentence2, return true if sentence1 and sentence2 are similar. Otherwise, return false.

### Example:
<b>Example 1:</b><br>
Input: sentence1 = "My name is Haley", sentence2 = "My Haley"<br>
Output: true<br>
Explanation:<br>
sentence2 can be turned to sentence1 by inserting "name is" between "My" and "Haley".<br>

<b>Example 2:</b><br>
Input: sentence1 = "of", sentence2 = "A lot of words"<br>
Output: false<br>
Explanation:<br>
No single sentence can be inserted inside one of the sentences to make it equal to the other.<br>

<b>Example 3:</b><br>
Input: sentence1 = "Eating right now", sentence2 = "Eating"<br>
Output: true<br>
Explanation:<br>
sentence2 can be turned to sentence1 by inserting "right now" at the end of the sentence.<br>

### Constraints:
- 1 <= sentence1.length, sentence2.length <= 100
- sentence1 and sentence2 consist of lowercase and uppercase English letters and spaces.
- The words in sentence1 and sentence2 are separated by a single space.