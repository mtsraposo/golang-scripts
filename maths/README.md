# Logic

### 1. Price per kg
    price_t = 0.8 * price_sjc
    price_sp = 1.2 * price_sjc
    price_t = 6
    => price_sp = 1.2 * price_t / 0.8 = 9.0

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
    Let the group of the first 17 teams be called G1, and the group of the last 3 teams be called G2.
    The last team in G1 and all of G2 will be downgraded.
    If, at every home-visitor pair of matches within G1, 
    each team wins only once, either as a visitor or as a host,
    all teams in G1 will have 3 * 16 = 48 points. If, in addition, each team in G1 always beats a team in G2,
    all teams in G1 will have a total of 48 (within G1) + 2 * 3 * 3 (G1 x G2) = 66 points.
    G2 teams will have between 0 and 2 * 2 * 3 = 12 points, depending on the outcomes of the matches within G2.
    The last team in G1 will be classified as 17th in the championship, according to any of the resolution criteria,
    and will therefore be downgraded with **66 points**.

# Probability

### 3. Dice
#### a. P(sum is odd)
    P(odd) = {1,3,5} / {1,2,3,4,5,6} = 1/2 = P(even)
    P(sum is odd) = P(only one of the dice is odd) = P(odd) * P(even) + P(even) * P(odd) = 1/2
#### b. P(product is odd)
    P(product is odd) = P(both dice are odd) = P(odd) * P(odd) = 1/4
#### c. P(sum <= 5)
    P(sum <= 5) = 2 * P({1,2}, {1,3}, {1,4}, {2,3}) + P({1,1}, {2,2}) 
    = 2 * 4 * 1 / 36 + 2 * 1 / 36 =   10 / 36 = 5 / 18 ~ 0.278 
### 4. Tennis
#### a. P(deuce)
    As the points are independent, we can say that the outcome of a point X 
    follows a binomial distribution, ie X ~ B(n,p), where n is the number of matches played.
    To get a deuce, 6 points must have been played, with each player winning 3 (k successes), 
    which translates to: f(k,n,p) = f(3,6,p) = P(3,6,p).
    From the binomial distribution: P(k,n,p) = C(n,k) * p^k * (1-p)^(n-k)
    So: P(deuce) = C(6,3) * p^3 * (1-p)^3 = 20 * p^3 * (1-p)^3.
#### b. P(first to deliver wins)
    For the first player to win, she may either win without a deuce or after a deuce.
    Without a deuce, (player-1, player-2) may score: {(4,2), (4, 1), (4,0)}
    With a deuce, the possible outcomes are: player-1 wins next two points or player-1 wins after another deuce.
    Assuming the notation for the binomial probability as P(n,k,p):
        P(first to deliver wins) = Sum( P(winning on the last point after n points) for n in {3,4,5} )
                                 + P(deuce) * P(winning after deuce)
            P(winning after deuce) = P(winning next two points) + P(winning from a subsequent deuce)
            P(winning after deuce) = p^2 + P(2,1,p) * P(winning after deuce)
            P(winning after deuce) = p^2 / (1-P(2,1,p))
        P(first to deliver wins) = Sum( P(n,3,p) * p for n in {3,4,5} )
                                 + 20 * p^3 * (1-p)^3 * p^2 / (1-P(2,1,p))
#### c. script
    cf. tennis.go

### 5. Ballot box
    There are three transitions to consider: 
        (i) when the game is "reset", ie the first ball of a sequence has not been drawn yet, 
        (ii) when the previous ball drawn was blue
        (iii) when the previous ball drawn was red
    For (i), the probability that the last ball will be blue may be defined as,
    assuming "i" red balls and "j" blue balls:
        P(i,j) = P(red) * P(i-1,j | R) + P(blue) * P(i, j-1 | B), 
            where P(i,j | C) denotes the conditional probability that 
            the last ball will be blue given that the previous ball was of color C
        P(i,j) = i / (i+j) * P(i-1,j | R) + j / (i+j) * P(i, j-1 | B)
    The conditional probabilities (ii) and (iii) are calculated similarly:
        P(i,j | B) = i / (i+j) * P(i,j) + j / (i+j) * P(i, j-1 | B)
        P(i,j | R) = i / (i+j) * P(i-1,j | R) + j / (i+j) * P(i, j),
            the P(i, j) terms indicating the cases where 
            a different ball is drawn and the game is reset
    With the boundary conditions:
        P(i,1) = 0, for i in [1, reds]
        P(1,j) = 1, for j in (1, blues]
    Starting from P(1,2) and moving either horizontally or vertically, 
    we can build a matrix of all possible probabilities that the last ball
    will be blue given a number of "i" red balls and "j" blues balls in the box.
    The result will be at position (reds, blues) of the matrix, at which point we may
    stop iterating.
    Please refer to ballot.go for the implementation.