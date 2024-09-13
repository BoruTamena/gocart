Feature: Update 

    update feature of a shopping cart is responsible for increasing and 
    decreasing the item quantity that pre-exist in the cart already


Scenario: Increase Quantity
When I update item in cart with <session_id>,  <product_id>  by increasing quantity, 
Then  The system should increase the item quantity and update cumulative item.
Examples:
    | session_id | product_id |
    | 3  | 1 | 
# Scenario: Decrease Quantity
# When I update item in cart with 1 product_id by decreasing  quantity 
# Then The system should decrease the item quantity and update cumulative item






