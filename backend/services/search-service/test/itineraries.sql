INSERT INTO itineraries (itineraryid, city, country, title, description, price, languages, tags, events, postid, username)
VALUES
(1, 'New York', 'USA', 'City Highlights', 'Explore iconic landmarks of NYC.', 250.00, ARRAY['English', 'Spanish'], ARRAY['sightseeing', 'food'], ARRAY['event1', 'event2'], 101, 'user1'),
(2, 'Sydney', 'Australia', 'Opera & Beaches', 'Experience Sydney''s opera house and beaches.', 300.00, ARRAY['English'], ARRAY['beach', 'opera'], ARRAY['event3', 'event4'], 102, 'user2'),
(3, 'Tokyo', 'Japan', 'Taste of Tokyo', 'Delve into the best sushi and ramen in Tokyo.', 150.00, ARRAY['Japanese', 'English'], ARRAY['food', 'sushi'], ARRAY['event5', 'event6'], 103, 'user1'),
(4, 'Tokyo', 'Japan', 'Night Owls', 'Enjoy Tokyo''s vibrant nightlife scene.', 180.00, ARRAY['Japanese', 'English'], ARRAY['nightlife', 'club'], ARRAY['event7', 'event8'], 104, 'user1'),
(5, 'Paris', 'France', 'Romantic Escape', 'A romantic journey through Paris.', 500.00, ARRAY['French', 'English'], ARRAY['romantic', 'food'], ARRAY['event9', 'event10'], 105, 'user3'),
(6, 'Los Angeles', 'USA', 'Hollywood & More', 'Discover the culture of Los Angeles.', 220.00, ARRAY['English', 'Spanish'], ARRAY['culture', 'food'], ARRAY['event11', 'event12'], 106, 'user3'),
(7, 'Melbourne', 'Australia', 'Art & Coffee', 'Explore Melbourne''s art galleries and coffee spots.', 200.00, ARRAY['English'], ARRAY['art', 'coffee'], ARRAY['event13', 'event14'], 107, 'user4'),
(8, 'Los Angeles', 'USA', 'Relax by the Ocean', 'Spend a relaxing day at the beaches of LA.', 100.00, ARRAY['English'], ARRAY['beach', 'surfing'], ARRAY['event15', 'event16'], 108, 'user2'),
(9, 'London', 'UK', 'Time Travel', 'Walk through the history of London.', 250.00, ARRAY['English'], ARRAY['history', 'museum'], ARRAY['event17', 'event18'], 109, 'user4'),
(10, 'Vancouver', 'Canada', 'Nature''s Best', 'Embark on an adventurous hiking trip.', 180.00, ARRAY['English'], ARRAY['hiking', 'nature'], ARRAY['event19', 'event20'], 110, 'user2'),
(11, 'Tokyo', 'Japan', 'Shopaholic''s Dream', 'Find the best shopping districts in Tokyo.', 200.00, ARRAY['Japanese', 'English'], ARRAY['shopping', 'fashion'], ARRAY['event21', 'event22'], 111, 'user5'),
(12, 'London', 'UK', 'Sounds of London', 'Join London''s famous music festival.', 150.00, ARRAY['English'], ARRAY['music', 'festival'], ARRAY['event23', 'event24'], 112, 'user5'),
(13, 'Sydney', 'Australia', 'Explore Trails', 'Discover scenic trails around Sydney.', 140.00, ARRAY['English'], ARRAY['hiking', 'adventure'], ARRAY['event25', 'event26'], 113, 'user2'),
(14, 'Tokyo', 'Japan', 'Traditions & More', 'Explore Tokyo''s rich traditions.', 200.00, ARRAY['Japanese', 'English'], ARRAY['culture', 'history'], ARRAY['event27', 'event28'], 114, 'user5'),
(15, 'Paris', 'France', 'Gastronomic Escape', 'Taste the best of French cuisine.', 350.00, ARRAY['French', 'English'], ARRAY['food', 'wine'], ARRAY['event29', 'event30'], 115, 'user3');