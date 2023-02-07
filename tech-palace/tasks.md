1. Create the welcome message

For most customers who scan their loyalty cards, the store owner wants to see <code>Welcome to the Tech Palace,</code> followed by the name of the customer in capital letters on the display.

Implement the function <code>WelcomeMessage</code> that accepts the name of the customer as a <code>string</code> argument and returns the desired message as a <code>string</code>.


2. Add a fancy border

For loyal customers that buy a lot at the store, the owner wants the welcome display to be more fancy by adding a line of stars before and after the welcome message. They are not sure yet how many stars should be in the lines so they want that to be configurable.

Write a function <code>AddBorder</code> that accepts a welcome message (a <code>string</code>) and the number of stars per line (type <code>int</code>) as arguments. It should return a <code>string</code> that consists of 3 lines, a line with the desired number of stars, then the welcome message as it was passed in, then another line of stars.



3. Clean up old marketing messages

Before installing this new display, the store had a similar display that could only show non-configurable, static lines. The owner would like to reuse some of the old marketing messages on the new display. However, the data already includes a star border and some unfortunate whitespaces. Your task is to clean up the messages so they can be re-used.

Implement a function <code>CleanUpMessage</code> that accepts the old marketing message as a string. The function should first remove all stars from the text and afterwards remove the leading and trailing whitespaces from the remaining text. The function should then return the cleaned up message.