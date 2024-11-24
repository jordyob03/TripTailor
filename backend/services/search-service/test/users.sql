INSERT INTO users (userid, username, email, password, dateofbirth, name, country, languages, tags, boards)
VALUES
(6, 'user1', 'user1@example.com', 'password1', '1990-01-01', 'John Doe', 'USA', ARRAY['English', 'Spanish'], ARRAY['travel', 'food'], ARRAY[101, 102]),
(7, 'user2', 'user2@example.com', 'password2', '1992-02-15', 'Jane Smith', 'Canada', ARRAY['English', 'French'], ARRAY['hiking', 'adventure'], ARRAY[103, 104]),
(8, 'user3', 'user3@example.com', 'password3', '1995-03-22', 'Alice Brown', 'UK', ARRAY['English'], ARRAY['art', 'culture'], ARRAY[105, 106]),
(4, 'user4', 'user4@example.com', 'password4', '1988-04-10', 'Bob Johnson', 'Australia', ARRAY['English'], ARRAY['music', 'nature'], ARRAY[107, 108]),
(5, 'user5', 'user5@example.com', 'password5', '1985-05-05', 'Charlie White', 'Japan', ARRAY['Japanese', 'English'], ARRAY['nightlife', 'food'], ARRAY[109, 110]);