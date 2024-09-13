Feature: Add

    add feature insert new product item in to cart 


    Scenario: Add New Item 
    # Given user has a session id 1,
    Then I add item a new product with a product_id,session_id and quantity 2,3,1,
    Then the system should add a new product into cart and return "new item add seccussfully".
    Examples:
        | product_id | session_id | quantity |
        | 2  | 3 | 1  |
        | 2  |  | 1  |
        

    # Scenario: Add Existing item 
    # Given user has a session id 1,
    # When I add existing item with product_id,session_id and quantity 2,1,1,
    # Then the system should increase a item quantity "new item add seccussfully ",




