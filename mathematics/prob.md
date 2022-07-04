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
    [...] could be the first of the second ...
    For the first player to win, she may either win without a deuce or after a deuce.
    Without a deuce, (player-1, player-2) may score: {(4,2), (4, 1), (4,0)}
    With a deuce, the possible outcomes are: player-1 wins next two points or player-1 wins after another deuce.
    Assuming the notation for the binomial probability as P(n,k,p):
        P(first to deliver wins) = Sum( P(winning on the last point after n points) for n in {3,4,5} )
                                 + Sum( P(deuce) * P(winning after deuce) )
            P(winning after deuce) = P(winning next two points) + P(winning from a subsequent deuce)
            P(winning after deuce) = p^2 + P(2,1,p) * P(winning after deuce)
            P(winning after deuce) = p^2 / (1-P(2,1,p))
        P(first to deliver wins) = Sum( P(n,3,p) * p for n in {3,4,5} )
                                 + P(deuce) * P(winning after deuce)
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