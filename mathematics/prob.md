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
    cf. prob.go

### 5. Ballot box
