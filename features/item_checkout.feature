Feature: Checkout 

    Checkout


    Scenario: Check Out Cart Item 

    When I check out the items I had in the my shopping cart <user_id>,
    Then the system should return "your order is created successfully".
    
    Examples:
        | user_id |
        | 1  |