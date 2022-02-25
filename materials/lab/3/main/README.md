Mariah Moore
Cyber Security
Lab 3b

Changed to allow user input to go to the next page of the query.
This was done in main.go with a for loop that will run until "yes" is not typed.
Added Timestamp to the host search method and then is printed out.

How to run:
go build main.go
SHODAN_API_KEY= ./main SEARCH
Example: SHODAN_API_KEY=zQYRQgl1nJrOcZjmykeHPDHuu619iCE5 ./main city:Highlands
Type yes to continue to the first page
To quit out type anything that is not yes