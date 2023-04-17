# bookshop-go

A simple book store API in need of input validation/sanitization.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:

- [Javascript](https://github.com/andey-robins/bookshop-js)
- [Rust](https://github.com/andey-robins/bookshop-rs)

## Versioning

`bookshop-go` is buit with:

- go version go1.19.3 darwin/arm64

## Usage

Start the api using `go run main.go`.

I recommend using [`httpie`](https://httpie.io) for testing of HTTP endpoints on the terminal. Tutorials are available elsewhere online, and you're free to use whatever tools you deem appropriate for testing your code.

## Analysis of Existing Code
There are currently two main issues with the bookstore application. First, you can edit the database at rest. This can lead to issues with reading customers shipping information and potentially exfiltrating it. Moreover, since account balances are also stored in the database, this information is vulnerable. The security fix includes encrypting the database. The second issue is that technically all the APIs are public. This means that someone could be hitting all of the APIs even though customers are only intended to use some. The security fix here includes setting up a private API key for the store owners to use. A couple of other issues concern a lack of database validation. First, you can add a new book with the same name and author as one that already exists in the database but with a lower price. This is concerning because the query will return the most recently added price for a given author and book. Therefore two identical books could exist in the database with different prices. A similar issue exists with allowing the same bookId and customerId to exist in the database. This causes problems when you check the order status. These issues should be resolved with database rules. Furthermore, if a book and author combination does not exist in the database, the price is returned as $0 instead of alerting the user it does not exist. This could easily be remedied with an error check.

## Changes
- `IsPOShipped` was querying the Books table instead of PurchaseOrders
- Database Error:
    - Need to wrap scan in a for loop when multiple results could match
- Added input validation using the `go-playground/validator` package
- Logged errors with the `log` package