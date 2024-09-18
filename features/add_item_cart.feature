Feature: Add

    add feature insert new product item in to cart 


    Scenario: Add New Item 
    # Given user has a session id 1,
    Then I add item a new product with a <product_id>,<session_id> and <quantity>,
    Then the system should add  product item into cart 
    Then the system should return "new item add seccussfully".
    Examples:
        | product_id | session_id | quantity |
        | 2 |  3| 1  |
    
        

   

