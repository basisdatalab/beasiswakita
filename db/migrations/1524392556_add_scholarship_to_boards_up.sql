ALTER TABLE student_boards
ADD scholarship_id int(11)
AFTER user_id,
ADD CONSTRAINT FK_scholarship_board
FOREIGN KEY (scholarship_id) REFERENCES scholarships (id);