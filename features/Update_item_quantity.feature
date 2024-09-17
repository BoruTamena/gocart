Feature: Update 

    update feature of a shopping cart is responsible for increasing and 
    decreasing the item quantity that pre-exist in the cart already


Scenario: Increase Quantity
When I update item in cart with <session_id>,  <product_id>  by increasing quantity, 
Then  The system should increase the item quantity 
Then the response should  return affected row <row>.
Examples:
    | session_id | product_id | row |
    | 3  | 1 | 1 |

Scenario: Decrease Quantity
When I update item in cart with <session_id>  and <product_id> by decreasing  quantity 
Then The system should decrease the item quantity 
Then the response should  return affected row <row>
Examples:
    | session_id | product_id |row|
    | 3  | 1 | 1|






