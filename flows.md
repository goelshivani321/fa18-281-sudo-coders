## UI Flows
1. User logs in-->user GoAPI
2. User checks ride History-->rides goAPI
3. user buys card--> card goAPI
                 --> payment goAPI
4. user gets a top-up --> payment GoAPI
                      -->balance stored in card goAPI
5. user checks balance --> card GoAPI
6. user buys a ride --> check balance --> if balance is there -->balance is deducted from cards db--> success entry made in cards db
                    --> check balance --> balance is not available --> prompt user to make a top up


## Schema Details
1. cards - user id, cards id, expiry date, balance of card

2. payment- card id and payment id ;
         account number, account display number, name on the card,

3. users - name,email id , address, address
