Feature: Remove 

    remove feature is responsible for removing each item from 
    shopping cart 


    Scenario:Remove Item 
    When The user <user_id> remove item with <product_id> from the cart .
    Then the system should remove the item for the cart 
    Then the system should  return "item removed successfully".
    Examples:
        | user_id | product_id|
        |  3  | 2 |  