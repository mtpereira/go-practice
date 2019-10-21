# Stock prices

## Summary

URL: https://www.interviewcake.com/

### Problem type

First, I wanna know how much money I could have made yesterday if I'd been trading Apple stocks all day.

So I grabbed Apple's stock prices from yesterday and put them in a list called stock_prices, where:

- The indices are the time (in minutes) past trade opening time, which was 9:30am local time.
- The values are the price (in US dollars) of one share of Apple stock at that time.

So if the stock cost \$500 at 10:30am, that means `stock_prices[60] = 500`.

### Question(s) asked

Write an efficient function that takes `stock_prices` and returns the best profit I could have made from one purchase and one sale of one share of Apple stock yesterday.
No "shorting" — you need to buy before you can sell. Also, you can't buy and sell in the same time step—at least 1 minute has to pass.
