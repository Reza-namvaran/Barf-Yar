INSERT INTO activities (message_id, title) VALUES (1, 'Activity 1');
INSERT INTO activities (message_id, title) VALUES (2, 'Activity 2');
INSERT INTO activities (message_id, title) VALUES (3, 'Activity 3');
INSERT INTO activities (message_id, title) VALUES (4, 'Activity 4');
INSERT INTO activities (message_id, title) VALUES (5, 'Activity 5');

INSERT INTO activity_prompts (activity_id, prompt_message_id) VALUES (1, 1);
INSERT INTO activity_prompts (activity_id, prompt_message_id) VALUES (2, 2);

INSERT INTO activity_supporters (activity_id, user_id) VALUES (1, 1001);
INSERT INTO activity_supporters (activity_id, user_id) VALUES (1, 1002);
INSERT INTO activity_supporters (activity_id, user_id) VALUES (2, 1003);
