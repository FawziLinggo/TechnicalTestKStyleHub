-- Insert into member table
INSERT INTO member (id, username, gender, skintype, skincolor, created_at, updated_at, deleted_at)
VALUES
    (UUID(), 'JohnDoe', 'Male', 'Oily', 'Fair', NOW(), NOW(), NULL),
    (UUID(), 'JaneSmith', 'Female', 'Dry', 'Medium', NOW(), NOW(), NULL),
    (UUID(), 'MikeJohnson', 'Male', 'Combination', 'Dark', NOW(), NOW(), NULL),
    (UUID(), 'EmilyBrown', 'Female', 'Normal', 'Fair', NOW(), NOW(), NULL),
    (UUID(), 'DavidWilson', 'Male', 'Oily', 'Medium', NOW(), NOW(), NULL),
    (UUID(), 'SophiaTaylor', 'Female', 'Dry', 'Dark', NOW(), NOW(), NULL),
    (UUID(), 'DanielAnderson', 'Male', 'Combination', 'Fair', NOW(), NOW(), NULL),
    (UUID(), 'OliviaClark', 'Female', 'Normal', 'Medium', NOW(), NOW(), NULL),
    (UUID(), 'MatthewWhite', 'Male', 'Oily', 'Dark', NOW(), NOW(), NULL),
    (UUID(), 'AvaThompson', 'Female', 'Dry', 'Fair', NOW(), NOW(), NULL),
    (UUID(), 'JamesMoore', 'Male', 'Combination', 'Medium', NOW(), NOW(), NULL),
    (UUID(), 'IsabellaWalker', 'Female', 'Normal', 'Dark', NOW(), NOW(), NULL),
    (UUID(), 'JosephHall', 'Male', 'Oily', 'Fair', NOW(), NOW(), NULL),
    (UUID(), 'MiaYoung', 'Female', 'Dry', 'Medium', NOW(), NOW(), NULL),
    (UUID(), 'ChristopherGreen', 'Male', 'Combination', 'Dark', NOW(), NOW(), NULL),
    (UUID(), 'AbigailHarris', 'Female', 'Normal', 'Fair', NOW(), NOW(), NULL),
    (UUID(), 'AndrewMartin', 'Male', 'Oily', 'Medium', NOW(), NOW(), NULL),
    (UUID(), 'SofiaLewis', 'Female', 'Dry', 'Dark', NOW(), NOW(), NULL),
    (UUID(), 'BenjaminKing', 'Male', 'Combination', 'Fair', NOW(), NOW(), NULL),
    (UUID(), 'CharlotteHall', 'Female', 'Normal', 'Medium', NOW(), NOW(), NULL);

-- Insert into product table
INSERT INTO product (id, name, price, created_at, updated_at, deleted_at)
VALUES
    (UUID(), 'TOVEGAN – Red Remedy Toner (150ml)', 433707, NOW(), NOW(), NULL),
    (UUID(), 'TOVEGAN – Camellia Pure Cushion 21 Light Beige(12ml)', 224331, NOW(), NOW(), NULL),
    (UUID(), 'Theonleaf – Vegan Greenery Repair Cica Cream', 463618, NOW(), NOW(), NULL),
    (UUID(), 'GINGER6 – Active Water Cream 50ml', 299108, NOW(), NOW(), NULL),
    (UUID(), 'SHUADAM – Silky Body Lotion 150ml', 239286, NOW(), NOW(), NULL),
    (UUID(), 'DINTO – Blur Finish Lip Tint – 4color', 224331, NOW(), NOW(), NULL),
    (UUID(), 'TOVEGAN – Orange Oasis Serum(50ml)', 553350, NOW(), NOW(), NULL),
    (UUID(), '76N1 – BE HAPPY Double Action Mist', 373885, NOW(), NOW(), NULL),
    (UUID(), 'NICHE STITCH – Pocket dress perfume 42ml', 164509, NOW(), NOW(), NULL);

-- Insert into product_review table
INSERT INTO product_review (id, product_id, member_id, desc_review, created_at, updated_at, deleted_at)
VALUES
    (UUID(), (SELECT id FROM product WHERE name = 'TOVEGAN – Red Remedy Toner (150ml)'), (SELECT id FROM member WHERE username = 'JohnDoe'), 'Great product!', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = 'TOVEGAN – Camellia Pure Cushion 21 Light Beige(12ml)'), (SELECT id FROM member WHERE username = 'JaneSmith'), 'Average product.', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = 'Theonleaf – Vegan Greenery Repair Cica Cream'), (SELECT id FROM member WHERE username = 'MikeJohnson'), 'Disappointed with the quality.', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = 'GINGER6 – Active Water Cream 50ml'), (SELECT id FROM member WHERE username = 'EmilyBrown'), 'Excellent product!', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = 'SHUADAM – Silky Body Lotion 150ml'), (SELECT id FROM member WHERE username = 'DavidWilson'), 'Moisturizes well.', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = 'DINTO – Blur Finish Lip Tint – 4color'), (SELECT id FROM member WHERE username = 'SophiaTaylor'), 'Long-lasting color.', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = 'TOVEGAN – Orange Oasis Serum(50ml)'), (SELECT id FROM member WHERE username = 'DanielAnderson'), 'Works wonders!', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = '76N1 – BE HAPPY Double Action Mist'), (SELECT id FROM member WHERE username = 'OliviaClark'), 'Refreshing mist.', NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product WHERE name = 'NICHE STITCH – Pocket dress perfume 42ml'), (SELECT id FROM member WHERE username = 'MatthewWhite'), 'Lovely scent.', NOW(), NOW(), NULL),
(UUID(), (SELECT id FROM product WHERE name = 'TOVEGAN – Red Remedy Toner (150ml)'), (SELECT id FROM member WHERE username = 'MikeJohnson'), 'Amazing toner!', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'TOVEGAN – Red Remedy Toner (150ml)'), (SELECT id FROM member WHERE username = 'EmilyBrown'), 'Refreshes my skin.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'TOVEGAN – Camellia Pure Cushion 21 Light Beige(12ml)'), (SELECT id FROM member WHERE username = 'DavidWilson'), 'Provides good coverage.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'Theonleaf – Vegan Greenery Repair Cica Cream'), (SELECT id FROM member WHERE username = 'SophiaTaylor'), 'Helped with my skin irritation.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'GINGER6 – Active Water Cream 50ml'), (SELECT id FROM member WHERE username = 'DanielAnderson'), 'Hydrates my skin.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'SHUADAM – Silky Body Lotion 150ml'), (SELECT id FROM member WHERE username = 'OliviaClark'), 'Smooth texture.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'DINTO – Blur Finish Lip Tint – 4color'), (SELECT id FROM member WHERE username = 'MatthewWhite'), 'Gorgeous shades.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'TOVEGAN – Orange Oasis Serum(50ml)'), (SELECT id FROM member WHERE username = 'JohnDoe'), 'Improved my skin texture.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = '76N1 – BE HAPPY Double Action Mist'), (SELECT id FROM member WHERE username = 'JaneSmith'), 'Feels refreshing on my face.', NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product WHERE name = 'NICHE STITCH – Pocket dress perfume 42ml'), (SELECT id FROM member WHERE username = 'MikeJohnson'), 'Unique fragrance.', NOW(), NOW(), NULL);


-- Insert into like_review table
INSERT INTO like_review (id, product_review_id, member_id, created_at, updated_at, deleted_at)
VALUES
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Great product!'), (SELECT id FROM member WHERE username = 'JohnDoe'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Average product.'), (SELECT id FROM member WHERE username = 'JaneSmith'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Disappointed with the quality.'), (SELECT id FROM member WHERE username = 'MikeJohnson'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Excellent product!'), (SELECT id FROM member WHERE username = 'EmilyBrown'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Moisturizes well.'), (SELECT id FROM member WHERE username = 'DavidWilson'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Long-lasting color.'), (SELECT id FROM member WHERE username = 'SophiaTaylor'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Works wonders!'), (SELECT id FROM member WHERE username = 'DanielAnderson'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Refreshing mist.'), (SELECT id FROM member WHERE username = 'OliviaClark'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Lovely scent.'), (SELECT id FROM member WHERE username = 'MatthewWhite'), NOW(), NOW(), NULL),
    (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Great product!'), (SELECT id FROM member WHERE username = 'JohnDoe'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Average product.'), (SELECT id FROM member WHERE username = 'JaneSmith'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Disappointed with the quality.'), (SELECT id FROM member WHERE username = 'MikeJohnson'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Excellent product!'), (SELECT id FROM member WHERE username = 'EmilyBrown'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Moisturizes well.'), (SELECT id FROM member WHERE username = 'DavidWilson'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Long-lasting color.'), (SELECT id FROM member WHERE username = 'SophiaTaylor'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Works wonders!'), (SELECT id FROM member WHERE username = 'DanielAnderson'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Refreshing mist.'), (SELECT id FROM member WHERE username = 'OliviaClark'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Lovely scent.'), (SELECT id FROM member WHERE username = 'MatthewWhite'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Amazing toner!'), (SELECT id FROM member WHERE username = 'JohnDoe'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Refreshes my skin.'), (SELECT id FROM member WHERE username = 'JaneSmith'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Provides good coverage.'), (SELECT id FROM member WHERE username = 'MikeJohnson'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Helped with my skin irritation.'), (SELECT id FROM member WHERE username = 'EmilyBrown'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Hydrates my skin.'), (SELECT id FROM member WHERE username = 'DavidWilson'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Smooth texture.'), (SELECT id FROM member WHERE username = 'SophiaTaylor'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Gorgeous shades.'), (SELECT id FROM member WHERE username = 'DanielAnderson'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Improved my skin texture.'), (SELECT id FROM member WHERE username = 'OliviaClark'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Feels refreshing on my face.'), (SELECT id FROM member WHERE username = 'MatthewWhite'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Unique fragrance.'), (SELECT id FROM member WHERE username = 'JohnDoe'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Great product!'), (SELECT id FROM member WHERE username = 'JaneSmith'), NOW(), NOW(), NULL),
  (UUID(), (SELECT id FROM product_review WHERE desc_review = 'Moisturizes well.'), (SELECT id FROM member WHERE username = 'MikeJohnson'), NOW(), NOW(), NULL);
