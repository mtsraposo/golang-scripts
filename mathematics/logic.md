# Logic

### 1. Price per kg
    price_t = 0.8 * price_sjc
    price_sp = 1.2 * price_sjc
    price_t = 6
    => price_sp = 1.2 * price_t / 0.8 = 9

### 2. Soccer championship
#### a. Number of matches per team
    Each team plays the remaining 19 twice (one at home and one as a visitor), 
    so 38 matches are played by each team  
#### b. Total number of matches
    Two matches between every team, so the total number of matches equals: 
    2 * C(20, 2) = 2 * 20! / [(20-2)!2!) = 20 * 19 = 380
#### c. Min score to become champion
    If the result of all matches is a draw, one of the teams will be picked randomly 
    and win with 38 points (1 point per match played).
#### d. Max score to be downgraded
    If at every home-visitor pair or matches between teams, 
    each team wins only once, either as a visitor or as a host, 
    all teams will have 3 * 19 = 57 points, and 4 will be picked randomly to be downgraded.