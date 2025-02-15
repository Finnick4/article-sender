# Goals
## Reading
- Get all .md files from a given directory
- Read the directory location from a .env file
- if a valid directory is included in the args, use this instead and add a disclaimer
- If the directory is invalid, promt the user for a new directory

## Usage
- First: Select a file
- Once selected: 
  - show table of content
  - Warning if there is no date attached (use current / set date)
  - Prompt the user for what to do next (view, send, reselect)
- Previewing a markdown file
- Load the file only into memory, when it is needed

## Sending the file
- Cooldown between sending each post request
- slicing each message after 1999 chars
  - slice at the latest space
- Allow only the whole article to be send
- formatting the message
  - table of content
  - title
  - date
  - subtitle
  - formating from HEAD.md file
- read the destination from a .env or ask the user (to then add it to the .env)

## .env
- directory
- destination url

# Future releases
- Add support for multiple destinations
- option for multiple destinations
- add the option for diffrent and custom formats 
  - "header" markdown file in the directory
  - e.g. "article___3.md" has the header file "article.md"
- option to generate a .tex and thus a pdf from the article
- support for polls
