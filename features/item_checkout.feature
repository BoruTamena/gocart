Feature: Checkout 

    Checkout


    Scenario: Check Out Cart Item 

    When I check out the items i have in the my shopping cart <user_id>
    Then The system should checkout the item and clear the cart 
    Examples:
        | user_id | 
        | 2  |