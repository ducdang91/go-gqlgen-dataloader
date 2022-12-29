https://david-yappeter.medium.com/the-importance-of-dataloader-in-graphql-go-4d5214869b20

mutation{
  transaction{
    create(input: {
      user_id: 1
      transaction_details:[
        {
          name: "Golang Course"
          price: 10000
          discount: 5000
        },
        {
          name: "Laravel Course"
          price: 50000
        }
      ]
    }){
      id
      created_at
      user_id
      summary{
        total_price
        total_discount
        transaction_details{
          id
          name
          description
          price
          discount
          transaction_id
        }
      }
    }
  }
}

query{
  users{
    id
    name
    age
    transactions{
      id
      created_at
      user_id
      summary{
        total_price
        total_discount
        transaction_details{
          id
          name
          description
          price
          discount
          transaction_id
        }
      }
    }
  }
}

go run github.com/vektah/dataloaden TransactionByIDLoader int 