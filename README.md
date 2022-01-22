# evermos-assesment-be
Evermos Assesment Backend

Task 1 : Online Store

Question :
1. Describe what you think happened that caused those bad reviews during our 12.12 event and why it happened. Put this in a section in your README.md file.
2. Based on your analysis, propose a solution that will prevent the incidents from occurring again. Put this in a section in your README.md file.
3. Based on your proposed solution, build a Proof of Concept that demonstrates technically how your solution will work.

Answer :
1. in my opinion, it happens at checkout with the same time. So there was a stock error, because there was 1 available stock, while more than 1 checkout at the same time
2. in my analysis, when you checkout, you must also check the available stock, if you checkout with empty stock, the checkout must be aborted
3.  a. when doing checkout use database transaction,
    b. Enter the required data at checkout
    c. Validate the stock with the checkout that has been entered, if it exceeds the available stock limit, cancel the checkout by rollback database
    d. return the checkout cancellation data to the client
    

Command Test Order:
- cd test/
- exec command " go test -v "


Question :
Aspects of the PoC that we will evaluate also include, but are not limited to:
  1. Database schema and entity design.
  2. API endpoints design.
  3. Logging and error handling.


Answer :
1. desain database : https://i.ibb.co/CbkRwgn/Evermos-Assessment.png
2. Swagger API : https://app.swaggerhub.com/apis/wawat7/evermos-assessment_api/1.0.0
3. 

